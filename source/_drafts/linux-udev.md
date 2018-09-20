---
title: udev入门
tags:
---
<!-- toc -->
# udev入门
UDev的目的是在Linux2.6 kernel之上,提供一个建立在用户空间上的动态/dev目录的解决方案.

## 简介 
再典型的linux系统上, /dev目录用于存储类似于文件的DeviceNodes, DeviceNode代表系统上的某个设备(Device), 这里的设备可能是看得见膜的着的真实设备,也可能是虚拟设备. 
用户空间程序,可以使用DevicesNodes作为与系统硬件互通的接口(Interface). 比如:
XServer 会收听`/dev/input/mice`这个DeviceNode上事件, 以便控制光标跟随鼠标移动.
所以早先的/dev目录下收录了所有提供支持的设备节点,
 1. 一方面这些设备节点不是谁都能用的着的
 2. 另一方面这导致/dev目录太大


显然我们不能容忍这种占毛坑不拉屎的设计.于是DevFs项目开始着说解决这种浪费. 办法也很容易想见,即计算机上连接了哪些设备,就在/dev下创建这些设备的设备节点.但是DevFs的动机十分正确,但实现略有欠缺,这导致DevFs不那么靠谱,而且很难变的更靠谱.
UDev发明了一种新的管理/dev目录的办法, 它依赖于SysFs提供的信息,以及用户定义的规则.
SysFs是2.6Kernels提供的新文件系统.由内核管理,将已知设备信息,映射到这个文件系统上, 这样Udev可以根据这些信息在/dev目录下建立设备节点. SysFs挂在/sys目录下. 
这篇文档将介绍udev的一些新的变化。从084版开始，udev能够代替hotplug和coldplug的所有功能。正因为这样，hotplug包已经从Arch仓库中去掉了。

## 通过UDev 规则, 你能干嘛? 
 * 重命名设备节点
 * 可以为每个设备节点提供一个可以修改/固定的名字(通过符号连接)
 * 通过程序的输出结果为设备节点命名
 * 当设备节点建立或删除时,你可以调用自定义脚本
 * 重命名网络接口(NetworkInterface)

UDev添加设备节点之前,会先看看Udev规则,如果Udev规则中没有其他特别的要求,那就按照设备的默认名字创建设备节点. 
如果你的机器上压根没有某个设备,你就是写了针对这个设备UDev规则, 也没用.
为设备提供一个固定的唯一名字,有诸多便利之处. 比如,你有个U盘和Usb摄像头, 如果你只是使用默认设备节点, 那么就会依赖于插U盘和USB摄像头的顺序. 比如: 你可能先接U盘,UDev为你创建了/dev/sdb , 先接USB摄像头, UDev还是创建/dev/sdb, 都是/dev/sdb, 确有可能是不同的设备, 这可不好. 如果你使用UDev规则为这两个设备建立一个固定的名字, 比如/dev/udisk , /dev/camera 就一目了然.

你可以看看自己机器上的存储设备:
```sh
$ ls -a /dev/disk
by-id by-label by-path by-uuid
$ ls -lR /dev/disk
...
/dev/disk/by-label:
total 0
lrwxrwxrwx 1 root root 10 Jun 18 10:48 AMAS -> ../../sda8
lrwxrwxrwx 1 root root 10 Jun 18 10:48 bin -> ../../sdc5
lrwxrwxrwx 1 root root 10 Jun 18 10:48 FINGER -> ../../sda2
lrwxrwxrwx 1 root root 10 Jun 18 10:48 INSIGHT -> ../../sda9
lrwxrwxrwx 1 root root 10 Jun 18 10:48 Libray -> ../../sdc8
lrwxrwxrwx 1 root root 10 Jun 18 10:48 LinuxBoot -> ../../sdb1
lrwxrwxrwx 1 root root 10 Jun 18 10:48 LinuxData -> ../../sdb9
lrwxrwxrwx 1 root root 10 Jun 18 10:48 LinuxHome -> ../../sdb8
lrwxrwxrwx 1 root root 10 Jun 18 10:48 LinuxRestore -> ../../sdb5
lrwxrwxrwx 1 root root 10 Jun 18 10:48 LinuxRoot -> ../../sdb2
lrwxrwxrwx 1 root root 10 Jun 18 10:48 LinuxSwap -> ../../sdb3
lrwxrwxrwx 1 root root 10 Jun 18 10:48 LinuxTmp -> ../../sdb6
lrwxrwxrwx 1 root root 10 Jun 18 10:48 LinuxUsr -> ../../sdb7
lrwxrwxrwx 1 root root 11 Jun 18 10:48 LinuxVar -> ../../sdb10
lrwxrwxrwx 1 root root 11 Jun 18 10:48 Movi -> ../../sdc11
lrwxrwxrwx 1 root root 10 Jun 18 10:48 Music -> ../../sdc9
lrwxrwxrwx 1 root root 11 Jun 18 10:48 src -> ../../sdc10
lrwxrwxrwx 1 root root 10 Jun 18 10:48 SRC -> ../../sda7
lrwxrwxrwx 1 root root 10 Jun 18 10:48 SWAP-sda11 -> ../../sda6
lrwxrwxrwx 1 root root 10 Jun 18 10:48 WinXP -> ../../sdc1
lrwxrwxrwx 1 root root 10 Jun 18 10:48 /usr -> ../../sda5
...
```
## 如何书写UDev规则文件 
> /etc/udev/rules.d 
 * 文件名必须以`.rules`为后缀
 * 故则文件按照LexicalOrder进行解析,所以10_*.rules要比50_*.rules先被执行.
 * `'#'`为注释
 * 一个设备可以对应多条Udev规则,  UDev找到合适的规则后,并不停止,而是继续查找下一条匹配规则.
 
