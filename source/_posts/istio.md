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

```yaml
apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: $vsvc
spec:
  hosts:
  - "*"
  gateways:
  - $gat4way
  http:
  - match:
    - headers:
        user-agent: ${regex}
      uri:
        prefix: $uri_path
    - uri:
        exact: $uri_path
    - uri:
        prefix: $uri_path
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
      delay:
        fixedDelay: ${time}
        percent: ${percent}  
      match:
      - headers:
```

## DESTINATION RULE

```yaml
apiVersion: "networking.istio.io/v1alpha3"
kind: "DestinationRule"
metadata:
  name: "default"
  namespace: "stock-trader"
spec:
  host: "*.stock-trader.svc.cluster.local"
  trafficPolicy:
    tls:
      mode: ISTIO_MUTUAL
    connectionPool:                     #----------------------------------[ 融短 ]
      tcp:
        maxConnections: ${n}
      http:
        http1MaxPendingRequest: ${n}
        maxRequestsPerConnection: ${n}
    outlierDecection:
      consecutiveErrors: ${n}
      interval: ${time}
      baseEjectionTime: ${time}
      maxEjectionPercent: ${percent}
```



## GATEWAY

```yaml
apiVersion: networking.istio.io/v1alpha3
kind: Gateway
metadata:
  name: $gateway
spec:
  selector:
  servers:
  - port:
      number:
      name:
      protocol: 
    hosts:
    - $hostname
    tls:
    	httpsRedirect:
	- port:
	...
  	
```



## SERVICE ENTRY

访问集群外部的服务

```yaml
apiVersion: networking.istio.io/v1alpha3
kind: ServiceEntry
metadata:
  name: ${service_entry}
spec:
  hosts:
  - ${host}
  ports:
  - number: ${port}
    name: https
    protocol: https
  resolution: DNS
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





## 出口流量

```bash
# 出口流量策略
$ kubectl get configmap istio -n istio-system -o yaml | grep -o 'mode: ALLOW_ANY'

# 设置出口流量为REGISTRY_ONLY
$ kubectl get configmap istio -n istio-system -o yaml | sed 's/mode: ALLOW_ANY/mode: REGISTRY_ONLY/g' | kubectl replace -n istio-system -f -configmap "istio" replaced
```



## 性能问题





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

