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



## 顺序执行模式

### Callback方式

### Promise方式

### Generator方式

## 并发执行模式

## 流模式

## 参考

- [关于Zalgo](http://blog.izs.me/post/59142742143/designing-apis-for-asynchrony)