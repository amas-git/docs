

# Docker

## Docker的技术背景

- Linux LXC

- Cgroup

- Namespaces (注意没有使用UID namespace)

  - PID namespaces
  - NET namespaces
  - IPC namespaces
  - MNT namespaces
  - UTS namespace

- FileSystem

  - BtrFS

    - 优点: 支持写时复制适合容器
    - 缺点:
      - 不支持PageCache
      - 不支持SELinux
      - 尚不被认为可以稳定运行在生产环境

  - AUFS

    - Docker支持的文件系统，很多Linux版本并不支持，常用于Ubuntu

  - VFS

  - DeviceMapper

  - OverlayFS


## containerd

DockerEngine使用containerd管理容器，containerd与具体的操作系统打交道，利用操作系统的特性完成容器的生命周期管理。containerd属于CNCF发起的开源项目Open Container Initiative (OCI)中的一部分。

## 安装

在linux上安装后会建立一个docker组， 为了方便可以将当前用户加入到docker组里:

```sh
$ sudo usermod -a -G dockder <user-name>
```



## Hello World

```bash
# 可以将当前用户加入到docker组中，省去很多sudo
$ useradd -aG docker $USER

$ docker run busybox echo hello
hello

# 进入cointainer
$ docker run -it busybox
[ /]#

# 简单的计时器
$ docker run busybox /bin/sh -c 'while true; do echo $(date); sleep 1; done' 
Mon Dec 30 15:40:51 UTC 2019
Mon Dec 30 15:40:52 UTC 2019
# <CTRL-C>
$ docker run -d busybox /bin/sh -c 'while true; do echo $(date); sleep 1; done'
5e968c81b422f5ba8305af2545c6e1ec1e1df2b9602fa1e087642aa134412cff

# 查看
$ docker ps
CONTAINER ID        IMAGE               COMMAND                  CREATED              STATUS              PORTS               NAMES
5e968c81b422        busybox             "/bin/sh -c 'while t…"   About a minute ago   Up About a minute                       clever_davinci

# 查看容器运行的日志, 通过容器ID或是随机生成的容器名都可以
$ docker logs 5e968c81b422
$ docker logs clever_davinci

# 结束容器
$ docker kill clever_davinci

```

>
>
>## docker run 和 docker container run
>
>```
>Prior to docker 1.13 the docker run command was only available. The cli commands were then refactored to have the form docker COMMAND SUBCOMMAND, where in this case the COMMAND is container and the SUBCOMMAND is run. This was done to have a more inituitive grouping of commands since the number of commands at the time has grown substantially.
>
>You can read more under CLI restructured.
>
>In short, use docker container run as it is the more modern way to run a container.
>```
>
>

### 运行一个简单的echo服务器

```bash
$ docker run -d -p 8888:80 busybox nc -lk -p 80 -e cat
b91e3d5d7c852aa733a0f763437a4bbe2fb863a59b9cb676da0b53349cfbfc37
$ docker ps
CONTAINER ID        IMAGE               COMMAND                 CREATED             STATUS              PORTS                  NAMES
b91e3d5d7c85        busybox             "nc -lk -p 80 -e cat"   51 seconds ago      Up 50 seconds       0.0.0.0:8888->80/tcp   infallible_hoover
# 注意: 0.0.0.0:8888->80/tcp 表示主机端口8888的tcp数据将被forward到容器的80端口
$ busybox nc localhost 8888
hello
hello
```



### 通过HTTPs(mTSL)访问docker registry

```bash
# 1. 初始化CA serial文件
$ echo 01 > sa.srl

# 2. 创建CA的公私钥对
$ openssl genrsa -des3 -out ca-key.pem 2048
$ tree
.
├── ca-key.pem
└── sa.srl

# 创建证书
$ openssl req -new -x509 -days 365 -key ca-key.pem -out ca.pem 
$ tree
.
├── ca-key.pem
├── ca.pem
└── sa.srl

# 3. 创建ServerKey和CertificateSigningRequest(CSR)
$ openssl genrsa -des3 -out server-key.pem 2048
$ openssl req -subj ‘/CN=<hostname here>’ -new -key server-key.pem -out server.csr

# 4. 用ServerKey给CA签名
$ openssl x509 -req -days 365 -in server.csr -CA ca.pem -CAkey ca-key.pem -out server-cert.pem

# 5. 创建Client证书
$ openssl genrsa -des3 -out key.pem 2048
$ openssl req -subj ‘/CN=<hostname here>’ -new -key key.pem -out client.csr

# 6. 创建ExtensionConfig文件
$ echo extendedKeyUsage = clientAuth > extfile.cnf

# 7. 给Key签名
$ openssl x509 -req -days 365 -in client.csr -CA ca.pem -CAkey ca-key.pem -out cert.pem -extfile extfile.cnf

# 8. 删除passphrase
$ openssl rsa -in server-key.pem -out server-key.pem
$ openssl rsa -in key.pem -out key.pem

# 9. 重启dockerd(不同操作系统可能会用到不同的方法)
$ sudo service docker.io stop
$ docker -d —tlsverify —tlscacert=ca.pem —tlscert=server-cert.pem —tlskey=server-key.pem -H=0.0.0.0:2376

# 10. 测试docker客户端
$ docker —tlsverify —tlscacert=ca.pem —tlscert=cert.pem —tlskey=key.pem -H=172.31.1.21:2376 version
```



## 容器运行

### docker  run 

#### --link container:alias

可以将容器之间处于同一个局域网，实际上它是通过配置hosts文件来实现的。

```bash
$ docker run -it --name arch1 -d base/archlinux
$ docker run -it --link arch2 base/archlinux cat ping arch1
PING arch1 (172.17.0.3) 56(84) bytes of data.
64 bytes from arch1 (172.17.0.3): icmp_seq=1 ttl=64 time=0.176 ms
...
```

通过link可以将几个容器置于一个局域网里。容器彼此之间可使用网络联通。

