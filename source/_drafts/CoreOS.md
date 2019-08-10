

# CoreOS

> CoreOS is a minimal Linux operation system built to run docker and rkt containers
>
> CoreOS = etcd + fleetd + systemd  + docker + cloud-config



![](/home/amas/2018-11-08-154640_1001x612_scrot.png)



- 没有安装包管理器
- 占用更少的内存
- 双分区升级策略
- docker和rkt容器作为CoreOS的应用



![](/home/amas/2018-11-07-231601_464x371_scrot.png)





CoreOS被设计为支持集群

![](/home/amas/2018-11-07-231844_498x315_scrot.png)



使用etcd或是fleet将多个CoreOS链接起来就构成了集群.



## 安装

```
$ sudo pacman -s vagrant 
$ git clone https://github.com/coreos/coreos-vagrant/
$ cd coreos-vagrant
$ mv user-data.sample user-data
```

```

#cloud-config

coreos:
  etcd2:
    name: core-01
    initial-advertise-peer-urls: http://$private_ipv4:2380
    listen-peer-urls: http://$private_ipv4:2380,http://$private_ipv4:7001
    initial-cluster-token: core-01_etcd
    initial-cluster: core-01=http://$private_ipv4:2380
    initial-cluster-state: new
    advertise-client-urls: http://$public_ipv4:2379,http://$public_ipv4:4001
    listen-client-urls: http://0.0.0.0:2379,http://0.0.0.0:4001
  fleet:
    public-ip: $public_ipv4
  flannel:
    interface: $public_ipv4
  units:
    - name: etcd2.service
      command: start
    - name: fleet.service
      command: start
    - name: flanneld.service
      drop-ins:
      - name: 50-network-config.conf
        content: |
          [Service]
          ExecStartPre=/usr/bin/etcdctl set /coreos.com/network/config '{ "Network": "10.1.0.0/16" }'
      command: start
    - name: docker-tcp.socket
      command: start
      enable: true
      content: |
        [Unit]
        Description=Docker Socket for the API

        [Socket]
        ListenStream=2375
        Service=docker.service
        BindIPv6Only=both

        [Install]
        WantedBy=sockets.target
```

- 参考:  https://github.com/rimusz/coreos-essentials-book/blob/master/Chapter1/user-data.



```zsh
$ vagrant up 
$ vagrant ssh 
Last login: Wed Nov  7 16:08:37 UTC 2018 from 10.0.2.2 on ssh
Container Linux by CoreOS alpha (1953.0.0)
core@core-01 ~ $ 
```





## etcd

> etcd isn’t designed for large-object storage or massive performance; its primary
> purpose is to make the state of the cluster singular



## 配置要求

 - CPU
    - 2-4 core: 一般使用
    - 8-16:

- Memory
  - 8G:
  - 16-64G: 上千watcher, 上百万的key
- Disk
  -  50 sequential IOPS (7200RPM Disk)
  -  50 sequential IOPS
  - 很多云存储厂商给的是concurrent IOPS, 因此还需要10x, 具体的IOPS可以用diskbench或fio来实地测量



> 参考: https://coreos.com/etcd/docs/latest/op-guide/hardware.html





## CoreOS工作流程

 	1. 建立应用层的Docker或rkt镜像, 以及如何启动镜像中服务的Systemd Units
 	2. 通过Fleet决定在哪些机器上运行这些容器, SystemdUnits通过etcd同步到容器上.

![CoreOS WorkFLow](/home/amas/2018-11-08-171611_890x856_scrot.png)



Nginx Systemd Template Units:



nginx@.service:

```ini
[Unit]
Description=My Nginx Server - %i
Requires=docker.service
After=docker.service

[Service]
ExecStartPre=-/usr/bin/docker kill nginx-%i
ExecStartPre=-/usr/bin/docker rm nginx-%i
ExecStartPre=/usr/bin/docker pull my/nginx:latest
ExecStart=/usr/bin/docker run --name mynginx-%i -p 80:80 my/nginx:latest

[X-Fleet]
Conflicts=nginx@*.service
```

