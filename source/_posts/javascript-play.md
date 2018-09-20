---
title: JavaScript玩闹
date: 2018-03-22 11:11:58
tags:
---

## 异步
javascript的引擎并不了解时间的概念，它只保证在需要的时候可以执行所需要的代码片段。
### 事件循环和tick
时间循环不断的处理发生的时间，一个循环又叫做一个tick。
对于setTimeout(...)这样的函数，工作原理就是设置一个定时器，在到达预订时间后将回调函数放到事件循环中。
由于并不能插队，所以智能排队等待着被处理。如果前面的事件比较多，或者比较消耗时间，那么setTimeout可能会不那么精确。

> ES6之前并没有描述应该如何实现事件循环，ES6中对此做了精确的定义。
> Promise的引入，要求对事件循环能够直接进行精准控制。
> Promise并不是在event queue中而是被放入了job queue，
> 而诸如setTimeout函数则是放入到了event queue中

## 回调
### 控制转移是callback hell的根本原因
 - 使用回调一个常见的问题是控制转移，回调函数是否被回调，何时回调以及回调多少次都是由第三方来控制的，因此相当于代码控制权部分交给了第三方。
 - 即使回调委托给自己的代码，仍然会有类似问题存在，因此对于回调之后的参数检查是非常必要的，仍然要考虑以下几个问题:
   - 是否被回调
   - 是否过早或过晚被回调
   - 是否回调了过多的次数
 - 分离回调: 将错误回调和正常回调分开处理， 对于同样一个异步API,你需要指定两个回调函数
 - Error-First: Node风格的回调函数设计， 回调函数的第一个参数用来保存错误对象，若没有错误发生，这个对像为空
 

```js
function f(err, data) {
    if(err) {
        console.error(err);
    } else {
        console.log(data);
    }
}
```
## 反转控制转移:Promise
假如我们能够收回控制权，不就可以避免callback hell了么？

Promise代表将来值，一个Promise对象有三种状态:
 - Pending: 执行中
 - Fulfilled: 完成
 - Rejected: 失败

### 如何用Promise实现超时模式: race + timeoutPromise
### 被忽略掉的Promise怎么处理
### 


一旦Promise对象变成Fulfilled或者Rejected, 那么这个状态就没法再改变了. 

```js
new Promise(function(resolve, reject) {
        // 假如执行成功，调用resolve(return-value)
        // 失败调用rejected
});
```
 - resolve/reject这俩函数是系统定义好传过来的，作用就是改变Promise对象的状态

## prototype
 - 在Javascript语言中，new命令后面跟的不是类，而是构造函数。
 - 只有函数具有prototype属性
 - 原型链的最终root是Object.prototype
 - Function.prototype是个不同于一般函数（对象）的函数（对象）
 - Object.prototype instansof Function.prototype == true
 - Function.prototype instansof Object.prototype == true
 - 普通函数继承于Function.prototype
 - Function.prototype继承于Object.prototype, 且Function.prototype.prototype是null
 - Object是一个构造函数, 是Function的实例，Object._proto_
 - 先有Object.prototype（原型链顶端），Function.prototype继承Object.prototype而产生，最后，Function和Object和其它构造函数继承Function.prototype而产生。


用函数来实现类，js是这么干的:
```js
function People(name) {
    this.name = name; // this指的是新生成的People实例
}

new People('Jonny'); // <- new 后面其实是构造函数, 
```

这么干之后一个新的问题来了，这俩对象如何共享一些属性，比如在java里面我们可以通过static成员来共享，但是目前这个思路不行，不仅没法共享，而且会造成大量的重复资源。
因此想到假如给构造函数加一个prototype属性，这个属性用来保存共享的属性和方法貌似就解决了问题.

## 函数: 所有的函数都是对象


## this
### this指向函数自身? [错误]
```js
function f(num) {
    this.count += 1;
    console.log(num);
}

f.count = 0;

for(let i=0; i<10; ++i) {
    f(i);
}

console.log(f.count); // 答案: 0
```
 - 为什么是0？ 不是10？
 - this.count究竟干了什么？

现在来稍微修改一下代码,把this换成函数的名字
```js
function f(num) {
    f.count += 1;
    console.log(num);
}

f.count = 0;

for(let i=0; i<10; ++i) {
    f(i);
}

console.log(f.count); // 答案: 10
```
 - 看来this不是指向函数自身的，而函数的名字是


我们再来一种方法:
```js
function f(num) {
    this.count += 1;
    console.log(num);
}

f.count = 0;

for(let i=0; i<10; ++i) {
    f.call(f,i);
}

console.log(f.count); // 答案: 10
```

### this指向函数的作用域？ [错误]
 - javascript中尽管作用域有些类似对象，但是它只是一个引擎内部的概念，没有代码可以引用作用域

### this究竟是个啥？
 - this是在运行的时候才绑定到一个`上下文对象`
 - 函数被调用的时候，会创建一个活动记录(或者也可以叫做上下文)
 - this就是这个上下文的一个属性而已

