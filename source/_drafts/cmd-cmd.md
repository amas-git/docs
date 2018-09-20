---
title: 各种命令行工具 
tags:
---
<!-- toc -->
##

## systemd
```
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

```

## nmap: 端口扫描
> 注意: 尽量不要使用Nmap扫描别人的机器。

### 分析本机开启了哪些网络服务
```bash
$ sudo nmap -sTU localhost
Starting Nmap 5.51 ( http://nmap.org ) at 2011-05-20 11:14 CST
Nmap scan report for localhost (127.0.0.1)
Host is up (0.00023s latency).
rDNS record for 127.0.0.1: localhost.localdomain
Not shown: 1995 closed ports
PORT     STATE SERVICE
22/tcp   open  ssh
80/tcp   open  http
3306/tcp open  mysql
6000/tcp open  X11
8600/tcp open  asterix
Nmap done: 1 IP address (1 host up) scanned in 0.18 seconds
```

## openssh
### 修改sshd登录端口
```sh
# 1. vi /etc/service, 22为默认端口，修改即可
ssh tcp/22
ssh udp/22
ssh stcp/22
# 2. vi /etc/ssh/sshd_config
Port 2622
```
### 在远程主机上执行命令
```
ssh username@host “remote_command” | “local_command”
```
比如:
```
# 列出192.168.1.100上的/etc目录
$ ssh amas@192.168.1.100 ls /etc
# 拉目录
$ ssh user@host “tar cf - /path” | tar –xf –
# 推文件方法1
$ scp localfile user@host:/path/to
# 推文件方法2
$ cat file | ssh user@host “cat > file”
# 推目录
$ tar –cf - /path | ssh user@host“tar –xf –”
```

## ptrace
```cpp
       #include <sys/ptrace.h>
       long ptrace(enum __ptrace_request request, pid_t pid, void *addr, void *data);
```
 * ptrace可以追踪并控制指定进程的执行过程
 * ptrace可以查看或者修改tracee进程的寄存器, VMA
 * 在内存中权限问题被忽略
由于ptrace过于强大，自然有人会打它的主意，拿ptrace干些偷鸡摸狗的事情。下面我们来分析一下。
### Fucking injectso是如何工作的?
常见的一种手段就是向目标进程空间注入DSO, 其实用mmap系统调用可以干这个事儿，但是这样简单粗暴的内存映射之后并不能立刻工作，因为DSO中的诸多对象需要重新计算地址。这本来是动态链接器的工作，可繁琐了。所以咱们还是用dlopen加载一个DSO进来更加方便一些。大概的思路如下:
 * 用ptrace粘住目标进程
 * 保存当前栈信息
 * 保存当前寄存器集合
 * 伪造函数调用的栈帧，返回值为0
 * 修改寄存器，准备调用dlopen
   * rip : &dlopen / dlopen函数的地址
   * rdi : DSO的名称, 注意不是PATH啊，所以可能会用LD_LIBARY_PATH指向你注入so的路径
   * rsi :  DSO打开模式，参看dlopen函数的原型及说明
 * let er rip, waitpid and it’ll segfault on returnto 0.
 * 恢复栈，寄存器集合，目标进程可以继续跑啦
从以上过程中可以看出，大概思路就是暂停目标进程的执行，保存上下文信息，伪造dlopen调用，完成注入，恢复上下文，让目标进程继续。
### Anti ptrace
```cpp
int stayalive ;
void trapcatch (int  ii) {
    stayalive = 1;
}
int main(void) {
...
stayalive = 1;
signal (SIGTRAP, trapcatch ) ;
while ( stayalive ) {
    stayalive = 0;
    kill(getpid(), SIGTRAP ) ;
    do_the_work() ;
}
```
### 参考
 * https://github.com/ice799/injectso64
 * http://mips42.altervista.org/ptrace.php

## nm: 查看elf文件
```sh
$ nm -g your-c-lib.so
$ nm -gC your-cpp-lib.so
```
如果.so库为Elf格式， 请使用[wiki:ElfUtils]
```sh
$ file x.so
x.so: ELF 32-bit LSB shared object, ARM, version 1 (SYSV), dynamically linked, stripped
$ readelf -Ws x.so
# 如果你只想看到符号名, 使用awk过滤下
$ readelf -Ws /usr/lib/libstdc++.so.6 | awk '{print $8}'
```


