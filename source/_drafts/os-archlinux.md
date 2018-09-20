---
title: archlinux
tags:
---
# ArchLinux
# 磁盘分区 ==
 * 注意最好给/boot单独分区
 * /etc必需与/在同一分区
# 调试网络 ==
```
$ ifconfig <ip> netmask <netmask> 
```
网关:
```
#!sh
$ route add default gw <gw ip>
```
如果使用dhcp:
```
#!sh
$ dhcpcd eth0
```
有些交换机会在网卡适配mtu时无法成功，可以尝试关闭dhcp的interface_mtu选项(`/etc/dhcpd.conf`)
# 无线网络 ==
查看设备:
```
#!sh
$ lspci | grep -i net
# 有些笔记本的网卡可能会是usb设备
$ lsudb 
```
如果设备驱动没有问题，那么开始安装无线网络工具:
```
#!sh
$ pacman -S wireless_tools
```
一些网卡使用前需要激活kernel interface
```
$ sudo ifconfig wlan0 up
```
现在我们来扫描周围有哪些可用热点:
```
#!sh
$ sudo wlist wlan0 scan
```
连接到热点，取决于热点的加密协议:
```
#!sh
# 不加密
$ iwconfig wlan0 essid "MyEssid"
# WEP
$ iwconfig wlan0 essid "MyEssid" key 1234567890
# WEP. ASCII Key:
$ iwconfig wlan0 essid "MyEssid" key s:asciikey
# WPA/WPA2
# You need to edit the /etc/wpa_supplicant.conf file as described in WPA_Supplicant. Then, issue this command:
$ wpa_supplicant -B -Dwext -i wlan0 -c /etc/wpa_supplicant.conf
# 这里假定你的设备使用wext驱动, 如果你不清楚，请看: https://wiki.archlinux.org/index.php/WPA_Supplicant
```
连接到热点后，你需要获得一个IP:
```
#!sh
$ dhcpcd wlan0
# 或者
$ ifconfig wlan0 192.168.0.2
$ route add default gw 192.168.0.1
```
# 使用NetworkProfile工具 ===
使用wicd管理网络连接
```
守护进程daemon
使用多种网络管理工具容易产生各种问题，因此，请只用一种网络管理工具来管理网络连接。 所以，在使用wicd前，必须先关闭其他网络管理工具。
1、使用以下命令手动关闭network、dhcdbd和networkmanager这些守护程序。
# rc.d stop network
# rc.d stop dhcdbd
# rc.d stop networkmanager
2、以root身份编辑/etc/rc.conf，注释掉（或者直接删除）相关的网络设置。（!network，之前加一个!就可以不加载了。）
A、在INTERFACES的参数前加一个(!) ，禁止使用，如下：
INTERFACES=(!eth0 !wlan0) #这里的lo不再需要了，因为其是由 /etc/rc.sysinit启动的。
B、禁止使用的各种网络管理守护进程，包括network, dhcdbd, 和 networkmanager：
DAEMONS=(syslog-ng @alsa !network dbus !dhcdbd !networkmanager wicd netfs ...)
C、添加守护进程dbus 和wicd 到DAEMONS中，使得守护进程参数列表看起来像这样：
DAEMONS=(syslog-ng dbus !network !dhcdbd !networkmanager wicd ...)
Note: 如果你使用了 hal,把 dbus 替换成hal，守护进程hal会自动启动 dbus 。
D、使用如下命令把你帐号加入到network组中，把$USERNAME替换成你自己帐号名称。
gpasswd -a $USERNAME network
E、最后，启动守护进程 dbus 和wicd :
# rc.d start dbus
# rc.d start wicd
Note: 如果守护进程dbus已经启动，需要重新启动它:
# rc.d restart dbus
启动并使用wicd图形管理工具
方法有两种，一种是命令行启动，一种是图形界面启动。图形界面只需要鼠标点击，下面介绍命令行启动，都不需要root权限。 命令行输入:
$ wicd-client
如果你不需要wicd出现在通知区，使用下面命令：
$ wicd-client -n
你也可以把wicd-client添加到你所使用的DE/WM 自启动列表中，这样每次登录就能自动启动图形管理界面。 Note that Wicd doesn't prompt you for a passkey. To use encrypted connections (WPA/WEP), expand the network you want to connect to, click 'Advanced', and enter the relevant info.
```
# 常用软件
Arch安装成功后，网络功能正常，则首先下载wget或curl等下载工具，然后开始更新/安装软件包
# wget 
# 配制/etc/pacman.d/mirrorlist
可以将距离较近的源排在前面
 1. China
 2. Any
 3. ...