- 这种@命名的Unit是模板Unit, 
- 启动的时候可以在@后面加上一个id, 比如: `fleetctl start nginx@n1.service`, 这个n1就被替换到%i的位置. 从而启动一个服务.

再准备一个nginx服务的sidekick:

```ini
[Unit]
Description=Register Nginx - %i
BindsTo=nginx@%i.service
After=nginx@%i.service

# 每45秒报告一次活跃, 一次报告持续60s+
[Service]
ExecStart=/bin/sh -c "while true; \
do etcdctl set /services/www/nginx@%i \
'{ \"host\": \"%H\", \"port\": 80 }' --ttl 60;sleep 45; \
done"
ExecStop=/usr/bin/etcdctl rm /services/www/nginx@%i

[X-Fleet]
MachineOf=nginx@%i.service
```



#### 实战六: Solid Failover

一个Nginx挂了之后, fleet会发现这件事情, 换衣台机器把它给起来, 这是一种基本的容灾. 但离高可用还差很远.

我们使用vagrant在本机利用visualbox创建有三个机器的集群.

```zsh
$ git clone https://github.com/coreos/coreos-vagrant/
$ cd coreos-vagrant
$ cp config.rb.sample config.rb
```

修改config.rb, 最主要是下面这个字段, 我们准备3个机器. 其他vm的配置需要修改的化可以看这个文件.

```ruby
$num_instances = 3 #初始化3台机器
```

```zsh
$ vagrant up
...
# 一切顺利的话, vagrant已经帮我们创建好了三台虚拟机
$ VBoxManage list vms
"coreos-vagrant_core-01_1541606898580_36820" {77f85ea7-28c7-4ab0-bc2d-7f7367ceb439}
"coreos-vagrant_core-02_1541674586355_26721" {1a947eff-1074-465d-9cc0-e312df7a125b}
"coreos-vagrant_core-03_1541674635515_42823" {de97f477-24e2-4047-b984-e8fde86d33b2}

# 登录机器
$ vagrant ssh core-01
```





## 实战一:  配置一个本地集群

1. 编译etcd





```zsh
$ etcd
...
$ export ETCDCTL_API=3
$ etcdctl put k1 v1
OK
$  etcdctl get k1
k1
v1

# 安装goreman
$ go get github.com/mattn/goreman
$ goreman
$ goreman -f Procfile start
...
```
一切顺利之后, 我们就会启动三个etcd, 分别监听

- 2379

 -  22379
 -  32379



另启一终端


```zsh
$ export ETCDCTL_API=3                                                                                                                                         
$ etcdctl --write-out=table --endpoints=localhost:2379 member list
+------------------+---------+--------+------------------------+------------------------+
|        ID        | STATUS  |  NAME  |       PEER ADDRS       |      CLIENT ADDRS      |
+------------------+---------+--------+------------------------+------------------------+
| 8211f1d0f64f3269 | started | infra1 | http://127.0.0.1:12380 |  http://127.0.0.1:2379 |
| 91bc3c398fb3c146 | started | infra2 | http://127.0.0.1:22380 | http://127.0.0.1:22379 |
| fd422379fda50e48 | started | infra3 | http://127.0.0.1:32380 | http://127.0.0.1:32379 |
+------------------+---------+--------+------------------------+------------------------+
# 测试写入
$ etcdctl put hello etcd
OK
$ etcdctl get hello
hello
etcd

# 从etcd2上读取
$ etcdctl  --endpoints=localhost:22379 get hello
hello
etcd

# 测试容错性, 把etcd2停掉
$ goreman run stop etcd2
$ etcdctl  --endpoints=localhost:22379 get hello 
Error: context deadline exceeded

# 仍然可以工作
$ etcdctl get hello
hello
etcd

# 启动etcd2, 重新读取, 恢复正常
$ goreman run restart etcd2
$ etcdctl  --endpoints=localhost:22379 get hello 
hello
etcd
```



