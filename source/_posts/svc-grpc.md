## gRPC

## 技术背景

- HTTP2
- Protocol Buffer

|                       |                |
| --------------------- | -------------- |
| 协议                  | 二进制，效率高 |
| HTTP/2                |                |
| 强类型                |                |
| 多语言支持（Polyglot) |                |
| 双工流                |                |
| 认证(Authentication)                |                |
| 加密                |                |
| 弹性                |    死期，超时，压缩，负载均衡，服务发现            |
| 双工流                |                |

## gRPC的缺点
 - 某些场合下不如REST/HTTP友好
 - 生态圈仍然有限
 - 调整协议需要更新客户端和服务端的代码

## gRPC / Thrift / GraphQL

|      | gRPC   | Thrift | GraphQL |
| ---- | ------ | ------ | ------- |
| 传输 | HTTP/2 |        |         |
| 流   |        |        |         |
| 性能 |        |        |         |
|      |        |        |         |

```
export GOPROXY=https://mirrors.aliyun.com/goproxy/
```



## Hello gRPC

- 建立echo.proto文件

```protobuf
syntax = "proto3";
package model;

service Echo {
    rpc say(Msg) returns (Msg);
}

message Msg {
    int32  id   = 1;
    string text = 2;
    enum Type {
        HIGH = 0;
        NORM = 1;
        LOW  = 2;
    }
    Type type = 3;
}
```



```sh
$ go get -u google.golang.org/grpc
$ tree
.
├── cmd
├── go.mod
├── go.sum
├── main.go
├── model
│   ├── msg.pb.go
│   └── msg.proto
├── svc
│   └── echosvc.go
├── test
└── utils
    └── utils.go
```



```bash
$ grpcurl -import-path model -proto model/msg.proto localhost:8888 list
model.Echo

$ grpcurl -import-path model -proto model/msg.proto localhost:8888 describe model.Echo  
model.Echo is a service:
service Echo {
  rpc say ( .model.Msg ) returns ( .model.Msg );
}

$ grpcurl -import-path model -proto model/msg.proto -plaintext -d '{"id":1, "text":"Hello gRPC"}' localhost:8888 model.Echo/say
{
  "id": 2
}
```



## ProtocolBuffer

### Message ID

- 范围: [1, 2^29-1] 
- [1,15] :  占用1字节，尽量分配给出现频率最高的字段
- [16,2047] : 占用2字节
- 19000-19999： 保留ID, 不要使用



### 常用类型（google.protobuf）

有很多类型google已经定义好了，我们不必重复造轮子

- https://developers.google.com/protocol-buffers/docs/reference/google.protobuf



### 函数映射

| PROTOBUF                            | GOLANG                                               |
| ----------------------------------- | ---------------------------------------------------- |
| rpc method(T1) returns (T2)         | func(Context, \*T1) (\*T2, error)                    |
| rpc method(T1) returns (stream T2); | func(Context, \*T1, \*${package}_methodServer) error |
|                                     |                                                      |
|                                     |                                                      |



### 交互方式

- Simple RPC
- 客户端Stream
- 服务端Stream
- 双向Stream

## 初始化

#### 服务端

```go
addr := ":8080" 
l, err := net.Listen("tcp", addr)
if err != nil {
  log.Fatalf("FAILED TO CREATE SERVER @%v : %v", addr, err)
}
svc := grpc.NewServer()
pb.Register/*${package}*/Server(svc, s)
if err := svc.Serve(l); err != nil {
  log.Fatalf("FAILED TO START: %v", err)
}
```



#### 客户端

```go
conn, err := grpc.Dial(addr, grpc.WithInsecure()/* ...options */)
defer conn.Close()

if err != nil {
  log.Fatalf("DID NOT CONNECT: %v", err)
}
c := pb.New/*${package}*/Client(conn)

ctx, cancel := context.WithCancel(context.Background())
defer cancel()

// DO RPC CALL
...
```



## Interceptors

在执行gRPC整个过程中可以插入一些代码，这就是Interceptors，有两种类型的Interceptors

- Unary Interceptor
- Stream Interceptor

### 服务端INTERCEPTOR

#### UNARY

```go
func(ctx context.Context, 
     req interface{}, 
     info *UnaryServerInfo,
	   handler UnaryHandler) 
 
(resp interface{}, err error)
```



