# istio

>
>
>## PROBLEM SPACE
>
>|                                             |      |
>| ------------------------------------------- | ---- |
>| APP Health Check / Performance Monitoring   | APP  |
>| APP Deployments                             | APP  |
>| APP Secret                                  | APP  |
>| Circuit Breaking                            | SVC  |
>| L7 Traffic Rounting : HTTP Redirect \| CORS |      |
>| Chaos Testing                               |      |
>| Canary Deployments                          |      |
>| Timeouts, Retres, Budgets,Deadlines         |      |
>| Per-request Routing                         |      |
>| Backpressure                                |      |
>| TLS                                         |      |
>| Identity, Access Control                    |      |
>| Quota management                            |      |
>| Protocal Transfer (REST, gRPC)              |      |
>| Policy                                      |      |
>| Service Performance Monitoring              | SVC  |
>| Cluster Management                          | LOW  |
>| Scheduling                                  | LOW  |
>| Orchestrator Update, Host Maintanance       | LOW  |
>| Service Discovery                           |      |
>| Networking & LB                             |      |
>| Stateful Services                           |      |
>| Multi-tenant , Multi-region                 |      |
>|                                             |      |
>

## 安装

```bash
$ curl -L https://istio.io/downloadIstio | sh -
# 安装指定版本
$ curl -L https://istio.io/downloadIstio | ISTIO_VERSION=1.4.3 sh -

# 进入到目录中
$ cd istion*
# 将bin/目录加入到PATH中

# 执行istioctrl开始安装
$ istioctl manifest apply --set profile=demo

# default空间开启istio注入
$ kubectl label namespace default istio-injection=enabled
$ kubectl apply -f samples/bookinfo/platform/kube/bookinfo.yaml

$ export INGRESS_PORT=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.spec.ports[?(@.name=="http2")].nodePort}')
$ export SECURE_INGRESS_PORT=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.spec.ports[?(@.name=="https")].nodePort}')

$ echo $INGRESS_PORT $SECURE_INGRESS_PORT
31386 30874

# 打开istio dashboard
$ istioctl dashboard kiali
```



```bash
# 根据service生成sidecar
$ istioctl kube-inject -f svc.yaml
```

## 注入

### 手动注入

```bash
# 手动注入
$ istioctl kube-inject -f $service.yaml | kubectl apply -f -

# 更细致的配置注入过程
$ istioctl kube-inject --injectConfigFile ${inject-config.yaml} \
                       --meshConfigFile ${inject-value.yaml}    \
                       --valuesFile ${mesh-config.yaml}         \
                       -f ${service.yaml}
```



```bash
# 注入POD模板
$ kubectl -n istio-system get configmap istio-sidecar-injector -o=jsonpath='{.data.config}'

# 注入POD的默认值
$ kubectl -n istio-system get configmap istio-sidecar-injector -o=jsonpath='{.data.values}'

# mesh的配置
$ kubectl -n istio-system get configmap istio -o=jsonpath='{.data.mesh}
```

### 自动注入

自动注入通过k8s的 [AddmissionControllers](https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/) 实现，

如果k8s关闭了此功能，可以参考: https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/#how-do-i-turn-on-an-admission-controller

```bash
# 需要给名字空间打一个标签，这样istio就可以通过addmission controllers来触发注入
$ kubectl label namespace ${ns:=default} istio-injection=enabled

$ kubectl get ns -L istion-injection
NAME                   STATUS   AGE     ISTIO-INJECTION
default                Active   7d      enabled
istio-system           Active   15h     disabled
kube-node-lease        Active   7d      
kube-public            Active   7d   
...
$ kubectl get ns default
apiVersion: v1
kind: Namespace
metadata:
  labels:
    istio-injection: enabled
  name: default
...

```

如何对某个POS关闭注入?

> sidecar.istio.io/inject:  "false"

```yaml
apiVersion: apps/v1
kind: Deployment
...
spec:
  template:
    metadata:
      annotations:
        sidecar.istio.io/inject: "false"  # <- check this one
    spec:
      containers:
      - name: ignored
        image: tutum/curl
        command: ["/bin/sleep","infinity"]
```



配置注入policy:

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: istio-sidecar-injector
data:
  config: |-
    policy: enabled
    neverInjectSelector:
      - matchExpressions:
        - {key: openshift.io/build.name, operator: Exists}
      - matchExpressions:
        - {key: openshift.io/deployer-pod-for.name, operator: Exists}
    template: |-
      initContainers:
