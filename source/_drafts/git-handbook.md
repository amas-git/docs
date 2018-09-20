---
title: GIT HANDBOOK 
tags:
---
# Git =
一种分布式版本控制系统(DCVS)
# 仓库 ==
# 安装 ==
# Linux/Unix ===
 * Ubuntu
```
#!sh
$ sudo apt-get build-dep git-core git-doc
```
 * Arch
```
#!sh
$ sudo pacman -S git
```
# Mac OS  ===
```
#!sh
$ sudo port install git-core +svn +doc
```
# Windows ===
 * Git on Cygwin
  * The “official” way to run Git in Windows is to use Cygwin
 * Git on MSys
# 配置Git ==
# 全局设定 "--global" ===
```
#!sh
# 配置个人信息
$ git config --global user.name "amas"
$ git config --global user.email "zhoujb.cn@gmail.com"
# 查看Git全局设定
$ git config --global --list                                                          
user.name=amas
user.email=zhoujb.cn@gmail.com
# 开启代码高亮
$ git config --global color.ui "auto"
```
# 使用GUI前端 ==
Git自带Tcl/Tk GUI前端，所以你必须先安装[wiki:Tcl/Tk]
```
#!sh
# git gui提供基本的Git操作，不提供仓库历史信息查询
$ git gui
# gitk 可以查看仓库的历史信息
$ gitk
# 查看全部分支的历史信息
$ gitk --all 
```
# 常用操作 ==
# 1. 建立仓库 : git init ===
```
#!sh
$ mkdir my-project                                                                
% cd my-project
% git init
Initialized empty Git repository in /home/amas/src/git/my-project/.git/
```
# 2. 添加文件到仓库中 : git add ===
假使有项目文件index.html:
```
#!html
<html>
<body>
    <h1>Hello World!</h1>
</body>
</html>
```
```
#!sh
# 添加index.html到仓库中
$ git add index.html
```
# 3. 提交变动到仓库: git commit -m "some-comment..." ===
每次commit意味着你的代码向前演进了一点。
```
#!sh
$ git commit -m "add index.html"
[master (root-commit) 7f59e70] add index.html
 1 files changed, 5 insertions(+), 0 deletions(-)
 create mode 100644 index.html
```
# 4. 查看历史记录 git log ===
```
$ git log 
commit 7f59e70b88bdcf7bd2eefa06bf035e462921250d
Author: amas <zhoujb.cn@gmail.com>
Date:   Sat Jul 31 14:20:50 2010 +0800
    add index.html
```
如果历史记录较多可用-n查看最近的n次变更记录
```
# 类似于svn log --limit 10
$ git log -10
```
快捷键:
|| q || 退出 ||
|| / || 查找 ||
```
#!div class=note
`为什么使用SHA-1 Hash 而非数字来标识仓库的版本?`
CVS或SVN使用 CentralizedRepository，通常用版本号作为每次变更的唯一标识，但是Git这样的 DistributedRepository 显然不能使用这种方法，比如版本号18在不同人的仓库中可能标识了不同的变更，因此Git需要其他的方法来表示版本变更，实际上，Git将使用
仓库的[wiki:Metadata]通过[wiki:SHA-1]计算出的Hash结果作为版本号。
```
# 5. 提交你的修改: git add 或 `git commit -a` ===
对index.html稍作修改:
```
#!diff
--- a/index.html
+++ b/index.html
@@ -1,5 +1,6 @@
 <html>
 <body>
     <h1>Hello World!</h1>
+    <h1>I'm amas!</h1>
 </body>
 </html>
```
使用`git status`查看发生的改变:
```
#!sh
$ git status
# On branch master
# Changed but not updated:
#   (use "git add <file>..." to update what will be committed)
#   (use "git checkout -- <file>..." to discard changes in working directory)
#
#       modified:   index.html
#
no changes added to commit (use "git add" and/or "git commit -a")
# Git发现你已经修改了index.html但它不知道改怎么处理，所以仅仅是列出被修改的文件。
$ git add index.html
$ git status
# On branch master
# Changes to be committed:
#   (use "git reset HEAD <file>..." to unstage)
#
#       modified:   index.html
#
$ git commit -m "add another header"
[master a5abe5e] add another header
 1 files changed, 1 insertions(+), 0 deletions(-)
# 再查看下历史
$ git log
commit a5abe5ef6a65390cc80c81f896cfd50278679582
Author: amas <zhoujb.cn@gmail.com>
Date:   Sat Jul 31 17:36:18 2010 +0800
    add another header
commit 7f59e70b88bdcf7bd2eefa06bf035e462921250d
Author: amas <zhoujb.cn@gmail.com>
Date:   Sat Jul 31 14:20:50 2010 +0800
    add index.html 
```
```
#!div class=note
提交记录中写什么好?
 1. 不要写你改了什么
 2. 写你为什么要修改
```
在Git中代码可能保存在三处:
 1. 源代码文件(废话)
 2. Staging area(是一个buffer), 可以使用它存放你想要提交的修改
 3. Adding Files
