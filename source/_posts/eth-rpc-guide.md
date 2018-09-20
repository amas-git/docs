---
title: ETH RPC GUIDE
date: 2018-01-05 20:04:57
tags:
---
## 
```zsh
$ geth --datadir eth-amas console --nodiscover --rpc                                                                                                                                                ~/.ethereum/ext
INFO [01-05|16:58:28] Starting peer-to-peer node               instance=Geth/v1.7.3-stable-4bb3c89d/linux-amd64/go1.9.2
INFO [01-05|16:58:28] Allocated cache and file handles         database=/home/amas/.ethereum/ext/eth-amas/geth/chaindata cache=128 handles=1024
INFO [01-05|16:58:28] Initialised chain configuration          config="{ChainID: 12345 Homestead: 0 DAO: <nil> DAOSupport: false EIP150: <nil> EIP155: 0 EIP158: 0 Byzantium: <nil> Engine: unknown}"
INFO [01-05|16:58:28] Disk storage enabled for ethash caches   dir=/home/amas/.ethereum/ext/eth-amas/geth/ethash count=3
INFO [01-05|16:58:28] Disk storage enabled for ethash DAGs     dir=/home/amas/.ethash                            count=2
INFO [01-05|16:58:28] Initialising Ethereum protocol           versions="[63 62]" network=1
INFO [01-05|16:58:28] Loaded most recent local header          number=0 hash=3b5eb5…99e545 td=20
INFO [01-05|16:58:28] Loaded most recent local full block      number=0 hash=3b5eb5…99e545 td=20
INFO [01-05|16:58:28] Loaded most recent local fast block      number=0 hash=3b5eb5…99e545 td=20
INFO [01-05|16:58:28] Loaded local transaction journal         transactions=0 dropped=0
INFO [01-05|16:58:28] Regenerated local transaction journal    transactions=0 accounts=0
INFO [01-05|16:58:28] Starting P2P networking
INFO [01-05|16:58:28] RLPx listener up                         self="enode://dc3f47a3869559b71760bfe06b4c969a5bdc4fa468d5fc4955eb4d81aa32322a20df5bff4bc081ba19be7ff06b77e8e00bc747c60e219b58605cf35bb5794ef5@[::]:30303?discport=0"
INFO [01-05|16:58:28] IPC endpoint opened: /home/amas/.ethereum/ext/eth-amas/geth.ipc  # IPC 接入点
INFO [01-05|16:58:28] HTTP endpoint opened: http://127.0.0.1:8545 # JSON-RPC HTTP 接入点
Welcome to the Geth JavaScript console!
```

```zsh
# 可以先测试一下8545端口是否可用
$ netcat -v 127.0.0.1 8545
localhost [127.0.0.1] 8545 open

$ curl -H "Content-Type: application/json" --data '{"jsonrpc":"2.0","method":"web3_clientVersion","params":[],"id":42}' http://127.0.0.1:8545
{"jsonrpc":"2.0","id":67,"result":"Geth/v1.7.3-stable-4bb3c89d/linux-amd64/go1.9.2"}
```

## 智能合约
### 啥是智能合约
智能合约就是一段可以被机器执行的代码

### EVM
EVM是用来执行智能合约的虚拟机, 执行智能合约需要消耗gas
* Solidity : JavaScript
* Serpent : Python Likee
* LLL : Lisp Like
* Mutan : 已经废弃

### DApps
任何人都可以向太坊网络中丢入自己的智能合约，如果把以太坊想象为一个应用市场，那么每个合约类似于一个App,
所以DApps很形象的说明了以太坊智能合约，他们是一个意思。

## DApps安装开发环境
## 安装testrpc
 * 公有链上发布DApps需要同步整个公有链，而且运行还需要以太币，对于穷人很不划算
 * 运行geth创建一个私有链，并在上面调试发布比较繁琐，这个后面再说
 * testrpc就是一个


## 安装truffle
truffle是号称最流行的以太坊合约开发，
```
$ sudo npm install -g truffle

# 安装solc
# 安装完truffle实际上就已经安装了solc, 不过有时候我们可能会使用solc的命令行工具，所以可以安装一下
$ sudo npm install -g solc
$ solcjs
```

## 建立truffle工程 

```
$ mkdir smart_contract
$ cd smart_contract
$ truffle init
Downloading...
Unpacking...
Setting up...
Unbox successful. Sweet!

Commands:

  Compile:        truffle compile
  Migrate:        truffle migrate
  Test contracts: truffle test

$ tree .
.
├── contracts
│   └── Migrations.sol
├── migrations
│   └── 1_initial_migration.js
├── test
├── truffle-config.js
└── truffle.js
```
> *.sol文件就是智能合约的源代码了

```js
pragma solidity ^0.4.17;

contract Migrations {
  address public owner;
  uint public last_completed_migration;

  modifier restricted() {
    if (msg.sender == owner) _;
  }

  function Migrations() public {
    owner = msg.sender;
  }

  function setCompleted(uint completed) public restricted {
    last_completed_migration = completed;
  }

  function upgrade(address new_address) public restricted {
    Migrations upgraded = Migrations(new_address);
    upgraded.setCompleted(last_completed_migration);
  }
}
```