## rsync
## sed & awk
```
#!sh
# 将所有/data/data下的目录保存为数组
packages=($(adb shell ls /data/data | sed -e 's///g')) 
```
SEE:[wiki:CRLF]
```
文本间隔：
# 在每一行后面增加一空行
sed G
awk '{printf("%s

",$0)}'
# 将原来的所有空行删除并在每一行后面增加一空行。
# 这样在输出的文本中每一行后面将有且只有一空行。
sed '/^$/d;G'
awk '!/^$/{printf("%s

",$0)}'
# 在每一行后面增加两行空行
sed 'G;G'
awk '{printf("%s


",$0)}'
# 将第一个脚本所产生的所有空行删除（即删除所有偶数行）
sed 'n;d'
awk '{f=!f;if(f)print $0}'
# 在匹配式样“regex”的行之前插入一空行
sed '/regex/{x;p;x;}'
awk '{if(/regex/)printf("
%s
",$0);else print $0}'
# 在匹配式样“regex”的行之后插入一空行
sed '/regex/G'
awk '{if(/regex/)printf("%s

",$0);else print $0}'
# 在匹配式样“regex”的行之前和之后各插入一空行
sed '/regex/{x;p;x;G;}'
awk '{if(/regex/)printf("
%s

",$0);else print $0}'
编号：
# 为文件中的每一行进行编号（简单的左对齐方式）。这里使用了“制表符”
# （tab，见本文末尾关于’	’的用法的描述）而不是空格来对齐边缘。
sed = filename | sed 'N;s/
/	/'
awk '{i++;printf("%d	%s
",i,$0)}'
# 对文件中的所有行编号（行号在左，文字右端对齐）。
sed = filename | sed 'N; s/^/     /; s/ *(.{6,})
/  /'
awk '{i++;printf("%6d  %s
",i,$0)}'
# 对文件中的所有行编号，但只显示非空白行的行号。
sed '/./=' filename | sed '/./N; s/
/ /'
awk '{i++;if(!/^$/)printf("%d %s
",i,$0);else print}'
# 计算行数 （模拟 “wc -l”）
sed -n '$='
awk '{i++}END{print i}'
文本转换和替代：
# Unix环境：转换DOS的新行符（CR/LF）为Unix格式。
sed 's/.$//'                     # 假设所有行以CR/LF结束
sed 's/^M$//'                    # 在bash/tcsh中，将按Ctrl-M改为按Ctrl-V
sed 's/$//'                  # ssed、gsed 3.02.80，及更高版本
awk '{sub(/$/,"");print $0}'
# Unix环境：转换Unix的新行符（LF）为DOS格式。
sed "s/$/`echo -e \`/"        # 在ksh下所使用的命令
sed 's/$'"/`echo \`/"         # 在bash下所使用的命令
sed "s/$/`echo \`/"           # 在zsh下所使用的命令
sed 's/$//'                    # gsed 3.02.80 及更高版本
awk '{printf("%s
",$0)}'
# DOS环境：转换Unix新行符（LF）为DOS格式。
sed "s/$//"                      # 方法 1
sed -n p                         # 方法 2
DOS环境的略过
# DOS环境：转换DOS新行符（CR/LF）为Unix格式。
# 下面的脚本只对UnxUtils sed 4.0.7 及更高版本有效。要识别UnxUtils版本的
# sed可以通过其特有的“–text”选项。你可以使用帮助选项（“–help”）看
# 其中有无一个“–text”项以此来判断所使用的是否是UnxUtils版本。其它DOS
# 版本的的sed则无法进行这一转换。但可以用“tr”来实现这一转换。
sed "s///" infile >outfile     # UnxUtils sed v4.0.7 或更高版本
tr -d  <infile >outfile        # GNU tr 1.22 或更高版本
DOS环境的略过
# 将每一行前导的“空白字符”（空格，制表符）删除
# 使之左对齐
sed 's/^[ 	]*//'                # 见本文末尾关于'	'用法的描述
awk '{sub(/^[ 	]+/,"");print $0}'
# 将每一行拖尾的“空白字符”（空格，制表符）删除
sed 's/[ 	]*$//'                # 见本文末尾关于'	'用法的描述
awk '{sub(/[ 	]+$/,"");print $0}'
# 将每一行中的前导和拖尾的空白字符删除
sed 's/^[ 	]*//;s/[ 	]*$//'
awk '{sub(/^[ 	]+/,"");sub(/[ 	]+$/,"");print $0}'
# 在每一行开头处插入5个空格（使全文向右移动5个字符的位置）
sed 's/^/     /'
awk '{printf("     %s
",$0)}'
# 以79个字符为宽度，将所有文本右对齐
# 78个字符外加最后的一个空格
sed -e :a -e 's/^.{1,78}$/ &/;ta'
awk '{printf("%79s
",$0)}'
# 以79个字符为宽度，使所有文本居中。在方法1中，为了让文本居中每一行的前
# 头和后头都填充了空格。 在方法2中，在居中文本的过程中只在文本的前面填充
# 空格，并且最终这些空格将有一半会被删除。此外每一行的后头并未填充空格。
sed  -e :a -e 's/^.{1,77}$/ & /;ta'                     # 方法1
sed  -e :a -e 's/^.{1,77}$/ &/;ta' -e 's/( *)//'  # 方法2
awk '{for(i=0;i<39-length($0)/2;i++)printf(" ");printf("%s
",$0)}'  #相当于上面的方法二
# 在每一行中查找字串“foo”，并将找到的“foo”替换为“bar”
sed 's/foo/bar/'                 # 只替换每一行中的第一个“foo”字串
sed 's/foo/bar/4'                # 只替换每一行中的第四个“foo”字串
sed 's/foo/bar/g'                # 将每一行中的所有“foo”都换成“bar”
sed 's/(.*)foo(.*foo)/bar/' # 替换倒数第二个“foo”
sed 's/(.*)foo/bar/'            # 替换最后一个“foo”
awk '{gsub(/foo/,"bar");print $0}'   # 将每一行中的所有“foo”都换成“bar”
# 只在行中出现字串“baz”的情况下将“foo”替换成“bar”
sed '/baz/s/foo/bar/g'
awk '{if(/baz/)gsub(/foo/,"bar");print $0}'
# 将“foo”替换成“bar”，并且只在行中未出现字串“baz”的情况下替换
sed '/baz/!s/foo/bar/g'
awk '{if(/baz$/)gsub(/foo/,"bar");print $0}'
# 不管是“scarlet”“ruby”还是“puce”，一律换成“red”
sed 's/scarlet/red/g;s/ruby/red/g;s/puce/red/g'  #对多数的sed都有效
gsed 's/scarlet|ruby|puce/red/g'               # 只对GNU sed有效
awk '{gsub(/scarlet|ruby|puce/,"red");print $0}'
# 倒置所有行，第一行成为最后一行，依次类推（模拟“tac”）。
# 由于某些原因，使用下面命令时HHsed v1.5会将文件中的空行删除
sed '1!G;h;$!d'               # 方法1
sed -n '1!G;h;$p'             # 方法2
awk '{A[i++]=$0}END{for(j=i-1;j>=0;j--)print A[j]}'
# 将行中的字符逆序排列，第一个字成为最后一字，……（模拟“rev”）
sed '/
/!G;s/(.)(.*
)/&/;//D;s/.//'
awk '{for(i=length($0);i>0;i--)printf("%s",substr($0,i,1));printf("
")}'
# 将每两行连接成一行（类似“paste”）
sed '$!N;s/
/ /'
awk '{f=!f;if(f)printf("%s",$0);else printf(" %s
",$0)}'
# 如果当前行以反斜杠“”结束，则将下一行并到当前行末尾
# 并去掉原来行尾的反斜杠
sed -e :a -e '/\$/N; s/\
//; ta'
awk '{if(/\$/)printf("%s",substr($0,0,length($0)-1));else printf("%s
",$0)}'
# 如果当前行以等号开头，将当前行并到上一行末尾
# 并以单个空格代替原来行头的“=”
sed -e :a -e '$!N;s/
=/ /;ta' -e 'P;D'
awk '{if(/^=/)printf(" %s",substr($0,2));else printf("%s%s",a,$0);a="
"}END{printf("
")}'
# 为数字字串增加逗号分隔符号，将“1234567”改为“1,234,567”
gsed ':a;s/B[0-9]{3}>/,&/;ta'                     # GNU sed
sed -e :a -e 's/(.*[0-9])([0-9]{3})/,/;ta'  # 其他sed
#awk的正则没有后向匹配和引用，搞的比较狼狈，呵呵。
awk '{while(match($0,/[0-9][0-9][0-9][0-9]+/)){$0=sprintf("%s,%s",substr($0,0,RSTART+RLENGTH-4),substr($0,RSTART+RLENGTH-3))}print $0}'
# 为带有小数点和负号的数值增加逗号分隔符（GNU sed）
gsed -r ':a;s/(^|[^0-9.])([0-9]+)([0-9]{3})/,/g;ta'
#和上例差不多
awk '{while(match($0,/[^.0-9][0-9][0-9][0-9][0-9]+/)){$0=sprintf("%s,%s",substr($0,0,RSTART+RLENGTH-4),substr($0,RSTART+RLENGTH-3))}print $0}'
# 在每5行后增加一空白行 （在第5，10，15，20，等行后增加一空白行）
gsed '0~5G'                      # 只对GNU sed有效
sed 'n;n;n;n;G;'                 # 其他sed
awk '{print $0;i++;if(i==5){printf("
");i=0}}'
选择性地显示特定行：
# 显示文件中的前10行 （模拟“head”的行为）
sed 10q
awk '{print;if(NR==10)exit}'
# 显示文件中的第一行 （模拟“head -1”命令）
sed q
awk '{print;exit}'
# 显示文件中的最后10行 （模拟“tail”）
sed -e :a -e '$q;N;11,$D;ba'
#用awk干这个有点亏，得全文缓存，对于大文件肯定很慢
awk '{A[NR]=$0}END{for(i=NR-9;i<=NR;i++)print A[i]}'
# 显示文件中的最后2行（模拟“tail -2”命令）
sed '$!N;$!D'
awk '{A[NR]=$0}END{for(i=NR-1;i<=NR;i++)print A[i]}'
# 显示文件中的最后一行（模拟“tail -1”）
sed '$!d'                        # 方法1
sed -n '$p'                      # 方法2
#这个比较好办，只存最后一行了。
awk '{A=$0}END{print A}'
# 显示文件中的倒数第二行
sed -e '$!{h;d;}' -e x              # 当文件中只有一行时，输出空行
sed -e '1{$q;}' -e '$!{h;d;}' -e x  # 当文件中只有一行时，显示该行
sed -e '1{$d;}' -e '$!{h;d;}' -e x  # 当文件中只有一行时，不输出
#存两行呗（当文件中只有一行时，输出空行）
awk '{B=A;A=$0}END{print B}'
# 只显示匹配正则表达式的行（模拟“grep”）
sed -n '/regexp/p'               # 方法1
sed '/regexp/!d'                 # 方法2
awk '/regexp/{print}'
# 只显示“不”匹配正则表达式的行（模拟“grep -v”）
sed -n '/regexp/!p'              # 方法1，与前面的命令相对应
sed '/regexp/d'                  # 方法2，类似的语法
awk '!/regexp/{print}'
# 查找“regexp”并将匹配行的上一行显示出来，但并不显示匹配行
sed -n '/regexp/{g;1!p;};h'
awk '/regexp/{print A}{A=$0}'
# 查找“regexp”并将匹配行的下一行显示出来，但并不显示匹配行
sed -n '/regexp/{n;p;}'
awk '{if(A)print;A=0}/regexp/{A=1}'
# 显示包含“regexp”的行及其前后行，并在第一行之前加上“regexp”所在行的行号 （类似“grep -A1 -B1”）
sed -n -e '/regexp/{=;x;1!p;g;$!N;p;D;}' -e h
awk '{if(F)print;F=0}/regexp/{print NR;print b;print;F=1}{b=$0}'
# 显示包含“AAA”、“BBB”和“CCC”的行（任意次序）
sed '/AAA/!d; /BBB/!d; /CCC/!d'   # 字串的次序不影响结果
awk '{if(match($0,/AAA/) && match($0,/BBB/) && match($0,/CCC/))print}'
# 显示包含“AAA”、“BBB”和“CCC”的行（固定次序）
sed '/AAA.*BBB.*CCC/!d'
awk '{if(match($0,/AAA.*BBB.*CCC/))print}'
# 显示包含“AAA”“BBB”或“CCC”的行 （模拟“egrep”）
sed -e '/AAA/b' -e '/BBB/b' -e '/CCC/b' -e d    # 多数sed
gsed '/AAA|BBB|CCC/!d'                        # 对GNU sed有效
awk '/AAA/{print;next}/BBB/{print;next}/CCC/{print}'
awk '/AAA|BBB|CCC/{print}'
# 显示包含“AAA”的段落 （段落间以空行分隔）
# HHsed v1.5 必须在“x;”后加入“G;”，接下来的3个脚本都是这样
sed -e '/./{H;$!d;}' -e 'x;/AAA/!d;'
awk 'BEGIN{RS=""}/AAA/{print}'
awk -vRS= '/AAA/{print}'
# 显示包含“AAA”“BBB”和“CCC”三个字串的段落 （任意次序）
sed -e '/./{H;$!d;}' -e 'x;/AAA/!d;/BBB/!d;/CCC/!d'
awk -vRS= '{if(match($0,/AAA/) && match($0,/BBB/) && match($0,/CCC/))print}'
# 显示包含“AAA”、“BBB”、“CCC”三者中任一字串的段落 （任意次序）
sed -e '/./{H;$!d;}' -e 'x;/AAA/b' -e '/BBB/b' -e '/CCC/b' -e d
gsed '/./{H;$!d;};x;/AAA|BBB|CCC/b;d'         # 只对GNU sed有效
awk -vRS= '/AAA|BBB|CCC/{print "";print}'
# 显示包含65个或以上字符的行
sed -n '/^.{65}/p'
cat ll.txt | awk '{if(length($0)>=65)print}'
# 显示包含65个以下字符的行
sed -n '/^.{65}/!p'            # 方法1，与上面的脚本相对应
sed '/^.{65}/d'                # 方法2，更简便一点的方法
awk '{if(length($0)<=65)print}'
# 显示部分文本——从包含正则表达式的行开始到最后一行结束
sed -n '/regexp/,$p'
awk '/regexp/{F=1}{if(F)print}'
# 显示部分文本——指定行号范围（从第8至第12行，含8和12行）
sed -n '8,12p'                   # 方法1
sed '8,12!d'                     # 方法2
awk '{if(NR>=8 && NR<12)print}'
# 显示第52行
sed -n '52p'                     # 方法1
sed '52!d'                       # 方法2
sed '52q;d'                      # 方法3, 处理大文件时更有效率
awk '{if(NR==52){print;exit}}'
# 从第3行开始，每7行显示一次
gsed -n '3~7p'                   # 只对GNU sed有效
sed -n '3,${p;n;n;n;n;n;n;}'     # 其他sed
awk '{if(NR==3)F=1}{if(F){i++;if(i%7==1)print}}'
# 显示两个正则表达式之间的文本（包含）
sed -n '/Iowa/,/Montana/p'       # 区分大小写方式
awk '/Iowa/{F=1}{if(F)print}/Montana/{F=0}'
选择性地删除特定行：
# 显示通篇文档，除了两个正则表达式之间的内容
sed '/Iowa/,/Montana/d'
awk '/Iowa/{F=1}{if(!F)print}/Montana/{F=0}'
# 删除文件中相邻的重复行（模拟“uniq”）
# 只保留重复行中的第一行，其他行删除
sed '$!N; /^(.*)
$/!P; D'
awk '{if($0!=B)print;B=$0}'
# 删除文件中的重复行，不管有无相邻。注意hold space所能支持的缓存大小，或者使用GNU sed。
sed -n 'G; s/
/&&/; /^([ -~]*
).*
/d; s/
//; h; P'  #bones7456注：我这里此命令并不能正常工作
awk '{if(!($0 in B))print;B[$0]=1}'
# 删除除重复行外的所有行（模拟“uniq -d”）
sed '$!N; s/^(.*)
$//; t; D'
awk '{if($0==B && $0!=l){print;l=$0}B=$0}'
# 删除文件中开头的10行
sed '1,10d'
awk '{if(NR>10)print}'
# 删除文件中的最后一行
sed '$d'
#awk在过程中并不知道文件一共有几行，所以只能通篇缓存，大文件可能不适合，下面两个也一样
awk '{B[NR]=$0}END{for(i=0;i<=NR-1;i++)print B[i]}'
# 删除文件中的最后两行
sed 'N;$!P;$!D;$d'
awk '{B[NR]=$0}END{for(i=0;i<=NR-2;i++)print B[i]}'
# 删除文件中的最后10行
sed -e :a -e '$d;N;2,10ba' -e 'P;D'   # 方法1
sed -n -e :a -e '1,10!{P;N;D;};N;ba'  # 方法2
awk '{B[NR]=$0}END{for(i=0;i<=NR-10;i++)print B[i]}'
# 删除8的倍数行
gsed '0~8d'                           # 只对GNU sed有效
sed 'n;n;n;n;n;n;n;d;'                # 其他sed
awk '{if(NR%8!=0)print}' |head
# 删除匹配式样的行
sed '/pattern/d'                      # 删除含pattern的行。当然pattern可以换成任何有效的正则表达式
awk '{if(!match($0,/pattern/))print}'
# 删除文件中的所有空行（与“grep ‘.’ ”效果相同）
sed '/^$/d'                           # 方法1
sed '/./!d'                           # 方法2
awk '{if(!match($0,/^$/))print}'
# 只保留多个相邻空行的第一行。并且删除文件顶部和尾部的空行。
# （模拟“cat -s”）
sed '/./,/^$/!d'        #方法1，删除文件顶部的空行，允许尾部保留一空行
sed '/^$/N;/
$/D'      #方法2，允许顶部保留一空行，尾部不留空行
awk '{if(!match($0,/^$/)){print;F=1}else{if(F)print;F=0}}'  #同上面的方法2
# 只保留多个相邻空行的前两行。
sed '/^$/N;/
$/N;//D'
awk '{if(!match($0,/^$/)){print;F=0}else{if(F<2)print;F++}}'
# 删除文件顶部的所有空行
sed '/./,$!d'
awk '{if(F || !match($0,/^$/)){print;F=1}}'
# 删除文件尾部的所有空行
sed -e :a -e '/^
*$/{$d;N;ba' -e '}'  # 对所有sed有效
sed -e :a -e '/^
*$/N;/
$/ba'        # 同上，但只对 gsed 3.02.*有效
awk '/^.+$/{for(i=l;i<NR-1;i++)print "";print;l=NR}'
# 删除每个段落的最后一行
sed -n '/^$/{p;h;};/./{x;/./p;}'
#很长，很ugly，应该有更好的办法
awk -vRS= '{B=$0;l=0;f=1;while(match(B,/
/)>0){print substr(B,l,RSTART-l-f);l=RSTART;sub(/
/,"",B);f=0};print ""}'
特殊应用：
# 移除手册页（man page）中的nroff标记。在Unix System V或bash shell下使
# 用’echo’命令时可能需要加上 -e 选项。
sed "s/.`echo \`//g"    # 外层的双括号是必须的（Unix环境）
sed 's/.^H//g'             # 在bash或tcsh中, 按 Ctrl-V 再按 Ctrl-H
sed 's/.//g'           # sed 1.5，GNU sed，ssed所使用的十六进制的表示方法
awk '{gsub(/./,"",$0);print}'
# 提取新闻组或 e-mail 的邮件头
sed '/^$/q'                # 删除第一行空行后的所有内容
awk '{print}/^$/{exit}'
# 提取新闻组或 e-mail 的正文部分
sed '1,/^$/d'              # 删除第一行空行之前的所有内容
awk '{if(F)print}/^$/{F=1}'
# 从邮件头提取“Subject”（标题栏字段），并移除开头的“Subject:”字样
sed '/^Subject: */!d; s///;q'
awk '/^Subject:.*/{print substr($0,10)}/^$/{exit}'
# 从邮件头获得回复地址
sed '/^Reply-To:/q; /^From:/h; /./d;g;q'
#好像是输出第一个Reply-To:开头的行？From是干啥用的？不清楚规则。。
awk '/^Reply-To:.*/{print;exit}/^$/{exit}'
# 获取邮件地址。在上一个脚本所产生的那一行邮件头的基础上进一步的将非电邮地址的部分剃除。（见上一脚本）
sed 's/ *(.*)//; s/>.*//; s/.*[:<] *//'
#取尖括号里的东西吧？
awk -F'[<>]+' '{print $2}'
# 在每一行开头加上一个尖括号和空格（引用信息）
sed 's/^/> /'
awk '{print "> " $0}'
# 将每一行开头处的尖括号和空格删除（解除引用）
sed 's/^> //'
awk '/^> /{print substr($0,3)}'
# 移除大部分的HTML标签（包括跨行标签）
sed -e :a -e 's/<[^>]*>//g;/</N;//ba'
awk '{gsub(/<[^>]*>/,"",$0);print}'
# 将分成多卷的uuencode文件解码。移除文件头信息，只保留uuencode编码部分。
# 文件必须以特定顺序传给sed。下面第一种版本的脚本可以直接在命令行下输入；
# 第二种版本则可以放入一个带执行权限的shell脚本中。（由Rahul Dhesi的一
# 个脚本修改而来。）
sed '/^end/,/^begin/d' file1 file2 ... fileX | uudecode   # vers. 1
sed '/^end/,/^begin/d' "$@" | uudecode                    # vers. 2
#我不想装个uudecode验证，大致写个吧
awk '/^end/{F=0}{if(F)print}/^begin/{F=1}' file1 file2 ... fileX
# 将文件中的段落以字母顺序排序。段落间以（一行或多行）空行分隔。GNU sed使用
# 字元“”来表示垂直制表符，这里用它来作为换行符的占位符——当然你也可以
# 用其他未在文件中使用的字符来代替它。
sed '/./{H;d;};x;s/
/={NL}=/g' file | sort | sed '1s/={NL}=//;s/={NL}=/
/g'
gsed '/./{H;d};x;y/
//' file | sort | sed '1s///;y//
/'
awk -vRS= '{gsub(/
/,"",$0);print}' ll.txt | sort | awk '{gsub(//,"
",$0);print;print ""}'
# 分别压缩每个.TXT文件，压缩后删除原来的文件并将压缩后的.ZIP文件
# 命名为与原来相同的名字（只是扩展名不同）。（DOS环境：“dir /b”
# 显示不带路径的文件名）。
echo @echo off >zipup.bat
dir /b *.txt | sed "s/^(.*).TXT/pkzip -mo  .TXT/" >>zipup.bat
DOS 环境再次略过，而且我觉得这里用 bash 的参数 ${i%.TXT}.zip 替换更帅。
下面的一些 SED 说明略过，需要的朋友自行查看原文。
```

