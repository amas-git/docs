# HELLO GO

![](/src/cmcm/ice_river/docs/assets/2019-07-31-103739_815x522_scrot.png)

## 开发环境配置

```bash
# 安装好go
$ go env
GOBIN=""
GOCACHE="/home/amas/.cache/go-build"
GOEXE=""
GOFLAGS=""
GOHOSTARCH="amd64"
GOHOSTOS="linux"
GOOS="linux"
GOPATH="/home/amas/go"
GOPROXY=""
GORACE=""
GOROOT="/usr/lib/go"
GOTMPDIR=""
GOTOOLDIR="/usr/lib/go/pkg/tool/linux_amd64"
GCCGO="gccgo"
CC="gcc"
CXX="g++"
CGO_ENABLED="1"
GOMOD=""
CGO_CFLAGS="-g -O2"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-g -O2"
CGO_FFLAGS="-g -O2"
CGO_LDFLAGS="-g -O2"
PKG_CONFIG="pkg-config"
GOGCCFLAGS="-fPIC -m64 -pthread -fmessage-length=0 -fdebug-prefix-map=/tmp/go-build276003034=/tmp/go-build -gno-record-gcc-switches"
```

```bash
# 设置环境变量
GOPATH="/home/amas/go"
GO_BIN=$GOPATH/bin
```



### 快速体验

如果你想立刻感受一下go，可以直接打开浏览器体验。

- https://play.golang.org/

### VSCODE

vscode安装go插件，排名最高的那个就可以，然后随便建立一个test.go文件， 用vscode打开， vscode会提示你安装gotools, 点全部安装即可。

因为golang.org被wall, 需要开vpn, 或者从github上下载源代码，直接安装。

```
Installing 11 tools at /home/amas/go/bin
  gocode
  gopkgs
  go-outline
  go-symbols
  guru
  gorename
  dlv
  gocode-gomod
  godef
  goreturns
  golint

Installing github.com/mdempsky/gocode SUCCEEDED
Installing github.com/uudashr/gopkgs/cmd/gopkgs SUCCEEDED
Installing github.com/ramya-rao-a/go-outline SUCCEEDED
Installing github.com/acroca/go-symbols SUCCEEDED
Installing golang.org/x/tools/cmd/guru SUCCEEDED
Installing golang.org/x/tools/cmd/gorename SUCCEEDED
Installing github.com/go-delve/delve/cmd/dlv SUCCEEDED
Installing github.com/stamblerre/gocode SUCCEEDED
Installing github.com/rogpeppe/godef SUCCEEDED
Installing github.com/sqs/goreturns SUCCEEDED
Installing golang.org/x/lint/golint SUCCEEDED
```

这些包的源代码被下载到`$GOHOME/src`下面, 所以如果因为被wall而安装失败，可找到安装失败的那个工具的github, 然后直接clone到对应的路径

> 注意: 路径要保持一致

hello.go:

```go
package main

import (
	"fmt"
	"net/http"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {
	http.HandleFunc("/", greet)
	http.ListenAndServe(":8080", nil)
}

```

```bash
# 使用code runner直接运行， 或者输入命令;
$ go run hello.go
Hello World! 2019-07-31 09:10:43.562093238 +0800 CST m=+18.202586775
```

## 历史

2007年GOOGLE的Rovert Griesemer, Rob Pike和Ken Tomsaon老汉发起，GO的设计思想是:

1. 简单易于使用的语法
2. 具有动态语言的编程体验
3. 面向对象编程
3. 静态类型
5. 编译为native binary, 效率更高
6. 接近0的编译时间，感觉如同解释型语言
7. 简单易用的并发支持，可充分发挥多CPU多核心的性能
8. 垃圾回收，自动内存管理



```
We wanted a language with the safety and performance of statically compiled languages such as
C++ and Java, but the lightness and fun of dynamically typed interpreted languages such as
Python.
—Rob Pike, Geek of the Week
```



## 语言核心

