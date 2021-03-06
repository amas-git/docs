---
title: ZSH补全系统
date: 2018-09-19 11:43:20
tags:
---
<!-- toc -->
## Completion Widget
## 补全系统
### Context
Context是一个字符串, 用来标识用户输入光标正在什么位置. 所有的补全动作都是跟上下文相关的. 这个是补全系统的核心概念.



#### 如何查看Context?
``` 
_complete_help这个函数可以帮助我们打印Context. 这个函数其实是一个ZleWidget, 你直接像正常命令一样输入它是不会得到有用的结果的. 使用ZleWidget有两种方法:

 1. ALT+x 然后输入ZleWidget的名字
 2. 将ZleWidget绑定到一个快捷键上, 通过快捷键调用
 
默认_complete_help会绑定到CTRL-x h这个快捷键上, 你也可以使用bindkey来查看一下快捷键设置.

$ bindkey | grep _complete_help
"^Xh" _complete_help

^就是CTRL, 这个意思就是你要先按下CTRL-x, 然后再按下h
```


``` zsh
# 先输入ls空格, 然后按CTRL-x h
$ ls  
tags in context :completion::complete:ls::
    argument-rest options  (_arguments _ls)
tags in context :completion::complete:ls:argument-rest:
    globbed-files  (_files _arguments _ls)
          ｜                 |
          ｜                 + 函数调用栈，自右向左_ls调用了_arguments, _arguments调用了_files
          ＋－－－－：当前的tag 
```

这一串字符串就是输入ls后的Context

```zsh
context :completion::complete:ls::
context :completion::complete:ls:argument-rest:
```

#### Context的结构
> : completion : function : completer : [command|special_context] : argument : tag:
 - completion
 - function
 - completer
 - command 或 special_context
 - argument
 - tag

#### completer
每当系统需要在当前上下文处理补全的时候, 都会调用_main_complete, _main_complete会找到合适的completer处理这次补全. 可以指定多个completer依次处理, 直到完成补全或者全部失败.

 - 每个completer对应一个函数, 在completer名字前面加上下划线就是函数的名字
 - 一个上下文之下可以指定多个completer, 按照顺序依次调用completer函数, 若返回0则表示补全完成.否则继续调用后面的completer
 - 可以使用zstyle指定在什么样的上下文之下使用哪些completer



``` zsh
# 对所有的补全使用complete和correct两种补全
$ zstyle ':completion:*' completer _complete _correct


# 你也可以实现一个补全函数, 下面这个函数啥正事儿都没干, 但可以助你理解completer如何工作. 
$ function completer_spy() { print $argv; return -1; }
$ zstyle ':completion:*' completer completer_spy _complete _correct

```


completer有一下几种:
##### complete
这个为SpecialContext提供补全, 提供整个系统最基本的补全功能.


##### all_matches
这个一般作为第一个completer, 可以将所有的补全结果拼成一个字符串. 

##### approximate
通常配置在complete后面, 允许补全系统出错的时候进行一定数量的尝试, 你可以在context中看到
approximate-1这样的状态, 就表示approximate在尝试一次修正.

这个补全的代价是非常高的, 但是可以通过max-error限制尝试的次数:
``` zsh
$ zstyle ':completion:*:approximate:*' max-errors 1 numeric

# 优先使用path-directories (默认优先使用local-directories), tag-order后面可以配置多个tag, 
# 如果最后设置了-, 则表示只尝试已经配置的tags即可，不必再尝试其他的tags了。


$ zstyle ':completion:*:cd:*' tag-order path-directories
```

#### match
通常配置在complete后面, 

 - canonical_paths
 - cmdambivalent
 - cmdstring
 - correct
 - expand
 - expand_alias
 - extensions
 - external_pwds
 - history
 - ignored
 - list

 - menu
 - oldlist
 - precommand
 - prefix
 - user_expand