#### STREAM

```go
func(srv interface{}, 
     ss      ServerStream, 
     info    *StreamServerInfo,
     handler StreamHandler) 
error
```



#### 安装方式

```go
s := grpc.NewServer(grpc.UnaryInterceptor(${interceptor}))
s := grpc.NewServer(grpc.StreamInterceptor(${interceptor}))
```



### 客户端INTERCEPTOR

#### UNARY

```go
func(ctx context.Context, 
     method  string, 
     req     interface{}, 
     reply   interface{},
     cc      *ClientConn,
     invoker UnaryInvoker, 
     opts ...CallOption) 
error
```

#### STREAM

```go
func(ctx context.Context, 
     desc     *StreamDesc, 
     cc       *ClientConn,
     method   string, 
     streamer Streamer,
     opts     ...CallOption) 
(ClientStream, error)
```

#### 安装方式

```go
conn, err := grpc.Dial(address, grpc.WithStreamInterceptor(${interceptor}))
conn, err := grpc.Dial(address, grpc.WithUnaryInterceptor(${interceptor}))
```



## 截止时间和超时

超时(Timeout)是客户端所能容忍的最长等待时间.

> ALWAYS SET DEADLINE ! 具体设置多长时间没有固定标准，根据实际情况出发

截止时间默认是一个非常大的值，具体多大取决于各个语言对gRPC的实现。如果在截止时间之内没有完成gRPC调用，则会返回DEADLINE_EXCEEDED错误。

设置方法:

```go
// GO语言利用Context包处理
ctx, cancel := context.WithDeadline(context.Background(), ${deadline}) 
```



## 取消

## 错误处理

| CODE                | VALUE | DESC     |
| ------------------- | ----- | -------- |
| OK                  | 0     | 成功     |
| CANCELLED           | 1     | 取消     |
| UNKNOWN             | 2     | 未知错误 |
| INVALID_ARGUMENT    | 3     | 参数错误 |
| DEADLINE_EXCEEDED   | 4     |          |
| NOT_FOUND           | 5     |          |
| ALREADY_EXISTS      | 6     |          |
| PERMISSION_DENIED   | 7     |          |
| UNAUTHENTICATED     | 16    |          |
| RESOURCE_EXHAUSTED  | 8     |          |
| FAILED_PRECONDITION | 9     |          |
| ABORTED             | 10    |          |
| OUT_OF_RANGE        | 11    |          |
| UNIMPLEMENTED       | 12    |          |
| INTERNAL            | 13    |          |
| UNAVAILABLE         | 14    |          |
| DATA_LOSS           | 15    |          |
|                     |       |          |

服务端：

```go
// 简单的错误
...
	if msg.Id < 0 {
		return nil, status.Error(codes.InvalidArgument, "Id must > 0")
	}
...

// 复杂的错误，可以携带额外的数据
errrStatus := status.New(codes.InvalidArgument, "Id must > 0")
ds, err := errStatus.WithDetails {
    &errdetails.BadRequest_FieldViolation {
        Field: "Id",
        Description: fmt.Sprintf("Id = %v, MUST > 0", msg.Id)
    }
} 
```



```sh
$  grpcurl -import-path model -proto model/msg.proto  -d '{"id":-1, "text":"Hello gRPC"}'  -cacert cert/svc.crt  localhost:8888 model.Echo/say
ERROR:
  Code: InvalidArgument
  Message: Id must > 0
```



客户端：

```go
st, ok := status.FromError(err)
if !ok {
	...
	// st.Code(), st.Message()
}

// 或者
st := status.Convert(err)
for _, detail := range st.Details() {
    switch t := detail.(type) {
    case *errdetails.BadRequest:
        fmt.Println("Oops! Your request was rejected by the server.")
        for _, violation := range t.GetFieldViolations() {
            fmt.Printf("The %q field was wrong:\n", violation.GetField())
            fmt.Printf("\t%s\n", violation.GetDescription())
        }
    }
}
```



## 多路复用
Multiplexing
多个gRPC调用可以复用一条HTTP2连接

## 携带元数据

```go
metadata.New(map[string]string{"key1": "val1", "key2": "val2"}).
```



## 名字解析

## 数据压缩

```go
import "google.golang.org/grpc/encoding/gzip"
```



## 负载均衡

- 代理负载均衡
- 客户端负载均衡