### 规则语法 
 * Key-Value构成,逗号分割
 * 一条规则中至少要包含一个MatchKey和一个AssignmentKey.

比如:
```
KERNEL'sdb', NAME="my_udisk"
```

| MATCH KEY  | 作用 |
|------------|------------------|
| KERNEL     | 匹配设备的KernelName
| SUBSYSTEM  | 匹配设备的Subsystem
| DRIVER   | 匹配设备的后台驱动名

| Assiggment Key | 作用 |
|------------|------------------|
| NAME       | 该设备节点的名字 
| SYMLINK    | 该设备节点的别名

 * UDev将使用NAME创建一个唯一的实际的设备节点
 * 设备节点的别名都是符号连接

现在回头看:
```
KERNEL'sdb', NAME="my_udisk"
```
这个规则是说当发现KernelName是sdb的时候,将这个设备的设备节点建立为`/dev/my_udisk`.


```
KERNEL'sdb', NAME="my_udisk", SYSMLIK+"amas_flash_disk"
```
除了上面说的,再建立一个/dev/amas_flash_disk的设备节点别名.

```
KERNEL"hdc", SYMLINK+="cdrom cdrom0"
```
以上是基本的用法.

### 匹配SysFs属性
还记得SysFs么?它是所有设备信息的映像, 为了正确识别设备,我们需要匹配其中的设备信息.
```
SUBSYSTEM"block", ATTR{size}"234441648", SYMLINK+="my_disk"
```
### 设备结构
Linux kernel将所有设备视为树形结构.(比之前的基本MatchKey多个`'S'`)

| MATCH KEY   | 作用 |
|-------------|------------------|
| KERNELS     | 匹配设备的kernelname, 或任何父设备的kernalname |
| SUBSYSTEMS  | 匹配设备Subsystem, 或任何父设备的Subsystem     |
| DRIVERS     | match against the name of the driver backing the device, or the name of the driver backing any of the parent devices |
| ATTRS       | match a sysfs attribute of the device, or a sysfs attribute of any of the parent devices |

### 字符串替换
| 操作符号 | 作用 |
|----------|---------------|
| %k       | 代表KernalName|
| %n       | 代表KernalNumber|
比如:
```
KERNEL"mice", NAME="input/%k"
KERNEL"loop0", NAME="loop/%n", SYMLINK+="%k"
```
The first rule ensures that the mice device node appears exclusively in the /dev/input directory (by default it would be at /dev/mice). The second rule ensures that the device node named loop0 is created at /dev/loop/0 but also creates a symbolic link at /dev/loop0 as usual.
The use of the above rules is questionable, as they all could be rewritten without using any substitution operators. The true power of these substitutions will become apparent in the next section.
## 字符串匹配 =

| *  | |
| ?  | |
| [] | |

