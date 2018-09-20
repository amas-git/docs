---
title: 理解动态链接技术
tags:
---
# 动态链接
静态链接，方法简单，但是不能大规模的使用。因为：
 1. 浪费资源。相同的代码不光消耗存储空间，而且浪费内存空间。
 2. 维护上很让人捉急，静态库的任何变动都要会导致整个程序重新编译，发布，相当不灵活。
前辈们引入了动态链接来弥补静态链接的不足。
# Dll Hell
动态链接的负作用之一便是让使用同一个共享库的不同程序之间产生了耦合。 如果缺乏成熟的动态库管理方法，就会出现当某个程序更新共享库之后导致其他依赖旧版本库的应用无法运行。这事儿在早期windows上时有发生，被成为DllHell.
# 如何才能实现动态链接
# DSO
能够动态载入VMA的目标文件(ObjectFile)成为DSO(DynamicSharedObjects). 在linux上通常以.so为后缀名。
# 延迟绑定(Lazy Binding)
在需要访问共享库代码的时候才使用动态链接器将动态库载入VMA的技术成为`延迟绑定`。
# 编写和使用共享库
lib.h:
```c
#ifndef LIB_H
#define LIB_H
    void func();
#endif
```
lib.c:
```c
#include <stdio.h>
void func(char * text) {
    puts(text);
}
```
p.c:
```c
#include "lib.h"
int main() {
    func("hello this is p");
    return 0;
}
```
q.c
```c
#include "lib.h"
int main() {
    func("hello this is q");
    return 0;
}
```
编译:
```sh
$ gcc -fPIC -shared lib.c -o lib.so
$ gcc p.c ./lib.so -o p
$ gcc q.c ./lib.so -o q
```
运行:
```sh
$ ./p
hello this is p
$ ./q
hello this is q
```
为了更加仔细的观察p和q的实际运行情况，我们来修改一下lib.c:
```c
#include <stdio.h>
void func(char * text) {
    puts(text);
    sleep(-1);
}
```
动态链接优势立现，我们只需要重新编译lib.so就可以了:
```sh
$ gcc -fPIC -shared lib.c -o lib.so
$ ./p &
[1] 1155
hello this is p                                                                                                                                                        
$ ./q &
[2] 1157
hello this is q 
```
接下来，我们分析一下p,q这两只进程:
||= p(pid=1155) =||= q(pid=1157) =||
```td
```
08048000-08049000  r-xp  00000000  08:17  9846308  /data/src/c/so/p
08049000-0804a000  rw-p  00000000  08:17  9846308  /data/src/c/so/p
b75b9000-b75ba000  rw-p  00000000  00:00  0
b75ba000-b7765000  r-xp  00000000  08:12  919634   /usr/lib/libc-2.17.so
b7765000-b7767000  r--p  001ab000  08:12  919634   /usr/lib/libc-2.17.so
b7767000-b7768000  rw-p  001ad000  08:12  919634   /usr/lib/libc-2.17.so
b7768000-b776b000  rw-p  00000000  00:00  0
b7783000-b7784000  rw-p  00000000  00:00  0
b7784000-b7785000  r-xp  00000000  08:17  9846310  /data/src/c/so/lib.so
b7785000-b7786000  rw-p  00000000  08:17  9846310  /data/src/c/so/lib.so
b7786000-b7787000  rw-p  00000000  00:00  0
b7787000-b7788000  r-xp  00000000  00:00  0        [vdso]
b7788000-b77a8000  r-xp  00000000  08:12  919676   /usr/lib/ld-2.17.so
b77a8000-b77a9000  r--p  0001f000  08:12  919676   /usr/lib/ld-2.17.so
b77a9000-b77aa000  rw-p  00020000  08:12  919676   /usr/lib/ld-2.17.so
bfc66000-bfc88000  rw-p  00000000  00:00  0        [stack]
```
```
```td
```
08048000-08049000  r-xp  00000000  08:17  9846309  /data/src/c/so/q
08049000-0804a000  rw-p  00000000  08:17  9846309  /data/src/c/so/q
b75f8000-b75f9000  rw-p  00000000  00:00  0
b75f9000-b77a4000  r-xp  00000000  08:12  919634   /usr/lib/libc-2.17.so
b77a4000-b77a6000  r--p  001ab000  08:12  919634   /usr/lib/libc-2.17.so
b77a6000-b77a7000  rw-p  001ad000  08:12  919634   /usr/lib/libc-2.17.so
b77a7000-b77aa000  rw-p  00000000  00:00  0
b77c2000-b77c3000  rw-p  00000000  00:00  0
b77c3000-b77c4000  r-xp  00000000  08:17  9846310  /data/src/c/so/lib.so
b77c4000-b77c5000  rw-p  00000000  08:17  9846310  /data/src/c/so/lib.so
b77c5000-b77c6000  rw-p  00000000  00:00  0
b77c6000-b77c7000  r-xp  00000000  00:00  0        [vdso]
b77c7000-b77e7000  r-xp  00000000  08:12  919676   /usr/lib/ld-2.17.so
b77e7000-b77e8000  r--p  0001f000  08:12  919676   /usr/lib/ld-2.17.so
b77e8000-b77e9000  rw-p  00020000  08:12  919676   /usr/lib/ld-2.17.so
bfe09000-bfe2b000  rw-p  00000000  00:00  0        [stack]
```
```
----
||= p =||
```td
```sh
$ readelf -l p
Elf file type is EXEC (Executable file)
Entry point 0x8048420
There are 8 program headers, starting at offset 52
Program Headers:
  Type           Offset   VirtAddr   PhysAddr   FileSiz MemSiz  Flg Align
  PHDR           0x000034 0x08048034 0x08048034 0x00100 0x00100 R E 0x4
  INTERP         0x000134 0x08048134 0x08048134 0x00013 0x00013 R   0x1
      [Requesting program interpreter: /lib/ld-linux.so.2]
  LOAD           0x000000 0x08048000 0x08048000 0x006bc 0x006bc R E 0x1000
  LOAD           0x0006bc 0x080496bc 0x080496bc 0x00120 0x00124 RW  0x1000
  DYNAMIC        0x0006c8 0x080496c8 0x080496c8 0x000f0 0x000f0 RW  0x4
  NOTE           0x000148 0x08048148 0x08048148 0x00044 0x00044 R   0x4
  GNU_EH_FRAME   0x0005e0 0x080485e0 0x080485e0 0x0002c 0x0002c R   0x4
  GNU_STACK      0x000000 0x00000000 0x00000000 0x00000 0x00000 RW  0x4
 Section to Segment mapping:
  Segment Sections...
   00     
   01     .interp 
   02     .interp .note.ABI-tag .note.gnu.build-id .gnu.hash .dynsym .dynstr .gnu.version .gnu.version_r .rel.dyn .rel.plt .init .plt .text .fini .rodata .eh_frame_hdr .eh_frame 
   03     .init_array .fini_array .jcr .dynamic .got .got.plt .data .bss 
   04     .dynamic 
   05     .note.ABI-tag .note.gnu.build-id 
   06     .eh_frame_hdr 
   07 
