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
  NAME amas
  AMO  100
  TAG  xrp
  STATUS [good|bad|disabled]
  TX extends ATOM
    $TIME $ID $AMO $RENAMINGSPACE
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

