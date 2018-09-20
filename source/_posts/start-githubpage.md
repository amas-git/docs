---
title: 使用Github配置+Hexo构建你的博客
date: 2018-01-05 14:06:18
tags: 
---
# Github Pages
## 安装
 1. 必须有github帐号
 2. [怎么开通](https://pages.github.com/)

## 第一次使用Hexo
```
$ mkdir amas.docs
$ hexo init amas.docs
$ cd amas.docs
$ hexo server
INFO  Start processing
INFO  Hexo is running at http://localhost:4000/. Press Ctrl+C to stop.
```
浏览器打开 http://localhost:4000/ 就可以看到你的blog了

```
$ tree -L 1 .
.
├── _config.yml
├── db.json
├── node_modules
├── package.json
├── package-lock.json
├── scaffolds
├── source
└── themes

```
* _config.yml : 配置文件

## 开始写作
```
$ hexo new [layout] <title>
```
layout可以是
 * post : 默认值，保存到source/_post中
 * draft: 草稿, 保存到source/_draft
 * page: 页面, 保存到source

```
$ hexo new geth 
INFO  Created: /src/amas/amas.docs/source/_posts/geth.md
```

你也可以直接编译mk文件:
```
$ hexo generate
$ hexo g
```

## 部署
部署到gitpage:
```
$ npm install hexo-deployer-git --save
$ hexo deploy
$ hexo d
```
## 删除页面
有时候你可能想把一些文章给删掉，聪明的你可能会直接跑到_posts目录下把文章干掉，结果当你部署以后兴冲冲的打开页面时，发现那些文章还在WTF
```
$ hexo clean
```
## 安装主题
```
# 1. 你需要将主题下载到themes目录下
# 2. 修改_config.yml的theme字段为主题目录名即可
```
 * [主题浏览](https://hexo.io/themes/)
 * PolarBear Prince
 * 目前正在使用: https://github.com/yiliashaw/hexo-theme-prince