```bash
$ docker run --rm --name srv1 -h srv1.com -td base/archlinux
$ docker ps
CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS              PORTS               NAMES
bebb21b942fb        base/archlinux      "/bin/bash"              3 seconds ago       Up 2 seconds                            srv1

# 我们启动一个新容器，然后link到之前启动的srv1上
$ docker run --rm --link srv1 -it base/archlinux
[root@1c7d75c20a2f /]# ping srv1
[root@1c7d75c20a2f /]# ping srv1.com
[root@1c7d75c20a2f /]# cat /etc/hosts
127.0.0.1	localhost
...
172.17.0.4	srv1 srv1.com
172.17.0.5	1c7d75c20a2f

```

> Docker容器之间的网络连接未来会使用publish services。但是link仍然是支持的。

#### -[-d]etach: 以后台方式运行容器

#### --rm: 无痕运行容器

容器运行结束后自动删除

#### --name $name: 给容器命名



#### -[-e]nv $key=value: 设置容器运行时环境变量

```sh
# 多个环境变量继续追加-e即可
$ docker container run --rm --env MSG=hello busybox env 
...
MSG=hello
...
```



#### -[-p]ublish 

	- -P: 将容器的端口随映射到主机的一个随机端口上
	- 

#### -v host-dir:container-dir

让我来创建一个卷：

```bash
$ docker run --name VOL1 -v /data -it base/archlinux
[root@5a6058eb6cd9 /] # cd /data
# VOL1容器启动后多了一个/data目录

$ docker inspect VOL1
[{volume 20a8f6...  /var/lib/docker/volumes/20a8f6.../_data /data local  true }]
# 1. /var/lib/docker是docker内部虚拟机中目录
# 2. /data是容器中的目录

# 我们在本地先构造些数据
$ mkdir host_data
$ echo "hello" > host_data/message.db

# 让后我们将host_data目录作为容器的/data卷，此时容器已经可以访问到我们构造的数据
$ docker run --rm -v $(pwd)/host_data:/data base/archlinux cat /data/messagedb
hello

# 我们再启动一个容器追加一些数据
$ docker run --rm -v $(pwd)/host_data:/data base/archlinux touch /data/hello
# 容器虽然在执行命令后酒消失了，但是它改变了本地数据
$ ls host_data
messagedb hello

```



#### -e name=value

```bash
$ docker run -e "MSG=hello" -e "VER=1.0"  base/archlinux env
```



#### -it 或 --[i]nteractive -[-t]ty

``` bash
$ docker container run busybox
/ # 
```



## 工作状态查询

### docker ps

```bash
$ docker ps
CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS              PORTS               NAMES
bf98c46e26d5        base/archlinux      "/bin/bash"         6 seconds ago       Up 5 seconds                            youthful_wilson
```

|              |                                                          |
| ------------ | -------------------------------------------------------- |
| CONTAINER ID | 容器ID                                                   |
| IMAGE        | 镜像名                                                   |
| STATUS       |                                                          |
| PORTS        | 端口映射                                                 |
| NAMES        | 容器名，可用--name指定，如果不指定那么系统会自动生成一个 |
|              |                                                          |

一旦你退出容器，这个容器其实会作为历史保存起来

```bash
# 查看过去运行的所有容器
$ docker ps -a
...

# 获取已经停止的容器ID
$ docker ps -aq -f status=exited

# 获得最近使用的容器ID
$ docker ps -lq
```



### docker inspect

```bash
# docker inspect --format {{<json-path>}} [container-id|container-name]

$ docker run -it -h amas.org
$ docker inspect  --format {{.Config.Hostname}} youthful_wilson
amas.org
```



### docker diff

进入容器后, 可以通过diff命令查看镜像发生的变化, 你可以清楚的看到每一步操作对文件系统的改变.

```bash
$ docker run -it --name hello busybox sh
/ # touch /tmp/hello
/ #

# 观察image发生了那些变化
$ docker diff hello
C /root
A /root/.ash_history
C /tmp
A /tmp/hello
```



### docker logs

```bash
$ docker run -it --name hello busybox sh
/ # echo hello
hello
/ # date
Mon Jan  6 20:36:16 UTC 2020

# 查看容器的stdout
$ docker logs hello
/ # echo hello
hello
/ # date
Mon Jan  6 20:36:16 UTC 2020
```



### docker port

### docker top

### docker rm

```sh
# 强制删除所有容器（包括正在运行的）
$ docker container rm --force $(docker container ls --all --quiet)
$ docker container rm -f $container
```



### docker stats $containerId
查看容器的系统资源使用情况

```zsh
# 获取全部容器的资源使用情况
$ docker stats $(docker inspect -f {{.Name}} $(docker ps -q))
```



### 进入正在运行的容器

```sh
$ docker container exec -it $container_id sh
```




## 容器构建

### docker build

> docker build -t name:tag
>
> docker tag name:tag user/name:tag



### docker commit

```bash
$ docker run -h hello -it --name hello base/archlinux
[root#hello] pacman -Sy
...
[root#hello] pacman -S zsh

# 
$ docker commit hello test/hello
sha256:2583d367b988ab8bc92afc4eb718fae3cf9b870dd95483879c9adbe98cd31823
$ docker run test/hello /bin/zsh -c "print hello zsh"
hello zsh
```



### docker export

导出文件系统，一些Docker的元数据则不会导出。

### docker history

我们可以用docker history查看容器是如何一步步被构建出来的

```bash
$ docker history busybox:latest
IMAGE               CREATED             CREATED BY                                      SIZE                COMMENT
b534869c81f0        4 weeks ago         /bin/sh -c #(nop)  CMD ["sh"]                   0B                  
<missing>           4 weeks ago         /bin/sh -c #(nop) ADD file:884f543fc51111835…   1.22MB  

# 我们可以知道busybox这个image有2层
# <missing> 这个是通过ADD指令建立的
# b534869c81f0 是通过CMD建立的


# 既然有history,那么我们理论上就可以回到任何一个history的image里，方法是用tag命令打个标签就好
$ docker tag <IMAGE> <IMAGE>:<TAG>
```



### docker images

- 搜索image: https://registry.hub.Docker.com/

### docker image

```sh
# 模糊搜索
$ docker image ls 'x*'

# 查看历史
$ docker image history $target
```



### docker import

### docker load

### docker rmi

删除镜像

### docker save

保存一个镜像, 保存好的镜像可以拷贝到其他的机器上再通过**docker load**命令加载进去

```bash
$ docker save -o image.tgz name:tag
$ docker load image.tgz
```

### docker tag

