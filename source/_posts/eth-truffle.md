---
title: Truffle以太坊智能合约开发环境
date: 2018-01-15 16:14:35
tags:
---
# Truffle
```
Truffle v4.0.4 - a development framework for Ethereum

Usage: truffle <command> [options]

Commands:
  init      Initialize new Ethereum project with example contracts and tests
  compile   Compile contract source files
  migrate   Run migrations to deploy contracts
  deploy    (alias for migrate)
  build     Execute build pipeline (if configuration present)
  test      Run Mocha and Solidity tests
  debug     Interactively debug any transaction on the blockchain (experimental)
  opcode    Print the compiled opcodes for a given contract
  console   Run a console with contract abstractions and commands available
  develop   Open a console with a local TestRPC
  create    Helper to create new contracts, migrations and tests
  install   Install a package from the Ethereum Package Registry
  publish   Publish a package to the Ethereum Package Registry
  networks  Show addresses for deployed contracts on each network
  watch     Watch filesystem for changes and rebuild the project automatically
  serve     Serve the build directory on localhost and watch for changes
  exec      Execute a JS module within this Truffle environment
  unbox     Unbox Truffle project
  version   Show version number and exit

```
## 创建工程
 - contracts/: Directory for Solidity contracts
 - migrations/: Directory for scriptable deployment files
 - test/: Directory for test files for testing your application and contracts
 - truffle.js: Truffle configuration file


## traffle migrate
```
$ truffle migrate                                                                           ~bc/metacoin
Using network 'development'.

Running migration: 1_initial_migration.js
  Deploying Migrations...
  ... 0x70f8116df07e6a0d5af11f580e3d9aeb520083f1731aac2f0437bf47e32bd3a8
  Migrations: 0x8cdaf0cd259887258bc13a92c0a6da92698644c0
Saving successful migration to network...
  ... 0xd7bc86d31bee32fa3988f1c1eabce403a1b5d570340a3a9cdba53a472ee8c956
Saving artifacts...
Running migration: 2_deploy_contracts.js
  Deploying ConvertLib...
  ... 0x14b9754cdeb4310a13e9f9b1713d05b35738cf7535d711d08b081c10b671277b
  ConvertLib: 0x345ca3e014aaf5dca488057592ee47305d9b3e10
  Linking ConvertLib to MetaCoin
  Deploying MetaCoin...
  ... 0xeb23f169616cbc25e4c2175f9dd4da35bcdba60cf6172958f13a0541729f7e4b
  MetaCoin: 0xf25186b5081ff5ce73482ad761db0eb0d25abfbf
Saving successful migration to network...
  ... 0x059cf1bbc372b9348ce487de910358801bbbd1c89182853439bec0afaee6c7db
Saving artifacts...
```

## 交易本质
> APPLAY(S,Tx) -> S'
    - S : 状态
    - Tx: 交易

### Transaction & Call
    - 写区块链 -> Transaction / 改变状态
    - 读区块链 -> Call / 不改变任何状态

Transaction和Call的区别
||是否消耗Gas| 是否有返回值 | 是否立刻执行 |
|---|---|---|---|
| Transaction | 是 | 是，只返回交易ID | 否 |
| Call        | 否 | 是               | 是 |

## truffle：以太坊的瑞士军刀
    - [HOME](http://truffleframework.com/)
    - 开发，编译，链接以太坊智能合约
    - 合约自动化测试
    - 脚本化部署以及迁移
    - 私网，公网部署管理
    - 可以引入其他外部的nodejs包
    - 控制台
    - Script Runner
  
## truffle-contract
以太坊合约更好的封装，nodejs
```
#源代码
$ git clone https://github.com/trufflesuite/truffle-contract

# 安装
$ npm install truffle-contract
```

特点:
    - 同步控制流
    - 无回调机制
    - 可以设置交易的默认值
    - 返回日志，交易收据