# 交互模式 git add -i ====
```
#!sh
$ git add -i                                                           
           staged     unstaged path
  1:    unchanged        +1/-0 index.html
*** Commands ***
  1: status       2: update       3: revert       4: add untracked
  5: patch        6: diff         7: quit         8: help
What now> 
# 2 update 
What now> 2
           staged     unstaged path
  1:    unchanged        +1/-0 index.html
# 进入update命令, stage 1, 输入1
Update>> 1
           staged     unstaged path
* 1:    unchanged        +1/-0 index.html
Update>> 
updated one path
# 再来查看状态
What now> 1
           staged     unstaged path
  1:        +1/-0      nothing index.html
*** Commands ***
  1: status       2: update       3: revert       4: add untracked
  5: patch        6: diff         7: quit         8: help
# unstaged 已经没有了, 显示为`nothing`
# staged 多了1个文件, 显示为`+1/-0`
# 这时候如果运行`git diff`你将什么都看不到，因为这次修改已经stage了，与commit不同，这种修改并没有历史记录，如果你决定纳入产品，就要commit,并加点儿提交信息
# 你可以认为这些修改被暂时缓存起来了，之后你可以选择提交或是放弃
```
 1. status 查看工作目录的状态，实际上就是`git add -i`刚进来那屏
 2. update stage 某个文件
 3. revert unstage 某个文件
 6. diff 比较修改
 7. quit  退出
# git commit ====
你可以指定用哪个编辑器编辑commit log:
 1. GIT_EDITOR environment variable.
 2. core.editor Git configuration value.
 3. VISUAL environment variable.
 4. EDITOR environment variable.
 5. Git tries vi if nothing else is set.
```
#!div class=note
假如你用惯了svn, 你可以添加shotcut:
```
#!sh
$ git config --global alias.ci "commit"
$ git config --global alias.st "status"
```
```
# git diff ===
```
#!sh
# 查看working tree
$ git diff 
# 查看staging area
$ git diff --cached
# 跟某个分支比较
$ git diff HEAD
```
# 不需要git cp ===
`svn cp`命令用于建立分支或标记Tag, git有专门的命令干这些事情，所以没有
# 让Git忽略特定的文件: .gitignore ===
比如对.csv和.swp文件不予以管理，需要将下面的内容加到`.gitignore`或`.git/info/exclude`中
```
*.csv
*.swp
```
如果这种限制仅仅是你个人的，请保存到`.git/info/exclude`中，否则
保存到`.gitignore`中，并提交
# 改名: git mv <origin-file> <new-file> ===
```
#!sh
$ git mv index.html index.htm
```
# 使用分支 ===
# Create new branches ====
```
#!sh
# 基于当前分支，建立新的new-branch分支
$ git branch new-branch
# 切换到new-branch分支
$ git checkout new-branch
# 建立分支，并马上切换,可用一条命令完成
$ git checkout -b new-branch master
Switched to a new branch 'new-branch'
```
关于创建分支，它不是一门科学，更像是艺术，一下是一些经验
 1. Experimental changes:
  * 算法优化
  * 重构一部分代码,弄成某种模式
 2. New features:
  * 开始加入新的功能
 3. Bug fixes:
  * 代码已经加了release的tag,但发现了bug
