---
title: 玩转以太坊客户端
date: 2018-08-03 18:26:22
tags:
---
<!-- toc -->
# 玩转以太坊客户端
## geth
### geth console

```sh 
$ geth console
INFO [08-03|19:46:12] Maximum peer count                       ETH=25 LES=0 total=25
INFO [08-03|19:46:12] Starting peer-to-peer node               instance=Geth/v1.8.2-stable-b8b9f7f4/linux-amd64/go1.10
INFO [08-03|19:46:12] Allocated cache and file handles         database=/home/amas/.ethereum/geth/chaindata cache=768 handles=512
INFO [08-03|19:46:12] Writing default main-net genesis block 
INFO [08-03|19:46:12] Persisted trie from memory database      nodes=12356 size=2.34mB time=34.918604ms gcnodes=0 gcsize=0.00B gctime=0s livenodes=1 livesize=0.00B
INFO [08-03|19:46:12] Initialised chain configuration          config="{ChainID: 1 Homestead: 1150000 DAO: 1920000 DAOSupport: true EIP150: 2463000 EIP155: 2675000 EIP158: 2675000 Byzantium: 4370000 Constantinople: <nil> Engine: ethash}"
INFO [08-03|19:46:12] Disk storage enabled for ethash caches   dir=/home/amas/.ethereum/geth/ethash count=3
INFO [08-03|19:46:12] Disk storage enabled for ethash DAGs     dir=/home/amas/.ethash               count=2
INFO [08-03|19:46:12] Initialising Ethereum protocol           versions="[63 62]" network=1
INFO [08-03|19:46:12] Loaded most recent local header          number=0 hash=d4e567…cb8fa3 td=17179869184
INFO [08-03|19:46:12] Loaded most recent local full block      number=0 hash=d4e567…cb8fa3 td=17179869184
INFO [08-03|19:46:12] Loaded most recent local fast block      number=0 hash=d4e567…cb8fa3 td=17179869184
INFO [08-03|19:46:12] Regenerated local transaction journal    transactions=0 accounts=0
INFO [08-03|19:46:12] Starting P2P networking 
INFO [08-03|19:46:14] UDP listener up                          self=enode://61c5ad55c2d77cffb8da106be4eb1050cef8b913f189ea885c116c6d17fde6f7c475de5a29cd4485585fb25e52c387c0ad5e2131ac3b514b4d37a4619c3bd8dd@[::]:30303
INFO [08-03|19:46:14] RLPx listener up                         self=enode://61c5ad55c2d77cffb8da106be4eb1050cef8b913f189ea885c116c6d17fde6f7c475de5a29cd4485585fb25e52c387c0ad5e2131ac3b514b4d37a4619c3bd8dd@[::]:30303
INFO [08-03|19:46:14] IPC endpoint opened                      url=/home/amas/.ethereum/geth.ipc
Welcome to the Geth JavaScript console!

instance: Geth/v1.8.2-stable-b8b9f7f4/linux-amd64/go1.10
 modules: admin:1.0 debug:1.0 eth:1.0 miner:1.0 net:1.0 personal:1.0 rpc:1.0 txpool:1.0 web3:1.0

>

$ tree ~/.ethereum
/home/amas/.ethereum
├── geth
│   ├── chaindata
│   │   ├── 000001.log
│   │   ├── CURRENT
│   │   ├── LOCK
│   │   ├── LOG
│   │   └── MANIFEST-000000
│   ├── LOCK
│   ├── nodekey
│   ├── nodes
│   │   ├── 000001.log
│   │   ├── CURRENT
│   │   ├── LOCK
│   │   ├── LOG
│   │   └── MANIFEST-000000
│   └── transactions.rlp
├── geth.ipc
└── keystore

```
### geth attach
#### admin.nodeInfo
```json
{
  enode: "enode://61c5ad55c2d77cffb8da106be4eb1050cef8b913f189ea885c116c6d17fde6f7c475de5a29cd4485585fb25e52c387c0ad5e2131ac3b514b4d37a4619c3bd8dd@[::]:30303",
  id: "61c5ad55c2d77cffb8da106be4eb1050cef8b913f189ea885c116c6d17fde6f7c475de5a29cd4485585fb25e52c387c0ad5e2131ac3b514b4d37a4619c3bd8dd",
  ip: "::",
  listenAddr: "[::]:30303",
  name: "Geth/v1.8.2-stable-b8b9f7f4/linux-amd64/go1.10",
  ports: {
    discovery: 30303,
    listener: 30303
  },
  protocols: {
    eth: {
      config: {
        byzantiumBlock: 4370000,
        chainId: 1,
        daoForkBlock: 1920000,
        daoForkSupport: true,
        eip150Block: 2463000,
        eip150Hash: "0x2086799aeebeae135c246c65021c82b4e15a2c451340993aacfd2751886514f0",
        eip155Block: 2675000,
        eip158Block: 2675000,
        ethash: {},
        homesteadBlock: 1150000
      },
      difficulty: 17179869184,
      genesis: "0xd4e56740f876aef8c010b86a40d5f56745a118d0906a34e69aec8c0db1cb8fa3",
      head: "0xd4e56740f876aef8c010b86a40d5f56745a118d0906a34e69aec8c0db1cb8fa3",
      network: 1
    }
  }
}
```
#### admin.peers
```json 
[{
    caps: ["eth/62", "eth/63"],
    id: "05b12b8ef42e1e066468ccd5cd0a32da442cc83289c9f07cfc10c7e8a554af6adc686ac7218d84d1b250325d4add9ccb4a690cde0d42ad62f97adc2185d8624b",
    name: "Geth/v0.1.1-akroma-5fa3ee8c/linux-amd64/go1.10.1",
    network: {
      inbound: false,
      localAddress: "10.60.113.123:44490",
      remoteAddress: "45.77.90.228:30303",
      static: false,
      trusted: false
    },
    protocols: {
      eth: "handshake"
    }
}, {
    caps: ["eth/62", "eth/63"],
    id: "29215d0950d7f1807e831b55c8497821edcff228e49bb5321251caa3a7a43e1aec915c7467155827a14f35bd13332dd50b6ed0d7655c44cf8b7ce48e27248596",
    name: "Geth/v1.8.10-stable-eae63c51/linux-amd64/go1.10",
    network: {
      inbound: false,
      localAddress: "10.60.113.123:58974",
      remoteAddress: "47.91.209.156:30303",
      static: false,
      trusted: false
    },
    protocols: {
      eth: {
        difficulty: 5.765374924733686e+21,
        head: "0x96f8f9bbffea3de2f5156b2e0974f48d40c5a4200efc8a0744ff98d326b26c8f",
        version: 63
      }
    }
}, {
    caps: ["eth/62", "eth/63"],
    id: "42d4febd3ed46f39bc33da80822f0b681b6dba4ff23f20c48b5682556e0a1e5bffc50af88fc68b92ae7e8a1a4d59d7e731ea2aa461a09aa3b40d48263eb42708",
    name: "Geth/v1.8.6-stable-12683fec/linux-amd64/go1.10",
    network: {
      inbound: false,
      localAddress: "10.60.113.123:46734",
      remoteAddress: "128.199.45.106:30303",
      static: false,
      trusted: false
    },
    protocols: {
      eth: {
        difficulty: 5.765274976706829e+21,
        head: "0x22127969ec067298ac10b42d536e3958dd6d43af3789eae2d0881bffeb2ce8e7",
        version: 63
      }
    }
}, {
    caps: ["eth/62", "eth/63", "par/1", "par/2", "par/3", "pip/1"],
    id: "97b6ef4f927a0fc7b4d08251dc594cd1a5518a02ab0682f071e28a39ed80813c7b263788f29241a77683ba8f155d148982b556aef9511882789fd9e825647a53",
    name: "Parity/v1.11.8-stable-c754a02-20180725/x86_64-linux-gnu/rustc1.27.2",
    network: {
      inbound: false,
      localAddress: "10.60.113.123:39624",
      remoteAddress: "185.64.116.14:30303",
      static: false,
      trusted: false
    },
    protocols: {
      eth: {
        difficulty: 5.765285321878454e+21,
        head: "0xf23856426327ec303444800ed785d0bd6c5402d435f5bcd3ae44196e4c5d037d",
        version: 63
      }
    }
}]
```

