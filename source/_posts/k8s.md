# k8s

## 安装

### 使用minikube

```bash
$ pacman -S minikube
# 或者从官方安装最新
$ curl -Lo minikube https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64 && chmod +x minikube
$ ./minkube start
😄  minikube v1.6.1 on Arch 
✨  Automatically selected the 'virtualbox' driver (alternates: [none])
💿  Downloading VM boot image ...
    > minikube-v1.6.0.iso.sha256: 65 B / 65 B [--------------] 100.00% ? p/s 0s
    > minikube-v1.6.0.iso: 150.93 MiB / 150.93 MiB  100.00% 687.51 KiB p/s 3m45
🔥  Creating virtualbox VM (CPUs=2, Memory=2000MB, Disk=20000MB) ...
🐳  Preparing Kubernetes v1.17.0 on Docker '19.03.5' ...
💾  Downloading kubelet v1.17.0
💾  Downloading kubeadm v1.17.0
🚜  Pulling images ...
🚀  Launching Kubernetes ... 

# 也可以指定minikube使用的cpu和内存
$ minikube delete; minikube start --extra-config=kubelet.authentication-token-webhook=true --cpus 4 --memory 8192

# 查看集群的IP
$ minkube ip
192.168.99.101

# 部署echo server
$ kubectl create deployment hello --image=k8s.gcr.io/echoserver:1.10
deployment.apps/hello created

# 查看deployment
$ kubectl get deployments hello -o yaml 
```

### Deployment
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  annotations:
    deployment.kubernetes.io/revision: "1"
  creationTimestamp: "2019-12-18T09:57:30Z"
  generation: 1
  labels:
    app: hello
  name: hello
  namespace: default
  resourceVersion: "80943"
  selfLink: /apis/apps/v1/namespaces/default/deployments/hello
  uid: abf0036c-b91f-4512-996c-3ec325005843
spec:
  progressDeadlineSeconds: 600
  replicas: 1
  revisionHistoryLimit: 10
  selector:
    matchLabels:
      app: hello
  strategy:
    rollingUpdate:
      maxSurge: 25%
      maxUnavailable: 25%
    type: RollingUpdate
  template:
    metadata:
      creationTimestamp: null
      labels:
        app: hello
    spec:
      containers:
      - image: k8s.gcr.io/echoserver:1.10
        imagePullPolicy: IfNotPresent
        name: echoserver
        resources: {}
        terminationMessagePath: /dev/termination-log
        terminationMessagePolicy: File
      dnsPolicy: ClusterFirst
      restartPolicy: Always
      schedulerName: default-scheduler
      securityContext: {}
      terminationGracePeriodSeconds: 30
status:
  availableReplicas: 1
  conditions:
  - lastTransitionTime: "2019-12-18T09:58:40Z"
    lastUpdateTime: "2019-12-18T09:58:40Z"
    message: Deployment has minimum availability.
    reason: MinimumReplicasAvailable
    status: "True"
    type: Available
  - lastTransitionTime: "2019-12-18T09:57:30Z"
    lastUpdateTime: "2019-12-18T09:58:40Z"
    message: ReplicaSet "hello-76dfd64498" has successfully progressed.
    reason: NewReplicaSetAvailable
    status: "True"
    type: Progressing
  observedGeneration: 1
  readyReplicas: 1
  replicas: 1
  updatedReplicas: 1
```

```zsh
# 我们把Deployment通过Service暴露给外部
$ kubectl expose deployment hello --type=NodePort --port=8080

# 查看Service
$ kubectl get svc
NAME         TYPE           CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
hello        NodePort       10.96.221.248   <none>        8080:30112/TCP   11m
$ kubectl get svc hello -o yaml
$ kubectl get endpoints hello -o yaml
```

### Service 
```yaml
apiVersion: v1
kind: Service
metadata:
  creationTimestamp: "2019-12-18T10:09:53Z"
  labels:
    app: hello
  name: hello
  namespace: default
  resourceVersion: "82411"
  selfLink: /api/v1/namespaces/default/services/hello
  uid: 37bf6aee-6dbb-461b-9624-0ec88ca53153
