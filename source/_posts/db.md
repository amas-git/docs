# DB



## 数据库中常用的数据结构

## 分布式数据库

### 复制Replication

- Node
  - Leader
  - Follower
- p复制方式
  - 同步
  - 异步
  - 半同步

- Node Outage
  -  Leader故障
    - 如何判断故障？
      - 心跳超时
    - 如何处理？Fallover
      - Leader选举
      - 同步客户端转移写请求到新的Leader
    - 故障可能出现的问题？
      - 丢数据
      - 脑裂，多个节点认为自己都是Leader
      - 多长时间认为需要更换Leader?
  - Follower故障
    - catch-up recovery

### 复制日志的实现方法

- Statement-based（如Raw SQL）
  - 面临的问题？
    - 如何处理NOW(), RAND()?
    - 如何保证自增ID的顺序一致？
    - 如何控制触发器，存储过程再不同节点上执行的副作用？
  - 现状
    - MYSQL: 5.1之前以这种方式为主，如今转向基于行的复制
- WAL:Write-ahead Log
  - 现状: PostgreSQL, Oracle
- Logical log(row-based)
  - 主要解决日志与存储引擎解耦
  - 现状: MySQL可以配置binlog使用row-based
- Trigger-based
  - 数据发生改变的时候出发可以自定义的代码
  - 现状: PostgresSQL Bucardo. Oracle的Databus

### 复制延迟问题（Replication Lag)

- 基于1个Leader的复制
  - 在同步复制的架构下，Follower故障会导致复制失败，也会出现单点问题
  - 异步复制架构下，可能会从Follower节点读到过时的数据，只能保证最终一致性
  - read-after-write consistency或read-your-writes consistency，写主之后读从碰到复制延迟拿到了旧数据
    - 如何在基于Leader写的系统里做到read-after-write?
      - 当读一个可能被修改的数据时，总是从主节点读
      - 对于写频繁的系统，第一种方案可能会让Leader压力过大，如果可以知道ReplicationLag多久，那么可以等待一会从Follower读
      - 客户端可以记住最近修改的时间戳（本质上是写入顺序），如果发现比这个旧则不用
      - 如果是多数据中心，则问题变得更复杂，写Leader操作需要路由到包含Leader的数据中心
      - 跨设备访问带来了 *cross-device* read-after-write 一致性问题, 在一个设备上操作之后另外一个设备上要看到结果
        - 怎么知道最后的修改什么时间发生？需要一个中心化的时间戳
        - 多个设备的操作如何路由到同一数据中心？
  - 如何解决？
    - 事务, 使用者简单，操作代价极大

- 多Leader复制

  - Collaborative editing问题
  - 如何解决写冲突?
    - 同步冲突检测
    - 异步冲突检测
    - 避免冲突：
    - 写入亲和性，写到固定的Leader
  - 自动解决冲突？
    - *Conflict-free replicated datatypes* (CRDTs)
    - Mergeable persistent data structure
    - Operatioal transformation

  - 多Leader复制拓扑结构
    - 星形
    - 循环结构
    - All-to-All

- Leaderless复制

  - Dynamo-style databases
  - 客户端直接发送数据到几个节点
  - 无Leader的情况下，一个掉线的节点恢复之后怎么catch-up?
    - 客户端同时读多个节点，筛选最新的数据
    - Anti-entropy process整理数据
  - Quorums for reading and writing
    - n个node
    - w个写入成功
    - r个读成功
    - 通常配置为 w = r = (n+1) / 2 (向上取整)

##  分片

> Shard, Region, Tablet, VNode, vBucket = Partitioning



## 事务

> Transaction = ACID = Atomicity Consistency Isolation Durabliity
>
> A: 一组操作无法分割成更小的操作，不会有人看到这组操作的中间状态，另外这组操作要么全部成功，要么就跟没有发生一样
>
> C:
>
> 	-  Replication Consistency
> 	-  Consistency Hash
> 	-  CAP Theory's Consistency = 
>
> I: 并发执行的事务如果是操作同一组数据，不会相互影响，一个事务执行的一组写操作，在其他事务看来要么什么都没发生，要么这些操作已经执行完毕。I主要是来对抗并发带来的各种问题，是非常有难度是事情。有很多数据库提供弱I, 
>
> 	-  Read Commit Isolation
> 	- Snapshot Isolation (MVCC)
>
> D: 一旦写操作执行成功，则数据不会丢
>
> 	- 可以是落盘
> 	- 可以是写到WAL中
> 	- 可以是复制到足够多的节点上
> 	- 不管怎样不存在绝对可靠的D
>
> ACID每一个都很难实现，因此不要认为一个数据库的ACID就是理想状态，要看具体对ACID支持到什么样的程度

