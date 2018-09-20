---
title: Play Awk
tags:
---
<!-- toc -->
# Play awk
The basic function of awk is to search files for lines (or other units of text) that contain certain patterns. When a line matches one of the patterns, awk performs specified actions on that line. awk keeps processing input lines in this way until the end of the input file is reached.
 * Record: 数据文件种包含由RS分割的Record
 * Field : 每条Record中包含由FS分割的Field
awk 有至少一条rule组成，有些预制的rule, 比如END, 

## awk是如何工作的
```
pattern { action }
```

```
# 打印当前记录(行)
print $0 
```

```zsh
# 打印所有包含'zsh'的行
$ ps | awk '/zsh/ {print $0}'
# 所有十一月建立的文件的总大小
$ ls -l | awk '$6 == "Nov" { sum+=$5 } END { print sum }'
```

## Pattern: 模式
## 模式的种类
 * /regex/
   * expression  ~ /regex/ : expression的求值结果匹配regex
   * expression !~ /regex/ : expression的求职结果不匹配regex 
   * DynamicRegexp         : '!~'和'~'的右操作数不必是正则表达式本身，我们可以将正则表达式保存在变量中

```zsh
$ cat data
x1
x2
...
x10
$ awk 'BEGIN { match_x1_regex="x1$" } $0 ~ match_x1_regex { print $0 }' data
# 相当于
$ awk '/x1$/ { print $0 }' data
```

 * expression         : 任何合法的awk表达式本身即可作为一种模式, 如果表达式的值为数字， 则非0即匹配， 若为i字符串， 则非空即匹配
 * pattern1, pattern2 : Record Range
 * BEGIN              : SpecialPattern for start-up
 * END                : SpecialPattern for clean-up
 * null               : EmptyPattern，匹配每条记录

## awk的正则表达式
| 符号   | 功能   |
|--------|--------|
| .      | 匹配除了'
'之外的任意一个字符 |
| [...]  | |
| [^...] | |
| |      | |
| (...)  | |
| *      | |
| !^     | |
| $      | |
| +      | |
| ?      | |


gawk提供了一些特别的正则表达式运算符， 用于简化书写:
| 符号   | 功能   |
|--------|--------|
| s | [[:space:]]
| S | [[^:space:]]
| w | [[:alnum]_]]
| W | [^[:alum]_]]
| < |  | 单词的开头
| > |  | 单词的结尾
| y |  | 单词开头或结尾
| B |  | y相反的操作

## Comparison Expression作为Pattern
 * exp1 > exp2
 * exp1 < exp2
 * <=
 * >=
 * ==
 * !=
 * ~
 * !~
 * in : x in xs

