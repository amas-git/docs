---
title: Haskell
tags:
---
<!-- toc -->
# Bindings 和 Definitions
## Bindings:  <identifier> `=` <expression>
这是Bindings:
```hs
ten = 1 + 2 + 3 + 4
```
为表达式指定标识名的过程就是定义
## Local Bindings : `let` <definitions> `in` <expression>
在let中定义的绑定只在in中有效.
```hs
> let m = 2
> m
2
> let m = 1 in m - 1
0
> m
2
```
let..in..本身也是<expression>, 它的值就是 `in` 表达式的值.
## Local Bindings for Definitions : <identifier> =  <expression> where 
在定义中也可以使用`where`来产生LocalBinding. 
```div class=note
# 注意:
where-clause只能出现在定义中，它不是一个表达式, 不能单独作为表达式来使用.
​```hs
> x+1 where x = 1 (错误)
```
```
[[TOC]]
# Burstall Darlington Transformation
# Haskell Cabal 
[[TOC]]
Cabal是HaskellPlateform的一部分.
# install 
​```sh
$ sudo pacman -S cabal-install 
```
源码:
```
https://github.com/haskell/cabal/
```
## Cabal在哪里？
安装后，cabal会建立自己的目录，*Nix下: `~/.cabal`
```sh
$ tree -L 1 .cabal 
.cabal
|-- bin
|-- config
|-- lib
|-- logs
|-- packages
|-- setup-exe-cache
|-- share
`-- world
```
# 使用
首先我们来更新一下cabal的仓库
```sh
$ cabal update
Config file path source is default config file.
Config file /home/amas/.cabal/config not found.
Writing default configuration to /home/amas/.cabal/config
Downloading the latest package list from hackage.haskell.org
```
使用cabal安装开发包
```
 cabal install                             Package in the current directory
 cabal install foo                         Package from the hackage server
 cabal install foo-1.0                     Specific version of a package
 cabal install 'foo < 2'                   Constrained package version
 cabal install foo bar baz                 Several packages at once
 cabal install foo --dry-run               Show what would be installed
 cabal install foo --constraint=bar==1.0   Use version 1.0 of package bar
```
```
#!sh
$ cd source-dir
$ cabal build
$ [sudo] cabal install
```
如果是从本地编译安装, 则要使用runhaskell脚本:
```sh
$ runhaskell Setup configure
$ runhaskell Setup build
$ runhaskell Setup install
$ runhaskell Setup configure --user --prefix=$HOME
# 查看编译选项
$ runhaskell Setup configure --help
# 指定编译器
$ runhaskell Setup configure --with-compiler=ghc-6.8.2
$ runhaskell Setup build
```
# cabal sandbox
一个模块会有多个版本，但是默认的cabal不支持多版本。 这样就产生一个问题， 如果M1依赖于Mx1.0, M2依赖于Mx2.0, 而Mx1.0和Mx2.0不能同时存在，这就麻烦了。 更加严重的问题是，当你更新了某个模块会导致其他依赖有冲突的模块无法使用。 这就是Cabal Hell.
sandbox是一个临时的解决方案，通过sandbox建立局部依赖数据库，提供一个干净，独立的构建环境，其代价是冗余存储和重复编译所必须忍受的漫长时间，但总比CabalHell要好。
未来的Cabal将提供一种
----
# 参考
 * http://coldwa.st/e/blog/2013-08-20-Cabal-sandbox.html
[[TOC]]
# Categories
函数f将集合A映射为集合B
```
f: A -> B 
```
`composition`, 如果:
```
 f: A -> B
 g: B -> C
则:
 g . f -> A -> C 
```
对于集合A, 它的`identity function`
```
id : A -> A
```
# Category 的定义
```
C = (O, M, T, . , id)
```
 * C: Category C
 * O: Objects, 通常我们认为O是集合, 而且里面的元素也是集合
 * M : `morphisms`, 也是一个函数的集合
 * T: 类型信息(Type Infomation), T 属于 M x O x O 叫做C的类型信息,
  *  我们希望M中的每个元素都有一个类型, 所以对于给定的属于集合O的A和B, 对于任何一个属于M的函数f, 必然满足(f, A, B)属于T
  * 我们希望类型是唯一的, 所以如果f: A -> B 且 f: A' -> B' 可以确定 A等于A' 且 B等于B'
 * . : 
# Haskell CheetSheet
[[TOC]]
# Concurrent Haskell
Concurrent Haskell is the collective name for the facilities that Haskell provides for programming with multiple threads of control. Unlike parallel programming, where the goal is to make the program run faster by using more CPUs, the goal in concurrent programming is usually to write a program with multiple interactions. These interactions might be with the user via a user interface of some kind, with other systems, or indeed between different subsystems within the same program. Concurrency allows us to write a program in which each of these interactions is described separately but all happen at the same time. As we shall see, concurrency is a powerful tool for structuring programs with multiple interactions.
In many application areas today, some kind of concurrency is a necessity. A typical user-facing application will have an interface that must remain responsive while the application is downloading data from the network or calculating some results. Often these applications may be interacting with multiple servers over the network at the same time; a web browser, for example, will have many concurrent connections open to the sites that the user is browsing, while all the time maintaining a responsive user interface. Server-side applications also need concurrency in order to manage multiple client interactions simultaneously.
Haskell takes the view that concurrency is a useful abstraction because it allows each interaction to be programmed separately, resulting in greater modularity. Abstractions should not be too expensive because then we won’t use them—hence GHC provides lightweight threads so that concurrency can be used for a wide range of applications, without needing to worry about the overhead.
Haskell’s philosophy is to provide a set of very simple but general features that you can use to build higher-level functionality. So while the built-in functionality may seem quite sparse, in practice it is general enough to implement elaborate abstractions. Furthermore, because these abstractions are not built in, you can make your own choices about which programming model to adopt, or to program down to the low-level interfaces for performance.
Therefore, to learn Concurrent Haskell, we can start from the low-level interfaces and then explore how to combine them and build on top to create higher-level abstractions, which is exactly the approach taken in this book. The aim is that by building up the implementations of higher-level abstractions using the low-level features, the higher-level abstractions will be more accessible and less mysterious than if we had just described an API. Furthermore, by seeing examples of how to build higher-level abstractions, you should be able to go away and build your own variations or entirely new libraries.
Haskell does not take a stance on which concurrent programming model is best: actors, shared memory, and transactions are all supported, for example. (Conversely, Haskell does take a stance on parallel programming; we strongly recommend that you use one of the deterministic programming models from Part I for parallel programming.) Haskell provides all of these concurrent programming models and more—but this flexibility is a double-edged sword. The advantage is that you can choose from a wide range of tools and pick the one best suited to the task at hand, but the disadvantage is that it can be hard to decide which tool is best for the job. Hopefully by demonstrating a series of examples using each of the programming models that Haskell provides, this book will help you develop an intuition for which tool to pick for your own projects.
In the following chapters we’re going on a tour of Concurrent Haskell, starting with the basics of threads and communication in Chapter 7 through Chapter 10, moving on to some higher-level abstractions in Chapter 11, and then we’ll look at how to build multithreaded network applications in Chapter 12. Chapter 13 deals with using Concurrent Haskell to achieve parallelism, and in Chapter 14 we look at writing distributed programs that run on multiple computers. Finally, Chapter 15 will present some techniques for debugging and performance-tuning and talk about the interaction between Concurrent Haskell and foreign code.
[[TOC]]
# Thread
```hs
forkIO :: IO () -> IO ThreadId
```
!HelloThread.hs:
```hs
import Control.Concurrent
import Control.Monad
import System.IO
main = do
  hSetBuffering stdout NoBuffering         --  关闭输出缓冲
  forkIO (replicateM_ 100 (putChar 'A'))   --  打印100个A
  replicateM_ 100 (putChar 'B')            --  打印100个B
```
```sh
$ runhaskell HelloThread.hs
ABABABABAB..AB
```
 * 不同的编译器可能会有不同的输出顺序
 * GHC默认实现为完全公平调度(fairness policy)
# threadDelay :: Int -> IO ()
 * 阻塞当前进程一定时间
 * 延迟时间单位为微秒(10^6^ ms = 1s)
```
-- 阻塞1秒后,打印'hola!'
> do {threadDelay  (10^6) ; print "hola!"}
```
# MVar
```hs
data MVar a  -- abstract
newEmptyMVar :: IO (MVar a)
newMVar      :: a -> IO (MVar a)
takeMVar     :: MVar a -> IO a
putMVar      :: MVar a -> a -> IO ()
```
```
> do { m <- newEmptyMVar ; forkIO $ putMVar m 'a'; c <- takeMVar m; print c}
'a'
```
 * MVar是一个安全的容器,只有一个位置, 用来在线程之间安全的传递消息
 * takeMVar一方的线程, 拿到消息之后,可以对里面的内容(Shared Mutable State)进行处理, 再用putMVar塞回给对方
 * 有时候SMS可能是在C代码中的某种结构, 这个时候需用`()`来构造出MVar, 保护SMS.
 * 正如锁一样, MVar是Haskell构建其它并发编程组件的基础元素.
# MVar单通道
多线程日志实现1:
```hs
import Control.Concurrent
import Control.Monad
import System.IO
    
loop :: (MVar String) -> IO () 
loop m = do
    message <- takeMVar m       -- 阻塞直到有新的消息到来
    putStrLn $ "> " ++ (show message)    -- 打印消息
    loop m                
-- 初始化一个日志线程, 返回与之通讯的MVar
newLogger :: IO (MVar String) 
newLogger = do
    m <- newEmptyMVar
    forkIO $ do loop m
    return m
-- 记录一条日志, 只是向MVar中写入一条消息
logMessage :: MVar String -> String -> IO ()
logMessage m message = do
    putMVar m message
main = do 
    m <- newLogger
    forkIO $ do id <- myThreadId ; logMessage m ("message from " ++ (show id))
    forkIO $ do id <- myThreadId ; logMessage m ("message from " ++ (show id))
    forkIO $ do id <- myThreadId ; logMessage m ("message from " ++ (show id))
    forkIO $ do id <- myThreadId ; logMessage m ("message from " ++ (show id))
    id <- myThreadId ; logMessage m ("message from " ++ (show id))
```
```hs
$ runhaskell logger.hs
> "message from ThreadId 16"
> "message from ThreadId 17"
> "message from ThreadId 18"
> "message from ThreadId 19"
> "message from ThreadId 14
```
我们可以进一步抽象出Logger类型:
```hs
data Logger = Logger (MVar String)
```
几个函数相应的需要修改: 
```hs
loop :: Logger -> IO () 
loop (Logger m) = do
    message <- takeMVar m       -- 阻塞直到有新的消息到来
    putStrLn $ "> " ++ (show message)    -- 打印消息
    loop (Logger m)                
-- 初始化一个日志线程, 返回与之通讯的MVar
newLogger :: IO Logger 
newLogger = do
    m <- newEmptyMVar
    forkIO $ do loop (Logger m)
    return (Logger m)
-- 记录一条日志, 只是向MVar中写入一条消息
logMessage :: Logger -> String -> IO ()
logMessage (Logger m) message = do
    putMVar m message
```
新的版本,看上去比老版本更直观了. 但是仍然存在一个问题, 我们怎么才能终止日志线程呢?
因为我们只能向Logger线程发送String类型的数据, 可以约定特定的字符串消息来表示停止线程, 所以似乎修改一下loop函数即可:
```hs
loop :: Logger -> IO () 
loop (Logger m) = do
    message <- takeMVar m       -- 阻塞直到有新的消息到来
    if message == "exit" 
        then do putStrLn $ "Logger Exist!"; return ()
        else do putStrLn $ "> " ++ (show message)    -- 打印消息
                loop (Logger m)
-- 用来停止日志线程
stopLogger l = logMessage l "exit"
```
这样的确可以工作, 但同时也带来一个问题, 如果使用者不小心写下了"exit"日志, 就会造成非预期的效果. 所以还得想其它办法. 这个问题主要是因为使用String类型无法携带其它类型的消息所导致的, 所以咱们可以扩展一下Logger类型, 使之除了可以携带一般的消息之外, 还可以携带一些必要的控制信息.
```hs
data Logger     = Logger (MVar LogContent)
data LogContent = Message String | Stop
```
相关的几个函数需要修改:
```hs
loop :: Logger -> IO () 
loop (Logger m) = do
    c <- takeMVar m       -- 阻塞直到有新的消息到来
    case c of
        Message message -> do 
                            putStrLn $ "> " ++ (show message)
                            loop (Logger m)
        Stop            -> return () 
-- 初始化一个日志线程, 返回与之通讯的MVar
newLogger :: IO Logger 
newLogger = do
    m <- newEmptyMVar
    forkIO $ do loop (Logger m)
    return (Logger m)
-- 记录一条日志, 只是向MVar中写入一条消息
logMessage :: Logger -> String -> IO ()
logMessage (Logger m) message = do
    putMVar m (Message message)
```
增加一个用于终止Logger的函数
```hs
stopLogger :: Logger -> IO ()
stopLogger (Logger m) = putMVar m Stop
```
一切顺利, 咱们来测试一下:
```hs
main = do 
    m <- newLogger
    forkIO $ do id <- myThreadId ; logMessage m ("message from " ++ (show id))
    forkIO $ do id <- myThreadId ; logMessage m ("message from " ++ (show id))
    stopLogger m
    forkIO $ do id <- myThreadId ; logMessage m ("message from " ++ (show id))
    forkIO $ do id <- myThreadId ; logMessage m ("message from " ++ (show id))
    id <- myThreadId ; logMessage m ("message from " ++ (show id))
```
stopLogger还有一个小问题, 当我们调用stopLogger的时候, 实际上只是向Logger传递了一条消息, 但是并没有等到Logger真正结束才返回, 为了等到Logger结束, 需要做一个同步处理:
```hs
data Logger     = Logger (MVar LogContent)
data LogContent = Message String | Stop (MVar ())
loop :: Logger -> IO () 
loop (Logger m) = do
    c <- takeMVar m 
    case c of
        Message message -> do 
                            putStrLn $ "> " ++ (show message)
                            loop (Logger m)
        Stop s          -> do
                            putMVar s () -- 通知调用进程, 已经终止
                            return () 
stopLogger :: Logger -> IO ()
stopLogger (Logger m) = do
    s <- newEmptyMVar
    putMVar m (Stop s)
    takeMVar s   -- 等待Logger终止
```
最终的版本:
```hs
import Control.Concurrent
import Control.Monad
import System.IO
data Logger     = Logger (MVar LogContent)
data LogContent = Message String | Stop (MVar ())
    
loop :: Logger -> IO () 
loop (Logger m) = do
    c <- takeMVar m       -- 阻塞直到有新的消息到来
    case c of
        Message message -> do 
                            putStrLn $ "> " ++ (show message)
                            loop (Logger m)
        Stop s          -> do
                            putMVar s ()
                            return () 
stopLogger :: Logger -> IO ()
stopLogger (Logger m) = do
    s <- newEmptyMVar
    putMVar m (Stop s)
    takeMVar s
-- 初始化一个日志线程, 返回与之通讯的MVar
newLogger :: IO Logger 
newLogger = do
    m <- newEmptyMVar
    forkIO $ do loop (Logger m)
    return (Logger m)
-- 记录一条日志, 只是向MVar中写入一条消息
logMessage :: Logger -> String -> IO ()
logMessage (Logger m) message = do
    putMVar m (Message message)
main = do 
    m <- newLogger
    forkIO $ do id <- myThreadId ; logMessage m ("message from " ++ (show id))
    forkIO $ do id <- myThreadId ; logMessage m ("message from " ++ (show id))
    stopLogger m
    forkIO $ do id <- myThreadId ; logMessage m ("message from " ++ (show id))
    forkIO $ do id <- myThreadId ; logMessage m ("message from " ++ (show id))
    id <- myThreadId ; logMessage m ("message from " ++ (show id))
```
实际上仍然存在一些问题, 当日志服务终止后, 如果再试图通过MVar向其中写数据, 就会一直阻塞住调用进程.
```
$ runhaskell logger.hs
> "message from ThreadId 17"
> "message from ThreadId 18"
logger.hs: thread blocked indefinitely in an MVar operation
```
 * "logger.hs: thread blocked indefinitely in an MVar operation" , 看来有个线程一直被阻塞, 而无法退出了

# GraphReduction
# 参考
 * http://en.wikipedia.org/wiki/Graph_reduction
 * http://en.wikibooks.org/wiki/Haskell/Graph_reduction
# Effciency
[[TOC]]
# Recuction Order
# $! : Strick Evaluaction
尽管Haskell是LazyEvaluation的，但是也可以支持严格求值。严格求值本质上就是将指定的表达式转化为NormalForm。
严格求值在Haskell中有时会带来好处，比方说可以节省LazyEvauation环节中保存的非NormalForm。 
[[TOC]]
# Setp-counting Analysis
[[TOC]]
# Space Efficency
# 累计空间消耗(Accumulated space analysis)
# 最大占用空间消耗(Largest space analysis)
# Space Leak
[[TOC]]
# 函数等价
# 参考
 * http://www.haskell.org/pipermail/haskell/2004-November/014939.html
 * http://www.haskell.org/pipermail/haskell-cafe/2004-December/007766.html
 * http://stackoverflow.com/questions/9906628/equality-of-functions-in-haskell
# 错误处理
# 参考
 * http://blog.ezyang.com/2011/08/8-ways-to-report-errors-in-haskell-revisited/
[[TOC]] 
# Expressions
 * `Expressions` are things you want the computer to calculate.
 * `Values` are the results of calculating expressions.
 * `Types`
表达式有2个重要属性:
 1. 表达式有值
 2. 表达式有类型
Each expression and value has a data type.
# Types
并非只有表达式有类型，定义也是有类型的，这个也很好理解，因为定义本身便是为表达式建立符号定义的过程.
# 为表达式指定类型 :`::` <expressions>
```hs
> 1::Integer -- 整数类型的1
```
# 如何定义一种类型
```hs
data Color = Red | Blue | Green
```
 * Red, Blue, Green 叫做 Color 类型的 constructors
 * 类型名和Constructors名必须以大写字母开头
```hs
> data Shape = Circle Float | Rect Float Float
> Circle 1
...
    No instance for (Show Shape) arising from a use of `print'
    Possible fix: add an instance declaration for (Show Shape)
    In a stmt of an interactive GHCi command: print it
-- 计算机不知道该如何打印这个表达式的值
> data Shape = Circle Float | Rect Float Float deriving Show
> Circle 1
Circle 1.0
```
# foldM
```
-- this is not valid Haskell code, it is just for illustration
foldM f a1 [x1,x2,...,xn] = do a2 <- f a1 x1
                               a3 <- f a2 x2
                               ...
                               f an xn
```
# Functional Parser
[[TOC]]
# Functional Parsers
# Haskell Functions =
 * CurriedFunctions
 * InfixOperators
 * LambdaAbstractions
# Functions are Non-strict ==
```
bot = bot
bot 是非终结表达式(Non-terminating expression), 
如果你运行这个函数，结果是你得不到什么结果，但也不会结束，就像无限循环那样。
```
```
In other words, bot is a non-terminating expression. Abstractly, we denote the value of a non-terminating expression as _|_ (read "bottom"). Expressions that result in some kind of a run-time error, such as 1/0, also have this value. Such an error is not recoverable: programs will not continue past these errors. Errors encountered by the I/O system, such as an end-of-file error, are recoverable and are handled in a different manner. (Such an I/O error is really not an error at all but rather an exception. Much more will be said about exceptions in Section 7.)
```
 * f是strict 仅当f为非终结表达式
 * 绝大多数语言中函数为strict, 但Haskell中并不一定