## 安全

### Oneway TLS

> 客户端验证服务器

- server.key
- server.crt | server.pem

```sh
$ openssl req -nodes -x509 -newkey rsa:4096 -keyout svc.key -out svc.crt -days 365
Generating a RSA private key
................................................++++
..............................................................................................................................++++
writing new private key to 'svc.key'
-----
You are about to be asked to enter information that will be incorporated
into your certificate request.
What you are about to enter is what is called a Distinguished Name or a DN.
There are quite a few fields but you can leave some blank
For some fields there will be a default value,
If you enter '.', the field will be left blank.
-----
Country Name (2 letter code) [AU]:.
State or Province Name (full name) [Some-State]:.
Locality Name (eg, city) []:.
Organization Name (eg, company) [Internet Widgits Pty Ltd]:.
Organizational Unit Name (eg, section) []:.
Common Name (e.g. server FQDN or YOUR name) []:localhost # 注意CN要设置正确，如果你想在本地测试可以使用localhost
Email Address []:. 
```



服务端

```go
import (
  "crypto/tls"
  "google.golang.org/grpc/credentials"
  ...
}
  
cert, err := tls.LoadX509KeyPair(${server.crt},${server.key}) 
opts := []grpc.ServerOption{
  grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
}
  
s := grpc.NewServer(opts...) 
pb.Register${package}Server(s, &server{})
  
l, err := net.Listen("tcp", port)
s.Serve(l)
  
```

客户端

```go
creds, err := credentials.NewClientTLSFromFile(${server.crt}, ${hostname}) 
if err != nil {
  log.Fatalf("failed to load credentials: %v", err)
}
opts := []grpc.DialOption{
  grpc.WithTransportCredentials(creds), 
}

conn, err := grpc.Dial(address, opts...) 
if err != nil {
  log.Fatalf("did not connect: %v", err)
}
defer conn.Close() 
c := pb.New${package}Client(conn) 
// DO RPC CALL
```

```bash
$ grpcurl -import-path model -proto model/msg.proto  -d '{"id":100, "text":"Hello gRPC"}' localhost:8888 model.Echo/say   
Failed to dial target host "localhost:8888": x509: certificate is valid for amas.org, not localhost

# 使用-insecure参数，不检查CA信任链
$ grpcurl -import-path model -proto model/msg.proto -insecure -d '{"id":100, "text":"Hello gRPC"}' localhost:8888 model.Echo/say   
{
  "id": 101,
  "text": "Hello gRPC"
}
# 将服务端的证书作为根证书
# -cacert cert/svc.crt 
$  grpcurl -import-path model -proto model/msg.proto  -d '{"id":100, "text":"Hello gRPC"}' -cacert cert/svc.crt localhost:8888 model.Echo/say 
{
  "id": 101,
  "text": "Hello gRPC"
}
```



### mTLS

> 客户端服务器双向验证

- server.key
- server.crt
- client.key
- client.crt
- ca.crt

### OAuth2

### JWT



## 测试

### 压测ghz

```bash
$  ghz --proto model/msg.proto --cacert cert/svc.crt --call model.Echo/say -n 1000 -d '{"id":100, "text":"Hello gRPC"}' localhost:8888  

Summary:
  Count:        1000
  Total:        38.03 ms
  Slowest:      11.30 ms
  Fastest:      0.14 ms
  Average:      1.77 ms
  Requests/sec: 26293.81

Response time histogram:
  0.144 [1]     |
  1.260 [557]   |∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  2.376 [320]   |∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
  3.492 [61]    |∎∎∎∎
  4.608 [11]    |∎
  5.724 [0]     |
  6.840 [0]     |
  7.957 [0]     |
  9.073 [0]     |
  10.189 [0]    |
  11.305 [50]   |∎∎∎∎

Latency distribution:
  10 % in 0.67 ms 
  25 % in 0.91 ms 
  50 % in 1.19 ms 
  75 % in 1.65 ms 
  90 % in 2.60 ms 
  95 % in 10.51 ms 
  99 % in 10.97 ms 

Status code distribution:
  [OK]   1000 responses
```



## 部署

### DOCKER

### K8S

```
deployments
├── docker
│   └── Dockerfile
├── echosvc-deployment.k8s.yaml
└── echosvc-service.k8s.yaml
```

