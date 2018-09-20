---
title: Zsh Array
date: 2018-03-27 11:10:17
tags:
---
# Zsh Array

## 数组
### 声明数组:
```zsh
set -A name
typeset -a name
```

### 定义数组:
```zsh
set -A name 1 2 3 4 5 6
typeset -a name ; set -A name 1 2 3 4 5 6
name=(1 2 3 4 5 6)
```

### 索引
假如: 
```zsh
xs=(a b c d)
```

```zsh
$ echo $xs[1]
a
$ echo $xs[-1]
a
$ echo ${xs[3]}
c
$ echo ${xs[$((1+1))]}
b
$ echo $xs[1,3] 
a b c
```

### 数组和字符串
字符串可以看作字符数组.
```zsh
$ x=1234567
$ echo $x[1,3]
123
```

### 数组的长度
```zsh
$ xs=({1..10})
$ echo ${#xs}
10
$ xs=({1..100})
$ echo ${#xs}
100
```

```zsh
# 从数组中随机挑选元素
$ birthday=({1..12})
$ echo "$birthday[$[${RANDOM}%${#birthday}+1]]"
```


### 追加元素: +=
> array+=element|array
比如:
```zsh
$ nums=({1..3})  ; echo $nums
1 2 3 
$ nums+=4        ; echo $nums
1 2 3 4
$ nums+=({5..9}) ; echo $nums
1 2 3 4 5 6 7 8 9
```

### 查找元素的下标: $array[(i)pattern]
 - 从左向右依次查找(反向查找使用:I)
 - 如果匹配pattern, 则返回该元素的下标
 - 如果零匹配，则返回数组长度+1(即: $#array+1)
 
```zsh
$ x=(apple age ago)
$ print $x[(i)ago]
3
$ print $x[(i)notfound]
4
$ print $x[(i)a*]
1

# 通常都是这么判断数组中是否包含某个元素
$ [[ $x[(i)ago] -le $#x ]] && print "found: ago"
found: ago
```

### 查找元素: $array[(r)pattern]: 
 - 从左向右依次查找(反向查找使用:R)
 - 如果匹配pattern, 则返回该元素
 - 如果无法匹配，则返回空

```zsh
$ x=(apple age ago)
$ print $x[(r)ago]
ago
$ print $x[(r)notfound]

$ print $x[(i)a*]
ago

# 通常都是这么判断数组中是否包含某个元素
$ [[ -n $x[(r)ago] ]]  && print "found: ago"
found: ago
```

### 删除元素: $array[n]=()
```zsh
$ x=(a b c)
$ x[2]=
$ print $x
a c
```

### 获取全部元素: $array[@] 和 $array[*]
 - 他们都会扩展为数组的所有元素
 - "$X[@]" 等价于 $X[*]
 - "$X[@]" 与 "$X[*]"意义不同`
  - "$X[@]" : 相当于 "$1" "$2" ... "$N"` (N个参数)
  - "$X[*]" : 相当于 "$1 $2 ... $N"` (1个参数)

例如:
我们用echo-args打印参数明细:
```
#!sh
#----------------------
#!/bin/zsh

msgI() {
    echo "$FG[green]$*"
}

msgI '[ARGS VALUES]':$*
msgI '[ARGS COUNTS]':$#
#----------------------
```

做如下试验:
```
$ ./echo-args $xs[*]                                                    /data/src/zsh/array
[ARGS VALUES]:a b c d
[ARGS COUNTS]:4
$ ./echo-args $xs[@]                                                    /data/src/zsh/array
[ARGS VALUES]:a b c d
[ARGS COUNTS]:4
$ ./echo-args "$xs[*]"                                                  /data/src/zsh/array
[ARGS VALUES]:a b c d
[ARGS COUNTS]:1
$ ./echo-args "$xs[@]"                                                  /data/src/zsh/array
[ARGS VALUES]:a b c d
[ARGS COUNTS]:4
```

### 笛卡尔积: $^array
 - 数组本身构成了一个数据集合(不考虑有相同元素的情形)，2个集合以上可进行求笛卡尔积运算. '$!^^算符即可完成此运算.
 - $^告诉zsh在expansion时按照笛卡尔积展开

```zsh
$ x=(a b c)
$ y=(1 2 3)
$ z=(A B C)
$ print $^x$^y
a1 a2 a3 b1 b2 b3 c1 c2 c3
$ print $^x$^y$^z
a1A a1B a1C a2A a2B a2C a3A a3B a3C b1A b1B b1C b2A b2B b2C b3A b3B b3C c1A c1B c1C c2A c2B c2C c3A c3B c3C
```

## 关联数组
关联数组是特殊的数组, 其元素个数总是偶数, 保存多个key/value. 其数据类型需要通过typeset -A 来指定, 指定后可以使用数组的赋值方式.
```zsh
typeset -A name
name=(k1 v1 k2 v2 ... kx vx)
```

### 定义关联数组:
```zsh
$ typeset -A people 
$ set -A people name x sex F 
$ echo $people
x F

# 等同于
$ typeset -A people
$ people=(name x sex F)

# 等同于
$ typeset -A people
$ people[name]=x
$ people[sex]=F

# 等同于
$ typeset -A people
$ typeset "people[name]=x"
$ typeset "people[sex]=F"
```

### 打印所有key:
```zsh
$ echo ${(k)people}
```

### 打印所有value:
```zsh
$ echo ${(v)people}
$ echo $people
$ echo ${people[*]}
$ echo ${people[1,-1]}

# for可以绑定两个以上的变量
for key value in $people; do
    print $key=$value
done

# 健壮性更好的方式如下，可以兼容value为空的情况
for k v in "${(Pkv)${map}[@]}"; do
    print $k $v
done
```

### 删除键值
```zsh
$ typeset -A people
$ people=(:name x :sex F :age 18)
$ print $people
18 x F
$ unset "people[:sex]" # 注意不带'$'号!!!
$ print $people
18 x
```

### 反转数组: ${(Oa)array}
```zsh
$ x=(a b c d e)
$ echo $x
a b c d e
$ echo ${(Oa)x}
e d c b a
```

## Brace Expansion
讲讲数组的BraceExpansion
比如:
```zsh
$ xs={m n q}
$ echo X{x,$xs}Y
$ echo X${^xs}Y
$ echo X${xs}Y
```
