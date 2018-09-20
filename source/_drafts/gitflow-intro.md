---
title: gitflow_intro
tags:
---
# Git Flow =
# 安装 
# 从网络安装
```
#!sh
$ wget -q – http://github.com/nvie/gitflow/raw/develop/contrib/gitflow-installer.sh –no-check-certificate
```
# 从源码安装
```
#!sh
$ git clone --recursive git://github.com/nvie/gitflow.git
$ cd gitflow 
$ sudo make install
```
# 安装Shell补全插件
```
#!sh
$ git clone http://github.com/bobthecow/git-flow-completion
$ cd git-flow-completion ; ls -1
git-flow-completion.bash
git-flow-completion.zsh
LICENSE
README.markdown
```
如需要，请详细阅读`README.markdown`
其中有两种补全插件，`.bash`为bash提供补全，`.zsh`为Zsh补全， 安装方法非常简单， 以后者为例， 你需要在自己的`~/.zshrc`中引入补全插件:
```
...
source ~/.zsh.d/module/git-flow-completion.zsh
...
```
# GitFlow模型介绍
在GitFlow中，一切被划分为分支。 当你开始一个新特性的时候，你会基于develop分离出一个新的分支。 如果你在进行hotfix, 那么你是从master上分离的，理解GitFlow模型，关键在于明白某个分支是从哪个分支分离出来，最终应该合并到哪些分支去。
[[Image(git-flow-overview.jpg)]]
master分支:
 * 总是发布稳定版本, 这反应了软件的最新/最稳定的状态
develop 分支:
 * 用于集成已完成的特性
 * 用于集成来自针对master分支的hostfix
 * CIServer需要真对此分支进行日构建/运行测试/
临时分支:
 * 多人开发时可能需要在远程临时共享某些分支
# 新特性开发: Feature分支
Feature branches are branches created to do new feature development and should be branched off the develop branch.  
 1. 用于开发新功能
 2. 只能从develop分支拉出
 3. 当特性开发完毕，需要从feature分支合并回归到develop分支， 同时在合并后删除feature分支
```
#!div class=note
# 建议 ===
 * 尽量将单一且完整的功能分配到一个feature分支中
 * 个人或小组在一天之内便可完成
```
# 集成: Release分支
当 develop  分支上积累了足够多的待发布功能时， 可能需要做一次发布。 此时， 可以在release分支上进行.
 * 拉出release分支后，此分支不再集成新的功能，进入ReleaseHardening阶段， 只接受bugfix, 直到足够稳定。 
 * QA Team需要针对此分支开展各种测试，以便确认产品质量达到发布标准.
 * 发布完成后，release分支应当合并到master分支
 * 发布完成后，release分支应当合并到develop分支，保证所有的bugfix都同步到develop分支.
 * 最后，删除release分支
```
#!div class=note
# 建议 ===
 * ReleaseNote可从该分支生成，应尽量自动化。
 * 当开发人员修复bug后，应当设法通知UAT(User Acceptable Testing)Team
```
# 紧急修复: Hotfix分支
实际中master分支上的产品也可能不尽稳定，有时会有一些影响使用的严重bug, 也有可能实际交付的产品某些细节并不令用户满意， 凡此种种我们都需要修改master分支上的代码， 为了保证master分支总是记录了稳定发布版， 我们需要使用bugfix分支来进行实际的修复。
 * 问题修复后hotfix分支的代码需要同时提交到master分支以及develop分支
 * 合并后的master分支，将标记版本号，格式为:`version.hotfix`