```



注入控制优先级:

>  sidecar.istio.io/inject > neverInjectSelector > awaysInjectSelector > policy

卸载自动注入机制

```bash
$ kubectl delete mutatingwebhookconfiguration istio-sidecar-injector
$ kubectl -n istio-system delete service istio-sidecar-injector
$ kubectl -n istio-system delete deployment istio-sidecar-injector
$ kubectl -n istio-system delete serviceaccount istio-sidecar-injector-service-account
$ kubectl delete clusterrole istio-sidecar-injector-istio-system
$ kubectl delete clusterrolebinding istio-sidecar-injector-admin-role-binding-istio-system
```



## VirtualService

> 控制一个名字下的流量如何路由到目的地集合

VirtualService可以采用如下方式路由和控制流量:

- URI中的属性
- Headers
- 请求scheme
- 请求的目标端口
- 重试
- 请求超时

```yaml
apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: $vsvc
spec:
  hosts:
  - ${istio_internal_name}      # e.g.: foo.default.svc.cluster.local
                                #       *.foo.com
  gateways:                     #-------------------------------------[ 绑定Gateway ]
  - $gatway
  [http|tcp|tls]:               # 根据什么协议路由流量
  - match:
    - headers:
        user-agent: ${regex}
        cookie:
      uri:
        prefix: $uri_path
    - uri:
        exact: $uri_path
    - uri:
        regex: ${regex}
    - port: ${port}             # tcp
      destinationSubsets:
    - sniHosts:                 # tls
    route:
    - destination:
        host: $host
        subset: ${subset}
        port:
          number: $port
      weight: ${n}
    retries:
      attempts: ${n}
      perTryTimeout: ${time}
    timeout: ${time}            #------------------------------------[ 超时 ]
    fault:                      #------------------------------------[ FAULT INJECTION ]
      abort:
        httpStatus: ${http_code}
        percentage: ${percent}
      delay:
        fixedDelay: ${time}
        percentage: ${percent}  
      match:
      - headers:
# 没有被分流的请求会返回404      
```

> 注意： 实际开发中可以采用VirtualService分层的方式，这样可以避免多个team编辑同一个VirtualService定义



利用VirtualService我们可以实现

- 蓝绿发布
- 金丝雀发布
- 生产环境切分测试流量(在请求头中加入测试标识)



## DESTINATION RULE

配置如何与某个名字进行通讯,通过DestinationRule找到对应的ServiceEntry进而找到ServiceEntry中的Endpoint然后结束整个流量的路由。

ISTIO采用客户端LoadBalance取代反向代理，一方面使得系统的弹性更加强，第二方面客户端可以根据服务端的反馈动态调整自己的行为，如停止向不断出错的终端发送数据。



> OUTLIER DETECTION: 触发Endpoint的Lame-Ducking, 将lame-duck从active的LB中摘除

- secrets
- load-balancing 
  - 一致性哈希
- cricuit breaking
- (L4,L7)connection pool
- TLS

```yaml
apiVersion: "networking.istio.io/v1alpha3"
kind: "DestinationRule"
metadata:
  name: ${name}
  namespace: ${ns}
spec:
  host: ${istio_internal_name}
  subsets:
  - name:
    labels:
      version:
    trafficPolicy:
    ...
  - name:
    ...
  trafficPolicy:
    tls:
      mode: [ISTIO_MUTUAL|SIMPLE|DISABLED|MUTUAL]
            # ISTIO_MUTUAL: 使用istio管理mTLS证书
            # MUTUAL      : mTLS
            # SIMPLE      : TLS
            # DISABLED    : 不使用TLS
      clientCertificate: ${path}
      privateKey: ${path}
      caCertificates: ${path}
    connectionPool:                     #----------------------------------[ 短路 ]
      tcp:
        maxConnections: ${n}
      http:
        http1MaxPendingRequest: ${n}
        http2MaxRequests: ${n}
        maxRequestsPerConnection: ${n}
    outlierDecection:                  #-----------------------------------[ 破脚鸭检测 ]
      consecutiveErrors: ${n}
      interval: ${time}
      baseEjectionTime: ${time}
      maxEjectionPercent: ${percent}
    loadBalancer:
      simple: [LEAST_CONN|ROUND_ROBIN]
      consistentHash:
        useSourceIp: true
```

```bash
$ kubectl get destinationrules.networking.istio.io
```





## GATEWAY

> 向mesh外部暴露名字
>
> Gateway可以被VirtualService引用 (显式绑定VS， 也叫MeshGateway)
>
> VirtualService声明的hostname恰好被某个Gateway暴露了  (隐式绑定VS)
>
> 可以使用Gateway构建mTLS隧道，多个独立的mesh可以通过L3层网络链接到一起(实现跨地区，跨IDC的mesh)

```yaml
apiVersion: networking.istio.io/v1alpha3
kind: Gateway
metadata:
  name: $gateway
