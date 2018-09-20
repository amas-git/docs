---
title: Rsync使用手册
tags:
---
<!-- toc -->
# Rsync
Rsync是一个小而强悍的同步工具,其核心算法非常有价值.
 * Rsync仅会传输文件的变更部分
 * Rsync支持压缩/加密

## rsyncd 
## 参数 
 * -P 
 * -a : archive mode,保护所有备份文件的文件属性/用户/组/连接等等 (类似于cp -a)
 * -e : -e ssh 使用ssh 
 * -r : 递归备份 
 * -z : 压缩传输
 * --delete : 同步源删除的文件要从备份中删除
 * -t pattern : 指定备份的文件名模式,比如备份所有java文件:`-t *.java`

## 配置文件: /etc/rsyncd.conf 
## 常见用法 
### 本地同步 
rsync可做为高级的cp单独使用,提供高效的拷贝功能.
```sh
# 同步src文件到dest文件
$ rsync -P  <src> <dest>
# -r/--recursive: 递归同步
$ rsync -Pr <src> <dest>
```

### 作为备份工具
用rsync作为备份工具,备份效率极高.
```sh
# 将/path/to/source目录以及该目录下的全部内容备份到/path/to/backup
$ rsync -avz --delete /path/to/source /path/to/backup
# 注意: 备份源目录后面多了一个'/', 意思是将/path/to/source目录下的所有内容备份到/path/to/backup
$ rsync -avz --delete /path/to/source/ /path/to/backup
```

### 远程备份:
```sh
$ rsync -avz --delete -e ssh /path/to/source  remoteuser@remotehost:/path/to/backup
```
过滤器
``` sh
$ rsync -avz --delete-excluded -exclude-from=backup.lst <source> <backup>
```

backup.list:
```
# 需要非分的内容
+ /dev/console
+ /dev/initctl
+ /dev/null
+ /dev/zero
# 无需备份的内容
- /dev/*
- /proc/*
- /sys/*
- /tmp/*
- lost+found/
- /media/backup/*
```

### 增量备份:
```
Since making a full copy of a large filesystem can be a time-consuming and expensive process, 
it is common to make full backups only once a week or once a month, and store only changes on
the other days. These are called "incremental" backups, and are supported by the venerable old 
dump and tar utilities, along with many others.
```

# 参考 
 * http://www.mikerubel.org/computers/rsync_snapshots/
 * http://rsync.samba.org/examples.html
 * http://troy.jdmz.net/rsync/index.html

