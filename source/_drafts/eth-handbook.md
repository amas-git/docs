

## 以太坊客户端程序

### geth

### parity

## 使用geth

### 安装

### 启动

```bash
$ geth version
Geth
Version: 1.8.17-stable
Git Commit: 8bbe72075e4e16442c4e28d999edee12e294329e
Architecture: amd64
Protocol Versions: [63 62]
Network Id: 1
Go Version: go1.11.1
Operating System: linux
GOPATH=
GOROOT=/usr/lib/go
```

启动geth, 链接到以太坊主网

```
$ geth
```



- --networkid
  - 1: Frontier (默认值)
  - 2: Morden (已经不用了)
  - 3: Ropsten  (PoW)
  - 4: Rinkeby (PoA)
  - 其他: 私有网络id

- --datadir : 区块链数据保存的路径
- --syncmode: 同步数据的方式
  - fast
  - full
  - light
- --rpc
- --rpcaddr
- --rpcport
- --rpcapi

- --port geth默认监听的端口(默认是30303)



启动节点后, 可以使用geth attach命令链接到节点

```
$ geth attach [http://localhost:8485]
Welcome to the Geth JavaScript console!

instance: Geth/v1.8.17-stable-8bbe7207/linux-amd64/go1.11.1
 modules: admin:1.0 debug:1.0 eth:1.0 ethash:1.0 miner:1.0 net:1.0 personal:1.0 rpc:1.0 txpool:1.0 web3:1.0
 
 >
```



admin.peer: 

geth启动以后就会自动寻找伙伴节点然后开始下载区块链数据

```
> admin.peers
[{
    caps: ["eth/62", "eth/63", "par/1", "par/2", "par/3", "pip/1"],
    enode: "enode://061f6682095b31f824345450b25a23166fe672f015e98effa6c5438f48fc28dd377f3c009f01f2c62730801172717b55920913a1d72872983718265b7a6c537b@101.132.151.199:30303",
    id: "76fedc0bdcec3a617ed521249e356b78414829ada35f19d0e84f61d161be880b",
    name: "Parity-Ethereum/v2.2.11-stable-8e31051-20190220/x86_64-linux-gnu/rustc1.32.0",
    network: {
      inbound: false,
      localAddress: "10.60.113.123:52106",
      remoteAddress: "101.132.151.199:30303",
      static: false,
      trusted: false
    },
    protocols: {
      eth: {
        difficulty: 9.872041327696772e+21,
        head: "0x1219362f1768c1449b2d916f5b3ea07861471980e95a5b2abba5a1941d24c2b3",
        version: 63
      }
    }
}]

```