## ssh-copy-id: ssh免密码登录
如果你想通过ssh登录192.168.100.2这台服务器，纯手工的做法是将自己的公钥(.pub)文件拷贝到服务器上blablabla...
然而有了ssh-copy-id你将一条命令搞定此事
```
$ ssh-copy-id <user>@host
```

## systemtap
在规划服务器的内存使用的时候经常需要知道应用在理想情况下会使用多少的pagecache, 我们好预先把这个内存预留出来.
这个值操作系统没有提供可查看的管道,我们只能自己写个脚本来实现.
下面的systemtap脚本每隔N秒显示下当前os下头10个文件占用多少的pagecache, 降序排列.
```
view sourceprint?
$ cat > pagecache.stp
global __filenames
global pagecache
 
probe vfs.add_to_page_cache
{
  pagecache[ino]++;
}
 
probe vfs.remove_from_page_cache
{
  pagecache[ino]--;
}
 
probe generic.fop.open
{
__filenames[ino]=filename
}
 
function find_filename(ino)
{
  if (ino in __filenames)
    return __filenames[ino];
  else return sprintf("N/A ino:%d", ino);
}
 
probe timer.s($1)
{
  ansi_clear_screen();
  printf ("%50s %10s
", "FILENAME", "COUNT")
 foreach( pages = ino in pagecache- limit 10)
 {
   if(pages)
   printf("%50s %10d
", find_filename(ino), pages);
 }
 
}
 
CTRL+D
$ sudo sysctl vm.drop_caches=3
 
$ sudo stap pagecache.stp  1
                                          FILENAME      COUNT
                                   librpmdb-4.4.so        173
                               libpython2.4.so.1.0        153
                                 libxml2.so.2.6.26        107
                                  N/A ino:68781310        100
                                     __m2crypto.so         91
                           libglib-2.0.so.0.1200.3         64
                                        libperl.so         53
                                     librpm-4.4.so         52
                                        pyexpat.so         45
                                libreadline.so.5.1         38
 
#拷贝个文件看看pagecache的变化
$ dd if=/dev/zero of=test.dat count=1024 bs=4096
```

