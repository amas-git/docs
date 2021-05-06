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