spec:
  selector:
  servers:
    - hosts: 
      - ${hostname}     
	  port:
	    number: ${port}
	    name:
	    protocol: [HTTP|HTTPS]
	  tls:              #----------------------[ 这个证书需要从CA申请长期证书 ]
	    mode: [SIMPLE|PASSTHROUGH]
	    serverCertificate:
	    privateKey:
	    httpsRedirect: [true|false] # true: 配合HTTP使用，将HTTP请求重定向到HTTPS，更加安全
	...
  	
```

```

```



## SERVICE ENTRY

> 定义新的名字，这个名字在mesh内部是被所有proxy知道的，可以访问到的，并不会添加到K8S的DNS中
>
> K8S的Service会

```yaml
apiVersion: networking.istio.io/v1alpha3
kind: ServiceEntry
metadata:
  name: ${service_entry} 
spec:
  hosts:   
  - ${host}                 # 这是一个在mesh内部可以访问的名字
  addresses: ${vips}        # Mesh内部可访问的虚拟IP
  location: [MESH_EXTERNAL] # 这个名字是MESH内部的还是外部的，如果外部可话可以用DNS解析或者静态IP
  ports:
  - number: ${port}
    name: https
    protocol: [https|http|tcp|udp|tls|redis|mongo|http2|grpc]
  resolution: [DNS|STATIC]  # 静态解析将使用下面定义的endpoints
  endpoints:
  - address: 2.2.2.2
  subjectAltNames:
  - "spiffe://..."
```

```bash
$ kubectl get serviceentries.networking.istio.io 
```

```bash
# 检测Mesh是否允许访问外部
$ kubectl get configmap istio -n istio-system -o yaml 
# global.outboundTrafficPolicy.mode:
#  - ALLOW_ANY
#  - REGISTRY_ONLY
$ kubectl edit configmap istio -n istio-system -o yaml 
```



## POLICY
```yaml
apiVersion: authentication.istio.io/v1alpha1
kind: Policy
metadata:
  name: $name
  namespace: $ns
spec:
  targets: # 如果不配置，则Policy作用于制定的${ns}

  - name: ${host:=${svc}-${ns}-svc-${cluster_domain}}
    port:
      name: ${port_name}
    peers:
  - mtls:
      mode: [STRICT | {} | PERMISSIVE]
    origins:
  - jwt:
    issuer:
    audiences:
    jwksUri:
    jwt_headers:
    principalBindings:  

```

## RBACCONFIG

```yaml
apiVersion: "rbac.istio.io/v1alpha1"
kind: RBACConfig
metadata:
  name: default
  namespace: istio-system
spec:
  mode: [ON|OFF|ON_WITH_INCLUSION|ON_WITH_EXCLUSION]
```



## CLUSTERRBACCONFIG

```yaml
apiVersion: "rbac.istio.io/v1alpha1"
kind: ClusterRBACConfig
metadata:
  name: default
  namespace: istio-system
spec:
  mode: ON_WITH_INCLUSION
  inclusion:
    services:
    - bar.bar.svc.cluster.local
    namespaces:
    - default
```





## Citadel

Citadel负责创建，颁发，回收证书， 管理整个mesh的keys和certs

Citadel使用 [SPIEFF](https://spiffe.io/) 格式构建StrongIdentities,用x.509证书编码

SVID(SPIEFF Verifiable Identity):

```
spiffe://${domain}/ns/${ns}/sa/${svc_account}
```


```bash
$ pod=$(kubectl get pod -l ${key}=${value} -o jsonpath{.items..metadata.name} -n $pod)
$ istioctl authn tls-check ${pod}.${svc} ${svc}.svc.cluster.local
```

```
[Register API] -> [Identity Register]
                        |
[Author]   ->   [ Authzer         ] -> [ Issuer ]      
   |                                       |
[                    CA SERVICE               ]

   |                                       |
   Requests                               Certs