# Merge changes between branches ====
# Straight merges  =====
take the history of one branch and the history of the other branch and attempt to merge them together.
这个也是比较常用的一种merge,将分支A的全部历史Merge到分支B
```
$ git checkout master
$ git merge BA 
```
# Squashed commits =====
take the history of one branch and compress or “squash”—it into one commit on top of another branch.
(FeatureBranch)经常在新的分支上开发新功能，开发完毕后，我们想merge到主线上，但是在分支上修正的bug我们并不关心，这个时候需要进行Squashed Merge, 简单说，其实就是将分支首版本merge到master上，然后再手工提交。
```
#!sh
$ git merge --squash B_new_feature
$ git status
...
$ git commit -m "add new feature"
...
```
# Cherry-picking a commit =====
pulls a single commit from a different branch and applies it to the current branch.
有时候你只是需要把分支A上的某次提交Merge到分支B上，这就叫Cherry-Picking(挑樱桃,够形象吧?)
```
#!sh
# 在new-branch上操作
$ git commit -m "add n1"                                          
[new-branch 4f04430] add n1
 0 files changed, 0 insertions(+), 0 deletions(-)
 create mode 100644 n1
$ git commit -m "add n2"
[new-branch 0aceb5f] add n2
 0 files changed, 0 insertions(+), 0 deletions(-)
 create mode 100644 n2
# 切到master branch 
$ git checkout master
# 我们只需要摘n1那颗樱桃 
$ git cherry-pick 4f04430
...
# 需要多个樱桃怎么办?
$ git cherry-pick -n 4f04430
$ git cherry-pick -n ...
$ git status
$ git commit
...
```
# Handle conflicts ====
做Merge怎么可能不遇到代码冲突呢?
先来制造冲突:
 1. git branch A 
  * 添加文件about,第一行xxx
 2. git branch B A
  * 修改文件A,第一行yyy,提交修改
 3. git checkout about ; git merge B
出现冲突了，大概是这个样子
```
$ git merge B                                                   
Auto-merging about
CONFLICT (content): Merge conflict in about
Automatic merge failed; fix conflicts and then commit the result.
$ cat about
<<<<<<< HEAD
xxx
=======
yyy
>>>>>>> B
# 解决冲突, git会帮你在系统中找merge工具，你需要用工具解决冲突，就这么简单
$ git mergetool
```
# Delete branches ====
```
#!sh
# 如果分支中包含没有merge到master的修改则删除失败
$ git branch -d <branch-name>
# 强制删除分支
$ git branch -D <branch-name>
```
# Rename branches ====
```
#!sh
# 将master分支改名为my-master
$ git branch -m master my-master
$ git branch
  REL_1.0
* my-master # 当前工作分支
```
# Handle release branches ====
从[wiki:MasterBranch master branch]上创建了RB_1.0 branch.
```
#!sh 
$ git branch RB_1.0 master
# RB : Release Branch
```
建立了这个Branch后，当时仓库的状态就被保存下来，你可以继续开发了，
直到你觉得RB_1.0才是值得发布的后，你需要仓库暂时回到RB_1.0的状态
```
#!sh
# 如果你是zsh用户，狂按Tab就行了
$ git checkout RB_1.0
```
# 使用Tag: `git tag <tag-name> <branch-name?` ===
```
#!sh
$ git tag 1.0 RB_1.0
```
# 打包仓库: `git archive` ===
# 怎样分析仓库的历史记录: git log ===
```
$ git log --since="5 hours"
# 某个版本到HEAD之间的历史记录
$ git log 6ae6271..HEAD
$ git log 6ae6271..
# 自定义历史记录的格式
% git log --pretty=format:"%h %s" 1.0..
e751e68 add file b
f12ea96 add n1
...
$ git log --pretty=oneline 1.0..                   
e751e68a586933cf4dd99369bf2dbb67e22fb821 add file b
f12ea9678d69570fe67148405af409e3b52ec419 add n1
...
# ^ 意思是 - ^^表示后退两个版本，以此类推, 太多了使用~N, 参看下面的例子
$ git log --pretty=oneline 1.0..HEAD^
f12ea9678d69570fe67148405af409e3b52ec419 add n1
...
# ~N 意思是之前的第N个版本
% git log --pretty=oneline HEAD~2..
e751e68a586933cf4dd99369bf2dbb67e22fb821 add file b
f12ea9678d69570fe67148405af409e3b52ec419 add n1
```
# git blame <file-name>  ===
如果你想看看某个文件的每行都是由谁在哪个版本上改动的，可以用git blame 
```
#!sh
% git blame index.htm 
^7f59e70 index.html (amas 2010-07-31 14:20:50 +0800  1) <html>
25ab1ca6 index.html (amas 2010-07-31 18:14:07 +0800  2) <head>Release 1.0</head>
a8723955 index.html (amas 2010-07-31 17:54:31 +0800  3) <head>AMAS</head>
...
# ^ 表示最早的历史，一般是第一次提交这个文件时就有的内容
# 查看某个区间内的变动
$ git blame -L 1,5  index.htm
$ git blame -L 3,1  index.htm
$ git blame -L 3,-1 index.htm
# 按正则表达式匹配特定的行
$ git blame -L "/html/",+2 index.htm
$ git blame -L "/html/",+2 HEAD^^ -- index.htm
```
# Following Content ===
Git
# 合并分支: `git rebase <branch-name>` ===
```
#!sh
# 假设有RB_1.0分支，需要merge到master上
$ git checkout master 
$ git rebase RB_1.0
```
如果出现冲突，你有如下三个选择:
 * `git rebase --continue` 解决冲突后运行此命令继续合并
 * `git rebase --skip` 干脆跳过冲突，不接受产生冲突的patch
 * `git rebase --abort` 算了，不merge了
