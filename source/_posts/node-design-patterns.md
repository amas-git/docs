---
title: Node设计模式
date: 2018-09-24 00:42:51
tags:
---



## 回调函数最后原则

回调总是置于最后一个参数

```javascript
function fn(arg1, arg2, ..., argN, callback) { ... }
```



## 回调函数错误第一原则

回调函数中第一个参数总是用来传递错误，并且错误只能是**Error**类型。

```js
(error, arg1, arg2, ... argN) => { ... }
```

## 同步回调或异步回调



- 同步回调: 在当前线程进行回调
- 异步回调: 在其他线程进行回调

如果函数的一个参数是回调函数， 实现回调时要么采用同步回调，要么采用异步回调，不要两种模式都有。两种模式同时存在也叫释放Zalgo。

![zalgo](https://i0.kym-cdn.com/photos/images/newsfeed/000/193/872/ZALGO_by_Cheezyspam.jpg?1320110443)

```js
function fn(callback) {
	callback                             // 同步回调 
	setTimeout(() => { callback }, 100); // 异步回调
}
 

# 如果fn处理callback时，有时候是同步方式，有时候是异步方式，那么当callback引用外部变量时因为调用顺序的不确定而遭遇不测。
let cache = undefined;
function fn(callback) {
    if(cache) {
        // sync callback
        callback(cache);
        return;
    }
    // async callback
    setTimeout(() => {
        cache=1
        callback(cache);
    }, 0);
}

let flag = false;
fn((r) => {
    assert(flag === true);
});
flag = true
```



## 异步顺序执行模式

最简单的异步顺序执行模式就是回调，现在我们来举个例子。

```js
function randomIn(min, max) {
    return Math.round(Math.random() * (max - min))  + min;
}

function print(tag, o) {
    console.log(`${tag} : ${typeof o === 'string' ? o : JSON.stringify(o) }`);
}

class Task {
    constructor(name, time = randomIn(1,100)) {
       this.name = name;
       this.time = time;
    }

    process(params, callback) {
        setTimeout(() => {
            callback(this._do(params));
        }, this.time);
    }

    _do(params) {
        let result = { task: this.name, result : randomIn(10,99) };
        print(`CALL Task@${this.name} WITH ${JSON.stringify(params)}`, ` = ${JSON.stringify(result)}`);
        return result;
    }
}
```

假如有3个Task，我们如何让这三个任务按照顺序依次执行呢？

```
let task1 = new Task("task1");
let task2 = new Task("task2");
let task3 = new Task("task3");
```



### Callback方式

```js
task1.process("init", (result) => {
     task2.process(result, (result) => {
         task3.process(result, (result) => {
             print(`ALL DONE: ${result}`);
         }) ;
     });
});
```

不考虑异常处理，我们用callback来实现的版本就是这个样子。它的特点是回调套回调，要分析代码的执行顺序，我们不得不一层一层的代码的看，时刻要留意自己是在哪个callback里，层级过多的时候非常难于阅读和理解。这种情况我们也叫做回调地狱（Callback Hell）。其时我们我们更习惯于代码的书写顺序即是程序的执行顺序，即：

```
task1.process(...)
task2.process(...)
task3.process(...)
```

要做到这一点就需要promise或者generator。

### Promise方式

```js
_process(params) {
    return new Promise( resolve => {
        setTimeout(() => {
            resolve(this._do(params));
        }, this.time);
    });
}
```

我们需要改写一个promise 方式的处理函数，我们叫这个函数为_process(), 它的代码如下，返回了一个promise对象。promise对象有一个then方法可以，你在这个then 方法里面可以得到上个promise的执行结果，同时返回下一个需要执行的promise，假如你有10个task，你只要不断的then下去就可以了。

```js
task1._process("init")
    .then(result => {
        return task2._process(result);
    })
    .then(result => {
        return task3._process(result);
    }).then((result) => {
        print("ALL DONE", result);
    });
```

经过Promise我们终于可以把逐层嵌套的callback 改写成单层的then方法序列。但是这样写仍然很啰嗦，每处理一个任务，就需要写一个then方法和一个callback。我们想要继续简化代码就需要ES6提供的async/await关键字了。在两个语法糖的帮助下，我们可以这么实现：

```js
(async () => {
    let r1 = await task1._process("init");
    let r2 = await task2._process(r1);
    let r3 = await task3._process(r2);
    print("ALL DONE", r3);
})();
```

- await只能在async函数中被使用
- 用async修饰的函数总是返回一个Promise对象，即便你返回的不是Promise对象，也会被转换为Promise对象。
- 使用await你就可以等待一个Promise执行完毕再继续执行下面的代码

### Generator方式

generator是一种特殊的函数，普通的函数要么不执行，要么全部执行。generator则可以部分的执行,  通过一个叫做yield的关键字你可以暂停函数的执行并且返回一结果。

```js
function * g() {
	yield  1;
	let i = yield  2;
	return 3;
}

let G = g();
console.log("%o", G.next()); // { value: 1, done: false }
console.log("%o", G.next()); // { value: 2, done: false }
console.log("%o", G.next()); // { value: 3, done: true }
```

- generator函数直接调用后你可以得到一个对象，这个对象可以控制内部代码的执行

- next可以恢复generator的执行，并且执行到下一个yiled。

- next也向generator传递参数

```js
function * add() {
    let a = yield "STEP 1";
    let b = yield "STEP 2";
    return a + b;
}

let Add = add();
console.log("%o", Add.next());  // {"done":false}
console.log("%o", Add.next(1)); // {"done":false}
console.log("%o", Add.next(2)); // {"value":3,"done":true}
```

那这些特性和控制异步代码执行的顺序有什么关系呢？既然每当执行到yield, generator的代码就不再继续执行了，其时这就是一个同步的机制。我们可以发起一个异步调用，然后将代码yield住，当这个异步调用完成的时候通过generator的next方法将执行结果返回给我们，并且继续下面的代码。为此我们得将generator 交给异步执行代码。

``` js
function async(generatorFunction) {
    function callback(...argv) {
        generator.next(argv);
    }

    let generator = generatorFunction(callback);
    generator.next();
}

async(function * (callback) {
    let r;
    r = yield task1.process("init", callback);
    r = yield task2.process(r, callback);
    r = yield task3.process(r, callback);
});
```



## 并发执行模式

## 流模式

## 参考

- [关于Zalgo](http://blog.izs.me/post/59142742143/designing-apis-for-asynchrony)