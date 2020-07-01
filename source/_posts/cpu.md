# cpu



查看CPU信息:

```bashv
$ lscpu
Architecture:                    x86_64
CPU op-mode(s):                  32-bit, 64-bit
Byte Order:                      Little Endian
Address sizes:                   39 bits physical, 48 bits virtual
CPU(s):                          8
On-line CPU(s) list:             0-7
Thread(s) per core:              2
Core(s) per socket:              4
Socket(s):                       1
NUMA node(s):                    1
Vendor ID:                       GenuineIntel
CPU family:                      6
Model:                           158
Model name:                      Intel(R) Core(TM) i7-7700 CPU @ 3.60GHz
Stepping:                        9
CPU MHz:                         800.049
CPU max MHz:                     4200.0000
CPU min MHz:                     800.0000
BogoMIPS:                        7202.00
Virtualization:                  VT-x
L1d cache:                       128 KiB
L1i cache:                       128 KiB
L2 cache:                        1 MiB
L3 cache:                        8 MiB
NUMA node0 CPU(s):               0-7
Vulnerability Itlb multihit:     KVM: Mitigation: Split huge pages
Vulnerability L1tf:              Mitigation; PTE Inversion; VMX conditional cache flushes, SMT vulnerable
Vulnerability Mds:               Vulnerable: Clear CPU buffers attempted, no microcode; SMT vulnerable
Vulnerability Meltdown:          Mitigation; PTI
Vulnerability Spec store bypass: Vulnerable
Vulnerability Spectre v1:        Mitigation; usercopy/swapgs barriers and __user pointer sanitization
Vulnerability Spectre v2:        Mitigation; Full generic retpoline, STIBP disabled, RSB filling
Vulnerability Tsx async abort:   Vulnerable: Clear CPU buffers attempted, no microcode; SMT vulnerable
Flags:                           fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush dts acpi mmx fxsr sse sse2 ss ht tm pbe syscall nx pdpe1gb rdtscp lm constant_tsc art arch_perfmon 
                                 pebs bts rep_good nopl xtopology nonstop_tsc cpuid aperfmperf pni pclmulqdq dtes64 monitor ds_cpl vmx smx est tm2 ssse3 sdbg fma cx16 xtpr pdcm pcid sse4_1 sse4_2 x2apic movbe pop
                                 cnt tsc_deadline_timer aes xsave avx f16c rdrand lahf_lm abm 3dnowprefetch cpuid_fault invpcid_single pti tpr_shadow vnmi flexpriority ept vpid ept_ad fsgsbase tsc_adjust bmi1 hle
                                  avx2 smep bmi2 erms invpcid rtm mpx rdseed adx smap clflushopt intel_pt xsaveopt xsavec xgetbv1 xsaves dtherm ida arat pln pts hwp hwp_notify hwp_act_window hwp_epp

```





## Cache Misses

BY TYPE:

- Cold miss
- Conflict miss
- Capacity miss



功能分:

- D-Cache (Data Cache)
- I-Cache (Instruction Cache)
- TLB(Translation Lookaside Buffer): 保存虚拟地址到物理地址的映射关系



Cache的体系结构:

| 类型                 | 缓存对象        | 在哪里缓存 | 延迟（时钟周期） | 被谁管理 |
| -------------------- | --------------- | ---------- | ---------------- | -------- |
| Register             | 4-8 bytes words | CPU        | 0                | 编译器   |
| TLB                  | 地址转换        | TLB芯片    | 0                | 硬件     |
| L1 Cache             | 64bytes block   | L1芯片     | 1                | 硬件     |
| L2 Cache             | 64bytes block   | L2芯片     | 10               | 硬件     |
| Vitrual Memory       | 4KB page        | 主存       | 100              | 硬件+OS  |
| Page Cache           | 文件的一部分    | 主存       | 100              | OS       |
| Disk Cache           | Disk sectors    | 磁盘控制器 | 100,000          | 硬盘     |
| Network Buffer Cache | 文件的一部分    | 本地磁盘   | 10,000,000       | AFS/NFS  |
| 浏览器缓存           | 网页            | 本地磁盘   | 10,000,000       | 浏览器   |
| CDN                  | 网页            | 远程磁盘   | 1,000,000,000    | CDN      |

举个例子:

```
Intel Core i7-9xx处理器(4 Core 8 HW Thread)
 - 32KB  L1 I-Cache / core
 - 32KB  L1 D-Cache / core
 - 256KB L2 Cache   / core
 - 8MB   L3 Cache   / 4 core
 
 -----------------------------------|
 C                                  |
 O  T0  [L1 I-Cache]|               |
 R                  |----[L2 Cache] |
 E  T1  [L1 D-Cache]|               |
 0                                  |
 -----------------------------------|
 C                                  |
 O  T0  [L1 I-Cache]|               |
 R                  |----[L2 Cache] |
 E  T1  [L1 D-Cache]|               |
 1                                  |
 -----------------------------------|----[L3 Cache]----[Main Memory]
 C                                  |
 O  T0  [L1 I-Cache]|               |
 R                  |----[L2 Cache] |
 E  T1  [L2 D-Cache]|               |
 2                                  |
 -----------------------------------|
 C                                  |
 O  T0  [L1 I-Cache]|               |
 R                  |----[L2 Cache] |
 E  T1  [L2 D-Cache]|               |
 3                                  |
 -----------------------------------|
```



Cache性能指标(Cache Performance Metrics)

- Miss Rate

  - MissRate = 1 - HitRate
  - L1 Cache: 3-10% 
  - L2 Cache: < 1%

- Hit Time（单位：cycles）

  - 将一条cache line发送给处理器的时间（包含检测数据是否存在与cache line的时间）
  - L1：1-2 cycles
  - L2：5-20 cycles

- Miss Penalty（单位是：cycles）

  - 缓存失效导致增加的额外时间开销
  - 50-200 cycles, 从主存去拿数据

  

> 为什么使用Miss Rate?
>
> 1% Miss Rate = 99% Hit Rate
>
> 3% Miss Rate = 97% Hit Rate
>
> 97% Hit Rate： 1 cycle + 0.01 * 100 cycles = 2 cycles
>
> 99% Hit Rate： 1 cycle + 0.03 * 100 cycles = 4 cycles
>
> 结果就是97%的Hit Rate比99%的Hit Rate要平均慢一倍， 所以用Miss Rate来避免误导



> #### AMAT(Average Memory Access Time): 内存平均访问时间是包含Cache在内访问内存平均所需要的时间
>
> AMAT = HitTime + MissRate * MissPenalty



> L1 Cache关注降低HitTime, L2 Cache关注降低Miss Rate



## Cache的结构

Line是缓存管理的最小单位，也就是说主存和Cache之间的数据读写是以Line为单位的。

CPU从内存中加载Line的同时还会进行预加载(prefetching)

> Cache由Line构成，每个Line = n个bytes, 通常Intel和AMD的CPU的line是64bytes(64*8 = 512bit)
>
> ```sh
> # Linux操作系统可以通过以下方式获取cache和line的大小
> $ getconf LEVEL1_DCACHE_LINESIZE # 一级数据缓存的Line大小
> 64
> $ getconf LEVEL1_DCACHE_SIZE     # 一级指令缓存的大小，可知道L1缓存包含512个Line
> 32768
> 
> # 所有和cache相关的配置，可以这么看
> $ getconf -a | grep CACHE
> LEVEL1_ICACHE_SIZE                 32768
> LEVEL1_ICACHE_ASSOC                8
> LEVEL1_ICACHE_LINESIZE             64
> LEVEL1_DCACHE_SIZE                 32768
> LEVEL1_DCACHE_ASSOC                8
> LEVEL1_DCACHE_LINESIZE             64
> LEVEL2_CACHE_SIZE                  262144
> LEVEL2_CACHE_ASSOC                 4
> LEVEL2_CACHE_LINESIZE              64
> LEVEL3_CACHE_SIZE                  8388608
> LEVEL3_CACHE_ASSOC                 16
> LEVEL3_CACHE_LINESIZE              64
> LEVEL4_CACHE_SIZE                  0
> LEVEL4_CACHE_ASSOC                 0
> LEVEL4_CACHE_LINESIZE              0
> ```
>
> 



## Cache友好的设计

之于数据：

1. 使用连续内存（arrays 优于 lists）
   1. 根据缓存加载的策略，连续内存会更快速的加载到cache line中
   2. arrays搜索可以打败在Heap中分配的BSTs
   3. 基于有序arrays的二分查找可以打败在堆中的HashMap
   4. Big-O在大数据规模的情况下最终会取胜，但是最先主导性能的一定是Cache
2. 选择尽可能内存开销小的数据结构
3. 避免Cacheline ping-pong
   1. 如果可能，用同一个CPU进行写
   2. 将线程绑定到指定的CPU
4. 避免在不同线程之间共享line
5. small == fast
6. 编写prefetch-friendly的代码
7. 尽可能利用cache line
8. 我不懂数据结构，但是我知道数组可以打败他们
9. 小心多线程环境下的False Sharing问题

之于代码:

1. Fit working set in cache
2. Make “fast paths” branch-free sequences
3. Inline cautiously
4. 利用编译器的PGO和WPO
   - PGO：Profile-Guided Optimiztion, PGO通过减小代码体积，减少分支错误预测，识别代码布局来减少指令缓存的问题
     - PGO包含3个步骤
       1. 给目标程序插入代码
       2. 运行插入特殊指令的代码，产生一些运行时的追踪数据
       3. 根据产生的数据找到关键路径，采取优化 
   - WPO：Whole Program Optimization
     - 见: https://docs.microsoft.com/zh-cn/archive/msdn-magazine/2015/february/compilers-what-every-programmer-should-know-about-compiler-optimizations