## taskset: 让你的程序只使用单个CPU
```sh
# Start a command on only one CPU core
# This is useful if you have a program which doesn't work well with multicore CPUs.
# With taskset you can set its CPU affinity to run on only one core.
$ taskset -c 0 command
```

## telnet
### 使用TELNET手工操作 IMAP 查看邮件 =
IMAP 协议收信与POP收信有很大的不同，最明显的一点就是发送的每条命令(命令是不区分大小写的)，前面都要带有一个标签/标志，发送一条命令后可以紧接着发送另一条[[BR]]命令，服务器端返回命令处理结果的顺序是未知的，取决于各条命令的执行时间。所以返回的结果中，将带有所发送命令的标签。如下面示例中的 A01, A02 等等。大致上[[BR]]分为几步：登录服务器，验证用户名和密码，列出服务器上所有box，进入到某一个box，发送相应的命令（下载邮件，同步状态等等），下面看具体用法吧：
登录到服务器 :
```
#!sh
$ telnet imap.aol.com 143                                                                                           
Trying 64.12.143.164...
Connected to nginx.gnginx.aol.com.
Escape character is '^]'.
* OK IMAP4 ready
```
验证用户名和密码 :
```
#!sh
$ A01 login patest000@aol.com chimei   ＃login为命令，patest000＠aol.com为用户名，chimei为密码，中间用空格隔开
A01 OK LOGIN completed                 ＃ 服务器返回登录成功                                                     
```
列出服务器上所有的mail box :
```
#!sh
$ A02 list "" *                        ＃列出服务器上所有的信箱
* LIST (HasNoChildren Noinferiors) "/" INBOX
* LIST (HasNoChildren Noinferiors) "/" "Sent Items"
* LIST (HasNoChildren Noinferiors) "/" VOICEMAIL
* LIST (HasNoChildren Noinferiors) "/" Spam
* LIST (HasChildren) "/" Saved
* LIST (HasNoChildren Noinferiors) "/" Drafts
* LIST (HasNoChildren Noinferiors) "/" Saved/SavedIMs
* LIST (HasNoChildren Noinferiors) "/" Saved/testsstst
A02 OK LIST completed                                                     
```
选择收件箱 :
```
#!sh
$ A03 select inbox             ＃选择收件箱
* 246 EXISTS
* OK [UIDVALIDITY 1] UID validity status
* OK [UIDNEXT 25659948] predicted next UID
* 0 RECENT
* FLAGS (Seen Deleted XAOL-READ XAOL-GOODCHECK-DONE XAOL-CLIENT-BULK XAOL-RECEIVED XAOL-SENT $Forwarded XAOL-GOOD Draft Flagged Answered $MDNSent XAOL-VOICEMAIL XAOL-BULK XAOL-PRIORITY-MAIL XAOL-CERTIFIED-MAIL XAOL-BILLPAY-MAIL XAOL-OFFICIAL-MAIL XAOL-VIRUS-REPAIRED XAOL-VIRUS-NOT-SCANNED)
* OK [PERMANENTFLAGS (Seen Deleted XAOL-GOODCHECK-DONE XAOL-CLIENT-BULK XAOL-RECEIVED $Forwarded XAOL-GOOD Draft Flagged Answered $MDNSent)] Permanent flags
A03 OK [READ-WRITE] SELECT completed                                                   
```
查询收件箱所有邮件 
```
#!sh
$ A04 search all                       ＃ 查询收件箱所有邮件
* SEARCH 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30
* SEARCH 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57
* SEARCH 58 59 60 61 62 63 64 65 66 67 68 69 70 71 72 73 74 75 76 77 78 79 80 81 82 83 84
* SEARCH 85 86 87 88 89 90 91 92 93 94 95 96 97 98 99 100 101 102 103 104 105 106 107 108
* SEARCH 109 110 111 112 113 114 115 116 117 118 119 120 121 122 123 124 125 126 127 128
* SEARCH 129 130 131 132 133 134 135 136 137 138 139 140 141 142 143 144 145 146 147 148
* SEARCH 149 150 151 152 153 154 155 156 157 158 159 160 161 162 163 164 165 166 167 168
* SEARCH 169 170 171 172 173 174 175 176 177 178 179 180 181 182 183 184 185 186 187 188
* SEARCH 189 190 191 192 193 194 195 196 197 198 199 200 201 202 203 204 205 206 207 208
* SEARCH 209 210 211 212 213 214 215 216 217 218 219 220 221 222 223 224 225 226 227 228
* SEARCH 229 230 231 232 233 234 235 236 237 238 239 240 241 242 243 244 245 246
A04 OK SEARCH completed
```
查询收件箱中的新邮件 :
```sh
$ A05 search new                       ＃ 查询收件箱所有新邮件
A05 OK SEARCH completed                ＃ 查询结束，没有新邮件                                     
```
获取第5封邮件的邮件头 :
```sh
$ A06 fetch 5 full                     ＃ 获取第5封邮件的邮件头
* 5 FETCH (INTERNALDATE "18-Dec-2009 04:30:09 -0500" RFC822.SIZE 1680 ENVELOPE ("Fri, 18 Dec 2009 17:29:50 +0800" "tang7" (("tanghaibo" NIL "haibo.tang" "borqs.com")) (("tanghaibo" NIL "haibo.tang" "borqs.com")) (("tanghaibo" NIL "haibo.tang" "borqs.com")) ((NIL NIL "patest000" "aol.com")) NIL NIL NIL "<1261128590.18284.1.camel@tang-desktop>") BODY ("TEXT" "PLAIN" ("CHARSET" "US-ASCII") NIL NIL "7BIT" 11 3) FLAGS (Seen XAOL-READ XAOL-GOODCHECK-DONE))
A06 OK FETCH completed                           
```
获取第5封邮件的完整内容 :
```sh
$ A07 Fetch 5 rfc822                   ＃ 获取第5封邮件的完整内容
* 5 FETCH (RFC822 {1680}
Return-Path: <haibo.tang@borqs.com>
Received: from rly-db07.mx.aol.com (rly-db07.mail.aol.com [172.19.130.82]) by air-db01.mail.aol.com (v126.13) with ESMTP id MAILINDB012-ad84b2b4b8f3c2; Fri, 18 Dec 2009 04:30:09 -0500
Received: from n23a.bullet.mail.mud.yahoo.com (n23a.bullet.mail.mud.yahoo.com [68.142.207.189]) by rly-db07.mx.aol.com (v125.7) with ESMTP id MAILRELAYINDB076-ad84b2b4b8f3c2; Fri, 18 Dec 2009 04:29:51 -0500
Received: from [68.142.194.243] by n23.bullet.mail.mud.yahoo.com with NNFMP; 18 Dec 2009 09:29:51 -0000
Received: from [68.142.201.65] by t1.bullet.mud.yahoo.com with NNFMP; 18 Dec 2009 09:29:51 -0000
Received: from [127.0.0.1] by omp417.mail.mud.yahoo.com with NNFMP; 18 Dec 2009 09:29:51 -0000
X-Yahoo-Newman-Id: 562278.71547.bm@omp417.mail.mud.yahoo.com
Received: (qmail 98905 invoked from network); 18 Dec 2009 09:29:51 -0000
Received: from  (haibo.tang@122.200.68.238 with plain)
        by smtp105.biz.mail.sp1.yahoo.com with SMTP; 18 Dec 2009 01:29:50 -0800 PST
X-Yahoo-SMTP: TfaA1_iswBCT10k0GvJw0zydevUTmDrWpNf4lHPzuw--
X-YMail-OSG: RIeQgREVM1kE0gHJ.iLiMjGDwooofRGoV1X_AuoiD5Uo9glK5kT2P0FW_yRsDisu_BJ6dGXNfRPipXbfAlX0RaEPyiXvce6oEBfClbAe9uBxle4F.pGDTnwWtS._7XBApftxoqUxcUfYllW87RGLwrKhQKZBLYmNtpcuP936ndPuvmfpP36ZG3EcF6DMDiSQypQccEAJOatekQAPbPWeafgWWEHQwD4kO2MVobmQWFEXOaRr.5ScKg--
X-Yahoo-Newman-Property: ymail-3
Subject: tang7
From: tanghaibo <haibo.tang@borqs.com>
To: patest000@aol.com
Content-Type: text/plain
Date: Fri, 18 Dec 2009 17:29:50 +0800
Message-Id: <1261128590.18284.1.camel@tang-desktop>
Mime-Version: 1.0
X-Mailer: Evolution 2.26.1 
Content-Transfer-Encoding: 7bit
X-AOL-IP: 68.142.207.189
tang7
)
A07 OK FETCH completed                                  
```
查询第5封邮件的标志位 :                     
```sh
$ A08 fetch 5 flags                     ＃ 查询第5封邮件的标志位
* 5 FETCH (FLAGS (Seen XAOL-READ XAOL-GOODCHECK-DONE))
A08 OK FETCH completed             
```
设置第5封邮件的标志位为删除 :                     
```sh
$ A09 Store 5 +flags.silent (deleted)  ＃ 设置第5封邮件的标志位为删除
A09 OK STORE completed            
```
永久删除当前邮件箱中所有设置了deleted标志的信件 :                     
```sh
$ A10 Expunge                           ＃ 永久删除当前邮件箱中所有设置了deleted标志的信件
* 5 EXPUNGE
A10 OK EXPUNGE completed          
```
退出 :                     
```sh
$ A11 Logout                            ＃ 退出
* BYE IMAP server terminating connection
A11 OK completed
Connection closed by foreign host.
```
### Talk with Google Imap Server Using openssl toolkit ==
```sh
$ openssl s_client -crlf -quiet -connect imap.gmail.com:993
depth=1 /C=US/O=Google Inc/CN=Google Internet Authority
verify error:num=20:unable to get local issuer certificate
verify return:0
* OK Gimap ready for requests from 122.200.68.247 32if9202003pzk.98
# login with account
1 login borqsmail@gmail.com testtest
* CAPABILITY IMAP4rev1 UNSELECT LITERAL+ IDLE NAMESPACE QUOTA ID XLIST CHILDREN X-GM-EXT-1 UIDPLUS COMPRESS=DEFLATE
1 OK borqsmail@gmail.com authenticated (Success)
# list folders
2 list "" * 
* LIST (HasNoChildren) "/" "Deleted Items"
* LIST (HasNoChildren) "/" "Drafts"
* LIST (HasNoChildren) "/" "INBOX"
* LIST (HasNoChildren) "/" "Junk E-mail"
* LIST (HasNoChildren) "/" "Sent"
* LIST (HasNoChildren) "/" "Sent Items"
* LIST (HasNoChildren) "/" "Trash"
* LIST (Noselect HasChildren) "/" "[Gmail]"
* LIST (HasNoChildren) "/" "[Gmail]/All Mail"
* LIST (HasNoChildren) "/" "[Gmail]/Drafts"
* LIST (HasNoChildren) "/" "[Gmail]/Sent Mail"
* LIST (HasNoChildren) "/" "[Gmail]/Spam"
* LIST (HasNoChildren) "/" "[Gmail]/Starred"
* LIST (HasChildren HasNoChildren) "/" "[Gmail]/Trash"
* LIST (HasNoChildren) "/" "sb"
* LIST (HasNoChildren) "/" "&V4NXPpCuTvY-"
2 OK Success
# idle mode
3 idle
+ idling
```

