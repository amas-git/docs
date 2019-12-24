k8s-network

## Kubernetes Network

分层:

1. 主机网络层
   - namespace
   - iptables
   - routing
   - IPVLAN
2. 容器网络层
   1. Single-host bridge
   2. Multihost
   3. IP-per-container
3. 服务发现和容器编排层



![](/src/amas-git/docs/source/_posts/k8s-network.assets/2019-12-24-232200_859x583_scrot.png)

### 容器网络层

![](/src/amas-git/docs/source/_posts/k8s-network.assets/2019-12-24-232336_878x583_scrot.png)

1. Host:Container = 1:N
   1. Facebook: 每个主机平均运行10~40个容器
   2. Mesosphere: 物理机上运行250个容器不成问题

## Docker Single-host Networking Mode

单机网络有四种模式:

	- Bridget Networking Mode
	- Host Networking Mode
	- Container Networking Mode
	- No Networking Mode

###  Bridge Mode Networking

Docker创建虚拟的docker0以太网网桥，这也是Docker默认的方式, 在生环境中建议使用网桥模式，配合SDN解决方案更好，如果想控制容器之间的通讯可以用`--iptables`和`--icc`

![](/src/amas-git/docs/source/_posts/k8s-network.assets/2019-12-24-233513_1034x643_scrot.png)

```bash
$ docker run -d -P --net=bridge nginx
```

### Host Mode Networking

容器共享了主机网络namespace.

```bash
$ docker run -d -P --net=host nginx
```



### Container Mode Networking

容器共享指定容器网络namespace， k8s便是利用这种机制。

```bash
$ docker run -d -P --net=bridge nginx
$ docker exec -it target_container ip addr
$ docker run -it --net=container:target_container ubuntu:14.04
```

### No Networking

容器不能访问网络，适合执行一些不需要网络的Job

```bash
$ docker run -d -P --net=none nginx
```



## Docker Multihost Networking Mode

当容器越来越多的时候，单机资源不足以支撑，势必需要在多个主机上运行容器，也就是容器集群。我们面临如下问题:

1. 容器如何跨主机通讯？
2. 如何控制容器与外部世界的通讯？
3. 如何追踪管理集群中的IP分配？
4. 如何保证安全？

> Batteries included but replaceable: 总是提供默认的功能，如果你对某些解决方案不满也有选择的余地



### Overlay

2015年docker公司发布了SDN方案SocketPlane, 最终重新命名为DockerOverlayDriver, 作为Multihost Networking的默认解决方案。DockerOverlayDriver扩展了网桥模式用于点对点通讯，使用kv存储(支持zookeeper,etcd.consul)记录集群状态。

### Flannel

CoreOS的Flannel实现了一个虚拟网络，为每个主机准备一个子网用于运行容器。可以为每个容器分配唯一的IP, 这样集群内部的容器就可以相互访问。同时Flannel也支持注入VXLAN, AWS VPC等。Flannel的优点在于降低了端口映射的复杂性。

### Weave

Weaveworks的Weave，2层OverlayNetwork，可以看作一个switch

### Calico

Metaswitch的Calico使用了标准的IP路由: BorderGatewayProtocal:RFC1105, 3层OverlayNetwork. 这个主要为数据中心设计的。

### Open vSwitch

虚拟switch, 可被NetFlow,IPFIX,LACP,802.lag等协议控制， 类似于Vmware的vNetworkDistributedvSwitch或思科的Nexus1000V

### Pipework

Pipework由Docker著名工程师Jérôme Petazzoni发起，目标是要实现Linux容器的SDN。使用cgroup+namespace,支持Docker和LXC容器。

### OpenVPN

OpenVPN所创建的网络也可以解决容器的跨主机通讯问题。可以参考:

https://www.digitalocean.com/community/tutorials/how-to-run-openvpn-in-a-docker-container-on-ubuntu-14-04

### More:

docker network命令 1.9引入， 可以让容器动态的访问其他网络。

- http://developerblog.info/2015/11/16/splendors-and-miseries-of-docker-network/
- https://www.oreilly.com/learning/docker-networking-service-discovery
- https://www.weave.works/blog/docker-networking-1-9-weave-plugin/
- https://www.docker.com/blog/tag/service-discovery/

#### IPVLAN