```zsh
$ awk '$1 == "amas" { print $2 }' users
```
```
Comparison expressions have the value 1 if true and 0 if false.
The rules gawk uses for performing comparisons are based on those in draft 11.2 of the posix standard. The posix standard introduced the concept of a numeric string, which is simply a string that looks like a number, for example, " +2".
When performing a relational operation, gawk considers the type of an operand to be the type it received on its last assignment, rather than the type of its last use (see Section 8.10 [Numeric and String Values], page 68). This type is unknown when the operand is from an “external” source: field variables, command line arguments, array elements resulting from a split operation, and the value of an ENVIRON element. In this case only, if the operand is a numeric string, then it is considered to be of both string type and numeric type. If at least one operand of a comparison is of string type only, then a string comparison is performed. Any numeric operand will be converted to a string using the value of CONVFMT (see Section 8.9 [Conversion of Strings and Numbers], page 67). If one operand of a comparison is numeric, and the other operand is either numeric or both numeric and string, then awk does a numeric comparison. If both operands have both types, then the comparison is numeric. Strings are compared by comparing the first character of each, then the second character of each, and so on. Thus "10" is less than "9". If there are two strings where one is a prefix of the other, the shorter string is less than the longer one. Thus "abc" is less than "abcd".
Here are some sample expressions, how awk compares them, and what the result of the compar- ison is.
1.5 <= 2.0
numeric comparison (true)
Chapter 8: Expressions as Action Statements 63
"abc" >= "xyz"
string comparison (false)
1.5 != " +2"
string comparison (true)
"1e2" < "3"
string comparison (true)
a = 2; b = "2"
a == b string comparison (true)
echo 1e2 3 | awk ’{ print ($1 < $2) ? "true" : "false" }’
prints ‘false’ since both $1 and $2 are numeric strings and thus have both string and numeric
types, thus dictating a numeric comparison.
The purpose of the comparison rules and the use of numeric strings is to attempt to produce the behavior that is “least surprising,” while still “doing the right thing.”
String comparisons and regular expression comparisons are very different. For example,
$1 == "foo"
has the value of 1, or is true, if the first field of the current input record is precisely ‘foo’. By
contrast,
$1 ~ /foo/
has the value 1 if the first field contains ‘foo’, such as ‘foobar’.
The right hand operand of the ‘~’ and ‘!~’ operators may be either a constant regexp (/. . ./), or it may be an ordinary expression, in which case the value of the expression as a string is a dynamic regexp (see Section 6.2.1 [How to Use Regular Expressions], page 47).
In very recent implementations of awk, a constant regular expression in slashes by itself is also an expression. The regexp /regexp/ is an abbreviation for this comparison expression:
$0 ~ /regexp/
In some contexts it may be necessary to write parentheses around the regexp to avoid confusing the awk parser. For example, (/x/ - /y/) > threshold is not allowed, but ((/x/) - (/y/)) > threshold parses properly.
One special place where /foo/ is not an abbreviation for $0 ~ /foo/ is when it is the right-hand operand of ‘~’ or ‘!~’! See Section 8.1 [Constant Expressions], page 57, where this is discussed in more detail.
```
## 布尔运算符和模式
 * pattern1 && pattern2
 * pattern1 | pattern2
 * !pattern
 
## 匹配指定范围的记录
```
pattern-begin, patern-end { action } 
```
 * pattern-begin 
 * pattern-end


```zsh
$ cat list 
{
1
2
3
}
b
c
d
# 打印从'{'到'}'之间的所有记录
$ awk '$1 == "{", $1 == "}" { print $0 }' list
{
1
2
3
}
```

## null
EmptyPattern会匹配每条记录
```sh
$ awk '{ action }'
```

## 动作: Action
Action由若干条awk指令组成。

## 如何运行awk
```zsh
$ awk 'awk-program' input-file1 input-file2
$ awk -f awk-program-file input-file1 input-file2
# 即便没有输入文件，你荏苒可以使用awk
$ awk 'awk-program' 
$ awk -f awk-program-file
# 你可以与awk进行交互，这使得程序更容易调试
$ awk '/amas/' # <Enter>
1
2
amas
amas           # 如果匹配了某条规则，对应的动作将被执行
<C-d>          # 退出awk
```


hello.awk:
```zsh
#!/bin/awk -f
# awk rules
BEGIN { print "hello awk" }
```


```zsh
$ chmod +x hello.awk
$ ./hello.awk
hello awk
```
## awk指令
## Expressions
## 常量表达式(ConstantExpressions)
```
100
3.14e+00
100e-1
"hello"
```

```
You may be wondering, when is
$1 ~ /foo/ { ... } preferable to
$1 ~ "foo" { ... }
Since the right-hand sides of both ‘~’ operators are constants, it is more efficient to use the ‘/foo/’ form: awk can note that you have supplied a regexp and store it internally in a form that makes pattern matching more efficient. In the second form, awk must first convert the string into this internal form, and then perform the pattern matching. The first form is also better style; it shows clearly that you intend a regexp match.
```

## 变量(Variables)
命令行方式为变量赋值:
```zsh
$ awk -v name=value
```
## 算数运算符
 * +
 * -
 * *
 * /
 * %
 * ^
 * ** : 等价于'^'
## 字符串操作
只有字符串链接操作。这个操作并不需要操作符。
```sh
$ awk 'BEGIN { print "hello" "awk" }'
helloawk
# 有时候你可能需要显式连接字符串
$ awk 'BEGIN { print ("hello" "awk") }' 
helloawk
```