```
> admin.datadir
"/home/amas/.ethereum"
> net.listening
true
> net.version
"1"
> net.peerCount
4
```


```bash
# 执行js命令
$ geth --verbosity 0 attach --exec 'admin.peers'


# attach
$ geth attach ipc:/some/path
$ geth attach rpc:http://host:8545

# 执行js代码
$ geth --verbosity 0  js <(echo 'console.log("hello eth")')

# 启动一个孤立节点
$ geth --maxpeers 0 --nodiscover --networdid 3301 

# 运行RPC服务
$ geth --rpc=true --rpcport 8000 --rpccorsdomain '"*"'

# 挖矿
$ geth --mine --minerthreads 4 --datadir /usr/local/share/ethereum/30303 --port 30303
$ geth --mine --minerthreads 4 --datadir /usr/local/share/ethereum/30304 --port 30304

# 设置bootnode
$ geth --bootnodes "enode://pubkey1@ip1:port1 enode://pubkey2@ip2:port2 enode://pubkey3@ip3:port3"


# 检测节点的连接, net.listening=true且peerCount大于0
# 通常有两种情况会造成链接失败
#  1. 系统时钟有误差(ntpdate -s time.nist.gov),12s的误差即可造成0peers
#  2. 防火墙的设置禁掉了UDP，这时你只能手动添加peers了
> net.listening
true
> net.peerCount
4
> admin.peers    // 获取peers详细信息

# 私有链，通过-—networkid可以设定私有链
$ geth --networkid=1983 console
# 可以通过--genesis指定一个创世块的JSON文件
$ geth --networkid=1983 --genesis 1983.json console


# 建立新的账户
$ geth account new    # 或者调用: personal.newAccount
$ geth account update # 升级账户存储文件的格式
$ geth cccount list   # 查看所有账户: eth.accounts

```

### 自定义创世块
```json 
{
  "alloc": {
    "dbdbdb2cbd23b783741e8d7fcf51e459b497e4a6": { 
        "balance": "1606938044258990275541962092341162602522202993782792835301376"
    },
    "e6716f9544a56c530d868e4bfbacb172315bdead": {
      "balance": "1606938044258990275541962092341162602522202993782792835301376"
    },
    ...
  },
  "nonce": "0x000000000000002a",
  "difficulty": "0x020000",
  "mixhash": "0x0000000000000000000000000000000000000000000000000000000000000000",
  "coinbase": "0x0000000000000000000000000000000000000000",
  "timestamp": "0x00",
  "parentHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
  "extraData": "0x",
  "gasLimit": "0x2fefd8"
}
```
### 设定静态节点
<datadir>/static-nodes.json:
```json
[
  "enode://f4642fa65af50cfdea8fa7414a5def7bb7991478b768e296f5e4a54e8b995de102e0ceae2e826f293c481b5325f89be6d207b003382e18a8ecba66fbaf6416c0@33.4.2.1:30303",
  "enode://pubkey@ip:port"
]
```

## abigen
## bootnode
## evm
## gethrpctest
## rlpdump
## swarm
## puppeth
