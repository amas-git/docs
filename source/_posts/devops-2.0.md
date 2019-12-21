devops-2.0

## pets versus cattle meme:

 	1. 把基础设施当宠物
      	1. 认为每个机器都是独特的
      	2. 给每个机器都认真起一个名字，比如db-prod-2
      	3. 应用手工部署到这些机器上，当然出了毛病也要手动修复
 	2. 把基础设施当牲口
      	1. 机器只有编号
      	2. 应用可以自动部署到任意的机器上，当出了问题也不用太担心
      	3. 或者换掉机器在其他机器上自动部署，或是替换掉故障机器上的一部分，比如磁盘等

把基础设施当牲口面临两大问题:

1. 社会问题，去看看凤凰项目这本书
2. 技术问题
   1. 如何自动部署?ansible,chef,pupet?
   2. 如何管理容器?k8s?



## 容器网络栈

 	1. 操作系统层
      	1. iptables
      	2. routing
      	3. IPVLAN
      	4. Linux Namespace
 	2. 容器网络层
      	1. Single-host Bridge Networking
      	2. Multi-host
      	3. IP-per-container
 	3. 容器编排层
      	1. Service Discovery
      	2. CNI
      	3. K8S Network



## SDN (Software-Defined Networking)

简单说就是可以通过程序配置网络，而不是手工配置

- https://www.cisco.com/c/en/us/solutions/software-defined-networking/overview.html
- https://www.sdxcentral.com/networking/sdn/definitions/what-the-definition-of-software-defined-networking-sdn/