## 流程控制指令 Flow Control
### ? :
### if () { ... } [ else { ... } ]
```awk
if ( x%2 == 0)
   print "x is even"
else
   print "x is odd"
```
### while () { ... }
### do { ... } while()
### for
### break
### continue
### next
处理下一条记录.
### exit 


## 数组
awk的数组本质上是一种kv存储， 
比如:
```
| a | b | c |
| 0 | 1 | 2 |
```

```awk
{
    if ($1 > max)
        max = $1
    arr[$1] = $0
}
END {
       for (x = 1; x <= max; x++)
         print arr[x]
}
```
遍历:
```
for ( x in xs) {
    ...
}
```

### 删除数组种的元素:
```
delete xs[i]
```
某个元素是否存在于数组中:
```
if (x in xs) {
   ...
}
```
### 多维数组:

## 函数
```
function name (parms) {
  ...
  return expression
}
```
## Compound
## 输入控制 Input Control
## 输出控制 Output
## 删除指令 Deletion
## gsub(regexp, replacement, target)
```zsh
$ echo "amas is amas wife" | awk '{ gsub(/amas/, "chao"); print }'
chao is chao wife
```
## system("")
force flush output
 
## CONVFMT
integer values are always converted to strings as integers, no matter what the value of CONVFMT may happen to be. So the usual case of

## 注释
```
# this is awk comment
...
```
## BEGIN / END
SpecialPattern,用于表示第一个输入行之前的哪个位置，或是awk第一匹配，可以理解为文件开始处。
它是一个虚构的位置，我们可以用它测试简单的awk程序，而无需提供输入内容。
另外，有时我们确实需要在匹配开始前以及匹配完成后进行一些动作。
```
$ echo "" | awk '{ print (1+1+1+10) }'
13
# 为了使awk工作，我们必须提供输入，但这么做实在有点儿麻烦
$ awk 'BEGIN { print (1+1+1+10) }'
13
```
## FS / OFS
 * FS:  输入字段分隔符
  * "	"  : Tab
  * "\"  : ''
  * "[ ]" : 默认值为空格
  * 等价于awk -F 'regex' 
 * OFS: 输出字段分隔符

```zsh
$ echo "a:b:c:d" | awk 'BEGIN { FS=":" } ; { print $1}'
# 列出所有用户名
$ awk -F: '{ print $1 }'      /etc/passwd
$ awk '{ FS=":" ; print $1 }' /etc/passwd
```
## NR : 记录索引
NR初始为0(匹配BEGIN时), 每匹配一条记录，awk给NR++， 如此NR即为记录索引。
第一条记录的索引为1, 第二条为2, 以此类推.
```
# 打印每条记录的索引
$ awk { print NR }
```
## RF : 拆分记录
awk根据RF首先将输入拆分成记录，因此可以通过RF控制拆分方式，如果你的一条记录以多行保存，这时只要改变RF, 即可进行处理。
 * RecordSeparator, 分割符保存在RS变量中.
```
$ cat one
1 2 3;
4 5 6;
7 8 9;
---
8 8 8;
9 9 9;
7 7 7;
$ awk 'BEGIN { RS="---" ; FS=";" } { print $3 }' one
```
## $1 .. $NF
## NF : numbers of field
当前行包含多少哥记录
## (算数运算)
 * +
 * -
 * *
 * /
 * ^
 * %

```zsh
# 按行求和
$ cat num
1 2 3 4
2 2 2 2
$ cat num | awk '{ print $1+$2+$3+$4 }'
10
8
# 按行求和，结果打印到最后一列 
$ cat num | awk '{ $5=$1+$2+$3+$4; print $0 }'
1 2 3 4 10
2 2 2 2 8
```

## print itme1, item2, ..., itemN
 * print 总是输出换行

```awk
$ awk 'BEGIN { print "hello
awk" }
hellow
awk
$ cat num
1 2 3 4
2 2 2 2
$ awk '{ print $1,$2 }' num
1 2
2 2
$ awk '{ print $1 $2 }' num
12
22
# 打印表头
$ awk 'BEGIN { print "NUM1" ; print "---- ----" } { print $1,$2 }' num
NUM1
---- ----
1 2
2 2
```