### Talk with QQ Imap Server using openssl toolkit ==
貌似证书不太对，可还能连接成功
```
#!sh
$ openssl s_client -crlf -quiet -connect imap.qq.com:993                                                              ~dragon
depth=0 /C=CN/ST=Guangdong/L=Shenzhen/O=Tencent Technology(Shenzhen) Company Limited/OU=R&D/OU=Terms of use at www.verisign.com/rpa (c)05/CN=imap.qq.com
verify error:num=20:unable to get local issuer certificate
verify return:1
depth=0 /C=CN/ST=Guangdong/L=Shenzhen/O=Tencent Technology(Shenzhen) Company Limited/OU=R&D/OU=Terms of use at www.verisign.com/rpa (c)05/CN=imap.qq.com
verify error:num=27:certificate not trusted
verify return:1
depth=0 /C=CN/ST=Guangdong/L=Shenzhen/O=Tencent Technology(Shenzhen) Company Limited/OU=R&D/OU=Terms of use at www.verisign.com/rpa (c)05/CN=imap.qq.com
verify error:num=21:unable to verify the first certificate
verify return:1
* OK [CAPABILITY IMAP4 IMAP4rev1 AUTH=LOGIN NAMESPACE] QQMail IMAP4Server ready
. login 85194354 woaiguozi 
. OK Success login ok
. select inbox
* 71 EXISTS
* 0 RECENT
* OK [UNSEEN 8]
* OK [UIDVALIDITY 1273502441] UID validity status
* OK [UIDNEXT 77] Predicted next UID
* FLAGS (Answered Flagged Deleted Draft Seen)
* OK [PERMANENTFLAGS (* Answered Flagged Deleted Draft Seen)] Permanent flags
. OK [READ-WRITE] SELECT complete
```