```
```
||= lib.so =||
```td
```sh
$ readelf -l lib.so
Elf file type is DYN (Shared object file)
Entry point 0x420
There are 6 program headers, starting at offset 52
Program Headers:
  Type           Offset   VirtAddr   PhysAddr   FileSiz MemSiz  Flg Align
  LOAD           0x000000 0x00000000 0x00000000 0x00610 0x00610 R E 0x1000
  LOAD           0x000610 0x00001610 0x00001610 0x00120 0x00124 RW  0x1000
  DYNAMIC        0x00061c 0x0000161c 0x0000161c 0x000e0 0x000e0 RW  0x4
  NOTE           0x0000f4 0x000000f4 0x000000f4 0x00024 0x00024 R   0x4
  GNU_EH_FRAME   0x000590 0x00000590 0x00000590 0x0001c 0x0001c R   0x4
  GNU_STACK      0x000000 0x00000000 0x00000000 0x00000 0x00000 RW  0x4
 Section to Segment mapping:
  Segment Sections...
   00     .note.gnu.build-id .gnu.hash .dynsym .dynstr .gnu.version .gnu.version_r .rel.dyn .rel.plt .init .plt .text .fini .eh_frame_hdr .eh_frame 
   01     .init_array .fini_array .jcr .dynamic .got .got.plt .data .bss 
   02     .dynamic 
   03     .note.gnu.build-id 
   04     .eh_frame_hdr 
   05
```
```
看起来so文件和可执行的elf文件没太大区别啊。
# 可执行文件的Entry Point为什么是0x80xxxxx呢？
```div class=note
On 386 systems, the text base address is 0x08048000, which permits a reasonably large stack below the text while still staying above address 0x08000000, permitting most programs to use a single second-level page table. (Recall that on the 386, each second-level table maps 0x00400000 addresses.)
See: [http://www.iecc.com/linker/linker04.html]
```
 1. 入口点即整个ELF的执行开始处，如果一切准备就绪，OS只需要把PC计数器设为入口点地址，整个程序就可以开始执行了
 2. ELF的EntryPoint是一个偏移量，这个值加上0x08000000， 就是整个程序的入口，或者说.text段的VMA.
