---
title: Android Gradle 作弊手册
tags:
---
# Android Gradle 作弊手册
<!-- toc -->
## 优化资源
```
    defaultConfig {
        ...
        resConfigs "en", "zh"
    }

```

## 如何打一个可执行的jar包
```
apply plugin: 'java'

dependencies {
    compile fileTree(include: ['*.jar'], dir: 'libs')
    //compile group: 'org.smali', name: 'dexlib2', version: '2.0.3'
    compile files('libs/libhello.jar')
    compile files('libs/jsr305-1.3.9.jar')
    compile files('libs/dexlib2-2.2.1.jar')
    compile files('libs/guava-18.0.jar')
    compile files('libs/util-2.2.1.jar')
    compile files('libs/jcommander-1.72.jar')
}

def mainClassName = "com.example.Utils"

jar {
    manifest {
        attributes "Main-Class": "$mainClassName"
    }

    from {
        configurations.compile.collect { print "--> "+it; it.isDirectory() ? it : zipTree(it) }
    }
}

```



## 如何引入外部依赖?
```
# libs目录下面的jar文件也加入编译
compile	fileTree(dir:	‘libs’,	include:	[‘*.jar’])

# 引用子模块
compile project(':libjava')

compile files('libs/foo.jar', 'libs/bar.jar')\

# 这俩是一个意思
compile 'com.example.android:app-magic:12.3'
compile group: 'com.example.android', name: 'app-magic', version: '12.3'


# 不同的varint引入不同的模块
dependencies {
  // Adds the 'debug' varaint of the library to the debug varaint of the app
  debugCompile project(path: ':my-library-module', configuration: 'debug')

  // Adds the 'release' varaint of the library to the release varaint of the app
  releaseCompile project(path: ':my-library-module', configuration: 'release')
}
```


如何查看依赖?
```
$ ./gradlew -q dependencies your-app-project:dependencies
```

## 如何配置签名?
```
android {
    
    signingConfigs	{
    	release	{
            storeFile file(“keystore.release”)
            storePassword System.getenv(“KEYSTOREPASSWD”)
            //或通过终端输入 System.console().readLine (“\nEnter Passwords:“)						
            keyAlias “your key alias here”
            keyPassword System.getenv(“KEYPASSWD”)
        }
    }   
    buildTypes {
        
    }
}
```


## 如何通过gradlew打包?
```
$ ./gradlew <task>
$ ./gradlew assembleDebug
$ ./gradlew assembleRelease

## 如何用命令行打包flavors版本
```
## Build Variant
```
android {
   productFlavors {
        cn {
            applicationId "a.m.a.s.ego_cn"
            versionName "cn_1.0"
        }

        en {
            applicationId "a.m.a.s.ego_en"
            versionName "en_1.0"
        }
    }
}
```
 * cn/en就是flavors的id, 这个玩意回头要建立对应的目录保存各自的资源/代码
 * 修改了当前的buildflavor后, 可以编译一下

### 一个例子
 * Init.java可以有各自的定义
 * activity_main.xml也可以有各自的定义
```
└── src
    ├── cn
    │   ├── java
    │   │   └── a
    │   │       └── m
    │   │           └── a
    │   │               └── s
    │   │                   └── ego
    │   │                       └── custom
    │   │                           └── Init.java
    │   └── res
    │       └── layout
    │           └── activity_main.xml
    ├── en
    │   ├── java
    │   │   └── a
    │   │       └── m
    │   │           └── a
    │   │               └── s
    │   │                   └── ego
    │   │                       └── custom
    │   │                           └── Init.java
    │   └── res
    │       └── layout
    │           └── activity_main.xml
    ├── main
    │   ├── AndroidManifest.xml
    │   ├── java
    │   │   └── a
    │   │       └── m
    │   │           └── a
    │   │               └── s
    │   │                   └── ego
    │   │                       ├── custom
    │   │                       │   └── IInit.java
    │   │                       └── MainActivity.java
    │   └── res


```

### Android的gradle插件在哪里?
已经包含在android-studio中, 你可以到这里找到:
```
$ cd android-studio
$ cd /gradle/m2repository/com/android/tools/build/builder/2.3.1/
$ ls *
gradle-api-2.3.1.jar
```
在编辑gradle文件的时候可以F3查看在插件中是如何实现的

### 我们来看看flavor在android gradle plugin中是怎么定义的
```
package com.android.build.gradle.internal.dsl;
...

public class ProductFlavor extends DefaultProductFlavor implements CoreProductFlavor {
    protected final Project project;
    protected final Logger logger;
    private final NdkOptions ndkConfig;
    private final ExternalNativeBuildOptions externalNativeBuildOptions;
    private final ErrorReporter errorReporter;
    private final JackOptions jackOptions;
    private final JavaCompileOptions javaCompileOptions;
    private final ShaderOptions shaderOptions;

    public ProductFlavor(String name, Project project, Instantiator instantiator, Logger logger, ErrorReporter errorReporter) {
        super(name, (DefaultVectorDrawablesOptions)instantiator.newInstance(VectorDrawablesOptions.class, new Object[0]));
        this.project = project;
        this.logger = logger;
        this.errorReporter = errorReporter;
        this.ndkConfig = (NdkOptions)instantiator.newInstance(NdkOptions.class, new Object[0]);
        this.externalNativeBuildOptions = (ExternalNativeBuildOptions)instantiator.newInstance(ExternalNativeBuildOptions.class, new Object[]{instantiator});
        this.jackOptions = (JackOptions)instantiator.newInstance(JackOptions.class, new Object[0]);
        this.javaCompileOptions = (JavaCompileOptions)instantiator.newInstance(JavaCompileOptions.class, new Object[]{instantiator});
        this.shaderOptions = (ShaderOptions)instantiator.newInstance(ShaderOptions.class, new Object[0]);
    }
```

更加全面的文档可以参考: [GoogleGradleDSL](http://google.github.io/android-gradle-dsl/)

## APK的优化
必要条件:
 * SDK Tools 25.0.10 or higher
 * Android plugin for Gradle 2.0.0 or higher

### 代码层面
基于Proguard,删掉无用的类方法等等
```
android {
    buildTypes {
        release {
            minifyEnabled true
            ...
        }
    }
    ...
}
```

### 资源层面
去掉无用的资源

有时候有一些特殊的资源不希望被删掉,可以这样配置:
```
<?xml version="1.0" encoding="utf-8"?>
<resources xmlns:tools="http://schemas.android.com/tools"
    tools:keep="@layout/l_used*_c,@layout/l_used_a,@layout/l_used_b*"
    tools:discard="@layout/unused2" />
```
保存在这个文件中就可以了 res/raw/keep.xml. 

```
Note: The resource shrinker currently does not remove resources defined in a values/ folder (such as strings, dimensions, styles, and colors). This is because the Android Asset Packaging Tool (AAPT) does not allow the Gradle Plugin to specify predefined versions for resources. For details, see issue 70869.

通过名字引用资源也想给删掉?
<?xml version="1.0" encoding="utf-8"?>
<resources xmlns:tools="http://schemas.android.com/tools"
    tools:shrinkMode="strict" />
```
