# 云原生应用



```
The most important property of a program is whether it accomplishes the intention of
its user.1
                             — C.A.R. Hoare, Communications of the ACM (October 1969)
                             
# quicksort和CSP的发明人                             
```



## 什么是云原生?

云原生要解决服务可靠性的问题.

```
if cloud native has to be a synonym for anything, it would be idempotent.
if cloud native has to be a synonym for anything, it would be dependability.

The dependability of a computer system is its ability to avoid failures that are more
frequent or more severe, and outage durations that are longer, than is acceptable to the
user(s).8
—Fundamental Concepts of Computer System Dependability (2001)
```



## Dependability的内涵

- Attributes
  - Availability
  - Reliability
  - Maintainability
- Threats
  - Faults
  - Errors
  - Failures
- Means
  - Faults Prevetion
    - 最佳编程实践
    - 开发语言的特性
    - Scaling
      - Scaling Up (垂直扩容)
      - Scaling Out (水平扩容)
    - 松耦合
  - Faults Removal
    - 验证和测试
      - 静态分析
      - 动态分析
    - 管理能力
      - 可以方便灵活的控制程序运行的行为, 以适应外部的变化
  - Faults Tolerance
    - self-repair
    - self-healing,
    - resilience
  - Faults Forcasting
    - 故障模式
    - 故障影响分析
    - 压力测试

###  Dependability != Reliability

Dependability并不好量化, 所以SRE工程师, 里的R代表Reliability, 是可以量化的



## 12 Factor

```
2010年,Heroku起草了The Twelve-Factor App, 以解决如下问题:
1. 最小化新成员进入项目的时间
2. 最大化兼容各种运行环境
3. 适合运行在云平台上
4. 减少生产环境和开发环境的差异,最大化敏捷
5. 可以很容易扩容
```

### 1. 代码库

### 2. 依赖管理

### 3. 配置管理

将配置项保存到环境变量中

```
$ go get github.com/spf13/viper
```

```
Viper会按照下面的优先级。每个项目的优先级都高于它下面的项目:

显示调用Set设置值
命令行参数（flag）
环境变量
配置文件
key/value存储
默认值
重要： 目前Viper配置的键（Key）是大小写不敏感的。目前正在讨论是否将这一选项设为可选。
```



```
viper.SetDefault("ContentDir", "content")
viper.SetDefault("LayoutDir", "layouts")
viper.SetDefault("Taxonomies", map[string]string{"tag": "tags", "category": "categories"})
```

```go
// Viper需要最少知道在哪里查找配置文件的配置。Viper支持JSON、TOML、YAML、HCL、envfile和Java properties格式的配置文件。Viper可以搜索多个路径，但目前单个Viper实例只支持单个配置文件。Viper不默认任何配置搜索路径，将默认决策留给应用程序。
// 下面是一个如何使用Viper搜索和读取配置文件的示例。不需要任何特定的路径，但是至少应该提供一个配置文件预期出现的路径。

viper.SetConfigFile("./config.yaml") // 指定配置文件路径
viper.SetConfigName("config") // 配置文件名称(无扩展名)
viper.SetConfigType("yaml") // 如果配置文件的名称中没有扩展名，则需要配置此项
viper.AddConfigPath("/etc/appname/")   // 查找配置文件所在的路径
viper.AddConfigPath("$HOME/.appname")  // 多次调用以添加多个搜索路径
viper.AddConfigPath(".")               // 还可以在工作目录中查找配置
err := viper.ReadInConfig() // 查找并读取配置文件
if err != nil { // 处理读取配置文件的错误
	panic(fmt.Errorf("Fatal error config file: %s \n", err))
}


if err := viper.ReadInConfig(); err != nil {
    if _, ok := err.(viper.ConfigFileNotFoundError); ok {
        // 配置文件未找到错误；如果需要可以忽略
    } else {
        // 配置文件被找到，但产生了另外的错误
    }
}

// 配置文件找到并成功解析



// 写配置文件
viper.WriteConfig() // 将当前配置写入“viper.AddConfigPath()”和“viper.SetConfigName”设置的预定义路径
viper.SafeWriteConfig()
viper.WriteConfigAs("/path/to/my/.config")
viper.SafeWriteConfigAs("/path/to/my/.config") // 因为该配置文件写入过，所以会报错
viper.SafeWriteConfigAs("/path/to/my/.other_config")


// 监控配置文件的改变, 实时调整程序
viper.WatchConfig()
viper.OnConfigChange(func(e fsnotify.Event) {
  // 配置文件发生变更之后会调用的回调函数
	fmt.Println("Config file changed:", e.Name)
})


// 绑定环境变量
SetEnvPrefix("spf") // 将自动转为大写
BindEnv("id")

os.Setenv("SPF_ID", "13") // 通常是在应用程序之外完成的

id := Get("id") // 13
```