## Docker Registry

```sh
$ docker info | grep Registry 
 Registry: https://index.docker.io/v1/
$ docker container run -d -p 5000:5000 --restart always registry
$ docker image push $registry:$port/$id/$image:$version
```



Image的三个类型:

- Verified Publishers: 注入微软，Google等在DockerHub上认证过的
- Official Images: 一般是开源项目，通常经过安全扫描并且经常更新
- Golden Images: 使用OfficalImages作为base的Images, 在一个组织内或项目内通常需要对OfficalImages进一步配置，得到一个可以在项目内使用的base, 这个叫做GoldenImages, 通常你可以将这些images保存到golden路径下，这样对于CI可以进行扫描，保证所有的Images都基于GoldenImages来构建

### docker pull

### docker push

### docker login

```sh
$ docker login --username $dockerId
```



### docker logout

### docker seaarch 

简单搜索image

```bash
$ docker search ubuntu
```



### docker service

```

```

## 容器

- 基于AUFS, 最多127层
- 每一条ADD, COPY, RUN命令都会增加一层
- 状态:
  - created: 新创建未运行过
  - restarting
  - running
  - paused
  - exited: 曾经运行过

## 容器存储

Docker Volume是一个存储单元，你可以认为它是一个给容器准备的U盘。Volumes有自己的生命周期，可以attached到容器之上。

有两种方法可以定义Volumes:

- 创建一个Volume然后attached到容器上
- 使用VOLUME $target-dir在Dockerfile中定义Volumes





```sh
$ docker volume ls
$ docker container inspect --format '{{.Mounts}}' $target
```



```dockerfile
FROM busybox
VOLUME /src
```

```sh
$ docker build -t vol .
$ docker run --rm -it -name voltest vol
$ docker inspect --format '{{.Mounts}}'
[{volume a77d2b5bddb0fb16ecbfc191381d04604f51de40c24eeaa6ae4e457b0faf422c /var/lib/docker/volumes/a77d2b5bddb0fb16ecbfc191381d04604f51de40c24eeaa6ae4e457b0faf422c/_data /src local  true }]
# volume对应本机路径/var/lib/docker/volumes/a77d2b5.../_data

$ docker volume ls
DRIVER              VOLUME NAME
local               a77d2b5bddb0...
```



````sh
# 手动创建一个volume
$ docker volume create src
# 在Host机上对应的volume目录中创建一个文件
$ sudo touch /var/lib/docker/volumes/src/_data/hello
# 在新的容器中使用这个volume
$ docker run -it --rm -v src:/src  busybox ls /src 
hello
````



使用bind mount是一种容器与Host主机共享文件系统的方式，需要注意的是这种方式虽然方便，但存在安全隐患，如果容器不需要写，那么可以用只读的方式。

```sh
# 读写方式
$ docker container run --mount type=bind,source=$source,target=$target $image

# 只读方式
$ docker container run --mount type=bind,source=$source,target=$target,readonly $image
```

使用Host存储另外一个需要了解的，假如在image中有同样的目录，然后同时mount了同样的主机目录，结果会如何？

1. Host 目录的内容替换容器中的目录？
2. 容器中的目录？
3. Host和容器目录合并？

使用Host存储，在linux和windows系统上会有差异，比如windows不支持单个文件的bind mount

虽然同是一个路径，可以背后的文件系统有差异，并非一定支持所有的操作，因此在这一点上会使得同样的dockerfile在不同host机上运行效果不同，比如有些文件系统fat32，azure file不支持

>存储层级： APP -> 可写层 -> Mount层 -> 镜像层

1. 可写层（writable layer)：每个容器有唯一的可写层，容器生命周期结束后回收
2. Local Bind Mounts层：用于Host机与容器间共享数据
3. Distributed Bind Mounts层： 用户网络存储与容器间共享数据
4. Volumes  Mounts：容器与docker管理的存储（Volumes）之间共享数据
5. Images层：只读数据

### 容器存储是如何工作的？



## Docker网络

Docker内置DNS系统，如果名字是容器的，那么返回容器IP, 否则交给操作系统的DNS进行解析？

## 镜像是怎么构建出来的？

```
Image = Dockerfile + BuildContext
```

- Dockerfile就是一个文件，里面是一些指令，描述如何构建出你想要的Image
- BuildContext就是一些本地文件，和Dockerfile放在一起，在Dockerfile中使用COPY或ADD放入镜像里

```bash
$ docker build -t .
# . 就是BuildContext是当前目录
# BuildContext也可以是URL
```





Base Image

- 首先我们要选择一个基础镜像
- 理想情况是，不去创建新的镜像，只是把软件的配置按需写好复制到已有的镜像中

下面是一些不错的BaseImage:

- alpine: 一个非常精简的linux系统
- ubuntu
- scratch
  - 空文件系统
  - 你需要将可执行程序拷贝进去并作为入口
  - 不支持动态库，不支持外部命令连shell都木有
- phusion/baseimage-docker
  - 没有init service
  - cron默认会被启动
  - 默认不安装或者不启动ssh

当build镜像的时候，BaseImage会先从缓存中获得，如果没有才去pull并且给缓存到本地. 因此BaseImage要记得更新，这样可以有效避免安全问题。也可以**docker build --pull**l来强制拉取新镜像, 



```
一些有争议的观点：
对于启动服务（init service）：一个容器一个程序，最好一个进程，一个进程也就不需要什么启动服务，但是砍掉服务init服务又会造成僵尸进程

是不是
```



准备好我们去写一个Dockerfile

```
FROM
```





### Dockerfile

```dockerfile
FROM
MAINTAINER # 过时了，可以用LABEL代替
```



> 注意: Docker Image 最多可以支持147层

cowsay/Dockerfile:

```
FROM base/archlinux
RUN  pacman -Sy --noconfirm &&  pacman -S --noconfirm cowsay
```

```sh
$ mkdir cowsay ; cd cowsay
$ docker build -t test/cowsay .
...
$ docker run test/cowsay cowsay hello
 _______
< hello >
 -------
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||
```

我们稍加修改:

```sh
FROM base/archlinux
RUN  pacman -Sy --noconfirm &&  pacman -S --noconfirm cowsay
ENTRYPOINT ["cowsay"]
```

