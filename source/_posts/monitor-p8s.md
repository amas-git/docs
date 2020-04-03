# Prometheus

```bash
$ docker run -d --name p8s -p 9090:9090 prom/prometheus:latest
```

- http://localhost:9090/ 



```bash
$ docker inspect prom/prometheus | grep CMD
 "CMD [\"--config.file=/etc/prometheus/prometheus.yml\" \"--storage.tsdb.path=/prometheus\" \"--web.console.libraries=/usr/share/prometheus/console_libraries\" \"--web.console.templates=/usr/share/prometheus/consoles\"]"

```

启动选项:

- --config.file
- --storage.tsdb.path
- --web.console.libraries=/usr/share/prometheus/console_libraries
- --web.console.templates=/usr/share/prometheus/consoles





启动node-exporter

```bash
# 这个只是用着玩儿，实际在机器上不要用容器运行以取得更准确的指标
$ docker run --rm -p 9100:9100 prom/node-exporter  
```

- http://localhost:9100

----



- node_cpu_seconds_total {cpu="0",mode="idle"}
- avg without(cpu, mode)(rate(node_cpu_seconds_total{mode="idle"}[1m]))

## 在K8S上部署P8S

```bash
# 建立monitoring namespace
$ kubectl create namespace monitoring
$ kubectl create -f clusterRole.yaml
$ kubectl create -f config-map.yaml
$ kubectl create  -f prometheus-deployment.yaml 
# 飞一会
$ kubectl get deployments --namespace=monitoring
NAME                    READY   UP-TO-DATE   AVAILABLE   AGE
prometheus-deployment   1/1     1            1           78s

# 用port-forward试试, http://localhost:8080
$  kubectl port-forward --namespace monitoring prometheus-deployment-77cb49fb5d-glh9k 8080:9090 

# 换成service node port方式
$ kubectl create -f prometheus-service.yaml --namespace=monitoring
```



### 安装kube-state-matrics

```bash
$ git clone https://github.com/devopscube/kube-state-metrics-configs.git
$ kubectl apply -f kube-state-metrics-configs/
# 成功之后可以从p8s里面查看这个target已经UP
```

### 安装Alert  Manager

p8s的prometheus.yml中：

```yml
    alerting:
      alertmanagers:
      - scheme: http
        static_configs:
        - targets:
          - "alertmanager.monitoring.svc:9093" # 报警服务的地址
```

所有报警规则保存在`prometheus.rules`中

```yml
    groups:
    - name: devopscube demo alert
      rules:
      - alert: High Pod Memory
        expr: sum(container_memory_usage_bytes) > 1
        for: 1m
        labels:
          severity: slack
        annotations:
          summary: High Memory Usage
```



故障排除

```bash
$ kubectl describe --namespace=monitoring pods alertmanager-5c94c6bf89-t9s4v
  Warning  FailedMount  19s (x8 over 83s)  kubelet, minikube  MountVolume.SetUp failed for volume "templates-volume" : configmap "alertmanager-templates" not found

```





### 安装Grafana



### 服务发现

静态监控, 可在配置文件中添加

```yml
scrape_configs:
  - job_name: prometheus
	static_configs:
	 - targets:
      - localhost:9090
      - web:8888
```



文件形式的服务发现:

```yml
scrape_configs:
  - job_name: file
    file_sd_configs:
     - files:
     - '*.json'
```

```json
[
    {
    "targets": [ "host1:9100", "host2:9100" ],
    "labels": {
    "team": "infra",
    "job": "node"
    },
    ...
]
```

- 文件形式的服务发现，可以是通过CMS生成文件，也可以是用cron job定期获取

### 监控节点

> 监控主机节点时，不要用容器运行expoter,以获得更准确的监控数据

### APP监控

1. APP需要以http/https方式暴露/metrics
2. k8s需要在Service中配置anotation, 以便p8s可以检索到这些需要被监控的服务，以及如何监控

## 自定义Exporter

## PromQL

> Gauges: 意思是状态的快照

```sh
# 不带device,fstype,mountpoint标签的node_filesystem_size_bytes求和
sum without(device, fstype, mountpoint)(node_filesystem_size_bytes)

# 求平均
avg (without(a,b,c))(x)

# 求最大
max (without(a,b,c))(x)

# 求最小
min (without(a,b,c))(x)

# 过去5分钟的rate, sum(x[5m])/5m*60
rate(x[5m])

# 5分钟之前的总请求数量
sum(http_requests_total{method="GET"} offset 5m) 

# Selector
x{a="" b!="" d=~"regex" e=!~"regex"}

# TimeUnit: ms s m h d w y

# by: 带有a,b,c标签的
by(a,b,c)(x)

# 统计gauges的长度
count(x)

#
stddev(x)

#
stdvar(x)

# topk, TOP n
topk(n,x)

# bottomk, BOTTOM n
bottomk(n, x)

# quantile: 90%分位点
quantile(0.9, x)

# count_values, 按照value分组，计算出value出现的次数
software_version{instance="a",job="j"} 7
software_version{instance="b",job="j"} 4
software_version{instance="c",job="j"} 8
software_version{instance="d",job="j"} 4
software_version{instance="e",job="j"} 7
software_version{instance="f",job="j"} 4
count_values withoutt(instance)("version", software_version)
{job="j",version="7"} 2
{job="j",version="8"} 1
{job="j",version="4"} 3

# 二元操作
x / y
x + y
x - y
x * y
x % y
^e

==
!=
>
<
>=
<=
```

直方图

> 用来分析事件大小的分布
>
> ```
> histogram_quantile(
> 0.90,
> rate(prometheus_tsdb_compaction_duration_seconds_bucket[1d]))
> 
> 结果：
> {instance="172.17.0.10:9090",job="kubernetes-service-endpoints",kubernetes_name="prometheus-service",kubernetes_namespace="monitoring"} 0.9
> 
> 意思是90%的延迟在0.9s
> ```



## 报警



## 参考

- https://opencensus.io/guides/grpc/go/#1
- https://linuxacademy.com/blog/kubernetes/running-prometheus-on-kubernetes/
- https://github.com/grpc-ecosystem/go-grpc-prometheus/issues/4
- https://github.com/philips/grpc-gateway-example/blob/master/cmd/serve.go
- https://sysdig.com/blog/kubernetes-monitoring-prometheus-operator-part3/
- https://devopscube.com/setup-prometheus-monitoring-on-kubernetes/