



# 行业动态 2018-10-23

## BCH分叉 (谨慎支持)

### 背景

- https://cn.bitcoin.com/archives/14287

### 交易支持情况

此前宣称已获得两家交易所支持

- [CoinEx 9月1日发布声名支持BitcoinSV](https://www.coinex.com/announcement/detail?id=124&lang=en_US)

```
1. On the basis of “snapshot” during the potential forking, we will allocate BSV to your accounts against a 1:1 ratio on your BCH assets.

2. After the fork, we will release a BSV/BCH trading pair in CoinEx so you can buy or sell BSV.

3. CoinEx will not support BSV deposit and withdrawal until its chain is stabilized and all supporting services e.g. wallet are available. Please refer to our official announcement for more information.
```

> 该交易所只是发布了声名, 并没有后续动作,至少交易币对还没有打开

- 比特亚洲(https://www.bitasiaex.com/notice/index.html) 

根据Google缓存网页的结果, 可以发现确实有这么回事, 但是目前这篇9月2号的文章已经被删除
```
Sep 2, 2018 - We are more than happy to see more teams (Bitcoin SV) join the community to compete and inject value into the BCH community and bring the 
```


### 官方动态

#### 里程碑:

1. 11月15日发布正式版本
2. 10月22日SV矿池对外开放 (https://unhashed.com/cryptocurrency-news/bitcoin-cash-mining-pool-svpool-opens-public/)
3. 09月07日发布测试版本

```
Roadmap
=========
A comprehensive roadmap is in development, but the first goal is to have the initial Bitcoin SV release ready for testing by first week of September 2018.  The code will be based off Bitcoin ABC v0.17.2.  The initial release will contain as minimal a changeset as possible to support the November 15, 2018 BCH protocol upgrade because the first priority for Bitcoin SV is to establish security and QA best practices.
```

#### 官方公告:

- 10月12日: [A BCH Protocol Implementation of the Satoshi Vision](https://nchain.com/en/blog/bitcoin-sv-bch-protocol-implementation-satoshi-vision/)
- 08月16日: [Bitcoin SV Full Node Implementation Launched to Fully Restore Original Bitcoin Protocol](



### 负面新闻

- https://cryptonomist.ch/en/2018/08/24/bitcoin-sv/


## 交易所动态

|          |                                         |      |
| -------- | --------------------------------------- | ---- |
| 10月24日 | DCR/BNB                                 | 币安 |
| 10月15日 | TUSD, USDC, GUSD, PAX / BTC, USDT       | OKEx |
| 10月17日 | VNT Chain (VNT) 暂缓开通VNT/ETH交易比对 | OKEx |

- [USDC稳定币](https://www.circle.com/cn/USDC)
- [VNT](http://www.vntchain.io/)
- VNT在CMC上并无介绍
- https://nchain.com/en/blog/bitcoin-sv-launch/)





<<<<<<< HEAD
https://nimiq.com/



=======
>>>>>>> aeb5a35156dca0febbadf4a0aca7aaca3598e648
## DCR: Decred

- https://www.decred.org/



## BCD: 比特币钻石

- https://btcd.io/zh-hans/



### RVN: 

- 于2018年7月3日发布, 这天是比特币发布9周年. 该项目修改改了出块时间和挖矿算法(X16R).

- https://ravencoin.org/

- [X16R](https://ravencoin.org/wp-content/uploads/2018/01/X16R-Whitepaper-3.pdf)

|          |        |      |
| -------- | ------ | ---- |
| Bitcoin  | SHA256 |      |
| Litecoin | Scrypt | 目前还没有对应的ASIC矿机 |
| Ethereum | Equihash |                          |
| DarkCoin / Dash | X11    |  |
| Dash     | X13    |      |
| Dash     | X13    |      |
| Rvencoin    | X16R |      |



挖矿算法对挖矿行业的影响:

> PCs-> GPUs -> FPGAs -> ASICs

- 算力等于利益
- 为了抵御ASICs, 有两种方法
  - 采用内存敏感的挖矿算法(如: Scrypt, Equihash)
  - 增加ASICs的开发成本, 比如采用哈希串联(X11)

### X11算法

X11算法是一种POW哈希函数, Evan Duffield最早发明在并在DarkCoin中使用. 后来的Dash币最初也采用这种算法. 

X11算法是由11种哈系算法串联而成的, 这就是11的含义. 这11种哈希算法是:

1. BLAKE
2. BLUE MIDNIGHT WISH (BMW)
3. Grøst
4. JH
5. Keccak
6. Skein
7. Luffa
8. CubeHash
9. SHAvite-3
10. SIMD
11. ECHO

这些算法你可以在[NIST](https://www.nist.gov/)查到, 是美国为了选择下一代哈希算法而征集的. 每一种算法都经过了第一轮的竞争. 为了搞定X11你需要支持11种哈希算法, 这样成本就大大提高了.

X11有两个优点:

1. 比SHA256更加安全, 因为毕竟采用了11种哈希算法.
2. 有一定的抗ASICs性, 但也不是做不出来.

好了, 明白X11之后你可以去写白皮书了, 发明一个X666算法应该不难.