# 与中心仓库的同步 ==
git pull 到你的本地master分支，这个命令实际上是两个命令的组合，分别是git fetch origin和git merge origin/master, 在你需要更新本地代码时执行这个操作。
```
#!sh
$ git pull origin master
```
git push 到中心仓库，将你的代码提交到中心仓库
```
#!sh
$ git push origin master
```
当中心仓库新建了分支，你需要首先同步中心仓库的所有分支到本地
```
#!sh
$ git fetch
# 或
$ git fetch origin remote_branch_name:local_branch_name
```
# Git Server ==
 * Gitosis
 * Gitolite : http://www.mmtek.com/dp20090929/node/58
# 客户端 ==
 * gitk
 * tig
# 使用bundle打包仓库
假设你想将仓库A的改动打包成一个文件，发送给你的伙伴，可以这样:
```sh
$ git bundle create /tmp/b-2012 HEAD
```
你的同事可以使用pull命令将这些数据拉到它的仓库中:
```sh
$ git pull ~/b-2012
```
# git log
```
$ git log --stat --author=amas
```
# Reference ==
 * http://code.google.com/p/msysgit/
 * http://gitready.com/
 * http://library.edgecase.com/git_immersion/ (非常不错的教程)

# Git Internal =
Git的设计观念更接近于文件系统，而非传统的VCS。
# Git对象 ==
 1. Blob
 2. Tree
 3. Commit
 4. tree-ish
# 使用Git ==
# 初始配置 ===
```
#!sh
$ git config --global user.name “Scott Chacon”
$ git config --global user.email “schacon@gmail.com”
# 查看配置
$ cat ~/.gitconfig
$ git config -l
```
如果你想对不通的项目使用不通的配置，你可以在Git仓库下执行不带'--global'参数的'git config'命令， 这些配置将保存在'.git/config'文件中。
# 创建仓库 ===
# 查看Log ===
```
#!sh
$ git log --pretty=online
```
--pretty:
 * oneline
 * short
 * medium
 * full
 * fuller
 * email
 * raw
 * format:"string"
