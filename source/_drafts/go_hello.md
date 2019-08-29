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

### 常量

```go
const PI := 3.1415

const (
	OK       = 200       // iota = 0
	CREATED  = OK + iota // iota = 1
	ACCEPTED             // iota = 2
	NOAUTH               // iota = 3

	REDIRECT = 300             // iota = 4
	MC       = REDIRECT + iota // iota = 5
)
```



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

### 定义新的类型: type

> type name literal

如:

```go
type Age int
type Months map[striing]int

type Color byte
type Box struct {
    width, height, depth float64
    color Color
}
```





### struct

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

// 初始化
b := box{11, 12}
c := box{height: 1}


// 赋值
var a box
a.width = 1
a.height = 2
```



匿名struct:

```go
p := struct{name string; age int}{"amas", 18}
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



> 注意: 是不是可以通过receiver来扩展内置类型呢？ 
>
> ```go
> func (s *string) Hello() {
>     // 编译器会告诉你: cannot define new methods on non-local type string
> }
> ```
>
> 什么是

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
>
> 组合之后子struct的所有公开函数都可以被父struct直接使用， 只要名字不冲突子struct,子子struct的方法都可以不必理会

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



### interface

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

判断是否实现了接口element.(Type)

> go中并没有接口实现声明，这样做的好处是接口函数的声明和实现可以不依赖于interface的存在， 我觉得是一种自底向上的过程，更加符自然，从具体到抽象，使得具体可以先于抽象

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

> duck typing
>
> 如果它走起路来像鸭子，会嘎嘎的叫，那么他就是鸭子

go的接口就是duck typing, 只要你有A接口的方法，那么你就是A接口类型， 跟你自身是什么类型没有关系



定义一个interface变量

```go
var r io.Reader
tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
if err != nil {
    return nil, err
}
r = tty
```

	 - r 背后对应(value, type)这样一个二元组，也就是(tty, *os.File), os.File的关联方法很多，但是r只能调用Read
	 - value中包含了全部的类型信息, 因此r实际上也是可以引用到value本身的， 这样我们也可以将r转化为其他可用的接口类型，比如

```go
var w io.Writer  // tty中也实现了Writer接口
w = r.(io.Write) // 看上去r是Reader接口变量，但是type assertion可以检测r引用的对象是不是有Writer接口
```

- w背后对应的是(tty, *io.File), 其实没有变化，这个就是静态类型系统

> 接口变量就好像一个方法集合的视图，接口像是调用白名单一样

更进一步, 这个你也就可以理解了

```go
var empty interface{}    // 空接口变量
empty = w.(interface {}) // type assertion
empty = w                // 但是我们知道空接口表示允许使用任何方法，所以type assertion可以省去
```

- empty背后还是(*tty, *ioFile)

到这里，应该明白为什么可以把任何类型的变量赋给空接口类型变量了吧

> 重点: 不管你把一个对象通过接口变量怎么折腾，背后仍然保留的是那个对象， 接口就是这个对象的`调用视图`



下面我们可以谈谈反射

> 反射只是一种允许程序从(value, type)中获取信息的能力

反射三定律

	1. 从interface value到reflection object (Interface -> Value)
 	2. 从reflection object到interface value (Value -> Interface)
 	3. 若想修改reflection object, value必须是settable

```go 
// 1. 从interface value到reflect object
x := 12
reflect.TypeOf(x)  // int
// Q: 这哪有interface的事？
// A: TypeOf()的签名: TypeOf func(i interface{}) Type,
//    - 1. 先new一个interface{}, 然后把x放到里面，这时候就可以作为参数传递给TypeOf了
//    - 2. 
reflect.TypeOf(reflect.TypeOf(x)) // *reflect.rtype
reflect.TypeOf(reflect.ValueOf(x)) // *reflect.Value

// 2. 从refelect object到interface value
// func (v Value) Interface() interface{}
```

底层类型Kind

```
type Age int
reflect.TypeOf(1)              // int
reflect.TypeOf(Age(1))         // main.Age
reflect.TypeOf(1).Kind()       // int
reflect.TypeOf(Age(1)).Kind()  // int
```



### make和new

- make(T, args)为go内置类型分配内存，比如map,slice, channel , 不会返回空， 因为这些内置类型都需要各自的初始化工作

- new(T)为类型T分配一个地址，返回值是指针(*T)

  

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