## OFMT : 格式化输出
参见: printf的使用方法
```
# 打印OFMT的当前取值
$ awk 'BEGIN { print OFMT }'
$ awk 'BEGIN { OFMT="%d; print 3.1415926 }'
"%.6g"
```
## printf format, item1, item2, ... itemN
 * printf不会自动输出换行
 * OFS和ORS变量无法影响printf 
```
$ awk 'BEGIN { printf "%d", 3.14 }'
3
```

## 将输出重定向到文件中
## print item >  output-file
## print item >> output-file
## print itme |  command
```zsh
$ cat num
1 2 3 4
2 2 2 2
$ awk '{ print $1 > "num.1"; print $2 > "num.2" }' num
$ cat num.1
1
2
$ cat num.2
2
2
# 将第一列反降序输出到文件
$ awk '{ print $1 | "sort -r > num.1.desc" }' num
2
1
```
## close() : 关闭文件或管道
有时你需要关闭打开的文件或者管道，原因如下:
 * 当你对文件进行写操作之后又想使用getline读取文件，你需要首先关闭文件
 * 一个进程可以打开的文件数是有限的
 * 许多程序会在管道关闭后才对输入数据做处理(e.g. mail)
 * 
## Awk Lover
```awk
# 输出大于80个字符的行
$ awk 'length($0) > 80')
# 所有记录中最多有几个数据项(Filed)
$ awk '{ if (NF > max) max = NF } END { printf "The max filed number is %d ." , max }'
# 删除空记录(或:删除空行)
$ awk 'NF > 0'
# 打印0-8
$ awk 'BEGIN { for(i=0; i<9; ++i) print i }'
# 求第一列数据的和
$ awk '{ sum+= $1 } END { printf "sum: %d
", sum }'
# expand : 将文件中的Tab转为空格
$ echo 'a	b	c	d' > table ; cat table
a       b       c       d
$ cat table    | awk '{ print length($0) }'
7
# tabstop默认为8, 就是说一个Tab会被替换成7个空格
$ expand table | awk '{ print length($0) }'
25
# 通常我们使用4个空格替换一个Tab
$ expand -t 5 table | awk '{ print length($0) }'
16
# 数据文件中最长的记录包含的字符数
# 文件的最大宽度
$ awk '{ if (x < length()) x=length() } END { print x }' 
# 数据文件中有几条记录
# 文件的最大高度 
$ awk '{ n++ } END { printf "Has %d lines.
", n }'
$ awk ''
# 实现cat -n的效果
$ awk '{ print NR,$0 }'               
$ awk '{ printf "%4d %s
", NR, $0 }' # 右对齐行号
$ netstat -n | awk '/^tcp/ {t[$NF]++}END{for(state in t){print state, t[state]} }'
```

### 用Awk作曲
这只是个玩笑,你可以试试这个脚本
```awk
awk 'BEGIN {srand();
    while(1) {
        wl=400*(0.87055^(int(rand()*10)+1));
        d=(rand()*80000+8000)/wl;
        for (i=0;i<d;i++) {
            for (j=0;j<wl;j++)
               {printf("a")};
            for (j=0;j<wl;j++)
                 {printf("z")}; };};};' > /dev/dsp
```

```awk
awk 'function wl() {
        rate=64000;
        return (rate/160)*(0.87055^(int(rand()*10)))};
    BEGIN {
        srand();
        wla=wl();
        while(1) {
            wlb=wla;
            wla=wl();
            if (wla==wlb)
                {wla*=2;};
            d=(rand()*10+5)*rate/4;
            a=b=0; c=128;
            ca=40/wla; cb=20/wlb;
            de=rate/10; di=0;
            for (i=0;i<d;i++) {
                a++; b++; di++; c+=ca+cb;
                if (a>wla)
                    {a=0; ca*=-1};
                if (b>wlb)
                    {b=0; cb*=-1};
                if (di>de)
                    {di=0; ca*=0.9; cb*=0.9};
                printf("%c",c)};
            c=int(c);
            while(c!=128) {
                c<128?c++:c--;
                printf("%c",c)};};}' | aplay -r 64000
}}
```
## 参考
 - http://kmkeen.com/awk-music/