```
$ docker built -t test/cowsay
...
$ docker run test/cowsay hello
 _______
< hello >
 -------
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||
```



```
$ ls
Dockerfile hello.sh
$ cat hello.sh
#!/bin/bash
cowsay HELLO

$ cat Dockerfile
FROM base/archlinux
RUN  pacman -Sy --noconfirm &&  pacman -S --noconfirm cowsay
COPY hello.sh /
ENTRYPOINT ["/hello.sh"]

$ docker build -t test/cowsay
...

$ docker run test/cowsay
 _______
< HELLO >
 -------
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||
```

### .dockerignore

避免不必要的文件被装进Image中。



玩玩redis:

```bash
# 获取redis
$ docker pull redis

# 启动一个名为myredis的容器并在后台运行 
$ docker run -it --name myredis -d redis

# --link <running-container>:<current-container>
$ docker run --rm -it --link myredis:rediss redis /bin/bash
root@a05fd98b73dc:/data# redis-cli -h redis -p 6379
```



```bash
$ docker run --rm -it redis cat /etc/hosts
127.0.0.1	localhost
::1	localhost ip6-localhost ip6-loopback
fe00::0	ip6-localnet
ff00::0	ip6-mcastprefix
ff02::1	ip6-allnodes
ff02::2	ip6-allrouters
172.17.0.3	6fe3c188ce41

$ docker run --rm -it --link myredis:redis redis cat /etc/host
27.0.0.1	localhost
::1	localhost ip6-localhost ip6-loopback
fe00::0	ip6-localnet
ff00::0	ip6-mcastprefix
ff02::1	ip6-allnodes
ff02::2	ip6-allrouters
172.17.0.2	redis fbb5e8adbff9 myredis
172.17.0.3	c5b2461df330
```

#### 

将host的目录挂载到容器的挂载点，这个用来为容器配置存储 

#### --volumes-from container

### docker network

```sh
# windows版需要手工创建nat网络
$ docker network create nat
$ docker container run --network nat $target
```



### docker build

### Build Context

当使用docker build命令时，当前目录就是Build Context.



#### FROM image:tag

MAINTAINER

```
# 使用这个命令查看
$ docker inspect -f {{.Author}}
```



RUN

#### CMD [cmd, arg1, arg2, ..., argN]

LABEL

#### EXPORSE port [port/protocol..]

```
EXPOSE <port> [<port>/<protocol>...]
```

用来设置容器暴露的端口

```bash
# Host机上的8000端口会推到容器的80端口上
$ docker run --name web1 -p 8000:80 nginx
$ docker port web1
0/tcp -> 0.0.0.0:8000
```

- http://localhost;8000



ENV

ADD

#### COPY src dst

#### COPY ["src", "dst"] 

ENTRPOINT

#### VOLUME 

VOLUME命令之后，你就无法再修改卷的权限，所属用户等等。

USER

#### WORKDIR dir

ARG

ONBUILD

STOPSIGNAL

HEALTHCHECK

SHELL





## 利用Layer Cache加速构建

我们知道在build镜像的时候，Docker是按照Dockerfile中指令的顺序逐步构建的，对于每一个指令产生的中间镜像会尽可能利用缓存，因此如果我们把不太发生变化的指令放到Dockerfile前面，那么意味着可以利用缓存加速构建的过程. 

例如:

```dockerfile
FROM node
ENV  X=1                     # 不变
WORKDIR /app                 # 不变
COPY app.js                  # 变
CMD ["node", "/app/app.js"]  # 不变
```

可以将CMD这种不太发生变化的指令调整到前面

```dockerfile
FROM node
ENV  X=1                     # 不变
CMD ["node", "/app/app.js"]  # 不变
WORKDIR /app                 # 不变
COPY app.js                  # 变
```

## 分部构建（Muilt）

```dockerfile
FROM diamol/base AS build-stage
RUN echo 'Building...' > /build.txt

FROM diamol/base AS test-stage
COPY --from=build-stage /build.txt /build.txt
RUN echo 'Testing...' >> /build.txt

FROM diamol/base
COPY --from=test-stage /build.txt /build.txt
CMD cat /build.txt
```

```dockerfile
FROM diamol/maven AS builder

WORKDIR /usr/src/iotd
COPY pom.xml .
RUN mvn -B dependency:go-offline

COPY . .
RUN mvn package

# app
FROM diamol/openjdk

WORKDIR /app
COPY --from=builder /usr/src/iotd/target/iotd-service-0.1.0.jar .

EXPOSE 80
ENTRYPOINT ["java", "-jar", "/app/iotd-service-0.1.0.jar"]
```



## 底层技术

### cgroups

- 控制内存的使用
- 控制CPU的使用
- 冻结容器
- 解冻容器



### naming space

- 隔离容器

### UFS(Union File System)



### Docker容器的生命周期

![image-20181103002229626](/Users/amas/Library/Application Support/typora-user-images/image-20181103002229626.png)

## 日常使用

> The advantage of containers, DevOps, microservices, and continuous delivery essentially comes down to the idea of a fast feedback loop. By iterating quicker, we can
> develop, test, and validate systems of higher quality in shorter time periods.

### 实战一

```
FROM ubuntu:14.04
MAINTAINER amas<zhoujb.cn@gmail.com>

# Python
RUN apt-get install -y libreadline-gplv2-dev libncursesw5-dev libssl-dev libsqlite3-dev tk-dev libgdbm-dev libc6-dev libbz2-dev
RUN cd /usr/src ; wget https://www.python.org/ftp/python/3.4.3/Python-3.4.3.tgz ; tar xzf Python-3.4.3.tgz ; cd Python-3.4.3 ; ./configure ; make altinstall
# be sure it's 3.4 and not 3.5
RUN ! ls /usr/bin/python3.4 && ls /usr/src/Python-3.4.3/python && cp /usr/src/Python-3.4.3/python /usr/bin/python3.4 ; exit 0
# replace python version to have 3.4.4 as default
RUN rm -f /usr/bin/python
RUN rm -f /usr/bin/python3
RUN ln -s /usr/bin/python3.4 /usr/bin/python
RUN ln -s /usr/bin/python3.4 /usr/bin/python3
# Pip
RUN apt-get install -y python3-pip
RUN pip3 uninstall pep8 ; pip3 install pep8 ; pip3 install --upgrade pep8
```



