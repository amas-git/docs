---
title: 什么是TCP KeepAlive
tags:
---
# TCP Keepalive
<!-- toc -->

## 简介
Keepalive的概念非常简单: 当你建立了TCP链接，这个链接将与一个Timer集合相关联。 一些Timer用于处理keepalive。 当KeepaliveTimer变为0时， 你将发送给对端一个KeepaliveProbe报文（设置了ACK标志位，但是不包含任何数据的报文）, 这样做完全符合TCP/IP协议， 对端收到KeepaliveProbe后，根据TCP/IP协议，将会发送给你一个ACK确认，这个ACK确认也不包含任何其他数据。
如果你的KeepaliveProbe得到了回应，说明当前的TCP/IP链接一切正常。
这个过程的实际意义在于，当对端失去链接时(比如: 对端的网线断了或是机器突然重启)， 即便当前与对端无任何数据交换，你仍然能够察觉链接出了问题，可以及时做出处理。
应用层完全察觉不到以上过程，因为整个过程并无TCP/IP上层协议的数据交换。
## Keepalive的两个使命
### 1. Checking for dead peers
已经建立了TCP链接，但是因为各种原因出现故障，并且无法得到对端的通知，这时候Keepalive机制可以发现死链接。
这种故障的原因有很多:
 * 对端KernalPanic
 * 对端服务崩溃，进程终止
 * 物理层出了问题，比如网线断掉了

### 2. Preventing disconnection due to network inactivity
当你置于NatProxy或是防火墙后面时，经常会被中断链接而得不到任何通知。 这是因为，NatProxy和防火墙的链接跟踪机制需要追踪全部的链接，这需要以消耗内存为代价，一但系统资源不足，
NatProxy或防火墙就要取消对一些链接的追踪，关闭一些最不常用的链接，你很可能在其中。如果你启用了Keepalive, 定时发送的ACK可以确保你的链接在没有任何数据交换的情况下，仍然
能够被NatProxy或防火墙认为有一定的活跃度，因而不会被轻易淘汰, 或者即便被淘汰也可及时发现(见:1)。

## Linux中的keepalive
Linux本身支持keepalive. 你需要
有三个内核参数
 * tcp_keepalive_time : 发送最后一条数据到发送KeepaliveProbe的时间间隔(KeepaliveProbe不算数据), 单位: 秒
 * tcp_keepalive_intvl : 两个KeepaliveProbe之间的间隔时间, 单位: 秒
 * tcp_keepalive_probes : KeepaliveProbe失效N次后才认为链接已死，通知应用层
另外，即便内核已经开启Keepalive(一般默认关闭), 程序仍然需要通过设置Socket选项才能使用Keepalive.

### 配置内核
 * procfs interface
 * sysctl interface

### 查看
查看内核对Keepalive的当前设定:
```sh
$ cat /proc/sys/net/ipv4/tcp_keepalive_time
7200
$ cat /proc/sys/net/ipv4/tcp_keepalive_intvl
75
$ cat /proc/sys/net/ipv4/tcp_keepalive_probes
9
```
或者使用sysctl命令查看:
```sh
$ sysctl net.ipv4.tcp_keepalive_time net.ipv4.tcp_keepalive_intvl net.ipv4.tcp_keepalive_probes
```
### 修改
直接修改内核文件即可配置Keepalive:
```sh
$ echo 600 > /proc/sys/net/ipv4/tcp_keepalive_time
$ echo 60  > /proc/sys/net/ipv4/tcp_keepalive_intvl
$ echo 20  > /proc/sys/net/ipv4/tcp_keepalive_probes
```
或者使用`sysctl -w`命令查看
```sh
$ sysctl -w  net.ipv4.tcp_keepalive_time=600 net.ipv4.tcp_keepalive_intvl=60 net.ipv4.tcp_keepalive_probes=20
```
如果打算长久使用这些配置，可以加入到init脚本中，或者直接配置到`/etc/sysctl.conf`文件中