#### 4. 将服务视为资源, 用URL表示

```
Treat backing services as attached resources.
—The Twelve-Factor App
```



#### 5. 构建, 发布, 运行

```
Strictly separate build and run stages.
—The Twelve-Factor App
```

- BUILD
  - 从代码仓库获取指定版本
  - 获取依赖
  - 编译
  - 每次build需要有唯一ID和时间戳
- RELEASE
  - 发布的时候, 是将build和对应的Deployment的配置所绑定
- RUN
  - release被发布到Deployment环境中

```
CODE -> BUILD  ->
                | -> RELEASE -> RUN
        CONFIG ->
```





#### 6. 进程

```
Execute the app as one or more stateless processes.
—The Twelve-Factor App
```

服务进程无状态, 状态都保存到后端服务中(如数据库,缓存服务)

#### 7. 数据隔离

服务是自给自足的, 如果需要外部访问数据, 可以提供API.

#### 8. 尽可能的水平扩展

#### 9. 一次性

快速启动, 优雅关机

为什么要这么干呢? 云环境复杂易变, 服务迁移非常频繁, 所以我们的服务尽可能小, 可以快速启动, 优雅停止.这样会最大限度的减少停机时间.

#### 10. 保持开发/生产环境尽可能相同

- 开发分支生命周期尽可能端, 迅速合并到主干, 快速上生产
- 开发/生产环境使用一致的软件,利用容器解决

- 开发者熟悉部署, 熟悉生产环境, 解决运营研发孤岛问题

#### 11. 日志

将日志视为事件流

服务不间断运行, 服务不断产生日志即为服务运行时产生的时间, 随进程开始开始, 随进程终结终结. 将日志写 到STDOUT, 免于落本地盘, 省去日志管理成本,另一方面将STDOUT流交给日志服务, 统一管理, 分析,方便查询.

#### 12. 管理进程

用一次性进程管理服务, 而不是手动管理.





### 缓存

- LRU : https://github.com/hashicorp/golang-lru



### 插件

- go plugin
- hashcorp plugin

### Hexagonal Architecture





### 重试

- backoff算法

### 限流

```sh
$ go get golang.org/x/time/rate@latest
```



```
Before we start writing any code, let’s take a moment to explain how token-bucket rate
limiters work. The description from the official x/time/rate documentation says:
A Limiter controls how frequently events are allowed to happen. It implements a “token
bucket” of size b , initially full and refilled at rate r tokens per second.
Putting that into the context of our API application…

 1. We will have a bucket that starts with b tokens in it.
 2. Each time we receive a HTTP request, we will remove one token from the bucket.
 3. Every 1/r seconds, a token is added back to the bucket — up to a maximum of b total
tokens.
 4.  If we receive a HTTP request and the bucket is empty, then we should return a
429 Too Many Requests response
```



### 优雅降级

### 优雅关机

### 熔断

 token buket

### 超时

如果你认为失败, 那么要快速失败



### Idempotence

```

```

### Redundancy



### Auto Scaling



### 健康检查

- Liveness: 进程活着
- Shallow: 进程活着且本地状态正常
- Deep: 进程或者且可以正常服务

### 服务的管理能力

- 配置和控制
  - 配置文件标准化: JSON/YAML/TOML
  - 利用环境变量进行配置
  - 配置与代码分离
  - 配置变更需要有版本管理
- 监控,报警,日志
- 发布和更新
- 服务发现

```go
github.com/spf13/cobra
github.com/spf13/viper


package main
import (
    "fmt"
    "os"
    "github.com/spf13/cobra"
)
var strp string
var intp int
var boolp bool
var rootCmd = &cobra.Command{
    Use: "flags",
    Long: "A simple flags experimentation command, built with Cobra.",
    Run: flagsFunc,
}

func init() {
    rootCmd.Flags().StringVarP(&strp, "string", "s", "foo", "a string")
    rootCmd.Flags().IntVarP(&intp, "number", "n", 42, "an integer")
    rootCmd.Flags().BoolVarP(&boolp, "boolean", "b", false, "a boolean")
}

func flagsFunc(cmd *cobra.Command, args []string) {
    fmt.Println("string:", strp)
    fmt.Println("integer:", intp)
    fmt.Println("boolean:", boolp)
    fmt.Println("args:", args)
}


func main() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
```

### 性能测试

```sh
 $ go test -gcflags=-N -benchmem -test.count=3 -test.cpu=1 -test.benchtime=1s -bench=.
```