### 实战二: ELK Stack

Elasticsearch : 搜索引擎

Logstash : 收集处理日志

Kibna: Elasticsearch的前端

1. 如何将Docker的日志发送给Logstash? 使用Logspout



![image-20181103001825545](/Users/amas/Library/Application Support/typora-user-images/image-20181103001825545.png)

```yaml
# docker-compose.yml file
version: '3.7'

services:
  es:
    labels:
      com.example.service: "es"
      com.example.description: "For searching and indexing data"
    image: docker.elastic.co/elasticsearch/elasticsearch:6.4.2
    container_name: E1
    volumes:
      - ./esdata:/usr/share/elasticsearch/data/
    ports:
      - "9200:9200"

  logstash:
    labels:
      com.example.service: "logstash"
      com.example.description: "For logging data"
    image: logstash
    container_name: L1
    volumes:
      - ./:/logstash_dir
    command: logstash -f /logstash_dir/logstash.conf
    depends_on:
      - es
    ports:
      - "5959:5959"

  kibana:
    labels:
      com.example.service: "kibana"
      com.example.description: "Data visualisation and for log aggregation"
    image: docker.elastic.co/kibana/kibana:6.4.2
    container_name: K1
    ports:
      - "5601:5601"
    environment:
      - ELASTICSEARCH_URL=http://es:9200
    depends_on:
      - es
```

```
esdata         # es数据目录        
logstash.conf  # logsdash配置文件 
logstash_dir   # logsdash目录 
main.yml       # docker
$ docker-compose -f main.yml up
```







### 优化镜像体积

来制造点大文件：

```
FROM debian:wheezy
RUN dd if=/dev/zero of=/bigfile count=1 bs=50MB
RUN rm /bigfile
```

```
$ docker build -t filetest .
...
$ docker images filetest
REPOSITORY TAG IMAGE ID CREATED VIRTUAL SIZE filetest latest e2a98279a101 8 seconds ago 135 MB

# 看看这135MB是怎么来的
$ docker history filetest

```

```
FROM debian:wheezy
RUN dd if=/dev/zero of=/bigfile count=1 bs=50MB && rm /bigfile
```

优化的思路：

1. 尽量减少镜像的层级(Layer)
   - docker export $(docker ps -lq) | docker import 
2. 在同一个层级中最后要删除无用文件，比如缓存，安装包等等。

```
Minimize the number of layers
In older versions of Docker, it was important that you minimized the number of layers in your images to ensure they were performant. The following features were added to reduce this limitation:

In Docker 1.10 and higher, only the instructions RUN, COPY, ADD create layers. Other instructions create temporary intermediate images, and do not directly increase the size of the build.

In Docker 17.05 and higher, you can do multi-stage builds and only copy the artifacts you need into the final image. This allows you to include tools and debug information in your intermediate build stages without increasing the size of the final image
```



### 持续集成和测试

```
Using Containers for Fast Testing
All tests, and in particular unit tests, need to run quickly in order to encourage devel‐ opers to run them often without getting stuck waiting on results. Containers repre‐ sent a fast way to boot a clean and isolated environment, which can be useful when dealing with tests that mutate their environment. For example, imagine you have a suite of tests that make use of a service3 that has been prepopulated with some test data. Each test that uses the service is likely to mutate the data in some way, either adding, removing, or modifying data. One way to write the tests is to have each test attempt to clean up the data after running, but this is problematic; if a test (or the clean-up) fails, it will pollute the test data for all following tests, making the source of the failure difficult to diagnose and requiring knowledge of the service being tested (it is no longer a black box). An alternative is to destroy the service after each test and start with a fresh one for each test. Using VMs for this purpose would be far too slow, but it is achievable with containers.
Another area of testing where containers shine is running services in different envi‐ ronments/configurations. If your software has to run across a range of Linux distribu
tions with different databases installed, set up an image for each configuration and you can fly through your tests. The caveat of this approach is that it won’t take into account kernel differences between distributions.
```



```
Testing and Microservices
If you’re using Docker, there’s a good chance you’ve also adopted a microservice architecture. When testing a microservice architecture, you will find that there are more levels of testing that are possible, and it is up to you to decide how and what to test. A basic framework might consist of:
Unit tests
Each service7 should have a comprehensive set of unit tests associated with it. Unit tests should only test small, isolated pieces of functionality. You may use test doubles to replace dependencies on other services. Due to the number of tests, it is important that they run as quickly as possible to encourage frequent testing and avoid developers waiting on results. Unit tests should make up the largest proportion of tests in your system.
Component tests
These can be on the level of testing the external interface of individual services, or on the level of subsystem testing of groups of services. In both cases, you are likely to find you have dependencies on other services, which you may need to replace with test doubles as described earlier. You may also find it useful to expose metrics and logging via your service’s API when testing, but make sure this is kept in a separate namespace (e.g., use a different URL prefix) to your functional API.
End-to-end tests
Tests that ensure the entire system is working. Since these are quite expensive to run (in terms of both resources and time), there should only be a few of these— you really don’t want a situation where it takes hours to run the tests, seriously
7 Normally, there will be one container per service, or multiple containers per service if more resources are needed.
  Hosted CI Solutions | 133
delaying deployments and fixes (consider scheduled runs, which we describe shortly). Some parts of the system may be impossible or prohibitively expensive to run in testing and may still need to be replaced with test doubles (launching nuclear missiles in testing is probably a bad idea). Our identidock test falls under end-to-end testing; the test runs the full system from end to end with no use of test doubles.
In addition, you may want to consider:
Consumer-contract tests
These tests, which are also called consumer-driven contracts, are written by the consumer of a service and primarily define the expected input and output data. They can also cover side effects (changing state) and performance expectations. There should be a separate contract for each consumer of the service. The pri‐ mary benefit of such tests is that it allows the developers of a service to know when they risk breaking compatability with consumers; if a contract test fails, they know to they need to either change their service, or work with the develop‐ ers of the consumer to change the contract.
Integration tests
These are tests to check that the communication channels between each compo‐ nent are working correctly. This sort of testing becomes important in a microser‐ vice architecture where the amount of plumbing and coordination between components is an order of magnitude greater than monolithic architectures. However, you are likely to find that most of your communication channels are covered by your component and end-to-end testing.
Scheduled runs
Since it’s important to keep the CI build fast, there often isn’t enough time to run extensive tests, such as testing against unusual configurations or different plat‐ forms. Instead, these tests can be scheduled to run overnight when there is spare capacity.
Many of these tests can be classified as preregistry and postregistry, depending on whether they occur prior to adding the image to the registry. For example, unit test‐ ing is preregistry: no image should be pushed to the registry if it fails a unit test. The same goes for some consumer contract tests and some component tests. On the other hand, an image will have already been pushed to a registry before it can be end-to- end tested. If a postregistry test fails, there is a question about what to do next. While any new images should not be pushed to production (or should be rolled back if they have already been deployed), the fault may actually be due to other, older images or the interaction between new images. These sort of failures may require a greater level of investigation and thought to handle correctly.
```