```
$ sudo npm install -g ethereumjs-testrpc
$ testrpc
EthereumJS TestRPC v6.0.3 (ganache-core: 2.0.2)

Available Accounts
==================
(0) 0xab995a6bf26db4dfcfc927d01967ef5973e9da63
(1) 0xb04d5a66e078cb8e8f031c414619216c72d4f3a3
(2) 0xaf7f4656106b25c6093233ab66f40d8e03f93926
(3) 0xd584c07964850c66ce26efd69a7e6bc153e8ece6
(2) 0xaf7f4656106b25c6093233ab66f40d8e03f93926
(3) 0xd584c07964850c66ce26efd69a7e6bc153e8ece6
(4) 0x807845d6d44de79cda137068e3e72175ce0d27ac
(5) 0x60097a1896719d0b73233c6d0159984f65d7d55b
(6) 0xf9cb27e1018ec970d004bf7cf319e0218e63d739
(7) 0x0f8a0ed850e6b75a6bd712f7b0e03a4255b0e035
(8) 0x3cc21e470a8d900d1a932eb2383390fdb19fdfa1
(9) 0xa0df7ad1041d9eb64a96f8a02f28ffb202373f1b

Private Keys
==================
(0) 3505325226a3e46493312c7014f1fc471095f7b9ddaefb2b55c9ffdfd7630810
(1) c0c4b54755c0edd0ae85a9009faaf10e1fbccbd4ce092c0ca34a12acea00ed2c
(2) 386bf24b68a4381681bc386449772cc4518dddf46979f4801cf0795f19c4c81f
(3) 4b18f70f1d4159ce3c980ea139c66c7fc42c7f9e5af117327cababfc32d347cd
(4) db189ec770c5659b1a011ae3a014729cfb95acae756f23b8e45044450ad42bc5
(5) dc91f3ce1c33d89caed7a60b824dd3212eeec741509cc07a24cc5976b9b73393
(6) d08e72f7a511542ef787fd9e03cbf9b8738ccd1d03a7d70cb8933c458d26321e
(7) 68ad92ca07665653c7f23d3a38db23ca500f0d5070d3050372070e7648f4daf7
(8) 89c2328653b460fc6cbadd48ee8e5151a8f944fa190e443ab5d4a4a8a6497cff
(9) ffabcd6873502895dbaa13adb32fb7e7ea48082c9f0fd589fac081e594a9a2e9

HD Wallet
==================
Mnemonic:      reduce oblige various segment myth few extra gentle achieve flash tape small
Base HD Path:  m/44'/60'/0'/0/{account_index}

Listening on localhost:8545
```
编译:
```
$ truffle compile
Compiling ./contracts/Migrations.sol...
Writing artifacts to ./build/contracts

# 编译后的结果保存在build目录中
$ tree.build
build
└── contracts
    └── Migrations.json
```