比如:
```
bot = bot
const1 x = 1
const bot 
结果是: 1
```
因为const1总返回1，这跟参数没一点儿关系，所以参数压根就不会被求值， 就跟没bot一样
Non-strict或lazily或by need有很多好处，你不必关心求值顺序，你不用担心某个你不使用的子函数会因必须求值而产生执行期错误。 
# 无限数据结构 ==
数据构造函数也可以是non-strict的，利用这点可以定义无限大小的数据结构，
[[TOC]]
# Fun
一些有趣儿的Haskell代码片段
# 凯撒加密器
```hs
caesaEncoding :: Int -> String -> String
caesaEncoding n xs = [ shift n x | x <-xs ]
  where
    isLower c = elem c ['a'..'z']
    isUpper c = elem c ['A'..'Z']
    cch x n xs = let n' = (fromEnum x - fromEnum (head xs) + n) `mod` (length xs) in head $ rotate n' xs
    shift n c
      | isLower c = cch c n ['a'..'z']
      | isUpper c = cch c n ['A'..'Z']
      | otherwise = c
caesaDecoding :: Int -> String -> String
caesaDecoding n = caesaEncoding (-n)
```
```hs
> caesaEncoding 2 "Hello world"
"Jgnnq yqtnf"
> caesaDecoding 2 $ caesaEncoding 2 "Hello world"
"Hello world"
```
[[TOC]]
# Arithmetic
http://www.haskell.org/haskellwiki/99_questions/31_to_41
# 31. 判定一个数是否为质数
例如：
```
P31> isPrime 7
True
```
# 32. 最大公约数
求2个正整数的最大公约数，根据欧几里德算法(http://en.wikipedia.org/wiki/Euclidean_algorithm)
例如：
```
[myGCD 36 63, myGCD (-3) (-6), myGCD (-3) 6]
[9,3,3]
```
# 33. 互质
判定2个正整数是否互质，互质的定义是2数的最大公约数为1。
例如：
```
* coprime 35 64
True
```
# 34. 欧拉商数phi(m)
欧拉商数是指，对于正整数r，1<=r<=m，若r与m互质，则phi(m)返回值为这样的r的个数。
如，m = 10: r = 1,3,7,9; 则phi(m) = 4. 注意一个特例，phi(1) = 1。
例如：
```
totient 10
4
```
# 35. 质因数
求一个正整数的质因数(一个数的因数，同时为质数)。返回一个正序排列的质因数列表。
例如：
```
> primeFactors 315
[3, 3, 5, 7]
```
# 36. 质因数
求一个正整数的质因数(一个数的因数，同时为质数)。返回一个(质因数, 该质因数出现的个数)的列表。
例如：
```
> prime_factors_mult 315
[(3,2),(5,1),(7,1)]
```
# 37. 改进版欧拉商数phi(m)
看34题欧拉商数phi(m)的定义，结合36题的结果，如果((p1 m1) (p2 m2) (p3 m3) ...) 为给定m的质因数及出现次数，phi(m)可以用以下公式方便求出：
```
phi(m) = (p1 - 1) * p1 ** (m1 - 1) * 
         (p2 - 1) * p2 ** (m2 - 1) * 
         (p3 - 1) * p3 ** (m3 - 1) * ...
```
其中，**代表幂运算。
# 38. 比较2种欧拉商数phi(m)算法
使用34和37题的函数来比较2种算法。以化简的次数作为比较算法效率的依据，比如，试着计算phi(10090)。
# 39. 质数列表
通过上界与下界给定一个数的范围，返回一个列表，包含这个范围内所有质数。
例如：
```
> primesR 10 20
[11,13,17,19]
```
# 40. Goldbach's conjecture 哥德巴赫猜想
哥德巴赫猜想表述的是，任何大于2的正偶数，都是2个质数的和，如28=5+23。
这是数论里最有名的一个事实，但还没有通用的证明。
这个猜想已经在很大的数上进行了数值上的验证（远大于Prolog系统可以支持的数值）。
写一个函数来求构成一个正偶数的2个质数。
例如：
```
goldbach 28
(5, 23)
```
# 41. 哥德巴赫列表
通过上界与下界给定一个数的范围，打印出所有正偶数，及它们的哥德巴赫成分。
绝大多数情况下，一个正偶数由2个质数构成，其中一个特别小。只有很少的情况下，2个质数都大于50。
尝试算出，在2..3000的范围内，有多少这样的特例。
例如：
```
*Exercises> goldbachList 9 20
[(3,7),(5,7),(3,11),(3,13),(5,13),(3,17)]
*Exercises> goldbachList' 4 2000 50
[(73,919),(61,1321),(67,1789),(61,1867)]
```
```
--P 14
dupli :: [a] -> [a] 
dupli [] = []
dupli (x:xs) = x:x:[] ++ dupli xs
--P 15
repli' :: [a] -> Int -> [a] 
repli' [] y = []  
repli' (x:xs) y = subrepli' x y ++ repli' xs y
    where subrepli' x y = if y == 1 then [x] else [x] ++ subrepli' x (y-1)
--P 16
helper [] _ _ = []
helper (x:xs) y z = if mod z y == 0 then helper xs y (z + 1) else [x] ++ helper xs y (z + 1)
dropEvery [] _ = []
dropEvery xs y = helper xs y 1 
--P 17
getA [] _ = []
getA (x:xs) y = if y == 0 then [] else [x] ++ getA xs (y - 1)
getB [] _ = []
getB (x:xs) y = if y == 1 then xs else getB xs (y - 1)
split xs 0 = ([], xs) 
split xs y = (getA xs y, getB xs y)
--P 18
helper' [] _ _ _ = []
helper' (x:xs) i j k 
    | k < i = helper' xs i j (k + 1)
    | k > j = helper' xs i j (k + 1)
    | otherwise = [x] ++ helper' xs i j (k + 1)
slice [] _ _ = []
slice xs y z = helper' xs y z 1 
--P 19
rotate [] _ = []
rotate all@(xs) y = if y >= 0 then (getB xs y) ++ (getA xs y) else (getB xs (length(all) + y)) ++ (getA xs (length(all) + y)) 
--P 20
getOther _ [] = []
getOther y (x:xs) = if y == 1 then getOther (y - 1) xs else [x] ++ getOther (y - 1) xs
getX _ [] = []
getX y (x:xs) = if y == 1 then [x] else getX (y - 1) xs
removeAt y xs = (getX y xs, getOther y xs)
```
```hs
-- decodeModified.hs
data ItemTest x = Single x | Multiple Int x
    deriving Show
isSingle (Single _) = True
isSingle (Multiple _ _) = False
getVale (Single x) = x
getVale (Multiple _ x) = x
getCount (Multiple x _) = x
flat (Multiple 0 x) xs = xs
flat (Multiple c x) xs = flat (Multiple (c-1) x) (x:xs)
decodeModified [] = []
decodeModified xs = foldr ( b a -> if isSingle b then (getVale b):a 
                                    else (flat b [])++a)
                    []
                    xs
```
```hs
-- dropEvery.hs
dropEvery [] _ = []
dropEvery xs c = reverse (dropEvery' xs c [] [])
               where dropEvery' [] c ys zs = zs
                     dropEvery' (x:xs) c ys zs = dropEvery' xs c (x:ys) (if (mod (length (x:ys)) c) == 0 then zs 
                                                                         else x:zs)
```
```hs
-- dupli.hs
dupli [] = []
dupli xs = foldr ( b a -> (b:b:a)) [] xs
```
```hs
-- encodeDirect.hs
encodeDirect [] = []
encodeDirect xs = encodeDirect' xs [] []
                where encodeDirect' [] xs _ = xs
                      encodeDirect' (x:xs) rs ts = :
```
```hs
-- encode_modify.hs
data ItemTest x = Single x | Multiple Int x
    deriving Show
encode [] = error "error empty list"
encode xs = foldr ( b a -> if a==[] then [(1,b)] 
                            else if snd (head a) == b then (((fst (head a))+1,b):tail a) 
                            else ((1,b):a)) [] xs
encode_modify [] = []
encode_modify xs = map ( (a,b) -> if a>1 then Multiple a b 
                                   else Single b) 
                   (encode xs)
```
```hs
-- removeAt.hs
removeAt [] _ = []
removeAt xs c = removeAt' xs [] c (length xs)
              where removeAt' [] ys c t = reverse ys
                    removeAt' (x:xs) ys c t = if (t-length xs) == c then removeAt' xs ys c t
                                            else removeAt' xs (x:ys) c t
```
```hs
-- repli.hs
flat 0 b xs = xs
flat c b xs = flat (c-1) b (b:xs)
repli [] _ = []
repli xs c = foldr ( b a -> (flat c b [])++a) [] xs
```
```hs
-- rotate.hs
rotate' (x:xs) ys c = if length ys == (c-1) then xs++(reverse (x:ys))
                      else rotate' xs (x:ys) c
rotate [] _ = []
rotate xs 0 = xs
rotate xs c | c>=0 = rotate' xs [] c
            | otherwise = rotate' xs [] (length xs +c) 
```
```hs
-- slice.hs
slice [] _ _ = []
slice xs f e = if (length xs < e || f <= 0) then error "out of bounds"
               else slice' xs 1 [] f e
               where slice' (x:xs) c rs f e = if c == e then reverse (x:rs)
                                              else if c >= f then slice' xs (c+1) (x:rs) f e
                                              else slice' xs (c+1) rs f e
```
```hs
-- split.hs
split' [] (ls,rs) c = (reverse ls,reverse rs)
split' (x:xs) (ls,rs) c = if (c>=1) then split' xs ((x:ls),rs) (c-1)
                          else split' xs (ls,(x:rs)) c
split [] _ = error "empty list"
split xs c = if (c > (length xs) || c <= 0)  then (error "out of bounds") 
             else (split' xs ([],[]) c)
```
```hs
duplicate [] = []
duplicate (x:xs) = x:(x:duplicate xs)
-- duplicate = foldr ( xs -> x:x:xs) []
-- 1. 在haskell中已经有replicate函数, 但跟这个意思不太一样
-- 2. 注意标准函数的实现方式，一般都将list放置在最后一个参数
replicate' [] _ = []
replicate' (x:xs) n = (take n (repeat x)) ++ (replicate'' xs n) 
replicate'' [] _ = []
replicate'' xs n = foldr ( xs -> (take n (repeat x)) ++ xs) [] xs
dropEvery xs 0 = xs
dropEvery [] _ = []
dropEvery xs n = (take (n-1) xs) ++ (dropEvery (drop n xs) n)
-- Data.List中有个 splitAt
splite xs n = [ (take n xs) , (drop n xs) ]
splite' n xs = acc n 0 [] [] xs
  where
    acc _ _ xs1 _ [] = [reverse xs1,[]]
    acc n i xs1 xs2 (x:xs)
      | i < n = acc n (i+1) (x:xs1) xs2 xs
      | otherwise = [reverse xs1, x:xs]
slice l r xs = acc l r l [] (drop l xs)
  where
    acc _ _ _ rx [] = reverse rx
    acc l r i rx (x:xs)
      | i < r = acc l r (i+1) (x:rx) xs
      | otherwise = reverse rx
-- a mod b + (-a) mod b = b
-- [1,2,3,4,5,6]
-- [3,4,5,6,1,2] 左移2 OR 右移(6-2)
rotate 0 xs = xs
rotate n xs = (drop n' xs) ++ (take n' xs)
  where n' = n `mod` (length xs)
-- 根据split修改而成
removeAt' n xs = acc n 0 [] xs
  where
    acc _ _ xs1 [] = reverse xs1
    acc n i xs1 (x:xs)
      | i < n = acc n (i+1) (x:xs1) xs
      | otherwise = reverse xs1 ++  xs
removeAt n xs = acc [] 1 n xs
  where
    acc rx _ _ [] = reverse rx
    acc rx i n (x:xs)
        | i == n = acc rx (i+1) n xs 
        | otherwise = acc (x:rx) (i+1) n xs 
-- 语法糖
range l r = [l..r]
range' l r
  | l > r  = []
  | l == r = [l]
  | otherwise = l:(range' (succ l) r)
-- splite' == splitAt (Data.List模块中)
insertAt x n xs = let xss = splite' (n-1) xs
                      xs1 = head xss
                      xs2 = last xss
                  in xs1 ++ (x:xs2)
```
```hs
-- insertAt.hs
insertAt [] _ _ = []
insertAt xs a 1 = (a:xs)
insertAt xs i c = foldr ( b a-> if (length xs - length a) == (c-1) then (b:i:a)
                                 else (b:a))
                  []
                  xs
```
```hs
-- range.hs
range f e = reverse (range' f e [])
          where range' f e xs = if f < e then range' (f+1) e (f:xs)
                                else (f:xs)
                                 
```
[[TOC]]
# 1.  last
Find the last element of a list.
例如:
```
> last' [1..100]
100
> last' []
error
```
A1: 递归
```hs
last' [] = error "Empty list"
last' (x:[]) = x
last' (x:xs) =  last' xs
```
A2:
```hs
last' = head.reverse
```
A3:
```hs
last' = 
```
# 2.  butlast
Find the last but one element of a list.
例如:
```
> butlast [1,2,3,4]
3
> butlast ['a'..'z']
y
> butlast []
error
```
```hs
butlast [] = error "empty list"
butlast [_] = error "less than 2 element" 
butlast (x:(_:[])) = x
butlast (x:xs) = butlast xs
butlast' = last.init
```
# 3.  nth
Find the K'th element of a list. The first element in the list is number 1.
```
> nth 1 [1..200] 
1
> nth 4 [1,2,3]
error
```
```hs
-- 利用累加的方式，如果希望下标从0开始，计数从0开始即可
nth n xs = _nth xs n 1 
  where _nth [] _ _   = error "not found"
        _nth (x:xs) n m
          | (n == m)  = x
          | otherwise = _nth xs n (m+1)
```
# 4.  length
Find the number of elements of a list.
A1: recursive
```hs
length' :: [a] -> Int
length' [] = 0
length' (x:xs) = 1 + length' xs
```
A2: curring with foldr
```hs
length' = foldr ( n -> n+1) 0
```
A3: curring with foldl
```hs
length' = foldl (
 x -> n+1) 0
```
A4: curring with sum/map
```hs
length' = sum.map ( > 1)
```
# 5.  reverse
Reverse a list.
```hs
reverse' :: [a] -> [a]
reverse' [] = []
reverse' (x:xs) = (reverse' xs)++[x]
```
# 6.  palindrome
Find out whether a list is a palindrome. A palindrome can be read forward or backward; e.g. (x a m a x).
```
> palindrome "1983891"
True
> palindrome "amas"
False
```
A1:
```hs
palindrome xs = xs == (reverse xs)
```
A2:
```hs
palindrome [] = True
palindrome [_] = True
palindrome [x,y] = x == y
palindrome (x:xs) = (x == (last xs)) && (palindrome (init xs))
-- 最后一个函数可以改写为一个效率更高的版本
palindrome (x:xs)
  | (x == (last xs)) = palindrome (init xs)
  | otherwise = False
```
# 7. flatten
Flatten a nested list structure.
Transform a list, possibly holding lists as elements into a `flat' list by replacing each list with its elements (recursively).
```hs
data NestedList a = Elem a | List [NestedList a]
> flatten (Elem 5)
[5]
> flatten (List [Elem 1, List [Elem 2, List [Elem 3, Elem 4], Elem 5]])
[1,2,3,4,5]
> flatten (List [])
[]
```
A1:
```hs
flatten (Elem x) = [x]
flatten (List []) = []
flatten (List (x:xs)) = (flatten x) ++ (flatten (List xs))
```
# 8.  compress
合并列表中连续重复的元素.
```hs
> compress []
[]
> compress [1,1,9,8,8]
[1,9,8]
```
A1: 递归
```hs
compress []  = []
compress [x] = [x]
compress (x:xs)
  | x == head xs = compress xs
  | otherwise = x:(compress xs)
```
A2: foldr + lambda
注意foldr初值`[last xs]`的使用, foldr是从最后列表中最后一个元素开始, 从后往前计算的. 因此第一个参与比较的元素并不是xs的第一个元素.
```hs
compress' [] = []
compress' [x] = [x]
compress' xs = foldr ( xs ->
                           if x == (head xs) then xs
                           else x:xs)
                   [last xs] xs
```
理解错了,写出了uniq
```hs
uniq :: Ord a => [a] -> [a]
uniq xs = _uniq xs []
  where _uniq [] ys   = reverse ys
        _uniq (x:xs) ys
          | x `elem` ys = _uniq xs ys
          | otherwise = _uniq xs (x:ys)
```
# 9.  pack
Pack consecutive duplicates of list elements into sublists. If a list contains repeated elements they should be placed in separate sublists.
```hs
> pack ['a', 'a', 'a', 'a', 'b', 'c', 'c', 'a', 'a', 'd', 'e', 'e', 'e', 'e']
["aaaa","b","cc","aa","d","eeee"]
```
```hs
pack [] = []
pack xs = [ys] ++ (pack (drop (length ys) xs))
  where ys = takeEq xs
        takeEq [x] = [x]
        takeEq (x:xs)
          | x == (head xs) = x:(takeEq xs)
          | otherwise = [x]
```
# 10.  run-length encode
Run-length encoding of a list. Use the result of problem P09 to implement the so-called run-length encoding data compression method. Consecutive duplicates of elements are encoded as lists (N E) where N is the number of duplicates of the element E.
 * 对于任意元素E, 其连续重复出现次数为N, 编码为: `(E, N)`
```hs
> encode "aaaabccaadeeee"
[(4,'a'),(1,'b'),(2,'c'),(2,'a'),(1,'d'),(4,'e')]
```
```hs
-- map + lambda
encode xss = map ( s->(length xs, head xs)) (pack xss)
-- 使用列表生成式
encode xss = [(length xs, head xs) | xs<-(pack xss)]
```
# 11. run-length encode'
基于问题10,稍微修改一下编码策略, 对于列表中任意元素E, 其连续重复出现次数为N:
 *  若N>1则: 编码为: `(E, N)`
 *  否则: 编码为 `E`
例如:
```hs
> encode' "aaaabccaadeeee"
[Multiple 4 'a',
 Single     'b',
 Multiple 2 'c',
 Multiple 2 'a',
 Single     'd',
 Multiple 4 'e']
```
!EncodeElement:
```hs
data EncodeElement a = Multiple Int a | Single a deriving (Show)
```
A1:
```hs
encodeModified :: [[a]] -> [EncodeElement a]
encodeModified [] = []
encodeModified (x:xs) = (encElem x) : encodeModified xs   
	where encElem x = case length(x) of
		1 -> Single(head x)
		_ -> Multiple (length x) (head x)
```
# 12. decode'
解码由11生成的列表.
例如:
```hs
> decode' 
[Multiple 4 'a',
 Single     'b',
 Multiple 2 'c',
 Multiple 2 'a',
 Single     'd',
 Multiple 4 'e']
"aaaabccaadeeee"
```
A1:
```hs
decodeModified :: [EncodeElement a] -> [a]
decodeModified [] = []
decodeModified (x:xs) = (decElem x) ++ decodeModified xs
	where 	decElem :: EncodeElement a -> [a]
		decElem (Single a)	= [a]
		decElem (Multiple n a) 	= take n (repeat a) 
```
# 13. run-length encode!''
基于10得到的结果, 只要将形式为`(1 E)`的元素转化为E, 即可实现run-length编码方法.
A1:
```hs
encodeDirect :: (Eq a) => [a] -> [EncodeElement a]
encodeDirect [] = []
encodeDirect (x:xs) = encodeHelper 1 x xs
encodeHelper n y [] = [encElem n y]
encodeHelper n y (x:xs) | y == x	= encodeHelper (n+1) y xs
			| otherwise	= encElem n y : (encodeHelper 1 x xs)
encElem 1 y = Single y
encElem n y = Multiple n y
```
# 14. duplicate
复制列表中的每个元素.
例如:
```hs
> duplicate [1,2,3]
[1,1,2,2,3,3]
```
A1: 递归方式
```
duplicate [] = []
duplicate (x:xs) = x:(x:duplicate xs)
```
# 15. duplicateN
例如:
```hs
> duplicateN 1 [1,2,3]
[1,2,3]
> duplicateN 2 [1,2,3]
[1,1,2,2,3,3]
> duplicateN 3[1,2,3]
[1,1,1,2,2,2,3,3,3]
```
A1: 递归方式
```hs
duplicateN _ []  = []
duplicateN n (x:xs) = (take n (repeat x)) ++ (duplicateN n xs) 
```
A2: fold方式
```hs
duplicateN _ [] = []
duplicateN n xs = foldr ( xs -> (take n (repeat x)) ++ xs) [] xs
```
# 16. drop every
给定列表L和自然数N, 丢弃列表中第N, 2*N, ..., m*N (m*N<length L && (m+1)*N>length L)个元素.
例如:
```hs
> dropEvery 2 [1,2,3,4,5,6]
[1,2,5]
```
A1: take + drop
```hs
dropEvery 0 xs = xs
dropEvery _ [] = []
dropEvery n xs = (take (n-1) xs) ++ (dropEvery n (drop n xs))
```
# 17. split N
将指定列表一分为二, 前N个元素构成列表一, 其余为列表二.
例如:
```hs
> split 2 [1,2,3,4,5]
[[1,2],[3,4,5]]
```
A1: take + drop
```hs
split n xs = [ (take n xs) , (drop n xs) ]
```
A2: 累加方式
```hs
split xs n = acc 0 [] xs
  where
    acc _ rs [] = [reverse rs,[]]
    acc i rs (x:xs)
      | i < n = acc (i+1) (x:rs) xs
      | otherwise = [reverse rs, x:xs]
```
# 18. slice
给定列表L, 自然数M和N(M<N<=length L), 将第M个元素到第N个元素取出,形成一个新的列表.
例如:
```hs
> slice 1 4 [1,2,3,4,5] 
[1,2,3,4]
> slice 4 5 [1,2,3,4,5] 
[4,5]
> slice 3 4 "hello world" 
"ll"
```
A1: 累加方式
```hs
slice l r xs = acc l r l [] (drop l xs)
  where
    acc _ _ _ rx [] = reverse rx
    acc l r i rx (x:xs)
      | i < r = acc l r (i+1) (x:rx) xs
      | otherwise = reverse rx
```
# 19. rotate
 * N > 0 : 循环左移列表中的N个元素
 * N < 0 : 循环右移列表中的N个元素
```hs
> rotate 3 ['a','b','c','d','e','f','g','h']
"defghabc"
> rotate  (-2) ['a','b','c','d','e','f','g','h']
"ghabcdef"
```
```hs
rotate 0 xs = xs
rotate n xs = (drop n' xs) ++ (take n' xs)
  where n' = n `mod` (length xs)
```
```div class=note
# 注意:
对于List ![1,2,3,4,5,6]通过位移变为
![3,4,5,6,1,2], 既可以认为是循环左移动2位, 也可以认为是循环右移4位.
```
# 20. remove at
```hs
> removeAt 1 "hello"
"ello"
```
A1: 累加方式
```hs
removeAt n xs = acc [] 1 xs
  where
    acc rx _ [] = reverse rx
    acc rx i (x:xs)
        | i == n = acc rx (i+1) xs 
        | otherwise = acc (x:rx) (i+1) xs 
```
# 21 insert at
在第列表中第N个位置插入元素。
```hs
> insertAt 'D' 4 "abcef" 
"abcDef"
```
A1:  配合split
```hs
insertAt x n xs = let xss = split (n-1) xs
                      xs1 = head xss
                      xs2 = last xss
                  in xs1 ++ (x:xs2)
  where split n xs = [ (take n xs) , (drop n xs) ]