```
One True Brace Style:
使用go你会少打很多分号，但是这些分号其实还是存在的，只不过编译器预处理帮你插入了分号。但并不是无代价的，你必须遵守编程固定的编程风格，没太多的选择反而挺好，less is more是不是？


```





### 不支持的特性

1. 参数默认值



### BASIC

### 变量

```
var age uint = 1
var r,g,b uint8 = 0, 88, 88
```



### Short Declaration

```
var n = 100
n := 100
```

短声明节省了var的使用，另外也有助于写出更加简洁的代码

```
for n := 0; n < 100; n++ {
	// just do it
}
```

> 注意: 无法在package scope中使用短声明



### 表达式

	- 不支持 --i， ++i
	- i++， i--不能作为表达式

### 简单语句

```go
x := 1   //短声明
x += 1  //
f()  // 函数调用
<- c // channels receive
i++, i--
-> c // channels send
```

其他的都是非简单语句， 简单语句可以用在for/if后面



### 函数

### 接口



### 循环和迭代

#### for

```go
// forever
for  {

}

for i:=0; i<100; i++ {
    
}
```



#### range

```go
s := "hello world"

for i,c := range s {

}
```



### 分支跳转

#### switch/case

> case默认不穿透，要想穿透需要用fallthrough

#### if/else

```go
if n--; n > 0 {

} else {

}
```



#### for/range

#### type/switch

#### select/case/default

```
switch n := rand.Intn(100); n % 7 {
	case 1:
	case 2
	default:
}
```



### 类型系统

#### 实数

```go
pi := 3.1415  // float64
n := 100   // int
int[8|32|64]
uint[8|32|64]


```

Type Alias

```go
type byte uint8
type rune int32
```



#### 大整数

- big.Int
- big.Float
- big.Rat

```
n := big.NewInt(99999999)

n := new (big.Int)
n.SetString("9999999999999999", 10);
```

#### String

#### nil

### Structure

```go
var point struct {
	x int
	y int
}

// 使用structure定义新类型
type location struct {
	x int
	y int
}

type box struct {
    width  int
    height int
}

var a box
a.width = 1
a.height = 2

b := box{11, 12}

c := box{height: 1}

fmt.Println(a, b, c)
```



#### struct输出为json

```go
	type person struct {
		Name string `json:"name"` 
		Age  int    `json:"age"`
	}
	students := []person{
		{"zhou", 11},
		{"bob", 12},
		{"amas", 13},
	}
	bytes, err := json.Marshal(students)
	fmt.Println(bytes, err)
	fmt.Println(string(bytes)) // FIXME: NOT WORK???
```

* 只有大写属性才会输出
* 通过struct tag可以

#### 绑定函数到struct上:

```go
type person struct {
  Name string `json:"name"` 
  Age  int    `json:"age"`
}


// (p person) 叫做receiver
func (p person) print() {
  fmt.Printf("NAME %s IS $d years old", p.Name, p.Age)
}
```

> 注意: struct是值传递
>
> ```go
> type person struct {
> 	Name string `json:"name"`
> 	Age  int    `json:"age"`
> }
> 
> // 值传递
> func (p person) print() {
> 	fmt.Printf("[NAME:%8s | AGE: %2d]\n", p.Name, p.Age)
> }
> 
> amas := person{"amas", 110}
> amas.print()
> 
> // 指针传递
> func(p person) {
>   p.Age *= 2
>   p.print()  // age: 220
> }(amas)
> 
> amas.print() // age: 110
> 
> func(p *person) {
>   p.Age *= 3
>   p.print()  // age: 330
> }(&amas)
> amas.print() // age: 330
> ```

#### 构造函数

GO里面没有构造函数，但是可以用一般函数代替构造函数， 按照约定这类函数用new或New开头

```go
func NewPerson(name string, age int) Person {
  return Person{name, age}
}
```

#### struct组合

> composition 与 inheritance是不同的概念
>
> 组合比继承具有更好的可重用性

