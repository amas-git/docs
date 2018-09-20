---
title: tcpdump
tags:
---
<!-- toc -->
## tcpdump
> tcpdump  tcp|udp    src|dst    host    port
### -v
### -vv
### -D [eth0 | any | lo]
### -n 
### -q
### -i <interface>
### -F /path/to/file
### -X
### -S
### 协议
 * ether
 * fddi
 * ip
 * arp
 * rarp
 * decnet
 * lat
 * sca
 * moprc
 * mopdl
 * tcp 
 * udp

### 通讯方向
 * src
 * dst
 * src and dst
 * src or dst
```
host 192.168.1.100 等价于 src or dst 192.168.1.100
```
### 主机
 * net
 * port
 * host
 * portrange
```
src 10.1.1.1" is equivalent to "src host 10.1.1.1
```
### 逻辑运算符
 * and
 * or
 * not

```
"not tcp port 3128 and tcp port 23" is equivalent to "(not tcp port 3128) and tcp port 23".
"not tcp port 3128 and tcp port 23" is NOT equivalent to "not (tcp port 3128 and tcp port 23)".
```

```sh
# The tcpdump is simple command that dump traffic on a network. However, you need good understanding of TCP/IP protocol to utilize this tool. For.e.g to display traffic info about DNS, enter:
$ tcpdump -i eth1 'udp port 53'
# To display all IPv4 HTTP packets to and from port 80, i.e. print only packets that contain data, not, for example, SYN and FIN packets and ACK-only packets, enter:
$ tcpdump 'tcp port 80 and (((ip[2:2] - ((ip[0]&0xf)<<2)) - ((tcp[12]&0xf0)>>2)) != 0)'
# To display all FTP session to 202.54.1.5, enter:
$ tcpdump -i eth1 'dst 202.54.1.5 and (port 21 or 20'
# To display all HTTP session to 192.168.1.5:
$ tcpdump -ni eth0 'dst 192.168.1.5 and tcp and port http'
# Use wireshark to view detailed information about files, enter:
$ tcpdump -n -i eth1 -s 0 -w output.txt src or dst port 80
```
### 示例
```sh
$ tcpdump -s 0 -X 'tcp dst port 80' 
```
### 参考
 * [http://danielmiessler.com/study/tcpdump/ 非常好的入门文档]

