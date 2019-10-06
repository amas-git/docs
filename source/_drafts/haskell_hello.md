# Heskell
## 开发环境
我们需要两个东西
	- 编译器ghc
	- 构建工具cable

## ghcup
管理haskell的编译环境，包括安装设置
```
# install the last known "best" GHC version
ghcup install
# install a specific GHC version
ghcup install 8.2.2
# set the currently "active" GHC version
ghcup set 8.4.4
# install cabal-install
ghcup install-cabal
# update cabal-install
cabal new-install cabal-install
```

## Stack
stack是一个haskell开发环境管理软件，主要解决几个问题
 - 多个ghc版本的管理
 - 各种库的管理

总之有了stack,让你轻松起步, 现在就来安装它:

```sh
$ wget -qO- https://get.haskellstack.org/ | sh
```

## 第一个stack工程
```
$ stack new hello
$ tree hello
hello
├── ChangeLog.md
├── LICENSE
├── README.md
├── Setup.hs
├── app
│   └── Main.hs
├── hello.cabal
├── package.yaml
├── src
│   └── Lib.hs
├── stack.yaml
└── test
    └── Spec.hs
    
# 构建工程
$ cd hello
$ stack build
...
inking .stack-work/dist/x86_64-osx/Cabal-2.4.0.1/build/hello-exe/hello-exe # 注意这句话，这个就是我们想要的结果

# 一顿操作之后，我们多了一个.stack-work目录，瞧瞧吧
$ tree -d
.stack-work
│   ├── dist
│   │   └── x86_64-osx
│   │       └── Cabal-2.4.0.1
│   │           ├── build
│   │           │   ├── autogen
│   │           │   └── hello-exe
│   │           │       ├── autogen
│   │           │       └── hello-exe-tmp
│   │           ├── package.conf.inplace
│   │           └── stack-build-caches
│   │               └── 87b2b02c1d77144a2d8d67813e28883a8a4c4832a3998a267bbb7bfcf3620ff5
│   └── install
│       └── x86_64-osx
│           └── 87b2b02c1d77144a2d8d67813e28883a8a4c4832a3998a267bbb7bfcf3620ff5
│               └── 8.6.5
│                   ├── bin
│                   ├── doc
│                   │   └── hello-0.1.0.0
│                   ├── lib
│                   │   └── x86_64-osx-ghc-8.6.5
│                   │       └── hello-0.1.0.0-8ryclCAfnaR40fmnvrxHg0
│                   └── pkgdb

# 既然构建成功了，咱们如何运行一下呢?
$ .stack-work/dist/x86_64-osx/Cabal-2.4.0.1/build/hello-exe/hello-exe
someFunc

# 太麻烦了吧，没关系我们可以用另外一个命令
$ stack exec hello-exe
someFunc

# 太麻烦了，没关系可以使用另外一个命令
$ stack run
someFunc

# 如果你没有合适的构建环境，可以执行
$ stack setup

# 清理工作
# 删除构建过程中产生的文件
$ stack clean 

# 删除全部stack产生的文件，变成一个非常干净的工程，相当于stack clean --full
$ stack purge 
```

 - stack.yaml : 构建环境的配置
 - package.yaml : 如何构建当前的工程，利用这个文件生成.cable文件，让cable处理具体的构建任务
 - .cabal: 这些文件就不要修改了，具体的构建细节可以从里面看到

通常使用一些常用的包只要修改package.yaml的depandency就好,因为stack已经帮你构建好这些包了
```yaml
dependencies:
- base >= 4.7 && < 5
- text
```

但是，有时候我们会使用一些没有被stack构建好的包或是一个放在github上的包等等，这个时候除了修改package.yaml之外，我们还需要配置stack.yaml的extra-deps字段，让stack了解上哪里去搞到你需要这个包。
```
extra-deps:
- acme-missiles-0.3
- git: https://github.com/commercialhaskell/stack.git
    commit: e7b331f14bcffb8367cd58fbfc8b40ec7642100a
```