## 实战二: 使用etcd存取信息

etcd的使用方式有两种:

1. 通过etcdctl
2. 通过http/https gRPC

etcd能干的事儿也是分简单:

1. CURD
2. lease
3. watch
4. Batch Operation



## 实战三: 配置内网集群

假如我们想在内网建立三台机器组成的etcd集群, 假设我们的机器配置如下:

|       |              |           |
| ----- | ------------ | --------- |
| etcd1 | 192.168.1.10 | etcd1.org |
| etcd2 | 192.168.1.20 | etcd2.org |
| etcd3 | 192.168.1.30 | etcd3.org |

1. 启动内网集群我们只需要搞明白几个必要的参数, 然后在每个机器上启动etcd即可.
2. 一旦节点在每个机器上成功起来, 以后这些flag或者环境变量就不需要了



```
ETCD_INITIAL_CLUSTER="etcd1=http://192.168.1.10:2380,etcd2=http://192.168.1.20:2380,etcd3=192.168.1.30:2380"
ETCD_INITIAL_CLUSTER_STATE=new
```

```
--initial-cluster "etcd1=http://192.168.1.10:2380,etcd2=http://192.168.1.20:2380,etcd3=192.168.1.30:2380"
--initial-cluster-state new
```

> 如果在内网中建立多个集群, 需要指定initial-cluster-token, 这样可以避免成员意外访问到其他集群的机器



下面这俩通常都是Local Address, 注意只能使用IP
--listen-peer-urls:   etcd集群各个member之间用于通讯的端口
--listen-client-urls: etcd客户端, 响应各种请求的端口


下面这俩提供的端口和IP必须是外网可以访问的
--advertise-client-urls
--initial-advertise-peer-urls

--initial-cluster-token 

- 参看

## 实战四: 配置测试生产环境

## 实战五: 使用paz

- https://github.com/paz-sh/paz-web

- https://www.packet.net



## 实战六:  CoreOS的热升级

四种策略:

1. best-effort
2. etcd-lock
3. reboot
4. off

 配置:

```yaml
#cloud-config
coreos:
  update:
    group: stable
    reboot-strategy: best-effort # 配置升级策略
```





## rkt

> rkt (pronounced "rock it") is a container runtime for applications made by CoreOS
> and is designed for composability, speed, and security.



## Runtime reconfiguration

etcd动态重配置分为两步

1. 通知所有成员新的配置
2. 启动新成员

> 参考: https://github.com/etcd-io/etcd/blob/master/Documentation/op-guide/runtime-reconf-design.md





## Public discovery service

这个服务的主要作用是帮助建立集群, 一开始成员之间还不知道彼此的IP. 一旦成员之间了解彼此的IP, 这时PublicDiscoverService就不再需要了.  这个时候添加或者退出集群用Runtime Reconfiguration API



## fleet: 已经不用再看了, 官方已经不再维护

> https://coreos.com/blog/migrating-from-fleet-to-kubernetes.html , 如这个公告
>
> CoreOS已于2018年2月移除了fleet

>   The fleet is a cluster manager that controls systemd at the cluster level

- fleet的unit文件是在systemd units文件基础上增加了一些额外的属性.到



![](/home/amas/2018-11-08-113255_562x282_scrot.png)



### X-Fleet

#### MachineID

#### MachineOf

#### Confilicts

#### Global