- Travis, Wercker, CircleCI, and drone.io. 

### DinD

### 部署

> It is perfectly possible and reasonable to use containers in production
> today.



### 日志和监控

```sh
# 通过p8s监控dockerd的健康状况
$ cat /etc/docker/daemon.json
{ 
        "metrics-addr" : "127.0.0.1:9323",
        "experimental": true
}
$ sudo systemctl restart docker.service 
$ curl http://127.0.0.1:9323/metrics
```





### 健康检查

````dockerfile
HEALTHCHECK [OPTIONS] CMD command
[OPTIONS]
  --interval=DURATION (default 30s)
  --timeout=DURATION (default 30s)
  --retries=N (default 3)
  
# 需要注意，CMD执行后exit code为1表示失败，0表示成功, 所以下面这个声明你就明白为什么要加exit 1
HEALTHCHECK CMD curl --fail http://localhost:9000/guid/ || exit 1  
````

```sh
# 也可以在docker命令行中设置
$ docker run -d --name db --health-cmd "curl --fail http://localhost:8091/pools || exit 1" --health-interval=5s --timeout=3s
```



```sh
$ docker inspect $container | grep FailingStreak
                "FailingStreak": 7,
$ docker inspect $container  | grep Status  
            "Status": "running",
                "Status": "unhealthy",
$ docker inspect --format "{{json .State.Health }}"
$ docker events --filter event=health_status
$ docker container ls 
CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS                      PORTS                  NAMES
0f582d8dd273        hping               "nc -lk -p 80 -e hos…"   47 minutes ago      Up 47 minutes (unhealthy)   0.0.0.0:9111->80/tcp   hping
```

> 注意：
>
> Docker只能按照你的要求进行健康检查，并且记录健康状况，但是并不会采取任何措施，主要原因是需不需要重启，什么时候重起，怎么重启都可能会影响你的业务，并不是应该由Docker来自动处理的

## 、安全

## 容器的局限

### Docker Content Trust



1. 

## 周边工具

- Universal Control Plane (UCP)： https://docs.docker.com/ee/ucp/
- Portainer

### Swarm 

Docker’s clustering solution. Swarm can group together several Docker hosts, allowing the user to treat them as a unified resource. See Chapter 12 for more information. 

### Docker Compose 

通常我们的服务需要由多个容器支撑，如何将多个容器组合到一起提供服务便是DockerCompose所解决的问题。我们可以把这些参数放到一个名为docker-compose.yml的文件中，利用docker-compose up命令来代替原来的启动方法。docker-compose的配置项目的定义规则与Dockerfile中的一致

> 比如: ports : container-port:host-port 等价的命令行参数则是 p host-port:container-port
>
> 当一个参数用来描述host与container的映射关系时，命令行参数总是host在前，而Dockerfile和docker-compose总是container在前. 
>
> 那么为什么会这样设计呢?
>
> 问题在于, docker命令是在宿主机使用的, 所以把宿主机的信息放在前面便于使用, 举个例子, 比如我想在80端口
>
> 上提供服务, 那么在启动的时候只要加上-p 80就可以了, 而EXPOSE命令则更加关注的是容器本身, 所以如果容器
>
> 内部提供的端口是8888, 只要加上EXPOSE [8888]j就可以了. 

```yaml
version: '3'
services:
  web:
    build: .
    ports:
     - "5000:5000"
  redis:
    image: "redis:alpine"
```

	- https://docs.docker.com/compose/

```yaml
version: ''
services:
    $name:
      image: $image
      ports:
        - ${local-port}:${container-expose-port}
      depend_on:
        - ${name1}
        - ${nameN}
      environment:
        - ${key}=${value}
      restart: [always]  
      secrets:
        - source: ${source_name}
          target: ${file}
      healthcheck:
          test: $cmd
          interval: ${time}
          timeout: ${time}
          retries: ${n}
volumes:          
```



```sh
$ docker-compose up
$ docker-compose up -d # --detach, 后台运行
$ docker-compose up -d --scale ${name}=${n} # 运行3个${name}容器
$ docker-compose down  # 停止相关容器，删除相关网络，Volumes
$ docker-compose stop  # 停止容器
$ docker-compose start # 启动容器
$ docker-compose run ${container_name}

```



```sh
$ tree .
.
├── docker-compose.yaml
├── Dockerfile
└── password.txt
$ cat password.txt
hello
```

```dockerfile
FROM busybox
EXPOSE 80
ENTRYPOINT ["nc", "-lk", "-p", "80", "-e", "hostname"]
```

```yaml
version: '3.7'
services:
  ping:
    build: .
    image: ping
    ports: 
      - "8888:80"
    secrets:
        - source: password
          target: /app/config/password
secrets:
  password: 
    file: ./password.txt
```

```
$ docker-compose up -d
$ docker container exec ping_1 cat /app/config/password
hello

```



### 使用compose构建多个环境

Docker Compose可以支持合并，我们可以将基础的配置放到一个文件中，而后将变化部分放到另外一个文件中。



```sh 
# 检测合并后的配置是否合法
$ docker-compose -f docker-compose.yml -f docker-compose-v2.yml config
```



### Docker Machine 

Docker Machine用来辅助快速在各种机器上部署Docker环境.

在Windows和Mac上使用Docker