spec:
  clusterIP: 10.96.221.248
  externalTrafficPolicy: Cluster
  ports:
  - nodePort: 30112
    port: 8080
    protocol: TCP
    targetPort: 8080
  selector:
    app: hello
  sessionAffinity: None
  type: NodePort
status:
  loadBalancer: {}
```
### Endpoint
```yaml
apiVersion: v1
kind: Endpoints
metadata:
  annotations:
    endpoints.kubernetes.io/last-change-trigger-time: "2019-12-18T10:09:53Z"
  creationTimestamp: "2019-12-18T10:09:53Z"
  labels:
    app: hello
  name: hello
  namespace: default
  resourceVersion: "82412"
  selfLink: /api/v1/namespaces/default/endpoints/hello
  uid: 91cced1d-7202-45a2-a229-3a4f77813e79
subsets:
- addresses:
  - ip: 172.17.0.5
    nodeName: minikube
    targetRef:
      kind: Pod
      name: hello-76dfd64498-674z2
      namespace: default
      resourceVersion: "80941"
      uid: f92988e8-3095-4a4d-a6e7-764cb6be2b14
  ports:
  - port: 8080
    protocol: TCP
```

```zsh
# 确保Deployment的Pod已经启动成功
$ kubectl get pod 
NAME                          READY   STATUS    RESTARTS   AGE
hello-76dfd64498-674z2        1/1     Running   0          1m

# 查看Pod
$ kubectl get pod hello-76dfd64498-674z2 -o yaml
```

### Pod
```yaml
apiVersion: v1
kind: Pod
metadata:
  creationTimestamp: "2019-12-18T09:57:30Z"
  generateName: hello-76dfd64498-
  labels:
    app: hello
    pod-template-hash: 76dfd64498
  name: hello-76dfd64498-674z2
  namespace: default
  ownerReferences:
  - apiVersion: apps/v1
    blockOwnerDeletion: true
    controller: true
    kind: ReplicaSet
    name: hello-76dfd64498
    uid: 2fdb6791-dca8-4bd7-9108-a2668f408887
  resourceVersion: "80941"
  selfLink: /api/v1/namespaces/default/pods/hello-76dfd64498-674z2
  uid: f92988e8-3095-4a4d-a6e7-764cb6be2b14
spec:
  containers:
  - image: k8s.gcr.io/echoserver:1.10
    imagePullPolicy: IfNotPresent
    name: echoserver
    resources: {}
    terminationMessagePath: /dev/termination-log
    terminationMessagePolicy: File
    volumeMounts:
    - mountPath: /var/run/secrets/kubernetes.io/serviceaccount
      name: default-token-vc276
      readOnly: true
  dnsPolicy: ClusterFirst
  enableServiceLinks: true
  nodeName: minikube
  priority: 0
  restartPolicy: Always
  schedulerName: default-scheduler
  securityContext: {}
  serviceAccount: default
  serviceAccountName: default
  terminationGracePeriodSeconds: 30
  tolerations:
  - effect: NoExecute
    key: node.kubernetes.io/not-ready
    operator: Exists
    tolerationSeconds: 300
  - effect: NoExecute
    key: node.kubernetes.io/unreachable
    operator: Exists
    tolerationSeconds: 300
  volumes:
  - name: default-token-vc276
    secret:
      defaultMode: 420
      secretName: default-token-vc276
status:
  conditions:
  - lastProbeTime: null
    lastTransitionTime: "2019-12-18T09:57:30Z"
    status: "True"
    type: Initialized
  - lastProbeTime: null
    lastTransitionTime: "2019-12-18T09:58:40Z"
    status: "True"
    type: Ready
  - lastProbeTime: null
    lastTransitionTime: "2019-12-18T09:58:40Z"
    status: "True"
    type: ContainersReady
  - lastProbeTime: null
    lastTransitionTime: "2019-12-18T09:57:30Z"
    status: "True"
    type: PodScheduled
  containerStatuses:
  - containerID: docker://1d98eb60c3fdfe6459a3c0ad1ea63d8ac32128c611ce5cf30936e3e1c1cc7eae
    image: k8s.gcr.io/echoserver:1.10
    imageID: docker-pullable://k8s.gcr.io/echoserver@sha256:cb5c1bddd1b5665e1867a7fa1b5fa843a47ee433bbb75d4293888b71def53229
    lastState: {}
    name: echoserver
    ready: true
    restartCount: 0
    started: true
    state:
      running:
        startedAt: "2019-12-18T09:58:40Z"
  hostIP: 192.168.99.101
  phase: Running
  podIP: 172.17.0.5
  podIPs:
  - ip: 172.17.0.5
  qosClass: BestEffort
  startTime: "2019-12-18T09:57:30Z"
