PLAY k8s



```zsh
# 查看集群信息
$ kubectl cluster-info

# 查看节点信息
$ kubectl get node

# 查看所有资源对象
$ kubectl get all -A

# 查看default命名空间下的资源对象
$ kubectl get all
```

```zsh
$ kubectl create deployment busybox --image=busybox
deployment.apps/busybox created

# 开启柯南模式...
$ kubectl describe pods busybox
Events:
  Type     Reason     Age                     From               Message
  ----     ------     ----                    ----               -------
  Normal   Scheduled  <unknown>               default-scheduler  Successfully assigned default/busybox-7d657df9dc-pqrn5 to minikube
  Normal   Pulled     8m3s (x4 over 8m56s)    kubelet, minikube  Successfully pulled image "busybox"
  Normal   Created    8m3s (x4 over 8m56s)    kubelet, minikube  Created container busybox
  Normal   Started    8m3s (x4 over 8m56s)    kubelet, minikube  Started container busybox
  Normal   Pulling    7m14s (x5 over 9m)      kubelet, minikube  Pulling image "busybox"
  Warning  BackOff    3m49s (x24 over 8m50s)  kubelet, minikube  Back-off restarting failed container
  
# 1. Pod已经分配了节点
# 2. 镜像也准备好了
$ kubectl logs pod busybox-7d657df9dc-pqrn5 
# 3. 没有看到有任何日志
# 4. why?
# 因为busybox镜像没有Entrypoint, 所以容器启动完了就正常退出结束了

# 我们让容器飞个3000秒
$ kubectl run busybox-1 --image=docker.io/library/busybox:latest --replicas=2 -- sleep 3000

# 看起来正常了，我们attach到容器上试试，是可以的,不再出现CrashBackOff
$ kubectl exec -it busybox-7d657df9dc-pqrn5  -- /bin/sh 

$ docker image inspect busybox

# 去掉sleep大法，我们直接运行shell, 这样容器就不会退出了，k8s也不会尝试不断重启Pod
$ kubectl run busybox-1 --image=docker.io/library/busybox:latest --replicas=2 -- /bin/sh

# 明白原因之后我们基于busybox稍微改动一下，提供一个简单的echo-server
# https://blog.hasura.io/sharing-a-local-registry-for-minikube-37c7240d0615/
# 切换到minikube上的docker
$ eval $(minikube docker-env)

# 建立一个新的镜像
$ cat Dockerfile
FROM busybox
EXPOSE 8888
CMD ["nc","-ll","-p","8888","-e","/bin/cat"]

$ docker build -t simple-echo .
$ kubectl create deployment simple-echo --image=simple-echo
$ docker get pod
NAME                          READY   STATUS         RESTARTS   AGE
simple-echo-5994f57d7d-dbxpt          0/1     ErrImagePull   0          14s
# ErrImagePull? minikube怎么会找不到这个镜像呢？

$ docker image ls
REPOSITORY                                      TAG                 IMAGE ID            CREATED             SIZE
simple-echo                                     latest              1692d9a08912        About an hour ago   1.22MB

# 你的镜像在minikube的docker中，但是不在默认的registry中，尴尬了吧
$ docker pull simple-echo
Using default tag: latest
Error response from daemon: pull access denied for simple-echo, repository does not exist or may require 'docker login': denied: requested access to the resource is denied

# 为什么会这样？镜像明明已经在本地了，确还要去pull, k8s是不是傻？
# 1. 如果镜像只给名字，那么k8s会认为你实际需要的镜像是:latest的，所以会自动调用docker pull拉取镜像，那么你需要保证registry里有这个镜像，否则就会pull失败
# 2. 现在我们终于明白为何k8s非要拉取镜像了，只要我们重新build一个非:latest的镜像应该就可以了
# 3. 注意: 不要在生产环境中使用:latest镜像
$ docker built -t simple-echo:v1.0.0 .

# 删掉之前的deploymennt
$ kubectl delete deployment simple-echo

# 使用tagged镜像重新部署
$ kubectl create deployment simple-echo --image=simple-echo:v1.0.0
$ kubectl get pod
NAME                          READY   STATUS    RESTARTS   AGE
simple-echo-b9cd8d689-9v9n2   1/1     Running   0          1m

# pod已经起来了
# 需要用Service给集群之外提供访问
$ kubectl expose deployment simple-echo --type=NodePort --port=8888

$ kubectl get service
NAME                  TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
service/simple-echo   NodePort    10.96.76.249    <none>        8888:30202/TCP   21s
#  如果你对这个service的详细信息感兴趣，可以用下面的命令查看
$ kubectl get service simple-echo -o yaml

# 通过30202就可以访问服务啦
$ nc $(minikube ip) 30202
hello
hello
# === (完) ===
```

> Task: 接下来我们要运行多个Pod副本来提供echo服务
>
> ```zsh
> # 首先我们要修改一下simple-echo服务，让它返回主机信息
> ```
>
> 