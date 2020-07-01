# GOLANG



```sh
$ go test -bench=. -benchtime 2s -count 2 -benchmem -cpu 4
```



## 长期的性能度量标准

- RED
  - Requests
  - Errors
  - Duration
- USE
  - Utilization
  - Saturation
  - Erros
- Goden Signal
  - Latency
  - Errors
  - Traffic
  - Saturation

## 常见操作的性能

| 操作                       | 所需的时间    |      |
| -------------------------- | ------------- | ---- |
| 执行指令                   | 1ns           |      |
| 读L1缓存                   | 0.5ns         |      |
|                            |               |      |
| 读L2缓存                   | 7ns           |      |
| 锁                         | 25ns          |      |
| 读主存                     | 100ns         |      |
| 网络发送2kb数据(1Gbps网络) | 20,000ns      |      |
| 从主存顺序读取1M数据       | 250,000ns     |      |
| 读机械硬盘                 | 8,000,000ns   |      |
| 从机械硬盘读连续的1M数据   | 20,000,000ns  |      |
| 美国到欧洲往返1个packet    | 150,000,000ns |      |

## 性能优化金字塔

|      |                          |                                                         |
| ---- | ------------------------ | ------------------------------------------------------- |
| 0    | 基于特定平台，架构的优化 |                                                         |
| 1    | Runtime                  |                                                         |
| 2    | 汇编                     |                                                         |
| 3    | 编译                     |                                                         |
| 4    | 构建                     |                                                         |
| 5    | 源代码                   |                                                         |
| 6    | 算法和数据结构           | O(1) > O(log n) > O(n) > O(n * log n) > O(n^2) > O(2^n) |
| 7    | 设计                     |                                                         |

## GC

- 算法: tri-color, mark-sweep

## 编译

```bash
# 交叉编译
# OXS
$ GOOS=darwin GOARCH=amd64 go build -o myapp.osx
# ARM
$ GOOS=linux GOARCH=arm go build -o myapp.linuxarm
# 
$ go tool dist list -json
```

## 测试和性能对比

- -benchtime $time
- -count $n
- -benchmem
- -cpu x,y,z

```go
import "testing"

func BenchmarkHello(b *testing.B) {
  
}
```



```
$ go test bench=. -benchtime 2s -count 2 -cpu 4
```



## GO调度器

 - M: Machine
    - Sechduling Stack; go
    - Thread Local Storage
    - P
 - P: Proccesor, 逻辑处理器，可通过GOMAXPROC设置
   	- P Id
      	- 可运行的G: runq
      	- M
         	- defer pool
   	- 可用的G: gFree
 - G: Goruntine
   	- 当前stack指针(stack.lo / stack.hi)
      	- stackguard0 / stackguard1
      	- M



## GO CHANNEL

- src: https://golang.org/src/runtime/chan.go#L32



> 什么情况下会panic:
>
> 1. 向已经closed的channel发送数据
> 2. 没有close的channel, 产生了block







## BIG O

```bash

# O(log(n))
1 + 1/2 + 1/3 + ... + 1/N
```



|                   |            |      |
| ----------------- | ---------- | ---- |
| Binary Search     | O(log n)   |      |
| Dictionary Search | O(log n)   |      |
| Quick Sort        | O(n log n) |      |
| Merge Sort        | O(n log n) |      |
| Heap Sort         | O(n log n) |      |
| Tim Sort          | O(n log n) |      |
| Bubble Sort       | O(n^2)      |      |
| Insertion Sort    | O(n^2)      |      |
| Selection Sort    | O(n^2)      |      |
| Recursive Fibonacci       | O(2^n)      |      |
| 汉诺塔    | O(n2)      |      |
| 旅行商问题       | O(n2)      |      |



排序算法复杂度

|                | 最佳       | 平均       | 最差   | 空间复杂度 |      |
| -------------- | ---------- | ---------- | ------ | ---------- | ---- |
| quick sort     | O(n log n) | O(n log n) | O(n^2) |            |      |
| merge sort     | O(n log n) | O(n log n) | O(n log n) |            |      |
| tim sort       | O(n) | O(n log n) | O(n log n) |            |      |
| heap sort      | O(n log n) | O(n log n) | O(n log n) |            |      |
| bubble sort    | O(n) | O(n^2) | O(n^2) |            |      |
| insertion sort | O(n) | O(n^2) | O(n^2) |            |      |
| selection sort | O(n^2) | O(n^2) | O(n^2) |            |      |
| tree sort      | O(n log n) | O(n log n) | O(n^2) |            |      |
| shell sort     | O(n log n) | ? | ? |            |      |
| bucket sort    | O(n+k) | O(n+k) | O(n^2) |            |      |
| radix sort     | O(nk) | O(nk) | O(nk) |            |      |
| counting sort  | O(n+k) | O(n+k) | O(n+k) |            |      |
| cube sort  | O(n) | O(n log n) | O(n log n) |            |      |

## 内存管理

基础:

- MMU: 硬件，虚拟地址到物理地址的转换
- 虚拟地址的作用
  - 允许给内存设置权限(rwx)
  - 使得内存更容易的交换到磁盘
  - 使得内存更加容易移动
  - 共享内存
- 度量:
  - VSZ(Virtual Memory Size) 一个进程可以访问的全部内存，单位KB, 包括swap
  - RSS(Resident Set Size): RAM中分配了多少，不包括swap (stack + heap + shared)

```bash
# 不打包libc,获得更小的size
$ go build -ldflags '-libgcc=none' simpleServer.go
$ go build -gcflag '-m'
$ go build -gcflag '-m -m'
$ go build -gcflag '-m -m -m'
$ go tool compile -help
```

- 源代码: https://golang.org/src/runtime/malloc.go

  

>  span: 大于8kb的连续内存

- mspan: 管理主要的xpan分配
  - next: 下一个span在列表中的位置
  - previous: 前一个span
  - list: span list用于debug
  - startAddr:
  - npages: span中包含的耶码

- mheap
  - lock
  - free
  - scav
  - sweepgen
  - sweepdone
  - sweepers
  
- 内存对象分类

  - Tiny: 小于16b

    - 分配算法:
      1. 如果P的M有空间，就用M的空间
      2. 找到一个已有的对象？然后扩展到8,4,2 byte?
      3. 把这个tiny对象放进去

  - Small: 16b ~ 32 kb

    - 向上对齐
    - 从P的mcache中找到一个足够用的mspan
    - 如果mcache中不够，则从mcentral中的mspan列表中拿一个新的mspan，如果没有则从mheap中分配出页，若果分配不出页，从OS中分配页出来，这个操作比较昂贵，至少一次获取1M的内存
    - 从mspan中释放一个对象的过程是
      - mspan不用的还给mcache
      - mspan idle, 没有任何对象占用，则还给mheap
      - mspan idle几个周期之后，mspan中的页还给OS

  - Large: 大于32byte

    - 大的对象不用mcache和mcetral, 直接使用mheap

    

>
>
>强制GC:
>
>```go
>FreeOSMemory()
>runtime.GC()
>```
>
> 



## CPU

- Cgo: go中调用c
- GPU
- CUDA on GCP
- CUDA





## 开发工具

- godoc : 文档工具
- gofmt: 代码格式化
- https://godoc.org/golang.org/x/perf/cmd/benchstat
- go:generate、
- viper (库)
- cobra  (库)
- 模板
  - text/template
  - html/tmp;late
  - sprig
- goweight: 分析依赖包的大小

## 参考

- The Go reference specification: https://golang.org/ref/spec
- How to write Go code: https://golang.org/doc/code.html
- Effective Go: https://golang.org/doc/effective_go.html

