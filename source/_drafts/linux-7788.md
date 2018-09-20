---
title: Linuc7788
tags:
---
# Linux Admin
[[TOC]]
# Anti Exploitation Techniques
# AAAS (ASCII Armored Address Space)
AAAS is a very interesting idea. The idea is to load libraries (and more 
generally any ET_DYN object) in the 16 first megabytes of the address 
space. As a result, all code and data of these shared libraries are located
at addresses beginning with a NULL byte. It naturally breaks the 
exploitation of the particular set of overflow bugs in which an improper 
use of the NULL byte character leads to the corruption (for example 
strcpy() functions and similar situations). Such a protection is 
intrinsically not effective against situations where the NULL byte is not 
an issue or when the return address used by the attacker does not contain a
NULL byte (like the PLT on Linux/*BSD x86 systems). Such a protection is 
used on Fedora distributions.
# ESP (Executable Space Protection)   
The idea of this protection mechanism is very old and simple. 
Traditionally, overflows are exploited using shellcodes which means the 
execution of user supplied 'code' in a 'data' area. Such an unusual 
situation is easily mitigated by preventing data sections (stack, heap, 
.data, etc.) and more generally (if possible) all writable process memory 
from executing. This cannot however prevent the attacker from calling 
already loaded code such as libraries or program functions. This led to the
classical return-into-libc family of attacks. Nowadays all PAE or 64 bits 
x86 linux kernel are supporting this by default.
# ASLR (Address Space Layout Randomization)
The idea of ASLR is to randomize the loading address of several memory 
areas such as the program's stack and heap, or its libraries. As a result 
even if the attacker overwrites the metadata and is able to change the 
program flow, he doesn't know where the next instructions (shellcode, 
library functions) are. The idea is simple and effective. ASLR is enabled 
by default on linux kernel since linux 2.6.12.
# Stack canaries (canaries of the death)
This is a compiler mechanism, in contrast to previously kernel-based 
described techniques. When a function is called, the code inserted by the 
compiler in its prologue stores a special value (the so-called cookie) on 
the stack before the metadata. This value is a kind of defender of 
sensitive data. During the epilogue the stack value is compared with the 
original one and if they are not the same then a memory corruption must 
have occurred. The program is then killed and this situation is reported in
the system logs. Details about technical implementation and little arm race
between protection and bypassing protection in this area will be explained 
further.
# 参考
 * http://www.phrack.org/issues.html?issue=67&id=13#article
[[TOC]]
# Linux Command Cook Book =
# 全文替换 ==
```
#!sh
# 使用perl
$ perl -p -i -e 's/foo/bar/g' filename
# 去掉android源码中的@hide API
$ find . -type f -name '*.java' -print0 | xargs -0 perl -p -i -e 's/@hide/___amas___/g'
```
# 添加用户到某个组 ==
```
#!sh
$ gpasswd -a user group
```
 * network
 * power
 * wheel
# 将光盘制作成ISO文件 ==
```
#!sh
# 如果光盘已经挂载，请务必先卸载之
$ dd if=/dev/dvd of /path/to/dvd.iso
```
# 如何确认某个组已经存在 ==
```
#!sh
$ groups groupname
```
# 查看当前网络使用详情 (端口 进程) ==
```
$ udo netstat -tlunp
```
# 打印进程树: pstree ==
```
#!sh
$ pstree
init─┬─6*[agetty]
     ├─chromium───6*[chromium───{chromium}]
     ├─crond
     ├─2*[dbus-daemon]
     ├─dbus-launch
     ├─dhcpcd
     ├─fcitx───{fcitx}
     ├─hald─┬─hald-runner─┬─hald-addon-acpi
     │      │             ├─hald-addon-inpu
     │      │             └─hald-addon-stor
     │      └─{hald}
     ├─httpd───11*[httpd]
     ├─slim─┬─X
     │      └─xmonad-i386-lin─┬─chromium─┬─chromium
     │                        │          └─25*[{chromium}]
     │                        ├─emacs
     │                        ├─urxvt───screen───screen─┬─zsh───sudo───pacman───wget
     │                        │                         └─zsh───pstree
     │                        ├─wicd-client
     │                        └─xfce4-panel─┬─xfce4-clipman-p
     │                                      ├─xfce4-cpugraph-
     │                                      ├─xfce4-dict-plug
     │                                      ├─xfce4-fsguard-p
     │                                      ├─xfce4-mixer-plu───{xfce4-mixer-pl}
     │                                      ├─xfce4-netload-p
     │                                      ├─xfce4-notes-plu
     │                                      ├─xfce4-screensho
     │                                      ├─xfce4-systemloa
     │                                      ├─xfce4-timer
     │                                      └─xfce4-weather-p
     ├─sshd
     ├─syslog-ng───syslog-ng
     ├─udevd───2*[udevd]
     ├─wicd─┬─dhcpcd
     │      └─wicd-monitor
     └─xfconfd
```
# 列出所有属于用户amas的进程PID: pgrep ==
```
#!sh
$ pgrep -u amas
# 列出属于root进程的sshd的PID
$ pgrep -u root sshd
```
# 给输出加上行号: nl ==
```
#!sh
$ cat ~/.profile | nl
     1  export JAVA_HOME=/opt/java
     2  export CLASSPATH=.:$JAVA_HOME/lib
     3  export ANDROID_HOME=/opt/android
     4  export TEXLIVE_HOME=/usr/local/texlive/2009
     5  PATH=$JAVA_HOME/bin:$ANDROID_HOME/tools:$TEXLIVE_HOME/bin/i386-linux:$PATH
# cat -n 也能为输出加上行号
$ cat -n ~/.profile
```
# 寻找依赖的动态库: ldd ==
```
#!sh
$ ldd /bin/ls
        linux-gate.so.1 =>  (0xb76fb000)
        librt.so.1 => /lib/librt.so.1 (0xb76d7000)
        libcap.so.2 => /lib/libcap.so.2 (0xb76d3000)
        libacl.so.1 => /lib/libacl.so.1 (0xb76cc000)
        libc.so.6 => /lib/libc.so.6 (0xb7581000)
        libpthread.so.0 => /lib/libpthread.so.0 (0xb7566000)
        libattr.so.1 => /lib/libattr.so.1 (0xb7561000)
        /lib/ld-linux.so.2 (0xb76fc000)
```
# 查看打开了哪些文件: lsof ==
```
#!sh
$ lsof
# 查看emacs打开了哪些文件
$ lsof | grep emacs
```
# 系统信息 ==
# 处理器架构: arch ===
```
#!sh
$ arch
i686
```
# 当前用户都属于哪些组: groups
```
#!sh
$ groups
http network video audio floppy storage users vboxusers
```
# 主机名: hostname
```
#!sh
$ hostname
myhost
```
# 负载: uptime 
# 还有谁登录了这台主机: who
# 取文件的前N行: head 
||= -n[N] =||前N行||
||= -c[N] =||前N个字符||
```
#!sh
$ ifconfig | head -n1 
eth0      Link encap:Ethernet  HWaddr 00:22:15:87:2B:08
$ ifconfig | head -c4
etho
```
# watch 
# sort
# xargs
```
#!sh
$ mkdir dir{1..4} ;  ls
dir1  dir2  dir3  dir4
$ ls | xargs
dir1 dir2 dir3 dir4
$ ls | xargs -0
dir1
dir2
dir3
dir4
```
# 全文替换
```
#!/bin/sh
#
# greplace
#
# Globally replace one string with another in a set of files
#
/usr/bin/perl -p -i -e "s/$1/$2/g" $3
```
mysql中携带的replace工具也非常好用
```
#!sh
$ replace from to -- filename
```
# lsof
```
#!sh
$ lsof /usr/bin/evince    #谁在使用evince编辑某个文件
$ lsof -u amas            #列出某个用户打开的所有文件
$ lsof -p 1234            #列出有个pid对应的程序打开的文件
$ lsof -i tcp:80          #列出某个使用tcp端口的进程
```
# 递归删除BOM
```sh
$ find . -type f -exec sed -i -e '1s/^﻿//' {} ;
$ find . -type f -print0 | xargs -0r awk '/^﻿/ {print FILENAME} {nextfile}'
```
# ls打印数值形式的权限(octal)
```sh
$ ls -l | sed -e 's/--x/1/g' -e 's/-w-/2/g' -e 's/-wx/3/g' -e 's/r--/4/g' -e 's/r-x/5/g' -e 's/rw-/6/g' -e 's/rwx/7/g' -e '
```
# 找出10个最常使用的命令
```sh
$ history | awk '{print $2}' | sort | uniq -c | sort -rn | head
$ history | cut -c8- | sort | uniq -c | sort -rn | head
```
# 打印0到9
```sh
$ seq 0 9
0
1
2
3
4
5
6
7
8
9
# 格式化方式
$ seq -f "你好: %f" 0 9
你好: 0.000000
你好: 1.000000
你好: 2.000000
你好: 3.000000
你好: 4.000000
你好: 5.000000
你好: 6.000000
你好: 7.000000
你好: 8.000000
你好: 9.000000
```
# 监听8000端口，将收到的请求打印到终端
```sh
# 1. 
$ ncat -l 8000
# 2. 打开浏览器，输入http://localhost:8000
# 3. 观察输出
$ ncat -l 8000
GET / HTTP/1.1
Host: localhost:8000
Connection: keep-alive
User-Agent: Mozilla/5.0 (X11; Linux i686) AppleWebKit/534.24 (KHTML, like Gecko) Chrome/11.0.696.71 Safari/534.24
Accept: application/xml,application/xhtml+xml,text/html;q=0.9,text/plain;q=0.8,image/png,*/*;q=0.5
Accept-Encoding: gzip,deflate,sdch
Accept-Language: en-US,en;q=0.8
Accept-Charset: ISO-8859-1,utf-8;q=0.7,*;q=0.3
Cookie: android_developer_guide_lastpage=/android-docs/guide/practices/screens_support.html; android_developer_android-docs_width=243px; android_developer_reference_lastpage=/android-docs/reference/android/app/Activity.html
```
# 生成一个矩阵
```sh
$ echo 'graph{node[shape=record];rankdir=LR;matrix[label="{1|2|3}|{4|5|6}|{7|8|9}",color=red]}' | dot -Tpng | display
```
结果是这样的:
```graphviz
graph{node[shape=record];rankdir=LR;matrix[label="{1|2|3}|{4|5|6}|{7|8|9}",color=red]}
```
# 使用指定应用打开某个文件
参看: xdg-open
```sh
$ alias o='xdg-open "$@" 2>/dev/null'
$ o x.png
```
# 使用vim读取STDIN
```sh
$ ls | vim +'set bt=nowrite' -
```
# MySQL 相关
```sh
# 查看当前数据库连接数
$ mysql -u root -p -e"show processlist;"|awk '{print $3}'|awk -F":" '{print $1}'|sort|uniq -c
```
# 查看系统所有用户名
```sh
$ awk -F ':' '{print $1 | "sort";}' /etc/passwd
# 或
$ cut -d: -f1 /etc/passwd | sort
```
# 格式化xml/html
```sh
# 从网络获取html
$ curl --silent http://exsample.com/ | xmllint --html --format - | more
# 本地
$ xmllint --html --format
```
# ls全路径
```sh
$ ls -d1 $PWD/{.*,*}
$ ls -d $PWD/*
$ printf "$PWD/%s
" *
# 列出隐藏文件的全路径
$ printf "$PWD/%s
" .*
```
# 查看哪些进程正在监听80端口
```sh
$ [sudo] fuser -v 80/tcp
```
# 查看zip文件的详细信息
```sh
$ zipinfo x.zip
```
# 使用iptables打开指定端口
```sh
$ sudo iptables -I INPUT -p tcp --dport 3000 -j ACCEPT
```
# 使用ssh在远程主机上执行命令
```sh
$ ssh user@host 'ls -al ~/'
# 打包远程服务器(hudson)上的家目录到本地
$ ssh hudson 'tar zcf - ~' > local.tgz
# 或者
$ ssh hudson 'tar cf - /path/to/dir' > local.tgz
```
# 通过ssh查看远程主机上的log
```sh
$ ssh -t remotebox "tail -f /var/log/remote.log"
```
# 查看某个进程的限制项
```sh
$ cat /proc/$PID/limits
Limit                     Soft Limit           Hard Limit           Units     
Max cpu time              unlimited            unlimited            seconds   
Max file size             unlimited            unlimited            bytes     
Max data size             unlimited            unlimited            bytes     
Max stack size            8388608              unlimited            bytes     
Max core file size        0                    unlimited            bytes     
Max resident set          unlimited            unlimited            bytes     
Max processes             25929                25929                processes 
Max open files            1024                 1024                 files     
Max locked memory         65536                65536                bytes     
Max address space         unlimited            unlimited            bytes     
Max file locks            unlimited            unlimited            locks     
Max pending signals       25929                25929                signals   
Max msgqueue size         819200               819200               bytes     
Max nice priority         0                    0                    
Max realtime priority     0                    0                    
Max realtime timeout      unlimited            unlimited            us   
```
# 16进制转10进制
```sh
$ echo $[0xFF]
255
```
# 打印第N行到第M行之间的内容
```sh
$ sed -n '$N,${M}p' /path/to/file
# 或者
$ awk 'NR >= ${N} && NR <= ${M}' /path/to/file
```
# 将字符串翻译为16进制
```sh
$ hex() { printf "%X
" $1; }
```
# 在终端上创建一个dialog
```sh
$ dialog --msgbox "hello world" 100 100 
```
# 删除尺寸小于500x500的图片
```sh
$ dentify -format '%w %h %f
' *.jpg | awk 'NF==3&&$1<500&&$2<500{print $3}' | xargs -r rm
```
# pv
 * 研究下
# ddate
```sh
$ ddate
Today is Prickle-Prickle, the 14th day of Chaos in the YOLD 3178
```
# 查看网卡信息
See: ethtool
```
$ sudo ethtool -i eth0
driver: tg3
version: 3.116
firmware-version: sb
bus-info: 0000:02:00.
```
# 绘制内核模块的依赖关系
```sh
$ lsmod | perl -e 'print "digraph "lsmod" {";<>;while(<>){@_=split/s+/; print ""$_[0]" -> "$_"
" for split/,/,$_[3]}print "}"' | dot -Tpng | display -
```
# cowsay
```sh
$ cowsay hello
 _______ 
< hello >
 ------- 
           ^__^
           (oo)_______
            (__)       )/\
                ||----w |
                ||     ||
# cowsay 电子表
$ watch -tn1 'figlet -f slant `date +%T` | cowsay -n -f telebears'
```
# shuf
随机置换文件中的行
```sh
$ cat adb.log
--- adb starting (pid 1554) ---
adb server killed by remote request
--- adb starting (pid 1883) ---
$ shuf adb.log
--- adb starting (pid 1883) ---
adb server killed by remote request
--- adb starting (pid 1554) ---
$ shuf adb.log 
adb server killed by remote request
--- adb starting (pid 1554) ---
--- adb starting (pid 1883) ---
```
# 踢掉用户(不好使)
```sh
$ killall -u username
```
# 创建随机密码
```sh
$ read -s pass; echo $pass | md5sum | base64 | cut -c -16
$ cat /dev/urandom | tr -dc A-Za-z0-9 | head -c 32
```
# 抓选中窗口
```sh
$ scrot -s /tmp/file.png
# 用鼠标点击窗口
```
# 建立PDF文件的预览图
```sh
$ evince-thumbnailer --size 600 x.pdf x.png
# 或者
$ convert -resize 600  'lexyacc.pdf[0]' x.png  
```
# 将man手册保存为PDF
```sh
$ man -t awk | ps2pdf - awk.pdf
```
# 查询500M ~ 1G之间的文件
```sh
$ find / -type f -size +500M -size -1G
```
# 加密文件
```sh
$ gpg -c <filename>
```
# 打印keycode
```sh
$ showkey -a
```
# Linux Command =
# 替换 ==
```sh
#!/bin/sh
#
# greplace
#
# Globally replace one string with another in a set of files
#
/usr/bin/perl -p -i -e "s/$1/$2/g" $3
```
mysql中携带的replace工具也非常好用
```
#!sh
$ replace from to -- filename
```
[[TOC]]
# Linux File Permissions 
# 概述
 * 在GnuLinux中，每个用户都有自己的Linux帐号， 这个帐号属于一个或多个[LinuxUserGroup 用户组]
 * 每个文件也属于一个Linux帐号和一个[LinuxUserGroup 用户组]
 * 每个文件可以定义三种基本权限:
   * 读
   * 写
   * 执行
# 查看文件的属性 
你可以使用`ls -l`命令观察以上提到的这些文件属性.
```
#!div class=note
因为`ls -l`使用频率非常之高，所以你可以将它定义一个别名，放在你所使用的Shell配置文件中, 比如: .bashrc
```
alias ll='ls -l'
```
```
现在我们来观察以下ls命令的输出:
```
#!sh
$ ls -l 
drwxr-xr-x 3 amas users 4096 Jun 30 10:56 etc
-rwxr-xr-x 1 amas users  144 Apr 19 14:11 loop.sh
```
关注一下第一栏信息:
|| 符号 || 文件类型 ||
||-	       || 普通文件 ||
||d	       || 目录 ||
||l	       || [SymbolicLinc 符号链接] ||
||s	       || Socket ||
||p        || [NamedPipe 命名管道] ||
||c        || [CharacterDevice 字符设备] ||
||b        || [BlockedDevice 块设备]   ||
 
# 修改文件的权限
可以使用chmod命令修改文件属性，在此之前，让我们来介绍一下
```
        chmod [ugoa][+-] permission file
        chmod NNN file
```
permission:
|| u ||	用户帐号 ||
|| g ||	用户组 ||
|| o ||	其他用户 ||
|| a ||	所有权限组 ||
|| + || 添加指定的权限 ||
|| - || 去掉指定的权限 ||
来看几个例子:
```
#!sh
# 将文件file设置为其他用户可读/可写/可执行的
$ chmod o+rwx file.txt
# 将build.sh设为本用户不可执行的
$ chmod u-x build.sh
```
# 文件权限的数字表示法
||= r =||= w =||= x =||= - =||
||  4  ||  2  ||  1  || 0   ||
||= --- =|| 0+0+0=0 ||= 0 =|| 啥都不能   ||
||= --x =|| 0+0+1=1 ||= 1 =|| 只能执行   ||
||= -w- =|| 0+2+0=2 ||= 2 =|| 只能写     ||
||= -wx =|| 0+2+1=3 ||= 3 =|| 可写可执行  ||
||= r-- =|| 4+0+0=4 ||= 4 =|| 只读      ||
||= r-x =|| 4+0+1=5 ||= 5 =|| 可读可执行 ||
||= rw- =|| 4+2+0=6 ||= 6 =|| 可读可写   ||
||= rwx =|| 4+2+1=7 ||= 7 =|| 都行      ||
 
||||||=用户=||||||=组=||||||=其他=||
||||||=u=||||||=g||||||=o=||
||=r=||=w=||=x=||=r=||=w=||=x=||=r=||=w=||=x=||
几个例子:
```
        755 = -rwxr-xr-x
        660 = -rw-rw----
```
# 修改文件的所有权
```
        chown user:group file
```
# 文件的StickyBit
> 如果你需要一个任何用户都可以读写的目录，怎么能保证用户之间不能修改彼此的文件?
我们来观察一下`/tmp`目录的属性:
```
#!sh
$ ls -l / | grep tmp
drwxrwxrwt  31 root root  4096 Aug  3 16:48 tmp
```
注意，除了我们之前提到的文件属性， 这个`/tmp`目录最后一位是个't', 这个't'取代了原来的x权限， 它就是StickyBit, 有了它
用户可以在这个目录下自由的创建文件，但是彼此却不能修改.
你可以使用chmod来增加StickyBit:
```
#!sh
$ chmod +t shared-dir
$ chmod -t shared-dir
```
# 文件的SGID属性
默认情况下，用户所执行的程序只具有该用户，以及该用户默认组所拥有的权限， SGID可以使进程拥有目标文件拥有者的权限，而不是可执行文件本身的权限.
来看个例子:
```
$ ls amas-server
-rwxrwxrwx  10 amas root  4096 2006-03-10 12:50 logger.sh
-rwxrwx---  10 amas root  4096 2006-03-10 12:50 server.log
```
这里logger.sh是一个可执行文件， 它负责收集信息到server.log中， 我们希望:
 1. 任何登录用户都可以执行/修改logger.sh
 2. server.log只有管理员权限才能查看
如果按照以上的权限设置，用户x(非root组)执行logger.sh的时候，会遭遇无法写入server.log的错误， 因为当x执行logger.sh的时候， logger.sh对应的进程只拥有x的权限，它虽然可以执行logger.sh, 但是
这个进程却无法写入server.log文件。
这样的问题就需要使用SGID， 使x的logger.sh进程拥有logger.sh所属者(即:amas)的权限, 从而可以写入server.log
```
#!sh
$ chmod u+s logger.sh
```
同样，你也可以给组属性添加SGID:
```
#!sh
$ chmod g+s logger.sh
```
这样，用户x所执行的logger.sh进程将拥有root用户的权限.
# umask 
你在某个目录下，新建文件的权限属性由umask来决定:
你可以使用umask命令查看但前目录下的umask值
```
#!sh
$ umask 
022
```
如果你在该目录下建立了一个文件(或目录)，其默认权限属性可以通过一下公式计算:
```
        permission-of-dir  = 777 - umask
        permission-of-file = 666 - umask
```
本例: 777-022=755 即: `-rwxr-xr-x`
如果你想改变目录的umask值，可以使用
```
#!sh
$ umask 077
```
# Linux FileSystem
# 参考 
 * http://www.cyberciti.biz/files/linux-kernel/Documentation/filesystems/
[[TOC]]
# 内核参数调优
#  /proc/sys/vm/drop_caches
 1.  释放PageCache
 2.  释放Dentries和inodes:
 3.  释放1+2
```
$ echo 3 >  /proc/sys/vm/drop_caches
```
# /proc/<pid>/io
# rchar
-----
I/O counter: chars read
The number of bytes which this task has caused to be read from storage. This
is simply the sum of bytes which this process passed to read() and pread().
It includes things like tty IO and it is unaffected by whether or not actual
physical disk IO was required (the read might have been satisfied from
pagecache)
# wchar
-----
I/O counter: chars written
The number of bytes which this task has caused, or shall cause to be written
to disk. Similar caveats apply here as with rchar.
# read_bytes
----------
I/O counter: bytes read
Attempt to count the number of bytes which this process really did cause to
be fetched from the storage layer. Done at the submit_bio() level, so it is
accurate for block-backed filesystems. <please add status regarding NFS and
CIFS at a later time>
# write_bytes
-----------
I/O counter: bytes written
Attempt to count the number of bytes which this process caused to be sent to
the storage layer. This is done at page-dirtying time.
# /proc/loadavg : 系统负载
```sh
$ cat /proc/loadavg
0.92 0.74 0.54 2/435 20024
```
||=最近1分钟负载=||=最近5分钟内负载=||=最近5分钟内负载=||= 活动任务/等待调度的人速 =||= 最近启动的进程PID =||
|| 0.92 || 0.74 || 0.54 ||  2/435 ||  20024 ||
------------------------------------------------------------------------------
                       T H E  /proc   F I L E S Y S T E M
------------------------------------------------------------------------------
/proc/sys         Terrehon Bowden <terrehon@pacbell.net>        October 7 1999
                  Bodo Bauer <bb@ricochet.net>
2.4.x update	  Jorge Nerin <comandante@zaralinux.com>      November 14 2000
move /proc/sys	  Shen Feng <shen@cn.fujitsu.com>		  April 1 2009
------------------------------------------------------------------------------
Version 1.3                                              Kernel version 2.2.12
					      Kernel version 2.4.0-test11-pre4
------------------------------------------------------------------------------
fixes/update part 1.1  Stefani Seibold <stefani@seibold.net>       June 9 2009
Table of Contents
-----------------
  0     Preface
  0.1	Introduction/Credits
  0.2	Legal Stuff
  1	Collecting System Information
  1.1	Process-Specific Subdirectories
  1.2	Kernel data
  1.3	IDE devices in /proc/ide
  1.4	Networking info in /proc/net
  1.5	SCSI info
  1.6	Parallel port info in /proc/parport
  1.7	TTY info in /proc/tty
  1.8	Miscellaneous kernel statistics in /proc/stat
  1.9 Ext4 file system parameters
  2	Modifying System Parameters
  3	Per-Process Parameters
  3.1	/proc/<pid>/oom_adj & /proc/<pid>/oom_score_adj - Adjust the oom-killer
								score
  3.2	/proc/<pid>/oom_score - Display current oom-killer score
  3.3	/proc/<pid>/io - Display the IO accounting fields
  3.4	/proc/<pid>/coredump_filter - Core dump filtering settings
  3.5	/proc/<pid>/mountinfo - Information about mounts
  3.6	/proc/<pid>/comm  & /proc/<pid>/task/<tid>/comm
------------------------------------------------------------------------------
Preface
------------------------------------------------------------------------------
0.1 Introduction/Credits
------------------------
This documentation is  part of a soon (or  so we hope) to be  released book on
the SuSE  Linux distribution. As  there is  no complete documentation  for the
/proc file system and we've used  many freely available sources to write these
chapters, it  seems only fair  to give the work  back to the  Linux community.
This work is  based on the 2.2.*  kernel version and the  upcoming 2.4.*. I'm
afraid it's still far from complete, but we  hope it will be useful. As far as
we know, it is the first 'all-in-one' document about the /proc file system. It
is focused  on the Intel  x86 hardware,  so if you  are looking for  PPC, ARM,
SPARC, AXP, etc., features, you probably  won't find what you are looking for.
It also only covers IPv4 networking, not IPv6 nor other protocols - sorry. But
additions and patches  are welcome and will  be added to this  document if you
mail them to Bodo.
We'd like  to  thank Alan Cox, Rik van Riel, and Alexey Kuznetsov and a lot of
other people for help compiling this documentation. We'd also like to extend a
special thank  you to Andi Kleen for documentation, which we relied on heavily
to create  this  document,  as well as the additional information he provided.
Thanks to  everybody  else  who contributed source or docs to the Linux kernel
and helped create a great piece of software... :)
If you  have  any comments, corrections or additions, please don't hesitate to
contact Bodo  Bauer  at  bb@ricochet.net.  We'll  be happy to add them to this
document.
The   latest   version    of   this   document   is    available   online   at
http://tldp.org/LDP/Linux-Filesystem-Hierarchy/html/proc.html
If  the above  direction does  not works  for you,  you could  try the  kernel
mailing  list  at  linux-kernel@vger.kernel.org  and/or try  to  reach  me  at
comandante@zaralinux.com.
0.2 Legal Stuff
---------------
We don't  guarantee  the  correctness  of this document, and if you come to us
complaining about  how  you  screwed  up  your  system  because  of  incorrect
documentation, we won't feel responsible...
------------------------------------------------------------------------------
CHAPTER 1: COLLECTING SYSTEM INFORMATION
------------------------------------------------------------------------------
------------------------------------------------------------------------------
In This Chapter
------------------------------------------------------------------------------
* Investigating  the  properties  of  the  pseudo  file  system  /proc and its
  ability to provide information on the running Linux system
* Examining /proc's structure
* Uncovering  various  information  about the kernel and the processes running
  on the system
------------------------------------------------------------------------------
The proc  file  system acts as an interface to internal data structures in the
kernel. It  can  be  used to obtain information about the system and to change
certain kernel parameters at runtime (sysctl).
First, we'll  take  a  look  at the read-only parts of /proc. In Chapter 2, we
show you how you can use /proc/sys to change settings.
# 1.1 /proc/<pid> : Process-Specific Subdirectories
-----------------------------------
The directory  /proc  contains  (among other things) one subdirectory for each
process running on the system, which is named after the process ID (PID).
The link  self  points  to  the  process reading the file system. Each process
subdirectory has the entries listed in Table 1-1.
Table 1-1: Process specific entries in /proc
..............................................................................
 File		Content
# clear_refs	
Clears page referenced bits shown in smaps output
# cmdline
Command line arguments
# cpu		
Current and last cpu in which it was executed	(2.4)(smp)
# cwd		
Link to the current working directory
# environ
	Values of environment variables
# exe		
Link to the executable of this process
# fd		
Directory, which contains all file descriptors
# maps		
Memory maps to executables and library files	(2.4)
# mem		
Memory held by this process
# root		
Link to the root directory of this process
# stat		
Process status
# statm		
Process memory status information
# status		
Process status in human readable form
# wchan		
If CONFIG_KALLSYMS is set, a pre-decoded wchan
# pagemap	
Page table
# stack		
Report full stack trace, enable via CONFIG_STACKTRACE
# smaps		
a extension based on maps, showing the memory consumption of each mapping
..............................................................................
For example, to get the status information of a process, all you have to do is
read the file /proc/PID/status:
  >cat /proc/self/status
  Name:   cat
  State:  R (running)
  Tgid:   5452
  Pid:    5452
  PPid:   743
  TracerPid:      0						(2.4)
  Uid:    501     501     501     501
  Gid:    100     100     100     100
  FDSize: 256
  Groups: 100 14 16
  VmPeak:     5004 kB
  VmSize:     5004 kB
  VmLck:         0 kB
  VmHWM:       476 kB
  VmRSS:       476 kB
  VmData:      156 kB
  VmStk:        88 kB
  VmExe:        68 kB
  VmLib:      1412 kB
  VmPTE:        20 kb
  VmSwap:        0 kB
  Threads:        1
  SigQ:   0/28578
  SigPnd: 0000000000000000
  ShdPnd: 0000000000000000
  SigBlk: 0000000000000000
  SigIgn: 0000000000000000
  SigCgt: 0000000000000000
  CapInh: 00000000fffffeff
  CapPrm: 0000000000000000
  CapEff: 0000000000000000
  CapBnd: ffffffffffffffff
  voluntary_ctxt_switches:        0
  nonvoluntary_ctxt_switches:     1
This shows you nearly the same information you would get if you viewed it with
the ps  command.  In  fact,  ps  uses  the  proc  file  system  to  obtain its
information.  But you get a more detailed  view of the  process by reading the
file /proc/PID/status. It fields are described in table 1-2.
The  statm  file  contains  more  detailed  information about the process
memory usage. Its seven fields are explained in Table 1-3.  The stat file
contains details information about the process itself.  Its fields are
explained in Table 1-4.
(for SMP CONFIG users)
For making accounting scalable, RSS related information are handled in
asynchronous manner and the vaule may not be very precise. To see a precise
snapshot of a moment, you can see /proc/<pid>/smaps file and scan page table.
It's slow but very precise.
Table 1-2: Contents of the status files (as of 2.6.30-rc7)
..............................................................................
 Field                       Content
 Name                        filename of the executable
 State                       state (R is running, S is sleeping, D is sleeping
                             in an uninterruptible wait, Z is zombie,
			     T is traced or stopped)
 Tgid                        thread group ID
 Pid                         process id
 PPid                        process id of the parent process
 TracerPid                   PID of process tracing this process (0 if not)
 Uid                         Real, effective, saved set, and  file system UIDs
 Gid                         Real, effective, saved set, and  file system GIDs
 FDSize                      number of file descriptor slots currently allocated
 Groups                      supplementary group list
 VmPeak                      peak virtual memory size
 VmSize                      total program size
 VmLck                       locked memory size
 VmHWM                       peak resident set size ("high water mark")
 VmRSS                       size of memory portions
 VmData                      size of data, stack, and text segments
 VmStk                       size of data, stack, and text segments
 VmExe                       size of text segment
 VmLib                       size of shared library code
 VmPTE                       size of page table entries
 VmSwap                      size of swap usage (the number of referred swapents)
 Threads                     number of threads
 SigQ                        number of signals queued/max. number for queue
 SigPnd                      bitmap of pending signals for the thread
 ShdPnd                      bitmap of shared pending signals for the process
 SigBlk                      bitmap of blocked signals
 SigIgn                      bitmap of ignored signals
 SigCgt                      bitmap of catched signals
 CapInh                      bitmap of inheritable capabilities
 CapPrm                      bitmap of permitted capabilities
 CapEff                      bitmap of effective capabilities
 CapBnd                      bitmap of capabilities bounding set
 Cpus_allowed                mask of CPUs on which this process may run
 Cpus_allowed_list           Same as previous, but in "list format"
 Mems_allowed                mask of memory nodes allowed to this process
 Mems_allowed_list           Same as previous, but in "list format"
 voluntary_ctxt_switches     number of voluntary context switches
 nonvoluntary_ctxt_switches  number of non voluntary context switches
..............................................................................
Table 1-3: Contents of the statm files (as of 2.6.8-rc3)
..............................................................................
 Field    Content
 size     total program size (pages)		(same as VmSize in status)
 resident size of memory portions (pages)	(same as VmRSS in status)
 shared   number of pages that are shared	(i.e. backed by a file)
 trs      number of pages that are 'code'	(not including libs; broken,
							includes data segment)
 lrs      number of pages of library		(always 0 on 2.6)
 drs      number of pages of data/stack		(including libs; broken,
							includes library text)
 dt       number of dirty pages			(always 0 on 2.6)
..............................................................................
Table 1-4: Contents of the stat files (as of 2.6.30-rc7)
..............................................................................
 Field          Content
  pid           process id
  tcomm         filename of the executable
  state         state (R is running, S is sleeping, D is sleeping in an
                uninterruptible wait, Z is zombie, T is traced or stopped)
  ppid          process id of the parent process
  pgrp          pgrp of the process
  sid           session id
  tty_nr        tty the process uses
  tty_pgrp      pgrp of the tty
  flags         task flags
  min_flt       number of minor faults
  cmin_flt      number of minor faults with child's
  maj_flt       number of major faults
  cmaj_flt      number of major faults with child's
  utime         user mode jiffies
  stime         kernel mode jiffies
  cutime        user mode jiffies with child's
  cstime        kernel mode jiffies with child's
  priority      priority level
  nice          nice level
  num_threads   number of threads
  it_real_value	(obsolete, always 0)
  start_time    time the process started after system boot
  vsize         virtual memory size
  rss           resident set memory size
  rsslim        current limit in bytes on the rss
  start_code    address above which program text can run
  end_code      address below which program text can run
  start_stack   address of the start of the stack
  esp           current value of ESP
  eip           current value of EIP
  pending       bitmap of pending signals
  blocked       bitmap of blocked signals
  sigign        bitmap of ignored signals
  sigcatch      bitmap of catched signals
  wchan         address where process went to sleep
  0             (place holder)
  0             (place holder)
  exit_signal   signal to send to parent thread on exit
  task_cpu      which CPU the task is scheduled on
  rt_priority   realtime priority
  policy        scheduling policy (man sched_setscheduler)
  blkio_ticks   time spent waiting for block IO
  gtime         guest time of the task in jiffies
  cgtime        guest time of the task children in jiffies
..............................................................................
The /proc/PID/maps file containing the currently mapped memory regions and
their access permissions.
The format is:
address           perms offset  dev   inode      pathname
08048000-08049000 r-xp 00000000 03:00 8312       /opt/test
08049000-0804a000 rw-p 00001000 03:00 8312       /opt/test
0804a000-0806b000 rw-p 00000000 00:00 0          [heap]
a7cb1000-a7cb2000 ---p 00000000 00:00 0
a7cb2000-a7eb2000 rw-p 00000000 00:00 0
a7eb2000-a7eb3000 ---p 00000000 00:00 0
a7eb3000-a7ed5000 rw-p 00000000 00:00 0
a7ed5000-a8008000 r-xp 00000000 03:00 4222       /lib/libc.so.6
a8008000-a800a000 r--p 00133000 03:00 4222       /lib/libc.so.6
a800a000-a800b000 rw-p 00135000 03:00 4222       /lib/libc.so.6
a800b000-a800e000 rw-p 00000000 00:00 0
a800e000-a8022000 r-xp 00000000 03:00 14462      /lib/libpthread.so.0
a8022000-a8023000 r--p 00013000 03:00 14462      /lib/libpthread.so.0
a8023000-a8024000 rw-p 00014000 03:00 14462      /lib/libpthread.so.0
a8024000-a8027000 rw-p 00000000 00:00 0
a8027000-a8043000 r-xp 00000000 03:00 8317       /lib/ld-linux.so.2
a8043000-a8044000 r--p 0001b000 03:00 8317       /lib/ld-linux.so.2
a8044000-a8045000 rw-p 0001c000 03:00 8317       /lib/ld-linux.so.2
aff35000-aff4a000 rw-p 00000000 00:00 0          [stack]
ffffe000-fffff000 r-xp 00000000 00:00 0          [vdso]
where "address" is the address space in the process that it occupies, "perms"
is a set of permissions:
 r = read
 w = write
 x = execute
 s = shared
 p = private (copy on write)
"offset" is the offset into the mapping, "dev" is the device (major:minor), and
"inode" is the inode  on that device.  0 indicates that  no inode is associated
with the memory region, as the case would be with BSS (uninitialized data).
The "pathname" shows the name associated file for this mapping.  If the mapping
is not associated with a file:
 [heap]                   = the heap of the program
 [stack]                  = the stack of the main process
 [vdso]                   = the "virtual dynamic shared object",
                            the kernel system call handler
 or if empty, the mapping is anonymous.
The /proc/PID/smaps is an extension based on maps, showing the memory
consumption for each of the process's mappings. For each of mappings there
is a series of lines such as the following:
08048000-080bc000 r-xp 00000000 03:02 13130      /bin/bash
Size:               1084 kB
Rss:                 892 kB
Pss:                 374 kB
Shared_Clean:        892 kB
Shared_Dirty:          0 kB
Private_Clean:         0 kB
Private_Dirty:         0 kB
Referenced:          892 kB
Anonymous:             0 kB
Swap:                  0 kB
KernelPageSize:        4 kB
MMUPageSize:           4 kB
The first of these lines shows the same information as is displayed for the
mapping in /proc/PID/maps.  The remaining lines show the size of the mapping
(size), the amount of the mapping that is currently resident in RAM (RSS), the
process' proportional share of this mapping (PSS), the number of clean and
dirty private pages in the mapping.  Note that even a page which is part of a
MAP_SHARED mapping, but has only a single pte mapped, i.e.  is currently used
by only one process, is accounted as private and not as shared.  "Referenced"
indicates the amount of memory currently marked as referenced or accessed.
"Anonymous" shows the amount of memory that does not belong to any file.  Even
a mapping associated with a file may contain anonymous pages: when MAP_PRIVATE
and a page is modified, the file page is replaced by a private anonymous copy.
"Swap" shows how much would-be-anonymous memory is also used, but out on
swap.
This file is only present if the CONFIG_MMU kernel configuration option is
enabled.
The /proc/PID/clear_refs is used to reset the PG_Referenced and ACCESSED/YOUNG
bits on both physical and virtual pages associated with a process.
To clear the bits for all the pages associated with the process
    > echo 1 > /proc/PID/clear_refs
To clear the bits for the anonymous pages associated with the process
    > echo 2 > /proc/PID/clear_refs
To clear the bits for the file mapped pages associated with the process
    > echo 3 > /proc/PID/clear_refs
Any other value written to /proc/PID/clear_refs will have no effect.
The /proc/pid/pagemap gives the PFN, which can be used to find the pageflags
using /proc/kpageflags and number of times a page is mapped using
/proc/kpagecount. For detailed explanation, see Documentation/vm/pagemap.txt.
1.2 Kernel data
---------------
Similar to  the  process entries, the kernel data files give information about
the running kernel. The files used to obtain this information are contained in
/proc and  are  listed  in Table 1-5. Not all of these will be present in your
system. It  depends  on the kernel configuration and the loaded modules, which
files are there, and which are missing.
Table 1-5: Kernel info in /proc
..............................................................................
 File        Content                                           
 apm         Advanced power management info                    
 buddyinfo   Kernel memory allocator information (see text)	(2.5)
 bus         Directory containing bus specific information     
 cmdline     Kernel command line                               
 cpuinfo     Info about the CPU                                
 devices     Available devices (block and character)           
 dma         Used DMS channels                                 
 filesystems Supported filesystems                             
 driver	     Various drivers grouped here, currently rtc (2.4)
 execdomains Execdomains, related to security			(2.4)
 fb	     Frame Buffer devices				(2.4)
 fs	     File system parameters, currently nfs/exports	(2.4)
 ide         Directory containing info about the IDE subsystem 
 interrupts  Interrupt usage                                   
 iomem	     Memory map						(2.4)
 ioports     I/O port usage                                    
 irq	     Masks for irq to cpu affinity			(2.4)(smp?)
 isapnp	     ISA PnP (Plug&Play) Info				(2.4)
 kcore       Kernel core image (can be ELF or A.OUT(deprecated in 2.4))   
 kmsg        Kernel messages                                   
 ksyms       Kernel symbol table                               
 loadavg     Load average of last 1, 5 & 15 minutes                
 locks       Kernel locks                                      
 meminfo     Memory info                                       
 misc        Miscellaneous                                     
 modules     List of loaded modules                            
 mounts      Mounted filesystems                               
 net         Networking info (see text)                        
 pagetypeinfo Additional page allocator information (see text)  (2.5)
 partitions  Table of partitions known to the system           
 pci	     Deprecated info of PCI bus (new way -> /proc/bus/pci/,
             decoupled by lspci					(2.4)
 rtc         Real time clock                                   
 scsi        SCSI info (see text)                              
 slabinfo    Slab pool info                                    
 softirqs    softirq usage
 stat        Overall statistics                                
 swaps       Swap space utilization                            
 sys         See chapter 2                                     
 sysvipc     Info of SysVIPC Resources (msg, sem, shm)		(2.4)
 tty	     Info of tty drivers
 uptime      System uptime                                     
 version     Kernel version                                    
 video	     bttv info of video resources			(2.4)
 vmallocinfo Show vmalloced areas
..............................................................................
You can,  for  example,  check  which interrupts are currently in use and what
they are used for by looking in the file /proc/interrupts:
  > cat /proc/interrupts 
             CPU0        
    0:    8728810          XT-PIC  timer 
    1:        895          XT-PIC  keyboard 
    2:          0          XT-PIC  cascade 
    3:     531695          XT-PIC  aha152x 
    4:    2014133          XT-PIC  serial 
    5:      44401          XT-PIC  pcnet_cs 
    8:          2          XT-PIC  rtc 
   11:          8          XT-PIC  i82365 
   12:     182918          XT-PIC  PS/2 Mouse 
   13:          1          XT-PIC  fpu 
   14:    1232265          XT-PIC  ide0 
   15:          7          XT-PIC  ide1 
  NMI:          0 
In 2.4.* a couple of lines where added to this file LOC & ERR (this time is the
output of a SMP machine):
  > cat /proc/interrupts 
             CPU0       CPU1       
    0:    1243498    1214548    IO-APIC-edge  timer
    1:       8949       8958    IO-APIC-edge  keyboard
    2:          0          0          XT-PIC  cascade
    5:      11286      10161    IO-APIC-edge  soundblaster
    8:          1          0    IO-APIC-edge  rtc
    9:      27422      27407    IO-APIC-edge  3c503
   12:     113645     113873    IO-APIC-edge  PS/2 Mouse
   13:          0          0          XT-PIC  fpu
   14:      22491      24012    IO-APIC-edge  ide0
   15:       2183       2415    IO-APIC-edge  ide1
   17:      30564      30414   IO-APIC-level  eth0
   18:        177        164   IO-APIC-level  bttv
  NMI:    2457961    2457959 
  LOC:    2457882    2457881 
  ERR:       2155
NMI is incremented in this case because every timer interrupt generates a NMI
(Non Maskable Interrupt) which is used by the NMI Watchdog to detect lockups.
LOC is the local interrupt counter of the internal APIC of every CPU.
ERR is incremented in the case of errors in the IO-APIC bus (the bus that
connects the CPUs in a SMP system. This means that an error has been detected,
the IO-APIC automatically retry the transmission, so it should not be a big
problem, but you should read the SMP-FAQ.
In 2.6.2* /proc/interrupts was expanded again.  This time the goal was for
/proc/interrupts to display every IRQ vector in use by the system, not
just those considered 'most important'.  The new vectors are:
  THR -- interrupt raised when a machine check threshold counter
  (typically counting ECC corrected errors of memory or cache) exceeds
  a configurable threshold.  Only available on some systems.
  TRM -- a thermal event interrupt occurs when a temperature threshold
  has been exceeded for the CPU.  This interrupt may also be generated
  when the temperature drops back to normal.
  SPU -- a spurious interrupt is some interrupt that was raised then lowered
  by some IO device before it could be fully processed by the APIC.  Hence
  the APIC sees the interrupt but does not know what device it came from.
  For this case the APIC will generate the interrupt with a IRQ vector
  of 0xff. This might also be generated by chipset bugs.
  RES, CAL, TLB -- rescheduling, call and TLB flush interrupts are
  sent from one CPU to another per the needs of the OS.  Typically,
  their statistics are used by kernel developers and interested users to
  determine the occurrence of interrupts of the given type.
The above IRQ vectors are displayed only when relevent.  For example,
the threshold vector does not exist on x86_64 platforms.  Others are
suppressed when the system is a uniprocessor.  As of this writing, only
i386 and x86_64 platforms support the new IRQ vector displays.
Of some interest is the introduction of the /proc/irq directory to 2.4.
It could be used to set IRQ to CPU affinity, this means that you can "hook" an
IRQ to only one CPU, or to exclude a CPU of handling IRQs. The contents of the
irq subdir is one subdir for each IRQ, and two files; default_smp_affinity and
prof_cpu_mask.
For example 
  > ls /proc/irq/
  0  10  12  14  16  18  2  4  6  8  prof_cpu_mask
  1  11  13  15  17  19  3  5  7  9  default_smp_affinity
  > ls /proc/irq/0/
  smp_affinity
smp_affinity is a bitmask, in which you can specify which CPUs can handle the
IRQ, you can set it by doing:
  > echo 1 > /proc/irq/10/smp_affinity
This means that only the first CPU will handle the IRQ, but you can also echo
5 which means that only the first and fourth CPU can handle the IRQ.
The contents of each smp_affinity file is the same by default:
  > cat /proc/irq/0/smp_affinity
  ffffffff
The default_smp_affinity mask applies to all non-active IRQs, which are the
IRQs which have not yet been allocated/activated, and hence which lack a
/proc/irq/[0-9]* directory.
The node file on an SMP system shows the node to which the device using the IRQ
reports itself as being attached. This hardware locality information does not
include information about any possible driver locality preference.
prof_cpu_mask specifies which CPUs are to be profiled by the system wide
profiler. Default value is ffffffff (all cpus).
The way IRQs are routed is handled by the IO-APIC, and it's Round Robin
between all the CPUs which are allowed to handle it. As usual the kernel has
more info than you and does a better job than you, so the defaults are the
best choice for almost everyone.
There are  three  more  important subdirectories in /proc: net, scsi, and sys.
The general  rule  is  that  the  contents,  or  even  the  existence of these
directories, depend  on your kernel configuration. If SCSI is not enabled, the
directory scsi  may  not  exist. The same is true with the net, which is there
only when networking support is present in the running kernel.
The slabinfo  file  gives  information  about  memory usage at the slab level.
Linux uses  slab  pools for memory management above page level in version 2.2.
Commonly used  objects  have  their  own  slab  pool (such as network buffers,
directory cache, and so on).
..............................................................................
> cat /proc/buddyinfo
Node 0, zone      DMA      0      4      5      4      4      3 ...
Node 0, zone   Normal      1      0      0      1    101      8 ...
Node 0, zone  HighMem      2      0      0      1      1      0 ...
External fragmentation is a problem under some workloads, and buddyinfo is a
useful tool for helping diagnose these problems.  Buddyinfo will give you a 
clue as to how big an area you can safely allocate, or why a previous
allocation failed.
Each column represents the number of pages of a certain order which are 
available.  In this case, there are 0 chunks of 2^0*PAGE_SIZE available in 
ZONE_DMA, 4 chunks of 2^1*PAGE_SIZE in ZONE_DMA, 101 chunks of 2^4*PAGE_SIZE 
available in ZONE_NORMAL, etc... 
More information relevant to external fragmentation can be found in
pagetypeinfo.
> cat /proc/pagetypeinfo
Page block order: 9
Pages per block:  512
Free pages count per migrate type at order       0      1      2      3      4      5      6      7      8      9     10
Node    0, zone      DMA, type    Unmovable      0      0      0      1      1      1      1      1      1      1      0
Node    0, zone      DMA, type  Reclaimable      0      0      0      0      0      0      0      0      0      0      0
Node    0, zone      DMA, type      Movable      1      1      2      1      2      1      1      0      1      0      2
Node    0, zone      DMA, type      Reserve      0      0      0      0      0      0      0      0      0      1      0
Node    0, zone      DMA, type      Isolate      0      0      0      0      0      0      0      0      0      0      0
Node    0, zone    DMA32, type    Unmovable    103     54     77      1      1      1     11      8      7      1      9
Node    0, zone    DMA32, type  Reclaimable      0      0      2      1      0      0      0      0      1      0      0
Node    0, zone    DMA32, type      Movable    169    152    113     91     77     54     39     13      6      1    452
Node    0, zone    DMA32, type      Reserve      1      2      2      2      2      0      1      1      1      1      0
Node    0, zone    DMA32, type      Isolate      0      0      0      0      0      0      0      0      0      0      0
Number of blocks type     Unmovable  Reclaimable      Movable      Reserve      Isolate
Node 0, zone      DMA            2            0            5            1            0
Node 0, zone    DMA32           41            6          967            2            0
Fragmentation avoidance in the kernel works by grouping pages of different
migrate types into the same contiguous regions of memory called page blocks.
A page block is typically the size of the default hugepage size e.g. 2MB on
X86-64. By keeping pages grouped based on their ability to move, the kernel
can reclaim pages within a page block to satisfy a high-order allocation.
The pagetypinfo begins with information on the size of a page block. It
then gives the same type of information as buddyinfo except broken down
by migrate-type and finishes with details on how many page blocks of each
type exist.
If min_free_kbytes has been tuned correctly (recommendations made by hugeadm
from libhugetlbfs http://sourceforge.net/projects/libhugetlbfs/), one can
make an estimate of the likely number of huge pages that can be allocated
at a given point in time. All the "Movable" blocks should be allocatable
unless memory has been mlock()'d. Some of the Reclaimable blocks should
also be allocatable although a lot of filesystem metadata may have to be
reclaimed to achieve this.
..............................................................................
meminfo:
Provides information about distribution and utilization of memory.  This
varies by architecture and compile options.  The following is from a
16GB PIII, which has highmem enabled.  You may not have all of these fields.
> cat /proc/meminfo
MemTotal:     16344972 kB
MemFree:      13634064 kB
Buffers:          3656 kB
Cached:        1195708 kB
SwapCached:          0 kB
Active:         891636 kB
Inactive:      1077224 kB
HighTotal:    15597528 kB
HighFree:     13629632 kB
LowTotal:       747444 kB
LowFree:          4432 kB
SwapTotal:           0 kB
SwapFree:            0 kB
Dirty:             968 kB
Writeback:           0 kB
AnonPages:      861800 kB
Mapped:         280372 kB
Slab:           284364 kB
SReclaimable:   159856 kB
SUnreclaim:     124508 kB
PageTables:      24448 kB
NFS_Unstable:        0 kB
Bounce:              0 kB
WritebackTmp:        0 kB
CommitLimit:   7669796 kB
Committed_AS:   100056 kB
VmallocTotal:   112216 kB
VmallocUsed:       428 kB
VmallocChunk:   111088 kB
    MemTotal: Total usable ram (i.e. physical ram minus a few reserved
              bits and the kernel binary code)
     MemFree: The sum of LowFree+HighFree
     Buffers: Relatively temporary storage for raw disk blocks
              shouldn't get tremendously large (20MB or so)
      Cached: in-memory cache for files read from the disk (the
              pagecache).  Doesn't include SwapCached
  SwapCached: Memory that once was swapped out, is swapped back in but
              still also is in the swapfile (if memory is needed it
              doesn't need to be swapped out AGAIN because it is already
              in the swapfile. This saves I/O)
      Active: Memory that has been used more recently and usually not
              reclaimed unless absolutely necessary.
    Inactive: Memory which has been less recently used.  It is more
              eligible to be reclaimed for other purposes
   HighTotal:
    HighFree: Highmem is all memory above ~860MB of physical memory
              Highmem areas are for use by userspace programs, or
              for the pagecache.  The kernel must use tricks to access
              this memory, making it slower to access than lowmem.
    LowTotal:
     LowFree: Lowmem is memory which can be used for everything that
              highmem can be used for, but it is also available for the
              kernel's use for its own data structures.  Among many
              other things, it is where everything from the Slab is
              allocated.  Bad things happen when you're out of lowmem.
   SwapTotal: total amount of swap space available
    SwapFree: Memory which has been evicted from RAM, and is temporarily
              on the disk
       Dirty: Memory which is waiting to get written back to the disk
   Writeback: Memory which is actively being written back to the disk
   AnonPages: Non-file backed pages mapped into userspace page tables
      Mapped: files which have been mmaped, such as libraries
        Slab: in-kernel data structures cache
SReclaimable: Part of Slab, that might be reclaimed, such as caches
  SUnreclaim: Part of Slab, that cannot be reclaimed on memory pressure
  PageTables: amount of memory dedicated to the lowest level of page
              tables.
NFS_Unstable: NFS pages sent to the server, but not yet committed to stable
	      storage
      Bounce: Memory used for block device "bounce buffers"
WritebackTmp: Memory used by FUSE for temporary writeback buffers
 CommitLimit: Based on the overcommit ratio ('vm.overcommit_ratio'),
              this is the total amount of  memory currently available to
              be allocated on the system. This limit is only adhered to
              if strict overcommit accounting is enabled (mode 2 in
              'vm.overcommit_memory').
              The CommitLimit is calculated with the following formula:
              CommitLimit = ('vm.overcommit_ratio' * Physical RAM) + Swap
              For example, on a system with 1G of physical RAM and 7G
              of swap with a `vm.overcommit_ratio` of 30 it would
              yield a CommitLimit of 7.3G.
              For more details, see the memory overcommit documentation
              in vm/overcommit-accounting.
Committed_AS: The amount of memory presently allocated on the system.
              The committed memory is a sum of all of the memory which
              has been allocated by processes, even if it has not been
              "used" by them as of yet. A process which malloc()'s 1G
              of memory, but only touches 300M of it will only show up
              as using 300M of memory even if it has the address space
              allocated for the entire 1G. This 1G is memory which has
              been "committed" to by the VM and can be used at any time
              by the allocating application. With strict overcommit
              enabled on the system (mode 2 in 'vm.overcommit_memory'),
              allocations which would exceed the CommitLimit (detailed
              above) will not be permitted. This is useful if one needs
              to guarantee that processes will not fail due to lack of
              memory once that memory has been successfully allocated.
VmallocTotal: total size of vmalloc memory area
 VmallocUsed: amount of vmalloc area which is used
VmallocChunk: largest contiguous block of vmalloc area which is free
..............................................................................
vmallocinfo:
Provides information about vmalloced/vmaped areas. One line per area,
containing the virtual address range of the area, size in bytes,
caller information of the creator, and optional information depending
on the kind of area :
 pages=nr    number of pages
 phys=addr   if a physical address was specified
 ioremap     I/O mapping (ioremap() and friends)
 vmalloc     vmalloc() area
 vmap        vmap()ed pages
 user        VM_USERMAP area
 vpages      buffer for pages pointers was vmalloced (huge area)
 N<node>=nr  (Only on NUMA kernels)
             Number of pages allocated on memory node <node>
> cat /proc/vmallocinfo
0xffffc20000000000-0xffffc20000201000 2101248 alloc_large_system_hash+0x204 ...
  /0x2c0 pages=512 vmalloc N0=128 N1=128 N2=128 N3=128
0xffffc20000201000-0xffffc20000302000 1052672 alloc_large_system_hash+0x204 ...
  /0x2c0 pages=256 vmalloc N0=64 N1=64 N2=64 N3=64
0xffffc20000302000-0xffffc20000304000    8192 acpi_tb_verify_table+0x21/0x4f...
  phys=7fee8000 ioremap
0xffffc20000304000-0xffffc20000307000   12288 acpi_tb_verify_table+0x21/0x4f...
  phys=7fee7000 ioremap
0xffffc2000031d000-0xffffc2000031f000    8192 init_vdso_vars+0x112/0x210
0xffffc2000031f000-0xffffc2000032b000   49152 cramfs_uncompress_init+0x2e ...
  /0x80 pages=11 vmalloc N0=3 N1=3 N2=2 N3=3
0xffffc2000033a000-0xffffc2000033d000   12288 sys_swapon+0x640/0xac0      ...
  pages=2 vmalloc N1=2
0xffffc20000347000-0xffffc2000034c000   20480 xt_alloc_table_info+0xfe ...
  /0x130 [x_tables] pages=4 vmalloc N0=4
0xffffffffa0000000-0xffffffffa000f000   61440 sys_init_module+0xc27/0x1d00 ...
   pages=14 vmalloc N2=14
0xffffffffa000f000-0xffffffffa0014000   20480 sys_init_module+0xc27/0x1d00 ...
   pages=4 vmalloc N1=4
0xffffffffa0014000-0xffffffffa0017000   12288 sys_init_module+0xc27/0x1d00 ...
   pages=2 vmalloc N1=2
0xffffffffa0017000-0xffffffffa0022000   45056 sys_init_module+0xc27/0x1d00 ...
   pages=10 vmalloc N0=10
..............................................................................
softirqs:
Provides counts of softirq handlers serviced since boot time, for each cpu.
> cat /proc/softirqs
                CPU0       CPU1       CPU2       CPU3
      HI:          0          0          0          0
   TIMER:      27166      27120      27097      27034
  NET_TX:          0          0          0         17
  NET_RX:         42          0          0         39
   BLOCK:          0          0        107       1121
 TASKLET:          0          0          0        290
   SCHED:      27035      26983      26971      26746
 HRTIMER:          0          0          0          0
     RCU:       1678       1769       2178       2250
1.3 IDE devices in /proc/ide
----------------------------
The subdirectory /proc/ide contains information about all IDE devices of which
the kernel  is  aware.  There is one subdirectory for each IDE controller, the
file drivers  and a link for each IDE device, pointing to the device directory
in the controller specific subtree.
The file  drivers  contains general information about the drivers used for the
IDE devices:
  > cat /proc/ide/drivers
  ide-cdrom version 4.53
  ide-disk version 1.08
More detailed  information  can  be  found  in  the  controller  specific
subdirectories. These  are  named  ide0,  ide1  and  so  on.  Each  of  these
directories contains the files shown in table 1-6.
Table 1-6: IDE controller info in  /proc/ide/ide?
..............................................................................
 File    Content                                 
 channel IDE channel (0 or 1)                    
 config  Configuration (only for PCI/IDE bridge) 
 mate    Mate name                               
 model   Type/Chipset of IDE controller          
..............................................................................
Each device  connected  to  a  controller  has  a separate subdirectory in the
controllers directory.  The  files  listed in table 1-7 are contained in these
directories.
Table 1-7: IDE device information
..............................................................................
 File             Content                                    
 cache            The cache                                  
 capacity         Capacity of the medium (in 512Byte blocks) 
 driver           driver and version                         
 geometry         physical and logical geometry              
 identify         device identify block                      
 media            media type                                 
 model            device identifier                          
 settings         device setup                               
 smart_thresholds IDE disk management thresholds             
 smart_values     IDE disk management values                 
..............................................................................
The most  interesting  file is settings. This file contains a nice overview of
the drive parameters:
  # cat /proc/ide/ide0/hda/settings 
  name                    value           min             max             mode 
  ----                    -----           ---             ---             ---- 
  bios_cyl                526             0               65535           rw 
  bios_head               255             0               255             rw 
  bios_sect               63              0               63              rw 
  breada_readahead        4               0               127             rw 
  bswap                   0               0               1               r 
  file_readahead          72              0               2097151         rw 
  io_32bit                0               0               3               rw 
  keepsettings            0               0               1               rw 
  max_kb_per_request      122             1               127             rw 
  multcount               0               0               8               rw 
  nice1                   1               0               1               rw 
  nowerr                  0               0               1               rw 
  pio_mode                write-only      0               255             w 
  slow                    0               0               1               rw 
  unmaskirq               0               0               1               rw 
  using_dma               0               0               1               rw 
1.4 Networking info in /proc/net
--------------------------------
The subdirectory  /proc/net  follows  the  usual  pattern. Table 1-8 shows the
additional values  you  get  for  IP  version 6 if you configure the kernel to
support this. Table 1-9 lists the files and their meaning.
Table 1-8: IPv6 info in /proc/net
..............................................................................
 File       Content                                               
 udp6       UDP sockets (IPv6)                                    
 tcp6       TCP sockets (IPv6)                                    
 raw6       Raw device statistics (IPv6)                          
 igmp6      IP multicast addresses, which this host joined (IPv6) 
 if_inet6   List of IPv6 interface addresses                      
 ipv6_route Kernel routing table for IPv6                         
 rt6_stats  Global IPv6 routing tables statistics                 
 sockstat6  Socket statistics (IPv6)                              
 snmp6      Snmp data (IPv6)                                      
..............................................................................
Table 1-9: Network info in /proc/net
..............................................................................
 File          Content                                                         
 arp           Kernel  ARP table                                               
 dev           network devices with statistics                                 
 dev_mcast     the Layer2 multicast groups a device is listening too
               (interface index, label, number of references, number of bound
               addresses). 
 dev_stat      network device status                                           
 ip_fwchains   Firewall chain linkage                                          
 ip_fwnames    Firewall chain names                                            
 ip_masq       Directory containing the masquerading tables                    
 ip_masquerade Major masquerading table                                        
 netstat       Network statistics                                              
 raw           raw device statistics                                           
 route         Kernel routing table                                            
 rpc           Directory containing rpc info                                   
 rt_cache      Routing cache                                                   
 snmp          SNMP data                                                       
 sockstat      Socket statistics                                               
 tcp           TCP  sockets                                                    
 tr_rif        Token ring RIF routing table                                    
 udp           UDP sockets                                                     
 unix          UNIX domain sockets                                             
 wireless      Wireless interface data (Wavelan etc)                           
 igmp          IP multicast addresses, which this host joined                  
 psched        Global packet scheduler parameters.                             
 netlink       List of PF_NETLINK sockets                                      
 ip_mr_vifs    List of multicast virtual interfaces                            
 ip_mr_cache   List of multicast routing cache                                 
..............................................................................
You can  use  this  information  to see which network devices are available in
your system and how much traffic was routed over those devices:
  > cat /proc/net/dev 
  Inter-|Receive                                                   |[... 
   face |bytes    packets errs drop fifo frame compressed multicast|[... 
      lo:  908188   5596     0    0    0     0          0         0 [...         
    ppp0:15475140  20721   410    0    0   410          0         0 [...  
    eth0:  614530   7085     0    0    0     0          0         1 [... 
   
  ...] Transmit 
  ...] bytes    packets errs drop fifo colls carrier compressed 
  ...]  908188     5596    0    0    0     0       0          0 
  ...] 1375103    17405    0    0    0     0       0          0 
  ...] 1703981     5535    0    0    0     3       0          0 
In addition, each Channel Bond interface has its own directory.  For
example, the bond0 device will have a directory called /proc/net/bond0/.
It will contain information that is specific to that bond, such as the
current slaves of the bond, the link status of the slaves, and how
many times the slaves link has failed.
1.5 SCSI info
-------------
If you  have  a  SCSI  host adapter in your system, you'll find a subdirectory
named after  the driver for this adapter in /proc/scsi. You'll also see a list
of all recognized SCSI devices in /proc/scsi:
  >cat /proc/scsi/scsi 
  Attached devices: 
  Host: scsi0 Channel: 00 Id: 00 Lun: 00 
    Vendor: IBM      Model: DGHS09U          Rev: 03E0 
    Type:   Direct-Access                    ANSI SCSI revision: 03 
  Host: scsi0 Channel: 00 Id: 06 Lun: 00 
    Vendor: PIONEER  Model: CD-ROM DR-U06S   Rev: 1.04 
    Type:   CD-ROM                           ANSI SCSI revision: 02 
The directory  named  after  the driver has one file for each adapter found in
the system.  These  files  contain information about the controller, including
the used  IRQ  and  the  IO  address range. The amount of information shown is
dependent on  the adapter you use. The example shows the output for an Adaptec
AHA-2940 SCSI adapter:
  > cat /proc/scsi/aic7xxx/0 
   
  Adaptec AIC7xxx driver version: 5.1.19/3.2.4 
  Compile Options: 
    TCQ Enabled By Default : Disabled 
    AIC7XXX_PROC_STATS     : Disabled 
    AIC7XXX_RESET_DELAY    : 5 
  Adapter Configuration: 
             SCSI Adapter: Adaptec AHA-294X Ultra SCSI host adapter 
                             Ultra Wide Controller 
      PCI MMAPed I/O Base: 0xeb001000 
   Adapter SEEPROM Config: SEEPROM found and used. 
        Adaptec SCSI BIOS: Enabled 
                      IRQ: 10 
                     SCBs: Active 0, Max Active 2, 
                           Allocated 15, HW 16, Page 255 
               Interrupts: 160328 
        BIOS Control Word: 0x18b6 
     Adapter Control Word: 0x005b 
     Extended Translation: Enabled 
  Disconnect Enable Flags: 0xffff 
       Ultra Enable Flags: 0x0001 
   Tag Queue Enable Flags: 0x0000 
  Ordered Queue Tag Flags: 0x0000 
  Default Tag Queue Depth: 8 
      Tagged Queue By Device array for aic7xxx host instance 0: 
        {255,255,255,255,255,255,255,255,255,255,255,255,255,255,255,255} 
      Actual queue depth per device for aic7xxx host instance 0: 
        {1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1} 
  Statistics: 
  (scsi0:0:0:0) 
    Device using Wide/Sync transfers at 40.0 MByte/sec, offset 8 
    Transinfo settings: current(12/8/1/0), goal(12/8/1/0), user(12/15/1/0) 
    Total transfers 160151 (74577 reads and 85574 writes) 
  (scsi0:0:6:0) 
    Device using Narrow/Sync transfers at 5.0 MByte/sec, offset 15 
    Transinfo settings: current(50/15/0/0), goal(50/15/0/0), user(50/15/0/0) 
    Total transfers 0 (0 reads and 0 writes) 
1.6 Parallel port info in /proc/parport
---------------------------------------
The directory  /proc/parport  contains information about the parallel ports of
your system.  It  has  one  subdirectory  for  each port, named after the port
number (0,1,2,...).
These directories contain the four files shown in Table 1-10.
Table 1-10: Files in /proc/parport
..............................................................................
 File      Content                                                             
 autoprobe Any IEEE-1284 device ID information that has been acquired.         
 devices   list of the device drivers using that port. A + will appear by the
           name of the device currently using the port (it might not appear
           against any). 
 hardware  Parallel port's base address, IRQ line and DMA channel.             
 irq       IRQ that parport is using for that port. This is in a separate
           file to allow you to alter it by writing a new value in (IRQ
           number or none). 
..............................................................................
1.7 TTY info in /proc/tty
-------------------------
Information about  the  available  and actually used tty's can be found in the
directory /proc/tty.You'll  find  entries  for drivers and line disciplines in
this directory, as shown in Table 1-11.
Table 1-11: Files in /proc/tty
..............................................................................
 File          Content                                        
 drivers       list of drivers and their usage                
 ldiscs        registered line disciplines                    
 driver/serial usage statistic and status of single tty lines 
..............................................................................
To see  which  tty's  are  currently in use, you can simply look into the file
/proc/tty/drivers:
  > cat /proc/tty/drivers 
  pty_slave            /dev/pts      136   0-255 pty:slave 
  pty_master           /dev/ptm      128   0-255 pty:master 
  pty_slave            /dev/ttyp       3   0-255 pty:slave 
  pty_master           /dev/pty        2   0-255 pty:master 
  serial               /dev/cua        5   64-67 serial:callout 
  serial               /dev/ttyS       4   64-67 serial 
  /dev/tty0            /dev/tty0       4       0 system:vtmaster 
  /dev/ptmx            /dev/ptmx       5       2 system 
  /dev/console         /dev/console    5       1 system:console 
  /dev/tty             /dev/tty        5       0 system:/dev/tty 
  unknown              /dev/tty        4    1-63 console 
1.8 Miscellaneous kernel statistics in /proc/stat
-------------------------------------------------
Various pieces   of  information about  kernel activity  are  available in the
/proc/stat file.  All  of  the numbers reported  in  this file are  aggregates
since the system first booted.  For a quick look, simply cat the file:
  > cat /proc/stat
  cpu  2255 34 2290 22625563 6290 127 456 0 0
  cpu0 1132 34 1441 11311718 3675 127 438 0 0
  cpu1 1123 0 849 11313845 2614 0 18 0 0
  intr 114930548 113199788 3 0 5 263 0 4 [... lots more numbers ...]
  ctxt 1990473
  btime 1062191376
  processes 2915
  procs_running 1
  procs_blocked 0
  softirq 183433 0 21755 12 39 1137 231 21459 2263
The very first  "cpu" line aggregates the  numbers in all  of the other "cpuN"
lines.  These numbers identify the amount of time the CPU has spent performing
different kinds of work.  Time units are in USER_HZ (typically hundredths of a
second).  The meanings of the columns are as follows, from left to right:
- user: normal processes executing in user mode
- nice: niced processes executing in user mode
- system: processes executing in kernel mode
- idle: twiddling thumbs
- iowait: waiting for I/O to complete
- irq: servicing interrupts
- softirq: servicing softirqs
- steal: involuntary wait
- guest: running a normal guest
- guest_nice: running a niced guest
The "intr" line gives counts of interrupts  serviced since boot time, for each
of the  possible system interrupts.   The first  column  is the  total of  all
interrupts serviced; each  subsequent column is the  total for that particular
interrupt.
The "ctxt" line gives the total number of context switches across all CPUs.
The "btime" line gives  the time at which the  system booted, in seconds since
the Unix epoch.
The "processes" line gives the number  of processes and threads created, which
includes (but  is not limited  to) those  created by  calls to the  fork() and
clone() system calls.
The "procs_running" line gives the total number of threads that are
running or ready to run (i.e., the total number of runnable threads).
The   "procs_blocked" line gives  the  number of  processes currently blocked,
waiting for I/O to complete.
The "softirq" line gives counts of softirqs serviced since boot time, for each
of the possible system softirqs. The first column is the total of all
softirqs serviced; each subsequent column is the total for that particular
softirq.
1.9 Ext4 file system parameters
------------------------------
Information about mounted ext4 file systems can be found in
/proc/fs/ext4.  Each mounted filesystem will have a directory in
/proc/fs/ext4 based on its device name (i.e., /proc/fs/ext4/hdc or
/proc/fs/ext4/dm-0).   The files in each per-device directory are shown
in Table 1-12, below.
Table 1-12: Files in /proc/fs/ext4/<devname>
..............................................................................
 File            Content                                        
 mb_groups       details of multiblock allocator buddy cache of free blocks
..............................................................................
------------------------------------------------------------------------------
Summary
------------------------------------------------------------------------------
The /proc file system serves information about the running system. It not only
allows access to process data but also allows you to request the kernel status
by reading files in the hierarchy.
The directory  structure  of /proc reflects the types of information and makes
it easy, if not obvious, where to look for specific data.
------------------------------------------------------------------------------
------------------------------------------------------------------------------
CHAPTER 2: MODIFYING SYSTEM PARAMETERS
------------------------------------------------------------------------------
------------------------------------------------------------------------------
In This Chapter
------------------------------------------------------------------------------
* Modifying kernel parameters by writing into files found in /proc/sys
* Exploring the files which modify certain parameters
* Review of the /proc/sys file tree
------------------------------------------------------------------------------
A very  interesting part of /proc is the directory /proc/sys. This is not only
a source  of  information,  it also allows you to change parameters within the
kernel. Be  very  careful  when attempting this. You can optimize your system,
but you  can  also  cause  it  to  crash.  Never  alter kernel parameters on a
production system.  Set  up  a  development machine and test to make sure that
everything works  the  way  you want it to. You may have no alternative but to
reboot the machine once an error has been made.
To change  a  value,  simply  echo  the new value into the file. An example is
given below  in the section on the file system data. You need to be root to do
this. You  can  create  your  own  boot script to perform this every time your
system boots.
The files  in /proc/sys can be used to fine tune and monitor miscellaneous and
general things  in  the operation of the Linux kernel. Since some of the files
can inadvertently  disrupt  your  system,  it  is  advisable  to  read  both
documentation and  source  before actually making adjustments. In any case, be
very careful  when  writing  to  any  of these files. The entries in /proc may
change slightly between the 2.1.* and the 2.2 kernel, so if there is any doubt
review the kernel documentation in the directory /usr/src/linux/Documentation.
This chapter  is  heavily  based  on the documentation included in the pre 2.2
kernels, and became part of it in version 2.2.1 of the Linux kernel.
Please see: Documentation/sysctls/ directory for descriptions of these
entries.
------------------------------------------------------------------------------
Summary
------------------------------------------------------------------------------
Certain aspects  of  kernel  behavior  can be modified at runtime, without the
need to  recompile  the kernel, or even to reboot the system. The files in the
/proc/sys tree  can  not only be read, but also modified. You can use the echo
command to write value into these files, thereby changing the default settings
of the kernel.
------------------------------------------------------------------------------
------------------------------------------------------------------------------
CHAPTER 3: PER-PROCESS PARAMETERS
------------------------------------------------------------------------------
3.1 /proc/<pid>/oom_adj & /proc/<pid>/oom_score_adj- Adjust the oom-killer score
--------------------------------------------------------------------------------
These file can be used to adjust the badness heuristic used to select which
process gets killed in out of memory conditions.
The badness heuristic assigns a value to each candidate task ranging from 0
(never kill) to 1000 (always kill) to determine which process is targeted.  The
units are roughly a proportion along that range of allowed memory the process
may allocate from based on an estimation of its current memory and swap use.
For example, if a task is using all allowed memory, its badness score will be
1000.  If it is using half of its allowed memory, its score will be 500.
There is an additional factor included in the badness score: root
processes are given 3% extra memory over other tasks.
The amount of "allowed" memory depends on the context in which the oom killer
was called.  If it is due to the memory assigned to the allocating task's cpuset
being exhausted, the allowed memory represents the set of mems assigned to that
cpuset.  If it is due to a mempolicy's node(s) being exhausted, the allowed
memory represents the set of mempolicy nodes.  If it is due to a memory
limit (or swap limit) being reached, the allowed memory is that configured
limit.  Finally, if it is due to the entire system being out of memory, the
allowed memory represents all allocatable resources.
The value of /proc/<pid>/oom_score_adj is added to the badness score before it
is used to determine which task to kill.  Acceptable values range from -1000
(OOM_SCORE_ADJ_MIN) to +1000 (OOM_SCORE_ADJ_MAX).  This allows userspace to
polarize the preference for oom killing either by always preferring a certain
task or completely disabling it.  The lowest possible value, -1000, is
equivalent to disabling oom killing entirely for that task since it will always
report a badness score of 0.
Consequently, it is very simple for userspace to define the amount of memory to
consider for each task.  Setting a /proc/<pid>/oom_score_adj value of +500, for
example, is roughly equivalent to allowing the remainder of tasks sharing the
same system, cpuset, mempolicy, or memory controller resources to use at least
50% more memory.  A value of -500, on the other hand, would be roughly
equivalent to discounting 50% of the task's allowed memory from being considered
as scoring against the task.
For backwards compatibility with previous kernels, /proc/<pid>/oom_adj may also
be used to tune the badness score.  Its acceptable values range from -16
(OOM_ADJUST_MIN) to +15 (OOM_ADJUST_MAX) and a special value of -17
(OOM_DISABLE) to disable oom killing entirely for that task.  Its value is
scaled linearly with /proc/<pid>/oom_score_adj.
Writing to /proc/<pid>/oom_score_adj or /proc/<pid>/oom_adj will change the
other with its scaled value.
NOTICE: /proc/<pid>/oom_adj is deprecated and will be removed, please see
Documentation/feature-removal-schedule.txt.
Caveat: when a parent task is selected, the oom killer will sacrifice any first
generation children with seperate address spaces instead, if possible.  This
avoids servers and important system daemons from being killed and loses the
minimal amount of work.
3.2 /proc/<pid>/oom_score - Display current oom-killer score
-------------------------------------------------------------
This file can be used to check the current score used by the oom-killer is for
any given <pid>. Use it together with /proc/<pid>/oom_adj to tune which
process should be killed in an out-of-memory situation.
3.3  /proc/<pid>/io - Display the IO accounting fields
-------------------------------------------------------
This file contains IO statistics for each running process
Example
-------
test:/tmp # dd if=/dev/zero of=/tmp/test.dat &
[1] 3828
test:/tmp # cat /proc/3828/io
rchar: 323934931
wchar: 323929600
syscr: 632687
syscw: 632675
read_bytes: 0
write_bytes: 323932160
cancelled_write_bytes: 0
Description
-----------
rchar
-----
I/O counter: chars read
The number of bytes which this task has caused to be read from storage. This
is simply the sum of bytes which this process passed to read() and pread().
It includes things like tty IO and it is unaffected by whether or not actual
physical disk IO was required (the read might have been satisfied from
pagecache)
wchar
-----
I/O counter: chars written
The number of bytes which this task has caused, or shall cause to be written
to disk. Similar caveats apply here as with rchar.
syscr
-----
I/O counter: read syscalls
Attempt to count the number of read I/O operations, i.e. syscalls like read()
and pread().
syscw
-----
I/O counter: write syscalls
Attempt to count the number of write I/O operations, i.e. syscalls like
write() and pwrite().
read_bytes
----------
I/O counter: bytes read
Attempt to count the number of bytes which this process really did cause to
be fetched from the storage layer. Done at the submit_bio() level, so it is
accurate for block-backed filesystems. <please add status regarding NFS and
CIFS at a later time>
write_bytes
-----------
I/O counter: bytes written
Attempt to count the number of bytes which this process caused to be sent to
the storage layer. This is done at page-dirtying time.
cancelled_write_bytes
---------------------
The big inaccuracy here is truncate. If a process writes 1MB to a file and
then deletes the file, it will in fact perform no writeout. But it will have
been accounted as having caused 1MB of write.
In other words: The number of bytes which this process caused to not happen,
by truncating pagecache. A task can cause "negative" IO too. If this task
truncates some dirty pagecache, some IO which another task has been accounted
for (in its write_bytes) will not be happening. We _could_ just subtract that
from the truncating task's write_bytes, but there is information loss in doing
that.
Note
----
At its current implementation state, this is a bit racy on 32-bit machines: if
process A reads process B's /proc/pid/io while process B is updating one of
those 64-bit counters, process A could see an intermediate result.
More information about this can be found within the taskstats documentation in
Documentation/accounting.
3.4 /proc/<pid>/coredump_filter - Core dump filtering settings
---------------------------------------------------------------
When a process is dumped, all anonymous memory is written to a core file as
long as the size of the core file isn't limited. But sometimes we don't want
to dump some memory segments, for example, huge shared memory. Conversely,
sometimes we want to save file-backed memory segments into a core file, not
only the individual files.
/proc/<pid>/coredump_filter allows you to customize which memory segments
will be dumped when the <pid> process is dumped. coredump_filter is a bitmask
of memory types. If a bit of the bitmask is set, memory segments of the
corresponding memory type are dumped, otherwise they are not dumped.
The following 7 memory types are supported:
  - (bit 0) anonymous private memory
  - (bit 1) anonymous shared memory
  - (bit 2) file-backed private memory
  - (bit 3) file-backed shared memory
  - (bit 4) ELF header pages in file-backed private memory areas (it is
            effective only if the bit 2 is cleared)
  - (bit 5) hugetlb private memory
  - (bit 6) hugetlb shared memory
  Note that MMIO pages such as frame buffer are never dumped and vDSO pages
  are always dumped regardless of the bitmask status.
  Note bit 0-4 doesn't effect any hugetlb memory. hugetlb memory are only
  effected by bit 5-6.
Default value of coredump_filter is 0x23; this means all anonymous memory
segments and hugetlb private memory are dumped.
If you don't want to dump all shared memory segments attached to pid 1234,
write 0x21 to the process's proc file.
  $ echo 0x21 > /proc/1234/coredump_filter
When a new process is created, the process inherits the bitmask status from its
parent. It is useful to set up coredump_filter before the program runs.
For example:
  $ echo 0x7 > /proc/self/coredump_filter
  $ ./some_program
3.5	/proc/<pid>/mountinfo - Information about mounts
--------------------------------------------------------
This file contains lines of the form:
36 35 98:0 /mnt1 /mnt2 rw,noatime master:1 - ext3 /dev/root rw,errors=continue
(1)(2)(3)   (4)   (5)      (6)      (7)   (8) (9)   (10)         (11)
(1) mount ID:  unique identifier of the mount (may be reused after umount)
(2) parent ID:  ID of parent (or of self for the top of the mount tree)
(3) major:minor:  value of st_dev for files on filesystem
(4) root:  root of the mount within the filesystem
(5) mount point:  mount point relative to the process's root
(6) mount options:  per mount options
(7) optional fields:  zero or more fields of the form "tag[:value]"
(8) separator:  marks the end of the optional fields
(9) filesystem type:  name of filesystem of the form "type[.subtype]"
(10) mount source:  filesystem specific information or "none"
(11) super options:  per super block options
Parsers should ignore all unrecognised optional fields.  Currently the
possible optional fields are:
shared:X  mount is shared in peer group X
master:X  mount is slave to peer group X
propagate_from:X  mount is slave and receives propagation from peer group X (*)
unbindable  mount is unbindable
(*) X is the closest dominant peer group under the process's root.  If
X is the immediate master of the mount, or if there's no dominant peer
group under the same root, then only the "master:X" field is present
and not the "propagate_from:X" field.
For more information on mount propagation see:
  Documentation/filesystems/sharedsubtree.txt
3.6	/proc/<pid>/comm  & /proc/<pid>/task/<tid>/comm
--------------------------------------------------------
These files provide a method to access a tasks comm value. It also allows for
a task to set its own or one of its thread siblings comm value. The comm value
is limited in size compared to the cmdline value, so writing anything longer
then the kernel's TASK_COMM_LEN (currently 16 chars) will result in a truncated
comm value.
/proc/stat explained
Various pieces of information about kernel activity are available in the
/proc/stat file.
All of the numbers reported in this file are aggregates since the system first booted.
For a quick look, simply cat the file:
```sh
$ cat /proc/stat
cpu  2255 34 2290  6290 127 456
cpu0 1132 34 1441 11311718 3675 127 438
cpu1 1123 0 849 11313845 2614 0 18
intr 114930548 113199788 3 0 5 263 0 4 [... lots more numbers ...]
ctxt 1990473
btime 1062191376
processes 2915
procs_running 1
procs_blocked 0
```
The very first "cpu" line aggregates the numbers in all of the other "cpuN" lines.
These numbers identify the amount of time the CPU has spent performing different kinds of work. Time units are in USER_HZ or Jiffies (typically hundredths of a second).
The meanings of the columns are as follows, from left to right:
|| || user || nice || system || idle || iowait || irq || softirq ||
 * user: normal processes executing in user mode
 * nice: niced processes executing in user mode
 * system: processes executing in kernel mode
 * idle: twiddling thumbs
 * iowait: waiting for I/O to complete
 * irq: servicing interrupts
 * softirq: servicing softirqs
The "intr" line gives counts of interrupts serviced since boot time, for each
of the possible system interrupts. The first column is the total of all interrupts serviced; each subsequent column is the total for that particular interrupt.
The "ctxt" line gives the total number of context switches across all CPUs.
The "btime" line gives the time at which the system booted, in seconds since
the Unix epoch.
The "processes" line gives the number of processes and threads created, which includes (but is not limited to) those created by calls to the fork() and clone() system calls.
The "procs_running" line gives the number of processes currently running on CPUs.
The "procs_blocked" line gives the number of processes currently blocked, waiting for I/O to complete.
copied from the kernel documentation of the /proc filesystem
Note: On my 2.6.18 kernel, cpu lines have 8 numeric fields, not 7. 
Wonder what that one means...
Note:
The 8th column is called steal_time. It counts the ticks spent
executing other virtual hosts (in virtualised environments like Xen)
Note2:
With Linux 2.6.24 there is 9th column for (virtual) guest systems. See man 5 proc.
# vmstat
vmstat可以查看核心运行状态， 其本质是对
 * `/proc/stat`,
 * `/proc/meminfo`
 * `/proc/vmstat`内核信息进行监控.
# HZ
HZ为系统时钟的频率. 学名叫Tick Rate, 那么Tick又是什么呢？
# Tick
```
tick = 1 / HZ 
```
Linux核心每隔固定周期会发出timer interrupt (IRQ 0)，HZ是用来定义每一秒有几次timer interrupts。举例来说，HZ为1000，代表每秒有1000次timer interrupts。 HZ可在编译核心时设定，其中HZ可设定100、250、300或1000。以核心版本预设值为250，做实验：
       观察/proc/interrupt的timer中断次数，并于一秒后再次观察其值。理论上，两者应该相差250左右。
[[TOC]]
# Linux File Timestamp
A common mistake is that ctime is the file creation time. This is not correct, it is the inode/file change time. mtime is the file modification time. A often heard question is "What is the ctime, mtime and atime?".This is confusing so let me explain the difference between ctime, mtime and atime.
# ctime
ctime是inode发生变化的时间. 通常当你修改文件属性的时候, 比如: 修改权限, 所有者, 或者将文件移动到其他文件系统都会修改ctime.
```sh
$ touch x    ; ls -l
-rw-r--r-- 1 amas users 0 Jul 18 23:46 x
$ chmod +x x ; ls -lc
-rwxr-xr-x 1 amas users 0 Jul 18 23:56 x
$ ls -l x
-rwxr-xr-x 1 amas users 0 Jul 18 23:46
$ echo x  > x
$ ls -l
-rwxr-xr-x 1 amas users 2 Jul 18 23:58 x
$ ls -lc
-rwxr-xr-x 1 amas users 2 Jul 18 23:58 x
```
# mtime
mtime是文件内容最后被修改的时间. 多说文件的ctime和mtime都是相同的, 如果两者不同, 则说明最近修改了该文件的属性.
# atime
文件被访问的时间. 只要文件被打开, 这个时间就会更新.
----
# 参考
 * [wiki:Command/stat]
[[TOC]]
# 参考
 * http://fluxius.handgrep.se/
# Debug
# DUMP IO
```sh
# 需要root权限
$ echo 1 > /proc/sys/vm/block_dump 
$ cat /proc/kmsg
```
How to use epoll? A complete example in C
Thursday, 2 June 2011 @ 1238 GMT by Mukund Sivaraman
Network servers are traditionally implemented using a separate process or thread per connection. For high performance applications that need to handle a very large number of clients simultaneously, this approach won't work well, because factors such as resource usage and context-switching time influence the ability to handle many clients at a time. An alternate method is to perform non-blocking I/O in a single thread, along with some readiness notification method which tells you when you can read or write more data on a socket.
This article is an introduction to Linux's epoll(7) facility, which is the best readiness notification facility in Linux. We will write sample code for a complete TCP server implementation in C. I assume you have C programming experience, know how to compile and run programs on Linux, and can read manpages of the various C functions that are used.
epoll was introduced in Linux 2.6, and is not available in other UNIX-like operating systems. It provides a facility similar to the select(2) and poll(2) functions:
select(2) can monitor up to FD_SETSIZE number of descriptors at a time, typically a small number determined at libc's compile time.
poll(2) doesn't have a fixed limit of descriptors it can monitor at a time, but apart from other things, even we have to perform a linear scan of all the passed descriptors every time to check readiness notification, which is O(n) and slow.
epoll has no such fixed limits, and does not perform any linear scans. Hence it is able to perform better and handle a larger number of events.
An epoll instance is created by epoll_create(2) or epoll_create1(2) (they take different arguments), which return an epoll instance. epoll_ctl(2) is used to add/remove descriptors to be watched on the epoll instance. To wait for events on the watched set, epoll_wait(2) is used, which blocks until events are available. Please see their manpages for more info.
When descriptors are added to an epoll instance, they can be added in two modes: level triggered and edge triggered. When you use level triggered mode, and data is available for reading, epoll_wait(2) will always return with ready events. If you don't read the data completely, and call epoll_wait(2) on the epoll instance watching the descriptor again, it will return again with a ready event because data is available. In edge triggered mode, you will only get a readiness notfication once. If you don't read the data fully, and call epoll_wait(2) on the epoll instance watching the descriptor again, it will block because the readiness event was already delivered.
The epoll event structure that you pass to epoll_ctl(2) is shown below. With every descriptor being watched, you can associate an integer or a pointer as user data.
typedef union epoll_data
{
  void        *ptr;
  int          fd;
  __uint32_t   u32;
  __uint64_t   u64;
} epoll_data_t;
struct epoll_event
{
  __uint32_t   events; /* Epoll events */
  epoll_data_t data;   /* User data variable */
};
Let's write code now. We'll implement a tiny TCP server that prints everything sent to the socket on standard output. We'll begin by writing a function create_and_bind() which creates and binds a TCP socket:
static int
create_and_bind (char *port)
{
  struct addrinfo hints;
  struct addrinfo *result, *rp;
  int s, sfd;
  memset (&hints, 0, sizeof (struct addrinfo));
  hints.ai_family = AF_UNSPEC;     /* Return IPv4 and IPv6 choices */
  hints.ai_socktype = SOCK_STREAM; /* We want a TCP socket */
  hints.ai_flags = AI_PASSIVE;     /* All interfaces */
  s = getaddrinfo (NULL, port, &hints, &result);
  if (s != 0)
    {
      fprintf (stderr, "getaddrinfo: %s
", gai_strerror (s));
      return -1;
    }
  for (rp = result; rp != NULL; rp = rp->ai_next)
    {
      sfd = socket (rp->ai_family, rp->ai_socktype, rp->ai_protocol);
      if (sfd == -1)
        continue;
      s = bind (sfd, rp->ai_addr, rp->ai_addrlen);
      if (s == 0)
        {
          /* We managed to bind successfully! */
          break;
        }
      close (sfd);
    }
  if (rp == NULL)
    {
      fprintf (stderr, "Could not bind
");
      return -1;
    }
  freeaddrinfo (result);
  return sfd;
}
create_and_bind() contains a standard code block for a portable way of getting a IPv4 or IPv6 socket. It accepts a port argument as a string, where argv[1] can be passed. The getaddrinfo(3) function returns a bunch of addrinfo structures in result, which are compatible with the hints passed in the hints argument. The addrinfo struct looks like this:
struct addrinfo
{
  int              ai_flags;
  int              ai_family;
  int              ai_socktype;
  int              ai_protocol;
  size_t           ai_addrlen;
  struct sockaddr *ai_addr;
  char            *ai_canonname;
  struct addrinfo *ai_next;
};
We walk through the structures one by one and try creating sockets using them, until we are able to both create and bind a socket. If we were successful, create_and_bind() returns the socket descriptor. If unsuccessful, it returns -1.
Next, let's write a function to make a socket non-blocking. make_socket_non_blocking() sets the O_NONBLOCK flag on the descriptor passed in the sfd argument:
static int
make_socket_non_blocking (int sfd)
{
  int flags, s;
  flags = fcntl (sfd, F_GETFL, 0);
  if (flags == -1)
    {
      perror ("fcntl");
      return -1;
    }
  flags |= O_NONBLOCK;
  s = fcntl (sfd, F_SETFL, flags);
  if (s == -1)
    {
      perror ("fcntl");
      return -1;
    }
  return 0;
}
Now, on to the main() function of the program which contains the event loop. This is the bulk of the program:
#define MAXEVENTS 64
int
main (int argc, char *argv[])
{
  int sfd, s;
  int efd;
  struct epoll_event event;
  struct epoll_event *events;
  if (argc != 2)
    {
      fprintf (stderr, "Usage: %s [port]
", argv[0]);
      exit (EXIT_FAILURE);
    }
  sfd = create_and_bind (argv[1]);
  if (sfd == -1)
    abort ();
  s = make_socket_non_blocking (sfd);
  if (s == -1)
    abort ();
  s = listen (sfd, SOMAXCONN);
  if (s == -1)
    {
      perror ("listen");
      abort ();
    }
  efd = epoll_create1 (0);
  if (efd == -1)
    {
      perror ("epoll_create");
      abort ();
    }
  event.data.fd = sfd;
  event.events = EPOLLIN | EPOLLET;
  s = epoll_ctl (efd, EPOLL_CTL_ADD, sfd, &event);
  if (s == -1)
    {
      perror ("epoll_ctl");
      abort ();
    }
  /* Buffer where events are returned */
  events = calloc (MAXEVENTS, sizeof event);
  /* The event loop */
  while (1)
    {
      int n, i;
      n = epoll_wait (efd, events, MAXEVENTS, -1);
      for (i = 0; i < n; i++)
	{
	  if ((events[i].events & EPOLLERR) ||
              (events[i].events & EPOLLHUP) ||
              (!(events[i].events & EPOLLIN)))
	    {
              /* An error has occured on this fd, or the socket is not
                 ready for reading (why were we notified then?) */
	      fprintf (stderr, "epoll error
");
	      close (events[i].data.fd);
	      continue;
	    }
	  else if (sfd == events[i].data.fd)
	    {
              /* We have a notification on the listening socket, which
                 means one or more incoming connections. */
              while (1)
                {
                  struct sockaddr in_addr;
                  socklen_t in_len;
                  int infd;
                  char hbuf[NI_MAXHOST], sbuf[NI_MAXSERV];
                  in_len = sizeof in_addr;
                  infd = accept (sfd, &in_addr, &in_len);
                  if (infd == -1)
                    {
                      if ((errno == EAGAIN) ||
                          (errno == EWOULDBLOCK))
                        {
                          /* We have processed all incoming
                             connections. */
                          break;
                        }
                      else
                        {
                          perror ("accept");
                          break;
                        }
                    }
                  s = getnameinfo (&in_addr, in_len,
                                   hbuf, sizeof hbuf,
                                   sbuf, sizeof sbuf,
                                   NI_NUMERICHOST | NI_NUMERICSERV);
                  if (s == 0)
                    {
                      printf("Accepted connection on descriptor %d "
                             "(host=%s, port=%s)
", infd, hbuf, sbuf);
                    }
                  /* Make the incoming socket non-blocking and add it to the
                     list of fds to monitor. */
                  s = make_socket_non_blocking (infd);
                  if (s == -1)
                    abort ();
                  event.data.fd = infd;
                  event.events = EPOLLIN | EPOLLET;
                  s = epoll_ctl (efd, EPOLL_CTL_ADD, infd, &event);
                  if (s == -1)
                    {
                      perror ("epoll_ctl");
                      abort ();
                    }
                }
              continue;
            }
          else
            {
              /* We have data on the fd waiting to be read. Read and
                 display it. We must read whatever data is available
                 completely, as we are running in edge-triggered mode
                 and won't get a notification again for the same
                 data. */
              int done = 0;
              while (1)
                {
                  ssize_t count;
                  char buf[512];
                  count = read (events[i].data.fd, buf, sizeof buf);
                  if (count == -1)
                    {
                      /* If errno == EAGAIN, that means we have read all
                         data. So go back to the main loop. */
                      if (errno != EAGAIN)
                        {
                          perror ("read");
                          done = 1;
                        }
                      break;
                    }
                  else if (count == 0)
                    {
                      /* End of file. The remote has closed the
                         connection. */
                      done = 1;
                      break;
                    }
                  /* Write the buffer to standard output */
                  s = write (1, buf, count);
                  if (s == -1)
                    {
                      perror ("write");
                      abort ();
                    }
                }
              if (done)
                {
                  printf ("Closed connection on descriptor %d
",
                          events[i].data.fd);
                  /* Closing the descriptor will make epoll remove it
                     from the set of descriptors which are monitored. */
                  close (events[i].data.fd);
                }
            }
        }
    }
  free (events);
  close (sfd);
  return EXIT_SUCCESS;
}
main() first calls create_and_bind() which sets up the socket. It then makes the socket non-blocking, and then calls listen(2). It then creates an epoll instance in efd, to which it adds the listening socket sfd to watch for input events in an edge-triggered mode.
The outer while loop is the main events loop. It calls epoll_wait(2), where the thread remains blocked waiting for events. When events are available, epoll_wait(2) returns the events in the events argument, which is a bunch of epoll_event structures.
The epoll instance in efd is continuously updated in the event loop when we add new incoming connections to watch, and remove existing connections when they die.
When events are available, they can be of three types:
Errors: When an error condition occurs, or the event is not a notification about data available for reading, we simply close the associated descriptor. Closing the descriptor automatically removes it from the watched set of epoll instance efd.
New connections: When the listening descriptor sfd is ready for reading, it means one or more new connections have arrived. While there are new connections, accept(2) the connections, print a message about it, make the incoming socket non-blocking and add it to the watched set of epoll instance efd.
Client data: When data is available for reading on any of the client descriptors, we use read(2) to read the data in pieces of 512 bytes in an inner while loop. This is because we have to read all the data that is available now, as we won't get further events about it as the descriptor is watched in edge-triggered mode. The data which is read is written to stdout (fd=1) using write(2). If read(2) returns 0, it means an EOF and we can close the client's connection. If -1 is returned, and errno is set to EAGAIN, it means that all data for this event was read, and we can go back to the main loop.
That's that. It goes around and around in a loop, adding and removing descriptors in the watched set.
Download the epoll-example.c program.
Update1: Level and edge triggered definitions were erroneously reversed (though the code was correct). It was noticed by Reddit user bodski. The article has been corrected now. I should have proof-read it before posting. Apologies, and thank you for pointing out the mistake. :)
Update2: The code has been modified to run accept(2) until it says it would block, so that if multiple connections have arrived, we accept all of them. It was noticed by Reddit user cpitchford. Thank you for the comments. :)
[[TOC]]
# IO Scheduler
# 查看IO调度器
```sh
$  cat /sys/block/<DEV>/queue/scheduler
```
# IO 调度器
CFQ: This is the default algorithm in most Linux distributions. It attempts to distribute all I/O bandwidth evenly among all processes requesting I/O. It is ideal for most purposes.
NOOP: The noop algorithm attempts to use as little cpu as possible. It acts as a basic FIFO queue expecting the hardware controller to handle the performance operations of the requests.
Anticipatory: This algorithm attempts to reorder all disk I/O operations to optimize disk seeks. It is designed to increase performance on systems that have slow disks.
Deadline: This scheduling algorithm places I/O requests in a priority queue so each is guaranteed to be ran within a certain time. It is often used in real-time operating systems.
Since the 2.6.10 kernel you were given the option to easily adjust the disk scheduler on the fly without a reboot. You can see which scheduler you are currently running by typing:
```sh
$ cat /sys/block/sda/queue/scheduler
noop anticipatory deadline [cfq]
```
 * 当前正在使用`cfq`调度器
Changing schedulers on the fly allows you to test and benchmark the algorithms for your specific application .Once the change is issued, any current I/O operations will be executed before the new scheduler goes into effect, so the change will not be instantaneous. Also remember that once one is set and performs to your liking, be sure to set the change to be applied on subsequent reboots.
There are cases where cfq may not be the best scheduler for your system or application. An example is if you are running a raid disk array with a caching raid controller. In this instance many reads and writes will be cached by the disk controller to be executed based on its own scheduling algorithm. Scheduling algorithms such as noop may offer you better performance by not wasting host os cpu cycles reordering I/O operations that do not offer the disk subsystem any benefit. Solid state drives are another case where any algorithm that reorders disk operations may not offer any benefit since random access times are identical to sequential access times. It is often recommend to use noop or deadline on any SSD drive.
There is usually no definitive answer to which algorithm to use, so benchmarking each one will be your best option. Performance results will depend greatly on the applications that are running, the type and speed of hard drives being used, the disk controllers or raid cards / level in use, along with the overall system load and amount of disk I/O. Please also keep in mind that each of these algorithms has many additional configurable options to further tweak performance to your needs depending on how much time you wish to spend on them.
[[TOC]]
# Memory
 * Vss = virtual set size
 * Rss = resident set size
 * Pss = proportional set size
 * Uss = unique set size
一般Vss和Rss用处不大，
# Uss
一个进程拥有的唯一页集合，通常这个值也代表，如果你Kill掉这个进程所立刻能够回收的内存量
# Pss 
一个进程与其他进程共享的内存大小， 如果立刻Kill这个进程，Pss未必能够得到回收。
```div class=note
# Android系统: dump meminfo <pid>
```
Applications Memory Usage (kB):
Uptime: 16667503 Realtime: 22342640
** MEMINFO in pid 5342 [com.android.fileexplorer:remote] **
                         Shared  Private     Heap     Heap     Heap
                   Pss    Dirty    Dirty     Size    Alloc     Free
                ------   ------   ------   ------   ------   ------
       Native       20       20       20     2156     1767      144
       Dalvik     3065    12176     2712     9332     8918      414
       Cursor        0        0        0                           
       Ashmem        0        0        0                           
    Other dev        4       40        0                           
     .so mmap     1317     2696     1132                           
    .jar mmap        0        0        0                           
    .apk mmap       17        0        0                           
    .ttf mmap        0        0        0                           
    .dex mmap        1        0        0                           
   Other mmap     1653      372       80                           
      Unknown      391      328      388                           
        TOTAL     6468    15632     4332    11488    10685      558
 
 Objects
               Views:        0         ViewRootImpl:        0
         AppContexts:        3           Activities:        0
              Assets:        4        AssetManagers:        4
       Local Binders:        5        Proxy Binders:       14
    Death Recipients:        0
     OpenSSL Sockets:        0
 
 SQL
         MEMORY_USED:      184
  PAGECACHE_OVERFLOW:       32          MALLOC_SIZE:       62
 
 DATABASES
      pgsz     dbsz   Lookaside(b)          cache  Dbname
         4       24             63         1/16/2  /data/data/com.android.fileexplorer/databases/kssuser.db
         4       20             30         1/16/2  /data/data/com.android.fileexplorer/databases/kuaipan_trans.db
```
 * Shared dirty : 与其他进程共享的脏页大小
 * Private dirty : 该进程占有的脏页大小
```
# Linux Networks =
# 历史 ==
早期计算机是真的高科技，造价高，技术复杂，只有少数几个公司控制。后来吧，有人琢磨能不能把机器之间联系起来，这样就像打电话一样，两台机器之间也可以互通数据。 彼时掌握这种高科技的公司，自己设计硬件软件，自己设计网络，你想想这是一个啥局面？就好比各国之间文字不通，根本无法交流。起初大家各有各的做法，彼此也决不妥协，渐渐意识到封闭不划算。我们需要有个共同的标准，标准真的是很重要的东西呀，所以我们要感谢秦始皇，感谢维护标准的专业组织，是他们在混沌中建立了秩序，而建立秩序是件难事。
# 以太网 ===
以太网是常见的网络硬件，其它诸如光纤网络，蓝牙无线技术， ATM等硬件。 这么多硬件存在主要因为适用场合不通。 比如一般家用网络不要求有太快的速度，你要用光纤就成本太高了。
总之整个网络，有很多不同的硬件接口，以太网因为种种原因，它最流行，最常见。 简单说，施乐公司为了自家的硬件设备链接发明了以太网，大力推动成为业界标准后，硬件厂商扑上来大量制造，成本降低。再加上IBM个人计算机流行，凑在一块以太网硬件接口流行程度可想而知。
最早以太网标准802.3的IEEE 10BASE5, 这个标准的定义是:
 * 10: 代表传输速率 10Mbps
 * BASE: 表示利用基频信号进行传输
 * 5: 指的是网络节点之间最长距离是5km
后来IEEE 802.3u 100BASE-T, 我们现在用的CAT5网线可以支持这个100Mbps, 你看，这网线就好比道路，要想在上面跑的快，路面必须规整，土路肯定比柏油路难走，但是制作工艺和成本都低。高速网络对网线的要求是和速度成正比的。
路不平干扰你走路，线不好干扰你上网。 网线的好不好，就看他抗干扰能力强不强。线类可以分两种
 1. UTP: 无遮蔽双绞线， 听起来就很简陋:(
 2. STP: 屏蔽双绞线， 这个不错。
数据如何在以太网上传送？
现在你碰到的公司网络啊，家庭网啊，本质上都是个Star链接，中心是个Hub或Switch. 无论怎么样，在任意时间只能有1台机器正在使用网线/Hub。这就好比是条单车道。两个计算机想同时发言，这就乱套了。 为此以太网规定了个法律，叫CSMA/CD, 这里边规定，谁要想使用以太网设备传输信息，先要检测传输介质上有没有其它节点(所谓节点就是任何一个具有MAC的网络介质)正在使用。8过，在共享介质的环境下，网络忙活的时候，还是可能冲突的。
这个MAC是不能跨Rouder地。
8个线蕊，我们之用了1,2,3,6 四个.
||= 接头名称 =|| 1    || 2  || 3    || 4  || 5     || 6  || 7     || 8 ||
|| 568A       || 白绿 || 绿 || 白橙 || 蓝 || 白蓝 || 橙 || 白棕 || 棕 ||
|| 568B       || 白绿 || 橙 || 白橙 || 蓝 || 白蓝 || 绿 || 白棕 || 棕 ||
 * 并行线: 两头都是568A， 用于链接网卡和集线器
 * 跳线: 一头是568A另一头是568B， 用于网卡之间
# 软件 ==
# 监控软件 ===
 * nagios
 * cacti
 * zabbix
# 其它 ==
 * 如果需要让Linux实现宽带共享，用Router或者NAT即可满足需求
 * 如果需要了解每个用户经常使用的网站， 那么最好使用Proxy+分析软件
 * 一般我们通过tarball安装的程序，想设置为开机启动，一般都会利用/etc/rc.d/rc.local(/etc/rc.local).
 * 一定要注意密码的设置，不能过于简单，通过猜测密码获得root权限登录是很常见的一种入侵手段
  1. 避免使用简单密码
  2. 设置密码更换周期
  3. 利用pam模块来额外地进行密码验证工作
 * 利用 SuperDaemon 和 TCPWrappers管理服务权限
  通过xinetd和 SuperDaemon 或直接调用 TCPWrappers 函数库， 那么我们可以直接使用/etc/hosts.allow和/etc/hosts.deny来管理是否能登录系统的某个daemon的权限
 * 利用netfilter防火墙， LinuxKernel 2.4.xx版本以上的防火墙机制为iptables, 2.2.xx为ipchains, 如无特殊原因，我们推荐使用iptables，配置简单，功能强大
 * 主机的服务越单一越好，为啥呢？因为万一出了问题，方便分析日志，找出原因
# 历史 
# 以太网
# 以太网网线接头
# MAC地址
Media Access Control
 * MAC帧结构
||=前导码=||=目的地址=||=来源地址=||=数据域位通信=||=主要信息=||=校验码=||
|| 8      || 6        || 6        || 2            || 46~1500  || 4      ||
 
 * CSMA/CD
 * Hardware Address: 00:00:00:00:00:00~FF:FF:FF:FF:FF:FF
 * 我们常用MAC代之网卡卡号,即HadwareAddress      
 * MAC仅在局域网内有效,不能夸网段, 夸网段时,路由会修改MAC
 * 针对MAC地址的限制,只在局域网中有效
 * MAC帧能够最大能装1500Byte的数据, 所以大于1500Byte数据在以太网中传输时,
都会被拆分/编号/分组 发送, 到了Gigabit年代,Gigabit Ethernet可以支持更大的MAC帧(JumboFrame),通常都是9000 Bytes, 
 * 更少拆分,意味着传输效能改善,可以作为选择网卡的参考项.
# IP 与 MAC
我们知道在局域网里,数据是以帧的形式在网卡之间相互传递的,
TCP/IP通讯协议只要了解IP即可,但是TCP/IP数据必须通过MAC来传递,
因此IP与MAC之间,需要一个相互解析的功能.
当主机A想要找目的IP时,就会对整个网络广播, 目标机通过广播里面的IP发现主机A是想和自己通讯,那么目标机就会向主机A返回相关的MAC信息. 其他主机发现跟自己没关系,也就不做相应. 这样, 主机A就知道目标机的MAC/IP了, 当然,每次都这么广播实在麻烦, 我们将目标机的MAC缓存到ARP table中.
```
#!sh
$ arp -nv
Address                  HWtype  HWaddress           Flags Mask            Iface
192.168.1.100            ether   d8:a2:5e:96:8e:37   C                     eth0
192.168.1.1              ether   00:25:86:4c:1d:94   C                     eth0
Entries: 2      Skipped: 0      Found: 2
# 我们发现arp中有两条记录
# 这个表由OS维护,你也可以使用arp -s 进行强制绑定    
```
# IP的组成 
IP是一种数据包格式, 最大可达65535Bytes, 
其中表示32bit表示IP(IPv4):
```
00000000.00000000.00000000.00000000 : 0.0.0.0
11111111.11111111.11111111.11111111 : 255.255.255.255
```
# 网段
在32bit中,包行两部分信息:
||=Net ID=||=Host ID=||
比如:192.168.0.0~192.168.0.225这个网段为例子:
||=Net ID =||=Host ID=||
||192.168.0||x        ||
 IP网段:: 在同一个物理网段内,主机的IP具有相同Net ID, 且具有不同的Host ID,           这些IP即构成同一个IP网段. 我们一般说网段,即IP网段之意.
 物理网段:: 所有主机使用同一网络媒介链接在一起,即所有主机都在同一物理网段,同一物理网段又可分为多个IP网段.
在同一网段中, Net ID必然相同, Host ID不能重复, 此外:
 * Host ID 在二进制表示法中不能全为0或全为1, 比如: 192.168.1网段中, 192.168.1.0和192.168.1.255不能用作主机IP.
   1. 全为0的就表示该网段的网络
   2. 全为1的IP为该网段的广播地址
 * 同一网段中的主机之间可通过MAC帧的格式传递信息, 具体方法是:
   1. 主机通过ARP协议与广播数据包取得网段内MAC与IP的对应关系
   2. 主机利用MAC帧传递数据
 * 同一物理网段之内,两个网段之间的主机无法以MAC帧的形式交换数据, 因为广播数据包无法查到MAC与IP的对应
# IP地址分级
Net ID越大,则Host ID越少, 这意味着网段内可分配的IP数量就越少.
||=级别   =||= 二进制表示                         =||= 十进制表示        =|| 
||=A Class=|| 0xxxxxxxx.xxxxxxxx.xxxxxxxx.xxxxxxxx ||0.*.*.* ~ 126.*.*.*  ||
||=B Class=|| 10xxxxxxx.xxxxxxxx.xxxxxxxx.xxxxxxxx ||128.*.*.* ~ 191.*.*.*||
||=C Class=|| 110xxxxxx.xxxxxxxx.xxxxxxxx.xxxxxxxx ||192.*.*.* ~ 223.*.*.*||
 127.*.*.*:: 该网段属于A类地址,保留给OS的Loopback
# Netmask
在共享媒介上面,任何主机要发言,就必须先使用CSMA/CD方式进行网络监听, 如果网络上主机们太多了,比如像Class A那样级别的网段, 能够容纳256*256*256-2台主机, 每个主机想发言,都在进行CSMA/CD会导致严重的停顿问题.因为一旦出现数据包冲突, 或者ARP查询MAC与IP的对应关系时, 需要响应的主机数量太多了.
一般我们的局域网使用C类网段, 也就是最多253台主机, 这其实也有点儿多, 通常要0台机器以内,网络性能会较好.
IP地址分级的方法也可以用在某个网段上面,以实现子网划分.
这个网段的Net ID肯定是没法改的, 我们要在这个网段中再划分出子网,就只能划分Host ID, 比如, 我们将Host ID的几位拿出来做子网的Net ID. 余下的部分做子网段Host ID, 这样我们就可以得到更多的网段,每个网段上的主机数量则更少.
 
# Hub / Switcher
# OSI 七层协议
||=分层=||=名称=||=说明=||
|-------------------------
```
#!td
1
```
```
#!td
PhysicalLayer 
```
```
#!td
PHY
```
|-------------------------
```
#!td
2
```
```
#!td
DataLinkLayer
```
```
#!td
 * ARP
```
|-------------------------
```
#!td
3
```
```
#!td
OsiNetworkLayer
```
```
#!td
 * ARP
```
|-------------------------
```
#!td
4
```
```
#!td
TransportLayer
```
```
#!td
 * ARP
```
|-------------------------
```
#!td
5
```
```
#!td
SessionLayer
```
```
#!td
 * ARP
```
|-------------------------
```
#!td
6
```
```
#!td
PresentationLayer
```
```
#!td
 * ARP
```
|-------------------------
```td
7
```
```td
ApplicationLayer
```
```td
 * ARP
```
# /etc/login.defs
||CONSOLE       ||                 || ||
||ENV_SUPATH    || PATH=/sbin:/bin || 超级用户登录后使用的环境变量     ||
||ENV_PATH      || PATH=/bin       || 普通用户登录后使用的环境变量 ||
||PASS_MAX_DAYS || N               || 用户密码有效期为N天          ||
||PASS_MIN_DAYS || N               || 用户两次更改密码至少间隔N天  ||
||PASS_WARN_AGE || N               || 提前N天通知用户更改密码      ||
||UID_MIN       || N               || useradd命令添加用户最小UID   ||
||UID_MAX       || N               || useradd命令添加用户最大UID   ||
||LOGIN_RETRIES || N               || 最多允许N次尝试登录          ||
||LOGIN_TIMEOUT || N               || N秒后登录超时                ||
||DEFAULT_HOME  || no/yes          || 允许无家目录的用户登录       ||          
[[TOC]]
# Linux File Permissions 
# 概述
 * 在GnuLinux中，每个用户都有自己的Linux帐号， 这个帐号属于一个或多个[LinuxUserGroup 用户组]
 * 每个文件也属于一个Linux帐号和一个[LinuxUserGroup 用户组]
 * 每个文件可以定义三种基本权限:
   * 读
   * 写
   * 执行
# 查看文件的属性 
你可以使用`ls -l`命令观察以上提到的这些文件属性.
```
#!div class=note
因为`ls -l`使用频率非常之高，所以你可以将它定义一个别名，放在你所使用的Shell配置文件中, 比如: .bashrc
```
alias ll='ls -l'
```
```
现在我们来观察以下ls命令的输出:
```
#!sh
$ ls -l 
drwxr-xr-x 3 amas users 4096 Jun 30 10:56 etc
-rwxr-xr-x 1 amas users  144 Apr 19 14:11 loop.sh
```
关注一下第一栏信息:
|| 符号 || 文件类型 ||
||-	       || 普通文件 ||
||d	       || 目录 ||
||l	       || [SymbolicLinc 符号链接] ||
||s	       || Socket ||
||p        || [NamedPipe 命名管道] ||
||c        || [CharacterDevice 字符设备] ||
||b        || [BlockedDevice 块设备]   ||
 
# 修改文件的权限
可以使用chmod命令修改文件属性，在此之前，让我们来介绍一下
```
        chmod [ugoa][+-] permission file
        chmod NNN file
```
permission:
|| u ||	用户帐号 ||
|| g ||	用户组 ||
|| o ||	其他用户 ||
|| a ||	所有权限组 ||
|| + || 添加指定的权限 ||
|| - || 去掉指定的权限 ||
来看几个例子:
```
#!sh
# 将文件file设置为其他用户可读/可写/可执行的
$ chmod o+rwx file.txt
# 将build.sh设为本用户不可执行的
$ chmod u-x build.sh
```
# 文件权限的数字表示法
||= r =||= w =||= x =||= - =||
||  4  ||  2  ||  1  || 0   ||
||= --- =|| 0+0+0=0 ||= 0 =|| 啥都不能   ||
||= --x =|| 0+0+1=1 ||= 1 =|| 只能执行   ||
||= -w- =|| 0+2+0=2 ||= 2 =|| 只能写     ||
||= -wx =|| 0+2+1=3 ||= 3 =|| 可写可执行  ||
||= r-- =|| 4+0+0=4 ||= 4 =|| 只读      ||
||= r-x =|| 4+0+1=5 ||= 5 =|| 可读可执行 ||
||= rw- =|| 4+2+0=6 ||= 6 =|| 可读可写   ||
||= rwx =|| 4+2+1=7 ||= 7 =|| 都行      ||
 
||||||=用户=||||||=组=||||||=其他=||
||||||=u=||||||=g||||||=o=||
||=r=||=w=||=x=||=r=||=w=||=x=||=r=||=w=||=x=||
几个例子:
```
        755 = -rwxr-xr-x
        660 = -rw-rw----
```
# 修改文件的所有权
```
        chown user:group file
```
# 文件的StickyBit
> 如果你需要一个任何用户都可以读写的目录，怎么能保证用户之间不能修改彼此的文件?
我们来观察一下`/tmp`目录的属性:
```
#!sh
$ ls -l / | grep tmp
drwxrwxrwt  31 root root  4096 Aug  3 16:48 tmp
```
注意，除了我们之前提到的文件属性， 这个`/tmp`目录最后一位是个't', 这个't'取代了原来的x权限， 它就是StickyBit, 有了它
用户可以在这个目录下自由的创建文件，但是彼此却不能修改.
你可以使用chmod来增加StickyBit:
```
#!sh
$ chmod +t shared-dir
$ chmod -t shared-dir
```
# 文件的SGID属性
默认情况下，用户所执行的程序只具有该用户，以及该用户默认组所拥有的权限， SGID可以使进程拥有目标文件拥有者的权限，而不是可执行文件本身的权限.
来看个例子:
```
$ ls amas-server
-rwxrwxrwx  10 amas root  4096 2006-03-10 12:50 logger.sh
-rwxrwx---  10 amas root  4096 2006-03-10 12:50 server.log
```
这里logger.sh是一个可执行文件， 它负责收集信息到server.log中， 我们希望:
 1. 任何登录用户都可以执行/修改logger.sh
 2. server.log只有管理员权限才能查看
如果按照以上的权限设置，用户x(非root组)执行logger.sh的时候，会遭遇无法写入server.log的错误， 因为当x执行logger.sh的时候， logger.sh对应的进程只拥有x的权限，它虽然可以执行logger.sh, 但是
这个进程却无法写入server.log文件。
这样的问题就需要使用SGID， 使x的logger.sh进程拥有logger.sh所属者(即:amas)的权限, 从而可以写入server.log
```
#!sh
$ chmod u+s logger.sh
```
同样，你也可以给组属性添加SGID:
```
#!sh
$ chmod g+s logger.sh
```
这样，用户x所执行的logger.sh进程将拥有root用户的权限.
# umask 
你在某个目录下，新建文件的权限属性由umask来决定:
你可以使用umask命令查看但前目录下的umask值
```
#!sh
$ umask 
022
```
如果你在该目录下建立了一个文件(或目录)，其默认权限属性可以通过一下公式计算:
```
        permission-of-dir  = 777 - umask
        permission-of-file = 666 - umask
```
本例: 777-022=755 即: `-rwxr-xr-x`
如果你想改变目录的umask值，可以使用
```
#!sh
$ umask 077
```
[[TOC]]
# 进程有ID
# 进程有父亲
# 进程有FD
# 进程有资源使用上的限制
```
getrlimit
```
```sh
$ cat /proc/<pid>/limits
Limit                     Soft Limit           Hard Limit           Units     
Max cpu time              unlimited            unlimited            seconds   
Max file size             unlimited            unlimited            bytes     
Max data size             unlimited            unlimited            bytes     
Max stack size            8388608              unlimited            bytes     
Max core file size        0                    unlimited            bytes     
Max resident set          unlimited            unlimited            bytes     
Max processes             14952                14952                processes 
Max open files            2048                 4096                 files     
Max locked memory         67108864             67108864             bytes     
Max address space         unlimited            unlimited            bytes     
Max file locks            unlimited            unlimited            locks     
Max pending signals       14952                14952                signals   
Max msgqueue size         819200               819200               bytes     
Max nice priority         40                   40                   
Max realtime priority     0                    0                    
Max realtime timeout      unlimited            unlimited            us
```
# 进程拥有自己的环境信息
# 进程有名字
# 进程有参数列表
# 进程有退出码
# 如何退出进程?
# exit <num>
 * <num> : 0-255
# abort
# 复制进程: fork(2)
```c
#include<stdio.h>
int main(char ** argv) {
    int pid = fork();
    printf("fork()=%d
", pid);
    if(pid) {
        printf("1> PID : %5d PARENT : %5d
", getpid(), getppid());
    } else {
        printf("2> PID : %5d PARENT : %5d
", getpid(), getppid());
    }
    return 0;
}
```
 * fork(2) 返回两次
   * 调用进程返回fork出来子进程pid
   * fork在子进程中返回0
 * fork成功后, 子进程是对父进程的复制, 内存使用会成倍增长, 恶意使用就是[ForkBomb Fork炸弹]
```div class=note
咱们来分析一下经典的shell Fork炸弹
```sh
bomb() {
    bomb | bomb&
}
bomb
```
如果运行这个函数, 出现的后果就是bomb递归调用自己, 
如果改成一个更屌的版本,只要做些文字上的替换, 看上去是不是更加邪恶?
```sh
:(){:|&};:
```
```
ecuring your Linux server is important to protect your data, intellectual property, and time, from the hands of crackers (hackers). The system administrator is responsible for security Linux box. In this first part of a Linux server security series, I will provide 20 hardening tips for default installation of Linux system.
#1: Encrypt Data Communication
All data transmitted over a network is open to monitoring. Encrypt transmitted data whenever possible with password or using keys / certificates.
Use scp, ssh, rsync, or sftp for file transfer. You can also mount remote server file system or your own home directory using special sshfs and fuse tools.
GnuPG allows to encrypt and sign your data and communication, features a versatile key managment system as well as access modules for all kind of public key directories.
Fugu is a graphical frontend to the commandline Secure File Transfer application (SFTP). SFTP is similar to FTP, but unlike FTP, the entire session is encrypted, meaning no passwords are sent in cleartext form, and is thus much less vulnerable to third-party interception. Another option is FileZilla - a cross-platform client that supports FTP, FTP over SSL/TLS (FTPS), and SSH File Transfer Protocol (SFTP).
OpenVPN is a cost-effective, lightweight SSL VPN.
Lighttpd SSL (Secure Server Layer) Https Configuration And Installation
Apache SSL (Secure Server Layer) Https (mod_ssl) Configuration And Installation
#1.1: Avoid Using FTP, Telnet, And Rlogin / Rsh
Under most network configurations, user names, passwords, FTP / telnet / rsh commands and transferred files can be captured by anyone on the same network using a packet sniffer. The common solution to this problem is to use either OpenSSH , SFTP, or FTPS (FTP over SSL), which adds SSL or TLS encryption to FTP. Type the following command to delete NIS, rsh and other outdated service:
# yum erase inetd xinetd ypserv tftp-server telnet-server rsh-serve
#2: Minimize Software to Minimize Vulnerability
Do you really need all sort of web services installed? Avoid installing unnecessary software to avoid vulnerabilities in software. Use the RPM package manager such as yum or apt-get and/or dpkg to review all installed set of software packages on a system. Delete all unwanted packages.
# yum list installed
# yum list packageName
# yum remove packageName
OR
# dpkg --list
# dpkg --info packageName
# apt-get remove packageName
#3: One Network Service Per System or VM Instance
Run different network services on separate servers or VM instance. This limits the number of other services that can be compromised. For example, if an attacker able to successfully exploit a software such as Apache flow, he / she will get an access to entire server including other services such as MySQL, e-mail server and so on. See how to install Virtualization software:
Install and Setup XEN Virtualization Software on CentOS Linux 5
How To Setup OpenVZ under RHEL / CentOS Linux
#4: Keep Linux Kernel and Software Up to Date
Applying security patches is an important part of maintaining Linux server. Linux provides all necessary tools to keep your system updated, and also allows for easy upgrades between versions. All security update should be reviewed and applied as soon as possible. Again, use the RPM package manager such as yum and/or apt-get and/or dpkg to apply all security updates.
# yum update
OR
# apt-get update && apt-get upgrade
You can configure Red hat / CentOS / Fedora Linux to send yum package update notification via email. Another option is to apply all security updates via a cron job. Under Debian / Ubuntu Linux you can use apticron to send security notifications.
#5: Use Linux Security Extensions
Linux comes with various security patches which can be used to guard against misconfigured or compromised programs. If possible use SELinux and other Linux security extensions to enforce limitations on network and other programs. For example, SELinux provides a variety of security policies for Linux kernel.
#5.1: SELinux
I strongly recommend using SELinux which provides a flexible Mandatory Access Control (MAC). Under standard Linux Discretionary Access Control (DAC), an application or process running as a user (UID or SUID) has the user's permissions to objects such as files, sockets, and other processes. Running a MAC kernel protects the system from malicious or flawed applications that can damage or destroy the system. See the official Redhat documentation which explains SELinux configuration.
#6: User Accounts and Strong Password Policy
Use the useradd / usermod commands to create and maintain user accounts. Make sure you have a good and strong password policy. For example, a good password includes at least 8 characters long and mixture of alphabets, number, special character, upper & lower alphabets etc. Most important pick a password you can remember. Use tools such as "John the ripper" to find out weak users passwords on your server. Configure pam_cracklib.so to enforce the password policy.
#6.1: Password Aging
The chage command changes the number of days between password changes and the date of the last password change. This information is used by the system to determine when a user must change his/her password. The /etc/login.defs file defines the site-specific configuration for the shadow password suite including password aging configuration. To disable password aging, enter:
chage -M 99999 userName
To get password expiration information, enter:
chage -l userName
Finally, you can also edit the /etc/shadow file in the following fields:
{userName}:{password}:{lastpasswdchanged}:{Minimum_days}:{Maximum_days}:{Warn}:{Inactive}:{Expire}:
Where,
Minimum_days: The minimum number of days required between password changes i.e. the number of days left before the user is allowed to change his/her password.
Maximum_days: The maximum number of days the password is valid (after that user is forced to change his/her password).
Warn : The number of days before password is to expire that user is warned that his/her password must be changed.
Expire : Days since Jan 1, 1970 that account is disabled i.e. an absolute date specifying when the login may no longer be used.
I recommend chage command instead of editing the /etc/shadow by hand:
# chage -M 60 -m 7 -W 7 userName
Recommend readings:
Linux: Force Users To Change Their Passwords Upon First Login
Linux turn On / Off password expiration / aging
Lock the user password
Search for all account without password and lock them
Use Linux groups to enhance security
#6.2: Restricting Use of Previous Passwords
You can prevent all users from using or reuse same old passwords under Linux. The pam_unix module parameter remember can be used to configure the number of previous passwords that cannot be reused.
#6.3: Locking User Accounts After Login Failures
Under Linux you can use the faillog command to display faillog records or to set login failure limits. faillog formats the contents of the failure log from /var/log/faillog database / log file. It also can be used for maintains failure counters and limits.To see failed login attempts, enter:
faillog
To unlock an account after login failures, run:
faillog -r -u userName
Note you can use passwd command to lock and unlock accounts:
# lock account
passwd -l userName
# unlocak account
passwd -u userName
#6.4: How Do I Verify No Accounts Have Empty Passwords?
Type the following command
# awk -F: '($2 == "") {print}' /etc/shadow
Lock all empty password accounts:
# passwd -l accountName
#6.5: Make Sure No Non-Root Accounts Have UID Set To 0
Only root account have UID 0 with full permissions to access the system. Type the following command to display all accounts with UID set to 0:
# awk -F: '($3 == "0") {print}' /etc/passwd
You should only see one line as follows:
root:x:0:0:root:/root:/bin/bash
If you see other lines, delete them or make sure other accounts are authorized by you to use UID 0.
#7: Disable root Login
Never ever login as root user. You should use sudo to execute root level commands as and when required. sudo does greatly enhances the security of the system without sharing root password with other users and admins. sudo provides simple auditing and tracking features too.
#8: Physical Server Security
You must protect Linux servers physical console access. Configure the BIOS and disable the booting from external devices such as DVDs / CDs / USB pen. Set BIOS and grub boot loader password to protect these settings. All production boxes must be locked in IDCs (Internet Data Center) and all persons must pass some sort of security checks before accessing your server. See also:
9 Tips To Protect Linux Servers Physical Console Access.
#9: Disable Unwanted Services
Disable all unnecessary services and daemons (services that runs in the background). You need to remove all unwanted services from the system start-up. Type the following command to list all services which are started at boot time in run level # 3:
# chkconfig --list | grep '3:on'
To disable service, enter:
# service serviceName stop
# chkconfig serviceName off
#9.1: Find Listening Network Ports
Use the following command to list all open ports and associated programs:
netstat -tulpn
OR
nmap -sT -O localhost
nmap -sT -O server.example.com
Use iptables to close open ports or stop all unwanted network services using above service and chkconfig commands.
#9.2: See Also
update-rc.d like command on Redhat Enterprise / CentOS Linux.
Ubuntu / Debian Linux: Services Configuration Tool to Start / Stop System Services.
Get Detailed Information About Particular IP address Connections Using netstat Command.
#10: Delete X Windows
X Windows on server is not required. There is no reason to run X Windows on your dedicated mail and Apache web server. You can disable and remove X Windows to improve server security and performance. Edit /etc/inittab and set run level to 3. Finally, remove X Windows system, enter:
# yum groupremove "X Window System"
#11: Configure Iptables and TCPWrappers
Iptables is a user space application program that allows you to configure the firewall (Netfilter) provided by the Linux kernel. Use firewall to filter out traffic and allow only necessary traffic. Also use the TCPWrappers a host-based networking ACL system to filter network access to Internet. You can prevent many denial of service attacks with the help of Iptables:
Lighttpd Traffic Shaping: Throttle Connections Per Single IP (Rate Limit).
How to: Linux Iptables block common attack.
psad: Linux Detect And Block Port Scan Attacks In Real Time.
#12: Linux Kernel /etc/sysctl.conf Hardening
/etc/sysctl.conf file is used to configure kernel parameters at runtime. Linux reads and applies settings from /etc/sysctl.conf at boot time. Sample /etc/sysctl.conf:
# Turn on execshield
kernel.exec-shield=1
kernel.randomize_va_space=1
# Enable IP spoofing protection
net.ipv4.conf.all.rp_filter=1
# Disable IP source routing
net.ipv4.conf.all.accept_source_route=0
# Ignoring broadcasts request
net.ipv4.icmp_echo_ignore_broadcasts=1
net.ipv4.icmp_ignore_bogus_error_messages=1
# Make sure spoofed packets get logged
net.ipv4.conf.all.log_martians = 1
#13: Separate Disk Partitions
Separation of the operating system files from user files may result into a better and secure system. Make sure the following filesystems are mounted on separate partitions:
/usr
/home
/var and /var/tmp
/tmp
Create septate partitions for Apache and FTP server roots. Edit /etc/fstab file and make sure you add the following configuration options:
noexec - Do not set execution of any binaries on this partition (prevents execution of binaries but allows scripts).
nodev - Do not allow character or special devices on this partition (prevents use of device files such as zero, sda etc).
nosuid - Do not set SUID/SGID access on this partition (prevent the setuid bit).
Sample /etc/fstab entry to to limit user access on /dev/sda5 (ftp server root directory):
/dev/sda5  /ftpdata          ext3    defaults,nosuid,nodev,noexec 1 2
#13.1: Disk Quotas
Make sure disk quota is enabled for all users. To implement disk quotas, use the following steps:
Enable quotas per file system by modifying the /etc/fstab file.
Remount the file system(s).
Create the quota database files and generate the disk usage table.
Assign quota policies.
See implementing disk quotas tutorial for further details.
#14: Turn Off IPv6
Internet Protocol version 6 (IPv6) provides a new Internet layer of the TCP/IP protocol suite that replaces Internet Protocol version 4 (IPv4) and provides many benefits. Currently there are no good tools out which are able to check a system over network for IPv6 security issues. Most Linux distro began enabling IPv6 protocol by default. Crackers can send bad traffic via IPv6 as most admins are not monitoring it. Unless network configuration requires it, disable IPv6 or configure Linux IPv6 firewall:
RedHat / Centos Disable IPv6 Networking.
Debian / Ubuntu And Other Linux Distros Disable IPv6 Networking.
Linux IPv6 Howto - Chapter 19. Security.
Linux IPv6 Firewall configuration and scripts are available here.
#15: Disable Unwanted SUID and SGID Binaries
All SUID/SGID bits enabled file can be misused when the SUID/SGID executable has a security problem or bug. All local or remote user can use such file. It is a good idea to find all such files. Use the find command as follows:
#See all set user id files:
find / -perm +4000
# See all group id files
find / -perm +2000
# Or combine both in a single command
find / ( -perm -4000 -o -perm -2000 ) -print
find / -path -prune -o -type f -perm +6000 -ls
You need to investigate each reported file. See reported file man page for further details.
#15.1: World-Writable Files
Anyone can modify world-writable file resulting into a security issue. Use the following command to find all world writable and sticky bits set files:
find /dir -xdev -type d ( -perm -0002 -a ! -perm -1000 ) -print
You need to investigate each reported file and either set correct user and group permission or remove it.
#15.2: Noowner Files
Files not owned by any user or group can pose a security problem. Just find them with the following command which do not belong to a valid user and a valid group
find /dir -xdev ( -nouser -o -nogroup ) -print
You need to investigate each reported file and either assign it to an appropriate user and group or remove it.
#16: Use A Centralized Authentication Service
Without a centralized authentication system, user auth data becomes inconsistent, which may lead into out-of-date credentials and forgotten accounts which should have been deleted in first place. A centralized authentication service allows you maintaining central control over Linux / UNIX account and authentication data. You can keep auth data synchronized between servers. Do not use the NIS service for centralized authentication. Use OpenLDAP for clients and servers.
#16.1: Kerberos
Kerberos performs authentication as a trusted third party authentication service by using cryptographic shared secret under the assumption that packets traveling along the insecure network can be read, modified, and inserted. Kerberos builds on symmetric-key cryptography and requires a key distribution center. You can make remote login, remote copy, secure inter-system file copying and other high-risk tasks safer and more controllable using Kerberos. So, when users authenticate to network services using Kerberos, unauthorized users attempting to gather passwords by monitoring network traffic are effectively thwarted. See how to setup and use Kerberos.
#17: Logging and Auditing
You need to configure logging and auditing to collect all hacking and cracking attempts. By default syslog stores data in /var/log/ directory. This is also useful to find out software misconfiguration which may open your system to various attacks. See the following logging related articles:
Linux log file locations.
How to send logs to a remote loghost.
How do I rotate log files?.
man pages syslogd, syslog.conf and logrotate.
#17.1: Monitor Suspicious Log Messages With Logwatch / Logcheck
Read your logs using logwatch or logcheck. These tools make your log reading life easier. You get detailed reporting on unusual items in syslog via email. A sample syslog report:
 ################### Logwatch 7.3 (03/24/06) ####################
        Processing Initiated: Fri Oct 30 04:02:03 2009
        Date Range Processed: yesterday
                              ( 2009-Oct-29 )
                              Period is day.
      Detail Level of Output: 0
              Type of Output: unformatted
           Logfiles for Host: www-52.nixcraft.net.in
  ##################################################################
 --------------------- Named Begin ------------------------
 **Unmatched Entries**
    general: info: zone XXXXXX.com/IN: Transfer started.: 3 Time(s)
    general: info: zone XXXXXX.com/IN: refresh: retry limit for master ttttttttttttttttttt#53 exceeded (source ::#0): 3 Time(s)
    general: info: zone XXXXXX.com/IN: Transfer started.: 4 Time(s)
    general: info: zone XXXXXX.com/IN: refresh: retry limit for master ttttttttttttttttttt#53 exceeded (source ::#0): 4 Time(s)
 ---------------------- Named End -------------------------
  --------------------- iptables firewall Begin ------------------------
 Logged 87 packets on interface eth0
   From 58.y.xxx.ww - 1 packet to tcp(8080)
   From 59.www.zzz.yyy - 1 packet to tcp(22)
   From 60.32.nnn.yyy - 2 packets to tcp(45633)
   From 222.xxx.ttt.zz - 5 packets to tcp(8000,8080,8800)
 ---------------------- iptables firewall End -------------------------
 --------------------- SSHD Begin ------------------------
 Users logging in through sshd:
    root:
       123.xxx.ttt.zzz: 6 times
 ---------------------- SSHD End -------------------------
 --------------------- Disk Space Begin ------------------------
 Filesystem            Size  Used Avail Use% Mounted on
 /dev/sda3             450G  185G  241G  44% /
 /dev/sda1              99M   35M   60M  37% /boot
 ---------------------- Disk Space End -------------------------
 ###################### Logwatch End #########################
(Note output is truncated)
#17.2: System Accounting with auditd
The auditd is provided for system auditing. It is responsible for writing audit records to the disk. During startup, the rules in /etc/audit.rules are read by this daemon. You can open /etc/audit.rules file and make changes such as setup audit file log location and other option. With auditd you can answers the following questions:
System startup and shutdown events (reboot / halt).
Date and time of the event.
User respoisble for the event (such as trying to access /path/to/topsecret.dat file).
Type of event (edit, access, delete, write, update file & commands).
Success or failure of the event.
Records events that Modify date and time.
Find out who made changes to modify the system's network settings.
Record events that modify user/group information.
See who made changes to a file etc.
See our quick tutorial which explains enabling and using the auditd service.
#18: Secure OpenSSH Server
The SSH protocol is recommended for remote login and remote file transfer. However, ssh is open to many attacks. See how to secure OpenSSH server:
Top 20 OpenSSH Server Best Security Practices.
#19: Install And Use Intrusion Detection System
A network intrusion detection system (NIDS) is an intrusion detection system that tries to detect malicious activity such as denial of service attacks, port scans or even attempts to crack into computers by monitoring network traffic.
It is a good practice to deploy any integrity checking software before system goes online in a production environment. If possible install AIDE software before the system is connected to any network. AIDE is a host-based intrusion detection system (HIDS) it can monitor and analyses the internals of a computing system.
Snort is a software for intrusion detection which is capable of performing packet logging and real-time traffic analysis on IP networks.
#20: Protecting Files, Directories and Email
Linux offers excellent protections against unauthorized data access. File permissions and MAC prevent unauthorized access from accessing data. However, permissions set by the Linux are irrelevant if an attacker has physical access to a computer and can simply move the computer's hard drive to another system to copy and analyze the sensitive data. You can easily protect files, and partitons under Linux using the following tools:
To encrypt and decrypt files with a password, use gpg command.
Linux or UNIX password protect files with openssl and other tools.
See how to encrypting directories with ecryptfs.
TrueCrypt is free open-source disk encryption software for Windows 7/Vista/XP, Mac OS X and Linux.
Howto: Disk and partition encryption in Linux for mobile devices.
How to setup encrypted Swap on Linux.
#20.1: Securing Email Servers
You can use SSL certificates and gpg keys to secure email communication on both server and client computers:
Linux Securing Dovecot IMAPS / POP3S Server with SSL Configuration.
Linux Postfix SMTP (Mail Server) SSL Certificate Installations and Configuration.
Courier IMAP SSL Server Certificate Installtion and Configuration.
Configure Sendmail SSL encryption for sending and receiving email.
Enigmail: Encrypted mail with Mozilla thunderbird.
Other Recommendation:
Backups - It cannot be stressed enough how important it is to make a backup of your Linux system. A proper offsite backup allows you to recover from cracked server i.e. an intrusion. The traditional UNIX backup programs are dump and restore are also recommended.
How to: Looking for Rootkits.
Howto: Enable ExecShield Buffer Overflows Protection.
Subscribe to Redhat or Debian Linux security mailing list or RSS feed.
Recommend readings:
Red Hat Enterprise Linux - Security Guide.
Linux security cookbook- A good collections of security recipes for new Linux admin.
Snort 2.1 Intrusion Detection, Second Edition - Good introduction to Snort and Intrusion detection under Linux.
Hardening Linux - Hardening Linux identifies many of the risks of running Linux hosts and applications and provides practical examples and methods to minimize those risks.
Linux Security HOWTO.
In the next part of this series I will discuss how to secure specific applications (such as Proxy, Mail, LAMP, Database) and a few other security tools. Did I miss something? Please add your favorite system security tool or tip in the comments.
Featured Articles:
20 Linux System Monitoring Tools Every SysAdmin Should Know
20 Linux Server Hardening Security Tips
Linux: 20 Iptables Examples For New SysAdmins 
My 10 UNIX Command Line Mistakes
25 PHP Security Best Practices For Sys Admins
The Novice Guide To Buying A Linux Laptop
Top 5 Email Client For Linux, Mac OS X, and Windows Users
Top 20 OpenSSH Server Best Security Practices
Top 10 Open Source Web-Based Project Management Software
# 参考
 * http://www.cyberciti.biz/tips/linux-security.html
# systemd
# 配置文件
# /usr/lib/systemd/system: 存放unit配置
# 使用方法
# systemd enable
# systemd disable
# systemd start
# systemd stop
# systemd restart
# hostnamectl set-hostname
# localectl set-locale 
# localectl set-keymap 
# localectl set-x11-keymap
# timedatectl set-timezone
# 使用本地时钟: timedatectl set-local-rtc true
# 使用UTC: timedatectl set-local-rtc false
# 添加unit
 * socket
 * service
 * target
 * 
```div class=warn
===注意
ExecStart等配置项中的可执行文件必须以绝对路径的方式给出。否则
```
# 故障排除
```
# 查看日志
$ journalctl -xn
$ systemctl --system daemon-reload
```
# 例子
```ini
[Unit]
Description=A high performance web server and a reverse proxy server
After=syslog.target network.target
[Service]
Type=forking
PIDFile=/run/nginx.pid
ExecStartPre=/usr/sbin/nginx -t -q -g 'pid /run/nginx.pid; daemon on; master_process on;'
ExecStart=/usr/sbin/nginx -g 'pid /run/nginx.pid; daemon on; master_process on;'
ExecReload=/usr/sbin/nginx -g 'pid /run/nginx.pid; daemon on; master_process on;' -s reload
ExecStop=/usr/sbin/nginx -g 'pid /run/nginx.pid;' -s quit
[Install]
WantedBy=multi-user.target
```
# 参考
 * http://freedesktop.org/wiki/Software/systemd/
 * http://0pointer.de/blog/projects/systemd.html
[[TOC]]
# All about Linux signals
In most cases if you want to handle a signal in your application you write a simple signal handler like:
```c
void handler (int sig) 
{
   ...
}
// register signal handler
signal(SIGINT, handler);
```
and use the signal(2) system function to run it when a signal is delivered to the process. This is the simplest case, but signals are more interesting than that! Information contained in this article is useful for example when you are writing a daemon and must handle interrupting your program properly without interrupting the current operation or the whole program.
# What is signaled in Linux
Your process may receive a signal when:
 * From user space from some other process when someone calls a function like kill(2).
 * When you send the signal from the process itself using a function like abort(3).
 * When a child process exits the operating system sends the SIGCHLD signal.
 * When the parent process dies or hangup is detected on the controlling terminal SIGHUP is sent.
 * When user interrupts program from the keyboard SIGINT is sent.
 * When the program behaves incorrectly one of SIGILL, SIGFPE, SIGSEGV is delivered.
 * When a program accesses memory that is mapped using mmap(2) but is not available (for example when the file was truncated by another process) - really nasty situation when using mmap() to access files. There is no good way to handle this case.
 * When a profiler like gprof is used the program occasionally receives SIGPROF. This is sometimes problematic when you forgot to handle interrupting system functions like read(2) properly (errno == EINTR).
 * When you use the write(2) or similar data sending functions and there is nobody to receive your data SIGPIPE is delivered. This is a very common case and you must remember that those functions may not only exit with error and setting the errno variable but also cause the SIGPIPE to be delivered to the program. An example is the case when you write to the standard output and the user uses the pipeline sequence to redirect your output to another program. If the program exits while you are trying to send data SIGPIPE is sent to your process. A signal is used in addition to the normal function return with error because this event is asynchronous and you can't actually tell how much data has been successfully sent. This can also happen when you are sending data to a socket. This is because data are buffered and/or send over a wire so are not delivered to the target immediately and the OS can realize that can't be delivered after the sending function exits.
For a complete list of signals see the signal(7) manual page.
# The recommended way of setting signal actions: sigaction
The sigaction(2) function is a better way to set the signal action. It has the prototype:
```c
int sigaction (int signum, const struct sigaction *act, struct sigaction *oldact);
```
As you can see you don't pass the pointer to the signal handler directly, but instead a struct sigaction object. It's defined as:
```c
struct sigaction {
        void     (*sa_handler)(int);
        void     (*sa_sigaction)(int, siginfo_t *, void *);
        sigset_t   sa_mask;
        int        sa_flags;
        void     (*sa_restorer)(void);
};
```
For a detailed description of this structure's fields see the sigaction(2) manual page. Most important fields are:
 * sa_handler - This is the pointer to your handler function that has the same prototype as a handler for signal(2).
 * sa_sigaction - This is an alternative way to run the signal handler. It has two additional arguments beside the signal number where the siginfo_t * is the more interesting. It provides more information about the received signal, I will describe it later.
 * sa_mask allows you to explicitly set signals that are blocked during the execution of the handler. In addition if you don't use the SA_NODEFER flag the signal which triggered will be also blocked.
 * sa_flags allow to modify the behavior of the signal handling process. For the detailed description of this field, see the manual page. To use the sa_sigaction handler you must use SA_SIGINFO flag here.
What is the difference between signal(2) and sigaction(2) if you don't use any additional feature the later one provides? The answer is: portability and no race conditions. The issue with resetting the signal handler after it's called doesn't affect sigaction(2), because the default behavior is not to reset the handler and blocking the signal during it's execution. So there is no race and this behavior is documented in the POSIX specification. Another difference is that with signal(2) some system calls are automatically restarted and with sigaction(2) they're not by default.
# Example use of sigaction()
See example of using sigaction() to set a signal handler with additional parameters.
In this example we use the three arguments version of signal handler for SIGTERM. Without setting the SA_SIGINFO flag we would use a traditional one argument version of the handler and pass the pointer to it by the sa_handler field. It would be a replacement for signal(2). You can try to run it and do kill PID to see what happens.
In the signal handler we read two fields from the siginfo_t *siginfo parameter to read the sender's PID and UID. This structure has more fields, I'll describe them later.
The sleep(3) function is used in a loop because it's interrupted when the signal arrives and must be called again.
SA_SIGINFO handler
In the previous example SA_SIGINFO is used to pass more information to the signal handler as arguments. We've seen that the siginfo_t structure contains si_pid and si_uid fields (PID and UID of the process that sends the signal), but there are many more. They are all described in sigaction(2) manual page. On Linux only si_signo (signal number) and si_code (signal code) are available for all signals. Presence of other fields depends on the signal type. Some other fields are:
si_code - Reason why the signal was sent. It may be SI_USER if it was delivered due to kill(2) or raise(3), SI_KERNEL if kernel sent it and few more. For some signals there are special values like ILL_ILLADR telling you that SIGILL was sent due to illegal addressing mode.
For SIGCHLD fields si_status, si_utime, si_stime are filled and contain information about the exit status or the signal of the dying process, user and system time consumed.
In case of SIGILL, SIGFPE, SIGSEGV, SIGBUS si_addr contains the memory address that caused the fault.
We'll see more examples of use of siginfo_t later.
Compiler optimization and data in signal handler
```c
Let's see the following example:
#include <stdio.h>
#include <unistd.h>
#include <signal.h>
#include <string.h>
 
static int exit_flag = 0;
 
static void hdl (int sig)
{
	exit_flag = 1;
}
 
int main (int argc, char *argv[])
{
	struct sigaction act;
 
	memset (&act, '