echosvc-deployment.k8s.yaml:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: echosvc
spec:
  replicas: 1
  selector:
    matchLabels:
      app: echosvc
  template:
    metadata:
      labels:
        app: echosvc
    spec:
      containers:
      - name: echosvc
        image: echosvc:v1.0.0
        resources:
          limits:
            memory: "128Mi"
            cpu: "500m"
        ports:
        - containerPort: 8888
```

echosvc-service.k8s.yaml:

```yaml
apiVersion: v1
kind: Service
metadata:
  name: echosvc
spec:
  selector:
    app: echosvc
  ports:
  - nodePort: 30100  # NodePort,可通过该节点的$IP:30100访问
    port: 8888       # Service的端口
    targetPort: 8888 # 通过selector选中对象的端口
  type: NodePort
```

```sh
# 启动minikube
$ minikube start
# 使用minikube上的docker
$ eval $(minikube docker-env) 
# 构建镜像, 注意此处buildTag要设置
$ docker build -t echosvc:v1.0.0 -f deployments/docker/Dockerfile . 

$ kubectl apply -f echosvc-deployment.k8s.yaml:
$ kubectl apply -f deployments/echosvc-service.k8s.yaml
$ kubectl get svc  
NAME         TYPE        CLUSTER-IP    EXTERNAL-IP   PORT(S)          AGE
echosvc      NodePort    10.96.100.3   <none>        8888:30100/TCP   28m
kubernetes   ClusterIP   10.96.0.1     <none>        443/TCP          83m
$ kubectl get pod
NAME                       READY   STATUS    RESTARTS   AGE
echosvc-6585d48566-klpln   1/1     Running   0          42m
$ ./echoc
# 查看服务端日志
$ kubectl logs echosvc-6585d48566-klpln 
```





## 监控

OpenCensus 

Prometheus

## 日志

## 追踪

## gRPC生态

### gRPC网关

### HTTP/JSON Transcoding for gRPC

### gRPC Server Reflection Protocol

### gRPC Middleware

|                  |      |      |
| ---------------- | ---- | ---- |
|                  |      |      |
| grpc_auth        |      |      |
| grpc_ctxtags     |      |      |
| grpc_zap         |      |      |
| grpc_logrus      |      |      |
| grpc_prometheus  |      |      |
| grpc_opentracing |      |      |
| grpc_retry       |      |      |
| grpc_validator   |      |      |
| grpc_recovery    |      |      |
| grpc_ratelimit   |      |      |
|                  |      |      |

### 健康检查

### 健康探测



## 最佳实践

- https://gist.github.com/tcnksm/eb78363fda067fdccd06ee8e7455b38b

### API设计

### 错误处理

### 截至时间

- 截至时间是服务端和客户端都清楚应该合适终止操作
- 总是使用截至时间
- 客户端负责设置截至时间
- 服务端负责检查截至时间，并做恰当的处理

### 限速

- 服务端限速: grpc.IntapHandle(rateLimitter)
- 客户端可以实现调用限速

### 重试

- 官方计划在[gRFC A6]( https://github.com/grpc/proposal/blob/master/A6-client-retries.md)中支持
  - 通过服务端的配置实现
  - 计划支持
    - 顺序重试
    - 并发对冲请求(hedged requests)
- 目前使用客户端包装或Interceptor处理重试
- 重试可以用中间件的方式实现，避免重复开发

```go
 d, _ := ctx.Deadline()
 ctx1, cancel := context.WithDeadline(ctx, d.Add(-150*time.Millisecond))
```



### 内存管理

1. gRPC go版本不限制server使用的goroutines
   1. 限制网络监听器的数量: netutil.LimitListener
   2. 使用TapHandler处理淤积RPC
2. 服务端设置载荷大小
3. 使用StreamAPI

### 日志

### 监控

- 尽可能暴露一切必要的度量指标

## 工具

```sh
$ go get github.com/fullstorydev/grpcurl
$ go install github.com/fullstorydev/grpcurl/cmd/grpcurl
```



```

```



## 参考

-  https://grpc.io/docs/guides/
-  https://github.com/square/certstrap
-  VSCode插件: vscode-proto3
-  https://github.com/grpc-ecosystem/awesome-grpc
-  测试: https://github.com/bojand/ghz
-  gRFC: https://github.com/grpc/proposal
