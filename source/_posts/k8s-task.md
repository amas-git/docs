PLAY MINIKUBE



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
> xecho:
>
> ```sh
> #!/bin/sh
> busybox echo "echo server started"
> busybox env
> busybox echo -n "[REV] : "
> busybox cat
> ```
>
> Dockerfile:
>
> ```sh
> FROM busybox
> EXPOSE 8888
> COPY ./xecho /bin/xecho
> CMD ["nc","-ll","-p","8888","-e","/bin/xecho"]
> ```

```bash
$ kubectl create deployment xecho --image xecho:v1.0.0
$ kubectl get all
NAME                         READY   STATUS    RESTARTS   AGE
pod/xecho-67c74f4587-bt9b6   1/1     Running   0          5s

NAME                 TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)   AGE
service/kubernetes   ClusterIP   10.96.0.1    <none>        443/TCP   3m59s

NAME                    READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/xecho   1/1     1            1           5s

NAME                               DESIRED   CURRENT   READY   AGE
replicaset.apps/xecho-67c74f4587   1         1         1       5s
$ kubectl get services 
NAME         TYPE        CLUSTER-IP    EXTERNAL-IP   PORT(S)          AGE
kubernetes   ClusterIP   10.96.0.1     <none>        443/TCP          10m
xecho        NodePort    10.96.49.52   <none>        8888:31401/TCP   36s

# 扩容
$  kubectl scale --replicas=3  deployment xecho
deployment.apps/xecho scaled
$  kubectl get pods -o wide
NAME                     READY   STATUS    RESTARTS   AGE   IP           NODE       NOMINATED NODE   READINESS GATES
xecho-67c74f4587-bt9b6   1/1     Running   0          31m   172.17.0.2   minikube   <none>           <none>
xecho-67c74f4587-nrkzl   1/1     Running   0          20m   172.17.0.8   minikube   <none>           <none>
xecho-67c74f4587-tf5gg   1/1     Running   0          20m   172.17.0.3   minikube   <none>           <none>


$ kubectl describe services xecho
Name:                     xecho
Namespace:                default
Labels:                   app=xecho
Annotations:              <none>
Selector:                 app=xecho
Type:                     NodePort
IP:                       10.96.49.52
Port:                     <unset>  8888/TCP
TargetPort:               8888/TCP
NodePort:                 <unset>  31401/TCP
Endpoints:                172.17.0.2:8888,172.17.0.3:8888,172.17.0.8:8888
Session Affinity:         None
External Traffic Policy:  Cluster
Events:                   <none>

$ kubectl get endpoints xecho
NAME         ENDPOINTS                                         AGE
xecho        172.17.0.2:8888,172.17.0.3:8888,172.17.0.8:8888   13m
$ kubectl get endpoints xecho -o yaml
apiVersion: v1
kind: Endpoints
metadata:
  annotations:
    endpoints.kubernetes.io/last-change-trigger-time: "2019-12-21T03:50:13Z"
  creationTimestamp: "2019-12-21T03:44:53Z"
  labels:
    app: xecho
  name: xecho
  namespace: default
  resourceVersion: "185320"
  selfLink: /api/v1/namespaces/default/endpoints/xecho
  uid: 29feb0cd-1bc9-42d4-a4fe-53a5096e1404
subsets:
- addresses:
  - ip: 172.17.0.2
    nodeName: minikube
    targetRef:
      kind: Pod
      name: xecho-67c74f4587-bt9b6
      namespace: default
      resourceVersion: "183820"
      uid: 899253c5-1dda-49b4-bac2-de4f80f558c9
  - ip: 172.17.0.3
    nodeName: minikube
    targetRef:
      kind: Pod
      name: xecho-67c74f4587-tf5gg
      namespace: default
      resourceVersion: "185318"
      uid: beca461c-ea0c-4958-98e7-f77d420f3b5a
  - ip: 172.17.0.8
    nodeName: minikube
    targetRef:
      kind: Pod
      name: xecho-67c74f4587-nrkzl
      namespace: default
      resourceVersion: "185310"
      uid: cd2a481e-06e7-4b5d-8279-e94f76c7b536
  ports:
  - port: 8888
    protocol: TCP

$ minikube ip
192.168.99.101 

# 访问xecho服务10次，负载随机分配
$ repeat 10 echo hello | busybox nc -w 0 192.168.99.101 31401 | grep HOSTNAME
HOSTNAME=xecho-67c74f4587-tf5gg
HOSTNAME=xecho-67c74f4587-nrkzl
HOSTNAME=xecho-67c74f4587-bt9b6
HOSTNAME=xecho-67c74f4587-tf5gg
HOSTNAME=xecho-67c74f4587-tf5gg
HOSTNAME=xecho-67c74f4587-bt9b6
HOSTNAME=xecho-67c74f4587-nrkzl
HOSTNAME=xecho-67c74f4587-tf5gg
HOSTNAME=xecho-67c74f4587-nrkzl
HOSTNAME=xecho-67c74f4587-tf5gg
```

