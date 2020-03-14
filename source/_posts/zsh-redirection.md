---
title: Zsh重定向
tags:
---
<!-- toc -->
# Zsh 重定向
If a command is followed by & and job control is not active, then the default standard input for the command is the empty file /dev/null. Otherwise, the environment for the execution of a command contains the file descriptors of the invoking shell as modified by input/output specifications.
命令使用的StandardInput是空文件:`/dev/null`.
The following may appear anywhere in a simple command or may precede or follow a complex command. Expansion occurs before word or digit is used except as noted below. 
If the result of substitution on word produces more than one filename, redirection occurs for each separate filename in turn.
## 标准输入重定向: < word
只读方式打开文件word, 作为StandardInput.
## 标准输入输出重定向: <> word
读写方式打开文件word, 作为StandardInput. 如不存在, 则建立之.
## 标准输出重定向: > word
  - 写方式打开word文件,作为StandardOutput. 如不存在, 则建立之.
  - 如文件存在且没有设置CLOBBER选项, 则产生错误, 否则文件word被修改为空文件.
## >| word
## >! word
Same as >, except that the file is truncated to zero length if it exists, even if CLOBBER is unset.
如文件存在且CLOBBER选项开启, 则文件word强制
## 追加方式重定向: >> word
追加方式打开文件word, 作为命令的StandardOutput
## >>| word
## >>! word
Same as >>, except that the file is created if it does not exist, even if CLOBBER is unset.
## <<[-] word
The shell input is read up to a line that is the same as word, or to an end-of-file. No parameter expansion, command substitution or filename generation is performed on word. The resulting document, called a here-document, becomes the standard input.
If any character of word is quoted with single or double quotes or a `', no interpretation is placed upon the characters of the document. Otherwise, parameter and command substitution occurs, `' followed by a newline is removed, and `' must be used to quote the characters `', `$', ``' and the first character of word.
Note that word itself does not undergo shell expansion. Backquotes in word do not have their usual effect; instead they behave similarly to double quotes, except that the backquotes themselves are passed through unchanged. (This information is given for completeness and it is not recommended that backquotes be used.) Quotes in the form $'...' have their standard effect of expanding backslashed references to special characters.
If <<- is used, then all leading tabs are stripped from word and from the document.
## <<< word
Perform shell expansion on word and pass the result to standard input. This is known as a here-string. Compare the use of word in here-documents above, where word does not undergo shell expansion.
## <& number
## >& number
The standard input/output is duplicated from file descriptor number (see man page dup2(2)).
## 关闭标准输入: <& -
## 关闭标准输出: >& -
关闭StantardInput/StantardOutput
```
#!sh
# 你将无法看到dmesg命令输出的任何内容
$ dmesg >& -
```
# <& p
# >& p
The input/output from/to the coprocess is moved to the standard input/output.
# >& word
# &> word
(Except where `>& word' matches one of the above syntaxes; `&>' can always be used to avoid this ambiguity.) Redirects both standard output and standard error (file descriptor 2) in the manner of `> word'. Note that this does not have the same effect as `> word 2>&1' in the presence of multios (see the section below).
类似'> word', 将StandardOutput和StandardError重定向到文件word中. 等价于:
```
#!sh
#! 多IO方式
$ command > word 2>&1
```
# >&| word
# >&! word
# &>| word
# &>! word
Redirects both standard output and standard error (file descriptor 2) in the manner of `>| word'.
# >>& word
# &>> word
Redirects both standard output and standard error (file descriptor 2) in the manner of `>> word'.
>>&| word
>>&! word
&>>| word
&>>! word
# 例子 ===
```
#!sh
# 将命令的输出重定向到一个文件,并且打印到终端上
$ date | tee date.txt
$ date > date.txt | cat
# 给当前目录下每个.sh文件末尾追加一行"exit 0"
$ echo "exit 0" >> *.sh
```
文件x:
```
b
c
a
e
f
```
文件y:
```
1
2
3
8
7
6
5
4
```
现在我们将x/y文件中的内容排序:
```zsh
$ sort < x < y
$ sort < y < x
$ cat x | sort y
$ cat y | sort x
$ cat x y | sort
# 结果都一样,为:
1
2
3
4
5
6
7
8
a
b
c
e
f
# 将排序结果保存到xy.sorted文件中:
$ sort < x < y > xy.sorted
```
