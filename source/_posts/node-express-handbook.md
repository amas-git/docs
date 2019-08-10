---
title: node_express_handbook
date: 2019-03-12 16:22:57
tags:
---

# Express 使用手册 (>4.0)

## 4.0的变化

1. connect被移出了express, 很多中间件需要单独安装
2. body-parser是一个独立的模块了
3. app.configure 被移除, 可以使用app.get(env)
4. 可以去掉app.use(app.router), router不再是中间件?

其他迁移到4.0的注意事项: https://expressjs.com/en/guide/migrating-4.html

## 安装

## Hello Express

## 如何上传文件

## 中间件就是pipline

>  在中间件中你可以干两件事, 或调用next()结束中间件的执行, 或调用send返回response, 不然客户端就卡住了.

是如何工作的

1. Route Handler:  app.VERB路由可以认为是专门处理HTTP请求的中间件, 客户端请求过来先由他接住
2. Route Handler会要求设置两部分, 一部分是路由匹配规则, 本质就是正则表达式. 第二部分是一个回调函数, 这个回调函数的形式可以是以下三种
   1. (req, res)
   2. (req, res, next)
   3. (error, req, res, next)

```javascript
# 一个中间件长这样, 
app.use((req, res, next) => {
    
});

# 假如你有三个中间件
app.use(m1);
app.use(m2)
app.use(m3)

# 请求过来, 首先由Route Handler接住, 然后按照m1, m2, m3这样的顺序处理

```

## 常用的中间件

Express 4.0 把好多内置的中间件都搞到connect这个库里了.

```js
var connect = require(connect);


```



## 使用Node发送邮件

## 如何扩展(Scale Up 和 Scale Out)

## 什么是app cluster?





# D3 入门手册

如何利用gooogle找到炫酷的D3控件, 例如我们想找词云(wordcloud), 可以在google中输入

> wordcloud site:bl.ocks.org



## Scales: 比例尺

比例尺的作用是将数据集合映射到可视化区域集合的一个工具.

在地图上可以看到比例尺

![Maps scale visible in Mac OS X](http://cdn.osxdaily.com/wp-content/uploads/2014/07/maps-scale-mac.jpg)



比例尺是一个函数 S :: Domain -> Range, 将数据集合映射到可视化集合, 这种映射包括但不限于

1. 数据范围的伸缩

![an example of how scales work](http://www.jeromecukier.net/wp-content/uploads/2011/08/d3scale1.png)





```svg
<div style="width:400px;height:300px">
  <svg width="100%" height="100%" xmlns="http://www.w3.org/2000/svg">
    <defs>
      <pattern id="smallGrid" width="8" height="8" patternUnits="userSpaceOnUse">
        <path d="M 8 0 L 0 0 0 8" fill="none" stroke="gray" stroke-width="0.5"/>
      </pattern>
      <pattern id="grid" width="80" height="80" patternUnits="userSpaceOnUse">
        <rect width="80" height="80" fill="url(#smallGrid)"/>
        <path d="M 80 0 L 0 0 0 80" fill="none" stroke="gray" stroke-width="1"/>
      </pattern>
    </defs>

    <rect width="100%" height="100%" fill="url(#grid)" />
  </svg>
</div>
```


### Axis

### 基于D3的库

d3是一个数据化可视框架, 具体到某一特定领域, 可能开发效率更加重要, 

### NVD3 

致力于开发可重用图表的项目, 很多常用的图表你都不用从头搞起, 直接使用nvd3就好.

- http://nvd3.org/


### C3

- https://c3js.org/



### 参考

- https://github.com/mbostock/d3/wiki/Gallery
- https://github.com/mbostock/d3/wiki
- https://github.com/mbostock/d3
- https://groups.google.com/forum/?fromgroups#!forum/d3-js
- http://www.jeromecukier.net/2011/08/11/d3-scales-and-color/e
- 