Linux 3.19+引入的IP-per-container功能，为每一个主机上的容器分配一个唯一的IP地址。工作在L2和L3层。

####  IPAM

IP Address Management, 多个主机构成的集群如何分配IP资源。

### 容器和服务发现

 - ZK
   	- java
   	- R/W + nginx
 - etcd
    - raft
    - CoreOS
    - go
    - confd + nginx
 - Consul
    - raft
    - HashiCorp
    - go
 - DNS解决方案
    - Mesos-DNS
    - SkyDNS
    - WeaveDNS
    - CoreDNS

## 容器编排

	- 容器调度
	- 容器升级
	- 健康检查
	- 扩容
	- 服务发现
	- Organizational Primitives



## Kubernetes Service

通常Service可以指定一个ClusterIP(也叫做VIP), 一个ClusterIP对应一个或多个Endpoint对象(以前每个Endpoint对多保存100个IP, 多于100会创建新的Endpoint, 后来k8s引入EndpointSlice解决此问题)， 在Endpoint对象中保存了Pod的实际IP, 这样VIP的流量会被节点上的kube-proxy根据Endpoint对象按照负载均衡算法分配到PodIP上

Service也可以把ClusterIP设置为None, 这样就不会自动建立Endpoint对象，此时你有两个选择

1. 手动创建Endpoint对象
2. 不创建Endpoint对象，负载均衡算法由客户端自己解决，比如查询CoreDNS获得服务IP

## Kubernetes DNS

DNS在k8s中是以add-on的形式提供的，就是说没有DNS, k8s也可以使用。然而很少有不需要DNS的。所以基本上都会默认安装，以前是kube-dns, 现在是CoreDNS.

> Cluster Domain: 集群有一个默认的集群域名cluster.local, 启动k8s的时候你可以修改它

对于每一个ClusterIP, DNS系统会给它准备一个对应的域名

> <service-name>.<namespace>.svc.<cluster-domain>
>
> e.g.: hello.default.svc.cluster.local

对于每一个ClusterIP, DNS系统也会准备对应的PTR记录，支持IP查询域名

## CoreDNS

>为什么开发CoreDNS
>
>1. 安全，性能
>2. 更好解决容器化微服务的服务发现问题

```bash

$ kubectl get svc
NAME         TYPE        CLUSTER-IP    EXTERNAL-IP   PORT(S)          AGE
kubernetes   ClusterIP   10.96.0.1     <none>        443/TCP          2d5h
xecho        NodePort    10.96.49.52   <none>        8888:31401/TCP   2d4h

# 使用infoblox调试dns
$ kubectl run --restart=Never -it --image infoblox/dnstools dnstools
# 用nslookup查看
dnstools# nslookup xecho.default.svc.cluster.local
Server:         10.96.0.10      # DNS服务器IP
Address:        10.96.0.10#53   # DNS服务器IP,使用标准的53端口

Name:   xecho.default.svc.cluster.local
Address: 10.96.49.52            # xecho服务的ClusterIP

# 使用dig命令查看一下xecho服务对应的dns解析
dnstools# dig xecho.default.svc.cluster.local
; <<>> DiG 9.11.3 <<>> xecho.default.svc.cluster.local
;; global options: +cmd
;; Got answer:
;; WARNING: .local is reserved for Multicast DNS
;; You are currently testing what happens when an mDNS query is leaked to DNS
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 63260
;; flags: qr aa rd; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 1
;; WARNING: recursion requested but not available

;; OPT PSEUDOSECTION:
; EDNS: version: 0, flags:; udp: 4096
; COOKIE: 9982e265f6f4219e (echoed)
;; QUESTION SECTION:
;xecho.default.svc.cluster.local. IN    A

;; ANSWER SECTION:
xecho.default.svc.cluster.local. 30 IN  A       10.96.49.52

;; Query time: 0 msec
;; SERVER: 10.96.0.10#53(10.96.0.10)
;; WHEN: Mon Dec 23 08:42:19 UTC 2019
;; MSG SIZE  rcvd: 119
```



Headless Server实验

```
$ tree .
.
├── headless-deployment.yaml
└── headless-svc.yaml
```



headless-deployment.yaml: 

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
    name: headless
    namespace: default
