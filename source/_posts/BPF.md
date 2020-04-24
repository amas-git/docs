# BPF

## 历史

- 1992: Steven McCanne和Van Jacobson发表了《The BSD PacketFilter: A New Architecture for User-Level Packat Capture》
  - 比当时最先进的PacketFilter快20倍
  - 有两个重要的发明
    - 设计了一种新的VM，可以有效利用基于寄存器的CPU
    - 每个应用分配单独的Buffer
- 2014: Alexei Starovoitov扩展了BPF, 针对现代硬件进行了优化，使其效率大幅提升4倍
  - 将BPF VM的寄存器从32位换成64位
  - 增加了寄存器的数量

- 2014年6月，BPF从内核空间暴露给用户空间，是重要转折点，BPF成为一个内核的子系统，使用范围从原来的网络栈扩展到整个操作系统。BPF程序不需要你对内核重新进行编译，而且可以保证不会崩溃，不会崩溃，不会崩溃。内核开发者随之添加了bpf系统调用



## 架构

- BPF VM, 可以为BPF程序运行提供隔离
- 高级预言编译成BPF指令
- BPF Verifier确保编译后的程序不会把内核搞死
- 一旦经过确认，既可以把BPF程序加载到内核空间中, 内核会利用JIT对BPF运行时进行性能提升
- 内核中有很多AttachmentPoints, 你的BPF程序可以附着到这些AP上实现你的目标
- bmap负责内核与用户空间的通讯



## 编写BFP程序