```go
type AB struct {
	A
	B
}

type A struct {
	name string
}

type B struct {
	age int
}

func (a A) FA() {
	fmt.Printf("FA name %v\n", a.name)
}

func (b B) FB() {
	fmt.Printf("FB age %v\n", b.age)
}


ab := AB{A{"amas"}, B{19}}
fmt.Println(ab)
ab.FA() // ab.A.FA()
ab.FB() // ab.B.FB()
ab.name // amas
ab.age  // 19
```

> 注意:
>
> 组合的时候可能会出现Name Collisions, 这时候你需要在顶层实现这个方法

判断是否实现了接口i.(Type)

```go
type Say {
		say() string
}

var one Say = A{"one"}
  iSay, ok := one.(Say)
  if ok {
  fmt.Println(iSay.say())
}
```



#### interface

```go
var t interface {
		talk() string
}

type Comparable interface {
  Compare(a,b interface) int
}
```

```
Used together, composition and interfaces make a very powerful design tool.
                                                                     —Bill Venners, JavaWorld
```

> interface{}叫做空接口， 所有的类型都包含空接口，因此空接口可以指向任何类型的数据



### 基础类型

	- string
	- bool
	- int8(byte)
	- uint(byte)
	- int16
	- uint16
	- int32(rune)
	- uint32
	- int64
	- uint64
	- int
	- uint
 - uintptr
   	- 指针类型不能进行算数运算
      	- 指针类型不能转换为其他类型的指针
	- float32
	- float64
	- complex64
	- complex128



### 类型转换

```
n := 42
m := float64(n)
```



> 小心处理类型转换
>
> 1996年6月4日，阿利亚那5号发射升空37秒后自解体，事故的原因就是类型转换错误。
>
> float64转int16导致处理器trap, 损失3.7亿美元。虽然在go语言里这种转换不会导致异常，顶多数溢出，但是处理数据转换的时候最好进行范围检查。
>
> https://hownot2code.com/2016/09/02/a-space-error-370-million-for-an-integer-overflow/

### 定义新类型

```
type cm = int
```







## 函数

```

```

### return

go可以return多个值

```go
func f(a int, b int) (int, int) {
	return a,b
}

// 也可以为返回值命名， 这样最后调用return一并返回
function double(a , b int) (m, n int) {
	m = 2*a
	n = 2*b
    return
}
```

### defer

gp函数在调用reture后进入exiting phase

```

```

内置函数

	- print
	- println
	- real
	- imag

### 匿名函数

- alias: anonymous function , function literal

```
fn := func (a int, b int) {
	retirm a + b
}

// IIFE: Immediately Invoked Function Expression
func() {

}()
```





### 首字母大小写决定访问控制

```
函数/变量/其他标识符

首字母大写 = public
首字母小写 = private
```


可以存在多个init函数, 最终会被合并在一起执行

```go
func init() {

}
```

init函数的调用顺序是:

	- 先执行依赖包里的init
	- 再执行当前包里的init
	- 所有的init都执行完毕，再调用main



### 方法Methods

> Method是和特定类型绑定的函数， 类似于js的bind, 方法可以与一个对象绑定



## Collections

### 数组

```
xs := [3]int // 数组
xs := []int   // slice
```

### Slice

> 通常不会直接使用数组， 而是使用slice
>
> Slice实际上是一个数组的视图

```
name[start:end:capacity]

// 定义slice
s := []int{1,2,3,4}

s := make ([]int, 0, 100)
```

### map

```go
name := map[string]int{
	"a": 1,
	"b": 2,
	"c": 3
}

name['d'] = 15


// key是否存在？
if value,ok := name['key']; ok {

} else {

}

// map as set
set := map[string]bool{}
set["key1"] = true

if set["key1"] {

}
```

### 错误处理

> To err is human; to forgive, divine.
> —Alexander Pope, “An Essay on Criticism: Part 2”

```
Errors are values.
Don’t just check errors, handle them gracefully.
Don’t panic.
Make the zero value useful.
The bigger the interface, the weaker the abstraction.
interface{} says nothing.
Gofmt’s style is no one’s favorite, yet gofmt is everyone’s favorite.
Documentation is for users.
A little copying is better than a little dependency.
Clear is better than clever.
Concurrency is not parallelism.
Don’t communicate by sharing memory, share memory by communicating.
Channels orchestrate; mutexes serialize.
—Rob Pike, Go Proverbs
(see go-proverbs.github.io)
```