```
KERNEL"fd[0-9]*", NAME="floppy/%n", SYMLINK+="%k"
KERNEL"hiddev*", NAME="usb/%k"
```
## 从SysFs中获得设备信息 =
SysFs下设备的路径的顶级目录都包含`dev`这个文件,所以从SysFs中搜索设备信息的方法是:
```
##!sh
$ find /sys -name 'dev'
```
在我的机器上, 第一块硬盘的信息映射在:
```
/sys/devices/pci0000:00/0000:00:1f.2/host0/target0:0:0/0:0:0:0/block/sda
```
你可以查看它的大小:
```
$ cat /sys/devices/pci0000:00/0000:00:1f.2/host0/target0:0:0/0:0:0:0/block/sda/size
625140335
```
你可以简单的使用这个硬盘的大小作为MatchKey:
```
ATTR{size}"625140335"
```
因为每个属性文件中都是该属性的值,所以一个设备的诸多属性分布在很多属性文件下,看起来比较困难,你可以使用`udevadm info -a -p <udev path>`来列出所有设备属性.
```sh
$ udevadm info -a -p /sys/devices/pci0000:00/0000:00:1f.2/host0/target0:0:0/0:0:0:0/block/sda
$ udevadm info -a -p /sys/block/sdb
Udevadm info starts with the device specified by the devpath and then
walks up the chain of parent devices. It prints for every device
found, all possible attributes in the udev rules key format.
A rule to match, can be composed by the attributes of the device
and the attributes from one single parent device.
  looking at device '/devices/pci0000:00/0000:00:1f.2/host0/target0:0:0/0:0:0:0/block/sda':
    KERNEL"sda"
    SUBSYSTEM"block"
    DRIVER""
    ATTR{range}"16"
    ATTR{ext_range}"256"
    ATTR{removable}"0"
    ATTR{ro}"0"
    ATTR{size}"625140335"
    ATTR{alignment_offset}"0"
    ATTR{discard_alignment}"0"
    ATTR{capability}"50"
    ATTR{stat}"    1860     5490     9466     8993        0        0        0        0        0     1856     8993"
    ATTR{inflight}"       0        0"
  looking at parent device '/devices/pci0000:00/0000:00:1f.2/host0/target0:0:0/0:0:0:0':
    KERNELS"0:0:0:0"
    SUBSYSTEMS"scsi"
    DRIVERS"sd"
    ATTRS{device_blocked}"0"
    ATTRS{type}"0"
    ATTRS{scsi_level}"6"
    ATTRS{vendor}"ATA     "
    ATTRS{model}"ST3320820AS     "
    ATTRS{rev}"3.AA"
    ATTRS{state}"running"
    ATTRS{timeout}"30"
    ATTRS{iocounterbits}"32"
    ATTRS{iorequest_cnt}"0x767"
    ATTRS{iodone_cnt}"0x767"
    ATTRS{ioerr_cnt}"0x2"
    ATTRS{modalias}"scsi:t-0x00"
    ATTRS{evt_media_change}"0"
    ATTRS{queue_depth}"1"
    ATTRS{queue_type}"none"
  looking at parent device '/devices/pci0000:00/0000:00:1f.2/host0/target0:0:0':
    KERNELS"target0:0:0"
    SUBSYSTEMS"scsi"
    DRIVERS""
  looking at parent device '/devices/pci0000:00/0000:00:1f.2/host0':
    KERNELS"host0"
    SUBSYSTEMS"scsi"
    DRIVERS""
  looking at parent device '/devices/pci0000:00/0000:00:1f.2':
    KERNELS"0000:00:1f.2"
    SUBSYSTEMS"pci"
    DRIVERS"ata_piix"
    ATTRS{vendor}"0x8086"
    ATTRS{device}"0x3a20"
    ATTRS{subsystem_vendor}"0x1043"
    ATTRS{subsystem_device}"0x82d4"
    ATTRS{class}"0x01018f"
    ATTRS{irq}"19"
    ATTRS{local_cpus}"ff"
    ATTRS{local_cpulist}"0-7"
    ATTRS{modalias}"pci:v00008086d00003A20sv00001043sd000082D4bc01sc01i8f"
    ATTRS{dma_mask_bits}"32"
    ATTRS{consistent_dma_mask_bits}"32"
    ATTRS{broken_parity_status}"0"
    ATTRS{msi_bus}""
  looking at parent device '/devices/pci0000:00':
    KERNELS"pci0000:00"
    SUBSYSTEMS""
    DRIVERS""
```
有了这些设备信息,你就可以从中选择出一部分来作为MatchKey.

## 权限控制
 - GROUP
 - OWNER
 - MODE

```
KERNEL"fb[0-9]*", NAME="fb/%n", SYMLINK+="%k", GROUP="video"
KERNEL"fd[0-9]*", OWNER="john"
KERNEL"inotify", NAME="misc/%k", SYMLINK+="%k", MODE="0666"
```

## 使用其他程序命名设备
```
KERNEL"hda", PROGRAM="/bin/device_namer %k", SYMLINK+="%c"
KERNEL"hda", PROGRAM="/bin/device_namer %k", NAME="%c{1}", SYMLINK+="%c{2}"
KERNEL"hda", PROGRAM="/bin/device_namer %k", NAME="%c{1}", SYMLINK+="%c{2+}"
KERNEL"hda", PROGRAM="/bin/who_owns_device %k", GROUP="%c"
```

## 根据特定的UDev事件触发外部程序
```
KERNEL"sdb", RUN+="/usr/bin/my_program"
#When /usr/bin/my_program is executed, various parts of the udev environment are available as environment variables, including key values such as SUBSYSTEM. You can also use the ACTION environment variable to detect whether the device is being connected or disconnected - ACTION will be either "add" or "remove" respectively.
```
## 内建设备名
UDev已经命名了一些类型的设备. 所以你不必什么都搞.