# 产品支持: Support分支
这个support分支，实际上是从master分支上来的, 用于修复bug, 但是bug修复后只需要同步会master, 并不需要同步到develop分支。
比如: 你当前软件的版本为6.0, 但是仍有用户使用2.1, 为了继续维护该版本， 所有2.1的bugfix需要进行修复， 这一点类似于hotfix, 但是
因为版本2.1与当前版本可能差异巨大， 也许架构上完全不同，因为hotfix总是需要向master和develop线合并的, 而此时并不需要向develop分支提交这部分代码，故而hotfix并不适用，
所以便有了support分支。 你可以认为它是一种特殊的hotfix分支。
# Git Flow 命令
# git flow init
```
#!sh
$ git clone git@localhost:oops.git && cd oops
master> cat .git/config
[core]
        repositoryformatversion = 0
        filemode = true
        bare = false
        logallrefupdates = true
[remote "origin"]
        fetch = +refs/heads/*:refs/remotes/origin/*
        url = git@localhost:oops.git
[branch "master"]
        remote = origin
        merge = refs/heads/master
# 初始化gitflow仓库
master> git flow init
Which branch should be used for bringing forth production releases?
   - master
Branch name for production releases: [master] 
Branch name for "next release" development: [develop] 
How to name your supporting branch prefixes?
Feature branches? [feature/] 
Release branches? [release/] 
Hotfix branches? [hotfix/] 
Support branches? [support/] 
Version tag prefix? []  
develop> cat .git/config
[core]
        repositoryformatversion = 0
        filemode = true
        bare = false
        logallrefupdates = true
[remote "origin"]
        fetch = +refs/heads/*:refs/remotes/origin/*
        url = git@localhost:oops.git
[branch "master"]
        remote = origin
        merge = refs/heads/master
[gitflow "branch"]
        master = master
        develop = develop
[gitflow "prefix"]
        feature = feature/
        release = release/
        hotfix = hotfix/
        support = support/
        versiontag = 
```
 * 初始化完毕后，你就会在develop分支， 请注意保持这个分支总是有整洁的记录
