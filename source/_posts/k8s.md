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

## 参考

- https://github.com/kelseyhightower/kubernetes-the-hard-way/blob/master/docs/13-smoke-test.md
- https://github.com/kelseyhightower/kubernetes-the-hard-way