```
# 22. range
给定整数M,N(M<N), 返回由M到N之间所有元素所构成的列表。
```hs
> range 1 3 [1..100]
[1,2,3]
```
A1. haskell中没有提供range函数, 取而代之的是下面这个语法糖:
```hs
range l r = [l..r]
```
A2. 递归解法
```hs
range' l r
  | l > r  = []
  | l == r = [l]
  | otherwise = l:(range' (succ l) r)
```
# 23. random select
从列表中随机选取N个元素，构成一个新的列表。
```hs
> randomSelect [1..100] 3
[1,34,6]
```
# 24.  draw
生成N个大于0且小于M的随机数，形成一个列表。
```hs
> draw 5 10
[1,5,6,7,8]
```
# 25. random permutation
随机打乱一个给定的列表，产生一个新的排列。
```hs
> randomPermutation "hello"
"leohl"
```
# 26. combinations
列表L中没有重复元素， 从中任意挑选N个元素构成一个新的子列表， 求所有长度为N的子列表构成的列表。
```hs
> combinations 3 "abcdef"
["abc","abd","abe",...]
```
# 27.  disjoint
```div class=note
# 互斥集 / DISJOINT SETS / 并查集
 * http://www.csie.ntnu.edu.tw/~u91029/DisjointSets.html
```
将结合中的元素拆分成互斥集的集合.
9人分三组, 一组2人,一组3人,一组4人, 请枚举全部的分组方式.
```hs
> disjoint [1..9]
[
 [[1,2], [3,4,5], [6,7,8,9]],
 ...
]
```
# 28.
# 29.  lsort
```hs
> lsort ["abc","de","fgh","de","ijkl","mn","o"]
["o","de","de","mn","abc","fgh","ijkl"]
```
# 30. lfsort
```hs
> lfsort ["abc", "de", "fgh", "de", "ijkl", "mn", "o"]
["ijkl","o","abc","fgh","de","de","mn"]
```
# 总结
 1. 列表相关处理经常会使用以下三种模式
  * 递归方式
  * 累加器方式
  * fold方式
 2. 处理列表的函数, 在参数的顺序上应当进可能将列表参数置于最后(因为列表是函数处理的主要对象), 既可方便函数之间的相互组合, 也方便基于此函数添加HigherOrder函数.
 3. 在列表操作中谨慎使用`++`函数, 尽可能考虑使用`x:xs`这种O(1)的头插入方式, 在返回最终结果时reverse一下即可.
```
-- Q11 add a data type to improve run-length encode
data EncodeElement a = Multiple Int a | Single a deriving (Show)
{-
encodeModified :: [[a]] -> [EncodeElement a]
encodeModified [] = []
encodeModified (x:xs) = (encElem x) : encodeModified xs   
	where encElem x = case length(x) of
		1 -> Single(head x)
		_ -> Multiple (length x) (head x)
-}
encodeModified :: [(Int, a)] -> [EncodeElement a]
encodeModified [] = []
encodeModified (x:xs) = (encElem x) : encodeModified xs
	where encElem x = case (fst x) of
		1 -> Single(snd x) 
		n -> Multiple n (snd x)
		
-- Q12 decode run-length encoded list
decodeModified :: [EncodeElement a] -> [a]
decodeModified [] = []
decodeModified (x:xs) = (decElem x) ++ decodeModified xs
	where 	decElem :: EncodeElement a -> [a]
		decElem (Single a)	= [a]
		decElem (Multiple n a) 	= take n (repeat a) 
-- Q13 a direct solution of RLE
encodeDirect :: (Eq a) => [a] -> [EncodeElement a]
encodeDirect [] = []
encodeDirect (x:xs) = encodeHelper 1 x xs
encodeHelper n y [] = [encElem n y]
encodeHelper n y (x:xs) | y == x	= encodeHelper (n+1) y xs
			| otherwise	= encElem n y : (encodeHelper 1 x xs)
encElem 1 y = Single y
encElem n y = Multiple n y
		
-- Q14 duplicate the elements of a list
dupli :: [a] -> [a]
dupli [] = []
dupli (x:xs) = [x,x] ++ (dupli xs)
-- Q15 replicate the elements of a list a given number of times
repli :: [a] -> Int -> [a]
repli [] n = []
repli (x:xs) n = (repeat_helper x n) ++ (repli xs n)
	where repeat_helper x n = take n (repeat x)
-- Q16 drop every N'th element from a list
dropEvery :: [a] -> Int -> [a]
dropEvery xs n = dropHelper xs n 1 
	where dropHelper xs n m =
		case xs of
			[] -> []
			(x:xs) -> if not (m `mod` n == 0)	
				then x:dropHelper xs n (m+1)
				else dropHelper xs n (m+1)
-- Q17 split a list into two
split :: [a] -> Int -> [[a]]
split xs n = [(take n xs), (drop n xs)]
{-
split :: [a] -> Int -> [[a]]
split xs n 
	= splitHelper xs [] 0
		where
			splitHelper :: [a] -> [a] -> Int -> [[a]]
			splitHelper xs acc m =
				if (m==n) 
				then [ acc, xs] 
				else splitHelper (tail xs) (acc+[(head xs)]) (m+1)
-}
-- Q18 slice from a list
slice :: [a] -> Int -> Int -> [a]
slice xs i k = head (tail (split (head (split xs k)) (i-1)))
-- Q19 rotate a list N places to the left
rotate :: [a] -> Int -> [a]
rotate xs n = 
	(drop m xs) ++ (take m xs)
	where m = n `mod` length(xs)
-- Q20 REMOVE K'th element from a list
removeAt :: Int -> [a] -> (a, [a])
removeAt n xs = 
	((last first), (init first) ++ second)
	where 
		twoParts = split xs n
		first = head twoParts
		second = (head (tail twoParts))
```
[[TOC]]
# LogicAndCodes
http://www.haskell.org/haskellwiki/99_questions/46_to_50
# 问题46
定义方法 and/2, or/2, nand/2, nor/2, xor/2, impl/2 和 equ/2（逻辑等于）根据自身判断结果决定是真是假。比如and(A,B)只有在A和B都为真的情况下才为真。
一个有两个参数的逻辑表达式可以这样表示：and(or(A,B),nand(A,B))
现在写一个有两个参数的table/3，并且根据自身逻辑写出其真值表。
例如：
```
(table A B (and A (or A B)))
true true true
true fail true
fail true fail
fail fail fail
```
在haskell中：
```
> table ( b -> (and' a (or' a b)))
True True True
True False True
False True False
False False False
```
# 问题47
逻辑表达式的真值表
根据问题46里的逻辑方法，现在写一个像java一样更自然的写法。
例如：
```
* (table A B (A and (A or not B)))
true true true
true fail true
fail true fail
fail fail fail
```
在haskell中：
```
> table2 ( b -> a `and'` (a `or'` not b))
True True True
True False True
False True False
False False False
```
# 问题48
逻辑表达式的真值表
一般化问题P47使得所述逻辑表达式可以包含任意数量的逻辑变量。像定义方法table/2的方式，打出包含逻辑变量(List,Expr)的真值表。
例如：
```
* (table (A,B,C) (A and (B or C) equ A and B or A and C))
true true true true
true true fail true
true fail true true
true fail fail true
fail true true true
fail true fail true
fail fail true true
fail fail fail true
```
在haskell中：
```
> tablen 3 ([a,b,c] -> a `and'` (b `or'` c) `equ'` a `and'` b `or'` a `and'` c)
-- infixl 3 `equ'`
True  True  True  True
True  True  False True
True  False True  True
True  False False True
False True  True  True
False True  False True
False False True  True
False False False True
 
-- infixl 7 `equ'`
True  True  True  True
True  True  False True
True  False True  True
True  False False False
False True  True  False
False True  False False
False False True  False
False False False False
```
# 问题49
格雷码
n位的格雷码是按照一定的规则构成的n位串序列。
例如：
```
n = 1: C(1) = ['0','1'].
n = 2: C(2) = ['00','01','11','10'].
n = 3: C(3) = ['000','001','011','010',´110´,´111´,´101´,´100´].
```
找出构造规则，按照以下规格写个方法：
```
% gray(N,C) :- C is the N-bit Gray code
```
当重复使用的时候，你可以用“结果缓存”的方式让它更有效率吗？
在haskell中：
```
P49> gray 3
["000","001","011","010","110","111","101","100"]
```
# 问题50
哈弗曼编码
我们设一个符号和它的使用频率的相对应的列表，例如：
```
[fr(a,45),fr(b,13),fr(c,12),fr(d,16),fr(e,9),fr(f,5)]
```
我们的目的是构造一个列表hc(S,C),C代表符号S.根据我们的例子，结果应该是：
```
 Hs = [hc(a,'0'), hc(b,'101'), hc(c,'100'), hc(d,'111'), hc(e,'1101'), hc(f,'1100')] [hc(a,'01'),...etc.].
```
该任务应该以如下方式进行：
```
% huffman(Fs,Hs) :- Hs is the Huffman code table for the frequency table Fs
```
在haskell中：
```
*Exercises> huffman [('a',45),('b',13),('c',12),('d',16),('e',9),('f',5)]
[('a',"0"),('b',"101"),('c',"100"),('d',"111"),('e',"1101"),('f',"1100")]
```
# H99
 * http://www.haskell.org/haskellwiki/H-99:_Ninety-Nine_Haskell_Problems
[[TOC]]
# Haskell Breadcrumbs
# 原文
 * http://acm.wustl.edu/functional/hs-breads.php
[[TOC]]
# Haskell Platform
由以下几个部分组成:
 * ghc (ghc) — The compiler
 * cabal-install (cabal-install) — A command line interface for Cabal and Hackage
 * haddock (haddock) — Tools for generating documentation
 * happy (happy) — Parser generator
 * alex (alex) — Lexical analyzer generator
all
# any
[[TOC]]
# (.)
假如有函数f1 :: a -> b 又有函数f2 :: b -> c , 即f1的输出恰好是f2的输入, 于是我们可以将两个函数串联起来得到`(f2 (f1 ...))`,  在Haskell中可以用`(.)`函数优雅的干这件事情.
更为一般的, 对于a -> a 的任意函数, 可以以此方式任意组合形成新的函数, 这个函数的类型仍然是a -> a
```
(b -> c) -> (a -> b) -> (a -> c)
```
 * 优先级: 9
# 简单的示例
```hs
> let double x = x + x
> double 2
4
> (double . double . double) 2
16
```
# compose
```hs
compose :: [a -> a] -> (a -> a)
compose = foldr (.) id
```
```hs
> compose [double, double, double] 2
16
```
# dropWhile
# filter
[[TOC]]
交换指定函数的参数
# flip :: (a -> b -> c) -> (b -> a -> c)
```hs
flip' f a b = f b a
```
```hs
> (flip'(-)) 1 2
1
> (-) 1 2
-1
```
[[TOC]]
# foldl
# sum
```hs
sum = foldl (+) 0
```
# reverse
```hs
reverse = foldl ( s x -> x:xs) []
```
# product
```hs
product = foldl (*) 1
```
# or
```hs
or :: [Bool] -> Bool
or = foldl (||) True
```
# and
```hs
and :: [Bool] -> Bool
and = foldl (&&) True
```
# length
```hs
length :: [a] -> Int
length = foldl (
 _ -> n + 1) 0
```
# map
```hs
map f = reverse . foldl ( s x -> (f x):xs ) []
-- 或者
map f = (foldl ( s x -> x:xs) []) . foldl ( s x -> (f x):xs ) []
```
[[TOC]]
# foldr
# reverse
```hs
reverse = foldr ( xs -> xs ++ [x]) []
```
[[TOC]]
# Higher Order Function
一个函数, 至少满足以下任意一个条件, 便是高阶函数.
 * 输入参数至少有一个函数
 * 返回值是一个函数
# map :: (a -> b) -> [a] -> [b]
```hs
> map (1+) [1..10]
[2,3,4,5,6,7,8,9,10,11]
```
# filter :: (a -> Bool) -> [a] -> [a]
# all :: (a -> Bool) -> [a] -> Bool
```
> all ( > x `mod` 2 == 0) [2,4,6]
True
```
# any :: (a -> Bool) -> [a] -> Bool
```hs
> any null ["","hello","world"]
True
```
# takeWhile :: (a -> Bool) -> [a] -> [a]
```hs

```
# dropWhile :: (a -> Bool) -> [a] -> [a]
# (.) :: (b -> c) -> (a -> b) -> a -> c 
# ($!) :: (a -> b) -> a -> b
# ($) :: (a -> b) -> a -> b
# 参考
 * http://en.wikibooks.org/wiki/Haskell/Higher-order_functions_and_Currying
# takeWhile
# zipWith
```
zipWith :: (a -> b -> c) -> [a] -> [b] -> [c]
```
# zip和zipWith
```hs
zip :: [a] -> [b] -> [(a,b)]
zip = zipWith ( b ->(a,b))
```
# 实现
```
zip' :: [a] -> [b] -> [(a,b)]
zip' _ [] = []
zip' [] _ = []
zip' (x:xs) (y:ys) = [(x,y)] ++ (zip' xs ys)
```
在此基础上抽象一点点, 得到zipWith'
```hs
zipWith' :: (a -> b -> c) -> [a] -> [b] -> [c]
zipWith' f _ [] = []
zipWith' f [] _ = []
zipWith' f (x:xs) (y:ys) = [f x y] ++ (zipWith' f xs ys)
```
# 历史
# 1930s
In the 1930s, Alonzo Church developed the lambda calculus, a simple but
powerful mathematical theory of functions.
# 1950s
In the 1950s, John McCarthy developed Lisp (“LISt Processor”), generally
regarded as being the first functional programming language. Lisp had some
influences from the lambda calculus, but still adopted variable assignments
as a central feature of the language
# 1960s
In the 1960s, Peter Landin developed ISWIM (“If you See What I Mean”),
the first pure functional programming language, based strongly on the
lambda calculus and having no variable assignments.
# 1970s
In the 1970s, John Backus developed FP (“Functional Programming”), a
functional programming language that particularly emphasised the idea of
higher-order functions and reasoning about programs.
r Also in the 1970s, Robin Milner and others developed ML (“Meta-
Language”), the first of the modern functional programming languages,
which introduced the idea of polymorphic types and type inference.
r In the 1970s and 1980s, David Turner developed a number of lazy func-
tional programming languages, culminating in the commercially produced
language Miranda (meaning “admirable”).
# 1980s
In 1987, an international committee of researchers initiated the develop-
ment of Haskell (named after the logician Haskell Curry), a standard lazy
functional programming language.
# 2000s
In 2003, the committee published the Haskell Report, which defines a long-
awaited stable version of Haskell, and is the culmination of fifteen years of
work on the language by its designers.
----
It is worthy of note that three of the above researchers — McCarthy, Backus,
and Milner — have each received the ACM Turing Award, which is generally
regarded as being the computing equivalent of a Nobel prize.
[[TOC]]
# How to Learn Haskell
即便你已经掌握一些编程语言，学习Haskell也非一朝一夕之事，但愿你有始有终。
# 本文的目的
这是一篇关于如何学习Haskell的介绍。而非介绍Haskell本身，可谓元指南(meta-tutorial). 我们从宏观角度解释如何更好的学习Haskell。 你应该结合本文提出的其他材料完成整个学习过程。我们选择有用的资源，精心编排，让你可以循序渐进。但愿对你有所帮助，祝君好运。
# 出发
为什么学习Haskell?
 * Unlike some modern languages, Haskell is not a language you will pick up in two days and then be able to write your homework 10x faster in. 
 * Haskell will make you sweat to write simple programs, but it can also make writing things you thought were really complex quite a bit simpler. 
 * 多数情况下，这是一个很好的学习经历，因为这些开发模式和思考方式将有益于你从事的计算机科学技术相关工作的方方面面
 * 你可以Google一下`Why Haskell`, 可以粗略看看Hskell世界的轮廓，试着读读: [http://acm.wustl.edu/functional/whyfp.php Why Functional Programming]
# 安装并运行Haskell. 
安装Haskell编译/解释器，你有两个选择
 * HUGS
 * GHC
推荐你安装GHC, 她更强大和灵活。看起来已经成为实事上的标准。 
你可以在这儿:[http://tryhaskell.org/ 戳我] 线上体验一把Haskell。它是一个在线解释器，别指望它可以使用所有的库和函数。
资料:
 * [http://www.haskell.org/ The Haskell Wiki]: 各种手册，论文
 * [http://www.haskell.org/haskellwiki/Haskell_in_5_steps Haskell in 5 Steps]
 * [http://book.realworldhaskell.org/read/installing-ghc-and-haskell-libraries.html 如何安装GHC和Haskell库]

# 编辑器与GHCI
# EMACS + haskell-mode
  * C-c C-l: 对整个文件内容求值 
# VIM + GHCI

# GHCi
如果你还打算使用GHCi,
 * [http://www.haskell.org/ghc/docs/latest/html/users_guide/interactive-evaluation.html 如何使用GHCi]
用Haskell的话来讲，使用GHCi让你在I/O monad中书写Haskell代码。
If you want to use GHCI, you should learn how to use it, check out The GHCI page on the Haskell wiki, and realize that you're always writing code in the I/O monad when using GHCI.
GHCI is kind of weird. I said it. But, it's weird because it's awesome. As you start studying 
Haskell you'll soon find out all about the distinction between writing code "in a monad" and writing code outside of one. To give you a quick example of what I mean, compare:
GHCi有点诡异。虽说如此，但这正是它牛x之处。对于初学Haskell的你，很快就会发现什么是
 * writing code "in a monad" 
 * writing code outside of one
为此，我来举个栗子:
```hs
a = 5
```
```hs
let a = 5
```
```div class=note
在你理解I/O monad之前，还是建议你在文件中编写代码，用GHCi加载。当你足够了解I/O monad之后，你就不觉得GHCi诡异了。
```
# 基础篇
首先，忘记你之前所使用的编程语言。为此请对号入座。
 * CC++/Java: 
  * 算法实现的模式是完全不同的，你没有循环可以使用
  * 忘记类吧
  * return可不是你想的那样
  * 代码不是按照你输入的顺序执行的
  * 啥？哇，太TMD牛X了
 * Python/Perl:
  *  You might be in the same boat as the previous group, but you may have run into some useful functions like reduce, map, etc. that use functions as first-class objects, and use this feature to abstract away some common looping patterns. 
  * This shows up all the time in Haskell. If you're a Perl whiz, who knows what you've run into, but depending on whether you write Perl like a C coder, or write Perl like a Lisp coder you may find yourself accordingly prepared.
 * Ruby:
  * Ruby's mixins and duck-typing actually bear some similarities to Haskell's typeclass system, so lookout for parallels, but don't assume you know everything since Ruby is object oriented and Haskell is not. Another handy thing is that Ruby coders tend to use Ruby's inject, map, collect, etc. functions quite a bit, which all represent very common patterns in Haskell.
 * Sheme/Lisp:
  * 看来只有类型系统才能降伏你了
 * ML/Ocaml:
  * You've still got a few things to learn, but these languages aren't that different. The main distinction (which is not small) is Haskell's purity, which will also be freaking out everyone in the above groups.
所以目前你应该已经找到自己的位子了，现在我可以说说Haskell和其它编程语言的不同之处:
 * 语法不同
 * 求值顺序是non-intuitive的。Haskell是惰性求值(call-by-need)，有个不错的参考: [http://mitpress.mit.edu/sicp/
    SCIP 3.5章](如果你愿意学点儿Scheme的话)
 * 在其它绝大多数语言中一个变量的类型只不过是定义了它是如何在机器中表示的。在Haskell中类型更加一般，可以包含FunctionSignature,Emumeration,Tuple,List或者其它类型.
 * 类型表示不太相同。
 * Haskell中的Typeclasses有点儿像Java中的接口，或者CommonLisp的GenericMethodSpecifications.有很多不错的参考我们待会儿给出，包括[LearnYouAHeskellhttp://learnyouahaskell.com/chapters]
 * Haskell是纯函数的。这使得debug工作大量减少，但是你也不得不在书写简单的I/O和状态上下点儿工夫

# 一些Haskell代码中的常见模式
在继续深入之前，很有必要来了解一些有用的概念.
 * [http://en.wikipedia.org/wiki/Higher-order_function Higher-Order function]
 * [http://steve-yegge.blogspot.com/2006/03/execution-in-kingdom-of-nouns.html Verb-Centric]
首先，Haksell中没有数组，幸亏数组本来也不是一个好主意。 几乎所有的事情都是由列表完成的。甚至，字符串不过是字符列表而已。所以任何可以操作列表的函数，也可以操作字符串。
问题是，你该怎样遍历列表，然后做点自己的事情呢？递归! 大概是这样的过程:
 1. 从列表中得到第一个元素, 干点儿啥
 2. 对列表的其余部分进行上一步操作， 直到列表为空
绝大多数情况下，你逃不出以下三个模式:
 * map 
 * fold
 * filter
# 主要的资料
# 入门教程
 * [http://learnyouahaskell.com/ Learn you a Haskell for Great Good]
 * [http://lambda-the-ultimate.org/node/3642 Erik Meijer's Channel 9 Functional Programming Lectures]
 * [http://www.haskell.org/haskellwiki/Hitchhikers_guide_to_Haskell Hitchiker's Guide to Haskell]
 * [http://en.wikibooks.org/wiki/Write_Yourself_a_Scheme_in_48_Hours Write Yourself a Scheme in 48 Hours]
 * [http://www.haskell.org/haskellwiki/Roll_your_own_IRC_bot Roll your own IRC bot]
 * [http://lisperati.com/haskell/ Lisperati.com's Haskell Picnic]
# 书
 * [http://en.wikibooks.org/wiki/Haskell The Haskell Wikibook]
 * [http://book.realworldhaskell.org/read/ Real World Haskell]
# 视频
 * [http://video.s-inf.de/#FP.2005-SS-Giesl.(COt).HD_Videoaufzeichnung Professor Giesl's Videos]
 * [http://lambda-the-ultimate.org/node/3642 Erik Meijer's Channel 9 Functional Programming Lectures]
 * [http://acm.wustl.edu/functional/simonsvideos.php Simon Peyton-Jones's Videos]
# IRC
 * #haskell on irc.freenode.net

# 循序渐进学习Haskell
这些步骤没有严格意义上的先后顺序。每个人学习语言的方式有所不同，所以你得自己找到适合的教程或是书籍。下面这些是推荐顺序。
# 语法
 * http://www.cs.utep.edu/cheon/cs3360/pages/haskell-syntax.html
# 常用函数 + 基本编程技巧
 * [http://learnyouahaskell.com/chapters] (1-2章) 或者 [http://book.realworldhaskell.org/read/] (1-7章). (熟悉 [http://hackage.haskell.org/packages/archive/base/4.0.0.0/doc/html/Prelude.html Prelude]函数)
 * [http://acm.wustl.edu/functional/hs-breads.php Breadcrumbs RPN以上的部分]
 * [http://www.haskell.org/haskellwiki/H-99:_Ninety-Nine_Haskell_Problems  H99] 99个Haskell问题，同一个问题的解决方法不只一种，挑战一下。试试不同的Higher-Order函数。

# 惰性求值
```
Laziness isn't that tricky of a concept, but it can lead to some pretty strange (and beautiful) things so make sure you have a good grasp of it. If you read any book or full-language tutorial it's guaranteed to talk about laziness, but for the lazy (hah!) the following resources offer concise introductory explanations of Haskell's laziness (in order of depth):
```
一下资料由浅入深:
 * [http://book.realworldhaskell.org/read/types-and-functions.html#id580425 RWH] (第二章)
 * [http://www.haskell.org/haskellwiki/Haskell/Lazy_evaluation THW上的惰性求值]
 * [http://blog.interlinked.org/tutorials/haskell_laziness.html 一篇不错的教程]
 * [http://en.wikibooks.org/wiki/Haskell/Laziness Wikibook/Laziness]
高级货:
 * 使用惰性求值定义序列(Using laziness to define sequences)
    * http://acm.wustl.edu/functional/hs-breads.php#infinity
    * http://acm.wustl.edu/functional/hs-breads.php#advancedinfinitelists
    * http://mitpress.mit.edu/sicp/full-text/book/book-Z-H-24.html#%_sec_3.5.4 (SCIP Streem 你可以改写成Haskell)
    * File I/O: 惰性求值之于I/O另有妙处。如果你使用`text <- getContesnts`求值是立刻发生的。不妥之处在于有时候你必须考虑是不是在读一个巨大的文件，上面这种方法很可能会用尽你的内存。
# 类型
从下面这些材料入手:
 * [http://book.realworldhaskell.org/read/types-and-functions.html RWH第二章]
 * [http://learnyouahaskell.com/types-and-typeclasses LYAH]
直到你能够看懂:
```hs
foo :: Int -> [Char] -> (a,b) -> ((a,b) -> Int -> Char) -> [Char]
```
然后你就该试着定义自己的类型了。
 * [http://learnyouahaskell.com/making-our-own-types-and-typeclasses#algebraic-data-types LYAH]，搞明白Typeclasses和Types
 * [http://www.haskell.org/haskellwiki/Newtype newtype]运算符
 * GHCi中记得使用'`:i <TypeName>`'来观察类型之定义. 比如:
```hs
> :i Int
data Int = GHC.Types.I# GHC.Prim.Int#   -- Defined in `GHC.Types'
instance Bounded Int -- Defined in `GHC.Enum'
instance Enum Int -- Defined in `GHC.Enum'
instance Eq Int -- Defined in `GHC.Classes'
instance Integral Int -- Defined in `GHC.Real'
instance Num Int -- Defined in `GHC.Num'
instance Ord Int -- Defined in `GHC.Classes'
instance Read Int -- Defined in `GHC.Read'
instance Real Int -- Defined in `GHC.Real'
instance Show Int -- Defined in `GHC.Show
```
# Monads
接下来的重头戏是Monads. Monad是一个Typeclass. 这是必须首先了解的实事。
如果你没有掌握Haskell的核心，理解Monad的例子是非常困难的. 可以从以下材料入手:
 * http://haskell.org/haskellwiki/Tutorials#Recommended_tutorials (8.1和8.3)