```
History and motivations behind CoreOS
The concept of single system image ( SSI ) computing is an OS architecture that hasn’t
seen much activity since the 1990s, except for a few cases that have longstanding sup-
port to run legacy systems. SSI is an architecture that presents many computers in a
cluster as a single system. There is a single filesystem, shared interprocess communica-
tion ( IPC ) via shared runtime space, and process checkpointing/migration.
MOSIX /openMosix, Kerrighed, VMScluster, and Plan 9 (natively supported) are all
SSI systems. Plan 9 has probably received the most current development activity, which
should tell you something about the popularity of this computing model.
The main drawbacks of SSI are, first, that the systems are often extremely difficult
to configure and maintain and aren’t geared toward generic use. Second, the field has
stagnated significantly: there’s nothing new in SSI , and it has failed to catch on as a
popular model. I think this is because scientific and other Big Data computing have
embraced grid-compute, batch operating models like Condor, BOINC , and Slurm.
These tools are designed to run compute jobs in a cluster and deliver a result; SSI ’s
shared IPC provides little benefit for these applications, because the cost (in time) of
data transmission is eclipsed by the cost of the blocking batch process. In the world of
application server stacks, abstractions by protocols like HTTP and distributed queues
have also made shared IPC not worth investing in.
The problem space now for distributed computing is how to effectively manage
large-scale systems. Whether you’re working on a web stack or distributed batch pro-
cessing, you may not need shared IPC , but the other things that came with SSI have
more apparent value: a shared filesystem means you configure only one system, andABOUT THIS BOOK
xix
process checkpointing and migration mean nodes are disposable and more “cattle-
like.” Without shared IPC , these solutions can be difficult to implement. Some organi-
zations turn to configuration-management systems that apply configuration to many
machines, or set up extremely complicated monitoring systems full of custom logic. In
my experience, configuration-management systems fall short of the goal by only
ensuring any state exactly at runtime; after they’ve made their pass, the state becomes
unknown. These systems are more focused on repeatability than consistency, which is
a fine goal but doesn’t provide the reliability of a shared configuration via a distrib-
uted filesystem. Monitoring systems that attempt to also manage processes are often
either application-specific or hairy to implement and maintain.
Intentionally or not, container systems like Docker laid the groundwork for resur-
recting the advantages of SSI without having to implement shared IPC . Docker guaran-
tees runtime state and provides an execution model that’s abstracted from the OS .
“But Matt,” you may think, “this is the complete opposite of SSI . Every discrete system
now has an even more isolated configuration and runtime, not shared!” Yes, this
approach is orthogonal, but it achieves the same goals. If runtime state is defined only
once (in the Dockerfile, for example) and maintained throughout the life of the con-
tainer, you’ve reached the goal of a single point of configuration. And if you can
orchestrate the discrete process state both remotely and independently from the OS
and the cluster node it’s running on, you’ve achieved the goal of cluster-wide process
scheduling of generic services.
Realizing those possibilities is where there needs to be tooling independent of the
containerization system. This is where CoreOS and its suite of systems come in.
CoreOS provides just enough OS to run a few services; the rest is handled by the
orchestration efforts of etcd and fleet—etcd provides a distributed configuration from
which containers can define their runtime characteristics, and fleet manages distrib-
uted initialization and scheduling of containers. Internally, CoreOS also uses etcd to
provide a distributed lock to automatically manage OS upgrades, which in turn uses
fleet to balance services across the cluster so that a node can upgrade itself.


oreOS is here to solve your scale, availability, and deployment workflow problems. In
this chapter, we’ll go through a simple application deployment of NGINX (a popular
HTTP server) to illustrate how CoreOS achieves some of these solutions, and review
some essential systems. With CoreOS, you won’t have to manage packages, initiate
lengthy upgrade processes, plan out complex configuration files, fiddle with permis-
sions, plan significant maintenance windows (for the OS ), or deal with complicated
configuration schema changes. If you fully embrace CoreOS’s features, your cluster of
nodes will always have the latest version of the OS , and you won’t have any downtime.
These ideas can be a little difficult to grasp when you’re first getting started with
CoreOS, but they embody the philosophy of an immutable OS after boot, which cre-
ates an experience with the OS that you probably aren’t used to. CoreOS’s distributed
scheduler, fleet, manages the state of your application stack, and CoreOS provides the
platform on which those systems orchestrate your services. If you have a computer sci-
ence background, you can consider traditional configuration-management systems as
relying heavily on side effects to constantly manipulate the state of the OS , whereas in
CoreOS, the state of the OS is created once on boot, never changes, and is lost on
shutdown. This is a powerful concept, and it forces architectures with high degrees of
idempotence and no hidden side effects, the results of which are vastly improved cer-
tainty about the reliability of your systems and drastically reduced need for layers of
complex tooling to monitor and manage OS s. In this section, I provide an overview of
the parts that make CoreOS tick and how they complement each other.
```