因此this实际上要在运行的时候你才能搞明白它究竟绑定了什么，
为此有四条规则。当下面四条规则发生冲突的时候，靠后的规则具有更高的优先级.

#### 规则一: 默认绑定
非严格模式下，this可以绑定到全局对象，严格模式下则不行

```js
function foo() {
    console.log(this.a);
}

var a = 1;

foo(); // 严格模式下: 
```
 - foo()被调用时，没有任何修饰，所以智能用默认绑定，默认绑定就是在非严格模式下绑定到全局对象,
 如果严格模式，则this此时只能绑定到undefined

#### 规则二: 隐式绑定
```js
function foo() {
    console.log(this.a);
}

var o1 = {
    a: 1,
    foo: foo
}

var o2= {
    a: 2,
    o1: o1
}

o2.o1.foo(); //1, this绑定到o1上


fb = o2.o1.foo; // 我们定义一个fb指向o2.o1.foo会发生什么？
fb(); // undefined

setTimeout(o2.o1.foo, 0); //假如我们把函数作为callback给出去，结果也是undefined, 就如上面的例子一样
```

#### 规则三: 显示绑定
call和apply,这俩函数可以明确指定绑定对象，这就解决了默认绑定和隐式绑定的诸多问题。
```js
function foo() {
    console.log(this.a);
}

var o1 = {a:1};
var o2 = {a:2};

foo.call(o1); // this绑定到o1上，结果是1
foo.call(o2); // this绑定到o2上，结果是2
```
硬绑定: 现在我们用显式绑定来解决一下callback导致绑定丢失的问题:
```js
function foo() {
    console.log(this.a);
}

var o1 = {a:1};
var o2 = {a:2};

var fb = function() {
    foo.call(o1); // 显式绑定到o1上
}

setTimeout(fb,0); // 答案: 1, 这下不会有不确定的绑定问题了

// 经过这么一折腾，fb就被绑定到o1上了，这么做有些不那么灵活，所以我们稍微修改一下, 
// 这件事情的本质是提供一个明确this绑定的包装函数
function bind(fn, object) {
    return function() {
        fn.apply(object);
    };
}

setTimeout(bind(foo, o1)); // 答案: 1
setTimeout(bind(foo, o2)); // 答案: 2

// 如此常用的方法已经在ES5中提供了bind方法
setTimeout(foo.bind(o1), 0);
setTimeout(foo.bind(o2), 0);
```

既然可以显示指定绑定的上下文，那我们设计依赖于上下文执行的函数的时候，是不是可以明确指定一下绑定对象来避免这种混淆呢？ 结果我们多了一个参数，而少了一个this, 这恰恰说明this是怎样简化代码的。
```js
function foo(ctx) {
    console.log(ctx.a);
}

foo(o1);
foo(o2);
```

#### 规则四: new 绑定, 可以改变this绑定
回想一下java如何new一个object

```js
Object o = new Object();
```
 - Object()是一个特殊的函数，我们叫做构造函数
 - 在javascript的世界里，并没有所谓的构造函数，所有的函数都可以通过new来修饰调用，因此所谓的构造函数只不过是一个普通的函数，并不属于某个object.

通过new来调用函数， 实际上会发生如下的事情:
 1. 构建一个全新的对象
 2. 这个新对象进行prototype链接
 3. 新对象绑定到函数调用的this
 4. 如果函数没有返回其他对象，那么new调用函数后返回新对象

```js
function foo(a) {
    this.a = a;
}

foo("A"); // 普通的函数调用: this在严格模式下理应绑定到undefined上

new foo("A"); // new 调用， this被绑定到新创建的foo对象上, new影响了this的绑定行为
```

最后，我们来看看new绑定和显示绑定谁更霸道一些:
```js
var o1 = { a: 1}

function fx(a) {
    if(this.a != 1) {
        this.a = a;
    }
    console.log(this.a);
}
fx.prototype.a= 0;
var fx_hardbind = fx.bind(o1); // 显示绑定到o1上了
fx_hardbind(111)               // 因为绑定到o1上， o1.a == 1, 所以尽管传入111, 无法完成赋值，结果仍然是1
new fx_hardbind(111);          // 这回new改变了this的绑定，不是绑定到o1了，而是绑定到新创建的对象，a == 0, 所以可以进行赋值， 最终答案是:  111
```

## 模块
```js
// mylibs.js定义并导出方法
exports.hello = new function() {
    return "hello";
}

// 使用模块
var MyLibs = require('./mylibs');
console.log(MyLibs.hello());

// 获取模块
```
## 参考
 * [ECMA262/ECMAScript6](http://www.ecma-international.org/ecma-262/6.0/ECMA-262.pdf)
 * http://www.ruanyifeng.com/blog/2011/06/designing_ideas_of_inheritance_mechanism_in_javascript.html
