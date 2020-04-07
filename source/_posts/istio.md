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
$ curl -L https://git.io/getLatestIstio | sh
$ istioctl manifest apply --set profile=demo
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
    - uri:
        exact: $url_path
    - uri:
        prefix: $url_path
    route:
    - destination:
        host: $host
        port:
          number: $port
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



## 性能问题