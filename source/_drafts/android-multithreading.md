---
title: Android多线程编程
tags:
---
# Android多线程编程
<!-- toc -->
## Linux Thread Schduling
Linux线程的调度受两个东西的影响
 1. 优先级
 2. 控制组(cgroup/ControlGroup)
## 优先级
Java提供一个API可以修改优先级:
```
java.lang.Thread
    setPriority(int priority);
```
  * priority : 0-10 (值越大,优先级越高, 跟linux的nice正好相反)
Android也提供两个API:
```
android.os.Process
    Process.setThreadPriority(int priority);               
    Process.setThreadPriority(int threadId, int priority);
```
对应关系:
```
|| Linux Niceness || Java Priority ||
|| -8|| 10 (Thread.MAX_PRIORITY) ||
|| -7||                          ||
|| -6||  9                       ||
|| -5||  8                       ||
|| -4||  7                       ||
|| -3||                          ||
|| -2||  6
|| -1||
||  0||  5 (Thread.NORMAL_PRIORITY)
||  1||
||  2||
||  3||
||  4||
||  5||
||  6||
||  7||
||  9||
|| 10|| 4
|| 11||
|| 12||
|| 13|| 3
|| 14||
|| 15||
|| 16|| 2
|| 17||
|| 18||
|| 19|| 1 (Thread.MIN_PRIORITY)
```

## cgroup
Linux的cgroup用来控制CPU/内存资源的分配, 每个线程都属于某个cgroup.
Android定义了几个CGROUP, 其中比较重要的是ForegroundGroup和BackgroundGroup, 比如
```
# app a.m.a.s.demo 出于可见状态的时候
$ adb shell ps -P | grep a.m.a.s.demo
u0_a67    7348  1700  554272 56436 fg  ffffffff 00000000 S a.m.a.s.demo
# 当按Home键回到桌面后
$ adb shell ps -P | grep a.m.a.s.demo
u0_a67    7348  1700  544272 56436 bg  ffffffff 00000000 S a.m.a.s.demo
```
## Java线程间通信
## Java Pipe
 * PipeReader
 * PipeWriter
## SharedMemory(Heap)
这个比较直观, 多个线程之间可以共享同一个进程的数据.
## Signal
线程之间通过锁可以协作工作, 实际上也是隐含了一种通信方式. 这些都是依靠线程信号来通信的. 线程信号实在是太底层了,而且容易出错.因此实践中不会使用.
## Blocking Queue
多线程协同工作的常见模型就是生产者和消费者, BlockingQueue作为生产者和消费者通讯的渠道. 可以简化线程之间的通讯.
## Java Lock
## Intrinsic Lock 
synchronized基于这个, 每个对象实例都有Intrinsic Lock?
## Java Monitor
Intrinsic lock 是 Java的monitor, monitor 有三种状态:
 * BLOCKED
 * EXCUTING
 * WATING
```
什么是monitor, monitor是一种同步机制, 可以保证任一时间内, 只有一个线程可以执行临界区中的代码.
```
当一段代码被IntrinsicLock所保护的时候, 这段代码就在临界区内.
synchronized (this) {  // (1) Enter Monitor 
                       // (2) Acquire Lock
wait();                // (3) Release Lock & Wait
                       // (4) Acquire Lock After Signal (notify() /  notifyAll())
}                      // (5) Release Lock & Exit Monitor
## synchronized
### Method Level
```
synchronized void changeState() {
sharedResource++;
}
```
  * 可以保证整个方法都在Monitor中, 使用Object的内置锁, 粒度大, 使用简单.
### Block Level
```
void changeState() {
    synchronized(this) {
        sharedResource++;
    }
}
```
 * 可以保证指定区域在Monitor中, 使用Object的内置锁, 粒度可以控制, 只对必要的代码进行保护, 可以降低锁开销.
### Block Level With Other Object's IntrinsicLock
```
private final Object mLock = new Object();
void changeState() {
        synchronized(mLock) {
        sharedResource++;
    }
}
```
 * 使用另外一个对象的内置锁, 可以避免共用同一个对象的内置锁所产生的开销
### Method Level with Class's IntrinsicLock
```
synchronized static void changeState() {
    staticSharedResource++;
}
```
### Block Level with Class's IntrinsicLock
```
static void changeState() {
    synchronized(this.getClass()) {
    staticSharedResource++;
    }
}
```
通过以上几种方法, 开发者可以控制锁的粒度, 从而降低同步开销. 
### ReentrantLock 
ReentrantLock 和 synchronized 在语意级别上是等价的, 都是用于建立临界区, 保证任一时刻只有一个线程执行临界区中的代码.
### ReentrantReadWriteLock
ReentrantLock & synchronized 有时候会过分保护临界区的代码, 比如, 多个进程读取临界区中的状态但是没有写操作发生就是无害的. 
因此 ReentrantReadWriteLock 用来解决这种问题. 但是这个机制的实现太复杂了, 进入临界区比起ReentrantLock要增加很多计算开销, 
所以最好只在大量读线程, 而只有少量线程写, 并且你还是十分在意读性能的情况下才使用它.
### Android ps command
```
-t Shows thread information in the processes.
-x Shows time spent in user code (utime) and system code (stime) in “jiffies,” which typically is units of 10 ms.
-p Shows priorities.
-P Shows scheduling policy, normally indicating whether the application is executing in the foreground or background.
-c Shows which CPU is executing the process.
name|pid Filter on the application’s name or process ID. Only the last defined value is used.
```

