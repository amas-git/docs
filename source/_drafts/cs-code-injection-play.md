---
title: Hello代码注入
tags:
---
# Code Injection
```sh
$ gcc -fPIC -shared shared.c -o shared.so
$ gcc app.c ./shared.so -o app
```
```
$./app
$ gdb
 Address Flags   Offset Device   Inode     Rss     Pss Referenced Anonymous    Swap  Locked Description
08048000  r-xp 00000000  08:17 9846325       4       4          4         0       0       0 /data/src/c/injection/app
08049000  rw-p 00000000  08:17 9846325       4       4          4         4       0       0 /data/src/c/injection/app
b753c000  rw-p 00000000  00:00       0       4       4          4         4       0       0 
b753d000  r-xp 00000000  08:12  919634     220       4        220         0       0       0 /usr/lib/libc-2.17.so
b76e8000  r--p 001ab000  08:12  919634       8       8          8         8       0       0 /usr/lib/libc-2.17.so
b76ea000  rw-p 001ad000  08:12  919634       4       4          4         4       0       0 /usr/lib/libc-2.17.so
b76eb000  rw-p 00000000  00:00       0       8       8          8         8       0       0 
b7706000  rw-p 00000000  00:00       0       4       4          4         4       0       0 
b7707000  r-xp 00000000  08:17 9846319       4       4          4         0       0       0 /data/src/c/injection/shared.so
b7708000  rw-p 00000000  08:17 9846319       4       4          4         4       0       0 /data/src/c/injection/shared.so
b7709000  rw-p 00000000  00:00       0       4       4          4         4       0       0 
b770a000  r-xp 00000000  00:00       0       4       0          4         0       0       0 [vdso]
b770b000  r-xp 00000000  08:12  919676     104       2        104         0       0       0 /usr/lib/ld-2.17.so
b772b000  r--p 0001f000  08:12  919676       4       4          4         4       0       0 /usr/lib/ld-2.17.so
b772c000  rw-p 00020000  08:12  919676       4       4          4         4       0       0 /usr/lib/ld-2.17.so
bf9a8000  rw-p 00000000  00:00       0      12      12         12        12       0       0 [stack]
                                       ======= ======= ========== ========= ======= ======= 
                                           396      74        396        60       0       0 KB
```
```
# O_RDWR = 2
(gdb) call open("injection.o",2)
$1 = 3 
(gdb) call mmap(0, 1012, 1 | 2 | 4, 1, 3, 0)
$2 = -1217376256
```
```sh
$ pmap -X 1568
1568:   ./app
 Address Flags   Offset Device   Inode     Rss     Pss Referenced Anonymous    Swap  Locked Description
08048000  r-xp 00000000  08:17 9846325       4       4          4         0       0       0 /data/src/c/injection/app
08049000  rw-p 00000000  08:17 9846325       4       4          4         4       0       0 /data/src/c/injection/app
0958d000  rw-p 00000000  00:00       0       4       4          4         4       0       0 [heap]
b753c000  rw-p 00000000  00:00       0       4       4          4         4       0       0
b753d000  r-xp 00000000  08:12  919634     256       9        256         4       0       0 /usr/lib/libc-2.17.so
b76e8000  r--p 001ab000  08:12  919634       8       8          8         8       0       0 /usr/lib/libc-2.17.so
b76ea000  rw-p 001ad000  08:12  919634       4       4          4         4       0       0 /usr/lib/libc-2.17.so
b76eb000  rw-p 00000000  00:00       0       8       8          8         8       0       0
b7705000  rwxs 00000000  08:17 9846320       0       0          0         0       0       0 /data/src/c/injection/injection.o
b7706000  rw-p 00000000  00:00       0       4       4          4         4       0       0
b7707000  r-xp 00000000  08:17 9846319       4       4          4         0       0       0 /data/src/c/injection/shared.so
b7708000  rw-p 00000000  08:17 9846319       4       4          4         4       0       0 /data/src/c/injection/shared.so
b7709000  rw-p 00000000  00:00       0       4       4          4         4       0       0
b770a000  r-xp 00000000  00:00       0       4       0          4         0       0       0 [vdso]
b770b000  r-xp 00000000  08:12  919676     104       6        104         4       0       0 /usr/lib/ld-2.17.so
b772b000  r--p 0001f000  08:12  919676       4       4          4         4       0       0 /usr/lib/ld-2.17.so
b772c000  rw-p 00020000  08:12  919676       4       4          4         4       0       0 /usr/lib/ld-2.17.so
bf9a8000  rw-p 00000000  00:00       0      12      12         12        12       0       0 [stack]
                                       ======= ======= ========== ========= ======= =======
                                           436      87        436        72       0       0 KB
```
```sh
$  readelf -r app 
Relocation section '.rel.dyn' at offset 0x3c8 contains 1 entries:
 Offset     Info    Type            Sym.Value  Sym. Name
08049824  00000406 R_386_GLOB_DAT    00000000   __gmon_start__
Relocation section '.rel.plt' at offset 0x3d0 contains 5 entries:
 Offset     Info    Type            Sym.Value  Sym. Name
08049834  00000207 R_386_JUMP_SLOT   00000000   sleep
08049838  00000307 R_386_JUMP_SLOT   00000000   puts
0804983c  00000407 R_386_JUMP_SLOT   00000000   __gmon_start__
08049840  00000507 R_386_JUMP_SLOT   00000000   __libc_start_main
08049844  00000607 R_386_JUMP_SLOT   00000000   func
```
 * fun : 
  * 08049844
  * R_386_JUMP_SLOT
