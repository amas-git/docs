---
title: Linux的内存管理
tags:
---
<!-- toc -->

# 内存管理
我们来观察一下linux的init进程实际上是怎么使用虚拟地址空间的:
|| start-end || perm || offset || major:minor || inode || image ||
 * start-end: 起始VA和终止VA
 * 页权限: `rwx[s|p]`
   * rwx: 可读/可写/可执行
   * p|s: private 或 shared
后四项都是跟image文件相关的:
 * offset: 此页面所包含的内容在源文件中的偏移量(源文件即: image, 它的路径见最后一行)
 * major:minor: image所在存储设备的主设备号和从设备号(or device mappings, the major and minor numbers refer to the disk partition holding the device special file that was opened by the user, and not the device itself.)
 * inode: image文件的inode号
 * image: image的文件路径
```sh
$ sudo cat /proc/1/maps
08048000-0812d000 r-xp 00000000 08:02 11543323   /usr/lib/systemd/systemd
0812e000-0813c000 r--p 000e5000 08:02 11543323   /usr/lib/systemd/systemd
0813c000-0813d000 rw-p 000f3000 08:02 11543323   /usr/lib/systemd/systemd
09b6c000-09bd8000 rw-p 00000000 00:00 0          [heap]
b73ad000-b747a000 rw-p 00000000 00:00 0 
b747a000-b748f000 r-xp 00000000 08:02 11540729   /usr/lib/libz.so.1.2.7
b748f000-b7490000 r--p 00014000 08:02 11540729   /usr/lib/libz.so.1.2.7
b7490000-b7491000 rw-p 00015000 08:02 11540729   /usr/lib/libz.so.1.2.7
b7491000-b7495000 r-xp 00000000 08:02 11541306   /usr/lib/libattr.so.1.1.0
b7495000-b7496000 r--p 00003000 08:02 11541306   /usr/lib/libattr.so.1.1.0
b7496000-b7497000 rw-p 00004000 08:02 11541306   /usr/lib/libattr.so.1.1.0
b7497000-b7498000 rw-p 00000000 00:00 0 
b7498000-b749b000 r-xp 00000000 08:02 11536561   /usr/lib/libdl-2.17.so
b749b000-b749c000 r--p 00002000 08:02 11536561   /usr/lib/libdl-2.17.so
b749c000-b749d000 rw-p 00003000 08:02 11536561   /usr/lib/libdl-2.17.so
b749d000-b7646000 r-xp 00000000 08:02 11536514   /usr/lib/libc-2.17.so
b7646000-b7648000 r--p 001a8000 08:02 11536514   /usr/lib/libc-2.17.so
b7648000-b7649000 rw-p 001aa000 08:02 11536514   /usr/lib/libc-2.17.so
b7649000-b764c000 rw-p 00000000 00:00 0 
b764c000-b7663000 r-xp 00000000 08:02 11536555   /usr/lib/libpthread-2.17.so
b7663000-b7664000 r--p 00016000 08:02 11536555   /usr/lib/libpthread-2.17.so
b7664000-b7665000 rw-p 00017000 08:02 11536555   /usr/lib/libpthread-2.17.so
b7665000-b7667000 rw-p 00000000 00:00 0 
b7667000-b766e000 r-xp 00000000 08:02 11536517   /usr/lib/librt-2.17.so
b766e000-b766f000 r--p 00006000 08:02 11536517   /usr/lib/librt-2.17.so
b766f000-b7670000 rw-p 00007000 08:02 11536517   /usr/lib/librt-2.17.so
b7670000-b76b9000 r-xp 00000000 08:02 11542568   /usr/lib/libdbus-1.so.3.7.2
b76b9000-b76ba000 r--p 00048000 08:02 11542568   /usr/lib/libdbus-1.so.3.7.2
b76ba000-b76bb000 rw-p 00049000 08:02 11542568   /usr/lib/libdbus-1.so.3.7.2
b76bb000-b76bc000 rw-p 00000000 00:00 0 
b76bc000-b76d2000 r-xp 00000000 08:02 11542915   /usr/lib/libkmod.so.2.2.3
b76d2000-b76d3000 r--p 00015000 08:02 11542915   /usr/lib/libkmod.so.2.2.3
b76d3000-b76d4000 rw-p 00016000 08:02 11542915   /usr/lib/libkmod.so.2.2.3
b76d4000-b76d8000 r-xp 00000000 08:02 11541387   /usr/lib/libcap.so.2.22
b76d8000-b76d9000 rw-p 00003000 08:02 11541387   /usr/lib/libcap.so.2.22
b76d9000-b76e5000 r-xp 00000000 08:02 11540901   /usr/lib/libpam.so.0.83.1
b76e5000-b76e6000 r--p 0000b000 08:02 11540901   /usr/lib/libpam.so.0.83.1
b76e6000-b76e7000 rw-p 0000c000 08:02 11540901   /usr/lib/libpam.so.0.83.1
b76e7000-b76f8000 r-xp 00000000 08:02 11543299   /usr/lib/libudev.so.1.3.3
b76f8000-b76f9000 r--p 00010000 08:02 11543299   /usr/lib/libudev.so.1.3.3
b76f9000-b76fa000 rw-p 00011000 08:02 11543299   /usr/lib/libudev.so.1.3.3
b76fa000-b76fc000 r-xp 00000000 08:02 11543275   /usr/lib/libsystemd-daemon.so.0.0.10
b76fc000-b76fd000 r--p 00001000 08:02 11543275   /usr/lib/libsystemd-daemon.so.0.0.10
b76fd000-b76fe000 rw-p 00002000 08:02 11543275   /usr/lib/libsystemd-daemon.so.0.0.10
b7707000-b7712000 r-xp 00000000 08:02 11536582   /usr/lib/libnss_files-2.17.so
b7712000-b7713000 r--p 0000a000 08:02 11536582   /usr/lib/libnss_files-2.17.so
b7713000-b7714000 rw-p 0000b000 08:02 11536582   /usr/lib/libnss_files-2.17.so
b7714000-b7717000 rw-p 00000000 00:00 0 
b7717000-b7718000 r-xp 00000000 00:00 0          [vdso]
b7718000-b7738000 r-xp 00000000 08:02 11536557   /usr/lib/ld-2.17.so
b7738000-b7739000 r--p 0001f000 08:02 11536557   /usr/lib/ld-2.17.so
b7739000-b773a000 rw-p 00020000 08:02 11536557   /usr/lib/ld-2.17.so
bfb1f000-bfb40000 rw-p 00000000 00:00 0          [stack]
```
pmap工具可以辅助查看进程在
```sh
$ sudo pmap 1
1:   /sbin/init
08048000    916K r-x-- systemd
0812e000     56K r---- systemd
0813c000      4K rw--- systemd
09b6c000    432K rw---   [ anon ]
b73ad000    820K rw---   [ anon ]
b747a000     84K r-x-- libz.so.1.2.7
b748f000      4K r---- libz.so.1.2.7
b7490000      4K rw--- libz.so.1.2.7
b7491000     16K r-x-- libattr.so.1.1.0
b7495000      4K r---- libattr.so.1.1.0
b7496000      4K rw--- libattr.so.1.1.0
b7497000      4K rw---   [ anon ]
b7498000     12K r-x-- libdl-2.17.so
b749b000      4K r---- libdl-2.17.so
b749c000      4K rw--- libdl-2.17.so
b749d000   1700K r-x-- libc-2.17.so
b7646000      8K r---- libc-2.17.so
b7648000      4K rw--- libc-2.17.so
b7649000     12K rw---   [ anon ]
b764c000     92K r-x-- libpthread-2.17.so
b7663000      4K r---- libpthread-2.17.so
b7664000      4K rw--- libpthread-2.17.so
b7665000      8K rw---   [ anon ]
b7667000     28K r-x-- librt-2.17.so
b766e000      4K r---- librt-2.17.so
b766f000      4K rw--- librt-2.17.so
b7670000    292K r-x-- libdbus-1.so.3.7.2
b76b9000      4K r---- libdbus-1.so.3.7.2
b76ba000      4K rw--- libdbus-1.so.3.7.2
b76bb000      4K rw---   [ anon ]
b76bc000     88K r-x-- libkmod.so.2.2.3
b76d2000      4K r---- libkmod.so.2.2.3
b76d3000      4K rw--- libkmod.so.2.2.3
b76d4000     16K r-x-- libcap.so.2.22
b76d8000      4K rw--- libcap.so.2.22
b76d9000     48K r-x-- libpam.so.0.83.1
b76e5000      4K r---- libpam.so.0.83.1
b76e6000      4K rw--- libpam.so.0.83.1
b76e7000     68K r-x-- libudev.so.1.3.3
b76f8000      4K r---- libudev.so.1.3.3
b76f9000      4K rw--- libudev.so.1.3.3
b76fa000      8K r-x-- libsystemd-daemon.so.0.0.10
b76fc000      4K r---- libsystemd-daemon.so.0.0.10
b76fd000      4K rw--- libsystemd-daemon.so.0.0.10
b7707000     44K r-x-- libnss_files-2.17.so
b7712000      4K r---- libnss_files-2.17.so
b7713000      4K rw--- libnss_files-2.17.so
b7714000     12K rw---   [ anon ]
b7717000      4K r-x--   [ anon ]
b7718000    128K r-x-- ld-2.17.so
b7738000      4K r---- ld-2.17.so
b7739000      4K rw--- ld-2.17.so
bfb1f000    132K rw---   [ stack ]
 total     5140K
```
 1. 每个共享库都有三条描述记录:
```
b747a000     84K r-x-- libz.so.1.2.7 # 只读， 有执行权限，所以是.text段
b748f000      4K r---- libz.so.1.2.7 # 只读，应该是.data段
b7490000      4K rw--- libz.so.1.2.7 # 可读写，应该是.bss段
```
 2. 注意每段VM的大小，都是Page大小的整数倍
如果想观察更全面的信息:
```sh
$[sudo] pmap -XX 1
 Address Perm   Offset Device    Inode Size  Rss  Pss Shared_Clean Shared_Dirty Private_Clean Private_Dirty Referenced Anonymous AnonHugePages Swap KernelPageSize MMUPageSize Locked                VmFlagsMapping
08048000 r-xp 00000000  08:02 11543323  916  220  220            0            0           220             0        220         0             0    0              4           4      0    rd ex mr mw me dw  systemd
0812e000 r--p 000e5000  08:02 11543323   56    8    8            0            0             4             4          8         4             0    0              4           4      0    rd mr mw me dw ac  systemd
0813c000 rw-p 000f3000  08:02 11543323    4    4    4            0            0             0             4          4         4             0    0              4           4      0 rd wr mr mw me dw ac  systemd
09b6c000 rw-p 00000000  00:00        0  432  284  284            0            0             0           284        272       284             0    0              4           4      0    rd wr mr mw me ac  [heap]
b73ad000 rw-p 00000000  00:00        0  820  416  416            0            0             0           416        256       416             0    0              4           4      0    rd wr mr mw me ac  
b747a000 r-xp 00000000  08:02 11540729   84    0    0            0            0             0             0          0         0             0    0              4           4      0       rd ex mr mw me  libz.so.1.2.7
b748f000 r--p 00014000  08:02 11540729    4    4    4            0            0             0             4          4         4             0    0              4           4      0       rd mr mw me ac  libz.so.1.2.7
b7490000 rw-p 00015000  08:02 11540729    4    4    4            0            0             0             4          4         4             0    0              4           4      0    rd wr mr mw me ac  libz.so.1.2.7
b7491000 r-xp 00000000  08:02 11541306   16    0    0            0            0             0             0          0         0             0    0              4           4      0       rd ex mr mw me  libattr.so.1.1.0
b7495000 r--p 00003000  08:02 11541306    4    4    4            0            0             0             4          4         4             0    0              4           4      0       rd mr mw me ac  libattr.so.1.1.0
b7496000 rw-p 00004000  08:02 11541306    4    4    4            0            0             0             4          4         4             0    0              4           4      0    rd wr mr mw me ac  libattr.so.1.1.0
b7497000 rw-p 00000000  00:00        0    4    4    4            0            0             0             4          4         4             0    0              4           4      0    rd wr mr mw me ac  
b7498000 r-xp 00000000  08:02 11536561   12    0    0            0            0             0             0          0         0             0    0              4           4      0       rd ex mr mw me  libdl-2.17.so
b749b000 r--p 00002000  08:02 11536561    4    4    4            0            0             0             4          4         4             0    0              4           4      0       rd mr mw me ac  libdl-2.17.so
b749c000 rw-p 00003000  08:02 11536561    4    4    4            0            0             0             4          4         4             0    0              4           4      0    rd wr mr mw me ac  libdl-2.17.so
b749d000 r-xp 00000000  08:02 11536514 1700  220   11          220            0             0             0        220         0             0    0              4           4      0       rd ex mr mw me  libc-2.17.so
b7646000 r--p 001a8000  08:02 11536514    8    8    8            0            0             0             8          8         8             0    0              4           4      0       rd mr mw me ac  libc-2.17.so
b7648000 rw-p 001aa000  08:02 11536514    4    4    4            0            0             0             4          4         4             0    0              4           4      0    rd wr mr mw me ac  libc-2.17.so
b7649000 rw-p 00000000  00:00        0   12   12   12            0            0             0            12          8        12             0    0              4           4      0    rd wr mr mw me ac  
b764c000 r-xp 00000000  08:02 11536555   92   16    0           16            0             0             0         16         0             0    0              4           4      0       rd ex mr mw me  libpthread-2.17.so
b7663000 r--p 00016000  08:02 11536555    4    4    4            0            0             0             4          4         4             0    0              4           4      0       rd mr mw me ac  libpthread-2.17.so
b7664000 rw-p 00017000  08:02 11536555    4    4    4            0            0             0             4          4         4             0    0              4           4      0    rd wr mr mw me ac  libpthread-2.17.so
b7665000 rw-p 00000000  00:00        0    8    4    4            0            0             0             4          4         4             0    0              4           4      0    rd wr mr mw me ac  
b7667000 r-xp 00000000  08:02 11536517   28    0    0            0            0             0             0          0         0             0    0              4           4      0       rd ex mr mw me  librt-2.17.so
b766e000 r--p 00006000  08:02 11536517    4    4    4            0            0             0             4          0         4             0    0              4           4      0       rd mr mw me ac  librt-2.17.so
b766f000 rw-p 00007000  08:02 11536517    4    4    4            0            0             0             4          0         4             0    0              4           4      0    rd wr mr mw me ac  librt-2.17.so
b7670000 r-xp 00000000  08:02 11542568  292   16    4           16            0             0             0         16         0             0    0              4           4      0       rd ex mr mw me  libdbus-1.so.3.7.2
b76b9000 r--p 00048000  08:02 11542568    4    4    4            0            0             0             4          4         4             0    0              4           4      0       rd mr mw me ac  libdbus-1.so.3.7.2
b76ba000 rw-p 00049000  08:02 11542568    4    4    4            0            0             0             4          4         4             0    0              4           4      0    rd wr mr mw me ac  libdbus-1.so.3.7.2
b76bb000 rw-p 00000000  00:00        0    4    4    4            0            0             0             4          4         4             0    0              4           4      0    rd wr mr mw me ac  
b76bc000 r-xp 00000000  08:02 11542915   88    0    0            0            0             0             0          0         0             0    0              4           4      0       rd ex mr mw me  libkmod.so.2.2.3
b76d2000 r--p 00015000  08:02 11542915    4    4    4            0            0             0             4          4         4             0    0              4           4      0       rd mr mw me ac  libkmod.so.2.2.3
b76d3000 rw-p 00016000  08:02 11542915    4    4    4            0            0             0             4          4         4             0    0              4           4      0    rd wr mr mw me ac  libkmod.so.2.2.3
b76d4000 r-xp 00000000  08:02 11541387   16    0    0            0            0             0             0          0         0             0    0              4           4      0       rd ex mr mw me  libcap.so.2.22
b76d8000 rw-p 00003000  08:02 11541387    4    4    4            0            0             0             4          4         4             0    0              4           4      0    rd wr mr mw me ac  libcap.so.2.22
b76d9000 r-xp 00000000  08:02 11540901   48    0    0            0            0             0             0          0         0             0    0              4           4      0       rd ex mr mw me  libpam.so.0.83.1
b76e5000 r--p 0000b000  08:02 11540901    4    4    4            0            0             0             4          0         4             0    0              4           4      0       rd mr mw me ac  libpam.so.0.83.1
b76e6000 rw-p 0000c000  08:02 11540901    4    4    4            0            0             0             4          4         4             0    0              4           4      0    rd wr mr mw me ac  libpam.so.0.83.1
b76e7000 r-xp 00000000  08:02 11543299   68   44   23           40            0             4             0         44         0             0    0              4           4      0       rd ex mr mw me  libudev.so.1.3.3
b76f8000 r--p 00010000  08:02 11543299    4    4    4            0            0             0             4          4         4             0    0              4           4      0       rd mr mw me ac  libudev.so.1.3.3
b76f9000 rw-p 00011000  08:02 11543299    4    4    4            0            0             0             4          4         4             0    0              4           4      0    rd wr mr mw me ac  libudev.so.1.3.3
b76fa000 r-xp 00000000  08:02 11543275    8    0    0            0            0             0             0          0         0             0    0              4           4      0       rd ex mr mw me  libsystemd-daemon.so.0.0.10
b76fc000 r--p 00001000  08:02 11543275    4    4    4            0            0             0             4          4         4             0    0              4           4      0       rd mr mw me ac  libsystemd-daemon.so.0.0.10
b76fd000 rw-p 00002000  08:02 11543275    4    4    4            0            0             0             4          4         4             0    0              4           4      0    rd wr mr mw me ac  libsystemd-daemon.so.0.0.10
b7707000 r-xp 00000000  08:02 11536582   44    0    0            0            0             0             0          0         0             0    0              4           4      0       rd ex mr mw me  libnss_files-2.17.so
b7712000 r--p 0000a000  08:02 11536582    4    4    4            0            0             0             4          4         4             0    0              4           4      0       rd mr mw me ac  libnss_files-2.17.so
b7713000 rw-p 0000b000  08:02 11536582    4    4    4            0            0             0             4          4         4             0    0              4           4      0    rd wr mr mw me ac  libnss_files-2.17.so
b7714000 rw-p 00000000  00:00        0   12   12   12            0            0             0            12          8        12             0    0              4           4      0    rd wr mr mw me ac  
b7717000 r-xp 00000000  00:00        0    4    4    0            4            0             0             0          4         0             0    0              4           4      0    rd ex mr mw me de  [vdso]
b7718000 r-xp 00000000  08:02 11536557  128    0    0            0            0             0             0          0         0             0    0              4           4      0    rd ex mr mw me dw  ld-2.17.so
b7738000 r--p 0001f000  08:02 11536557    4    4    4            0            0             0             4          4         4             0    0              4           4      0    rd mr mw me dw ac  ld-2.17.so
b7739000 rw-p 00020000  08:02 11536557    4    4    4            0            0             0             4          4         4             0    0              4           4      0 rd wr mr mw me dw ac  ld-2.17.so
bfb1f000 rw-p 00000000  00:00        0  136   24   24            0            0             0            24         20        24             0    0              4           4      0 rd wr mr mw me gd ac  [stack]
                                       ==== ==== ==== ============ ============ ============= ============= ========== ========= ============= ==== ============== =========== ====== 
                                       5144 1404 1142          296            0           228           880       1208       880             0    0            212         212      0 KB 
```
 * [stack] : 栈区起始位置
 * [vdso] : 虚拟共享对象
 * [heap] : 堆区
