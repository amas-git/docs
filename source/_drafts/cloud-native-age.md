

### Cloud Native Age



- [ ] 安装minkube
- [ ] 使用KinD(好像最新的Docker已经包含了Kind)
- [ ] 啊那装k8s-dashboard
- [ ] 红绿发布
- [ ] 滚动升级
- [ ] 金丝雀发布
- [ ] CoreDNS实战
- [ ] Union文件系统实操
- [ ] Selector的使用
- [ ] Pods
- [ ] Services
- [ ] Headless Server
- [ ] Deployments
- [ ] Ingress
- [ ] Egress
- [ ] 存储实战
- [ ] 日志
- [ ] Ingress Controller

## 什么是Cloud Native

```
云原生技术有利于各组织在公有云、私有云和混合云等新型动态环境中，构建和运行可弹性扩展的应用。云原生的代表技术包括容器、服务网格、微服务、不可变基础设施和声明式API。

这些技术能够构建容错性好、易于管理和便于观察的松耦合系统。结合可靠的自动化手段，云原生技术使工程师能够轻松地对系统作出频繁和可预测的重大变更。

云原生计算基金会（CNCF）致力于培育和维护一个厂商中立的开源生态系统，来推广云原生技术。我们通过将最前沿的模式民主化，让这些创新为大众所用。
```

## Cloud Native解决哪些问题？

- Self-healing infrastructure
- Auto-scaling
- High-avaliablility with multi-server failover
- 灵活的storage backends
- Multi-cloud

## 解决方案

- k8s/openshift
- docker swarm
- apache mesos
- aws elastic container service

## k8s 

```
九之七（Seven of Nine）是航海家号上的一名特殊成员——她没有正式的官衔，但却对航海家号非常重要。航海家号曾遇到多次的极端险境，是在九之七的帮助之下才得以脱险
```

管理一个集群（可以使物理服务器，虚拟机，或者他们的混合）

尽全力保证被管理的服务按照描述的方式运行

### k8s master : contrller airplane

- k8s API
- etcd: 用于存储各种
- kube-scheduler: 负责管理节点，实例化各种服务
- controller-manager

### k8s node

- kubelet: 控制crt(container runtime)

- kube-proxy

  

### k8s dns (Core DNS)



### Pods = logical host

> Q: 一个Pod中包含了多个容器，那这些容器是被部署到同一个节点了么？



### 使用kubectl获取集群信息

````sh 
$ kubectl cluster-info         # 一般的信息
$ kubectl cluster-info dump    # 非常详细的集群信息
$ kubectl config view          # 集群配置信息
````



## k8s生态

- 谷歌GCE
- 亚马逊EKS
- 微软AKS
- 红帽OpenShift
- IBM Cloud Kubernetes

## 容器

### 容器技术发展历史

### CGROUP

### Union File System



## 实战

### 安装

#### Minkube

```sh
$ minikube dashboard  # 打开dashboard
$ minikube ip         # minikube的clusterIP
```



#### KinD (k8s in docker)

#### K3S

#### K8S in PI



## 日志

- Aggregators 

- Forwarders

| 解决方案   | 开发语言 | 初始内存 | 依赖     | 性能 | 插件  |
| ---------- | -------- | -------- | -------- | ---- | ----- |
| Fluentd    | C/Rubby  | 40M      | 一些Gems | 高   | 1000+ |
| Fluent Bit | C        | 650K     | 无       | 高   | 70+   |
| Filebeat   | Go       | 50~200M  | 无       | 高   |       |
| Logstash   | Java     |          |          | 低   |       |
|            |          |          |          |      |       |

