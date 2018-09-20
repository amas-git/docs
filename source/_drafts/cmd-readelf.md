---
title: readelf
tags:
---
# readelf
```sh
#查看dynamic段
$ readelf -d
```
```cpp
main() {}
```
```sh
$ gcc main.c
```
```sh
# 查看ELFHeader
$ readelf -h a.out
ELF 头：
  Magic：  7f 45 4c 46 01 01 01 00 00 00 00 00 00 00 00 00 
  Class:                             ELF32
  Data:                              2's complement, little endian
  Version:                           1 (current)
  OS/ABI:                            UNIX - System V
  ABI Version:                       0
  Type:                              EXEC (可执行文件)
  Machine:                           Intel 80386
  Version:                           0x1
  入口点地址：              0x80482d0
  程序头起点：              52 (bytes into file)
  Start of section headers:          2316 (bytes into file)
  标志：             0x0
  本头的大小：       52 (字节)
  程序头大小：       32 (字节)
  Number of program headers:         8
  节头大小：         40 (字节)
  节头数量：         35
  字符串表索引节头： 32
# 查看SegmentHeaders
$ readelf -l a.out
Elf 文件类型为 EXEC (可执行文件)
入口点 0x80482d0
共有 8 个程序头，开始于偏移量52
程序头：
  Type           Offset   VirtAddr   PhysAddr   FileSiz MemSiz  Flg Align
  PHDR           0x000034 0x08048034 0x08048034 0x00100 0x00100 R E 0x4
  INTERP         0x000134 0x08048134 0x08048134 0x00013 0x00013 R   0x1
      [正在请求程序解释器：/lib/ld-linux.so.2]
  LOAD           0x000000 0x08048000 0x08048000 0x0054c 0x0054c R E 0x1000
  LOAD           0x00054c 0x0804954c 0x0804954c 0x00114 0x00118 RW  0x1000
  DYNAMIC        0x000558 0x08049558 0x08049558 0x000e8 0x000e8 RW  0x4
  NOTE           0x000148 0x08048148 0x08048148 0x00044 0x00044 R   0x4
  GNU_EH_FRAME   0x000470 0x08048470 0x08048470 0x0002c 0x0002c R   0x4
  GNU_STACK      0x000000 0x00000000 0x00000000 0x00000 0x00000 RW  0x4
 Section to Segment mapping:
  段节...
   00     
   01     .interp 
   02     .interp .note.ABI-tag .note.gnu.build-id .gnu.hash .dynsym .dynstr .gnu.version .gnu.version_r .rel.dyn .rel.plt .init .plt .text .fini .rodata .eh_frame_hdr .eh_frame 
   03     .init_array .fini_array .jcr .dynamic .got .got.plt .data .bss 
   04     .dynamic 
   05     .note.ABI-tag .note.gnu.build-id 
   06     .eh_frame_hdr 
   07 
# 查看SectionHeaders
$ readelf -S a.out 
There are 35 section headers, starting at offset 0x90c:
Section Headers:
  [Nr] Name              Type            Addr     Off    Size   ES Flg Lk Inf Al
  [ 0]                   NULL            00000000 000000 000000 00      0   0  0
  [ 1] .interp           PROGBITS        08048134 000134 000013 00   A  0   0  1
  [ 2] .note.ABI-tag     NOTE            08048148 000148 000020 00   A  0   0  4
  [ 3] .note.gnu.build-i NOTE            08048168 000168 000024 00   A  0   0  4
  [ 4] .gnu.hash         GNU_HASH        0804818c 00018c 000020 04   A  5   0  4
  [ 5] .dynsym           DYNSYM          080481ac 0001ac 000040 10   A  6   1  4
  [ 6] .dynstr           STRTAB          080481ec 0001ec 000045 00   A  0   0  1
  [ 7] .gnu.version      VERSYM          08048232 000232 000008 02   A  5   0  2
  [ 8] .gnu.version_r    VERNEED         0804823c 00023c 000020 00   A  6   1  4
  [ 9] .rel.dyn          REL             0804825c 00025c 000008 08   A  5   0  4
  [10] .rel.plt          REL             08048264 000264 000010 08   A  5  12  4
  [11] .init             PROGBITS        08048274 000274 000023 00  AX  0   0  4
  [12] .plt              PROGBITS        080482a0 0002a0 000030 04  AX  0   0 16
  [13] .text             PROGBITS        080482d0 0002d0 000184 00  AX  0   0 16
  [14] .fini             PROGBITS        08048454 000454 000014 00  AX  0   0  4
  [15] .rodata           PROGBITS        08048468 000468 000008 00   A  0   0  4
  [16] .eh_frame_hdr     PROGBITS        08048470 000470 00002c 00   A  0   0  4
  [17] .eh_frame         PROGBITS        0804849c 00049c 0000b0 00   A  0   0  4
  [18] .init_array       INIT_ARRAY      0804954c 00054c 000004 00  WA  0   0  4
  [19] .fini_array       FINI_ARRAY      08049550 000550 000004 00  WA  0   0  4
  [20] .jcr              PROGBITS        08049554 000554 000004 00  WA  0   0  4
  [21] .dynamic          DYNAMIC         08049558 000558 0000e8 08  WA  6   0  4
  [22] .got              PROGBITS        08049640 000640 000004 04  WA  0   0  4
  [23] .got.plt          PROGBITS        08049644 000644 000014 04  WA  0   0  4
  [24] .data             PROGBITS        08049658 000658 000008 00  WA  0   0  4
  [25] .bss              NOBITS          08049660 000660 000004 00  WA  0   0  4
  [26] .comment          PROGBITS        00000000 000660 000038 01  MS  0   0  1
  [27] .debug_aranges    PROGBITS        00000000 000698 000020 00      0   0  1
  [28] .debug_info       PROGBITS        00000000 0006b8 000042 00      0   0  1
  [29] .debug_abbrev     PROGBITS        00000000 0006fa 000037 00      0   0  1
  [30] .debug_line       PROGBITS        00000000 000731 000035 00      0   0  1
  [31] .debug_str        PROGBITS        00000000 000766 00005e 01  MS  0   0  1
  [32] .shstrtab         STRTAB          00000000 0007c4 000146 00      0   0  1
  [33] .symtab           SYMTAB          00000000 000e84 000490 10     34  52  4
  [34] .strtab           STRTAB          00000000 001314 000246 00      0   0  1
Key to Flags:
  W (write), A (alloc), X (execute), M (merge), S (strings)
  I (info), L (link order), G (group), T (TLS), E (exclude), x (unknown)
  O (extra OS processing required) o (OS specific), p (processor specific)
# 查看所有的Header, 相当于-h -l -S
$ readelf -e a.out
```
读取ELF中的信息
```
# 读取符号表
$ readelf -s a.out
# 读取重定位信息
$ readelf -r a.out
# 读取.dynamic段，与共享库相关的一些信息保存在这个段中
$ readelf -d a.out
```
ELF信息转储:
```sh
$ readelf -x<N> a.out
# N是section number, 可以使用readelf -S来查看
$ readelf -x1 a.out
Hex dump of section '.interp':
  0x08048134 2f6c6962 2f6c642d 6c696e75 782e736f /lib/ld-linux.so
  0x08048144 2e3200                              .2.
# 也可以使用objdump来转储指定的section
$ objdump -s -j .interp
a.out:     file format elf32-i386
Contents of section .interp:
 8048134 2f6c6962 2f6c642d 6c696e75 782e736f  /lib/ld-linux.so
 8048144 2e3200                               .2
```
# 返汇编
```sh
$ objdump -d a.out
08048454 <_fini>:
 8048454:       push   %ebx
 8048455:       sub    $0x8,%esp
 8048458:       call   8048300 <__x86.get_pc_thunk.bx>
 804845d:       add    $0x11e7,%ebx
 8048463:       add    $0x8,%esp
 8048466:       pop    %ebx
 8048467:       ret
# 不想看十六进制指令，可以使用
$ objdump -d --no-show-raw-insn a.out
08048454 <_fini>:
 8048454:       53                      push   %ebx
 8048455:       83 ec 08                sub    $0x8,%esp
 8048458:       e8 a3 fe ff ff          call   8048300 <__x86.get_pc_thunk.bx>
 804845d:       81 c3 e7 11 00 00       add    $0x11e7,%ebx
 8048463:       83 c4 08                add    $0x8,%esp
 8048466:       5b                      pop    %ebx
 8048467:       c3                      ret  
# 显示每条指令的偏移量
$ objdump -d --prefix-addr a.out
Disassembly of section .fini:
08048454 <_fini> push   %ebx
08048455 <_fini+0x1> sub    $0x8,%esp
08048458 <_fini+0x4> call   08048300 <__x86.get_pc_thunk.bx>
0804845d <_fini+0x9> add    $0x11e7,%ebx
08048463 <_fini+0xf> add    $0x8,%esp
08048466 <_fini+0x12> pop    %ebx
08048467 <_fini+0x13> ret 
# 反汇编指定的section
$ objdump -d -j .fini a.out
# 在包含调试信息的情况下打印行号
$ objdeump -d -l a.out
...
080483d0 <main>:
main():
/data/src/c/main.c:1
 80483d0:       55                      push   %ebp
 80483d1:       89 e5                   mov    %esp,%ebp
/data/src/c/main.c:2
 80483d3:       5d                      pop    %ebp
 80483d4:       c3                      ret    
 80483d5:       66 90                   xchg   %ax,%ax
 80483d7:       66 90                   xchg   %ax,%ax
 80483d9:       66 90                   xchg   %ax,%ax
 80483db:       66 90                   xchg   %ax,%ax
 80483dd:       66 90                   xchg   %ax,%ax
 80483df:       90                      nop
...
# 如果既包含调试信息又能找到源码，可以使用-S来查看反汇编
$ objdump -d -S a.out
080483d0 <main>:
main() {
 80483d0:       55                      push   %ebp
 80483d1:       89 e5                   mov    %esp,%ebp
}
 80483d3:       5d                      pop    %ebp
 80483d4:       c3                      ret    
 80483d5:       66 90                   xchg   %ax,%ax
 80483d7:       66 90                   xchg   %ax,%ax
 80483d9:       66 90                   xchg   %ax,%ax
 80483db:       66 90                   xchg   %ax,%ax
 80483dd:       66 90                   xchg   %ax,%ax
 80483df:       90                      nop
```

