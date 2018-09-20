---
title: hello_hla
tags:
---
# hla
 * [http://net.pku.edu.cn/~course/cs201/2003/Page_hla/0_hla_dnld.html 下载主页]
# 安装
```sh
HLA_HOME=/data/opt/hla
export PATH=$HLA_HOME:$PATH
export hlalib=$HLA_HOME/hlalib/hlalib.a
export hlainc=$HLA_HOME/include
export hlatemp=/tmp/.hla
```
# 第一个程序
```
program hello;
#include ("stdlib.hhf")
begin hello;
    stdout.put("hello!",nl);
end hello;
```
```sh
$ hla hello.hla
$ ./hello
hello!
```
# 基本类型
# int8或i8
# int16或i16
# int32或i32
# boolean
# char
# 基本指令
# mov(src, dst);
在一个程序中， 25%-40%的指令都是mov
 * `src`可以是
  * 寄存器
  * 内存变量
  * 常量
 * `dst`可以是
  * 寄存器
  * 内存变量
```div class=note
 1. 从技术角度来说，80x86不允许src和dst同时为内存变量
 2. mov的`src`和`dst`位数必须相同，就是说，或者同为8位，16位或32位
```
# add(src, dst);
# sub(src, dst);
# flags
 * @c : 进位
 * @nc : 无进位
 * @z : 零
 * @nz : 非零
 * @o : 溢出
 * @no : 未溢出
 * @s : 符号
 * @ns : 无符号
# if-else
```
 if () then
 elseif() then
 else
 endif;
```

