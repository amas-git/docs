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



## Hello gRPC

- 建立proto文件

```protobuf
syntax = "proto3";
package goecho;

service Echo {
    rpc say(Msg) returns (Msg);
}

message Msg {
    int32 id = 1;
    string text = 2;
}
```



```sh
$ go get -u google.golang.org/grpc
$ tree .
.
└── goecho
    └── echo.proto
    
# 编译出go源文件
$ protoc -I goecho/ goecho/echo.proto --go_out=plugins=grpc:goecho   
$ tree .
.
└── goecho
    ├── echo.pb.go
    └── echo.proto
```



## 参考

	-  https://grpc.io/docs/guides/
	-  VSCode插件: vscode-proto3