spec:
    replicas: 4
    selector:
        matchLabels:
            app: headless
    template:
        metadata:
            labels:
                app: headless
        spec:
            hostname: myhost
            subdomain: headless
            containers:
                - image: nginx
                  name: nginx
                  ports:
                      - containerPort: 80
                        name: http
                        protocol: TCP
```



 headless-svc.yaml:

```yaml
apiVersion: v1
kind: Service
metadata:
    name: headless
spec:
    selector:
        app: headless
    type: ClusterIP
    clusterIP: None
    ports:
        - name: http
          port: 80
          protocol: TCP
```



```bash
$ kubectl get pods -l app=headless -o wide
NAME                        READY   STATUS    RESTARTS   AGE   IP            NODE       NOMINATED NODE   READINESS GATES
headless-5774f796c8-hjcff   1/1     Running   0          32m   172.17.0.12   minikube   <none>           <none>
headless-5774f796c8-r7tbm   1/1     Running   0          32m   172.17.0.10   minikube   <none>           <none>
headless-5774f796c8-vr9lm   1/1     Running   0          32m   172.17.0.11   minikube   <none>           <none>
headless-5774f796c8-vsp75   1/1     Running   0          32m   172.17.0.9    minikube   <none>           <none>

$ kubectl run --restart=Never -it --image infoblox/dnstools dnstools
dnstools# dig  myhost.headless.default.svc.cluster.local

; <<>> DiG 9.11.3 <<>> myhost.headless.default.svc.cluster.local
;; global options: +cmd
;; Got answer:
;; WARNING: .local is reserved for Multicast DNS
;; You are currently testing what happens when an mDNS query is leaked to DNS
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 47795
;; flags: qr aa rd; QUERY: 1, ANSWER: 4, AUTHORITY: 0, ADDITIONAL: 1
;; WARNING: recursion requested but not available

;; OPT PSEUDOSECTION:
; EDNS: version: 0, flags:; udp: 4096
; COOKIE: b81df53c5338bed8 (echoed)
;; QUESTION SECTION:
;myhost.headless.default.svc.cluster.local. IN A

;; ANSWER SECTION:
myhost.headless.default.svc.cluster.local. 30 IN A 172.17.0.10
myhost.headless.default.svc.cluster.local. 30 IN A 172.17.0.12
myhost.headless.default.svc.cluster.local. 30 IN A 172.17.0.11
myhost.headless.default.svc.cluster.local. 30 IN A 172.17.0.9

;; Query time: 1 msec
;; SERVER: 10.96.0.10#53(10.96.0.10)
;; WHEN: Mon Dec 23 11:35:07 UTC 2019
;; MSG SIZE  rcvd: 310

# A records中包含SRV records, 可以看下
dnstools# dig  -t srv myhost.headless.default.svc.cluster.local
; <<>> DiG 9.11.3 <<>> -t srv myhost.headless.default.svc.cluster.local
;; global options: +cmd
;; Got answer:
;; WARNING: .local is reserved for Multicast DNS
;; You are currently testing what happens when an mDNS query is leaked to DNS
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 29633
;; flags: qr aa rd; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 5
;; WARNING: recursion requested but not available

;; OPT PSEUDOSECTION:
; EDNS: version: 0, flags:; udp: 4096
; COOKIE: f11a012a85ce82a4 (echoed)
;; QUESTION SECTION:
;myhost.headless.default.svc.cluster.local. IN SRV

;; ANSWER SECTION:
myhost.headless.default.svc.cluster.local. 30 IN SRV 0 25 80 myhost.headless.default.svc.cluster.local.

;; ADDITIONAL SECTION:
myhost.headless.default.svc.cluster.local. 30 IN A 172.17.0.12
myhost.headless.default.svc.cluster.local. 30 IN A 172.17.0.11
myhost.headless.default.svc.cluster.local. 30 IN A 172.17.0.9
myhost.headless.default.svc.cluster.local. 30 IN A 172.17.0.10

;; Query time: 1 msec
;; SERVER: 10.96.0.10#53(10.96.0.10)
;; WHEN: Mon Dec 23 11:42:47 UTC 2019
;; MSG SIZE  rcvd: 412
```



## 参考

 - https://www.amazon.com/Using-Docker-Developing-Deploying-Containers/dp/1491915765

 - https://github.com/docker/docker-bench-security

 - https://www.weave.works/blog/automating-weave-discovery-docker/

   