# 并发计算模型

> The world is concurrent
>
> Concurrency is the key to responsive systems.

###  CPU Bit级别的并行

为什么32bit架构的计算机要比8bit的更快？因为32bit可以同时处理更多的bit

### CPU指令级别的并行

- 现代CPU高度并行的
  - pipelining
  - out-of-order execution
  - speculative execution

但如今CPU单核的处理能力很难再有提升， 迫使我们转向利用多核心

### 数据并行(SIMD:Single Instruction Multiple Data)

一种能够在大规模数据上应用单一运算的解决方案。典型的例子是图像处理，GPU对位图进行重复运算，比如调整亮度，这就是为什么GPUs可以充当数据并行处理器。但这种方法并不适合所有的场合。



### Task并行化，多处理器

- Shared Memory System
- Distributed Memory System

通过内存通讯更加的迅速可靠，所以基于SharedMemory的多处理器编码简单很多。很多异常不需要特别处理。但超过一定数目的处理器，共享内存就成为扩展的瓶颈。你就必须使用分布式内存系统，这就需要编写程序的时候考虑更多的容错，比共享内存系统开发难度更高。



### Concurrent

## Threads + Locks

- Mutual Exclusion: 计数问题
- Race Condition
- Deadlocks
  - Dining philosophers

## Functional Programming

## The Clojure Way : Separating Identity And State

## Actors

### CSP

## Data Parallelism

## Lambda