如果你需要切换到master分支， 可以:
```
#!sh
$ git checkout master
```
你可以使用如下命令保持develop分支为最新:
```
#!sh
develop> git pull origin develop
```
# git flow feature
开始一个新的feature:
```
#!sh
develop> git flow feature start F1
Switched to a new branch 'feature/F1'
Summary of actions:
- A new branch 'feature/F1' was created, based on 'develop'
- You are now on branch 'feature/F1'
Now, start committing on your feature. When done, use:
     git flow feature finish F1
# 来看看现在本地有哪些分支
feature/F1> git branch
  develop
* feature/F1
  master
```
接下来，我们在新的Feature分支中进行正常开发.
```
#!sh
feature/F1> echo "add apple" >> F1
feature/F1> git add -A && git commit -am "F1: add apple"
[feature/F1 d55703e] F1: add apple
 1 files changed, 1 insertions(+), 0 deletions(-)
 create mode 100644 F1
feature/F1> echo "add bus"   >> F1
feature/F1> git add -A && git commit -am "F1: add bus"
[feature/F1 07f6b56] F1: add bus
 1 files changed, 1 insertions(+), 0 deletions(-)
feature/F1> echo "add coco"  >> F1
feature/F1> git add -A && git commit -am "F1: add coco"
[feature/F1 1617f6c] F1: add coco
 1 files changed, 1 insertions(+), 0 deletions(-)
```
```
git flow feature publish [Branch Name]
```
可能会有众多成员参与到新特性的开发中，如是，你可能需要在remote上共享feature分支给其他成员。
```
#!sh
$ git flow feature publish F1
Counting objects: 10, done.
Delta compression using up to 2 threads.
Compressing objects: 100% (6/6), done.
Writing objects: 100% (9/9), 792 bytes, done.
Total 9 (delta 0), reused 0 (delta 0)
To git@localhost:oops.git
 * [new branch]      feature/F1 -> feature/F1
Already on 'feature/F1'
Summary of actions:
- A new remote branch 'feature/F1' was created
- The local branch 'feature/F1' was configured to track the remote branch
- You are now on branch 'feature/F1'
```
```
#!div class=note
# 此条命令，大致等同于:
```
#!sh
$ git push origin feature/F1
$ git config “branch.feature/F1.remote” “origin”
$ git config “branch.feature/F1.merge”  “refs/heads/featureAdd”
$ git checkout “feature/F1”
```
```
```
#!div class=warn
# 尽量不要使用`git pull`
git pull默认会将远程所有对象拉到本地，这可能会给多分支开发带来麻烦，所以使用GitFlow时应当遵循如下原则:
 1. 尽量深入了解并使用gitflow提供的命令
 2. 当gitflow命令无法满足需要，使用影响范围尽可能小的命令
比如: 同步feature/F1分支时，尽可能不要使用git pull, 尽量使用:
```
#!sh
$ git pull origin feature/F1
```
```
需要在共享分支工作的人员，可通过如下命令下载:
```
#!sh
develop> git flow feature pull origin F1
Created local branch feature/F1 based on origin's feature/F1.
feature/F1>  
# 或者
develop> git flow feature track F1
```
user2:
```
#!sh
feature/F1> echo "add dog" >> F1
feature/F1> git add -A && git commit -am "F1: add dog"
[feature/F1 10e6d24] F1: add dog
 1 files changed, 1 insertions(+), 0 deletions(-)
feature/F1> echo "add egg" >> F1
feature/F1> git add -A && git commit -am "F1: add egg"
[feature/F1 e079ca9] F1: add egg
 1 files changed, 1 insertions(+), 0 deletions(-)
feature/F1> git push origin feature/F1
Counting objects: 8, done.
Delta compression using up to 2 threads.
Compressing objects: 100% (4/4), done.
Writing objects: 100% (6/6), 562 bytes, done.
Total 6 (delta 0), reused 0 (delta 0)
To git@localhost:oops.git
   1617f6c..e079ca9  feature/F1 -> feature/F
```
user1:
```
#!sh
feature/F1> echo "add flower"  >> F1 
feature/F1> git add -A && git commit -am "F1: add flower" 
feature/F1> git pull origin feature/F1
remote: Counting objects: 8, done.
remote: Compressing objects: 100% (4/4), done.
remote: Total 6 (delta 0), reused 0 (delta 0)
Unpacking objects: 100% (6/6), done.
From localhost:oops
 * branch            feature/F1 -> FETCH_HEAD
Auto-merging F1
CONFLICT (content): Merge conflict in F1
Automatic merge failed; fix conflicts and then commit the result.
# 发生冲突，人工化解
feature/F1> git mergetools
feature/F1> git commit -am "F1: fixed conflict"
[feature/F1 146010d] F1: fixed conflict
```
结束feature分支开发:
```
#!sh
# 1. 首先同步develop分支， 为之后的合并做准备
feature/F1> git pull origin develop
# 2. 结束feature分支开发
feature/F1> git flow feature finish
Branches 'feature/F1' and 'origin/feature/F1' have diverged.
And local branch 'feature/F1' is ahead of 'origin/feature/F1'.
Switched to branch 'develop'
Merge made by recursive.
 F1 |    6 ++++++
 1 files changed, 6 insertions(+), 0 deletions(-)
 create mode 100644 F1
warning: not deleting branch 'feature/F1' that is not yet merged to
         'refs/remotes/origin/feature/F1', even though it is merged to HEAD.
error: The branch 'feature/F1' is not fully merged.
If you are sure you want to delete it, run 'git branch -D feature/F1'.
Summary of actions:
- The feature branch 'feature/F1' was merged into 'develop'
- Feature branch 'feature/F1' has been removed
- You are now on branch 'develop'
develop> git branch
* develop
  master
```
相当于:
```
git checkout develop
git merge –no-ff “ReadmeAdd”
git branch -d ReadmeAdd
```
如果遇到冲突， GitFlow会帮你切换到feature/F1分支下， 此时你需要解决冲突， 之后再次使用`git feature finish`
```
#!div class=note
# 别忘记删除服务器上的feature分支
```
#!sh
develop> git push origin :feature/F1
To git@localhost:oops.git
 - [deleted]         feature/F1