运行:
```
$ truffle migrate
Error: No network specified. Cannot determine current network.
    at Object.detect (/usr/lib/node_modules/truffle/build/cli.bundled.js:41338:23)
    at /usr/lib/node_modules/truffle/build/cli.bundled.js:202239:19
    at finished (/usr/lib/node_modules/truffle/build/cli.bundled.js:41266:9)
    at /usr/lib/node_modules/truffle/build/cli.bundled.js:200593:14
    at /usr/lib/node_modules/truffle/build/cli.bundled.js:63299:7
    at /usr/lib/node_modules/truffle/build/cli.bundled.js:165077:9
    at /usr/lib/node_modules/truffle/build/cli.bundled.js:161676:16
    at replenish (/usr/lib/node_modules/truffle/build/cli.bundled.js:162196:25)
    at iterateeCallback (/usr/lib/node_modules/truffle/build/cli.bundled.js:162186:17)
    at /usr/lib/node_modules/truffle/build/cli.bundled.js:162161:16
# 这个因为我们没有指定合约发布的网络, 需要在truffle.js中指定
```
  * [truffle配置说明](http://truffleframework.com/docs/advanced/configuration>)



truffle.js:
```
module.exports = {
    // See <http://truffleframework.com/docs/advanced/configuration>
    // to customize your Truffle configuration!
    networks: {
        development: {
            host: "localhost",
            port: 8545,     // 这个端口改成testrpc监听的端口就可以
            network_id: "*" // Match any network id
        }
    }
}
```

```
$ truffle migrate 
Using network 'development'.

Network up to date.
```


### 编写智能合约
```
$ truffle create contract HelloWorld 
# 多了一个contracts/HelloWorld.sol
```

HelloWorld.sol:
```
pragma solidity ^0.4.4;

contract HelloWorld {
  function HelloWorld() {
    // constructor
  }
}
```

### 函数种类
|类型|
|---|---|---|
|Read-only (constant) functions |
|Transactional functions        |




### 在私有链上发布DApps
```zsh
# 先来建一个私有链，稍后我们在私有链上编写智能合约
$ geth init helloworld-genesis.json --datadir bc_helloworld
WARN [01-09|10:56:25] No etherbase set and no accounts found as default 
INFO [01-09|10:56:25] Allocated cache and file handles         database=/src/bc/bc_helloworld/geth/chaindata cache=16 handles=16
INFO [01-09|10:56:25] Writing custom genesis block 
INFO [01-09|10:56:25] Successfully wrote genesis state         database=chaindata                            hash=3b5eb5…99e545
INFO [01-09|10:56:25] Allocated cache and file handles         database=/src/bc/bc_helloworld/geth/lightchaindata cache=16 handles=16
INFO [01-09|10:56:25] Writing custom genesis block 
INFO [01-09|10:56:25] Successfully wrote genesis state         database=lightchaindata                            hash=3b5eb5…99e545

# 启动geth运行起这个ETH节点
$ geth --datadir bc_helloworld --nodiscover console
WARN [01-09|10:58:29] No etherbase set and no accounts found as default
INFO [01-09|10:58:29] Starting peer-to-peer node               instance=Geth/v1.7.3-stable-4bb3c89d/linux-amd64/go1.9.2
INFO [01-09|10:58:29] Allocated cache and file handles         database=/src/bc/bc_helloworld/geth/chaindata cache=128 handles=1024
INFO [01-09|10:58:29] Initialised chain configuration          config="{ChainID: 19830709 Homestead: 0 DAO: <nil> DAOSupport: false EIP150: <nil> EIP155: 0 EIP158: 0 Byzantium: <nil> Engine: unknown}"
INFO [01-09|10:58:29] Disk storage enabled for ethash caches   dir=/src/bc/bc_helloworld/geth/ethash count=3
INFO [01-09|10:58:29] Disk storage enabled for ethash DAGs     dir=/home/amas/.ethash                count=2
INFO [01-09|10:58:29] Initialising Ethereum protocol           versions="[63 62]" network=1
INFO [01-09|10:58:29] Loaded most recent local header          number=0 hash=3b5eb5…99e545 td=20
INFO [01-09|10:58:29] Loaded most recent local full block      number=0 hash=3b5eb5…99e545 td=20
INFO [01-09|10:58:29] Loaded most recent local fast block      number=0 hash=3b5eb5…99e545 td=20
INFO [01-09|10:58:29] Loaded local transaction journal         transactions=0 dropped=0
INFO [01-09|10:58:29] Regenerated local transaction journal    transactions=0 accounts=0
INFO [01-09|10:58:29] Starting P2P networking
INFO [01-09|10:58:29] RLPx listener up                         self="enode://116ea2d28bdee9cec96e613e2f7984bdde85e44f1703a47cff778cb5f8088a51a473db6ca0adcc934a9f656787e676b796cb6c76caa424b6dfc7e8025358700f@[::]:30303?discport=0"
INFO [01-09|10:58:29] IPC endpoint opened: /src/bc/bc_helloworld/geth.ipc
Welcome to the Geth JavaScript console!

instance: Geth/v1.7.3-stable-4bb3c89d/linux-amd64/go1.9.2
 modules: admin:1.0 debug:1.0 eth:1.0 miner:1.0 net:1.0 personal:1.0 rpc:1.0 txpool:1.0 web3:1.0
>
# 

```

## Ethereum 协议
    * Message Call
    * Transaction Call : A kind of message call


## 交易
### 1. 构造交易对象
交易对象大概是这个样子:
```
{
  from: ...,
  to: ...,
  value: ...,
  gas: ...,
  data: ...,
  gasPrice: ...,
  nonce: ...
}
```
| 字段 | 含义 | 作用 |
|---|---|---|---|
| from     | 
| to       |
| gas      |
| data     |
| gasPrice |
| nonce    |
### 2. 为交易签名
### 3. 本地交易合法性验证
### 4. 广播合法交易
### 5. 矿工节点接受交易
### 7. 矿工

## 参考
 * [](https://blog.zeppelin.solutions/the-hitchhikers-guide-to-smart-contracts-in-ethereum-848f08001f05)
 * [Truffle Release Notes](https://github.com/trufflesuite/truffle/releases/tag/v4.0.0)
 * https://medium.com/@mvmurthy/full-stack-hello-world-voting-ethereum-dapp-tutorial-part-1-40d2d0d807c2
 * [MetaMask](https://metamask.io/)
 * [Truffle使用手册](http://truffleframework.com/docs/getting_started/installation)
 * [TestRpc使用手册](https://github.com/trufflesuite/ganache-cli#install)
 * [Parity](https://www.parity.io/)
 * [Ganache](http://truffleframework.com/ganache/)
 * [以太坊交易的生命周期](https://medium.com/blockchannel/life-cycle-of-an-ethereum-transaction-e5c66bae0f6e)
 * [JavaScript版EVM:etherumjs-vm](https://github.com/ethereumjs/ethereumjs-vm)