![Docker Machine on Mac and Windows](https://docs.docker.com/machine/img/machine-mac-win.png)



安装配置管理远程机器上的Docker

![Docker Machine for provisioning multiple systems](https://docs.docker.com/machine/img/provision-use-case.png)



> 通常我们说起Docker都是指Docker Engine
>
> Docker Machine 是用来帮助我们在各种机器上快速安装/管理Docker Engine的



#### 安装

#### 在本地建立一个Docker Machine

```zsh
# 使用virtualbox作为虚拟化引擎
# 执行下面命令之前需要安装virtualbox
# 实际上docke-machine在虚拟机上安装了一个linux: https://github.com/boot2docker/boot2docker
$ docker-machine  --debug create --driver virtualbox default
# 等价于
$ docker-machine create default
...
$ docker-machine ls

# 删除
$ docker-machine rm default


```





```
# 使用ssh登陆容器
$ docker-machine ssh default
```





>Error creating machine: Error in driver during machine creation: This computer doesn't have VT-X/AMD-v enabled. Enabling it in the BIOS is mandatory
>
>碰到这种问题说明CPU的虚拟化特性没有打开,  怎么看你的CPU是不是支持虚拟化技术呢?
>
>#### Intel CPU:
>$ cat /proc/cpuinfo | grep --color vmx
>
>$ lscpu也可以
>
>Virtualization:      VT-x
>
>#### AMD CPU:
>$ cat /proc/cpuinfo | grep --color svm





#### 如何使用SSH登录容器?

方法一:

```zsh
$ docker-machine ssh default
```



方法二:

```zsh
$ VBoxManage list vms
$ VBoxManage showvminfo f3c75c0a-0ea1-4a91-b775-70ad1d111278 | grep ssh
NIC 1 Rule(0):   name = ssh, protocol = tcp, host ip = 127.0.0.1, host port = 40275, guest ip = , guest port = 22
# 用户名: docker
#     密码: tcuser 
$ ssh docker@localhost -p 40275
docker@localhost's password:tcuser
```







Docker Machine installs and configures Docker hosts on local or remote resources. Machine also configures the Docker client, making it easy to swap between environments. See Chapter 9 for an example. 

```
    - `docker-machine config`
    - `docker-machine env`
    - `docker-machine inspect`
    - `docker-machine ip`
    - `docker-machine kill`
    - `docker-machine provision`
    - `docker-machine regenerate-certs`
    - `docker-machine restart`
    - `docker-machine ssh`
    - `docker-machine start`
    - `docker-machine status`
    - `docker-machine stop`
    - `docker-machine upgrade`
    - `docker-machine url`
    
    # docker本身就是虚拟机, 所以你也可以用本地的docker来驱动
    $ create --driver none --url=tcp://50.134.234.20:2376 custombox
    $ docker-machine create --driver digitalocean --digitalocean-access-token xxxxx docker-sandbox
    $ docker-machine create --driver amazonec2 --amazonec2-access-key AKI******* --amazonec2-secret-key 8T93C*******  aws-sandbox
```

- 目前支持的云平台: https://docs.docker.com/machine/drivers/



#### 在远程主机上部署Docker

首先你要确保三件事情:

- 远程主机打开ssh, 设置免密码登录(ssh-copy-id)
- 设置登录用户可以免密码sudo
- 检查防火墙, 确保2376端口可以访问

```zsh
$ sudo echo "${USER}  ALL=(ALL)       NOPASSWD:ALL" >> /etc/sudoer
```

```zsh
# 测试2376端口没有打开是这样的
$ nc -vz 10.60.81.232  2376 
10.60.81.232 2376 (docker-s): No route to host

# 在远程主机上打开2376端口
10.60.81.232$ sudo firewall-cmd --zone=public --add-port=2376/tcp

$ nc -vz 10.60.81.232  2376
10.60.81.232 2376 (docker-s) open
```



做完以上工作, 我们就可以执行下面的命令部署DockerEngine了. 

```zsh
$ docker-machine create \
	--driver generic \
	--generic-ip-address=10.60.81.232 \
	--generic-ssh-key ~/.ssh/id_rsa  \
	--generic-ssh-user=worker \
	d232
..
$  docker-machine ls 
NAME        ACTIVE   DRIVER       STATE     URL                         SWARM   DOCKER        ERRORS
ir232       -        generic      Running   tcp://10.60.81.232:2376             v18.06.1-ce 
```

#### 如何操作远程主机上的Docker?

```

```



### Kitematic 

Kitematic is a Mac OS and Windows GUI for running and managing Docker containers. 

### Docker Trusted Registry

比Local Registry更加便于使用的商业化版本

- Docker Trusted Registry 

- CoreOS Enterprise Registry 

### Local Registry

```zsh
$ docker run -d -p 5000:5000 registry:2
$ docker tag amouat/identidock:0.1 localhost:5000/identidock:0.1
$ docker push localhost:5000/identidock:0.1
```



```
$ docker pull 192.168.1.100:5000/identidock:0.1
Error response from daemon: unable to ping registry endpoint https://192.168.99.100:5000/v0/
v2 ping attempt failed with error: Get https://192.168.99.100:5000/v2/: tls: oversized record received with length 20527
     v1 ping attempt failed with error: Get https://192.168.99.100:5000/v1/_ping:
    tls: oversized record received with length 20527
Here I’ve substituted the IP address of the server for “localhost.” You will get this error whether you pull from a daemon on another machine or on the same machine as the registry.
So what happened? The Docker daemon is refusing to connect to the remote host because it doesn’t have a valid Transport Layer Security (TLS) certificate. The only reason it worked before is because Docker has a special exception for pulling from “localhost” servers. We can fix this issue in one of three ways:
1. Restart each Docker daemon that accesses the registry with the argument -- insecure-registry 192.168.1.100:5000, replacing the address and port as appropriate for your server.
2. Install a signed certificate from a trusted certificate authority on the host, as you would for hosting a website accessed over HTTPS.
3. Install a self-signed certificate on the host and a copy on every Docker daemon that needs to access the registry.
The first option is the easiest, but we won’t consider it here due to the security con‐ cerns. The second option is the best but requires you to obtain a certificate from a trusted certificate authority, which normally has an associated cost. The third option is secure but requires the manual step of copying the certificate to each daemon.
If you want to create your own self-signed certificate, you can use the OpenSSL tool. These steps should be carried out on a machine you want to keep running long term as a registry server. They were tested on an Ubuntu 14.04 VM running on Digital Ocean; there are likely to be differences on other operating systems.
root@reginald:~# mkdir registry_certs
root@reginald:~# openssl req -newkey rsa:4096 -nodes -sha256 \ > -keyout registry_certs/domain.key -x509 -days 365 \
          -out registry_certs/domain.crt
    Generating a 4096 bit RSA private key
    ....................................................++
    ....................................................++
    writing new private key to 'registry_certs/domain.key'
    -----
    
    root@reginald:~# ls registry_certs/
domain.crt domain.key
Creates a x509 self-signed certificate and a 4096-bit RSA private key. The certifi‐ cate is signed with a SHA256 digest and is valid for 365 days. OpenSSL will ask for information, you can input or leave at the default values.
The common name is important; it must match the name you want to access the server on and should not be an IP address (“reginald” is the name of my server).
At the end of this process, we have a certificate file called domain.crt that will be shared with clients and a private key domain.key that must be kept secure and not shared.

Addressing the Registry by IP Address
If you want to use an IP address to reach your registry, things are a little more compli‐ cated. You can’t simply use the IP address as the common name. You need to set up Subject Alternative Names (or SANs) for the IP address or addresses you want to use.
In general, I would advise against this approach. It’s better just to pick a name for your server and make it addressable by the name internally (in the worst case, you can always manually add the server name to /etc/hosts). This is generally easier to set up and doesn’t require retagging of all images should you want to change the IP address.
```

- 

### 作弊手册

建立mkdoc

```bash
#!/bin/bash
#git@gitbj.cmcm.com:CMFinance/docbase.git
repo=$1
host=$(dirname $repo  | cut -d: -f1)
home=$(basename $repo | cut -d. -f1)

echo repo=$repo
echo host=$host
echo "home=${home}"
echo "source=${home}/.docbase"

ssh -o StrictHostKeyChecking=no $host
git clone $repo

source $home/.docbase
echo doc=$doc
cd $home/$doc
source .docbase
mkdocs serve -a 0.0.0.0:8000
```



```dockerfile
FROM ubuntu:16.04 
ENV workdir=/data
RUN apt-get update  -y \
    && apt-get install -y python python-pip \
    && apt-get install -y git \
    && apt-get install -y netcat \
    && pip install mkdocs  \
    && pip install mkdocs-material \
    && mkdir -p ${workdir} \
    && mkdir /root/.ssh

ADD ssh/* /root/.ssh/
ADD start /usr/bin/
WORKDIR ${workdir}
```



防火墙:

```zsh
# centos 检测防火墙是否开启
$ sudo firewall-cmd --state
# 查看当前使用的zone, zone就是防火墙抽象出来的分组
$ firewall-cmd --get-active-zones
public
  interfaces: eth0
trusted
  interfaces: docker0
$ firewall-cmd --get-zones
block dmz drop external home internal public trusted work

# 现在我们想暴露一个端口改怎么办?
# 查看当前系统有哪些端口暴露
$ sudo firewall-cmd --zone=public --list-services
ssh dhcpv6-client
$ sudo firewall-cmd --zone=public --permanent --list-services

# 查看一下有哪些服务可以暴露
$ firewall-cmd --get-services 
... https
$ sudo cat /usr/lib/firewalld/https.xml
<?xml version="1.0" encoding="utf-8"?>
<service>
  <short>Secure WWW (HTTPS)</short>
  <description>HTTPS is ... </description>
  <port protocol="tcp" port="443"/>
</service>

# 测试一下打开http端口
$ firewall-cmd --zone=public --add-service=http
# 如果OK, 你需要永久保存一下这个配置
$ firewall-cmd --zone=public --permanent --add-service=http

# 当然有时候你需要打开的端口并没有对应的xml文件, 这时候可以
$ sudo firewall-cmd --zone=public --add-port=2376/tcp
```



```
apt-get update
apt-get install -y <package>
pip install <package>==<version>

设置locale, 在Dockerfile中指定以下几个环境变量
RUN locale-gen en_US.UTF-8
ENV LC_ALL en_US.UTF-8
ENV LANG en_US.UTF-8
ENV LANGUAGE en_US.UTF-8

```

```dockerfile
FROM python:3.4
RUN groupadd -r uwsgi && useradd -r -g uwsgi uwsgi
RUN pip install Flask==0.10.1 uWSGI==2.0.8 requests==2.5.1 WORKDIR /app
COPY app /app
COPY cmd.sh /
EXPOSE 9090 9191 USER uwsgi
CMD ["/cmd.sh"]
```

```
tRunning Multiple Process in a Container
The majority of containers only run a single process. Where multiple processes are needed, it’s best to run multiple containers and link them together, as we have done in this example.
However, sometimes you really do need to run multiple processes in a single con‐ tainer. In these cases, it’s best to use a process manager such as supervisord or runit to handle starting and monitoring the processes. It is possible to write a simple script to start your processes, but be aware that you will then be responsible for cleaning up the processes and forwarding any signals.
For more information on using supervisord inside containers, see this Docker article.
```

### Kubernetes

- http://kubernetes.io/docs/getting-started-guides/



## CI/CD



一个简单的build脚本

```sh
# the build stage in the Jenkinsfile- it switches directory, then runs two
# shell commands - the first sets up a script file so it can be executed
# and the second calls the script:
stage('Build') {
    steps {
      dir('ch11/exercises') {
      sh 'chmod +x ./ci/01-build.bat'
      sh './ci/01-build.bat'
    }
  }
}

# --pull所有镜像使用最新的
docker-compose -f docker-compose.yml -f docker-compose-build.yml build --pull
```




### 参考
- https://www.cyberciti.biz/faq/linux-xen-vmware-kvm-intel-vt-amd-v-support/
- [Best practices for writing Dockerfiles](https://docs.docker.com/develop/develop-images/dockerfile_best-practices/)
- monolithic architecture: 一种把所有系统都放在一起的架构方式，来源于建筑

- 防火墙: https://www.digitalocean.com/community/tutorials/how-to-set-up-a-firewall-using-firewalld-on-centos-7

- https://docs.docker.com/samples/library/mongo/ 