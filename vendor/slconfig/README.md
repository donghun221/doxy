# slconfig模块说明

  配置文件使用INI格式。

  由 [section] 的节构成, 节内部使用 name=value 键值对 如果没有指定节,则默认放入名为 [default] 的节当中. "#"或";"为注释的开头,可以放置于任意的单独一行中。如下所示

    # test config
    debug = true
    url = act.wiki
    
    ; redis config
    [redis]
    redis.key = push1,push2
    
    ; mysql config
    [mysql]
    mysql.dev.host = 127.0.0.1
    mysql.dev.user = root
    mysql.dev.pass = 123456
    mysql.dev.db = test
    
    mysql.master.hosT = 10.0.0.1
    mysql.master.user = root
    mysql.master.pass = 89dds)2$#d
    mysql.master.db = act
    
    ; math config
    [math]
    math.i64 = 64
    math.f64 = 64.1
    
    ; other config
    [other]
    name = ATC自动化测试^-^&($#……#
    key1 = test key

## 引用方式

    import {
        "slconfig"
    }

## 接口

	slconfig.NewConfig(confName string) (ConfigInterface, error)
	slconfig.ReloadConfig(ci ConfigInterface, confName string) (ConfigInterface, error) // reload配置文件的参数并添加程序中自己Set设置的新配置参数

  ConfigInterface类型包含以下方法

    String(key string) string
    Strings(key string) []string
    Bool(key string) (bool, error)
    Int(key string) (int, error)
    Int64(key string) (int64, error)
    Float64(key string) (float64,error)
    Set(key string, value string) error // 可以修改已有参数，也可以新增参数

## 测试DEMO

  单元测试模块slconfigtest.go，后续持续更新

    git clone http://git.code.oa.com/All-In-OneNAS/slconfig.git

    cd slconfig
    chmod +x ./build.sh
    ./build.sh
    ./slconfigtest  -configfile="./src/slconfigtest/slconfigtest.conf"

  执行结果

    load config file ./src/slconfigtest/slconfigtest.conf successed
    get config "debug":%!s(bool=true)
    get config "url":act.wiki
    get config "redis::redis.key":[push1 push2]
    get config "mysql::mysql.dev.host":127.0.0.1
    get config "mysql::mysql.master.host":10.0.0.1
    get config "math::math.i64":%!s(int64=64)
    get config "math::math.f64":%!s(float64=64.1)
    get config "other::name":ATC自动化测试^-^&($#……#
    get config "other::key1":test key
    get config "other::key2":test key2

    reload config file ./src/slconfigtest/slconfigtest.conf successed
    reload get config "other::name":ATC自动化测试^-^&($#……#
    reload get config "other::key1":test key
    reload get config "other::key2":test key2

## 配置热加载

### 命令行工具修改

  可以通过命令行工具(使用flags)直接动态更新主进程的配置，不更改配置文件的配置。具体使用参考测试DEMO

### 修改配置文件(to be continued)

  watch进程定期检查配置文件的修改时间和md5校验值，如发生改变，则重新加载配置文件或通过命令行工具更新主进程的配置。


### 配置管理/下发(to be continued)

  采用server-agent架构,server端配置不同配置并设置对应的服务器（或模块角色），agent去server拉取使用。

  后续进行，临时先在本地修改配置