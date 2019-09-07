# Kafka

> Highthrough Distributed Message Broker



- 集群服务，水平扩展
- KSQL
- Retenion Based: 消息默认都是持久化的，可以配置消息存放几个副本



![](/src/amas-git/docs/source/_drafts/assets/2019-09-07-083202_861x426_scrot.png)



- The streaming layer is fault-tolerant
- Each stream processor node can hold state of its own
- Each stream processor can write and store local state

>Life is a series of natural and spontaneous changes. Don't resist them—that only
>creates sorrow. Let reality be reality. Let things flow naturally forward.
>
>-- 老子

## Kafka集群

> 本质上就是一个保存了消息，分布在各个机器上的文件集合

### 限流:quotas

kafaka可以进行吞吐量控制，可以很好的控制带宽，避免被压垮。

### Synchronous Request/Response的缺点

当服务逐渐增多，同步通讯就会产生连锁反应，比如某个服务超时引起关联服务超时，然后运维就需要调查究竟哪个服务是问题的根源。为了解决这个问题，通常需要用SLA来要求关键的服务。必须达到某个水平从而降低这种事情的发生几率。

另外一个解决方案就是异步请求/响应

如果你是获取一个图片，或者某些资源，同步请求/响应非常的合适。但如果是你点击购买，这样简单操作背后其实对应的是非常负责的处理，这个时候异步请求/相应就更合适。

### Strong Ordering Guarantees

	- kafka在一个partiion中保证前后顺序(key-based ordering)
	- 客户端在提交更新的时候需要给出partionkey
	- 那如何实现全局的顺序呢？你需要用一台机器存放这个partiion, 尽管受到单机限制，但是绝大多数场景都够用

### Commands, Events, Queries

- commands: 就是动作，请求某个服务处理某件事情，有时候会改变服务内部的状态
- events: 已经发生的事情，单向流动，不必非要响应(所谓的: fire&forget)
- queries: 查询，不会改变服务内部状态

相比commands和queries, event具有loose coupling的特点:

> Loose coupling reduces the number of assumptions two parties make about one
> another when they exchange information.
>
> 松耦合可以减少信息交换过程中双方需要做出的假设数量。
>
> 或者: 松耦合减少了修改一个组件对其他组件产生的影响。

event-driven有助于我们构建松耦合系统，但并不是说tight coupling就不好，取决与场景。另外，任何业务有部分核心数据是没法避免耦合的。

loose coumpling会增强Pluggablility, 这在系统变得复杂的时候更好维护和迭代开发。

> 其实并不存在真正的解耦，所有的解耦都是改变事物联系的方式，实体1和实体2具有某种联系，必须通过某种介质才能发生了联系，介质决定了耦合程度。举个简单的例子，比如老张只能通过Facebook联系到，而我又必须联系老张，可是Facebook并不是人人都在使用，那么我和老张的联系就是因为Facebook而耦合的，比如换成电话，那么这种耦合其实就降低了，因为电话人人都有，如空气，水，如此的普遍以至于人们认为使用代价很低，低耦合等于低使用成本。S1和S2之前通过同步接口通讯，二者是耦合的, 如果让S1和S2既有联系的能力，又不相互直接联系，那么只要让S1和S2通过kafka通讯即可。

## QA

### Kafka是异步的REST?

- sync resquest/response call是一种最简单和直白的通讯机制
- async request/response 也是一种通讯机制
- 二者都有各自适合的场景
- 一些简单的资源访问直接用REST，无状态HTTP更加合适
- 当你需要对某个资源进行广播，存储，Kafka会更适合

### Kafka和Service Bus有什么区别？

- SB通常使用同步HTTP, 所以消息比较小且返回迅速的场景比较适合, 吞吐低
- Kafka完全异步，吞吐高

### Kafka是数据库么？

- 分布式存储+KSQL看起来就是数据库嘛？
- KSQL背后对ContinualComputaion进行了大量优化，而不是BatchComputaion
- Kafka实时处理第一，数据存储只能算老二

### Kafka和Spark有什么区别？

- 在Spark看来，Kafka就是一个数据源，Spark可以把Kafka的数据加工整到HDFS或者其他数据库里
- Spark支持ML和GraphProcessing, 具有更强的分析能力，更具灵活性， 实时性不如kafka



## 参考

	- services grows gradually: 服务逐渐增长
	- request driven
	- event driven
	- loose coupling / tight coupling
	- Google SLA: https://cloud.google.com/compute/sla
	- CQRS(Command and Query Responsibility Segregation)

