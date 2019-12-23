k8s-network

通常Service可以指定一个ClusterIP(也叫做VIP), 一个ClusterIP对应一个或多个Endpoint对象， 在Endpoint对象中保存了Pod的实际IP, 这样VIP的流量会被节点上的kube-proxy根据Endpoint对象按照负载均衡算法分配到PodIP上

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