```

1. Citadel以API的方式提供CA服务，Citadel需要从Pilot获取合法的名字
2. Citadel接受CSR, 然后进行一系列的认证，之后签名后以X.509 SVID的证书下发给NodeAgent
3. Citadel办法的证书通常只有1小时有效期，过了45分钟之后NodeAgent会为快要过期的证书发送CSR
4. NodeAgent会把证书发送到Envoy上 (SDS协议, Envoy 1.8.0+)
5. NodeAgent以无状态的方式工作，将所有的秘密保存到内存中，如果挂了，则由编排系统重新启动，启动之后从Citadel上同步所需要的数据
6. Pilot向Envoy发布配置，告诉目前有什么服务，如何连接
7. 所有Envoy的证书保存在/etc/certs目录下

## Pilot

> Pilot用来对数据平面编程
>
> Pilot根据Galley的要求，配合K8S的服务发现，将配置下发给Envoy, 完成对Mesh的配置

Pilot的配置包含三个方面 

- Mesh
  - istio各模块如何通讯
  - Proxy配置，Envoy初始化的配置
  - Mesh Networks
    - 如何使用Mixer
    - 如何配置proxy
    - 是否支持 k8s的Ingress
    - istio各组件的配置
- Networking
  - VirtualServices
  - ServiceEntries
  - DestinationRules
  - Gateways
- Service discovery

## Mixer

![](/src/amas/docs/source/_posts/istio.assets/DeepinScreenshot_select-area_20200409214750.png)

mixer是属性处理器，可以分为两类

- Policy Evaluation: checks，策略的二级缓存
  - ACLs
  - 配额管理
- Telemetry: reports, envoy将数据上报给mixer， mixer通过adpters把数据汇总给外部
  - 指标(metrics)
  - 日志(logs)
  - 分布式追踪(traces)

mixer是高可用的(可HA)， 无状态的服务

mixer这两个功能都包含在一个DockerImage中，编排的时候采用不同的命令和参数

mixer采用plug-ins架构，这些插件也叫做adaptor



```bash
$ kubectl -n istio-system get cm istio -o jsonpath="{@.data.mesh}" | grep disablePolicyChecks
$ kubectl -n istio-system get metrics requestduration -o yaml
```

> 属性:
>
> 属性是Mixer的关键概念，本质上是一些三元组(type,name,value)



> 问: 哪些策略来自Mixer?哪些策略来自Pilot
>
> 答: 影响流量的策略都由Pilot定义，需要增强认证的策略由Mixer定义



## 监控

> Envoy -> Mixer - adapter -> p8s -> grafana

需要配置三种资源:

- handler
- instance
- rules

## 出口流量

```bash
# 出口流量策略
$ kubectl get configmap istio -n istio-system -o yaml | grep -o 'mode: ALLOW_ANY'

# 设置出口流量为REGISTRY_ONLY
$ kubectl get configmap istio -n istio-system -o yaml | sed 's/mode: ALLOW_ANY/mode: REGISTRY_ONLY/g' | kubectl replace -n istio-system -f -configmap "istio" replaced
```

## istioctl

```sh
$ istioctl proxy-config bootstrap ${pod}
```



## 调试

- 不要使用UID 1337
- 不要使用HTTP1.0
- 每个POD至少关联到1个SERVICE上，如果关联到多个SERVICE, 则SERVICE中暴露的接口和协议保持一样
- DEPLOYMENT加上app和version标签
- POD必须开启NET_ADMIN, 不然istio没法完成注入, 如果使用了ISTIO CNI, 则NET_ADMIN不必要
- 

## 性能问题



## xDS协议

### LDS (Listener Discovory Service)

### RDS (Router Discovory Service)

### CDS (Cluster Discovory Service)

### EDS (Endpoint Discovory Service)

### ADS (Aggre Discovory Service)

### HDS (Health Discovory Service)

## 实验1

> 准备:
>
> 1. 安装好istio
> 2. 确保default名字空间开启istio自动注入

hello.go

```go
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func arg(i int, value string) string {
	if i > len(os.Args)-1 || os.Args[i] == "" {
		return value
	}
	return os.Args[i]
}

func main() {
	addr := arg(1, ":6666")
	tags := arg(2, "v1")

	fmt.Printf("START WITH %v [%v]\n", addr, tags)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("HELLO %s", tags)))
	})

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	log.Fatal(http.ListenAndServe(addr, nil))
}
```



Dockerfile:

```dockerfile
ARG GO_VERSION=1.14

FROM golang:${GO_VERSION}-alpine AS builder
WORKDIR /app
COPY .  /app
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "" cmd/hello.go