# Systemd





# Kubernetes

> Kubernetes, an open source container orchestration system.

google已经全面使用容器十多年了. google内部的系统叫做Borg. 可以认为是Kubernetes的前身.

### Kubernetes的组件

#### Master: Kube的大脑

- etcd cluster
- API Service
- Controller Manager
- Scheduler

#### Node: Kube集群中用来运行Pod的机器

集群worker, 可以是VM也可以是bare-metal server, Node由Master来管理. 每个Node上至少运行两个服务:

- Kubelet
- Proxy

#### Pod: 最小的部署单元

- Pod下面可以运行一个或多个容器
- 相同的pod智能运行一次. 如果死了, RC就启动一个新的
- 每个pod都有自己的IP
- 每次RC从模板启动的pod都会分配不同的IP

#### Replication Controllers

- 这个用来确保一定数量的pods正常运行
- 如果过pods太多就杀掉一些
- 如果pods太少就启动一些
- 即便是运行一个pods也要用RC来跑

#### Services

- 用来服务用户,或者服务其他Service的一组Pod
- 你可以给Service定义Label

#### Labels

#### Volumes

- 与特定容器关联的目录

#### Kubectl

- 容器管理工具



![](/home/amas/2018-11-08-191951_1103x823_scrot.png)



Sidecar pattern

Ambassador pattern

```
Sidecar pattern
The sidecar pattern is about co-locating another container in a pod in addition to
the main application container. The application container is unaware of the sidecar
container and just goes about its business. A great example is a central logging agent.
Your main container can just log to stdout , but the sidecar container will send all
logs to a central logging service where they will be aggregated with the logs from the
entire system. The benefits of using a sidecar container versus adding central logging
to the main application container are enormous. First, applications are not burdened
anymore with central logging, which could be a nuisance. If you want to upgrade or
change your central logging policy or switch to a totally new provider, you just need
to update the sidecar container and deploy it. None of your application containers
change, so you can't break them by accident.
Ambassador pattern
The ambassador pattern is about representing a remote service as if it were local and
possibly enforcing some policy. A good example of the ambassador pattern is if you
have a Redis cluster with one master for writes and many replicas for reads. A local
ambassador container can serve as a proxy and expose Redis to the main application
container on the localhost. The main application container simply connects to Redis
on localhost:6379 (Redis default port), but it connects to the ambassador running
in the same pod, which filters the requests, and sends write requests to the real
Redis master and read requests randomly to one of the read replicas. Just like with
the sidecar pattern, the main application has no idea what's going on. That can help
a lot when testing against a real local Redis. Also, if the Redis cluster configuration
changes, only the ambassador needs to be modified; the main application remains
blissfully unaware.


Adapter pattern
The adapter pattern is about standardizing output from the main application
container. Consider the case of a service that is being rolled out incrementally: it
may generate reports in a format that doesn't conform to the previous version. Other
services and applications that consume that output haven't been upgraded yet.
An adapter container can be deployed in the same pod with the new application
container and massage their output to match the old version until all consumers
have been upgraded. The adapter container shares the filesystem with the main
application container, so it can watch the local filesystem, and whenever the new
application writes something, it immediately adapts it.
```





- 安装kubectl: https://kubernetes.io/docs/tasks/tools/install-kubectl/

```zsh


$ kubectl run hello-minikube --image=k8s.gcr.io/echoserver:1.10 --port=8080
deployment.apps/hello-minikube created
$ kubectl expose deployment hello-minikube --type=NodePort
service/hello-minikube exposed
```



#### kubectl explain pods|svc|

## 使用Minikube在单机上使用Kube

