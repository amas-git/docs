# ARK

方舟流动性系统设计概念

## 核心概念

```
ATOM
  NAMINGSPACE
  ID
  NAME
```



## 什么是流动性系统?





- 流动性监控: 节点的存量监控
- 流量监控: 节点/节点组之间的流量转移
- 节点组: 一组功能类似, 或处于一定目的组织在一起的节点



### 流动性的最小单元

NODE: The minium container for hoding one kind of asset

```
Li: 蠡/lí/流动性(Liquidity)存储的最小单元, 小勺子的意思
```

```
LI extends ATOM
  ID
  NAMINGSPACE
  NAME BTC
  AMO  100
  TAG  xrp
  STATUS [good|bad|disabled]
  TX extends ATOM
    $TIME $ID $AMO $TAG $SRC $DST
```





EDGE

```
EID
  LNODE
  RNODE
  DIR    [LR|RL|BIN]
  TIME_WINDOW: [1s|1m|15m|30m|...]
HIST:
  100 LR $TIME
  200 RL $TIME
```



## 系统概览

```

```

### URI

```
li://$name.$namespace.$host:$port/$id/@amount
li://$name.$namespace.$host:$port/$id/@amount_in
li://$name.$namespace.$host:$port/$id/@amount_out
```



##  指令

- buy A:B  9000@19.4
- sell  A:B
- view A 







### 询价

- 某个交易所的价格
- bitrue持有BTC的价格
- 

```
$ 1.5btc
$time $price(usd)
$ 1.5btc:trx
$time $price(trx)
$ btc 24hh
$ btc 24hm
$ li://bitrue/btc 24hm
$ li://hotbit/btc 1m
$ li://hotbit/btc/@price
$ 
```