```go
fn := func (a int, b int) {
	retirm a + b
}

// IIFE: Immediately Invoked Function Expression
func() {

}()
```

### 函数的polymorphism

同名方法必须通过不同的接口来区分，同一个struct上不能有方法的重载

```go
type intInterface struct {
}

type stringInterface struct {
}

func (number intInterface) Add (a int, b int) int {
	return a + b;
}

func (text stringInterface) Add (a string, b string) string {
	return a + b
}
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

> 当你把一个函数绑定到某个类型上的时候， 这种函数叫methods

```
func (reciver Type) func_name(input) result 
```



### Black hole funcs

```go
func _() {}
func _() {}
```



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

### import

> import $importname $import_path
> import . $import_path
> import _ $import_path

```go
import "fmt"
import （
	“fmt”
	"io"
）

import . "fmt" // 可直接调用fmt中的函数而无需加上fmt
import _ "fmt" // 引用fmt包并调用包中的init方法
```

### Reflect

反射就是一种能够检测程序本身结构的能力， 反射通常是通过类型系统来实现的。为此我们先要了解Go的类型系统。

- https://blog.golang.org/laws-of-reflection
- Go是静态类型系统(Static Typed), 任何一个变量只能有一个固定类型，在编译的时候就已经确定
- type Age int提供了类型别名的功能，类型别名也是新类型， Age和int不是同一类型
- 接口类型代表有限方法的集合
- interface{}空接口表示，万物皆空，抽象的过程就是不断去掉个性保留共性，终极的共性就是空， 人们常说的go的接口是dynamically typed实际上是误导，实际上说的就是interface{}可以引用一切类型这件事，但是如果一个变量是空接口类型，那它就永远都是空接口类型，不可能是其他类型

### 内存管理

### 极速编译

### 内置测试

go内置了API和工具用于测试代码的覆盖率，性能测试等等。

### 内置文档

### 异常处理

go没有try/catch

1. 将错误当作程序的正常返回，区别处理
2. panic会打断程序的执行，进入到panic状态， 比如数组越界等
3. 也可以调用panic函数进入到panic状态中
4. 提供panic/recoveri机制
   1. 当在函数F调用panic后, F不再执行， 但是F的defer函数仍然会执行
   2. 如果想拦截住panic, 可以在defer函数中调用recover()

```go
panic("ERROR")


func testPanic() {
	f := func() {
		panic("I'm panic")
	}
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("catch panic: ", err)
		}
	}() // 只能在函数return后进行recover
	f()
}
```



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



### Go/CSP 和 Actor model的区别
>  Go/CSP和Actor想解决的问题: 
>
> CPU不能更快了， 想要更高的效率需要让程序能够并发运行，能够充分利用多核CPU

Actors: 

 - 最小的运算单元

 - 可以接受一些消息，然后做一些运算，每次只能处理一条消息

 - 与对象其实类似，不同的是，Actors之间是不共享内存的，这样Actor内部的状态没法直接被另两一个Actors之间修改

 - "one ant is not ant", 一个蚂蚁不能叫做蚂蚁，Actors也是一样的， 只有多个Actor存在相互配合才有意义

 - Actors之间通信通过messagebox, 每个actor都有自己的地址可以接受消息

 - message接收顺序是不确定的，每个消息将被尽最大可能(Best Effort)发送, 并且At Most Once

 - Actor收到消息后可以做三件事情

    - 创建更多的Actor (processing)

    - 发消息给其他Actor （communication）

    - 改变状态，决定如何处理下一条消息 (storage)

      

Actor Model : https://www.brianstorti.com/the-actor-model/

在Erlang,Scala中消息的传递使用的是Actor Model, 和Go/CSP有一点点区别，Go/CSP中必须创建Channel,
而Actor Model中Actor相互直接可以发消息, Actor模型中Channel可以认为是另外一个Actor

### race detection

```go
import (
	"fmt"
	"math/rand"
	"time"
)

var balance = 1000 // 总共的钱
var taken = 0 // 拿走的钱

// get money rnadom
func getMoney(seq, x int) (r bool) {
	defer func() {
		fmt.Printf("[%02d] : %5v (%02d/%03d)\n", seq, r, x, balance)
	}()

	if balance-x < 0 {
		r = false
		return
	}
	balance -= x
	r = true
	taken += x
	return
}