```zsh
$ minikube start
...
# 第一次启动的时候, minikube会建立~/.minikube目录, 并且去下载各种资源, 需要等待一会

$ minikube start 
Starting local Kubernetes v1.10.0 cluster...
Starting VM...
Getting VM IP address...
Moving files into cluster...
Setting up certs...
Connecting to cluster...
Setting up kubeconfig...
Starting cluster components...
Kubectl is now configured to use the cluster.
Loading cached images from config file.
# 看到这行就算成功启动了, 否则就是失败, 可以用minikube logs看看哪里出了问题, 
# 如果没什么线索可以尝试minikube stop ; minikube start 重新启动一把

$ kubectl version
Client Version: version.Info{Major:"1", Minor:"12", GitVersion:"v1.12.0", GitCommit:"0ed33881dc4355495f623c6f22e7dd0b7632b7c0", GitTreeState:"clean", BuildDate:"2018-09-27T17:05:32Z", GoVersion:"go1.10.4", Compiler:"gc", Platform:"linux/amd64"}
Server Version: version.Info{Major:"1", Minor:"10", GitVersion:"v1.10.0", GitCommit:"fc32d2f3698e36b93322a3465f63a14e9f0eaead", GitTreeState:"clean", BuildDate:"2018-03-26T16:44:10Z", GoVersion:"go1.9.3", Compiler:"gc", Platform:"linux/amd64"

$ kubectl cluster-info 
Kubernetes master is running at https://192.168.99.101:8443
CoreDNS is running at https://192.168.99.101:8443/api/v1/namespaces/kube-system/services/kube-dns:dns/proxy

To further debug and diagnose cluster problems, use 'kubectl cluster-info dump'.
# 主义这里的master: https://192.168.99.101:8443

$ kubectl get nodes
NAME       STATUS   ROLES    AGE   VERSION
minikube   Ready    master   6h    v1.10.0

# 登录到minikube里
$ minikube ssh

$ kubectl run echo --image=gcr.io/google_containers/echoserver:1.4 --port=8080                                             
deployment.apps/echo created

$ kubectl get pods 
NAME                              READY   STATUS              RESTARTS   AGE
echo-5c7b447f8c-tgsfd             0/1     ContainerCreating   0          59s
hello-minikube-7c77b68cff-sn4cw   1/1     Running             1          1h

$ kubectl expose deployment echo --type=NodePort
service/echo exposed
$ minikube ip
192.168.99.101


$ kubectl get service echo
NAME   TYPE       CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
echo   NodePort   10.105.182.79   <none>        8080:30358/TCP   1m

$ curl http://192.168.99.101:30358
CLIENT VALUES:
client_address=172.17.0.1
command=GET
real path=/
query=nil
request_version=1.1
request_uri=http://192.168.99.101:8080/

SERVER VALUES:
server_version=nginx: 1.10.0 - lua: 10001

HEADERS RECEIVED:
accept=*/*
host=192.168.99.101:30358
user-agent=curl/7.62.0
BODY:
-no body in request-%
```



### 使用kube dashboard

```zsh
# 首先看下kube-system下面都有哪些pod
$ kubectl -n kube-system get pod 
NAME                                    READY   STATUS    RESTARTS   AGE
coredns-c4cffd6dc-j5ggx                 1/1     Running   2          13d
etcd-minikube                           1/1     Running   0          18m
kube-addon-manager-minikube             1/1     Running   3          13d
kube-apiserver-minikube                 1/1     Running   0          18m
kube-controller-manager-minikube        1/1     Running   0          18m
kube-dns-86f4d74b45-qvl9z               3/3     Running   10         13d
kube-proxy-pj9l9                        1/1     Running   0          17m
kube-scheduler-minikube                 1/1     Running   0          18m
kubernetes-dashboard-6f4cfc5d87-ffx7g   1/1     Running   6          13d
storage-provisioner                     1/1     Running   6          13d

$ kubectl -n kube-system get service 
NAME                   TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)         AGE
kube-dns               ClusterIP   10.96.0.10     <none>        53/UDP,53/TCP   13d
kubernetes-dashboard   ClusterIP   10.104.168.9   <none>        80/TCP          13d
```