#### Special Contexts
 * -array-value- : The right hand side of an array-assignment (‘name=(...)’)
 * -brace-parameter- :The name of a parameter expansion within braces (‘${...}’)
 * -assign-parameter- :The name of a parameter in an assignment, i.e. on the left hand side of an ‘=’
 * -command- : A word in command position
 * -condition- : A word inside a condition (‘[[...]]’)
 * -default- :Any word for which no other completion is defined
 * -equal- : A word beginning with an equals sign
 * -first- :This is tried before any other completion function. The function called may set the _compskip parameter to one of various values: all: no further completion is attempted; a string containing the substring patterns: no pattern completion functions will be called; a string containing default: the function for the ‘-default-’ context will not be called, but functions defined for commands will be.
 * -math- : Inside mathematical contexts, such as ‘((...))’
 * -parameter- : The name of a parameter expansion (‘$...’)
 * -redirect- : The word after a redirection operator.
 * -subscript- : The contents of a parameter subscript.
 * -tilde- : 输入~之后
 * -value- : name=value之后

#### 使用_comps查找补全函数
当我们输入一条命令的时候, 补全系统会根据当前的上下文去调用指定的补全函数. 那么怎么查看呢?

``` zsh
# 查看ls的补全函数
$ print $_comps[ls]
_ls

# 想看看这个补全函数的源码?
$ builtin functions _ls
```

### Tag
Tag用来标识补全的对象是什么.

 - zsh相关
     - accounts: host
     - my-accounts
     - other-accounts
     - urls: url数据库中的url
     - all-expansions: 包含全部补全的字符串
     - arguments: 命令的参数
     - parameters: 参数名
     - options: 命令的选项
     - arrays: 数组
     - association-keys: 关联数组的key
     - indexes: 数组的下标
     - functions: 函数
     - history-words: 历史
     - jobs: 
     - modules: zsh模块
     - packages
     - original: 矫正器使用, 表示被纠正内容的原始值
     - path-directories:  cdpath中的目录
     - named-directories:
     - values: 列表里的某个值
     - strings
     - expansions: _expand用
     - styles: zstyle
     - widgets: zsh widgets
     - zsh-options: zsh options
     - builtins: 内置命令
     - characters: 
     - commands: 外部命令, 或者外部命令的子命令
     - contexts: zstyle专用
     - corrections: 可能的修正
     - default
     - descriptions: 表述信息的格式化定义
     - keymaps: zsh快捷键

 - 操作系统相关
     - devices: 设备名
     - all-files: 全部文件
     - files: 文件
     - other-files:
     - globbed-files
     - suffixes: 文件后缀名
     - file-descriptors: 文件描述符
     - directories: 目录
     - local-directories: 当前目录
     - directory-stack: pushd维护的目录栈
     - path-directories: 保存在cdpath中的路径
     - fstypes: 文件系统类型
     - users: 用户名
     - groups: 用户组
     - hosts: 
     - domains: 域名, 如ping命令之后
     - interfaces: 网络接口
     - libraries: 库
     - limits: 
     - manuals: man手册
     - printers: 打印队列
     - signals: 信号名
     - processes: pid
     - processes-names: pname
     - time-zones: 时区
     - ports: 端口号
     - paths: 

 - X相关
     - displays: XWndows
     - colormapids: XWindows的colormap中的color ID
     - cursors: X系统的光标名
     - colors: 颜色名
     - extensions: XWindows的扩展
     - fonts: XWindows的字体
     - keysyms: XWindows的keysyms
     - modifiers: 
     - visuals
     - windows

 - 其他
     - bookmarks:
     - email-plugin
     - mailboxes: 邮件目录
     - messages: 
     - warnings: 警告消息的格式
     - newsgroups
     - names
     - nicknames
     - targets: Makefile的target
     - maps
     - pods: perl pods
     - prefixes
     - sequences
     - sessions: zFTP sessions
     - types
     - variant
     - tags


