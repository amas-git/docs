---
title: Play Netcat
tags:
---
<!-- toc -->
# Netcat 
 * [http://netcat.sourceforge.net/ Home]
 * GPL
## 简介
Netcat(or 'nc') is a featured networking utility which reads and writes data across network connections, using the TCP/IP protocol.
It is designed to be a reliable "back-end" tool that can be used directly or easily driven by other programs and scripts. At the same time, it is a feature-rich network debugging and exploration tool, since it can create almost any kind of connection you would need and has several interesting built-in capabilities.
It provides access to the following main features:
 * Outbound and inbound connections, TCP or UDP, to or from any ports.
 * Featured tunneling mode which allows also special tunneling such as UDP to TCP, with the possibility of specifying all network parameters (source port/interface, listening port/interface, and the remote host allowed to connect to the tunnel.
 * Built-in port-scanning capabilities, with randomizer.
 * Advanced usage options, such as buffered send-mode (one line every N seconds), and hexdump (to stderr or to a specified file) of trasmitted and received data.
 * Optional RFC854 telnet codes parser and responder
netcat is possibly the "udp telnet-like" application you always wanted for testing your UDP-mode servers.

## 安装
## 为什么不使用telnet呢?
 * Telnet has the "standard input EOF" problem , so one must introduce calculated delays in driving scripts to allow network output to finish.
 * Telnet also will not transfer arbitrary binary data, because certain characters are interpreted as telnet options and are thus removed from the data stream.
 * Telnet also emits some of its diagnostic messages to standard output, where netcat keeps such things religiously separated from its *output* and will never modify any of the real data in transit unless you *really* want it to.
 * telnet is incapable of listening for inbound connections, or using UDP instead.
features:
 	* Outbound or inbound connections, TCP or UDP, to or from any ports
	* Full DNS forward/reverse checking, with appropriate warnings
	* Ability to use any local source port
	* Ability to use any locally-configured network source address
	* Built-in port-scanning capabilities, with randomizer
	* Built-in loose source-routing capability
	* Can read command line arguments from standard input
	* Slow-send mode, one line every N seconds
	* Hex dump of transmitted and received data
	* Optional ability to let another program service established connections
	* Optional telnet-options responder

## 使用方法
### 使用nc作为客户端
```sh
# nc HOST PORT 建立一个连接, 类似于telnet, 我们用这个方法通过HTTP协议取回www.baidu.com的首页
$ nc www.baidu.com 80
GET / HTTP/1.1
# ---------- 服务器返回的报文 ----------
HTTP/1.1 200 OK
Date: Mon, 16 Jan 2012 03:23:55 GMT
Server: BWS/1.0
Content-Length: 7677
...
```
### 使用nc作为服务器
ListenMode使得netcat等待InboundConnection.
```sh
#---------------------[ Server: localhost ]
$ cat file
hello client
# 一旦有客户端connect 8000端口，将file中的内容返回给客户端, 并断开连接
$ nc -v -l -p 8000 < file
#---------------------[ Client: localhost ]
$ nc localhost 8000
hello client
```
### 1. 一个简单的例子
现在我们使用nc进行文件传输试验:
```sh
#---------------------[ Server: 192.168.1.100 ]
$ nc -l -p 8000
#---------------------[ Client: 192.168.1.102 ]
$ cat README
hello nc
$ cat README | nc 192.168.1.100 8000
```
Server端:
```
hello nc
```

### 2. 传输一个文件
```sh
#---------------------[ Server: 192.168.1.100 ]
$ nc -l -p 8000 < file
#---------------------[ Client: 192.168.1.102 ]
# 接收文件
$ nc 192.168.1.100 8000 > file
```
 
### 3. 传输整个目录
```sh
#---------------------[ Server: 192.168.1.100 ]
$ nc -l -p 8000 | uncompress -c | tar xvfp -
#---------------------[ Client: 192.168.1.102 ]
# 我们将hello-nc目录下的全部文件传输到192.168.1.100
$ tar cfp - ./hello-nc | compress -c | nc -w 3 192.168.1.100 8000
```
## 参数
### -n 
HostIp可以是IP地址或域名，当指定-n, 则强制使用主机IP, netcat不做DNSLookups.
### port
OutboundConnections必须指定port, 可以是:
 * 数字,具体的端口号
 * `/etc/services`中配置的端口名
### -w seconds
-w用于指定连接超时, 交互超时
```sh
$ alias nc='nc -v -w 3'
$ nc localhost 80
127.0.0.1:80 (localhost.localdomain) open
# 如不3秒内输入请求，则连接断开
```
### -o log
你可以将整个交互过程dump到文件中
```sh
$ nc localhost 80 -o log
GET / HTTP/1.1
# 查看dump文件，其中包含数据的十六进制输出+AscII输出
# < : 表示请求(to the net)
# > : 表示响应(from the net)
$ cat log
< 00000000 47 45 54 20 2f 20 48 54 54 50 2f 31 2e 30 0a    | GET / HTTP/1.0.
< 0000000f 0a                                              | .
> 00000000 48 54 54 50 2f 31 2e 31 20 32 30 30 20 4f 4b 0d | HTTP/1.1 200 OK.
> 00000010 0a 44 61 74 65 3a 20 4d 6f 6e 2c 20 31 36 20 4a | .Date: Mon, 16 J
```
### -e prog
如果编译时指定了 `-DGAPING_SECURITY_HOLE`, 则连接成功后，
可以使用'-e'启动一个进程， 这类似于inetd, 但是
nc只是单例。
现在我们来做一个简单的BackdoorShell
宿主机上(192.168.1.100):
```sh
$ nc -l -p 8000 -e /bin/sh
```
客户端机(192.168.1.102)
```sh
$ echo 'cat /etc/passwd' | nc 192.168.1.100 8000
root:x:0:0:root:/root:/bin/bash
bin:x:1:1:bin:/bin:/bin/false
daemon:x:2:2:daemon:/sbin:/bin/false
...
```
### -i seconds
netcat默认使用8k读/写数据，如果你想指定读写间隔时间(通常用于调试), 可以使用-i.
这样读写总是以指定时间间隔进行。
### -z :  端口扫描
nc可以进行端口扫描, -z 参数的意思是'zero-I/O mode', nc会连接端口， 但不会通过TCP连接发送任何数据。
```sh
$ nc -v -w 2 -z $TARGET_HOST $PORT_MIN-$PORT_MAX
# 扫描本机80到1000端口
$ nc -v -w 2 -z localhost 80-1000
# 可以使用 -i 指定端口之间的扫描间隔时间
$ nc -v -w 2 -l 1 -z localhost 80-1000
# -r: 随机扫描
$ nc -v -w 2 -r -z localhost 80-1000
```

## 7788
使用nc进行简单的网络性能测试:
```sh
# 测试接收能力
$ yes AAAAAAAAAAAAAAAAAAAAAA | nc -v -v -l -p 2222 > /dev/null
# 测试发送能力
$ yes BBBBBBBBBBBBBBBBBBBBBB | nc othermachine 2222 > /dev/null
```
You can use netcat to protect your own workstation's X server against outside
access.  X is stupid enough to listen for connections on "any" and never tell
you when new connections arrive, which is one reason it is so vulnerable.  Once
you have all your various X windows up and running you can use netcat to bind
just to your ethernet address and listen to port 6000.  Any new connections
from outside the machine will hit netcat instead your X server, and you get a
log of who's trying.  You can either tell netcat to drop the connection, or
perhaps run another copy of itself to relay to your actual X server on
"localhost".  This may not work for dedicated X terminals, but it may be
possible to authorize your X terminal only for its boot server, and run a relay
netcat over on the server that will in turn talk to your X terminal.  Since
netcat only handles one listening connection per run, make sure that whatever
way you rig it causes another one to run and listen on 6000 soon afterward, or
your real X server will be reachable once again.  A very minimal script just
to protect yourself could be:
```sh
#!/bin/sh
while true ; do
    nc -v -l -s  -p 6000 localhost 2
done
```

## 简单的媒体流服务器
```sh
# server 
$ nc -l -p 9999 < music.mp3
# client
$ nc <server-ip> -p 9999 | mplayer -
```