从上面这些信息中我们大致可以看出程序对VMA的使用情况:
|| 0xC0000000 ||= 内核地址空间 =|| 内核空间与用户空间的地址边界由TASK_SIZE定义
||  ||             || 随机偏移
||  ||= ↓ 栈区 =|| Stack区大小由RLIMIT_STACK确定，通常为8M
||  || ||  random of mmap offset
||  ||= ↓ Memory Mapping Segment  =|| 动态库，anonymous 映射
|| || ||xxx
||  ||= ↑ 堆区 =|| 
||  ||  || random brk offset ||
|| ||= BSS Segment  =|| `static char * message;`
||  ||= DATA Segement =|| `static char * message="hello";`
||  ||= TEXT Segment =|| 指令
# size
可以通过size命令查看ELF文件主要段的大小。
||= hello.c ||= TEXT=||= DATA=||= BSS=||= DEC=||= HEX=||= FILE=||
```td
```cpp
#include <stdio.h>
int main() {
  return 0;
}
```
```
||1035    || 276||  4||  1315|| 523|| a.out ||
|------------------------------------------
```td
```cpp
#include <stdio.h>
int main() {
    static sM = 1;
    return 0;
}
```
```
|| 1035  || (`+4`)`280`|| 4||  1319||  527|| a.out||
|------------------------------------------
```td
```cpp
#include <stdio.h>
int main() {
    static sM = 1;
    static sX;
    static sY = 0;
    return 0;
}
```
```
|| 1035  || 280|| (`+8`)`12`||  1327||  52f|| a.out||
|------------------------------------------
```td
```cpp
#include <stdio.h>
int gX;
int main() {
    static sM = 1;
    static sX;
    static sY = 0;
    return 0;
}
```
```
|| 1035  || 280|| (`+4`)`16`||  1331||  533|| a.out||
|------------------------------------------
```td
```cpp
#include <stdio.h>
int gX;
int gM = 1;
int main() {
    static sM = 1;
    static sX;
    static sY = 0;
    return 0;
}
```
```
|| 1035  || (`+4`)`284`|| 16||  1335||  537|| a.out||
|------------------------------------------
```td
```cpp
#include <stdio.h>
int gM = 1;
int gX;
int gY = 0;
int main() {
    static sM = 1;
    static sX;
    static sY = 0;
    return 0;
}
```
```
|| 1035  || 284|| (`+4`)`20`||  1339||  53b|| a.out||
 * 已初始化的`静态变量`和`全局变量`(初始值不为0)被编译到`.data`段
 * 未初始化的`静态变量`和`全局变量`(初始值不为0)被编译到`.bss`段
