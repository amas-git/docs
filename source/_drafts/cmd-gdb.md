---
title: gdb
tags:
---
# gdb
# [b]reak
```
break <function>
break <line-number>
break <file-name>:<line-number>
break 
```
# [i]nfo
# info reg
# [p]rint
# [c]ontinue
继续运行，直到下一个断点. 
```
continue
continue <number> :  跳过<number>个断点
```
# p/<format>
|| x || 十六进制
|| d || 十进制
|| u || 无符号十进制
|| o || 八进制
|| t || 二进制
|| a || 地址
|| c || 显示为字符
|| f || 浮点数
|| s || 字符串
|| i || 显示为机器语言
查看程序计数器(intel):
```
(gdb) p $pc
$1 = (void (*)()) 0x80483dd <f+13>
(gdb) p $eip
$2 = (void (*)()) 0x80483dd <f+13>
```
# e[x]amining
打印指定内存中的内容
```
x/<number><format><unit><address>
```
<unit>:
|| b  || byte          || 1
|| h  || half word  || 2
|| w || word         || 4
|| g ||                  || 8
查看当前指令:
```
(gdb) x/i $pc
# 显示当前指令之后的10条指令
(gdb) x/10i $pc
```
# [disas]semble
 * disassemble : 反汇编当前整个函数
 * disassemble $pc : 反汇编PC所在的整个函数
 * disassemble <start-addr> <end-addr> : 反汇编某个地址范围内的指令
```
(gdb) disassemble
(gdb) disassemble $pc
(gdb) disassemble $pc $pc+10
```
# [n]ext
# [s]tep
# ignore
# finish
# until
# watch <expr>
# awatch <expr>
# rwatch <expr>
设置监视点，监视点是断点的一种，当程序改变某个地址的内容时
# [d]elete
删除断点
# clear
# disable
# enable
# command
到达断点后执行命令
# set variable <name>=<expr>
设置变量的值.
# generate-core-file 
生成内核转储文件
# attach <pid>
# condition
# 历史值
|| $ || 最近的值 ||
|| $n || 第n个历史值 ||
|| $$ || 倒数第2个值 ||
|| $$n || 倒数第n个值 ||
|| $_ || x命令显示的最后地址 ||
|| $__ || x命令显示的 ||
|| $_exitcode ||  调试中的程序的返回代码 ||
# 配置文件: .gidbinit
# 自定义命令
```
define <command>
    ...
end
```
# 为命令添加帮助文档
```
document <command>
   ...
end
```
然后就可以使用`help <command>`来查看说明了。
# source <gdb-file>
使用souce命令可以将gdb命令文件加载到当前的环境中。
# 什么情况下无法使用调试器
# StackFrame遭到破坏
# 其他
# 最常用的gdb命令
# 查看寄存器
```
(gdb) i r eip ebp
```
```
backtrace full: Complete backtrace with local variables
up, down, frame: Move through frames
watch: Suspend the process when a certain condition is met
set print pretty on: Prints out prettily formatted C source code
set logging on: Log debugging session to show to others for support
set print array on: Pretty array printing
finish: Continue till end of function
enable and disable: Enable/disable breakpoints
tbreak: Break once, and then remove the breakpoint
where: Line number currently being executed
info locals: View all local variables
info args: View all function arguments
list: view source
rbreak: break on function matching regular expression
gdb7.0:
* record -- 记录必要的信息以便可以反向调试
* reverse-continue ('rc') -- Continue program being debugged but run it in reverse
* reverse-finish -- Execute backward until just before the selected stack frame is called
* reverse-next ('rn') -- Step program backward, proceeding through subroutine calls.
* reverse-nexti ('rni') -- Step backward one instruction, but proceed through called subroutines.
* reverse-step ('rs') -- Step program backward until it reaches the beginning of a previous source line
* reverse-stepi -- Step backward exactly one instruction
* set exec-direction (forward/reverse) -- Set direction of execution.
thread apply all bt or thread apply all print $pc: For finding out quickly what all threads are doing.
```