- [FluentBit官方与Fluentd的特性对比](https://docs.fluentbit.io/manual/about/fluentd-and-fluent-bit)



### 什么是Deployments?

> - 重建（create）: 简单粗暴， 会有DownTime, 停止A, 启动B
> - 滚动升级(Ramped)：逐步用B换掉A
>   - k8s通过readinessProb来检测实例是不是启来了，以便完成切换
>   - 当服务需要支持HPA时，这是唯一可以使用的发布方式么？
>   - 滚动升级需要冗余一部分资源，如何规划？
> - 蓝绿发布：
>   - 系统中总是存在两个版本A和B, A用于提供服务，B版本迭代后经过测试后，将全部流量切给B, 此时A变成了可以继续开发和迭代的环境
>   - K8S用命名空间解决？
>   - 因为两个版本存在，同步数据是比较关键的问题
> - 金丝雀发布：（或Testing in Production System）
>   - 将一部分生产环境的流量切到B版本上
> - A/B测试：
> - 影子流量：复制部分A的流量到B, B不影响A
>
> 参考： https://thenewstack.io/deployment-strategies/

- 发布策略:
  - recreate
  - rolling update



### 什么是Ingress

作为一个服务，可以通过nodeIp + port方式向外部提供，因为有扩展的需求，那么势必会引入LoadBalancer, 

同时也会出于安全的目的将协议由http升级为https, 这些改动如果发声在应用层会让每个服务实现非常重复的功能，这样也会使得证书的管理变得没有统一标准，所以在实践中通常不会这么做。这些事情交给Ingress来完成。

服务本身只要关注业务逻辑就好。

k8s的Ingress controller可以感知到Ingress资源的变化，然后部署相应的Pod



### 什么是GitOps

- https://www.weave.works/

### KinD

- https://kind.sigs.k8s.io/

```sh
# 要求go 1.11+
$ go get sigs.k8s.io/kind@v0.9.0
$ kind
kind creates and manages local Kubernetes clusters using Docker container 'nodes'

Usage:
  kind [command]

Available Commands:
  build       Build one of [node-image]
  completion  Output shell completion code for the specified shell (bash, zsh or fish)
  create      Creates one of [cluster]
  delete      Deletes one of [cluster]
  export      Exports one of [kubeconfig, logs]
  get         Gets one of [clusters, nodes, kubeconfig]
  help        Help about any command
  load        Loads images into nodes
  version     Prints the kind CLI version

Flags:
  -h, --help              help for kind
      --loglevel string   DEPRECATED: see -v instead
  -q, --quiet             silence all stderr output
  -v, --verbosity int32   info log verbosity
      --version           version for kind

Use "kind [command] --help" for more information about a command.
```



## 多个k8s集群管理: kubefed

https://github.com/kubernetes-sigs/kubefed

## IaC : Infrustrate as code

把基础设施用代码描述出来，目前有两种思潮：
  -   declarative (defining the desired outcome) ：描述期望达到的状态
  -  imperative (defining the process of provisioning.) ：描述构建架构的过程



> 什么是Infrastructure:
>
> - vm
> - loadbanlancer
> - rounting
> - firewall
> - storage
> - switching
> - cdn
> - job schedualing queue

## 公司

- HashCorp

#### Terraform

```groovy
terraform {
  required_providers {
    azurerm = "=1.41.0"
  }
}


```



```json
provider "aws" {
  region = "us-east-1"
  access_key = ""
  secret_key = "" // 此处可以采用environment variables or AWS shared credentials(https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-configure.html#cli-config-files).  
}

// 10000~99999的随机整数
resource "random_integer" "rand" {
  min = 10000
  max = 99999
}

resource "aws_s3_bucket" "bucket" {
  name = "unique-name-${random_integer.rand.result}"
}
```

> 有关云平台的access_key和secret_key实际生产环境中不要写进terraform文件中, 可以用云平台的工具在本地管理这些, terraform直接使用即可

别名：

```json
provider "aws" {
  region = "us-east-1"
}

provider "aws" {
  region = "us-wast-1"
  alias  = "west"  
}

resource "aws_ec2_instance" "example_west" {
  name = "instance_west"
  provider = aws.west 
}

```



#### Terraform的环境变量

- TF_PLUGIN_CACHE_DIR: 插件缓存目录

#### Provisioners

尽量不要使用，这使得你的IaC不那么优雅

> Don’t use provisioners unless there is absolutely no other way to accomplish your goal.

通常有几种情况下不得不用：

- 将一些数据上传到VM
- 给VM安装配置管理工具

remote-exe provisioners : 可以通过ssh链接目标机器，执行命令

类似还有Chef, Puppet,Salt等配置管理工具的provisoners

local-exec provisioners可以与本机打交道



> 通常provisoner的用处在于构建资源, 一旦所期望的资源构建完毕,则可以干掉.
>
> 如果资源需要被更改,则不能
>
> provisoner可能会执行失败, 这样就没法构建出你所需要的资源
>
> 另外,可能同样一个resource, 可以指定多个provisoner来构建, 比如一个失败了用另一个之类的

总之, 一旦使用provisoner, 那么你将不得不沉入到具体如何创建资源的细节之中,因此引来的问题也要一一处理,因此你最好别用

### 什么是Terraform Graph

Terraform将资源和之间的依赖抽象为图, 每个节点都是唯一的资源,节点之间的边也可能是资源

```sh
# 生成graph, 输出结果是dot文件, 可以通过graphviz转化为图片
$ terraform graph > infra.dot
```



### Terraform实战

```
$ tree .
.
├── base
├── devl
├── prod
├── README.md
└── stag

```

配置文件为.tf或.tf.json后缀

- .tf: 人类友好, 文件格式参考https://github.com/hashicorp/hcl
- .tf.json:  机器友好



```json
provider "aws" {
  region = "us-east-1"
  access_key = ""
  secret_key = "" // 此处可以采用environment variables or AWS shared credentials(https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-configure.html#cli-config-files).  
}
```



定义实例资源:

```json
resource "aws_instance" "base" {
    ami = "ami-19491001"
    instance_type = "t1.micro"
}
```



定义好provider和资源之后, 我们接下来就让terraform创建资源, 这包含两个步骤:

1. plan: infra面临哪些改动?
2. apply: 实际进行改动



> terraform plan的结果可以输出到文件, 这个文件本身也是可以作为指令让terraform执行的.

```sh
$ terraform plan
$ terraform plan -out base-`date +'%s'`.plan # 保存plan
$ terraform apply *.plan                     # 执行plan
$ terraform plan -target ${target}           # 也可以plan某个具体的资源(target)
```

如果plan执行失败, 对应的资源就会被标记为tainted状态, 再次执行的时候, terraform会尝试重建该资源以





> terraform执行操作的时候只会读取当前目录下的.tf, .tf.json文件,不会递归执行子目录下的配置文件



```
+   : 增加资源
-   : 删除资源
-/+ : 先删除后怎加
~   : 修改资源

<computed> : 表示资源正在被创建,还不知道具体的信息
```



plan后,经过你的确认, 就可以进行apply, 实际操作infra了.



操作完毕后, infra的状态会保存在当前目录下的terraform.tfstate文件中, 因为这个文件非常重要,一但丢失就不得不重新构建一遍, 所以也会自动建立一个terraform.tfstate.backup的备份,  你可能想到要把这放到git里, 但是目前最好别这么做,因为这文件里有密码.



> 注意: 一旦你打算用terraform管理所有的事情,那么目前最好一条路走到黑, 因为你如果手动修改infra, 或者使用其它工具修改infra都会导致状态不同步. 我认为这可能也是terraform需要解决的问题, 如何能够优雅的同步改动

```sh
# 查看infra状态
$ terraform show
```





```sh
# 检验,格式化配置文件
$ terraform fmt 
```



销毁资源

```sh
$ terraform destory
$ terraform destory -target ${target}       # 销毁某个资源
$ terraform plan -destory -out destory.plan # 输出销毁计划
```



### Terraform变量

```json
variable ${name} {
  type        = "string|map"
  description = ""
  default     = ""
}

variable map {
  default = {
    key1 = "value1"
    key2 = "value2"
  }
}

variable array {
  default = [ "value1", "value2" ]
}
```



定义变量

```sh
# 环境变量
TF_VAR_${name}=${value}

# 命令行参数传入
$ terraform plan -var 'key1=value1' -var 'key2=value2' 

# 变量文件
$ terraform plan -var-file *.tfvar
```

### Terraform Modules

模块是用来封装可重用的配置

一个包含了.tf/.tf.json配置的目录就是一个模块



云服务现在已经变得非常复杂了, 你可以使用一些现成的封装模块来简化工作, 一下是一些有用的模块:

- https://github.com/terraform-aws-modules/terraform-aws-vpc 
- 可以使用这个工具画一个真正的架构图: https://www.cloudcraft.co/#Create-your-cloud
- 教材代码仓库: https://github.com/turnbullpress/tfb-cod



#### Terraform之Provisioner

这里主要解决的问题是Terraform如何与Ansible, Chef, Puppet这些工具配合,对资源进行更进一步的配置管理.

- [ ] Terraform + Ansible
- [ ] 将.tfstate存到s3里

#### Terraform的状态

因为terraform在本地维护infra的状态, 你可以将.tfstate这存到etcd/artifactory/consul/s3中, 这样团队协同的时候就可以获取一致的状态, 但这样如果大家没商量好同时提交了数据, 没有锁管理很可能导致数据冲突.



```sh
# 从infra同步状态到.tstate文件
$ terraform refresh
```



```sh
# 初始化时设置state使用的后端
$ terraform init -backend-config=s3_backend
```



## 参考

- [CNCF所定义的CloudNative](https://github.com/cncf/toc/blob/master/DEFINITION.md)
- nsenter命令的使用
- https://www.terraform.io/intro/index.html#infrastructure-as-code
- [PI k8s Cluster](https://www.pidramble.com/)



![](https://miro.medium.com/max/957/1*bosOS65vEyyF7Fu-ldjXtA.jpeg)

> ### 啥是AMI?
>
> Amazon Machine Image:
>
> 什么是 AMI : https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/AMIs.html
>
> 1. 一个AMI可以创建多个Instance
> 2. 



> 啥是: Amazon Elastic IP Address?
>
> https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/db_instance