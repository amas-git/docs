# 数据库





## 索引

### unique index / ordered index : (value, pointer)

通常使用B+Tree，

### B+Tree Index

- index node: 包含p个指针和p-1个key values
- p叫做B+Tree的order
- leaf index node: 包含三个信息key, 记录ID(RID), 指向下个叶子的指针
- block pointer: leaf index node之间用block pointer连接起来， 用于提高
- height: 树的层级
- intermmediate node: 


### clustered index

### hash index

当物理数据库表是无序的时候， 可以使用hash table index提高性能

### bitmap index

通常做secondary indexing索引多个属性值, 在非常大的数据库中，  通常包含一个bit向量， 代表某个属性之是否存在，通常用于数据非常稀疏的场合



## Materialized Views

当数据是从多个表中聚合而成结果，通常保存在MaterializedViews中， 与普通视图不同，他们更像是另外一个表，可以用来缓存查询结果



## Partitioning

将数据拆分到不同的硬盘上减轻硬件的压力



### MDC

Multidimensional Clustering是一种允许数据同时	

### 数据压缩

可以让数据更有效利用磁盘空间

### Data Striping

将数据保存在多个磁盘上， 需要访问的时候可以将数据从多个磁盘上组织到一起