```
# SB了，没输出符号表, 看一下func函数的地址
(gdb) p & func
$11 = (<text variable, no debug info> *) 0xb770754c <func>
```
# 成功注入
---- 
08048000  r-xp 00000000  08:17 9846325       4       4          4         0       0       0 /data/src/c/injection/app
08049000  rw-p 00000000  08:17 9846325       4       4          4         4       0       0 /data/src/c/injection/app
08bf1000  rw-p 00000000  00:00       0       4       4          4         4       0       0 [heap]
b75fa000  rw-p 00000000  00:00       0       4       4          4         4       0       0
b75fb000  r-xp 00000000  08:12  919634     256       9        256         4       0       0 /usr/lib/libc-2.17.so
b77a6000  r--p 001ab000  08:12  919634       8       8          8         8       0       0 /usr/lib/libc-2.17.so
b77a8000  rw-p 001ad000  08:12  919634       4       4          4         4       0       0 /usr/lib/libc-2.17.so
b77a9000  rw-p 00000000  00:00       0       8       8          8         8       0       0
b77c3000  rwxs 00000000  08:17 9846320       0       0          0         0       0       0 /data/src/c/injection/injection.o
b77c4000  rw-p 00000000  00:00       0       4       4          4         4       0       0
b77c5000  r-xp 00000000  08:17 9846319       4       4          4         0       0       0 /data/src/c/injection/shared.so
b77c6000  rw-p 00000000  08:17 9846319       4       4          4         4       0       0 /data/src/c/injection/shared.so
b77c7000  rw-p 00000000  00:00       0       4       4          4         4       0       0
b77c8000  r-xp 00000000  00:00       0       4       0          4         0       0       0 [vdso]
b77c9000  r-xp 00000000  08:12  919676     104       6        104         4       0       0 /usr/lib/ld-2.17.so
b77e9000  r--p 0001f000  08:12  919676       4       4          4         4       0       0 /usr/lib/ld-2.17.so
b77ea000  rw-p 00020000  08:12  919676       4       4          4         4       0       0 /usr/lib/ld-2.17.so
bfde4000  rw-p 00000000  00:00       0      12      12         12        12       0       0 [stack]
                                       ======= ======= ========== ========= ======= =======
                                           436      87        436        72       0       0 KB