# 删除本地分支
develop> git branch -D feature/F1
# 查看远程分支
develop> git remote show origin
```
```
# git flow release
```
#!sh
release/V1.0> git flow release finish -F [-p] V1.0
# 你需要填写版本信息, 作为release的tag, 这个tag可以从master上看到
Switched to branch 'develop'
Merge made by recursive.
 F1 |    1 +
 1 files changed, 1 insertions(+), 0 deletions(-)
warning: not deleting branch 'release/V1.0' that is not yet merged to
         'refs/remotes/origin/release/V1.0', even though it is merged to HEAD.
error: The branch 'release/V1.0' is not fully merged.
If you are sure you want to delete it, run 'git branch -D release/V1.0'.
Summary of actions:
- Latest objects have been fetched from 'origin'
- Release branch has been merged into 'master'
- The release was tagged 'V1.0'
- Release branch has been back-merged into 'develop'
- Release branch 'release/V1.0' has been deleted
```
This will bring up a screen where you’ll be expected to type in some final notes for this release branch.  After that the following will happen:
```
git checkout master
git fetch origin master
git merge –no-ff Version_1.0
git tag -a Version_1.0
git push origin master
git checkout develop
git fetch origin develop
git merge –no-ff Version_1.0
git push origin develop
git branch –d Version_1.0
```
||=参数=||= 说明 =|| 
||=-F=||执行操作前先执行fetch      ||
||=-s=||对新建的tag签名            ||
||=-u=||签名使用的GPG-key          ||
||=-m=||使用指定的注释作为tag的注释  ||
||=-p=||当操作结束后,push到远程仓库中||
||=-k=||保留分支                   ||
||=-n=||不创建tag                  ||
As you can see there are quite a few things that is done here.  To explain this simply you can read the following list:
 1. 将origin拉到本地
 2. Release分支merge到master分支
 3. Release分支被标记为Ver1.0
 4. Release分支被merge回develop分支
 5. release/Ver1.0分支被删除
 6. develop/ master分支 / Tags 被推到origin
 7. origin/release/Ver1.0分支被删除
# git flow hotfix
类似于git flow feature
 * git flow hotfix start
 * git flow hotfix start [hotfixName]
 * git flow hotfix start Version_1.1
 * git flow hotfix finish
 * git flow hotfix finish Version_1.1
# git flow support
# 场景 ==
||=开发人员A=||= 远程仓库git@localhost:oops.git =||=开发人员B=||
```td#----------------------------------[step 1]
```
```td
||=165e957=||
```
#!sh
# git@localhost:oops.git
$ git branch
  develop
* master
```
```
```td
```
|-----------------------------------------
```td#----------------------------------[step 2]
```
#!sh
# 1. clone远程仓库
$ git clone git@localhost:oops.git
# 2. 初始化gitflow
$ git flow init
Which branch should be used for bringing forth production releases?
   - master
Branch name for production releases: [master] 
Branch name for "next release" development: [develop] 
How to name your supporting branch prefixes?
Feature branches? [feature/] 
Release branches? [release/] 
Hotfix branches? [hotfix/] 
Support branches? [support/] 
Version tag prefix? [] 
master>  git pull origin develop # optional
# 3. 我们在develop分支上
develop> git branch
* develop
  master
```
```
```td
-
```
```td
```
#!sh
$ git clone git@localhost:oops.git
```
```
|-----------------------------------------
```td#----------------------------------[step 3]
```sh
# 1. 建立新的开发分支
develop> git flow feature start F1
Branches 'develop' and 'origin/develop' have diverged.
And local branch 'develop' is ahead of 'origin/develop'.
Switched to a new branch 'feature/F1'
Summary of actions:
- A new branch 'feature/F1' was created, based on 'develop'
- You are now on branch 'feature/F1'
Now, start committing on your feature. When done, use:
     git flow feature finish F1
```
```sh
# 2. 完成这个功能(make-some-commit F1 为一次提交)
feature/F1> make-some-commit F1
...
feature/F1> make-some-commit F1
...
feature/F1> make-some-commit F1
```
```sh
# 3. 结束F1的开发
feature/F1> git flow feature finish F1
Branches 'develop' and 'origin/develop' have diverged.
And local branch 'develop' is ahead of 'origin/develop'.
Switched to branch 'develop'
Merge made by recursive.
 F1 |    3 +++
 1 files changed, 3 insertions(+), 0 deletions(-)
 create mode 100644 F1