用objdump -d来看一下p, .text段始于0x8048420=0x8048000+420, 所以整个ELF的入口函数并不是main函数，而是_start函数.
```
Disassembly of section .text:
08048420 <_start>:
 8048420:       31 ed                   xor    %ebp,%ebp
 8048422:       5e                      pop    %esi
 8048423:       89 e1                   mov    %esp,%ecx
 8048425:       83 e4 f0                and    $0xfffffff0,%esp
 8048428:       50                      push   %eax
 8048429:       54                      push   %esp
 804842a:       52                      push   %edx
 804842b:       68 b0 85 04 08          push   $0x80485b0
 8048430:       68 40 85 04 08          push   $0x8048540
 8048435:       51                      push   %ecx
 8048436:       56                      push   %esi
 8048437:       68 1c 85 04 08          push   $0x804851c
 804843c:       e8 bf ff ff ff          call   8048400 <__libc_start_main@plt>
 8048441:       f4                      hlt    
 8048442:       66 90                   xchg   %ax,%ax
 8048444:       66 90                   xchg   %ax,%ax
 8048446:       66 90                   xchg   %ax,%ax
 8048448:       66 90                   xchg   %ax,%ax
 804844a:       66 90                   xchg   %ax,%ax
 804844c:       66 90                   xchg   %ax,%ax
 804844e:       66 90                   xchg   %ax,%ax
```
# 如何生成地址无关代码
# 模块内部跳转
# 模块内部的数据访问
# 模块间数据访问
# 模块间跳转
# 动态链接器的入口 .interp Setion
```sh
$ objdump -s -j .interp p
Contents of section .interp:
 8048134 2f6c6962 2f6c642d 6c696e75 782e736f  /lib/ld-linux.so
 8048144 2e3200
```
动态链接器由ELF文件中的.interp段来设定，我们看到`/lib/ld-linux.so`就是我们寻找的动态链接器。
# .dynamic Section
```sh
$ readelf -d p
Dynamic section at offset 0x6c8 contains 25 entries:
  Tag        Type                         Name/Value
 0x00000001 (NEEDED)                     Shared library: [./lib.so]
 0x00000001 (NEEDED)                     Shared library: [libc.so.6]
 0x0000000c (INIT)                       0x80483b4
 0x0000000d (FINI)                       0x80485b4
 0x00000019 (INIT_ARRAY)                 0x80496bc
 0x0000001b (INIT_ARRAYSZ)               4 (bytes)
 0x0000001a (FINI_ARRAY)                 0x80496c0
 0x0000001c (FINI_ARRAYSZ)               4 (bytes)
 0x6ffffef5 (GNU_HASH)                   0x804818c
 0x00000005 (STRTAB)                     0x8048298
 0x00000006 (SYMTAB)                     0x80481c8
 0x0000000a (STRSZ)                      193 (bytes)
 0x0000000b (SYMENT)                     16 (bytes)
 0x00000015 (DEBUG)                      0x0
 0x00000003 (PLTGOT)                     0x80497bc
 0x00000002 (PLTRELSZ)                   24 (bytes)
 0x00000014 (PLTREL)                     REL
 0x00000017 (JMPREL)                     0x804839c
 0x00000011 (REL)                        0x8048394
 0x00000012 (RELSZ)                      8 (bytes)
 0x00000013 (RELENT)                     8 (bytes)
 0x6ffffffe (VERNEED)                    0x8048374
 0x6fffffff (VERNEEDNUM)                 1
 0x6ffffff0 (VERSYM)                     0x804835a
 0x00000000 (NULL)                       0x0
$ readelf -d lib.so
Dynamic section at offset 0x61c contains 24 entries:
  Tag        Type                         Name/Value
 0x00000001 (NEEDED)                     Shared library: [libc.so.6]
 0x0000000c (INIT)                       0x3a0
 0x0000000d (FINI)                       0x57c
 0x00000019 (INIT_ARRAY)                 0x1610
 0x0000001b (INIT_ARRAYSZ)               4 (bytes)
 0x0000001a (FINI_ARRAY)                 0x1614
 0x0000001c (FINI_ARRAYSZ)               4 (bytes)
 0x6ffffef5 (GNU_HASH)                   0x118
 0x00000005 (STRTAB)                     0x234
 0x00000006 (SYMTAB)                     0x154
 0x0000000a (STRSZ)                      189 (bytes)
 0x0000000b (SYMENT)                     16 (bytes)
 0x00000003 (PLTGOT)                     0x1710
 0x00000002 (PLTRELSZ)                   32 (bytes)
 0x00000014 (PLTREL)                     REL
 0x00000017 (JMPREL)                     0x380
 0x00000011 (REL)                        0x340
 0x00000012 (RELSZ)                      64 (bytes)
 0x00000013 (RELENT)                     8 (bytes)
 0x6ffffffe (VERNEED)                    0x310
 0x6fffffff (VERNEEDNUM)                 1
 0x6ffffff0 (VERSYM)                     0x2f2
 0x6ffffffa (RELCOUNT)                   3
 0x00000000 (NULL)                       0x0
```