```
(gdb) p & func
$3 = (void (*)()) 0xb77c554c <func>
# 查看 func的地址
(gdb) p & func
$3 = (void (*)()) 0xb77c554c <func>
```
```
$ readelf -r app                                                                                                                                      ~src/c/injection
Relocation section '.rel.dyn' at offset 0x3c8 contains 1 entries:
 Offset     Info    Type            Sym.Value  Sym. Name
08049824  00000406 R_386_GLOB_DAT    00000000   __gmon_start__
Relocation section '.rel.plt' at offset 0x3d0 contains 5 entries:
 Offset     Info    Type            Sym.Value  Sym. Name
08049834  00000207 R_386_JUMP_SLOT   00000000   sleep
08049838  00000307 R_386_JUMP_SLOT   00000000   puts
0804983c  00000407 R_386_JUMP_SLOT   00000000   __gmon_start__
08049840  00000507 R_386_JUMP_SLOT   00000000   __libc_start_main
08049844  00000607 R_386_JUMP_SLOT   00000000   func
# 看看这个VMA对应的物理地址, 正好是func的地址，木错拉
(gdb) p/x * 0x08049844
$4 = 0xb77c554c
接下来我们把0xb77c554c物理地址的值，改为hello_func应该就可以OK了
看一下injection.o:
Section Headers:
  [Nr] Name              Type            Addr     Off    Size   ES Flg Lk Inf Al
  [ 0]                   NULL            00000000 000000 000000 00      0   0  0
  [ 1] .text             PROGBITS        00000000 000034 000014 00  AX  0   0  4
  [ 2] .rel.text         REL             00000000 0003dc 000010 08     11   1  4
  [ 3] .data             PROGBITS        00000000 000048 000000 00  WA  0   0  4
  [ 4] .bss              NOBITS          00000000 000048 000000 00  WA  0   0  4
  [ 5] .rodata           PROGBITS        00000000 000048 000012 00   A  0   0  1
  [ 6] .comment          PROGBITS        00000000 00005a 000012 01  MS  0   0  1
  [ 7] .note.GNU-stack   PROGBITS        00000000 00006c 000000 00      0   0  1
  [ 8] .eh_frame         PROGBITS        00000000 00006c 000038 00   A  0   0  4
  [ 9] .rel.eh_frame     REL             00000000 0003ec 000008 08     11   8  4
  [10] .shstrtab         STRTAB          00000000 0000a4 00005f 00      0   0  1
  [11] .symtab           SYMTAB          00000000 00030c 0000b0 10     12   9  4
  [12] .strtab           STRTAB          00000000 0003bc 00001d 00      0   0  1
先得找到.text段， 来看一下， .text段位于.o的0x000034这个偏移, 因此:
$ objdump -d injection.o                                                                                                                          ~src/c/injection:[1]
injection.o:     file format elf32-i386
Disassembly of section .text:
00000000 <hello_func>:
   0:   55                      push   %ebp
   1:   89 e5                   mov    %esp,%ebp
   3:   83 ec 18                sub    $0x18,%esp
   6:   c7 04 24 00 00 00 00    movl   $0x0,(%esp)
   d:   e8 fc ff ff ff          call   e <hello_func+0xe>
  12:   c9                      leave  
  13:   c3                      ret    
  看到了把, .text段开始就是hello_func, 所以hello_func的地址应该是:
  injection.o装载地址 + 0x34
08048000      4K r-x--  /data/src/c/injection/app
08049000      4K rw---  /data/src/c/injection/app
08bf1000    132K rw---    [ anon ]
b75fa000      4K rw---    [ anon ]
b75fb000   1708K r-x--  /usr/lib/libc-2.17.so
b77a6000      8K r----  /usr/lib/libc-2.17.so
b77a8000      4K rw---  /usr/lib/libc-2.17.so
b77a9000     12K rw---    [ anon ]
b77c3000      4K rwxs-  /data/src/c/injection/injection.o
b77c4000      4K rw---    [ anon ]
b77c5000      4K r-x--  /data/src/c/injection/shared.so
b77c6000      4K rw---  /data/src/c/injection/shared.so
b77c7000      4K rw---    [ anon ]
b77c8000      4K r-x--    [ anon ]
b77c9000    128K r-x--  /usr/lib/ld-2.17.so
b77e9000      4K r----  /usr/lib/ld-2.17.so
b77ea000      4K rw---  /usr/lib/ld-2.17.so
bfde4000    136K rw---    [ stack ]
 total     2172K
(gdb) set *0x08049844 = 0xb77c3000 + 0x000034
好了，跳转已经生效了，但是我们还有些工作.
$  readelf -S injection.o
There are 13 section headers, starting at offset 0x104:
Section Headers:
  [Nr] Name              Type            Addr     Off    Size   ES Flg Lk Inf Al
  [ 0]                   NULL            00000000 000000 000000 00      0   0  0
  [ 1] .text             PROGBITS        00000000 000034 000014 00  AX  0   0  4
  [ 2] .rel.text         REL             00000000 0003dc 000010 08     11   1  4
  [ 3] .data             PROGBITS        00000000 000048 000000 00  WA  0   0  4
  [ 4] .bss              NOBITS          00000000 000048 000000 00  WA  0   0  4
  [ 5] .rodata           PROGBITS        00000000 000048 000012 00   A  0   0  1
  [ 6] .comment          PROGBITS        00000000 00005a 000012 01  MS  0   0  1
  [ 7] .note.GNU-stack   PROGBITS        00000000 00006c 000000 00      0   0  1
  [ 8] .eh_frame         PROGBITS        00000000 00006c 000038 00   A  0   0  4
  [ 9] .rel.eh_frame     REL             00000000 0003ec 000008 08     11   8  4
  [10] .shstrtab         STRTAB          00000000 0000a4 00005f 00      0   0  1
  [11] .symtab           SYMTAB          00000000 00030c 0000b0 10     12   9  4
  [12] .strtab           STRTAB          00000000 0003bc 00001d 00      0   0  1
Key to Flags:
  W (write), A (alloc), X (execute), M (merge), S (strings)
  I (info), L (link order), G (group), T (TLS), E (exclude), x (unknown)
  O (extra OS processing required) o (OS specific), p (processor specific)
$ readelf -r injection.o
Relocation section '.rel.text' at offset 0x3dc contains 2 entries:
 Offset     Info    Type            Sym.Value  Sym. Name
00000009  00000501 R_386_32          00000000   .rodata
0000000e  00000a02 R_386_PC32        00000000   puts <-- puts中使用的常量
Relocation section '.rel.eh_frame' at offset 0x3ec contains 1 entries:
 Offset     Info    Type            Sym.Value  Sym. Name
00000020  00000202 R_386_PC32        00000000   .text
```
Note that system and print relocations are of R_386_PC32 type. This means that the value (resolved address) to be set into the relocation location should be calculated relatively to the PC program counter, that is relatively to the relocation location. Also R_386_PC32 relocation requires that the value that was stored in the relocation location before relocation resolution (addend) should be added to the resolved address. The R_386_32 .rodata relocation also adds the addend to its resolved address.
```
# 重新确定puts函数的地址
(gdb) p & puts
$8 = (<text variable, no debug info> *) 0xb7662e40 <puts>
# 既然puts函数位于0xb7662e40, injection.o的装载位置为: 0xb77c3000 .text段位于0x000034, .rodata段位于.text段起始处偏移00000009， 再 
(gdb) set *(0xb77c3000 + 0x000034 + 0x0000000e) = 0xb7662e40 - (0xb77c3000 + 0x000034 + 0x0000000e) - 4
# 重新确定.rodata段的地址
(gdb) set *(0xb77c3000 + 0x000034 + 0x00000009) = 0xb77c3000 + 0x000048
(gdb) quit
```
哎呀，func函数被我们换掉了
```
[1765] : n=0
[1765] : n=1
[1765] : n=2
[1765] : n=3
[1765] : n=4
from the hell ...
from the hell ...
from the hell ...
from the hell ...
from the hell ...
```