Deleted branch feature/F1 (was 360bfdc).
Summary of actions:
- The feature branch 'feature/F1' was merged into 'develop'
- Feature branch 'feature/F1' has been removed
- You are now on branch 'develop'
```
```
#!sh
# 我们又回到"develop"分支上了
develop> tig
2011-09-06 11:25 A               M─┐ [develop] Merge branch 'feature/F1' into develop    
2011-09-06 11:22 A               │ o [F1] make some change 17016:2742                    
2011-09-06 11:22 A               │ o [F1] make some change 17016:2741                    
2011-09-06 11:22 A               │ o [F1] make some change 17016:2738                    
2011-09-06 11:19 A               I
```
||=0bf42a5=||
```
```td
-
```
```td
```
|-----------------------------------------
```td#----------------------------------[step 4]
推到远程仓库上:
```
#!sh
$ git pull origin develop
$ git push origin develop
```
```
```td
||=0bf42a5=||
```
#!sh
2011-09-06 11:25 A               M─┐ [develop] Merge branch 'feature/F1' into develop
2011-09-06 11:22 A               │ o [F1] make some change 17016:2742
2011-09-06 11:22 A               │ o [F1] make some change 17016:2741
2011-09-06 11:22 A               │ o [F1] make some change 17016:2738
```
```
```td
-
```
|-----------------------------------------
```td#----------------------------------[step 5]
合作开发
```
```td
```
#!sh
# 'develop'分支的初始状态
2011-09-06 15:40 A               o [develop] =======[合作开发]=======
```
```
```td
-
```
|------------------------------------------
```td#----------------------------------[step 6]
```
#!sh
develop> git flow feature start Fs
Branches 'develop' and 'origin/develop' have diverged.
And local branch 'develop' is ahead of 'origin/develop'.
Switched to a new branch 'feature/Fs'
Summary of actions:
- A new branch 'feature/Fs' was created, based on 'develop'
- You are now on branch 'feature/Fs'
Now, start committing on your feature. When done, use:
     git flow feature finish Fs
```
```
```td
-
```
```td
-
```
|------------------------------------------
```td#----------------------------------[step 7]
```
#!sh
# 共享开发分支
feature/Fs> git flow feature publish Fs
Total 0 (delta 0), reused 0 (delta 0)
To git@localhost:oops.git
 * [new branch]      feature/Fs -> feature/Fs
Already on 'feature/Fs'
Summary of actions:
- A new remote branch 'feature/Fs' was created
- The local branch 'feature/Fs' was configured to track the remote branch
- You are now on branch 'feature/Fs'
```
```
```td
```
#!sh
$ git branch
  develop
  feature/Fs
* master
```
```
#!sh
# 表示[feature/Fs]分支从[develop]分支上拉出
2011-09-06 15:47 A               o [develop] [feature/Fs] =======[合作开发]=======
```
```
```td
```
|------------------------------------------
```td#----------------------------------[step 8]
```
#!sh
# 提交两个变动
feature/Fs> make-some-commit Fs
...
feature/Fs> make-some-commit Fs
...
```
```
```td
-
```
```td
```
#!sh
develop> git flow feature track Fs
Branch feature/Fs set up to track remote branch feature/Fs from origin.
Switched to a new branch 'feature/Fs'
Summary of actions:
- A new remote tracking branch 'feature/Fs' was created
- You are now on branch 'feature/Fs'
# 提交一个变动
feature/Fs> make-some-commit Fs
```
```
|------------------------------------------
```td#----------------------------------[step 9]
```
#!sh
$ git push origin feature/Fs
```
```
```td
```
```td
```
#!sh
# B同步A的改动, 冲突啦:P
feature/Fs> git pull origin feature/Fs                                  /tmp/gitflow/B/oops
remote: Counting objects: 7, done.
remote: Compressing objects: 100% (5/5), done.
remote: Total 6 (delta 0), reused 0 (delta 0)
Unpacking objects: 100% (6/6), done.
From localhost:oops
 * branch            feature/Fs -> FETCH_HEAD
Auto-merging Fs
CONFLICT (add/add): Merge conflict in Fs
Automatic merge failed; fix conflicts and then commit the result.
# 解决冲突
feature/Fs> git mergetools
feature/Fs> git commit -am "[Fs] B fixed conflict"
# 推到远程分支
feature/Fs> git push origin feature/Fs 
```
```
|------------------------------------------
```td#----------------------------------[step 10]
```
#!sh
# 将B的改同同步过来
feature/Fs> git pull origin feature/Fs  
feature/Fs> make-some-commit Fs
...
```
```
```td
```
#!sh
# feature/Fs分支的情况:
2011-09-06 15:58 A               o [Fs] make some change 17016:19300
2011-09-06 15:57 B               M─┐ [Fs] B fixed conflict
2011-09-06 15:55 A               │ o [Fs] make some change 17016:19166
2011-09-06 15:55 A               │ o [Fs] make some change 17016:19165
2011-09-06 15:56 B               o │ [Fs] make some change 17087:19136
2011-09-06 15:47 A               o─┘ [develop] [feature/Fs] =======[合作开发]=======
```
```
```td
a
```
|------------------------------------------
```td#----------------------------------[step 11]
```
#!sh
# 在结束前执行下列命令非常重要!!!
feature/Fs> git push origin feature/Fs
feature/Fs> git pull origin develop
# 结束特性开发
feature/Fs> git flow feature finish Fs
Switched to branch 'develop'
Merge made by recursive.
 Fs |    4 ++++
 1 files changed, 4 insertions(+), 0 deletions(-)
 create mode 100644 Fs