# git show ==
git show 这命令可以将git内部的对象翻译成人类的语言， 如果你开了高亮，看起来就更顺眼了。
```
#!sh
# git show 默认查看最近一次的变更记录
$ git show
commit 8c573ceb529b12f2ee363951ce7a98779fb406db
Author: amas <zhoujb.cn@gmail.com>
Date:   Wed Mar 23 10:28:17 2011 +0800
    message
...
# Show tree-ish object, 查看最近一次的前一次改变了什么
$ git show master^
```
# git  instaweb  ==
如果你在本地部署了WebServer, 你可以通过他们来直接查看git仓库。
```
#!sh
# 使用apache浏览仓库
$ git instaweb --httpd httpd
# 停止浏览
$ git instaweb --stop
```
# git grep ==
```
#!sh
$ git grep -n TODO
```
# git diff ==
```
#!sh
```
# 生成补丁: git diff > xxx.patch ==
```
#!sh
# 生成补丁
$ git diff > xxx.patch
# 打补丁
$ patch -p1 < xxx.patch
```
# git branch ==
```
#!sh
# 建立dev分支
$ git branch dev
# 查看有哪些分支
$ git brach
* master
* dev
# 我们来看看.git/refs/heads目录:
$ ls .git/refs/heads
master dev
# 看看refs/heads下文件的具体内容:
$ cat .git/refs/heads/master
886c3d0d14523acd3c9f88fe3a9f4fbee75508bc
$ cat .git/refs/heads/dev
886c3d0d14523acd3c9f88fe3a9f4fbee75508bc
# 内容完全一样， git建立分支的成本就只有这40个字节
# 切换到dev分支
$ git checkout dev
```
通常我们更喜欢建立分支，而后切换到新建立分支上开始工作，可以用如下命令:
```
$ git checkout -b newbranch
```
# git merge ==
# git reset ==
```
#!sh
# 撤销当前所有修改:
$ git reset --hard HEAD
# 插销最近一次提交
$ git reset --hard origin/HEAD
```
如果要使代码恢复到更早的版本需要使用git revert, 但是那很危险， 明白之前慎用!!!
# git rebase ==
git rebase的使用场景是这样的:
```
master: M1-M2-M3--M4
dev:       +--D1--D2
```
dev经过一段开发后，dev已经可以brach off, 此时，通常的做法是
将dev合并到master上。 这样连同dev分支上的历史记录也一同合并到master上了。
有时，为了保持master上历史记录的清晰，我们并不需要这些log, 此时可以使用rebase.
不需要到master上，在dev上执行`git rebase master`

# git stash — 暂存临时代码
当你不想提交当前完成了一半的代码，但是却不得不修改一个紧急Bug，那么使用`git stash`就可以将你当前未提交到本地(和服务器)的代码推入到Git的栈中，这时候你的工作区间和上一次提交的内容是完全一样的，所以你可以放心的修 Bug，等到修完Bug，提交到服务器上后，再使用`git stash apply`将以前一半的工作应用回来。
例：
当有一些修改，还没有提交时：
```
#!sh
develop> git stash
```
切换到另一分支上做review：
```
#!sh
master> {review}
```
再切换到原来的分支上通过git stash apply继续工作：
```
#!sh
develop>git stash apply
```
当git  stash多次的时候，你的栈里将充满了未提交的代码，这时候你会对将哪个版本应用回来有些困惑，`git stash list`命令可以将当前的Git栈信息打印出来，你只需要将找到对应的版本号，例如使用`git stash apply stash@!{1}`就可以将你指定版本号为stash@!{1}的工作取出来，当你将所有的栈都应用回来的时候，可以使用 `git stash clear`来将栈清空。
例：
如果有多次`git stash`，通过`git stash  list `列出Git栈信息：
```
#!sh
develop>git  stash  list
stash@{0}: WIP on feature/push: 873b0e7 Merge branch 'develop' of 192.168.48.12:pansi into develop
stash@{1}: WIP on feature/push: 873b0e7 Merge branch 'develop' of 192.168.48.12:pansi into develop
stash@{2}: WIP on feature/log.upload: d06f040 tmp
```
如果想返回`stash@!{1}`版本：
```
#!sh
develop> git stash apply stash@{1}
```
如果想清空Git栈信息（慎用），则
```
#!sh
develop> git  stash clear
```

