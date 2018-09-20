---
title: vmstat
tags:
---
# vmstat
<!-- toc -->
The command vmstat reports information about processes, memory, paging, block IO, traps, and cpu activity.
```sh
$ vmstat -m
Cache                       Num  Total   Size  Pages
ext4_groupinfo_4k          3744   3744    104     39
ext4_inode_cache          27112  31486    616     26
ext4_xattr                  680    680     48     85
ext4_free_data             9180   9180     40    102
ext4_allocation_context     288    288    112     36
ext4_prealloc_space         448    448     72     56Cache                       Num  Total   Size  Pages
ext4_groupinfo_4k          3744   3744    104     39
ext4_inode_cache          27112  31486    616     26
ext4_xattr                  680    680     48     85
ext4_free_data             9180   9180     40    102
ext4_allocation_context     288    288    112     36
ext4_prealloc_space         448    448     72     56
```
## 各监控数据的含义
 * Procs:
   * r: （running）运行队列中的进程数
   * b: （blocked）被阻塞的进程数，通常是进程等待IO动作
 * Memory(单位kB):
   * free:  系统空闲可用内存数
   * mapped: 设备和文件映射的内存大小
   * anon:  未映射的内存大小
   * slab:  内核数据缓存的大小
 * system:
   * in:  interrupts，在delay的时间内(默认是1秒)系统产生的中断数。
   * cs:  context switch，在delay的时间内(默认是1秒)系统的上下文切换次数。
   * flt:  major page faults，在delay的时间内(默认是1秒)系统产生的缺页错误次数。
 * cpu:
   * us: user time 用户部分消耗的CPU时间百分比
   * ni: 被提高优先级的进程在用户模式下的执行CPU时间百分比
   * sy: system time 系统内核部分消耗的CPU时间百分比
   * id: idle time 系统空闲时间百分比
   * wa: IO wait CPU在等待IO完成时间占比(do nothing but wait)
   * ir: interrupt 系统中断服务占比CPU百分比

## 如何分析
4. 问题现象与分析

## 运行进程队列(r)
 * 运行队列进程数远大于1（即CPU个数）
 * 说明系统运行的比较慢或系统比较忙，有线程在排队等待运行。
此时要根据运行环境/loading及系统整体状况(如频率设置是否有问题，系统中断是否过多，是不是跑了特别耗费资源的程序等)，分析系统是否正常

## 阻塞进程队列(b)
 * 运行队列进程数大于0
 * b值过大，等待IO的进程可能有很多
问题定位及分析方法
需要确定运行场景，是否有过多的文件操作(如浏览网页，市场下载安装apk等)，是否由于文件读写过慢或文件缓存被频繁刷出(如某个模块请求大量内存，刷出了文件cache)导致。

## 空闲内存(free)：
 * 该值忽大忽小，或比较小（如8M以下）
 * 可能有潜在的文件cache被换出
Android对Memory的使用，和Windows上不一致，Linux是尽量将memory都利用起来，当内存不够时，然后由oom killer来根据优先级将不太重要的后台cached进程杀掉以释放内存。Free值不是越大越好，而是在内存最大利用率时最好。当该值忽大忽小时，需要注意是否有文件cache被清理出来，文件cache对系统性能和流畅性影响很大。

## 中断个数(in)
 * 在系统运行过程中，该值变化很大
 * 可能系统中断过多
701B正常主界面静置状态下，中断数在500左右，下载和玩游戏时，在2000以下，如果比这个值明显高很多(如高出一倍以上)，需要检查一下哪个模块是否产生了太多中断，查看方法是观察 /proc/interrupts 下的中断增长情况。

## 上下文切换(cs)
 * 在系统运行过程中，该值变化很大
 * 可能是系统中断过多
先定位中断问题，若中断没有问题，需要确认调度策略及时钟中断是否有影响。

## 缺页错误(flt)
 * 该值大于0
 * 可能是文件cache被刷出导致

## 用户空间CPU时间占比(us)
 * 该值较大
 * 可能用户进程太忙或死循环
结合top命令查看，到底是哪个进程或线程太忙了，如果占比太高(如95%或以上)，需要查找该进程代码中是否陷入死循环。

## 内核空间CPU时间占比(sy)
 * 该值较大
 * 可能系统比较忙或使用过多的忙等待
该值一般不会太大，很少有超过20%的情况，如果过高需要检查驱动或内核模块，是否代码实现有问题(如使用过多忙等待)。

## CPU空闲时间占比(id)
 * 该值较小或为0
 * 系统比较忙
需要综合其他几个CPU占比，综合考虑；该值若大，说明系统有冗余，此时系统响应应该会比较好。

## CPU等待IO时间占比(wa)
 * 该值大于0
 * 系统的IO开销较大，系统在等待IO
系统在等待IO的时间内CPU无事可做，只能等待IO完成才能继续工作，该值若大，对系统流畅性体验的冲击很大，如果该值达到5%，根据以往经验，用户可以感觉到很明显的卡顿现象。出现此问题时，需要结合其他数据，定位是文件读写过慢导致，还是由于系统文件缓存(很大部分是app的资源文件及代码文件)被换出导致。注意：这个问题不是绝对的，还要看使用情况，比如在做文件读写测试，虽然IOW很高，但可能对系统冲击就比较小

## 中断服务CPU时间占比(ir)
 * 该值较大
 * 中断占比太高，消耗资源，影响系统响应
结合中断次数，定位是由于中断过多还是中断处理太长导致，需要进一步的分析。在正常情况下，该值极少能达到1%，基本上看不到超过1%的情况

## 性能分析
绝大多数的性能问题可以归结为三类:
 * IO Bound
 * CPU Bound
 * Memory Bound

那么如何识别出这三类问题呢?
### IO Bound System
 * 高b高wa: 有大量的进程在等待IO

### CPU Bound System
 * 高r或高us: 系统的频于应付用户进程，run queue中的进程数不见低，而CPU时间基本都花在用户空间，产生了us比重远远高于sy

### Memory Bound System
 * 高so高si: 系统内存不够，会触发Swap开始工作，有大量的页面换入换出swap.


## 参考
 * http://www.helpmehost.com/linux/reading-vmstat-in-linux-part-1/
 * http://nonfunctionaltestingtools.blogspot.hk/2013/03/vmstat-output-explained.html