5. 编写Brach-free 的代码: 无分支或少分支的代码



## 缓存的连贯性(Cache Coherency)

想一下：

- core 0 和 core 1都缓存了地址a, 当core 0修改了a的内容之后，core 1缓存的a也会随之更新（方法是core 1的a所在的cache line被更新，这需要一定时间的开销，应用程序无感知）
- 既然某个缓存的更新可以影响其他缓存，那么就会出现，多个core缓存了同一个line, 但是每个core更新line中的不同元素，这样导致每个core中的line频繁的同步, 结果就是这个line使用效率不高，这种现象叫做False Sharing, 是多线程环境下常见的性能杀手

### 预加载(Prefetching)

按照缓存内容分为：

- 数据预加载：数据访问模式的规律难以预测，因此比指令预加载更加困难
- 指令预加载：指令的执行模式可以预测，通过预加载提升性能

按照缓存预加载方式分为：

- 硬件预加载：硬件固化的算法，如StreamBuffers

- 软件预加载：通过编译器识别模式，插入预加载代码

  ```c++
  for (int i=0; i<1024; i++) {
      prefetch (array1 [i + k]); // k = MissPenalty / cycles(single iteration for loop)
      array1[i] = 2 * array1[i];
  }
  ```




## CPU的指令分支预测

分支预测，目的是让CPU流水线尽可能保持满负荷状态，即正在执行的线程以最高效率运行。



CPU流水线可分为4个阶段：

1. 指令预取
2. 指令解码
3. 指令执行
4. 结果回写

当程序中出现分支条件语句时， 下一条指令要等条件结果计算后才能确定。因此会出现CPU流水线无法对指令预取，也就是产生了流水线stall。CPU支持分支预测，就是想降低这种情况i下的stall。另外如果我们编程的时候



- IA-64： 几乎所有指令都是Predicated的
- X86-64：cmove指令
- ARM：几乎所有的指令都支持条件判断
  - ARM 15-bit Thumb指令集不支持分支检测

SPMD就是Single Program Multiple Data，硬件上多个ALU绑定在一起，共享一个PC, 就可以组成SPMD运算单元。如NV GPU的SMX, AMD的CU都属于SPMD运算单元。条件分支的出现会使SPMD运算单元的ALU执行路径不一致，由于ALU共享同一个PC, 所以ALU并行执行不同路径是不可能的。在SPMD上实现条件分支，需要编译器对条件分支程序进行等价变换，转化为brach-free的代码。

### 问题

> 1. 按照行遍历矩阵与按照列遍历同一个矩阵会有怎样的性能差距？



Loop1和Loop2哪个更快？(为什么是16？)

```c
int[] arr = new int[64 * 1024 * 1024];

// Loop 1
for (int i = 0; i < arr.Length; i++) arr[i] *= 3;

// Loop 2
for (int i = 0; i < arr.Length; i += 16) arr[i] *= 3;
```



K从1到1024

```c
for (int i = 0; i < arr.Length; i += K) arr[i] *= 3;
```



### 专家的声音

> Sergey Solyanik (from Microsoft): Linux was routing packets at ~30Mbps [wired], and wireless at ~20. Windows CE was crawling at barely 12Mbps wired and 6Mbps wireless. ... We found out Windows CE had a LOT more instruction cache misses than Linux. ... After we changed the routing algorithm to be more cache-local, we started doing 35MBps [wired], and 25MBps wireless - 20% better than Linux. 



> Jan Gray (from the MS CLR Performance Team): If you are passionate about the speed of your code, it is imperative that you consider ... the cache/memory hierarchy as you design and implement your algorithms and data structures.
>
> 

> Dmitriy Vyukov (developer of Relacy Race Detector): Cache-lines are the key! Undoubtedly! If you will make even single error in data layout, you will get 100x slower solution! No jokes

> Joe Duffy at Microsoft: During our Beta1 performance milestone in Parallel Extensions, most of our performance problems came down to stamping out false sharing in numerous places.

## 参考

- https://www.aristeia.com/TalkNotes/ACCU2011_CPUCaches.pdf

- https://software.intel.com/content/www/us/en/develop/documentation/cpp-compiler-developer-guide-and-reference/top/optimization-and-programming-guide/profile-guided-optimization-pgo.html

- [Branch Free Code](https://en.wikipedia.org/wiki/Branch_(computer_science)#Branch-free_code)

- [Branch Predication](https://homepages.dcc.ufmg.br/~douglas/research/pred_opt.html)

- [Gallery of Processor Cache Effects](http://igoro.com/archive/gallery-of-processor-cache-effects/)