# gRPC



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



### 工具

```sh
$ go get github.com/fullstorydev/grpcurl
$ go install github.com/fullstorydev/grpcurl/cmd/grpcurl
```



## 参考

	-  https://grpc.io/docs/guides/
	-  VSCode插件: vscode-proto3