Deleted branch feature/Fs (was b990542).
Summary of actions:
- The feature branch 'feature/Fs' was merged into 'develop'
- Feature branch 'feature/Fs' has been removed
- You are now on branch 'develop
# 自动回到"develop"分支， "feature/Fs"分支已经被删除
develop> git  branch 
* develop
  master
# 将结果推送到远程分支
develop> git push origin develop
```
```
```td
```
#!sh
# "develop" 分支
2011-09-06 16:05 A               M─┐ [develop] Merge branch 'feature/Fs' into develop
2011-09-06 15:58 A               │ o [feature/Fs] [Fs] make some change 17016:19300
2011-09-06 15:57 A               │ M─┐ [Fs] B fixed conflict
2011-09-06 15:55 A               │ │ o [Fs] make some change 17016:19166
2011-09-06 15:55 A               │ │ o [Fs] make some change 17016:19165
2011-09-06 15:56 B               │ o │ [Fs] make some change 17087:19136
2011-09-06 15:47 A               I─┴─┘ =======[合作开发]=======
```
```
```td
```
|------------------------------------------
```td#----------------------------------[step 12]
```
#!sh
# 删除远程的Feature分支
develop> git push origin :feature/Fs
To git@localhost:oops.git
 - [deleted]         feature/Fs
```
```
```td
```
#!sh
$ git branch
  develop
* master
```
```
```td
```
#!sh
# 同步最新的"develop"分支
feature/Fs> git pull origin develop 
feature/Fs> git branch
  develop
* feature/Fs
  master
# 删除无用的Feature分支
feature/Fs> git checkout develop
develop> git branch -D feature/Fs 
```
```
|------------------------------------------
```td#----------------------------------[step x]
-
```
```td
-
```
```td
-
```
# 常见错误应对
```
#!sh
$ git flow feature start F1
Branches 'develop' and 'origin/develop' have diverged.
Branches need merging first.
$ git pull origin develop
```
```
#!sh
feature/Fs> git flow feature finish Fs
Branches 'feature/Fs' and 'origin/feature/Fs' have diverged.
And local branch 'feature/Fs' is ahead of 'origin/feature/Fs'.
Switched to branch 'develop'
Merge made by recursive.
 Fs |    3 +++
 1 files changed, 3 insertions(+), 0 deletions(-)
 create mode 100644 Fs
warning: not deleting branch 'feature/Fs' that is not yet merged to
         'refs/remotes/origin/feature/Fs', even though it is merged to HEAD.
error: The branch 'feature/Fs' is not fully merged.
If you are sure you want to delete it, run 'git branch -D feature/Fs'.
Summary of actions:
- The feature branch 'feature/Fs' was merged into 'develop'
- Feature branch 'feature/Fs' has been removed
- You are now on branch 'develop'
develop> git branch
* develop
  feature/Fs
  master
# 一般出现在 git flow feature finish <name>时，本地<name>分支与远程分支不同步，为了安全，git flow在此种情形下不会帮你删除<name>分支， 你需要确认无遗漏后手工删除
# 1. 首先确认分支确实已经无用， 删除本地分支
develop> git branch -D feature/Fs
# 2. 删除远程分支, 注意origin与:feature/Fs之间有一个空格(' ')!!!
develop> git push origin :feature/Fs
To git@localhost:oops.git
 - [deleted]         feature/Fs
```
如果使用git flow release/feature start <name> 后感觉不对, 你可以使用如下命令将其删除:
```
#!sh
$ git  branch -D  feature/F1
```
# 参考
 * http://yakiloo.com/getting-started-git-flow/
 * http://danielkummer.github.io/git-flow-cheatsheet/