## 设置环境变量 
```
KERNEL"fd0", SYMLINK+="floppy", ENV{some_var}="value"
KERNEL"fd0", ENV{an_env_var}"yes", SYMLINK+="floppy"
```
> 重要提示 
切记，在使用udev加载任何modules（内核模块）之前（无论是否是启动时自动加载），您必须在/etc/rc.conf将MOD_AUTOLOAD选项设置为yes ，否则您必须手动加载这些modules。您可以修改rc.conf中的MODULES或者使用modprobe命令来手动加载您所需要的modules。另一种方法是用hwdetect --modules生成系统硬件的modules列表，然后将这个列表添加到rc.conf中让系统启动时自动加载这些modules。

## UDev 规则 
udev的规则保存在`/etc/udev/rules.d/`，其中的文件名要以.rules结尾。
```
##!sh
$ ls /etc/udev/rules.d
75-cd-aliases-generator.rules.optional
75-persistent-net-generator.rules.optional
```
## 相关工具 
## udevam
UDev管理工具.
获得设备的UDev路径:
```sh
$ udevadm info -q path -n [device name]
```

查询设备的属性列表:
```sh
$ devpath=$(udevadm info -q path -n [device name])
$ udevadm info -a -p $devpath
```
修改了UDev规则后,需要重新启动Udev以使其生效:
```
$ udevadm control --reload-rules
```
## 模块禁用列表 
udev也会犯错或加载错误的模块。可在`/etc/rc.conf`中进制加载特定模块.
```
MODULES=(!moduleA !moduleB)
```

## 监控udev事件
```sh
$ udevadm monitor
monitor will print the received events for:
UDEV - the event which udev sends out after rule processing
KERNEL - the kernel uevent
KERNEL[1332244611.326232] remove   /devices/pci0000:00/0000:00:1d.7/usb1/1-7/1-7:1.0/host18/target18:0:0/18:0:0:0/bsg/18:0:0:0 (bsg)
KERNEL[1332244611.326255] remove   /devices/pci0000:00/0000:00:1d.7/usb1/1-7/1-7:1.0/host18/target18:0:0/18:0:0:0/scsi_generic/sg3 (scsi_generic)
KERNEL[1332244611.326267] remove   /devices/pci0000:00/0000:00:1d.7/usb1/1-7/1-7:1.0/host18/target18:0:0/18:0:0:0/scsi_device/18:0:0:0 (scsi_device)
KERNEL[1332244611.326276] remove   /devices/pci0000:00/0000:00:1d.7/usb1/1-7/1-7:1.0/host18/target18:0:0/18:0:0:0/scsi_disk/18:0:0:0 (scsi_disk)
KERNEL[1332244611.326289] remove   /devices/pci0000:00/0000:00:1d.7/usb1/1-7/1-7:1.0/host18/target18:0:0/18:0:0:0/block/sdc/sdc4 (block)
UDEV  [1332244611.326851] remove   /devices/pci0000:00/0000:00:1d.7/usb1/1-7/1-7:1.0/host18/target18:0:0/18:0:0:0/scsi_device/18:0:0:0 (scsi_device)
UDEV  [1332244611.326986] remove   /devices/pci0000:00/0000:00:1d.7/usb1/1-7/1-7:1.0/host18/target18:0:0/18:0:0:0/bsg/18:0:0:0 (bsg)
UDEV  [1332244611.327122] remove   /devices/pci0000:00/0000:00:1d.7/usb1/1-7/1-7:1.0/host18/target18:0:0/18:0:0:0/scsi_disk/18:0:0:0 (scsi_disk)
UDEV  [1332244611.327136] remove   /devices/pci0000:00/0000:00:1d.7/usb1/1-7/1-7:1.0/host18/target18:0:0/18:0:0:0/scsi_generic/sg3 (scsi_generic)
UDEV  [1332244611.328845] remove   /devices/pci0000:00/0000:00:1d.7/usb1/1-7/1-7:1.0/host18/target18:0:0/18:0:0:0/block/sdc/sdc4 (block)
...
```
### 测试配置文件
```sh
$ udevadm test /dev/sdc
```
### 重新加载规则
```sh
$ udevadm control --reload-rules
```
### 查看设备信息
```sh
## udevinfo -a -p $(udevinfo -q path -n /dev/sda)
$ udevadm info -a -p $(udevadm info -q path -n /dev/sda)
```
### 其他 
 * [查看USB设备信息的GTK图形工具](http://www.kroah.com/linux/usb/ UsbView)
 * http://www.reactivated.net/writing_udev_rules.html Writing udev rules by Daniel Drake

