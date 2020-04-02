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

```

```



## 参考

- https://opencensus.io/guides/grpc/go/#1
- https://linuxacademy.com/blog/kubernetes/running-prometheus-on-kubernetes/
- https://github.com/grpc-ecosystem/go-grpc-prometheus/issues/4
- https://github.com/philips/grpc-gateway-example/blob/master/cmd/serve.go
- https://sysdig.com/blog/kubernetes-monitoring-prometheus-operator-part3/
- https://devopscube.com/setup-prometheus-monitoring-on-kubernetes/