> 注意: kubernetes-dashboard-* 这一行, 它的状态是Running



接下来就可以通过master的ip和port访问dashboard了:

 - https://192.168.99.101:8443/api/v1/namespaces/kube-system/services/https:kubernetes-dashboard:/proxy/



你应该会看到下面的信息:

```json
{
  "kind": "Status",
  "apiVersion": "v1",
  "metadata": {
    
  },
  "status": "Failure",
  "message": "services \"https:kubernetes-dashboard:\" is forbidden: User \"system:anonymous\" cannot get services/proxy in the namespace \"kube-system\"",
  "reason": "Forbidden",
  "details": {
    "name": "https:kubernetes-dashboard:",
    "kind": "services"
  },
  "code": 403
}
```



有两种方法可以解决这个问题,

方案1:

 -  使用minikube proxy建立一个代理
 -  然后使用chrome设置这个代理

```zsh
$ minikube proxy
Starting to serve on 127.0.0.1:8001

$ chromium --proxy-server=127.0.0.1:8001  
```

> 打开浏览器从这个地址访问:  http://localhost:8001/api/v1/namespaces/kube-system/services/kubernetes-dashboard/proxy/#!/cronjob?namespace=default

此方案比较简单, 但是不适合在生产环境中工作



### 创建一个pod

pod-nginx.yml:

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: web
spec:
  containers:
      - name: nginx-01
        image: nginx
        ports:
          - containerPort: 80
```



```zsh
$ kubectl create -f pod-nginx.yml
$ kubectl describe pods/web
Name:         web
Namespace:    default
Node:         minikube/10.0.2.15
Start Time:   Fri, 23 Nov 2018 22:16:07 +0800
Labels:       <none>
Annotations:  <none>
Status:       Running
IP:           172.17.0.8
Containers:
  nginx-01:
    Container ID:   docker://9c9ff976736c0da86184383069898acdb2915b9ec95172208bb30b4be58cc695
    Image:          nginx
    Image ID:       docker-pullable://nginx@sha256:31b8e90a349d1fce7621f5a5a08e4fc519b634f7d3feb09d53fac9b12aa4d991
    Port:           80/TCP
    Host Port:      0/TCP
    State:          Running
      Started:      Fri, 23 Nov 2018 22:16:27 +0800
    Ready:          True
    Restart Count:  0
    Environment:    <none>
    Mounts:
      /var/run/secrets/kubernetes.io/serviceaccount from default-token-h246d (ro)
Conditions:
  Type           Status
  Initialized    True 
  Ready          True 
  PodScheduled   True 
Volumes:
  default-token-h246d:
    Type:        Secret (a volume populated by a Secret)
    SecretName:  default-token-h246d
    Optional:    false
QoS Class:       BestEffort
Node-Selectors:  <none>
Tolerations:     node.kubernetes.io/not-ready:NoExecute for 300s
                 node.kubernetes.io/unreachable:NoExecute for 300s
Events:
  Type    Reason                 Age    From               Message
  ----    ------                 ----   ----               -------
  Normal  Scheduled              8m3s   default-scheduler  Successfully assigned web to minikube
  Normal  SuccessfulMountVolume  8m3s   kubelet, minikube  MountVolume.SetUp succeeded for volume "default-token-h246d"
  Normal  Pulling                8m2s   kubelet, minikube  pulling image "nginx"
  Normal  Pulled                 7m43s  kubelet, minikube  Successfully pulled image "nginx"
  Normal  Created                7m43s  kubelet, minikube  Created container
  Normal  Started                7m43s  kubelet, minikube  Started container

# 进入容器
$ kubectl exec -it web -- bash
```







## Creating multi-node cluster using kubeadm

## Creating clusters in the cloud

## Creating bare-metal clusters from scratch