FROM alpine:latest
LABEL maintainer=zhoujb.cn@gmail.com
USER nobody
WORKDIR /app
EXPOSE 6666
COPY --from=builder --chown=nobody /app/hello .
```



hello_a_depolyment.yaml:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: hello
spec:
  selector:
    matchLabels:
      app: hello
  replicas: 1
  template:
    metadata:
      labels:
        app: hello
        version: "a"
    spec:
      containers:
      - name: hello
        image: hello:v1.0.0
        command: ["/app/hello"]
        args: [":6666", "a"]
        resources:
          limits:
            memory: "128Mi"
            cpu: "500m"
        ports:
        - containerPort: 6666
        readinessProbe:
            httpGet:
              path: "/healthz" # 需要配置readiness probe, istio会检测
              port: 6666
```





hello_a_svc.yaml

```yaml
apiVersion: v1
kind: Service
metadata:
  name: hello
spec:
  selector:
    app: hello
  ports:
  - port: 80
    targetPort: 6666
```

```sh
$ kubectl -f hello_a_depolyment.yaml
$ kubectl -f hello_a_svc.yaml
$ kubectl get pods hello-66f85974fd-4czjl -o yaml > pod.inject.yaml 
```

inject操作需要注入2个容器，是同一个image, 采用不同的启动方式

1. istio-proxy的istio-iptables作为初始化容器
2. istio-proxy sidecar
   - ISTIO_META_POD_PORTS
   - ISTIO_META_CLUSTER_ID
   - ISTIO_META_CONFIG_NAMESPACE
   - ISTIO_META_INTERCEPTION_MODE
   - ISTIO_META_WORKLOAD_NAME
   - ISTIO_META_OWNER
   - ISTIO_META_MESH_ID



## 实验1: 运行sleep pod

```bash
$ kubectl -f samples/sleep/sleep.yaml
# 确认istio是否完成了注入
$ export SLEEP_POD=$(kubectl get pod -l app=sleep -o jsonpath={.items..metadata.name})
$ kubectrl describe $SLEEP_POD 
...
  Normal  Pulled     3m53s      kubelet, minikube  Container image "docker.io/istio/proxyv2:1.5.1" already present on machine
  Normal  Created    3m53s      kubelet, minikube  Created container istio-proxy
  Normal  Started    3m53s      kubelet, minikube  Started container istio-proxy
  
$ kubectl exec -it $SLEEP_POD -c sleep curl hello-a.default.svc.cluster.local
HELLO a

# 查看
$ kubectl get svc istio-ingressgateway -n istio-system
NAME                   TYPE           CLUSTER-IP       EXTERNAL-IP     PORT(S)                                      AGE
istio-ingressgateway   LoadBalancer   172.21.109.129   130.211.10.121  80:31380/TCP,443:31390/TCP,31400:31400/TCP   17h
# EXTERNAL-IP: 如果设置，则当前环境有一个外部的LB, 如果是<none>或<pending>则说明不提供外部LB, 直接使用NodeIP访问
```



## 实验2: 限制mesh之外的http/https(Egress控制)

```bash
# 1. 首先看下outboundTrafficPolicy.mode的设置
$ kubectl get configmap istio -n istio-system -o yaml > tmp.yaml
​```
      outboundTrafficPolicy:
        mode: ALLOW_ANY      # <- 注意这个，允许envoy转发不在mesh里的域名请求
​```

# 进入到pod的容器中, 注意注入之后的pod里面有两个容器，我们指定一下进入到我们自己的容器里
$ kubectl exec -it sleep-f8cbf5b76-pzprd -c sleep -- sh  
/ # curl -I www.baidu.com
HTTP/1.1 200 OK
...

# 我们把outboundTrafficPolicy改为REGISTRY_ONLY
$  kubectl get configmap istio -n istio-system -o yaml | sed 's/mode: ALLOW_ANY/mode: REGISTRY_ONLY/g' | kubectl replace -n istio-system -f -
configmap "istio" replaced

# 配置下方到envoy需要一点时间，过一会可以登录到mesh里再试试
$ kubectl exec -it sleep-f8cbf5b76-pzprd -c sleep -- sh  
/ # curl -I www.baidu.com
HTTP/1.1 502 Bad Gateway
...
```





## 常用命令

```bash
$ istioctl proxy-status
# SYNCED:   envoy已经拿到最新的配置
# NOT SENT: envoy没有收到任何配置
# STALE:    istiod发送了配置，但没收到envoy的确认

# 如指定了具体的名字，则显示当前配置与最新配置的差异
$ istioctl proxy-status ${pod} 

# 查看指定POD的proxy配置
$ istioctl proxy-config cluster -n istio-system ${pod}

# 使用了哪个版本的envoy
$ kubectl exec -it $pod -c istio-proxy pilot-agent request GET server_info
```



## 参考

- SNI:https://en.wikipedia.org/wiki/Server_Name_Indication