```

```zsh
#  万事具备，可以从外部访问服务了
$ minikube service hello --url
http://192.168.99.101:30112
# 或者使用kubectl获得service的端口
$ kubectl get service hello --output='jsonpath="{.spec.ports[0].nodePort}"'
"30112"
$ kubectl get service hello -o yaml | grep nodePort
  - nodePort: 30112

# 可以通过节点Port访问这个服务了
$ curl http://192.168.99.101:30112

Hostname: hello-76dfd64498-674z2

Pod Information:
        -no pod information available-

Server values:
        server_version=nginx: 1.13.3 - lua: 10008

Request Information:
        client_address=172.17.0.1
        method=GET
        real path=/
        query=
        request_version=1.1
        request_scheme=http
        request_uri=http://192.168.99.101:8080/

Request Headers:
        accept=*/*
        host=192.168.99.101:30112
        user-agent=curl/7.67.0

Request Body:
        -no body in request-
```

###  使用minikube的docker

```zsh
# 我们可以直接使用minkube的docker runtime, 这样就不用搞一个Registry了
$ minikube docker-env
export DOCKER_TLS_VERIFY="1"
export DOCKER_HOST="tcp://192.168.99.101:2376"
export DOCKER_CERT_PATH="/home/amas/.minikube/certs"
# Run this command to configure your shell:
# eval $(minikube docker-env)
$ eval $(minikube docker-env)
```

minikube创建了一个叫minikube的kubectrl context, 当我们想操作其他集群后，需要切换context, 想要再用回minikube可以

```zsh
$ kubectl config use-context minikube
# 或者明确设置kubectl的上下文
$ kubectl get pods --context=minikube
```



### 安装集群

- https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/install-kubeadm/

>你的任务:
>
>1. 安装单节点的k8s (single control-plane node)
>2. 安装多节点的k8s集群 (high availability )
>3. 通过: Certified k8s: https://github.com/cncf/k8s-conformance
>4. 通过vagrant+ansible部署到本地虚机
>5. 通过ansible部署到Pi集群
>6. 通过ansible部署到Aws
>7. 通过ansible部署到DigitalOcean
>8. 通过ansible部署到阿里云
>9. 使用minikube
>10. 使用minishift
>11. 使用iosti



 安装条件:

- OS

  - Ubuntu 16.04+
  - Debian 9+
  - CentOS 7
  - Red Hat Enterprise Linux (RHEL) 7
  - Fedora 25+
  - HypriotOS v1.0.1+
  - Container Linux (tested with 1800.6.0)

- 2GB+ RAM

- 2CPUs+

- hostname/MAC/product_uuid 每个节点必须唯一

- 保持对应的端口打开

  - Master

    - TCP inbound 6443* : k8s API server
    - TCP inbound 2379-2380: etcd server client API @usedby kube-apiserver, etcd
    - TCP inbound 10250: kublet API @usedby Self, Control plane
    - TCP inbound 10251: kube-scheduler @usedby self
    - TCP inbound 10252: kube-controller-manager @usedby selfWorker

    - TCP inbound 10250: kublet API @usedby self,Control plane
    - TCP inbound 30000-32767: NodePort Services **

- Disable Swap

- 确保iptables不使用nftables作为后端

  - Deian 10

  - Ubuntu 19.04

    ```sh
    sudo update-alternatives --set iptables /usr/sbin/iptables-legacy
    sudo update-alternatives --set ip6tables /usr/sbin/ip6tables-legacy
    sudo update-alternatives --set arptables /usr/sbin/arptables-legacy
    sudo update-alternatives --set ebtables /usr/sbin/ebtables-legacy
    ```

  - Fedora 29

    ```sh
    update-alternatives --set iptables /usr/sbin/iptables-legacy
    ```

- 安装Runtime

  - v1.6.0+, k8s默认用CRI(Container Runtime Interface)
  - v1.14.0+， kubeadm会自动检测容器runtime:
    - docker: /var/run/docker.sock
    - containerd: /run/containerd/containerd.sock
    - CRI-O: /var/run/crio/crio.sock
    - 总是尽量使用docker

- 安装kubeadm, kubelet, kubectl

  - kubeadm: 用来搭建k8s集群
  - kubelet: 运行在所有机器上，管理集群内部的各种组件
  - kubectl: 控制集群的命令
  - 注意: kubeadm不会安装和管理kubelet和kubectrl, 你必须保证三者之间尽可能使用相同的版本(见: https://kubernetes.io/docs/setup/release/version-skew-policy/)

- 配置cgroup diver, kublet会用到

  - 当你使用docker时，kubeadm会自动检测cgroup驱动，并且配置到/var/lib/kubelet/kubeadm-flags.env中

  - 如果你使用CRI: 则必须修改/etc/default/kubelet的cgroup-diver: 

    ```ini
    KUBELET_EXTRA_ARGS=--cgroup-driver=<value>
    ```

  > 注意： 只有CRI使用的cgroup驱动不是cgroupfs的时候，你才需要这么做

  - 重新启动kubelet你需要:

    ```sh
    systemctl daemon-reload
    systemctl restart kubelet
    ```

- k8s的发布周期为9个月
	- v1.6.x: 2019年9月～2020年7月
	

----
## LET'S GO

### 0x00: 安装docker

```bash
# 1. 安装docker的GPG key
- name: Add an apt signing key for Docker
   apt_key:
     url: https://download.docker.com/linux/ubuntu/gpg
     state: present
      
$ curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -

# 2. 将docker官方的源加入到apt中
- name: Add apt repository for stable version
   apt_repository:
     repo: deb [arch=amd64] https://download.docker.com/linux/ubuntu xenial stable
     state: present
$ sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"

# 3. 更新
$ sudo apt-get update
# 如果不放心，可以确认下docker-ce确实使用官方的源
$ apt-cache policy docker-ce 

# 4. 开始安装
  - name: Install docker and its dependecies
    apt: 
      name: "{{ packages }}"
      state: present
      update_cache: yes
    vars:
      packages:
      - docker-ce 
      - docker-ce-cli 
      - containerd.io
$ sudo apt-get install -y docker-ce docker-ce-cli containerd.io

# 5. 确认下docker运行的状态
$ sudo systemctl status docker

# 6. 将当前用户加入到docker group里
$ sudo usermod -aG docker ${USER}
# 需要重新登录下才能让权限生效
$ sudo su - ${USER} 

# 7. 检查下是否可以使用docker
$ docker info
```



关闭swap

```bash
$ sudo swapoff -a  
# 为了保证重启的时候也可以关闭swap, 需要检查/etc/fstab里面是否存在swap分区，有的话用sed干掉
$ sudo sed -i '/ swap / s/^/#/' /etc/fstab
```



### 0x01: 安装kubeadm, kubelet, kubectl

```bash
sudo apt-get update && sudo apt-get install -y apt-transport-https curl
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -
cat <<EOF | sudo tee /etc/apt/sources.list.d/kubernetes.list
deb https://apt.kubernetes.io/ kubernetes-xenial main
EOF
sudo apt-get update
sudo apt-get install -y kubelet kubeadm kubectl
sudo apt-mark hold kubelet kubeadm kubectl

# 可以查看下版本是不是匹配
$ kubeadm version
$ kubectl verison
$ kubelet --version
```

### 0x02: 初始化control-plane node
 1. ControlPlaneNode就是ControlPlane组建运行的节点，包括
    - etcd
    - API server
    
 2. 为了便于日后升级为高可用集群，可以加上
    
     - --control-plane-endpoint <ip>
     
 3. 安装Pod network add-on
     - SEE: https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/create-cluster-kubeadm/#pod-network
     
 >Pod Network Add-on
 >
 >这个东西可以让Pod之间可以互相访问, 简单的话使用Flannel, 生产环境为了安全可使用Calico

 4. kubeadm init

    - --apiserver-advertise-address
    - --control-plane-endpoint

```bash
# 初始化的时候要用root权限， 这个时间会比较长
# 注意--pod-network-cidr的设置和你选择的add-on插件有关
#  - calico: --pod-network-cidr=192.168.0.0/16
#  - flannel: --pod-network-cidr=10.244.0.0/16 
$ sudo kubeadm init --apiserver-advertise-address="192.168.50.10" --apiserver-cert-extra-sans="192.168.50.10"  --node-name k8s-master --pod-network-cidr=192.168.0.0/16
Your Kubernetes control-plane has initialized successfully!

To start using your cluster, you need to run the following as a regular user:

  mkdir -p $HOME/.kube
  sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
  sudo chown $(id -u):$(id -g) $HOME/.kube/config

You should now deploy a pod network to the cluster.
Run "kubectl apply -f [podnetwork].yaml" with one of the options listed at:
  https://kubernetes.io/docs/concepts/cluster-administration/addons/

Then you can join any number of worker nodes by running the following on each as root:

kubeadm join 192.168.50.10:6443 --token tl5guk.eks9pdcjvbsxxh1d \
    --discovery-token-ca-cert-hash sha256:f52e083623e59d8fdf0d9fbfbf2177d6db94ffa8ab8da42ebb7d024c8940a14d 
```

 ```bash
# 成功之后提示我们给当前用户配置一下，这样不必用root权限操作
$ mkdir -p $HOME/.kube
$ sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
$ sudo chown $(id -u):$(id -g) $HOME/.kube/config

# 最后根据提示，我们可以安装pod networking add-on了， 我们使用calico
$ kubectl apply -f https://docs.projectcalico.org/v3.8/manifests/calico.yaml

# 最后我们检查下是否OK
$ kubectl get pods --all-namespaces
NAMESPACE     NAME                                      READY   STATUS              RESTARTS   AGE
kube-system   calico-kube-controllers-55754f75c-hqnfg   0/1     Pending             0          38s
kube-system   calico-node-fr4j9                         0/1     PodInitializing     0          38s
kube-system   coredns-5644d7b6d9-7hwzb                  0/1     Pending             0          21m
kube-system   coredns-5644d7b6d9-96whr                  0/1     ContainerCreating   0          21m
kube-system   etcd-k8s-master                           1/1     Running             0          20m
kube-system   kube-apiserver-k8s-master                 1/1     Running             0          19m
kube-system   kube-controller-manager-k8s-master        1/1     Running             0          19m
kube-system   kube-proxy-bfx5c                          1/1     Running             0          21m
kube-system   kube-scheduler-k8s-master                 1/1     Running             0          20m

# 一且OK了，还需要做一件事情，处于安全和稳定的考虑，通常不会在control-plane上运行pod, 我们需要改变一下
# 删除node-role.kubernetes.io/master
# https://www.linode.com/docs/kubernetes/getting-started-with-kubernetes/
$ kubectl taint nodes --all node-role.kubernetes.io/master-
$ kubectl get nodes
$ kubectl get namespaces
 ```



### Smoke Test

一切就绪，我们来测试一下k8s是不是work

```zsh
$ kubectl create deployment nginx --image=nginx
$ kubectl get pods -l app=nginx

# 建立一个Pod
$ POD_NAME=$(kubectl get pods -l app=nginx -o jsonpath="{.items[0].metadata.name}")
$ kubectl port-forward $POD_NAME 8080:80
$ curl --head http://127.0.0.1:8080
$ kubectl exec -ti $POD_NAME -- nginx -v

# 查看日志
$ kubectl logs $POD_NAME

# 运行一个service
$ kubectl expose deployment nginx --port 80 --type NodePort
$ NODE_PORT=$(kubectl get svc nginx --output=jsonpath='{range .spec.ports[0]}{.nodePort}')
$ print $NODE_PORT
$ 31071

# 接下来我们可以在外部系统访问这个端口的http服务了
$ curl http://192.168.50.10:31071
```



```bash
# 清除kubeadm所做的事情
$ kubectl drain <node name> --delete-local-data --force --ignore-daemonsets
$ kubectl delete node <node name>

# 重启
$ kubeadm reset

# 撤销iptables设置
$ iptables -F && iptables -t nat -F && iptables -t mangle -F && iptables -X
# 如果是ipvs
$ ipvsadm -C
```



安装补全脚本:

```bash
# BASH
$ sudo apt-get install bash-completion
$ echo 'source /usr/share/bash-completion/bash_completion' >> ~/.bashrc
$ echo 'alias k=kubectl' >>~/.bashrc
$ echo 'complete -F __start_kubectl k' >>~/.bashrc
$ kubectl completion bash > /tmp/kubect
$ sudo cp /tmp/kubect /etc/bash_completion.d/
$ . ~/.bashrc
```





目前支持的Network Add-on:

- https://kubernetes.io/docs/concepts/cluster-administration/addons/
- [ACI](https://www.github.com/noironetworks/aci-containers) provides integrated container networking and network security with Cisco ACI.
- [Calico](https://docs.projectcalico.org/latest/getting-started/kubernetes/) is a secure L3 networking and network policy provider.
- [Canal](https://github.com/tigera/canal/tree/master/k8s-install) unites Flannel and Calico, providing networking and network policy.
- [Cilium](https://github.com/cilium/cilium) is a L3 network and network policy plugin that can enforce HTTP/API/L7 policies transparently. Both routing and overlay/encapsulation mode are supported.
- [CNI-Genie](https://github.com/Huawei-PaaS/CNI-Genie) enables Kubernetes to seamlessly connect to a choice of CNI plugins, such as Calico, Canal, Flannel, Romana, or Weave.
- [Contiv](http://contiv.github.io/) provides configurable networking (native L3 using BGP, overlay using vxlan, classic L2, and Cisco-SDN/ACI) for various use cases and a rich policy framework. Contiv project is fully [open sourced](http://github.com/contiv). The [installer](http://github.com/contiv/install) provides both kubeadm and non-kubeadm based installation options.
- [Contrail](http://www.juniper.net/us/en/products-services/sdn/contrail/contrail-networking/), based on [Tungsten Fabric](https://tungsten.io/), is an open source, multi-cloud network virtualization and policy management platform. Contrail and Tungsten Fabric are integrated with orchestration systems such as Kubernetes, OpenShift, OpenStack and Mesos, and provide isolation modes for virtual machines, containers/pods and bare metal workloads.
- [Flannel](https://github.com/coreos/flannel/blob/master/Documentation/kubernetes.md) is an overlay network provider that can be used with Kubernetes.
- [Knitter](https://github.com/ZTE/Knitter/) is a network solution supporting multiple networking in Kubernetes.
- [Multus](https://github.com/Intel-Corp/multus-cni) is a Multi plugin for multiple network support in Kubernetes to support all CNI plugins (e.g. Calico, Cilium, Contiv, Flannel), in addition to SRIOV, DPDK, OVS-DPDK and VPP based workloads in Kubernetes.
- [NSX-T](https://docs.vmware.com/en/VMware-NSX-T/2.0/nsxt_20_ncp_kubernetes.pdf) Container Plug-in (NCP) provides integration between VMware NSX-T and container orchestrators such as Kubernetes, as well as integration between NSX-T and container-based CaaS/PaaS platforms such as Pivotal Container Service (PKS) and OpenShift.
- [Nuage](https://github.com/nuagenetworks/nuage-kubernetes/blob/v5.1.1-1/docs/kubernetes-1-installation.rst) is an SDN platform that provides policy-based networking between Kubernetes Pods and non-Kubernetes environments with visibility and security monitoring.
- [Romana](http://romana.io/) is a Layer 3 networking solution for pod networks that also supports the [NetworkPolicy API](https://kubernetes.io/docs/concepts/services-networking/network-policies/). Kubeadm add-on installation details available [here](https://github.com/romana/romana/tree/master/containerize).
- [Weave Net](https://www.weave.works/docs/net/latest/kube-addon/) provides networking and network policy, will carry on working on both sides of a network partition, and does not require an external database.

> 注意: network必须最先部署， 然后才能安装CoreDNS
>
> 里国内外kubeadm只支持CNI(Container Network Interface)
>
> IPV6需要安装:
>
> - CNI v0.60+
> - CNI bridge
> - local-ipam
>
> 注意: kubeadm强制使用RBAC: https://kubernetes.io/docs/reference/access-authn-authz/rbac/, 你要确保network manifest支持RBAC

> 注意: Pod网络不能与主机网络重合, 你可以通过--pod-network-cidr设定合适的CIDR

```bash
$ kubectl apply -f <add-on.yaml>
```

## 故障排除

```bash
# 如果重启机器发现k8s没有启动，可以看kubelet日志
$ journalctl -xeu kubelet
```

## k8s的组成

	- kubernetes
	- containerd: 容器标准运行环境
	- coredns: DNS server/forwarder
	- cni: Container Network Interface
	- etcd: 可靠的分布式kv存储

## 生产环境选择

- 红帽的openshift的最小化部署https://github.com/MiniShift/minishift
- CoreOS的tectonic: https://coreos.com/tectonic/

## Kubectl 

```bash
# kubectl可以管理多个集群
# 查看
$ kubectl config current-context
minikube
# 修改
$ kubectl config set-context my-context --namespace=mystuff
# 配置文件： ~/.kube/config中
$ cat ~/.kube/config
# 使用
$ kubectl config use-context my-context --namespace=mystuff

# 查看
$ kubectl get ds --namespace=kube-system kube-proxy

# 以minikube为例，查看节点启动的全部k8s组件
$ kubectl get  --namespace=kube-system all 
NAME                                   READY   STATUS    RESTARTS   AGE
pod/coredns-6955765f44-2kk4k           1/1     Running   12         30d
pod/coredns-6955765f44-z5dxj           1/1     Running   12         30d
pod/etcd-minikube                      1/1     Running   12         30d
pod/kube-addon-manager-minikube        1/1     Running   12         30d
pod/kube-apiserver-minikube            1/1     Running   12         30d
pod/kube-controller-manager-minikube   1/1     Running   12         30d
pod/kube-proxy-884mj                   1/1     Running   14         30d
pod/kube-scheduler-minikube            1/1     Running   22         30d
pod/storage-provisioner                1/1     Running   22         30d

# CoreDNS
NAME               TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)                  AGE
service/kube-dns   ClusterIP   10.96.0.10   <none>        53/UDP,53/TCP,9153/TCP   30d

NAME                        DESIRED   CURRENT   READY   UP-TO-DATE   AVAILABLE   NODE SELECTOR                 AGE
daemonset.apps/kube-proxy   1         1         1       1            1           beta.kubernetes.io/os=linux   30d

NAME                      READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/coredns   2/2     2            2           30d

NAME                                 DESIRED   CURRENT   READY   AGE
replicaset.apps/coredns-6955765f44   2         2         2       30d


## DEBUG
# 登录到pod上
$ kubectl logs <pod-name>
$ kubectl exec -it <pod-name> -- <cmd>
$ kubectl attach -it <pod-name>
# 注意pod:后面的文件路径必须去掉'/'
$ kubectl cp xecho-67c74f4587-clfmk:etc/hostname 1.txt 
# 通过master开一个隧道链接到Pod
$ kubectl port-forward <pod-name>|service/<srv-name> <local-port>:<pod-port>

# Rank, 必须安装heapster
$ kubectl top pod
$ kubectl node pod

# 创建Pod
$ kubectl run --restart=Never -it --image infoblox/dnstools dnstools 
```

## kubectl label

```bash

```



## 参考

- https://github.com/kelseyhightower/kubernetes-the-hard-way/blob/master/docs/13-smoke-test.md
- https://github.com/kelseyhightower/kubernetes-the-hard-way