# 更新软件
```
#!sh
$ pacman -Syy
```
传入两个 --refresh 或 -y 标记强制 pacman 刷新所有软件包，不管是否被认为是最新的。只要镜像有变化就执行 pacman -Syy 是防止出现令人头疼问题的好习惯。
升级整个系统
```
$ pacman -Syu 
```
# 新建用户
```
$ adduser
... 根据提示来，需要注意添加分组时，一般输入以下分组: `audio,lp,optical,storage,video,wheel,games,power,scanner`
正如例子中所示，建议仅在Login name 和 Additional groups 输入内容，其它都留空。
Additional groups 中的列表是桌面系统的典型选择，特别推荐给新手：
audio - 让任务可以调用声卡以及相关软件
lp - 管理打印任务
optical - 管理光驱相关任务
storage - 管理存储设备
video - 视频任务以及硬件加速
wheel - 使用 sudo
games - 得到那些属于游戏组的权限，比如手柄
power - 笔记本用户需要这个
scanner - 使用扫描仪
接着会给出用户信息预览，可以取消或者继续。
```
```
#!div class=note
# 弄错用户重新来 ==
```
#!sh
$ userdel -r username
```
 * -r 会将所有用户数据清除
```
# sudo ===
也可以在安装arch时选择安装该软件包，这里就不用再安装了.
```
#!sh
$ visudo
```
# ===
```
#!sh
$ pacman -S zsh emacs gvim 
```
# 升级维护
# pacmatic
 * http://kmkeen.com/pacmatic/index.html
#  arch-wiki-lite
 * http://kmkeen.com/arch-wiki-lite/
# 参考
 * [http://pacnet.karbownicki.com/ Packnet]
# 时钟
# 硬件时钟和系统时钟
系统用两个时钟保存时间：硬件时钟和系统时钟。
 * 硬件时钟(即实时时钟 RTC 或 CMOS 时钟)仅能保存：年、月、日、时、分、秒这些时间数值，无法保存时间标准(UTC 或 localtime)和是否使用夏令时调节。
 * 系统时钟(即软件时间) 与硬件时间分别维护，保存了：时间、时区和夏令时设置。Linux 内核保存为自 UTC 时间 1970 年1月1日经过的秒数。初始系统时钟是从硬件时间计算得来，计算时会考虑/etc/adjtime的设置。系统启动之后，系统时钟与硬件时钟独立运行，Linux 通过时钟中断计数维护系统时钟。
```div class=note
因为系统时间是按 32 为整数保存的，最大只能记到 2038 年，所以 32 位 Linux 系统将在 2038 年停止工作。
```
# 获取时钟状态
```sh
 $ timedatectl status
      Local time: 日 2013-04-28 18:54:37 CST
  Universal time: 日 2013-04-28 10:54:37 UTC
        RTC time: 日 2013-04-28 10:54:37
        Timezone: Asia/Shanghai (CST, +0800)
     NTP enabled: no
NTP synchronized: yes
 RTC in local TZ: no
      DST active: n/a
```
# 设置系统时钟
```sh
$timedatectl set-time "2012-12-25 18:18:18"
```
# RTC clock
大部分操作系统的时间管理包括如下方面：
 1. 启动时根据硬件时钟设置系统时间
 2. 运行时通过 NTP 守护进程联网校正时间
 3. 关机时根据系统时间设置硬件时间。
# 时间标准
# UTC
# RTC