SEE: http://go-proverbs.github.io/

错误也是数据， 正常去处理就好了

### 包管理

源码组织:

```
project
    ├── bin
    ├── main.go
    ├── pkg
    └── src
```



#### 包管理工具

- go get

- glide: https://glide.readthedocs.io/en/latest/getting-started/

  - ```bash
    $ go get github.com/Masterminds/glide 
    $ go install github.com/Masterminds/glide 
    ```

    

- go dep

#### 定义自己的模块

- package $name

- main 函数必须定义到main包里

- 多个go文件可以存放在同一目录下，属于同一个包
- vendor作为特殊目录， 和go.mod文件所在的目录共享同一个名字
- module
  - GO111MODULE环境变量为on
  - go.mod文件
  - build时可以通过`-mod=vedor`来打包vendor目录下的代码
- go不支持循环import

```go
import $importname $import_path
import format "fmt"
```



### dot import

```go
import . "fmt"

Println("Let's go") 
```





#### 锁: sync.Mutex



### 内存管理

### 极速编译

### 内置测试

go内置了API和工具用于测试代码的覆盖率，性能测试等等。

### 内置文档



### 并发模型

#### goroutine

```
go some_function
```

- 状态:
  - running
    - blocking
  - 任意时间最多有不超过runtime.NumCPU个goroutine同时执行
- M-P-G模型: https://docs.google.com/document/d/1TTj4T2JO42uD5ID9e89oa0sLKhJYD0Y_kqxDv3I3XMw/edit
  - M : OS Thread
  - P:  实际上G到M的执行是由P来控制的， P中有一个runqueue保存G, 此外还有调度器来分配时间片
    - runtime.GOMAXPROCS(-1)来查看P的数量，Go1.5以后这个就等于runtime.NumCPU(),
    - 对于重IO业务，可以把这个调大
  - G: goroutine

> TODO:
>
> 当P=1时，如果panic了，会不会导致阻塞？





## sync

- Cond: 主要用来阻塞coroutine, 等待其他coroutine发来信号之后再WatiPass

  ```go
  // 忙等
  for something() == false {
  	time.Sleep(time.Second)
  }
  
  // Cond等, 等待的corutine调用c.Signal()
  var c sync.NewCond(&sync.Mutex{})
  c.L.Lock()
  if something() == false {
  	c.Wait()
  }
  
  c.L.Unlock()
  ```

  

- Mutex

  - Lock()制造临界区
  - UnLock()退出临界区

- Once

  - Do(func), 无论在多少个corotine中调用once.Do(), 其中的函数只被执行一次

- Map: 并发安全的Map

- WaitGroup: 安全计数器

  - Wait(): 阻塞当前coroutine
  - Add(n): 加n
  - Done(): 减1， 为0时WaitPass

- Pool: 用于创建并发安全对象池

- RWMutex: 区分读写操作的Mutex, 写的时候读



## Channels

CSP: Communicating Sequential Processes CSP的概念诞生于1978年，简单来说，就是把输入和输出作为程序设计语言的内置功能



> sync中所有的功能都可以概括为控制内存访问顺序，Channels的目的也是如此

Channels背后的概念是CSP, 并不是只有Go采用CSP, Erlang也是

### 定义channels

channels有三种类型, 这说法不够严谨，但方便记忆

	- chan :  双向可收发
	- chan <-: 只写
	- <- chan: 只读

```go
c := make(chan, int)

var c chan int
var writeOnly chan <- int
var readOnly <-chan  int
```





## 参考

- go 1.4.3是最后一版用C实现的go, 此后go用go语言实现， 这个很好的解释了先有鸡还是现有蛋的问题
- https://medium.com/rungo/the-anatomy-of-functions-in-go-de56c050fe11
- 本机Go的源码安装位置: `$ echo $(go env GOROOT)/src`