func main() {
	//runtime.GOMAXPROCS(2)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 100; i++ {
		go getMoney(i, rand.Intn(25)) // 随机取钱
	}
	// 简单的等待2秒
	time.Sleep(time.Second * 2)
	fmt.Printf("FINAL BALANCE : %d TAKEN %d", balance, taken)
}

```

go可以检测race condition
```sh
$ go run -race race.go

[00] :  true (23/977)
==================
WARNING: DATA RACE
Read at 0x0000005d7418 by goroutine 7:
  main.getMoney()
      /src/amas/docs/src/go/c10k/race.go:18 +0x81

Previous write at 0x0000005d7418 by goroutine 6:
  main.getMoney()
      /src/amas/docs/src/go/c10k/race.go:22 +0xdb

Goroutine 7 (running) created at:
  main.main()
      /src/amas/docs/src/go/c10k/race.go:33 +0x13c

Goroutine 6 (finished) created at:
  main.main()
      /src/amas/docs/src/go/c10k/race.go:33 +0x13c
==================
...
==================
[01] :  true (12/963)
...
[99] : false (10/002)
FINAL BALANCE : 2 TAKEN 986Found 3 data race(s)
exit status 66

# 拿走了986, 剩余2， 还有12不知道哪去了？ WHY?
```

现在我们使用sync.Mutex保护制造一个临界区，来避免race condition

```go
var mutex = new(sync.Mutex)

func getMoney2(seq, x int) (r bool) {
	mutex.Lock()
	defer mutex.Unlock()

	defer func() {
		fmt.Printf("[%02d] : %5v (%02d/%03d)\n", seq, r, x, balance)
	}()

	if balance-x < 0 {
		r = false
		return
	}
	balance -= x
	r = true
	taken += x
}
```

```
WARNING: DATA RACE
Read at 0x00000060dc08 by main goroutine:
  main.main()
      /src/amas/docs/src/go/c10k/race.go:58 +0x191

Previous write at 0x00000060dc08 by goroutine 28:
  main.getMoney2()
      /src/amas/docs/src/go/c10k/race.go:45 +0x17c

Goroutine 28 (finished) created at:
  main.main()
      /src/amas/docs/src/go/c10k/race.go:54 +0x13c
==================
```

> go的race检测是很靠谱的，我们使用Sleep等待这种方式也是有潜在问题的

我们再改造一下，使用WaitGroup来替代Sleep

```go
var mutex = new(sync.Mutex)
var wg sync.WaitGroup


func getMoney2(seq, x int) (r bool) {
  wg.Add(1)           // 增加等待
	defer wg.Done()     // 标记完成
	mutex.Lock()
	defer mutex.Unlock()

	defer func() {
		fmt.Printf("[%02d] : %5v (%02d/%03d)\n", seq, r, x, balance)
	}()

	if balance-x < 0 {
		r = false
		return
	}
	balance -= x
	r = true
	taken += x
}

func main() {
	//runtime.GOMAXPROCS(2)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 100; i++ {
		go getMoney2(i, rand.Intn(25))
	}

	
	wg.Wait() // 简单全部任务结束
	fmt.Printf("FINAL BALANCE : %d TAKEN %d", balance, taken)
}

```

> 再次用-race检测我们的程序， 这回不会有任何警告了

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
var writeOnly chan<- int
var readOnly <-chan  int
```

channels按照buffer的不同，也可分为三种

	- Unbuffered
	- Buffered
	- Unidirectional(单向`88)

### 超时处理

```go
import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func longTImeTask(ch chan string, max int) {
	time.Sleep(time.Second * time.Duration(max))
	ch <- strconv.Itoa(max)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ch := make(chan string)

	go longTImeTask(ch, rand.Intn(6))
	select {
	case t := <-ch:
		fmt.Printf("WORK DONE in %v\n", t)
	case <-time.After(time.Second * time.Duration(3)): // 3秒超时
		fmt.Print("TIMEOUT")
		close(ch)
	}
}

```



## 参考

- go 1.4.3是最后一版用C实现的go, 此后go用go语言实现， 这个很好的解释了先有鸡还是现有蛋的问题
- https://medium.com/rungo/the-anatomy-of-functions-in-go-de56c050fe11
- 本机Go的源码安装位置: `$ echo $(go env GOROOT)/src`
- https://blog.learngoprogramming.com/go-functions-overview-anonymous-closures-higher-order-deferred-concurrent-6799008dde7b
- https://research.swtch.com/interfaces
- http://luca.ntop.org/Teaching/Appunti/asn1.html