### Style
Style用说明以什么样的方式完成补全.

 * complete-word
 * delete-char-or-list
 * expand-or-complete
 * expand-or-complete-prefix
 * list-choices
 * menu-complete
 * menu-expand-or-complete
 * reverse-menu-complete

### 补全插件(Completion Widget)

> zle -C complete expand-or-complete completer

 * complete-word
 * expand-or-complete
 * expand-or-complete-prefix
 * menu-complete
 * menu-expand-or-complete
 * reverse-menu-complete
 * list-choices
 * delete-char-or-list

### 编写补全脚本

#### Hello World
```zsh
# 启动一个干净的zsh
$ zsh.baby

# 初始化补全系统
$ autoload -Uz compinit

# 定义一个补全函数, 按照约定如果你针对foo命令编写的补全函数的名字应该是_foo
$ _foo() { _message "HELLO WORLD" }

# 把_foo这个补全函数加入到补全系统中
$ compdef _foo foo

# 输入foo空格，然后按下TAB, 什么都没有发生？
$ foo 
# 打开并设置补全系统的消息格式
$ zstyle ':completion:*:messages' format "%d"
# 再试一下
$ foo #按下TAB
HELLO WORLD

# 简单的补全参数 
$ _foo() { _arguments -h }
$ foo -
-h

# 为option添加描述
$ _foo() { _arguments {-h,--help}"[show help message]" }
$ foo -
--help  -h  -- show help message

# 指定如何补全option的argument
$ touch 1.log
$ _foo() { _arguments --out:"output .log file":'_files "*.log"'}
$ foo --out <TAB>
1.log

# 枚举
$ _foo() { _arguments --age="[The age]":ages:"(18 19 20)" }
$ foo --age=<TAB>
18 19 20

# 为arguments添加描述
$ _foo() { _arguments --age="[The age]":ages:"((18\:A 19\:B 20\:C))" }
$ foo --age=<TAB>
18  -- A
19  -- B
20  -- C

# 多个arguments
$ _foo() { _arguments --point:x:(1 2 3):y:(4 5 6) }
$ foo --point <TAB>
1 2 3
$ foo --point 1 <TAB>
4 5 6

```
#### 描述命令行的options和arguments

```
_argumengs spec1 spec2 ... specN
spec    = optspec optarg
optspec = -optname
        = +optname
        = -optname-
        = -optname=
        = -optname=-
        = optspec[optdesc]
optarg  = :message:action
        = ::message::action
        = :pattern:message:action
        = :pattern::message:action
        = :pattern:::maessage:action
action  = (item1 item2 ... itmeN)
        = ((item1\:desc1 item2\:desc2 ... itemN\:descN))
        = ->string
        = {eval-string}
        = = action
```

#### 补全系统使用的变量

 - curcontext
 - words
 - state
 - state_descr
 - line
 - expl
 - opt_args
 - CURRENT
 - IPREFIX
 - ISUFFIX
 - PREFIX
#### _arguments
> OPTSPEC

| optspec | 功能 |  范例   |
|---------|-----|-----|
| -optname   |  -o   |  -o [arg]  |
| -optname-  |  -o-  |  -o[arg]   |
| -optname+  |  -o+  |  -o [arg] 或 -o[arg] |
| -optname=  |  -o=  |  -o [arg] 或 -o=[arg] |
| -optname=- |  -o=- |  -o=[arg] |

> OPTSPEC[DESC]
 - 补全系统的verbose设置为true的时候, 会显示DESC, 用来描述参数的功能

> OPTARG

| optarg | 功能 |  范例   |
|---------|-----|-----|
|  :message:action |
| ::message:action |
| :*pattern:message:action |
| :*pattern::message:action |
| :*pattern:::message:action |

### 配置补全系统
### 系统自带的补全函数在哪?

# 参考
 - http://zsh.sourceforge.net/Doc/Release/Completion-System.html
 - http://zsh.sourceforge.net/Doc/Release/Completion-Widgets.html
 - whence -v function-name : 查看函数的加载路径