理解Monads可分为三步:
 1. 为什么会出现Monads
  * http://www.iterasi.net/openviewer.aspx?sqrlitid=ixx7fcluvek_9lfolsxr_g
 2. 实际代码中如何运用Monads
  * http://www.haskell.org/tutorial/monads.html
  * http://www.haskell.org/all_about_monads/html/index.html
  * http://en.wikibooks.org/wiki/Haskell
  * http://acm.wustl.edu/functional/state-monad.php
 3. 人们可以使用Mondads完成哪些有趣儿的工作
  * http://blog.sigfpe.com/2009/01/haskell-monoids-and-their-uses.html
  * http://www.randomhacks.net/articles/2007/03/12/monads-in-15-minutes
  * http://www.randomhacks.net/articles/2007/02/21/refactoring-probability-distributions
  * [http://blog.sigfpe.com/2008/12/mother-of-all-monads.html Monads之母]
  * http://community.livejournal.com/evan_tech/220036.html
```
Before you start writing your own monads though, you should actually skip to our next section and do some I/O stuff as a way to practice using monads. 
```
 * http://www.haskell.org/haskellwiki/All_About_Monads (搞懂里面的例子)
继续后面的学习之前，请务必确认你可以做到:
 * 理解mapM/forM/foldM,sequence,liftM
 * 可以实现自定义monad
 * 可以使用内置StateMonad编写代码
 * 理解[http://acm.wustl.edu/functional/hs-breads.php#bind bind]和[http://acm.wustl.edu/functional/hs-breads.php#bind list monad]
 * 理解[http://acm.wustl.edu/functional/hs-breads.php#customstatemonad CustomStateMonad]
 * 看看这三个项目
  * http://acm.wustl.edu/functional/projects.php#snobol
  * http://acm.wustl.edu/functional/projects.php#sierpinski
  * http://acm.wustl.edu/functional/projects.php#bayesian-spam
# I/O
 * [http://learnyouahaskell.com/input-and-output LYH第9章]
 * [http://book.realworldhaskell.org/read/io.html RWH第7章]
 * [http://www.haskell.org/haskellwiki/Haskell_IO_for_Imperative_Programmers Haskell IO For Imperative Programmers]
# 高级货
 * [http://en.wikibooks.org/wiki/Haskell/Monad_transformers MonadTransformers]
 * [http://www.haskell.org/haskellwiki/Blow_your_mind Haskell "idioms"]
 * Arrows
 * [http://hackage.haskell.org/packages/hackage.html Hackage]
 * [http://acm.wustl.edu/functional/projects.php 一些有趣儿的项目]
# 原文
 * http://acm.wustl.edu/functional/haskell.php
# 参考
 * http://yannesposito.com/Scratch/en/blog/Haskell-the-Hard-Way/
 * http://www.seas.upenn.edu/~cis552/13fa/schedule.html
# 1. 推荐工具
Almost all new Haskell projects use the following tools. Each is intrinsically useful, but using a set of common tools also helps everyone by increasing productivity, and you're more likely to get patches.
# 1.1 版本控制
 * [wiki:Git]
 * [wiki:Darcs]
Use git or darcs unless you have a specific reason not to. Both are lightweight distributed revision control systems (and darcs is written in Haskell). Both have massive market share in the Haskell world, if you want to encourage contributions from other Haskell hackers git or darcs are the best. Darcs hosting is available on code.haskell.org and patch-tag. github for git is very popular.
# 1.2 编译系统
# Cabal
Use Cabal. You should read at least the start of section 2 of the Cabal User's Guide.
You should use cabal-install as a front-end for installing your Cabal library. Cabal-install provides commands not only for building libraries but also for installing them from, and uploading them to, Hackage. As a bonus, for almost all programs, it's faster than using Setup.hs scripts directly, since no time is wasted compiling the scripts. (This does not apply for programs that use custom Setup.hs scripts, since those need to be compiled even when using cabal-install.)
cabal-install is widely available, as part of the Haskell Platform, so you can probably assume your users will have it too.
# 1.3 文档
 * [wiki:Haddock]
For libraries, use Haddock. We recommend using the version of Haddock that ships with the Haskell Platform. Haddock generates nice markup, with links to source.
# 1.4 测试
 * QuickCheck
 * SmallCheck
You can use QuickCheck or SmallCheck to test pure code. To test impure code, use HUnit.
To get started, try Introduction to QuickCheck. For a slightly more advanced introduction, Simple Unit Testing in Haskell is a blog article about creating a testing framework for QuickCheck using some Template Haskell. For HUnit, see HUnit 1.0 User's Guide
# 1.5 发布
The standard mechanism for distributing Haskell libraries and applications is Hackage. Hackage can host your cabalised tarball releases, and link to any library dependencies your code has. Users will find and install your packages via "cabal install", and your package will be integrated into Haskell search engines, like hoogle
# 1.6 目标环境
If at all possible, depend on libraries that are provided by the Haskell Platform, and libraries that in turn build against the Haskell Platform. This set of libraries is designed to be widely available, so your end users will be able to build your software.
# 2 一个工程的结构
The basic structure of a new Haskell project can be adopted from HNop, the minimal Haskell project. It consists of the following files, for the mythical project "haq".
```
Haq.hs    -- the main haskell source file
haq.cabal -- the cabal build description
Setup.hs  -- build script itself
_darcs    -- revision control
README    -- info
LICENSE   -- license
```
Of course, you can elaborate on this, with subdirectories and multiple modules. See Structure of a Haskell project for an example of a larger project's directory structure.
Here is a transcript that shows how you'd create a minimal darcs and cabalised Haskell project for the cool new Haskell program "haq", build it, install it and release.
Note: The new tool "cabal init" automates all this for you, but you should understand all the parts even so.
We will now walk through the creation of the infrastructure for a simple Haskell executable. Advice for libraries follows after.
# 2.1 Create a directory
Create somewhere for the source:
```
#!sh
$ mkdir haq
$ cd haq
```
# 2.2 编写Haskell代码
Write your program:
```
#!sh
$ cat > Haq.hs
--
-- Copyright (c) 2006 Don Stewart - http://www.cse.unsw.edu.au/~dons/
-- GPL version 2 or later (see http://www.gnu.org/copyleft/gpl.html)
--
import System.Environment
 
-- | 'main' runs the main program
main :: IO ()
main = getArgs >>= print . haqify . head
 
haqify s = "Haq! " ++ s
```
# 2.3 Stick it in darcs
Place the source under revision control (you may need to enter your e-mail address first, to identify you as maintainer of this source):
$ darcs init
$ darcs add Haq.hs 
$ darcs record
addfile ./Haq.hs
Shall I record this change? (1/?)  [ynWsfqadjkc], or ? for help: y
hunk ./Haq.hs 1
+--
+-- Copyright (c) 2006 Don Stewart - http://www.cse.unsw.edu.au/~dons/
+-- GPL version 2 or later (see http://www.gnu.org/copyleft/gpl.html)
+--
+import System.Environment
+
+-- | 'main' runs the main program
+main :: IO ()
+main = getArgs >>= print . haqify . head
+
+haqify s = "Haq! " ++ s
Shall I record this change? (2/?)  [ynWsfqadjkc], or ? for help: y
What is the patch name? Import haq source
Do you want to add a long comment? [yn]n
Finished recording patch 'Import haq source'
And we can see that darcs is now running the show:
 $ ls
Haq.hs _darcs
# 2.4 加入编译系统
Create a .cabal file describing how to build your project:
```
$ cat > haq.cabal
Name:                haq
Version:             0.0
Description:         Super cool mega lambdas
License:             GPL
License-file:        LICENSE
Author:              Don Stewart
Maintainer:          dons@cse.unsw.edu.au
Build-Type:          Simple
Cabal-Version:       >=1.2
Executable haq
  Main-is:           Haq.hs
  Build-Depends:     base >= 3 && < 5
(If your package uses other packages, e.g. haskell98, you'll need to add them to the Build-Depends: field as a comma separated list.) Add a Setup.hs that will actually do the building:
$ cat > Setup.hs
import Distribution.Simple
main = defaultMain
Cabal allows either Setup.hs or Setup.lhs.
```
Now would also be a good time to add a LICENSE file and a README file. Examples are in the tarball for HNop.
Record your changes:
```
$ darcs add haq.cabal Setup.hs LICENSE README
$ darcs record --all
What is the patch name? Add a build system
Do you want to add a long comment? [yn]n
Finished recording patch 'Add a build system'
```
# 2.5 编译工程
Now build it! There are two methods of accessing Cabal functionality: through your Setup.hs script or through cabal-install. In most cases, cabal-install is now the preferred method.
使用[wiki:Haskell/Cabal]编译:
```
#!sh
$ cabal install --prefix=$HOME --user
```
使用传统的Setup.hs进行构建:
```
#!sh
$ runhaskell Setup configure --prefix=$HOME --user
$ runhaskell Setup build
$ runhaskell Setup install
```
This will install your newly minted haq program in $HOME/bin.
# 2.6 运行
And now you can run your cool project:
```
#!sh
$ haq me
"Haq! me"
```
You can also run it in-place, even if you skip the install phase:
 $ dist/build/haq/haq you
"Haq! you"
# 2.7 构建文档
Generate some API documentation into dist/doc/*
Using cabal install:
```
#!sh
$ cabal haddock
```
Traditional method: 
```
#!sh
$ runhaskell Setup haddock
```
which generates files in dist/doc/ including:
```
#!sh
$ w3m -dump dist/doc/html/haq/Main.html
haq Contents Index
Main
Synopsis
main :: IO ()
Documentation
main :: IO ()
main runs the main program
Produced by Haddock version 0.7
```
No output? Make sure you have actually installed haddock. It is a separate program, not something that comes with Cabal. Note that the stylized comment in the source gets picked up by Haddock.
# 2.8 (Optional) 改进你的代码: HLint
HLint can be a valuable tool for improving your coding style, particularly if you're new to Haskell. Let's run it now.
 $ hlint .
./Haq.hs:11:1: Warning: Eta reduce
Found:
  haqify s = "Haq! " ++ s
Why not:
  haqify = ("Haq! " ++)
The existing code will work, but let's follow that suggestion. Open Haq.hs in your favourite editor and change the line:
where haqify s = "Haq! " ++ s
to:
where haqify = ("Haq! " ++)
# 2.9 来一点自动测试: QuickCheck
# 2.9.1 QuickCheck v1
We'll use QuickCheck to specify a simple property of our Haq.hs code. Create a tests module, Tests.hs, with some QuickCheck boilerplate:
$ cat > Tests.hs
import Char
import List
import Test.QuickCheck
import Text.Printf

main  = mapM_ ((s,a) -> printf "%-25s: " s >> a) tests

instance Arbitrary Char where
    arbitrary     = choose (' ', '
8')
    coarbitrary c = variant (ord c `rem` 4)
Now let's write a simple property:
$ cat >> Tests.hs 
-- reversing twice a finite list, is the same as identity
prop_reversereverse s = (reverse . reverse) s == id s
    where _ = s :: [Int]

-- and add this to the tests list
tests  = [("reverse.reverse/id", test prop_reversereverse)]
We can now run this test, and have QuickCheck generate the test data:
 $ runhaskell Tests.hs
reverse.reverse/id       : OK, passed 100 tests.
Let's add a test for the 'haqify' function:
-- Dropping the "Haq! " string is the same as identity
prop_haq s = drop (length "Haq! ") (haqify s) == id s
    where haqify s = "Haq! " ++ s

tests  = [("reverse.reverse/id", test prop_reversereverse)
        ,("drop.haq/id",        test prop_haq)]
and let's test that:
 $ runhaskell Tests.hs
reverse.reverse/id       : OK, passed 100 tests.
drop.haq/id              : OK, passed 100 tests.
Great!
# 2.9.2 QuickCheck v2
If you're using version 2 of QuickCheck, the code in the previous section needs some minor modifications:
```
#!hs
import Char
import List
import Test.QuickCheck
import Text.Printf
 
main  = mapM_ ((s,a) -> printf "%-25s: " s >> a) tests
 
-- reversing twice a finite list, is the same as identity
prop_reversereverse s = (reverse . reverse) s == id s
    where _ = s :: [Int]
 
-- Dropping the "Haq! " string is the same as identity
prop_haq s = drop (length "Haq! ") (haqify s) == id s
    where haqify s = "Haq! " ++ s
 
tests  = [("reverse.reverse/id", quickCheck prop_reversereverse)
        ,("drop.haq/id",        quickCheck prop_haq)]
To run the test:
```
```
#!sh
$ runhaskell Tests.hs
reverse.reverse/id       : +++ OK, passed 100 tests.
drop.haq/id              : +++ OK, passed 100 tests.
Success!
```
# 2.10 Running the test suite from darcs
We can arrange for darcs to run the test suite on every commit that is run with the flag --test:
 $ darcs setpref test "runhaskell Tests.hs"
Changing value of test from  to 'runhaskell Tests.hs'
will run the full set of QuickChecks. If your test requires it, you may need to ensure other things are built too -- for example:darcs setpref test "alex Tokens.x;happy Grammar.y;runhaskell Tests.hs". You will encounter that this way a darcs patch is also accepted if a QuickCheck test fails. You have two choices to work around this:
Use quickCheck' from the package QuickCheck-2 and call exitWithFailure if it return False.
Keep the test program as it is, and implement the failure on the shell level:
runhaskell Tests.hs | tee test.log && if grep Falsifiable test.log >/dev/null; then exit 1; fi
Let's commit a new patch:
$ darcs add Tests.hs
$ darcs record --all --test
What is the patch name? Add testsuite
Do you want to add a long comment? [yn]n
Running test...
reverse.reverse/id       : OK, passed 100 tests.
drop.haq/id              : OK, passed 100 tests.
Test ran successfully.
Looks like a good patch.
Finished recording patch 'Add testsuite'
Excellent: now, patches must pass the test suite before they can be committed provided the --test flag is passed.
# 2.11 Tag the stable version, create a tarball, and sell it!
Tag the stable version:
 $ darcs tag
What is the version name? 0.0
Finished tagging patch 'TAG 0.0'
2.11.1 Create a tarball
You can do this using either Cabal or darcs, or even an explicit tar command.
2.11.1.1 Using Cabal
Since the code is cabalised, we can create a tarball with cabal-install directly (you can also use runhaskell Setup.hs sdist, but you need tar on your system [1]):
 $ cabal sdist
Building source dist for haq-0.0...
Source tarball created: dist/haq-0.0.tar.gz
This has the advantage that Cabal will do a bit more checking, and ensure that the tarball has the structure that HackageDB expects. Note that it does require the LICENSE file to exist. It packages up the files needed to build the project; to include other files (such as Test.hs in the above example, and our README), we need to add:
extra-source-files: Tests.hs README
to the .cabal file to have everything included.
2.11.1.2 Using darcs
Alternatively, you can use darcs:
 $ darcs dist -d haq-0.0
Created dist as haq-0.0.tar.gz
And you're all set up!
2.11.2 Check that your source package is complete
Just to make sure everything works, try building the source package in some temporary directory:
 $ tar xzf haq-0.0.tar.gz
$ cd haq-0.0
$ cabal configure
$ cabal build
and for packages containing libraries,
 $ cabal haddock
2.11.3 Upload your package to Hackage
Whichever of the above methods you've used to create your package, you can upload it to the Hackage package collection via a web interface. You may wish to use the package checking interface there first, and fix things it warns about, before uploading your package.
2.12 Summary
The following files were created:
   $ ls
   Haq.hs           Tests.hs         dist             haq.cabal
   Setup.hs         _darcs           haq-0.0.tar.gz
3 Libraries
The process for creating a Haskell library is almost identical. The differences are as follows, for the hypothetical "ltree" library:
3.1 Hierarchical source
The source should live under a directory path that fits into the existing module layout guide. So we would create the following directory structure, for the module Data.LTree:
   $ mkdir Data
   $ cat > Data/LTree.hs 
   module Data.LTree where
So our Data.LTree module lives in Data/LTree.hs
3.2 The Cabal file
Cabal files for libraries list the publically visible modules, and have no executable section:
   $ cat > ltree.cabal 
   Name:                ltree
   Version:             0.1
   Description:         Lambda tree implementation
   License:             BSD3
   License-file:        LICENSE
   Author:              Don Stewart
   Maintainer:          dons@cse.unsw.edu.au
   Build-Type:          Simple
   Cabal-Version:       >=1.2

   Library
     Build-Depends:     base >= 3 && < 5
     Exposed-modules:   Data.LTree
     ghc-options:       -Wall
We can thus build our library:
   $ cabal configure --prefix=$HOME --user
   $ cabal build    
   Preprocessing library ltree-0.1...
   Building ltree-0.1...
   [1 of 1] Compiling Data.LTree       ( Data/LTree.hs, dist/build/Data/LTree.o )
   /usr/bin/ar: creating dist/build/libHSltree-0.1.a
and our library has been created as a object archive. Now install it:
   $ cabal install
   Installing: /home/dons/lib/ltree-0.1/ghc-6.6 & /home/dons/bin ltree-0.1...
   Registering ltree-0.1...
   Reading package info from ".installed-pkg-config" ... done.
   Saving old package config file... done.
   Writing new package config file... done.
And we're done! To try it out, first make sure that your working directory is anything but the source directory of your library:
   $ cd ..
And then use your new library from, for example, ghci:
   $ ghci -package ltree
   Prelude> :m + Data.LTree
   Prelude Data.LTree> 
The new library is in scope, and ready to go.
3.3 More complex build systems
For larger projects, you may want to store source trees in subdirectories. This can be done simply by creating a directory -- for example, "src" -- into which you will put your src tree.
To have Cabal find this code, you add the following line to your Cabal file:
   hs-source-dirs: src
You can also set up Cabal to run configure scripts, among other features. For more information consult the Cabal user guide.
4 Automation
A tool to automatically populate a new cabal project is available:
   cabal init
Usage is:
 $ cabal init
Package name [default "haq"]? 
Package version [default "0.1"]? 
Please choose a license:
   1) GPL
   2) GPL-2
   3) GPL-3
   4) LGPL
   5) LGPL-2.1
   6) LGPL-3
 * 7) BSD3
   8) BSD4
   9) MIT
    10) PublicDomain
    11) AllRightsReserved
    12) OtherLicense
    13) Other (specify)
Your choice [default "BSD3"]? 
Author name? Henry Laxen
Maintainer email? nadine.and.henry@pobox.com
Project homepage/repo URL? http://somewhere.com/haq/
Project synopsis? A wonderful little module
Project category:
   1) Codec
   2) Concurrency
   3) Control
   4) Data
   5) Database
   6) Development
   7) Distribution
   8) Game
   9) Graphics
    10) Language
    11) Math
    12) Network
    13) Sound
    14) System
    15) Testing
    16) Text
    17) Web
    18) Other (specify)
Your choice? 3
What does the package build:
   1) Library
   2) Executable
Your choice? 1
Generating LICENSE...
Generating Setup.hs...
Generating haq.cabal...
You may want to edit the .cabal file and add a Description field.
5 Licenses
Code for the common base library package must be BSD licensed. Otherwise, it is entirely up to you as the author. Choose a licence (inspired by this). Check the licences of things you use (both other Haskell packages and C libraries), since these may impose conditions you must follow. Use the same licence as related projects, where possible. The Haskell community is split into 2 camps, roughly: those who release everything under BSD, and (L)GPLers. Some Haskellers recommend avoiding LGPL, due to cross-module optimisation issues. Like many licensing questions, this advice is controversial. Several Haskell projects (wxHaskell, HaXml, etc) use the LGPL with an extra permissive clause which gets round the cross-module optimisation problem.
6 Releases
It's important to release your code as stable, tagged tarballs. Don't just rely on darcs for distribution.
darcs dist generates tarballs directly from a darcs repository
For example:
$ cd fps
$ ls       
Data      LICENSE   README    Setup.hs  TODO      _darcs    cbits dist      fps.cabal tests
$ darcs dist -d fps-0.8
Created dist as fps-0.8.tar.gz
You can now just post your fps-0.8.tar.gz
You can also have darcs do the equivalent of 'daily snapshots' for you by using a post-hook.
put the following in _darcs/prefs/defaults:
 apply posthook darcs dist
 apply run-posthook
Advice:
Tag each release using darcs tag. For example:
$ darcs tag 0.8
Finished tagging patch 'TAG 0.8'
Then people can darcs pull --partial -t 0.8, to get just the tagged version (and not the entire history).
7 Hosting
Hosting for repos is available from the Haskell community server:
   http://community.haskell.org/
A Darcs repository can be published simply by making it available from a web page.
8 Web page
Create a web page documenting your project! An easy way to do this is to add a project specific page to the Haskell wiki
9 The user experience
When developing a new Haskell library, it is important to remember how the user expects to be able to build and use a library.
9.1 Introductory information and build guide
A typical library user expects to:
Visit Haskell.org
Find the library/program they are looking for:
if not found, try mailing list;
if it is hidden, try improving the documentation on haskell.org;
if it does not exist, try contributing code and documentation)
Download
Build and install
Enjoy
Each of these steps can pose potential road blocks, and code authors can do a lot to help code users avoid such blocks. Steps 1..2 may be easy enough, and many coders and users are mainly concerned with step 5. Steps 3..4 are the ones that often get in the way. In particular, the following questions should have clear answers:
Which is the latest version?
What state is it in?
What are its aims?
Where is the documentation?
Which is the right version for given OS and Haskell implementation?
How is it packaged, and what tools are needed to get and unpack it?
How is it installed, and what tools are needed to install it?
How do we handle dependencies?
How do we provide/acquire the knowledge and tool-chains needed?
The best place to answer these questions is a README file, distributed with the library or application, and often accompanied with similar text on a more extensive web page.
9.2 Tutorials
Generated haddock documentation is usually not enough to help new programmers learn how to use a library. You must also provide accompanying examples, and even tutorials about the library.
Please consider providing example code for your library or application. The code should be type-correct and well-commented.
10 Program structure
Monad transformers are very useful for programming in the large, encapsulating state, and controlling side effects. To learn more about this approach, try Monad Transformers Step by Step.
11 Publicity
The best code in the world is meaningless if nobody knows about it. The process to follow once you've tagged and released your code is:
11.1 Join the community
If you haven't already, join the community. The best way to do this is to subscribe to at least haskell-cafe@ and haskell@ mailing lists. Joining the #haskell IRC channel is also an excellent idea.
11.2 Announce your project on haskell@
Most important: announce your project releases to the haskell@haskell.org mailing list. Tag your email subject line with "ANNOUNCE: ...". This ensure it will then make it into the Haskell Weekly News. To be doubly sure, you can email the release text to the HWN editor.
11.3 Add your code to the public collections
Add your library or application to the Libraries and tools page, under the relevant category, so people can find it.
If your release is a Cabal package, add it to the Hackage database (Haskell's CPAN wanna-be).
11.4 Blog about it
Blog about it! Blog about your new code on Planet Haskell. Write about your project in your blog, then email the Planet Haskell maintainer (ibid on #haskell) the RSS feed url for your blog
12 Example
A complete example of writing, packaging and releasing a new Haskell library under this process has been documented.
# 参考
 *[http://www.haskell.org/haskellwiki/How_to_write_a_Haskell_program 原文]
# Infix Operators
中缀操作符
Infix operators are really just functions, 
and can also be defined using equations. 
For example, here is a definition of a list concatenation operator:
Infix Operators实际上是函数，也可以用方程式定义。
# Sections
因为以上原因，Infix Operators可以接收部分参数，这样不完整的Infix Operators叫做Sections
```hs
(+)  =  y -> x + y
(x+) = y -> x + y
(+y) =  > x + y
Prelude> map (+2) [1,2,3]
[3,4,5]
Prelude> map (2/) [1]
[2.0]
Prelude> map (/2) [1]
[0.5]
```
中缀操作符也可以这么用:
```hs
1 `add` 2
# 相当于
add 1 2
Prelude> mod 3 7
3
Prelude> 3 `mod` 7
3
x `elem` xs 要比 elem x xs 更直观，易懂
```
这纯粹是因为某些操作使用这种表达方式更加直观，易懂。
# Fixity Declarations ==
可以给Infix Operators指定优先级(precedence), 由0－9优先级逐渐增高
```hs
infixr 5 ++
infixr 6 -
infixl 6 -
```
[[TOC]]
# Haskell List =
# 使用`':'`构造List
下面演示了Haskell是如何建立List的, ![1,2,3,4]这样的List实际上等价于1:2:3:4:[].
```
#!hs
Prelude> let list=[]
Prelude> list
[]
Prelude> 0:1:2:3:list
[0,1,2,3]
```
# List `++` List
# List中的List
既然List可以包含一切, 当然也可以包含List, 需要注意的是: List中仍然只能容纳相同类型的元素或空List(`[]`).
```
#!hs
Prelude> [[1,2], ["a","b"]]
<interactive>:1:5:
    No instance for (Num [Char])
      arising from the literal `2'
    Possible fix: add an instance declaration for (Num [Char])
    In the expression: 2
    In the expression: [1, 2]
    In the expression: [[1, 2], ["a", "b"]]
Prelude> [[1,2], [1,2,3]]
[[1,2],[1,2,3]]
Prelude> [[1,2], [1,2,3], [1]]
[[1,2],[1,2,3],[1]]
```
```
#!div class=warn
# List中包含的每个元素必须为相同类型
如果违背这一规则, 情形如下:
```
#!hs
Prelude> let mixedList=[0, "hello"]
<interactive>:1:16:
    No instance for (Num [Char])
      arising from the literal `0'
    Possible fix: add an instance declaration for (Num [Char])
    In the expression: 0
    In the expression: [0, "hello"]
    In an equation for `mixedList': mixedList = [0, "hello"]
```

```
```
#!sh
ghci> let mylist=[1,2,3,4,5,6,7]
ghci> mylist 
[1,2,3,4,5,6,7]
ghci> [1,2] ++ [3,4]
[1,2,3,4]
ghci> ['h','o','l','a']
"hola"
ghci> "ho" ++ "la"
"hola" 
# 加入一个新元素
$ 1:[2,2,3]
[1,2,2,3]
```
# <list> `!!` <number> 按索引取出
```hs
$ [0,1,2,3,4] !! 1
1
# 越界啦
$ [0,1,2,3,4] !! 5
*** Exception: Prelude.(!!): index too large
```
# List之间的关系运算
```hs
# 怎样比较List哇？
$ [1,2,3] > [1,2]
True
$ [1,2,3] > [1,2,2]
True
$ [1,2,3] > [1,2,4]
False
$ [1,2,3] > [3,2]
False
```
# head <list> :取头
```sh
$ head [1, 2, 3] 
1
```
# tail <list> : 取尾
```
# 取尾, 注意啦，这个相当于Lisp中的cdr, 想取最后一个元素？用last
$ tail [1, 2, 3]
[2,3]
```
# last <list> : 取最后一个元素
```hs
$ last [1, 2, 3]
3
```
# init <list> : 除了最后一个我都要
```hs
$ init [5,4,3,2,1]  
[5,4,3,2]
```
# null <list> : 列表是空的么?
```hs
$ null [1,2,3]
False
# 字符串也是序列的一种，所以
$ length "1234" == 0
False
$ null "1234"
False
$ null ""
True
```
# reverse <list> : 逆序
```hs
# 翻转
$ reverse [0, 1, 2, 4, 5]
[5,4,2,1,0]
```
# take <number> <list> : 取出前number个元素
```hs
# 取前N个元素
$ take 1 [0, 1, 2, 3]
[0]
$ take 2 [0, 1, 2, 3]
[0,1]
$ take 3 [0, 1, 2, 3]
[0,1,2]
$ take 5 [0, 1, 2, 3]
[0,1,2,3]
```
# drop <number> <list> : 丢弃前<number>个元素
```hs
# 丢弃前N个元素
$ drop 0  [0, 1, 2, 3]
[0,1,2,3]
$ drop 1 [0, 1, 2, 3]
[1,2,3]
$ drop 2 [0, 1, 2, 3]
[2,3]
$ drop 5 [0, 1, 2, 3]
[]
```
# 取最大元素
$ maximum [1, -2, 4, 100]
100
# 取最小元素
$ minimum  [1, -2, 4, 100]
-2
# 求和(仅对Number)
$ sum [1,2,3]
6
# 求积
$ product [6,2,1,2]  
24
$ product [6,2,1,2,0]  
0
# 这个元素是不是在List中? 
$ 10 `elem` [1,2,3]
False
$ 10 `elem` [1,2,3,10]
True
# 生成序列, python中的range(1,10), zsh中的 {1..10}, 都是类似的概念
$ [1..10]
[1,2,3,4,5,6,7,8,9,10]
$ ['A'..'Z']
"ABCDEFGHIJKLMNOPQRSTUVWXYZ"
$ ['A'..'z']
"ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_`abcdefghijklmnopqrstuvwxyz"
# 步长
$ [10,9..1] 
[10,9,8,7,6,5,4,3,2,1]
# 注意: 1.步长为9-10=-1,上限=1,  假如: [1,3..10] 步长为3-1=2, 上限=10, 所以 2-10内所有偶数
$ [2,4..10]
[2,4,6,8,10]
# 注意: 2. 步长=0, 可能会出现两种情况
# 2.1 上限大于起点，就像原地踏步, 无限循环
# 2.3 返回空List
$ [2,2..10]
# 当心误差， 虽然range/step可以应用于浮点数，但最好不要这样
$ [0.1, 0.3 .. 1] 
[0.1,0.3,0.5,0.7,0.8999999999999999,1.0999999999999999]
# Haskell是懒惰的
$ [2,4..] 
... 无限循环,列表不断增长
$ take 10 [2,4..]
[2,4,6,8,10,12,14,16,18,20]
# 不难看出，当需要打印一个无限长的列表时，循环就无限进行下去，只想take 10 个元素时，就循环10次
# range/setp产生一个集合，(x*2 | x属于N, x小于10)
$ [x*2 | x <- [1..10]]  
[2,4,6,8,10,12,14,16,18,20]
# '|'以前的部分(x*2)叫做OutputFunction
$ [x*2 | x <- [1..10], x*2 >= 12] 
# (x*2 >= 12)叫做Predicate
[12,14,16,18,20]
# 结合以上所述，让我们来找出1到100以内可以被5整除的整数集合
$ [x | x<-[1..100], x `mod` 5 == 0]
[5,10,15,20,25,30,35,40,45,50,55,60,65,70,75,80,85,90,95,100]
# 多个Predicate
$ [ x | x <- [10..20], x /= 13, x /= 15, x /= 19]  
[10,11,12,14,16,17,18,20]
# Draw from several lists
$ [ x*y | x <- [2,5,10], y <- [8,10,11]]  
[16,20,22,40,50,55,80,100,110]
# 笛卡尔集
$ let colors=["red", "blue", "green"]
$ let cars=["Benz200", "QQ", "212"]
$ [color ++ " " ++ car | color <- colors, car <- cars]
["red Benz200","red QQ","red 212","blue Benz200","blue QQ","blue 212","green Benz200","green QQ","green 212"]
# 再来看个例子： 计算List的长度
$ sum [1 | _ <- [1..100]]
100
# 1. [1 | _ <- [1..100]] 产生了一个新的集合， 这个集合中每个元素都是1, 总共100个， ('_'表示这个变量我们不关心)
# 2. 求sum便得到了长度
# 写成函数: length' xs = sum [1 | _ <- xs] 
# 下面是个更复杂的例子, 从复杂列滤出全部偶数
$ let xxs = [[1,3,5,2,3,1,2,4,5],[1,2,3,4,5,6,7,8,9],[1,2,4,2,1,6,3,1,3,2,3,6]]  
$ [ [ x | x <- xs, even x ] | xs <- xxs]  
$ [[2,2,4],[2,4,6,8],[2,4,2,6,2,6]]   
```
!`elem`学名叫InfixFunction,
# map ==
 1. L1 = [ a1, a2, ... an ]
 2. f(a) = b , a 属于L1
 3. 对L1每个元素ax调用f(ax)得到结果bx， bx构成一个新的List L2=[ b1, b2, ... bn]
如果对某个List中的每个元素a进行同一种计算f(a)，结果为b, 那么就可以得到一个新的List, map就是基于这种想法。
```
#!sh
# map :: (a -> b) -> [a] -> [b]
# 将小写字母转换为大写字母
Prelude> map Char.toUpper "haskell"
"HASKELL"
```
# filter ==
filter从指定List中选取符合条件的元素(或 从List中过滤掉不符合条件的元素)
 * `filter :: (a -> Bool) -> [a] -> [a]`
```
#!sh
Prelude> filter Char.isLower "aAbBcCdD"
"abcd"
```
# foldr / foldl ==
先来说说foldr, 请看:
```
#!sh
foldr (-) 5 [1,2,3,4]
# 演算步骤
> foldr (-) 5 [1,2,3,4]
> 1 - (foldr (-) 5 [2,3,4])
> 1 - (2 - (foldr (-) 5 [3,4]))
> 1 - (2 - (3 - (foldr (-) 5 [4])))
> 1 - (2 - (3 - (4 - (foldr (-) 5 []))))
> 1 - (2 - (3 - (4 - (5))))
> 1 - (2 - (3 - (-1)))
> 1 - (2 - 4)
> 1 - (-2)
> 3
```
看了这附图之后，你就明白为什么叫foldr[ight] 了:
```
#!graphviz.dot
digraph G {
label="foldr (-) 5 [1,2,3,4]"
fontcolor="red"
node[shape="plaintext"]
edge[arrowhead="none"]
f0->1
f0->f1
f1->2
f1->f2
f2->3
f2->f3
f3->4
f3->5
f0[label="-"]
f1[label="-"]
f2[label="-"]
f3[label="-"]
}
```
再来计算foldl
```
#!sh
foldl (-) 5 [1,2,3,4]
# 演算步骤
> foldl (-) 5 [1,2,3,4]
> 
```

```
#!graphviz.dot
digraph G {
label="foldl (-) 5 [1,2,3,4]"
fontcolor="red"
node[shape="plaintext"]
edge[arrowhead="none"]
f0->f1
f0->4
f1->f2
f1->3
f2->f3
f2->2
f3->5
f3->1
f0[label="-"]
f1[label="-"]
f2[label="-"]
f3[label="-"]
}
```
# ++ ==
++用于连接两个Lists。
```
#!sh
Prelude> [1,2] ++ [3,4]
[1,2,3,4]
Prelude> "hello" ++ " " ++ "world!"
"hello world!"
```

```
Prelude> fst (5, "hello")
5
Prelude> snd (5, "hello")
"hello"
```
[[TOC]]
# zip xs (tail xs)
​```hs
> let xs = [1..9]
> zip xs (tail xs)
[(1,2),(2,3),(3,4),(4,5),(5,6),(6,7),(7,8),(8,9)]
```
斐波那契数列，又称黄金分割数列, 我们来看看怎样使用(zip/tail)定义它: 
```
1 1    -> zip / tail -> [(1,1)] -> map -> [1,1,2]
1 1 2  -> zip / tail -> [(1,1), (1,2)] -> map [1,1,2,3]
```
```hs
-- 当当当, 自然数
> let xs = [1..]
-- zip / tail 会发生什么?
> take 10 $ zip xs (tail xs)
[(1,2),(2,3),(3,4),(4,5),(5,6),(6,7),(7,8),(8,9),(9,10),(10,11)]
-- 呵呵, take / tail 可以将给定的list错开一位分组,  有点计算Fab的意思了
> let xs = [1, 1]
> zip xs (tail xs)
[(1, 1)]
> (map ((a,b) -> a + b) $ zip xs (tail xs))
> [2]
> let xs = [1, 1, 2]
> zip xs (tail xs)
[(1,1), (1,2)]
> (map ((a,b) -> a + b) $ zip xs (tail xs))
> [2, 3]
-- 假如有长度为n的斐波那契数列xs, 
-- 对其进行(map ((a,b) -> a + b) $ zip xs (tail xs)) 就可以得到
-- 长度为n+1的斐波那契数列
-- 我们知道最短的斐波那契数列就是[1,1]了, 因此可以得到迭代出下一个斐波那契数列的公式:
> let fibonacci' xs = (map ((a,b) -> a + b) $ zip xs (tail xs))
> fibonacci' [1,1]
[2]
> 1 : 1 : fibonacci' [1,1]
[1,1,2]
> 1 : 1 : fibonacci' [1,1,2]
[1,1,2,3]
...
-- 目前为止,借助fibonacci' 我们可以继续往下演算得到更长的斐波那契数列
-- 接下来, 借助一个优美的表达式, 可以重复上面的操作:
fibonacci = 1 : 1 : (fibonacci' fibonacci)
-- 最后,我们打算写成一行
fibonacci = 1 : 1 : (map ((a,b) -> a + b) $ zip fibonacci (tail fibonacci))
-- 生成一个长度为24的fab
> take 24 fibonacci
[1,1,2,3,5,8,13,21,34,55,89,144,233,377,610,987,1597,2584,4181,6765,10946,17711,28657,46368]
```
# Haskell Module =
# 怎样定义自己的模块 ==
下面的代码定义了MyModule模块,比如保存到my-module.hs文件中:
```
module MyModule
    where
-- Functions/Data/etc...
```
加载模块:
```
$ ghci
ghci> :l my-module.hs
MyModule>
# 重新载入模块
MyModule> :r
```
# Stand-alone Executables ==
.hs文件可以编译为单独的可执行文件，必要条件如下:
 1. 必须定义Main模块
 2. Main模块中必须有main函数
将以下内容保存到a.hs中:
```
module Main
    where
mkOddList x = [2,4..x]
main = putStrLn "Hello Amas!!!"
```
编译.hs文件:
```
#!sh
$ ghc --make a.hs -o a
$ ls
a a.hs
$ ./a
Hello Amas
```
[[TOC]]
# Monads
# The monad laws
 1. (return x) >>= f == f x
 2. m >>= return == m
 3. (m >>= f) >>= g == m >>= ( > f x >>= g)
# Monad Type Class in Haskell
```hs
class Monad m where
  (>>=) :: m a -> (a -> m b) -> m b
  (>>) :: m a -> m b -> m b
  return :: a -> m a
  fail :: String -> m a
        -- Defined in `GHC.Base'
instance Monad Maybe -- Defined in `Data.Maybe'
instance Monad (Either e) -- Defined in `Data.Either'
instance Monad [] -- Defined in `GHC.Base'
instance Monad IO -- Defined in `GHC.Base'
instance Monad ((->) r) -- Defined in `GHC.Base'
```
# >>=
```
(>>=) :: m a -> (a -> m b) -> m b
```
# >>
右面的函数不需要任何输入参数.
```
(>>) :: m a -> m b -> m b
m >> k = m >>= (_ -> k)
```
# return
# fail
# 从Monad中获取数据
通过模式匹配, 很容易将Monad中的数据取出来
```
> import Data.Maybe
> fromJust (Just 1)
1
```
# One-way Monads
One-way Monad允许值通过return或者fail进入到Monad内部, 但是不允许值从Monad中抽取回去. 这就是One-way Monads.这些值的运算只能通过`>>=`和`>>`在Monad中完成. 返回结果也必须带着Monad.
IO Monad就是这种类型. 
```
The wonderful feature of a one-way monad is that it can support side-effects in its monadic operations but prevent them from destroying the
functional properties of the non-monadic portions of the program.
```
# Andy Gill's Monad  Template Library
在GHC中, 这些Monad都属于`Control.Monad`的子模块
# Identity 
 * Identity monad在实际中基本不会用到
 * 这个最简单的Monad干的事儿你可能根本注意不到, 只是将函数的参数传递给函数
 * Monad变换定理中常常出现
```hs
newtype Identity a = Identity { runIdentity :: a }
instance Monad Identity where
    return a = Identity a
    (Identity x) >>= f = f x
-- i.e. return = id
-- i.e. x >>= f = f x
```
e.g:
A typical use of the Identity monad is to derive a monad from a monad transformer.
```hs
-- derive the State monad using the StateT monad transformer
type State s a = StateT s Identity a
```
# Maybe
 * 计算类型: 用于可能会返回Nothing值的运算.
 * 绑定策略: 转换过程中一旦出现Nothing值, 则计算结果为Nothing
 * 使用场景: 数据库查询等
定义:
```hs
data Maybe a = Nothing | Just a
instance Monad Maybe where
  return         = Just
  fail           = Nothing
  Nothing  >>= f = Nothing
  (Just x) >>= f = f x
instance MonadPlus Maybe where
  mzero             = Nothing
  Nothing `mplus` x = x
  x `mplus` _       = x
```
如果你的代码中出现了以下模式, 那么可以使用Maybe monad简化一下.
```hs
  case ... of
    Nothing -> Nothing
    Just a  -> case ... of
      Nothing -> Nothing
      Just a  -> case ... of
        Nothing -> Nothing
        Just a  -> Just a
```
# Error
也称为ExceptionMonad.
 * 运算类型: 计算中可能会fail并扔出异常的运算
 * 绑定策略: 跟Maybe类似, 计算中一点出现fail值, 则fail值穿过后续函数的边界返回
定义: 
```hs
  
```
```hs
do {  action1; action2; action3 } `catchError` handler
```
action函数可以用`throwError`来扔出异常, handler函数作为最后一个变换函数, 必须跟do-block类型一致.
```hs
instance MonadError (Either e) where
  throwError = Left
  (Left e) `catchError` handler = handler e
  a `catchError` _ -> a
```
# []
```hs
instance Monad [] where
  m >>= f  = concatMap f m
  return x = [x]
  fail   s = []
instance MonadPlus [] where
  mzero = []
  mplus = (++)
```
```hs
import Data.Char
import Control.Monad
data Parsed = Digit Integer | Hex Integer | Word String deriving Show
parseWord :: Parsed -> Char -> [Parsed]
parseWord (Word s) c
  | isAlpha c = return (Word (s ++ [c]))
  | otherwise = mzero
parseWord _ _ = mzero
parseDigit:: Parsed -> Char -> [Parsed]
parseDigit (Digit n) c
  | isDigit c = return (Digit ((n*10) + (toInteger (digitToInt c))))
  | otherwise = mzero
parseDigit _ _ = mzero
parse p c = (parseDigit p c) `mplus` (parseWord p c)
parseArg :: String -> [Parsed]
parseArg s = do init <- ((return (Digit 0)) `mplus` (return (Word "")))
                foldM parse init s
```
# IO
 * IO monad的定义是平台相关的
 * IO monad不提供任何构造函数, 任何从IO monad中删除数据的方法, 也就是one-way monad
```hs
import System.Environment
import System.IO
import Control.Monad.Error
translateString []   _  s = s
translateString _    [] s = s
translateString from to s = [ tr c | c <- s]
                            where dict = zip from to
                                  tr c = case (lookup c dict) of
                                         Nothing   -> c
                                         (Just c') -> c'
usage e = do
  putStrLn "oops!"
doTranslate :: IO ()
doTranslate = do
  [from, to] <- getArgs
  input <- getLine
  print $ translateString from to input
```
# State
# Reader
# Writer
# Cont
[[TOC]]
# Normal Form
An expression in normal form is fully evaluated, and no sub-expression could be evaluated any further (i.e. it contains no un-evaluated thunks).
当表达式已被充分求值，不可再化简，即为NormalForm.
These expressions are all in normal form:
```hs
42
(2, "hello")
 > (x + 1)
```
These expressions are not in normal form:
```hs
1 + 2                 -- we could evaluate this to 3
( > x + 1) 2       -- we could apply the function
"he" ++ "llo"         -- we could apply the (++)
(1 + 1, 2 + 2)        -- we could evaluate 1 + 1 and 2 + 2
```
# :sprint
可以使用sprint命令查看表达式的求值情况
```hs
> let x = 1 + 1
> :sp x
_
> let y = 2
> :sp y
2
```
 * `-` 表示未被求值
# thunk
 * Haskell语言被设计为惰性求值, 就是说表达式只在需要的时候才被求值
 * Haskell将未被求值的表达式保存在内存中, 这个就叫做thunk.
 * 一旦表达式需要求值, 实际上就是对thunk进行求值, 比如, 当我们要打印一个表达式的时候, 就会触发求值
 * 一旦thunk被求值, thunk将被求出的值代替
```hs
> x = 1 + 2
> :sp x
x = _ -- 还没有被求值, 内存中有对应的thunk
> x
3 
> :sp x
x = 3 -- 已经被求值了
```
```hs
> let x = 1 + 1
> let y = 1 + x
> let z = 1 + x
-- 对X求值
> x
2
> :sp x y z
x = 2 -- 求值
y = _ -- 未求值
z = _ -- 未求值
> let x = 1 + 1
> let y = 1 + x
> let z = 1 + x
-- 对y求值, 因为y依赖与x, 所以结果x和y都被求值
> y
3
> :sp x y z
x = 2
y = 3
z = _
```
# seq :: a -> b -> b
seq函数对a进行求值, 而后返回b, 利用它也可以对表达式立刻求值.
```hs
> let x = 1 + 1; y = x + 1
> seq y ()
()
> :sp x y
x = 2
y = 3
```
再来看一个稍微复杂一点的例子:
```hs
> let x = 1 + 1 ; xs = [x, x]
> :t x xs
x  = _
xs = [_,_]
> let t = tail xs
> seq t ()
()
> sp: t
[_] -- seq对t进行了求值, 但是结果并不是一个值
```
# WHNF :  Weak Head Normal Form
历史原因Haskell把没有被充分求值的表达叫做WHNF,  完全被求值的表达式交NormalForm
```hs
> let x = 1 + 1
> let xs = [x,x,x,x]
> :sp xs
[_,_,_,_] -- WHNF
-- 计算一下xs的长度
> length xs
4
-- xs中的表达式并没有被求值
> :sp xs
[_,_,_,_] -- WHNF
> sum xs
8
> :sp xs
[2,2,2,2] -- NF
```
```div class=warn
通常你不需要在乎到底是WHNF还是NF, 除非用Haskell做并行计算, 那个时候你必须十分清楚运算对象是不是已经是NF, 以便可以开始并行计算.
```
[[TOC]]
# Evaluation Strategies
Evaluation Strategies, or simply Strategies, are a means for modularizing parallel code by separating the algorithm from the parallelism. Sometimes they require you to rewrite your algorithm, but once you do so, you will be able to parallelize it in different ways just by substituting a new Strategy.
Concretely, a Strategy is a function in the Eval monad that takes a value of type a and returns the same value:
```
type Strategy a = a -> Eval a
```
The idea is that a Strategy takes a data structure as input, traverses the structure creating parallelism with rpar and rseq, and then returns the original value.
Here’s a simple example: Let’s create a Strategy for pairs that evaluates the two components of the pair in parallel. We want a function parPair with the following type
```
parPair :: Strategy (a,b)
```
From the definition of the Strategy type previously shown, we know that this type is the same as (a,b) -> Eval (a,b). So parPair is a function that takes a pair, does some computation in the Eval monad, and returns the pair again. Here is its definition:
strat.hs
```
parPair :: Strategy (a,b)
parPair (a,b) = do
  a' <- rpar a
  b' <- rpar b
  return (a',b')
```
This is similar to the rpar/rpar pattern that we saw in “The Eval Monad, rpar, and rseq”. The difference is that we’ve packaged it up as a Strategy: It takes a data structure (in this case a pair), creates some parallelism using rpar, and then returns the same data structure.
We’ll see this in action in a moment, but first we need to know how to use a Strategy. Using a Strategy consists of applying it to its input and running the Eval computation to get the output. We could write that directly with runEval; for example, to evaluate the pair (fib 35, fib 36) in parallel, we could write:
```
 runEval (parPair (fib 35, fib 36))
```
This works just fine, but it turns out to be much nicer to package up the application of a Strategy into a function named using:
```
using :: a -> Strategy a -> a
x `using` s = runEval (s x)
```
The using function takes a value of type a and a Strategy for a, and applies the Strategy to the value. We normally write using infix, as its definition suggests. Here is the parPair example above rewritten with using:
```
   (fib 35, fib 36) `using` parPair
```
Why write it this way? Well, a Strategy returns the same value that it was passed, so we know that aside from its performance, the above code is equivalent to just:
```
   (fib 35, fib 36)
```
So we’ve clearly separated the code that describes what the program does (the pair) from the code that adds the parallelism (`using` parPair). Indeed, everywhere we see x `using` s in our program, we can delete the `using` s part and the program should produce the same result.[6] Conversely, someone who is interested in parallelizing the program can focus on modifying the Strategy without worrying about breaking the program.
The example program strat.hs contains the parPair example just shown; try running it yourself with one and two processors to see it compute the two calls to fib in parallel.
# Parameterized Strategies
The parPair Strategy embodies a fixed policy: It always evaluates the components of the pair in parallel, and always to weak head normal form. If we wanted to do something different with a pair—fully evaluate the components to normal form, for example—we would have to write a completely new Strategy. A better way to factor things is to write a parameterized Strategy, which takes as arguments the Strategies to apply to the components of the data structure. Here is a parameterized Strategy for pairs:
strat.hs
```hs
evalPair :: Strategy a -> Strategy b -> Strategy (a,b)
evalPair sa sb (a,b) = do
  a' <- sa a
  b' <- sb b
  return (a',b')
```
This Strategy no longer has parallelism built in, so I’ve called it evalPair instead of parPair.[7] It takes two Strategy arguments, sa and sb, applies them to the respective components of the pair, and then returns the pair.
Compared with parPair, we are passing in the functions to apply to a and b instead of making fixed calls to rpar. So to define parPair in terms of evalPair, we can just pass rpar as the arguments:
```hs
parPair :: Strategy (a,b)
parPair = evalPair rpar rpar
```
This means we’re using rpar itself as a Strategy:
```hs
rpar :: Strategy a
```
The type of rpar is a -> Eval a, which is equivalent to Strategy a; rpar is therefore a Strategy for any type, with the effect of starting the evaluation of its argument while the enclosing Eval computation proceeds in parallel. (The rseq operation is also a Strategy.)
But parPair is still restrictive, in that the components of the pair are always evaluated to weak head normal form. What if we wanted to fully evaluate the components using force, for example? We can make a Strategy that fully evaluates its argument:
```!hs
rdeepseq :: NFData a => Strategy a
rdeepseq x = rseq (force x)
```
But how do we combine rpar with rdeepseq to give us a single Strategy that fully evaluates its argument in parallel? We need one further combinator, which is provided by Control.Parallel.Strategies:
```hs
rparWith :: Strategy a -> Strategy a
```
Think of rparWith s as wrapping the Strategy s in an rpar.
Now we can provide a parameterized version of parPair that takes the Strategies to apply to the components:
```hs
parPair :: Strategy a -> Strategy b -> Strategy (a,b)
parPair sa sb = evalPair (rparWith sa) (rparWith sb)
```
And we can use parPair to write a Strategy that fully evaluates both components of a pair in parallel:
```hs
  parPair rdeepseq rdeepseq :: (NFData a, NFData b) => Strategy (a,b)
```
To break down what happens when this Strategy is applied to a pair: parPair calls evalPair, and evalPair calls rparWith rdeepseq on each component of the pair. So the effect is that each component will be fully evaluated to normal form in parallel.
When using these parameterized Strategies, we sometimes need a way to say, "Don’t evaluate this component at all." The Strategy that does no evaluation is called r0:
```
r0 :: Strategy a
r0 x = return x
```
For example, we can write a Strategy over a pair of pairs that evaluates the first component (only) of both pairs in parallel.
```
  evalPair (evalPair rpar r0) (evalPair rpar r0) :: Strategy ((a,b),(c,d))
```
The first rpar applies to a and the first r0 to b, while the second rpar applies to c and the second r0 to d.
# A Strategy for Evaluating a List in Parallel
In Chapter 2, we defined a function parMap that would map a function over a list in parallel. We can think of parMap as a composition of two parts:
The algorithm: map
The parallelism: evaluating the elements of a list in parallel
And indeed, with Strategies, we can express it exactly this way:
```hs
parMap :: (a -> b) -> [a] -> [b]
parMap f xs = map f xs `using` parList rseq
```
The parList function is a Strategy on lists that evaluates the list elements in parallel. To define parList, we can take the same approach that we took with pairs earlier and first define a parameterized Strategy on lists, called evalList:
parlist.hs
```hs
evalList :: Strategy a -> Strategy [a]
evalList strat []     = return []
evalList strat (x:xs) = do
  x'  <- strat x
  xs' <- evalList strat xs
  return (x':xs')
```
Note that evalList walks the list recursively, applying the Strategy parameter strat to each of the elements and building the result list. Now we can define parList in terms of evalList, using rparWith:
parList :: Strategy a -> Strategy [a]
parList strat = evalList (rparWith strat)
In fact, both evalList and parList are already provided by Control.Parallel.Strategies so you don’t have to define them yourself, but it’s useful to see that their implementations are not mysterious.
As with parPair, the parList function is a parameterized Strategy. That is, it takes as an argument a Strategy on values of type a and returns a Strategy for lists of a. So parList describes a family of Strategies on lists that evaluate the list elements in parallel.
The parList Strategy covers a wide range of uses for parallelism in typical Haskell programs; in many cases, a single parList is all that is needed to expose plenty of parallelism.
Returning to our Sudoku solver from Chapter 2 for a moment: instead of our own hand-written parMap, we could have used parList:
sudoku5.hs
```hs
  let solutions = map solve puzzles `using` parList rseq
```
Using rseq as the Strategy for the list elements is enough here: The result of solve is a Maybe, so evaluating it to weak head normal form forces the solver to determine whether the puzzle has a solution.
This version has essentially the same performance as the version that used parMap in Chapter 2.
# Example: The K-Means Problem
Let’s look at a slightly more involved example. In the K-Means problem, the goal is to partition a set of data points into clusters. Figure 3-1 shows an example data set, and the circles indicate the locations of the clusters that the algorithm should derive. From the locations of the clusters, partitioning the points is achieved by simply finding the closest cluster to each point.
[[Image(http://orm-chimera-prod.s3.amazonaws.com/1230000000929/images/pcph_0301.png)]]
Finding an optimal solution to the problem is too expensive to be practical. However, there are several heuristic techniques that are fast, and even though they don’t guarantee an optimal solution, in practice, they give good results. The most well-known heuristic technique for K-Means is Lloyd’s algorithm, which finds a solution by iteratively improving an initial guess. The algorithm takes as a parameter the number of clusters to find and makes an initial guess at the center of each cluster. Then it proceeds as follows:
Assign each point to the cluster to which it is closest. This yields a new set of clusters.
Find the centroid of each cluster (the average of all the points in the cluster).
Repeat steps 1 and 2 until the cluster locations stabilize. We cut off processing after an arbitrarily chosen number of iterations, because sometimes the algorithm does not converge.
The initial guess can be constructed by randomly assigning each point in the data set to a cluster and then finding the centroids of those clusters.
The algorithm works in any number of dimensions, but we will use two for ease of visualization.
A complete Haskell implementation can be found in the directory kmeans in the sample code.
A data point is represented by the type Point, which is just a pair of Doubles representing the x and y coordinates respectively:[8]
```
data Point = Point !Double !Double
There are a couple of basic operations on Point:
```
kmeans/KMeansCore.hs
```
zeroPoint :: Point
zeroPoint = Point 0 0
sqDistance :: Point -> Point -> Double
sqDistance (Point x1 y1) (Point x2 y2) = ((x1-x2)^2) + ((y1-y2)^2)
```
We can make a zero point with zeroPoint, and find the square of the distance between two points with sqDistance. The actual distance between the points would be given by the square root of this value, but since we will only be comparing distances, we can save time by comparing squared distances instead.
Clusters are represented by the type Cluster:
```
data Cluster
  = Cluster { clId    :: Int
            , clCent  :: Point
            }
```
A Cluster contains its number (clId) and its centroid (clCent).
We will also need an intermediate type called PointSum:
data PointSum = PointSum !Int !Double !Double
A PointSum represents the sum of a set of points; it contains the number of points in the set and the sum of their x and y coordinates respectively. A PointSum is constructed incrementally, by repeatedly adding points using addToPointSum:
kmeans/kmeans.hs
```
addToPointSum :: PointSum -> Point -> PointSum
addToPointSum (PointSum count xs ys) (Point x y)
  = PointSum (count+1) (xs + x) (ys + y)
```
A PointSum can be turned into a Cluster by computing the centroid. The x coordinate of the centroid is the sum of the x coordinates of the points in the cluster divided by the total number of points, and similarly for the y coordinate.
```
pointSumToCluster :: Int -> PointSum -> Cluster
pointSumToCluster i (PointSum count xs ys) =
  Cluster { clId    = i
          , clCent  = Point (xs / fromIntegral count) (ys / fromIntegral count)
          }
```
The roles of the types Point, PointSum, and Cluster in the algorithm are as follows. The input is a set of points represented as [Point], and an initial guess represented as [Cluster]. The algorithm will iteratively refine the clusters until convergence is reached.
Step 1 divides the points into new sets by finding the Cluster to which each Point is closest. However, instead of collecting sets of Points, we build up a PointSum for each cluster. This is an optimization that avoids constructing the intermediate data structure and allows the algorithm to run in constant space. We’ll represent the output of this step as Vector PointSum.
The Vector PointSum is fed into step 2, which makes a Cluster from each PointSum, giving [Cluster].
The result of step 2 is fed back into step 1 until convergence is reached.
The function assign implements step 1 of the algorithm, assigning points to clusters and building a vector of PointSums:
```
assign :: Int -> [Cluster] -> [Point] -> Vector PointSum
assign nclusters clusters points = Vector.create $ do
    vec <- MVector.replicate nclusters (PointSum 0 0 0)
    let
        addpoint p = do
          let c = nearest p; cid = clId c
          ps <- MVector.read vec cid
          MVector.write vec cid $! addToPointSum ps p
    mapM_ addpoint points
    return vec
 where
  nearest p = fst $ minimumBy (compare `on` snd)
                        [ (c, sqDistance (clCent c) p) | c <- clusters ]
```
Given a set of clusters and a set of points, the job of assign is to decide, for each point, which cluster is closest. For each cluster, we build up a PointSum of the points that were found to be closest to it. The code has been carefully optimized, using mutable vectors from the vector package; the details aren’t important here.
The function makeNewClusters implements step 2 of the algorithm:
```
makeNewClusters :: Vector PointSum -> [Cluster]
makeNewClusters vec =
  [ pointSumToCluster i ps
  | (i,ps@(PointSum count _ _)) <- zip [0..] (Vector.toList vec)
  , count > 0
  ]
```
Here we make a new Cluster, using pointSumToCluster, from each PointSum produced by assign. There is a slight complication in that we have to avoid creating a cluster with no points, because it cannot have a centroid.
Finally step combines assign and makeNewClusters to implement one complete iteration:
```
step :: Int -> [Cluster] -> [Point] -> [Cluster]
step nclusters clusters points
   = makeNewClusters (assign nclusters clusters points)
To complete the algorithm, we need a loop to repeatedly apply the step function until convergence. The function kmeans_seq implements this:
kmeans_seq :: Int -> [Point] -> [Cluster] -> IO [Cluster]
kmeans_seq nclusters points clusters =
  let
      loop :: Int -> [Cluster] -> IO [Cluster]
      loop n clusters | n > tooMany = do                  -- 
        putStrLn "giving up."
        return clusters
      loop n clusters = do
        printf "iteration %d
" n
        putStr (unlines (map show clusters))
        let clusters' = step nclusters clusters points    -- 
        if clusters' == clusters                          -- 
           then return clusters
           else loop (n+1) clusters'
  in
  loop 0 clusters
tooMany = 80
```
The first argument to loop is the number of iterations completed so far. If this figure reaches the limit tooMany, then we bail out (sometimes the algorithm does not converge).
After printing the iteration number and the current clusters for diagnostic purposes, we calculate the next iteration by calling the function step. The arguments to step are the number of clusters, the current set of clusters, and the set of points.
If this iteration did not change the clusters, then the algorithm has converged, and we return the result. Otherwise, we do another iteration.
We compile this program in the same way as before:
```
$ cd kmeans
$ ghc -O2 -threaded -rtsopts -eventlog kmeans.hs
```
The sample code comes with a program to generate some input data, GenSamples.hs, which uses the normaldistribution package to generate a realistically clustered set of values. The data set is large, so it isn’t included with the sample code, but you can generate it using GenSamples:
```
$ ghc -O2 GenSamples.hs
$ ./GenSamples 5 50000 100000 1010
```
This should generate a data set of about 340,000 points with 5 clusters in the file points.bin.
Run the kmeans program using the sequential algorithm:
```
$ ./kmeans seq
```
The program will display the clusters at each iteration and should converge after 65 iterations.
Note that the program displays its own running time at the end; this is because there is a significant amount of time spent reading in the sample data at the beginning, and we want to be able to calculate the parallel speedup for the portion of the runtime spent computing the K-Means algorithm only.
Parallelizing K-Means
How can this algorithm be parallelized? One place that looks profitable to parallelize is the assign function because it is essentially just a map over the points, and indeed that is where we will concentrate our efforts. The operations are too fine-grained here to use a simple parMap or parList as we did before; the overhead of the parMap will swamp the parallelism, so we need to increase the size of the operations. One way to do that is to divide the list of points into chunks, and process the chunks in parallel. First we need some code to split a list into chunks:
```
split :: Int -> [a] -> [[a]]
split numChunks xs = chunk (length xs `quot` numChunks) xs
chunk :: Int -> [a] -> [[a]]
chunk n [] = []
chunk n xs = as : chunk n bs
  where (as,bs) = splitAt n xs
```
So we can split the list of points into chunks and map assign over the list of chunks. But what do we do with the results? We have a list of Vector PointSums that we need to combine into a single Vector PointSum. Fortunately, PointSums can be added together:
```
addPointSums :: PointSum -> PointSum -> PointSum
addPointSums (PointSum c1 x1 y1) (PointSum c2 x2 y2)
  = PointSum (c1+c2) (x1+x2) (y1+y2)
And using this, we can combine vectors of PointSums:
combine :: Vector PointSum -> Vector PointSum -> Vector PointSum
combine = Vector.zipWith addPointSums
We now have all the pieces to define a parallel version of step:
parSteps_strat :: Int -> [Cluster] -> [[Point]] -> [Cluster]
parSteps_strat nclusters clusters pointss
  = makeNewClusters $
      foldr1 combine $
          (map (assign nclusters clusters) pointss
            `using` parList rseq)
```
The arguments to parSteps_strat are the same as for step, except that the list of points is now a list of lists of points, that is, the list of points divided into chunks by split. We want to pass in the chunked data rather than call split inside parSteps_strat so that we can do the chunking of the input data just once instead of repeating it for each iteration.
The kmeans_strat function below is our parallel version of kmeans_seq, the only differences being that we call split to divide the list of points into chunks () and we call parSteps_strat instead of steps ():
```
kmeans_strat :: Int -> Int -> [Point] -> [Cluster] -> IO [Cluster]
kmeans_strat numChunks nclusters points clusters =
  let
      chunks = split numChunks points                            -- 
      loop :: Int -> [Cluster] -> IO [Cluster]
      loop n clusters | n > tooMany = do
        printf "giving up."
        return clusters
      loop n clusters = do
        printf "iteration %d
" n
        putStr (unlines (map show clusters))
        let clusters' = parSteps_strat nclusters clusters chunks -- 
        if clusters' == clusters
           then return clusters
           else loop (n+1) clusters'
  in
  loop 0 clusters
```
Note that the number of chunks doesn’t have to be related to the number of processors; as we saw earlier, it is better to produce plenty of sparks and let the runtime schedule them automatically, because this should enable the program to scale over a wide range of processors.
Performance and Analysis
Next we’re going on an exploration of the performance of this parallel program. Along the way, we’ll learn several lessons about the kinds of things that can go wrong when parallelizing Haskell code, how to look out for them using ThreadScope, and how to fix them.
We’ll start by taking some measurements of the speedup for various numbers of cores. When running the program in parallel, we get to choose the number of chunks to divide the input into, and for these measurements I’ll use 64 (but we’ll revisit this in “Granularity”). The program is run in parallel like this:
```
$ ./kmeans strat 64 +RTS -N2
```
strat indicates that we want to use the Strategies version of the algorithm, and 64 is the number of chunks to divide the input data into. Here, I’m telling the GHC runtime to use two cores.
Here are the speedup results I get on my computer for the kmeans program I showed earlier.[9] For each measurement, I ran the program a few times and took the average runtime.[10]
We can see that speedup is quite good for two to three cores but starts to drop off at four cores. Still, a 2.6 speedup on 4 cores is reasonably respectable.
The ThreadScope profile gives us some clues about why the speedup might be less than we hope. The overall view of the four-core run can be seen in Figure 3-2.
[[Image(http://orm-chimera-prod.s3.amazonaws.com/1230000000929/images/pcph_0302.png)]]
We can clearly see the sequential section at the start, where the program reads in the input data. But that isn’t a problem; remember that the program emits its own timing results, which begin at the parallel part of the run. The parallel section itself looks quite good; all cores seem to be running for the duration. Let’s zoom in on the beginning of the parallel section, as shown in Figure 3-3.
[[Image(http://orm-chimera-prod.s3.amazonaws.com/1230000000929/images/pcph_0303.png)]]
There’s a segment between 0.78s and 0.8s where, although parallel execution has started, there is heavy GC activity. This is similar to what we saw in “Example: Parallelizing a Sudoku Solver”, where the work of splitting the input data into lines was overlapped with the parallel execution. In the case of kmeans, the act of splitting the data set into chunks is causing the extra work.
The sequential version of the algorithm doesn’t need to split the data into chunks, so chunking is a source of extra overhead in the parallel version. This is one reason that we aren’t achieving full speedup. If you’re feeling adventurous, you might want to see whether you can avoid this chunking overhead by using Vector instead of a list to represent the data set, because Vectors can be sliced in O(1) time.
Let’s look at the rest of the parallel section in more detail (see Figure 3-4).
[[Image(http://orm-chimera-prod.s3.amazonaws.com/1230000000929/images/pcph_0304.png)]]
The parallel execution, which at first looked quite uniform, actually consists of a series of humps when we zoom in. Remember that the algorithm performs a series of iterations over the data set—these humps in the profile correspond to the iterations. Each iteration is a separate parallel segment, and between the iterations lies some sequential execution. We expect a small amount of sequential execution corresponding to makeNewClusters, combine, and the comparison between the new and old clusters in the outer loop.
Let’s see whether the reality matches our expectations by zooming in on one of the gaps to see more clearly what happens between iterations (Figure 3-5).
[[Image(http://orm-chimera-prod.s3.amazonaws.com/1230000000929/images/pcph_0305.png)]]
There’s quite a lot going on here. We can see the parallel execution of the previous iteration tailing off, as a couple of cores run longer than the others. Following this, there is some sequential execution on HEC 3 before the next iteration starts up in parallel.
Looking more closely at the sequential bit on HEC 3, we can see some gaps where nothing appears to be happening at all. In the ThreadScope GUI, we can show the detailed events emitted by the RTS (look for the "Raw Events" tab in the lower pane), and if we look at the events for this section, we see:
```
0.851404792s HEC 3: stopping thread 4 (making a foreign call)
0.851405771s HEC 3: running thread 4
0.851406373s HEC 3: stopping thread 4 (making a foreign call)
0.851419669s HEC 3: running thread 4
0.851451713s HEC 3: stopping thread 4 (making a foreign call)
0.851452171s HEC 3: running thread 4
...
```
The program is popping out to make several foreign calls during this period. ThreadScope doesn’t tell us any more than this, but it’s enough of a clue: A foreign call usually indicates some kind of I/O, which should remind us to look back at what happens between iterations in the kmeans_seq function:
```
 loop n clusters = do
        printf "iteration %d
" n
        putStr (unlines (map show clusters))
        ...
```
We’re printing some output. Furthermore, we’re doing this in the sequential part of the program, and Amdahl’s law is making us pay for it in parallel speedup.
Commenting out these two lines (in both kmeans_seq and kmeans_strat, to be fair) improves the parallel speedup from 2.6 to 3.4 on my quad-core machine. It’s amazing how easy it is to make a small mistake like this in parallel programming, but fortunately ThreadScope helps us identify the problem, or at least gives us clues about where we should look.
# Visualizing Spark Activity
We can also use ThreadScope to visualize the creation and use of sparks during the run of the program. Figure 3-6 shows the profile for kmeans running on four cores, showing the spark pool size over time for each HEC (these graphs are enabled in the ThreadScope GUI from the "Traces" tab in the left pane).
[[Image(http://orm-chimera-prod.s3.amazonaws.com/1230000000929/images/pcph_0306.png)]]
The figure clearly shows that as each iteration starts, 64 sparks are created on one HEC and then are gradually consumed. What is perhaps surprising is that the sparks aren’t always generated on the same HEC; this is the GHC runtime moving work behind the scenes as it tries to keep the load balanced across the cores.
There are more spark-related graphs available in ThreadScope, showing the rates of spark creation and conversion (running sparks). All of these can be valuable in understanding the performance characteristics of your parallel program.
# Granularity
Looking back at Figure 3-5, I remarked earlier that the parallel section didn’t finish evenly, with two cores running a bit longer than the others. Ideally, we would have all the cores running until the end to maximize our speedup.
As we saw in “Example: Parallelizing a Sudoku Solver”, having too few work items in our parallel program can impact the speedup, because the work items can vary in cost. To get a more even run, we want to create fine-grained work items and more of them.
To see the effect of this, I ran kmeans with various numbers of chunks from 4 up to 512, and measured the runtime on 4 cores. The results are shown in Figure 3-7.
[[Image(http://orm-chimera-prod.s3.amazonaws.com/1230000000929/images/pcph_0307.png)]]
We can see not only that having too few chunks is not good for the reasons given above, but also having too many can have a severe impact. In this case, the sweet spot is somewhere around 50-100.
Why does having too many chunks increase the runtime? There are two reasons:
There is some overhead per chunk in creating the spark and arranging to run it on another processor. As the chunks get smaller, this overhead becomes more significant.
The amount of sequential work that the program has to do is greater. Combining the results from 512 chunks takes longer than 64, and because this is in the sequential part, it significantly impacts the parallel performance.
# GC’d Sparks and Speculative Parallelism
Recall the definition of parList:
```
parList :: Strategy a -> Strategy [a]
parList strat = evalList (rparWith strat)
And the underlying parameterized Strategy on lists, evalList:
evalList :: Strategy a -> Strategy [a]
evalList strat []     = return []
evalList strat (x:xs) = do
  x'  <- strat x
  xs' <- evalList strat xs
  return (x':xs')
```
As evalList traverses the list applying the strategy strat to the list elements, it remembers each value returned by strat (bound to x'), and constructs a new list from these values. Why? Well, one answer is that a Strategy must return a data structure equal to the one it was passed.
But do we really need to build a new list? After all, this means that evalList is not tail-recursive; the recursive call to evalList is not the last operation in the do on its right-hand side, so evalList requires stack space linear in the length of the input list.
Couldn’t we just write a tail-recursive version of parList instead? Perhaps like this:
```
parList :: Strategy a -> Strategy [a]
parList strat xs = do
  go xs
  return xs
 where
  go []     = return ()
  go (x:xs) = do rparWith strat x
                 go xs
```
After all, this is type-correct and seems to call rparWith on each list element as required.
Unfortunately, this version of parList has a serious problem: All the parallelism it creates will be discarded by the garbage collector. The omission of the result list turns out to be crucial. Let’s take a look at the data structures that our original, correct implementations of parList and evalList created (Figure 3-8).
[[Image(http://orm-chimera-prod.s3.amazonaws.com/1230000000929/images/pcph_0308.png)]]
At the top of the diagram is the input list xs: a linked list of cells, each of which points to a list element (x1, x2, and so forth). At the bottom of the diagram is the spark pool, the runtime system data structure that stores references to sparks in the heap. The other structures in the diagram are built by parList (the correct version, not the one I most recently showed). Each strat box represents the strategy strat applied to an element of the original list, and xs' is the linked list of cells in the output list. The spark pool contains pointers to each of the strat boxes; these are the pointers created by each call to rparWith.
The GHC runtime regularly checks the spark pool for any entries that are not required by the program and removes them. It would be bad to retain entries that aren’t needed, because that could cause the program to hold on to memory unnecessarily, leading to a space leak. We don’t want parallelism to have a negative impact on performance.
How does the runtime know whether an entry is needed? The same way it knows whether any item in memory is needed: There must be a pointer to it from something else that is needed. This is the reason that parList creates a new list xs'. Suppose we did not build the new list xs', as in the tail-recursive version of parList above. Then the only reference to each strat box in the heap would be from the spark pool, and hence the runtime would automatically sweep all those references from the spark pool, discarding the parallelism. So we build a new list xs' to hold references to the strat calls that we need to retain.
The automatic discarding of unreferenced sparks has another benefit besides avoiding space leaks; suppose that under some circumstances the program does not need the entire list. If the program simply forgets the unused remainder of the list, the runtime system will clean up the unreferenced sparks from the spark pool and will not waste any further parallel processing resources on evaluating those sparks. The extra parallelism in this case is termed speculative, because it is not necessarily required, and the runtime will automatically discard speculative tasks that it can prove will never be required—a useful property!
Although the runtime system’s discarding of unreferenced sparks is certainly useful in some cases, it can be tricky to work with because there is no language-level support for catching mistakes. Fortunately, the runtime system will tell us if it garbage-collects unreferenced sparks. For example, if you use the tail-recursive parList with the Sudoku solver from Chapter 2, the +RTS -s stats will show something like this:
  SPARKS: 1000 (2 converted, 0 overflowed, 0 dud, 998 GC'd, 0 fizzled)
Garbage-collected sparks are reported as "GC’d." ThreadScope will also indicate GC’d sparks in its spark graphs.
If you see that a large number of sparks are GC’d, it’s a good indication that sparks are being removed from the spark pool before they can be used for parallelism. Unless you are using speculation, a non-zero figure for GC’d sparks is probably a bad sign.
All the combinators in the Control.Parallel.Strategies libraries retain references to sparks correctly. These are the rules of thumb for not shooting yourself in the foot:
 * Use using to apply Strategies instead of runEval; it encourages the right pattern, in which the program uses the results of applying the Strategy.
 * When writing your own Eval monad code, this is wrong:
```
do
    ...
    rpar (f x)
    ...
```
Equivalently, using rparWith without binding the result is wrong. However, this is OK:
```
  do
    ...
    y <- rpar (f x)
    ... y ...
```
And this might be OK, as long as y is required by the program somewhere:
```
  do
    ...
    rpar y
    ...
```
# Parallelizing Lazy Streams with parBuffer
A common pattern in Haskell programming is to use a lazy list as a stream so that the program can consume input while simultaneously producing output and consequently run in constant space. Such programs present something of a challenge for parallelism; if we aren’t careful, parallelizing the computation will destroy the lazy streaming property and the program will require space linear in the size of the input.
To demonstrate this, we will use the sample program rsa.hs, an implementation of RSA encryption and decryption. The program takes two command line arguments: the first specifies which action to take, encrypt or decrypt, and the second is either the filename of the file to read, or the character - to read from stdin. The output is always produced on stdout.
The following example uses the program to encrypt the message "Hello World!":
```
$ echo 'Hello World!' | ./rsa encrypt -
11656463941851871045300458781178110195032310900426966299882646602337646308966290
04616367852931838847898165226788260038683620100405280790394258940505884384435202
74975036125752600761230510342589852431747
```
And we can test that the program successfully decrypts the output, producing the original text, by piping the output back into rsa decrypt:
```
$ echo "Hello World!" | ./rsa encrypt - | ./rsa decrypt -
Hello World!
```
The rsa program is a stream transformer, consuming input and producing output lazily. We can see this by looking at the RTS stats:
```
$ ./rsa encrypt /usr/share/dict/words >/dev/null +RTS -s
   8,040,128,392 bytes allocated in the heap
      66,756,936 bytes copied during GC
         186,992 bytes maximum residency (71 sample(s))
          36,584 bytes maximum slop
               2 MB total memory in use (0 MB lost due to fragmentation)
```
The /usr/share/dict/words file is about 1 MB in size, but the program has a maximum residency (live memory) of 186,992 bytes.
Let’s try to parallelize the program. The program uses the lazy ByteString type from Data.ByteString.Lazy to achieve streaming, and the top-level encrypt function has this type:
```
encrypt :: Integer -> Integer -> ByteString -> ByteString
```
The two Integers are the key with which to encrypt the data. The implementation of encrypt is a beautiful pipeline composition:
rsa.hs
```
encrypt n e = B.unlines                                 --  1 
            . map (B.pack . show . power e n . code)   --  2
            . chunk (size n)                           --  3
```

 1. Divide the input into chunks. Each chunk is encrypted separately; this has nothing to do with  parallelism.
 2. Encrypt each chunk.
 3. Concatenate the result as a sequence of lines.
We won’t delve into the details of the RSA implementation here, but if you’re interested, go and look at the code in rsa.hs (it’s fairly short). For the purposes of parallelism, all we need to know is that there’s a map on the second line, so that’s our target for parallelization.
First, let’s try to use the parList Strategy that we have seen before:
rsa1.hs
```
encrypt n e = B.unlines
            . withStrategy (parList rdeepseq)        -- 
            . map (B.pack . show . power e n . code)
            . chunk (size n)
```
I’m using withStrategy here, which is just a version of using with the arguments flipped; it is slightly nicer in situations like this. The Strategy is parList, with rdeepseq as the Strategy to apply to the list elements (the list elements are lazy ByteStrings, so we want to ensure that they are fully evaluated).
If we run this program on four cores, the stats show something interesting:
```
   6,251,537,576 bytes allocated in the heap
      44,392,808 bytes copied during GC
       2,415,240 bytes maximum residency (33 sample(s))
         550,264 bytes maximum slop
              10 MB total memory in use (0 MB lost due to fragmentation)
```
The maximum residency has increased to 2.3 MB, because the parList Strategy forces the whole spine of the list, preventing the program from streaming in constant space. The speedup in this case was 2.2; not terrible, but not great either. We can do better.
The Control.Parallel.Strategies library provides a Strategy to solve exactly this problem, called parBuffer:
```
parBuffer :: Int -> Strategy a -> Strategy [a]
```
The parBuffer function has a similar type to parList but takes an Int argument as a buffer size. In contrast to parList which eagerly creates a spark for every list element, parBuffer N creates sparks for only the first N elements of the list, and then creates more sparks as the result list is consumed. The effect is that there will always be N sparks available until the end of the list is reached.
The disadvantage of parBuffer is that we have to choose a particular value for the buffer size, and as with the chunk factor we saw earlier, there will be a "best value" somewhere in the range. Fortunately, performance is usually not too sensitive to this value, and something in the range of 50-500 is often good. So let’s see how well this works:
rsa2.hs
```
encrypt n e = B.unlines
            . withStrategy (parBuffer 100 rdeepseq)             -- 
            . map (B.pack . show . power e n . code)
            . chunk (size n)
```
Here I replaced parList with parBuffer 100.
This programs achieves a speedup of 3.5 on 4 cores. Furthermore, it runs in much less memory than the parList version:
```
   6,275,891,072 bytes allocated in the heap
      27,749,720 bytes copied during GC
         294,872 bytes maximum residency (58 sample(s))
          62,456 bytes maximum slop
               4 MB total memory in use (0 MB lost due to fragmentation)
}}
We can expect it to need more memory than the sequential version, which required only 2 MB, because we’re performing many computations in parallel. Indeed, a higher residency is common in parallel programs for the simple reason that they are doing more work, although it’s not always the case; sometimes parallel evaluation can reduce memory overhead by evaluating thunks that were causing space leaks.
ThreadScope’s spark pool graph shows that parBuffer really does keep a constant supply of sparks, as shown in Figure 3-9.
[[Image(http://orm-chimera-prod.s3.amazonaws.com/1230000000929/images/pcph_0309.png)]]
The spark pool on HEC 0 constantly hovers around 90-100 sparks.
In programs with a multistage pipeline, interposing more calls to withStrategy in the pipeline can expose more parallelism.
# Chunking Strategies
When parallelizing K-Means in “Parallelizing K-Means”, we divided the input data into chunks to avoid creating parallelism with excessively fine granularity. Chunking is a common technique, so the Control.Parallel.Strategies library provides a version of parList that has chunking built in:
```
parListChunk :: Int -> Strategy a -> Strategy [a]
```
The first argument is the number of elements in each chunk; the list is split in the same way as the chunk function that we saw earlier in the kmeans example. You might find parListChunk useful if you have a list with too many elements to spark every one, or when the list elements are too cheap to warrant a spark each.
The spark pool has a fixed size, and when the pool is full, subsequent sparks are dropped and reported as overflowed in the +RTS -s stats output. If you see some overflowed sparks, it is probably a good idea to create fewer sparks; replacing parList with parListChunk is a good way to do that.
Note that chunking the list incurs some overhead, as we noticed in the earlier kmeans example when we used chunking directly. For that reason, in kmeans we created the chunked list once and shared it amongst all the iterations of the algorithm, rather than using parListChunk, which would chunk the list every time.
# The Identity Property
I mentioned at the beginning of this chapter that if we see an expression of this form:
```
  x `using` s
```
We can delete `using` s, leaving an equivalent program. For this to be true, the Strategy s must obey the identity property; that is, the value it returns must be equal to the value it was passed. The operations provided by the Control.Parallel.Strategies library all satisfy this property, but unfortunately it isn’t possible to enforce it for arbitrary user-defined Strategies. Hence we cannot guarantee that x `using` s == x, just as we cannot guarantee that all instances of Monad satisfy the monad laws, or that all instances of Eq are reflexive. These properties are satisfied by convention only; this is just something to be aware of.
There is one more caveat to this property. The expression x `using` s might be less defined than x, because it evaluates more structure of x than the context does. What does less defined mean? It means that the program containing x `using` s might fail with an error when simply x would not. A trivial example of this is:
```
print $ snd (1 `div` 0, "Hello!")
```
This program works and prints "Hello!", but:
```
print $ snd ((1 `div` 0, "Hello!") `using` rdeepseq)
```
This program fails with divide by zero. The original program didn’t fail because the erroneous expression was never evaluated, but adding the Strategy has caused the program to fully evaluate the pair, including the division by zero.
This is rarely a problem in practice; if the Strategy evaluates more than the program would have done anyway, the Strategy is probably wasting effort and needs to be modified.
----
[6] This comes with a couple of minor caveats that we’ll describe in “The Identity Property”.
[7] The evalPair function is provided by Control.Parallel.Strategies as evalTuple2.
[8] The actual implementation adds UNPACK pragmas for efficiency, which I have omitted here for clarity.
[9] A quad-core Intel i7-3770
[10] To do this scientifically, you would need to be much more rigorous, but the goal here is just to optimize our program, so rough measurements are fine.
[[TOC]]
# Parallelism
# The Eval Monad, rpar, and rseq
​```hs
data Eval a
instance Monad Eval
runEval :: Eval a -> a
rpar :: a -> Eval a
rseq :: a -> Eval a
```
 * rpar : 对参数a, 并行求值, 不等待结果返回
 * rseq : 对参数a同步求值, 等待结果返回
# rpar / rpar 模式
```hs
import Control.Parallel.Strategies
factor 1 = 1
factor n = n * (factor (n - 1))
rpar_rseq_Eval= runEval $ do
    a <- rpar (factor 10)
    b <- rseq (factor 15)
    return (a, b)
```
 * b计算完成, 函数rpar_rseq_Eval返回
# rpar / rpar / rseq
a/b完成后, 再返回
```hs
syncEval = runEval $ do
    a <- rpar (factor 10)
    b <- rpar (factor 15)
    rseq a
    rseq b
    return (a, b)
```
```hs
> rpar_rseq_Eval
(3628800,1307674368000)
```
# 编译运行
```sh
$ ghc -O2 rpar.hs  -threaded
$ ./rpar +RTS -N2 -s
```
 * +RTS -N2 : 使用双核运行程序
# force
rpar/rseq和seq函数求值行为类似, 只是将给定的表达式编程WHNF, 并不是充分求值. 这个有点蛋疼, WHNF不是我们最终想要的结果.
force就是用来弥补这个空缺的. 
```
force :: NFData a => a -> a
class NFData a where
  rnf :: a -> ()
  rnf a = a `seq` ()
force :: NFData a => a -> a
force x = x `deepseq` x
deepseq :: NFData a => a -> b -> b
deepseq a b = rnf a `seq` b
```
```div class=note
force/deepseq 对整个结构都要进行充分求值得到NF, 其代价很高, 所以要避免对同一个结构进行两次force/deepseq计算.
```
# 并行效率
如果你有两个核心, 把工作分成两份交给不同的核心来运算, 速度是不是会快上一倍呢?
答案并不一定, 等分数据并不等价于等分工作. 相同的数据量可能运算量不同. 可以使用ThreadScope观察到这种不均衡.  
所以, 要注意等量均分任务可能造成速度损失.
# Static Partitioning
A fixed division of work is often called static partitioning,
# Dynamic Partitioning 
distributing smaller units of work among processors at runtime is called dynamic partitioning.
GHC实现了Dynamic Partitioning , 所以为了让并行效率最高, 我们需要做的就是要把代码尽可能小粒度并行话. 简单说, 就是尽量用rpar拆分出更多的任务. 如果你使用了一次rpar, 那么就算有两个CPU也不能跑的更快, 因为只有一个不能分解的任务
想想一下, CPU就好像是炉子, 需要运算的对象就好比柴火, 你要做的就是不管有多少个炉子, 尽可能快的将柴火全部烧掉. 如果柴火就是一颗树干, 你有两个炉子, 想必你并不会把整个树干都塞到一个炉子里烧, 那样会导致另一个炉子无柴可烧. 所以你要把它砍成两段分别赛给不同的炉子, 但是问题又来了, 树干有粗细, 虽然你一斧头砍为两半, 但是发现粗的那个烧的慢. 所以咱们可以换个思路, 把整个树干分成四份, 十六份, 无等等份. 然后你边烧边观察, 哪里先烧完就往哪里加小柴.
说到这里, 回想一下rpar的参数, 它是个普通表达式, , 它有个很好文艺的名字: `spark`. rpar就是个斧头, 劈出来的spark保存在一个池子里, 等待有空闲的处理器就会被喂进去.
所以rpar的工作简单到不行, 就是把表达式的指针扔到一个队列中.
还记得运行程序时加了-s参数的输出结果么?
```
 SPARKS: 2 (1000 converted, 0 overflowed, 0 dud, 0 GC'd, 0 fizzled)
```
 * SPARKS  : sparks的个数
 * overflowed : spark pool固定长度, 满了之后再添加spark, 就会产生丢弃
 * dud : 当你把表达式丢给rpar时已经是NF了, 无需求值, 无需进入spark pool
 * GC'd : spark对程序来说没有用, 被runtime移除掉了
 * fizzled : spark pool中被间接求值变为NF的spark后被runtime移除
#  Amdahl’s law
```
1 / ((1 - P) + P/N)
```
 * P :  the portion of the runtime that can be parallelized
 * N : number of processors available.
# 并行编程的本质
回想一下砍柴烧火这件事情, 并行编程的本质其实与之神似. 如果没有并行编程, 我们写的Haskell程序就是一颗完整的树干, 为了能烧的快一点, 我们就要在合适的部位来上一斧头(rpar/rseq). 这样做并不是没有代价的, 本来一颗树干是非常完美的, 体现了一颗树的本质, 不多也不少, 但是被我们东一斧头西一斧头砍个稀烂,  虽然能烧的更快, 也变得更难理解.
```
   ... rpar
   ... rseq ... rseq
   ... rpar ... rpar
   ... rpar
```
有没有办法更优雅的方法来分解树干, 使之还可以保存原有的形态呢? 
有没有办法更优雅的来并行化程序, 使之还可以保留原有的结构呢?
# Evaluation Strategies : 求值策略
# 关键词
 * shallow evaluation
 * deep evaluation
 * WHNF
 * NF
# 参考
 * http://chimera.labs.oreilly.com/books/1230000000929/ch02.html#sec_par-eval-whnf
[[TOC]]
# Pattern Matching
# 函数的模式匹配
```hs
factorial :: Int -> Int
factorial 0 = 1
factorial n = n * factorial (n - 1)
```
# Tuples的模式匹配
```hs
ghci> let people=("amas", 16)
ghci> let name (name, _) = name
ghci> name people
"amas"
```
# As-patterns: <body>@(<pattern>)
```hs
firstLetter :: String -> String
firstLetter "" = "Empty string, whoops!"
firstLetter all@(x:xs) = "The first letter of " ++ all ++ " is " ++ [x]
```
In mainstream object oriented languages, subtype polymorphism is more widespread than parametric polymorphism. The subclassing mechanisms of C++ and Java give them subtype polymorphism. A base class defines a set of behaviours that its subclasses can modify and extend. Since Haskell isn't an object oriented language, it doesn't provide subtype polymorphism. 5 comments
Also common is coercion polymorphism, which allows a value of one type to be implicitly converted into a value of another type. Many languages provide some form of coercion polymorphism: one example is automatic conversion between integers and floating point numbers. Haskell deliberately avoids even this kind of simple automatic coercion.
# 参考
 * [Wadler89] Philip Wadler.  Theorems for free
[[TOC]]
# Haskell中的随机数生成器
在指令型语言中,随机函数生成器的API通常是一个函数或者方法, 这个函数返回
# System.Random
# 获得随机序列(RandomSequence)
# randomRs
```hs
> import System.Random
> g <- getStdGen
> take 10 (randomRs ('a', 'z') g)
"jfzhindezo"
> take 10 (randomRs ('a', 'z') g)
"jfzhindezo"
> take 10 (randomRs (1,100) g)
[6,4,56,24,55,74,56,75,48,77]
> take 10 (randomRs (1,100) g)
[6,4,56,24,55,74,56,75,48,77]
```
r.hs:
```
import System.Random
main = do
  g <- getStdGen
  print . take 10 $ (randomRs ('a', 'z') g)
  print . take 10 $ (randomRs ('a', 'z') g)
```
```hs
> :l r.hs
> main
"fkujeigwqg"
"fkujeigwqg"
> main
"fkujeigwqg"
"fkujeigwqg"
```
你应该已经发现,再获得随机数生成器之后,从其中抓出的元素都是固定的随机序列, 这是因为它是在Runtime启动的时候初始化的. 所以如果你ghc编译成可执行文件, 每次执行就会得到不同的结果. 
可以使用newStdGen代替getStdGen, 每次获得一个新的随机生成器.
r.hs:
```hs
import System.Random
main = do
  g <- newStdGen -- 注意这一行, 每次求值就会获得一个新的随机生成器
  print . take 10 $ (randomRs ('a', 'z') g)
  print . take 10 $ (randomRs ('a', 'z') g)
```
```hs
> :l r.hs
> main
"fybsiemkvu"
"fybsiemkvu"
> main
"tjwpwrowdl"
"tjwpwrowdl"
```
randomRs和下面提到的randoms都是纯函数, 这就意味着, 给定同一个随机数生成器, 它们总会得到相同的随机序列.
# randoms
如果你不需要设定随机元素的取值范围,可以使用randoms
```hs
> import System.Random
> take 3 (randoms g :: [Int])
[1021315375,-49446355,-1066240494]
> take 3 (randoms g :: [Bool])
[True,True,True]
```
# 抛硬币
从上面的例子中我们了解了如何使用随机生成器,以及相关的随机序列生成函数. 接下来, 我们模拟一下抛硬币.
某个类型可以由randoms/randomRs函数生成随机序列, 必须实现[Haskell/TypeClass/Random]中定义的两个方法:
 * randomR :: `RandomGen g => (a, a) -> g -> (a, g))`
 * random (通常情况下: random g = randomR (minBound, maxBound) g)

由于random中使用了类型的上界和下界, 所以新类型也必须是[Haskell/TypeClass/Bound]的实例.
draw_coin.hs:
```hs
import System.Random
data Coin = Heads | Tails deriving (Show, Enum, Bounded)
instance Random Coin where
  randomR (a, b) g =
    case randomR (fromEnum a, fromEnum b) g of
      (x, g') -> (toEnum x, g')
  random g = randomR (minBound, maxBound) g
drawCoin times = do
  g <- newStdGen
  print . take times $ (randoms g :: [Coin])
```
```hs
> :l draw_coin.hs
> drawCoin 1
[Tails]
> drawCoin 5
[Tails,Heads,Heads,Tails,Heads]
> drawCoin 5
[Heads,Tails,Heads,Heads,Tails]
> drawCoin 10
[Tails,Heads,Tails,Heads,Tails,Heads,Heads,Tails,Heads,Tails]
```
```hs
:i Random
class Random a where
  randomR   :: RandomGen g => (a, a) -> g -> (a, g)
  random    :: RandomGen g => g -> (a, g)
  randomRs  :: RandomGen g => (a, a) -> g -> [a]
  randoms   :: RandomGen g => g -> [a]
  randomRIO :: (a, a) -> IO a
  randomIO  :: IO a
```
# 参考
 * https://www.fpcomplete.com/school/starting-with-haskell/libraries-and-frameworks/randoms
 * [http://zh.wikipedia.org/wiki/%E8%92%99%E5%9C%B0%E5%8D%A1%E7%BE%85%E6%96%B9%E6%B3%95 MonteCarloMethod]
[[TOC]]
# Stack Overflow
 1. 在Haskell中没有调用栈(CallStack)
 2. 但是有一个模式匹配栈(PatternMatchingStack), 本质上是一些列的case表达式，等待求值过程中能够匹配预设的某个构造函数(WHNF)
There is no call stack in Haskell. Instead we find a pattern matching stack whose entries are essentially case expressions waiting for their scrutinee to be evaluated enough that they can match a constructor ([wiki:WHNF WeakHeadNormalForm]).
When GHC is evaluating a thunked expression it uses an internal stack. This inner stack for thunk evaluation is the one that can overflow in practice.
当GHC对表达式求值的时候会使用内部栈， 这个内部栈在实践过程中是会溢出的。
 * 如果你不是以[TailRecursive]的方式实现了某个递归函数，那么这个函数极有可能发生StackOverflow.
来看一个简单的例子，比如求列表的长度:
```hs
length' :: [a] -> Int
length' [] = 0
length' (x:xs) = len xs + 1
```
来测试一下极端的情况, 内存用尽。
```sh
> length' [1..100000000]
<interactive>: out of memory (requested 1048576 bytes)length' [1..100000000]
```
`length'` 的LastCall是`+`，而非其本身，所以会不断消耗模式匹配的栈空间. 
```hs

```
实际上这些情况根本不会发生。。。。
# 参考
 * http://www.haskell.org/haskellwiki/Stack_overflow
[[TOC]]
# Tail Recursion
# 参考
 * http://www.haskell.org/haskellwiki/Tail_recursion
 * http://www.haskell.org/pipermail/haskell-cafe/2009-March/058607.html
 * http://stackoverflow.com/questions/13042353/does-haskell-have-tail-recursive-optimization
[[TOC]]
# Tail Recursivity Optimization
A thunk is a value that is yet to be evaluated. It is used in Haskell systems, that implement non-strict semantics by lazy evaluation. A lazy run-time system does not evaluate a thunk unless it has to. Expressions are translated into a graph (not a tree, as you might have expected, this enables sharing, and infinite lists!) and a Spineless Tagless G-machine (STG, G for graph, I suppose?) reduces it, chucking out any unneeded thunk, unevaluated.
1 Why are thunks useful?
Well, if you don't need it, why evaluate it? Take for example the "lazy" && (and) operation. Two boolean expressions joined by && together is true if and only if both of them are. If you find out that one of them is false, you immediately know the joined expression cannot be true.
-- the first is false, so is the answer, don't even need to know what the other is
False && _ = False
-- so the first turns out to be true, hmm...
-- if the second is true, then the result is true
-- if it's false, so is the result
-- in other words, the result is the second!
True  && x = x
This function only evaluates the first parameter, because that's all that is needed. Even if the first parameter is true, you don't need to evaluate the second, so don't (so this version effectively is smarter than the explicit truth table). Who knows, the second parameter may get thrown out later as well!
Perhaps a more convincing example is a (naive but intuitive) algorithm to find out if a given number is prime.
-- the non-trivial factors are those who divide the number so no remainder
factors n = filter (m -> n `mod` m == 0) [2 .. (n - 1)]
-- a number is a prime if it has no non-trivial factors
isPrime n = n > 1 && null (factors n)
Fascinatingly, isPrime evaluates to False as soon as it finds a factor (due to the lazy definition of null), and discards the rest of the list.
2 When are thunks not so useful?
If you keep building up a very complicated graph to reduce later, it consumes memory (naturally), and can hinder performance, like, (from a blog comment made by Cale Gibbard)
foldl (+) 0 [1,2,3]
==> foldl (+) (0 + 1) [2,3]
==> foldl (+) ((0 + 1) + 2) [3]
==> foldl (+) (((0 + 1) + 2) + 3) []
==> ((0 + 1) + 2) + 3
==> (1 + 2) + 3
==> 3 + 3
==> 6
and by ==>, "is reduced to" is meant. Lucky, the example is not applied to [1 .. 2^40] because it would have taken a very long time to load this page then, and so would your program when run. For more involved (and subtle!) examples, see Performance/Strictness.
# Haskell Tuples
Tuples是一种存储多值的结构, 它允许你向其中存储多个值, 而无须考虑值的类型, 来看几个例子:
```
#!hs
Prelude> ("age", 18)
("age",18)
Prelude> ("x", True, 9)
("x",True,9)
```
 * Tuples中的元素个数我们称之为Tuples的大小
 * 两个元素的Tuples也叫Pair
 * 两个以上元素的Tuples亦可称为n-tuple
# 使用 fst 和 snd 取得2-tuple中的元素
```
#!hs
Prelude> fst (1,2)
1
Prelude> snd (1,2)
2
-- fst只能对2-tuple使用哈!
Prelude> fst (1,2,3)
<interactive>:1:5:
    Couldn't match expected type `(a0, b0)'
                with actual type `(t0, t1, t2)'
    In the first argument of `fst', namely `(1, 2, 3)'
    In the expression: fst (1, 2, 3)
    In an equation for `it': it = fst (1, 2, 3)
```
[[TOC]]
# Enum
```hs
class Enum a where
  succ :: a -> a
  pred :: a -> a
  toEnum :: Int -> a
  fromEnum :: a -> Int
  enumFrom :: a -> [a]
  enumFromThen :: a -> a -> [a]
  enumFromTo :: a -> a -> [a]
  enumFromThenTo :: a -> a -> a -> [a]
```
# Eq
```
class Eq a where
  (==) :: a -> a -> Bool
  (/=) :: a -> a -> Bool
        -- Defined in `GHC.Classes'
```
[[TOC]]
# Functor
Now it's time for my favorite type class!
Recall the map function on lists:
map :: (a -> b) -> [a] -> [b]
map takes a function and applies it to every element of a list, creating a new list with the results. We saw on Monday that the same pattern can be used for Trees:
treeMap :: (a -> b) -> Tree a -> Tree b
If you think a little, you'll realize that map makes sense for pretty much any data structure that holds a single type of values. It would be nice if we could factor this out into a class to keep track of the types that support map.
Behold, Functor:
class Functor f where
  fmap :: (a -> b) -> f a -> f b
Functor is a little different than the other classes we've seen so far. It's a "constructor" class, because the types it works on are constructors like Tree, Maybe and [] - they take another type as their argument. Notice how the f in the class declaration is applied to other types.
The standard library defines:
instance Functor [] where
  -- fmap :: (a -> b) -> [a] -> [b]
  fmap = map
And we can define an instance for our trees:
instance Functor Tree where
  -- fmap :: (a -> b) -> Tree a -> Tree b
  fmap = treeMap
The standard library also defines Functor instances for a number of other types. For example, Maybe is a Functor:
instance Functor Maybe where
   fmap _ Nothing  = Nothing
   fmap f (Just a) = Just (f a)
Functor is very useful, and you'll see many more examples of it in the weeks to come.
# Maybe
# 参考
 * http://www.haskell.org/haskellwiki/Maybe
 * http://withouttheloop.com/articles/2013-05-19-maybe-haskell/
[[TOC]]
# Type Class
TypeClass是一组接口, 是具有某种行为所必须实现的函数集合.
# Ord
```hs
class Eq a => Ord a where
  compare :: a -> a -> Ordering
  (<)  :: a -> a -> Bool
  (>=) :: a -> a -> Bool
  (>)  :: a -> a -> Bool
  (<=) :: a -> a -> Bool
  max  :: a -> a -> a
  min  :: a -> a -> a
        -- Defined in `GHC.Classes'
```
# Read
```
class Read a where
  readsPrec :: Int -> ReadS a
  readList :: ReadS [a]
  GHC.Read.readPrec :: Text.ParserCombinators.ReadPrec.ReadPrec a
  GHC.Read.readListPrec ::
    Text.ParserCombinators.ReadPrec.ReadPrec [a]
        -- Defined in `GHC.Read'
```
# Show
```
class Show a where
  showsPrec :: Int -> a -> ShowS
  show :: a -> String
  showList :: [a] -> ShowS
```
# type
# 具体类型: Concrete Type
如果某个类型的类型构造器不接受任何类型参数, 则此种类型为具体类型.
比如: Int
反之: Maybe
[[TOC]]
# Data Type
# 什么是类型?
 * 类型(Type)是一组值的集合
Bool类型:
```
data Bool = False | True
```
Int类型:
```
data Short = -32768 | -32767 | ...  | 0 | 1 | ... | 32767
```
既然类型是一个集合, 如果恰好是一个有限集合, 就可以通过枚举的方式来定义这个类型. 比方说上面的Bool和Short这两种类型的定义方式就属于这种情况.
# Type Constructors
如果某个类型包含数也数不清的值, 那么它就是无限集合, 我们该如何定义它呢? 通过枚举的方法显然是不行的.
比如: 圆形是一个类型, 那么它就包含了所有的圆形, 这恐怕是数不清的. 因此我们必须换一个定义方式, 我们可以通过圆心坐标和半径的长度来确定一个圆形. `这是一种新的定义值的方法. 它(类型构造器)是一个函数, 通过这个函数来产生属于这个集合的元素(值)`.
```
data Circle = Circle Float Float Float
```
考虑以下形状类型, 它可以包含圆形或者是矩形, 于是可以这样定义:
```
data Shape
  = Circle    Float Float Float 
  | Rectangle Float Float Float Float
```
有了数据类型, 接下来定义几个处理这种类型数据的函数, 对于Shape, 可以计算它的面积.
```hs
area:: Shape -> Float
area (Circle _ _ r) = 3.14 * r * r
area (Rectangle x1 y1 x2 y2) = (abs (x1 - x2)) * (abs (y1 - y2))
```
```div class=note
其中:
​```hs
area (Rectangle x1 y1 x2 y2) = (abs (x1 - x2)) * (abs (y1 - y2))
```
可以使用`$`函数来少些几个括号:
```hs
area (Rectangle x1 y1 x2 y2) = (abs $ x1 - x2) * (abs $ y1 - y2)
```
```hs
> :i ($)
($) :: (a -> b) -> a -> b       -- Defined in `GHC.Base'
infixr 0 $
```
 1. 首先, $函数的优先级是最低的
 2. 其次`function $ argument`等价于`function expr`, 如果expr是一个值, 这种写法没有任何方便之处,但是当expr是一个表达式, $就有些作用了, 因为它的优先级最低, 所以优先计算expr, 然后再作为参数传递给function, 于是避免显示指定优先级`function (expr)`. 最终你少写两个括号.
```
因为类型构造器本质上就是函数, 所以我们可以交给Higher-Order函数处理:
​```hs
> map (Rectangle 0 0 1) [1..5]
[Rectangle 0.0 0.0 1.0 1.0,Rectangle 0.0 0.0 1.0 2.0,Rectangle 0.0 0.0 1.0 3.0,Rectangle 0.0 0.0 1.0 4.0,Rectangle 0.0 0.0 1.0 5.0]
```
通过引入中间类型, 我们可以把Shape定义的更加清晰:
```hs
data Point = Point Float Float deriving (Show)
data Shape
  = Circle Point Float
  | Rectangle Point Point
  deriving (Show)
-- area函数也作必要修改
area :: Shape -> Float
area (Circle _ r) = pi * r ^ 2
area (Rectangle (Point x1 y1) (Point x2 y2)) = (abs $ x2 - x1) * (abs $ y2 - y1)
```
# Type Parameters
类型构造器不仅可以使用值(Value Parameters)作为生成其类型的参数, 还可以使用类型作为输入参数.
Maybe类型(See: Haskell/TypeClass/Maybe):
```hs
> :i Maybe
data Maybe a = Nothing | Just a
> :i Just
data Maybe a = ... | Just a
```
See: [Haskell/TypeSystem/ConcreteType]
# 实际的例子
我们来定义一个新的类型Day来代表一周7天.
```hs
data Day = Monday | Tuesday | Wednesday | Thursday | Friday | Saturday | Sunday
```
有了类型定义,接下来我们要通过简单的方式来扩展这个类型, 使他的具备一些实用的功能.
# 打印: Show
```hs
> Monday
<interactive>:34:1:
    No instance for (Show Day) arising from a use of `print'
    Possible fix: add an instance declaration for (Show Day)
    In a stmt of an interactive GHCi command: print it
>
```
这是应为, GHCI通过`print`函数来完成打印的, 我们来看看它的函数签名:
```hs
> :i print
print :: Show a => a -> IO ()   -- Defined in `System.IO'
```
必须是属于[Haskell/TypeClass/Show Show] TypeClass的才能够被print函数处理.
GHCI识别出错误之后,还会推荐我们如何处理这个错误, 她建议我们使用instance方式来实现Show中定义的函数, 当然更简便的方法是直接使用deriving关键字, 让系统为我们自动添加默认的处理方式.
接下来我们稍微修改以下定义:
```hs
data Day = Monday | Tuesday | Wednesday | Thursday | Friday | Saturday | Sunday
    deriving Show
```
再试一下:
```hs
> Monday
Monday
```
# 相等或是不等: Eq
```hs
data Day = Monday | Tuesday | Wednesday | Thursday | Friday | Saturday | Sunday
    deriving (Show, Eq)
```
```hs
> Monday == Monday
True
> Monday == Tuersday
False
```
# 有序的: Ord
```hs
data Day = Monday | Tuesday | Wednesday | Thursday | Friday | Saturday | Sunday
    deriving (Show, Eq, Ord)
```
```hs
> max Monday Tuesday
Tuesday
> min Monday Tuesday
Monday
> Monday > Tuesday
False
> Monday < Tuesday
True
```
# 有界的: Bounded
```hs
data Day = Monday | Tuesday | Wednesday | Thursday | Friday | Saturday | Sunday
    deriving (Show, Eq, Ord, Bounded)
```
```hs
> minBound :: Day
Monday
> maxBound :: Day
Sunday
```
# 可枚举的: Enum
```hs
data Day = Monday | Tuesday | Wednesday | Thursday | Friday | Saturday | Sunday
    deriving (Show, Eq, Ord, Bounded, Enum)
```
```hs
> succ Tuesday
Wednesday
> pred Tuesday
Monday
> toEnum 0 :: Day
Monday
> fromEnum Tuesday
1
> [minBound .. maxBound] :: [Day]
[Monday,Tuesday,Wednesday,Thursday,Friday,Saturday,Sunday]
> [Tuesday .. Saturday]
[Tuesday,Wednesday,Thursday,Friday,Saturday]
> [Monday ..]
[Monday,Tuesday,Wednesday,Thursday,Friday,Saturday,Sunday]
```
[[TOC]]
# 类型系统
# 基本概念
# 什么是类型?
A type is a collection of related values
# 强类型
# 静态类型
# 可以自动推导
# 强类型
# 静态
静态类型系统意味着编译器在编译之前就明确知道每个值或者表达式的具体类型。
静态类型有时让书写某些代码变得困难， 如Python中ducktyping
Haksell通过typeclasses来实现动态类型，它更加安全，也十分方便。
在传统的动态类型系统中，人然需要编写大量的测试用例来检测类型转换错误。一旦覆盖不到，就可能产生运行期错误。haskell的静态类型+Strong类型系统有效的避免了这些错误。
# 类型推导
编译器可以自动推导出所有表达式的类型。
# 从类型系统中获益
While strong, static typing makes Haskell safe, type inference makes it concise.
# 基本类型
# Char
Unicode字符
# Bool
True 或者 False
# Int: 固定长度的有符号整数
可能是32位，也可能是64位，Haksell标准可以保证至少为28位。
# Integer : 有符号无上界整数，大整数，效率和开销上都要大很多
# Double: 64位浮点数
另为有32位浮点数Float, 不过不推荐你使用它，Haskell编译器的作者们正在使劲优化double,所以它的效率要高一些。
# List
List是polymorphic类型，这种类型是通过TypeVariable来实现的，类型变量必须小写，实际上充当一个placeholder的角色，最终会被替换为实际的类型。
由于空List的类型是`[a]`， 着说明List中只能存放一种类型的元素。当a已知以后，我们就会得到更精确的类型签名:
```
ghci> :t [1,2,3]
[1,2,3] :: Num t => [t]
ghci> :t [[True],[False]]
[[True],[False]] :: [[Bool]]
```
```
类型名必须大写开头，而类型变量必须小写开头
```
# 查看类型
```
ghci> :type 'x'
'x' :: Char
ghci> :t 1
1 :: Num a => a
ghci> :t []
[] :: [a]
```
`:: Char`又被叫做TypeSignature
# 有关haskell类型系统的另外一个作用恐怕就是区分纯函数和非纯函数
在haskell中IO类型是一个特殊的类型，凡是在类型签名中带有IO的，这种函数都是非纯函数，所以Haskell类型系统可以避免你将纯函数和非纯函数混搭在一起。
# 添加新类型
```
data [<constraints> =>] <type-name> = <constuctor-1> | <constructor-2> ... | <constructor-n>
    deriving (<typeclass-1>, ... , <typeclass-n>)
```
# 重载
如果我们希望自定义类型属于某个Typeclass, 那么我们必须给出那个Typeclass所要求实现的接口(函数). 可以通过以下表达式来完成:
```
instance <typeclass> (<type-name> <type-variable>)
    where 
        <implement-function-1>
        ...
        <implement-function-n>
```
[[TOC]]
# Record
假如People类型
People:
 * name::String
 * age::Int
可以这么来定义:
```
data People = People String Int deriving (Show)
peopleName :: People -> String
peopleName (People name _)  = name
peopleAge :: People -> Int
peopleAge (People _ age) = age  
```
```hs
> let p=People "amas" 20
> p
People "amas" 20
> peopleName p
"amas"
> peopleAge p
20
```
 1. 如果数据项多一点的话, 我们很容易搞混淆. 
 2. 每当定义数据时, 我们必须按照定义的顺序来安排实际的参数,  这样一点儿都不方便. 
 3. 必须定义一系列的访问方法, 机械而无趣
使用Record可以克服这些问题:
```hs
data People = People { name::String , age::Int } deriving (Show)
```
```hs
> :t name
name :: People -> String
> :t age
age :: People -> Int
> let p=People{name="amas",age=10}
People {name = "amas", age = 10}
> name p
"amas"
> age p
10
```
[[TOC]]
# 递归数据结构 (Recursive Data Structures)
# List
```hs
data List a = Empty | Cons a (List a) deriving (Show, Read, Eq, Ord)
```
```hs
> Empty
Empty
> Cons 1 Empty
Cons 1 Empty
> Cons 1 (Cons 2 Empty)
Cons 1 (Cons 2 Empty)
> Cons 1 (Cons 2 (Cons 3 Empty)
Cons 1 (Cons 2 (Cons 3 Empty)
```
可以通过记录的方式修改一下定义:
```hs
data List a = Empty | Cons { listHead :: a, listTail :: List a}
             deriving (Show, Read, Eq, Ord)
```
来构造几个简单的List:
```hs
> Empty
Empty
> Cons 1 Empty
Cons 1 Empty
```
列表`[1,2,3]`可以表示为:
```hs
> Cons {listHead = 1, listTail = Cons {listHead = 2, listTail = Cons {listHead = 3, listTail = Empty}}}
Cons {listHead = 1, listTail = Cons {listHead = 2, listTail = Cons {listHead = 3, listTail = Empty}}}
-- 或者更为简单的型式, 我们用括号明确一下构造顺序
> Cons 1 (Cons 2 (Cons 3 Empty))
Cons {listHead = 1, listTail = Cons {listHead = 2, listTail = Cons {listHead = 3, listTail = Empty}}}
```
```div class=note
如果我们将`Cons`改为中缀表达式, 则` Cons 1 (Cons 2 (Cons 3 Empty))` 可表示为:
```
1 `Cons` (2 `Cons` (3 `Cons` Empty))
```
回忆一下Haskell中List的构造方法:
```
1:(2:(3:[]))
```
是不是看到二者本质上是一样的?
```
# 中缀式构造器
我们也可以使用运算符号来充当构造器, 必须以`:`开头后面可以是一系列的运算符号, 不能出现字母或数字. 
```hs
data List a = Empty | ::: a (List a)
            deriving (Show, Read, Eq, Ord)
```
以下型式都是合法的.
```
infixr 5 :@
data List a = Empty | a :@ (List a)
            deriving (Show, Read, Eq, Ord)
infixr 5 :--
data List a = Empty | a :-- (List a)
            deriving (Show, Read, Eq, Ord)
```
中缀式构造器仍然是左结合的, 在使用中存在一些不方便之处. 比方说: List的构造过程中不得不书写大量的括号, 非常之繁琐.
# 右结合的中缀式构造器和infixr
我们可以使用`infixr`来修改构中缀造器为右结合的
```hs
infixr 5 :::
data List a = Empty | a ::: (List a)
            deriving (Show, Read, Eq, Ord)
```
于是, List的构造可改写为以下型式, 节约了大量的括号.
```hs
> 1 ::: 2 ::: 3 ::: Empty
1 ::: (2 ::: (3 ::: Empty))
```
# 增加一些List操作函数
```hs
infix 5 +++
(+++) :: List a -> List a -> List a
Empty +++ xs = xs
xs +++ Empty = xs
(x ::: xs) +++ ys = x ::: (xs +++ ys) -- 注意: (x ::: xs) 所匹配的是 a ::: (List a) 这个构造器
--length of the List
listLength :: List a -> Int
listLength Empty = 0
listLength (x ::: xs) = 1 + (listLength xs)
```
```hs
> let xs=1 ::: 2 ::: 3 ::: Empty
> let ys=4 ::: 5 ::: 6 ::: Empty
> xs +++ ys
1 ::: (2 ::: (3 ::: (4 ::: (5 ::: (6 ::: Empty)))))
> listLength xs
3
```
```div class=warn
PatternMatching的本质是构造器的模式匹配
```
# Pair
```hs
data Pair a b = Pair a b deriving Show
```
```hs
> Pair 1 2
Pair 1 2
-- 如果构造器只接受两个参数, 则也可以使用中缀表达式
> 1 `Pair` 2
Pair 1 2
-- 使用括号明确构造顺序
> Pair 1 (Pair 3 (Pair 4 5))
Pair 1 (Pair 3 (Pair 4 5))
```
[[TOC]]
# Type Declarations
[[TOC]]
# Type Synonyms
类型同义词, 又可称为类型别名. 最典型的便是Haskell中的String类型, 其本质实为`[Char]`, 即字符数组.
```hs
> :i String
type String = [Char]    -- Defined in `GHC.Base'
```
# type
通过type关键字可以定义新的类型别名.
# 举个例子
# String Map
```hs
type StringMap = [(String, String)]
```
亦可定义为:
```hs
type Key   = String
type Value = String
type StringMap = [(Key, Value)]
```
# type + 类型变量
```hs
type Map k v = [(k, v)]
--| The 'contains' function search the specify key in map, if the key existed True otherwise False
contains :: Eq k => (Map k v) -> k -> Bool
contains [] _ = False
contains (x:xs) key
  | fst x == key = True
  | otherwise = contains xs key
```
```hs
> let map = [(1,"one"),(2,"two"),(3,"three")] :: Map Int String
[(1,"one"),(2,"two"),(3,"three")]
> contains 1 map
True
> contains 4 map
False
```
[[TOC]]
# Weak Head Normal Form 
 * An expression in weak head normal form has been evaluated to the outermost data constructor or lambda abstraction (the head). 
 * Sub-expressions may or may not have been evaluated. 
 * Therefore, every normal form expression is also in weak head normal form, though the opposite does not hold in general.
To determine whether an expression is in weak head normal form, we only have to look at the outermost part of the expression. If it's a data constructor or a lambda, it's in weak head normal form. If it's a function application, it's not.
These expressions are in weak head normal form:
```hs
(1 + 1, 2 + 2)       -- the outermost part is the data constructor (,)
 > 2 + 2          -- the outermost part is a lambda abstraction
'h' : ("e" ++ "llo") -- the outermost part is the data constructor (:)
```
As mentioned, all the normal form expressions listed above are also in weak head normal form.
These expressions are not in weak head normal form:
```hs
1 + 2                -- the outermost part here is an application of (+)
( > x + 1) 2      -- the outermost part is an application of ( > x + 1)
"he" ++ "llo"        -- the outermost part is an application of (++)
```
# 参考
 * http://stackoverflow.com/questions/6872898/haskell-what-is-weak-head-normal-form