进入main函数后，栈的状态如下:
|| null ||
|| 环境变量 ||
|| null ||
|| 参数 ||
|| argc ||
----

# HelloWorld
```cpp
#include <stdio.h>
int * calcArgc(char **argv);
int * calcArgc(char **argv) {
    int *p = (int)argv;
    return p - 1;
}
void printArgc(char **argv) {
    int *argc = calcArgc(argv);
    printf("addr=%p argc=%d  
", argc, *argc);
}
void printArgv(char **argv) {
    int argc = *(calcArgc(argv));
    int i;
    for(i=0; i<argc; ++i) {
        printf("addr=%p value='%s'
", argv+i, *(argv+i));
    }
}
void dump(char ** p) {
    if(p == NULL) {
        return;
    }
    while (*p != NULL) {
        printf("addr=%p value='%s'
", p, *p);
        p++;
    }
}
void printEnv(char **argv) {
    int * argc = calcArgc(argv);
    char ** env = argc + *argc + 2;
    dump(env);
}
int main(int argc, char **argv) {
    printArgc(argv);
    printArgv(argv);
    printEnv(argv);
    return 0;
}
```
```sh
$ gcc helloworld.c
$ export TEST="HelloWorld"
$ ./a.out a b c d e
addr=0xbfc83ac0 argc=6  
#-------------------------------------[ argv ]
addr=0xbfc83ac4 value='./a.out'
addr=0xbfc83ac8 value='a'
addr=0xbfc83acc value='b'
addr=0xbfc83ad0 value='c'
addr=0xbfc83ad4 value='d'
addr=0xbfc83ad8 value='e'
#-------------------------------------[ env ]
addr=0xbfc83ae0 value='STY=666.pts-0.lambda'
addr=0xbfc83b98 value='WINDOWID=23068678'
...
addr=0xbfc83b9c value='COLORFGBG=default;default;0'
addr=0xbfc83ba0 value='TERMINFO=/usr/share/terminfo'
addr=0xbfc83ba4 value='COLORTERM=rxvt-xpm'
addr=0xbfc83ba8 value='OLDPWD=/data/src/c/so'
addr=0xbfc83bac value='vcs_info_msg_0_='
addr=0xbfc83bb0 value='vcs_info_msg_1_='
addr=0xbfc83bb4 value='TEST=HelloWorld'
```
用gdb来仔细观察一下
```sh
$ gcc -g helloworld.c
$ gdb a.out a b c d e
(gdb) b main
(gdb) run
(gdb) frame
#0  main (argc=1, argv=0xbfffe8a4) at helloworld.c:43
43          printArgc(argv);
(gdb) info registers
eax            0x1      1
ecx            0xbfffe8a4       -1073747804
edx            0x8048517        134513943
ebx            0xb7fbf000       -1208225792
esp            0xbfffe7f0       0xbfffe7f0
ebp            0xbfffe808       0xbfffe808
esi            0x0      0
edi            0x0      0
eip            0x8048520        0x8048520 <main+9>
eflags         0x286    [ PF SF IF ]
cs             0x73     115
ss             0x7b     123
ds             0x7b     123
es             0x7b     123
fs             0x0      0
gs             0x33     51
(gdb) p &argc
$1 = (int *) 0xbfffe810
(gdb) p &argv
$2 = (char ***) 0xbfffe814
```
