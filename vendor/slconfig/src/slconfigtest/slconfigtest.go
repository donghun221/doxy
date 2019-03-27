// slconfigtest
package main

import (
	"flag"
	"fmt"
	"slconfig"
)

var configPath = flag.String("configfile", "slconfigtest.conf", "General configuration file")

func main() {
	flag.Parse()

	config, err := slconfig.NewConfig(*configPath)
	if err != nil {
		fmt.Printf("load config file %v failed, err:%v\n", *configPath, err.Error())
		return
	}
	fmt.Printf("load config file %v successed\n", *configPath)

	// default::key test
	if v, err := config.Bool("debug"); err != nil || v != true {
		fmt.Printf("Get failure: expected different value for debug (expected: [%#v] got: [%#v])\n", true, v)
	} else {
		fmt.Printf("get config \"debug\":%s\n", v)
	}
	if v := config.String("url"); v != "act.wiki" {
		fmt.Printf("Get failure: expected different value for url (expected: [%#v] got: [%#v])\n", "act.wiki", v)
	} else {
		fmt.Printf("get config \"url\":%s\n", v)
	}

	// reids::key test
	if v := config.Strings("redis::redis.key"); len(v) != 2 || v[0] != "push1" || v[1] != "push2" {
		fmt.Printf("Get failure: expected different value for redis::redis.key (expected: [%#v] got: [%#v])\n", "[]string{push1,push2}", v)
	} else {
		fmt.Printf("get config \"redis::redis.key\":%s\n", v)
	}
	if v := config.String("mysql::mysql.dev.host"); v != "127.0.0.1" {
		fmt.Printf("Get failure: expected different value for mysql::mysql.dev.host (expected: [%#v] got: [%#v])\n", "127.0.0.1", v)
	} else {
		fmt.Printf("get config \"mysql::mysql.dev.host\":%s\n", v)
	}
	if v := config.String("mysql::mysql.master.host"); v != "10.0.0.1" {
		fmt.Printf("Get failure: expected different value for mysql::mysql.master.host (expected: [%#v] got: [%#v])\n", "10.0.0.1", v)
	} else {
		fmt.Printf("get config \"mysql::mysql.master.host\":%s\n", v)
	}

	// math::key test
	if v, err := config.Int64("math::math.i64"); err != nil || v != 64 {
		fmt.Printf("Get failure: expected different value for math::math.i64 (expected: [%#v] got: [%#v])\n", 64, v)
	} else {
		fmt.Printf("get config \"math::math.i64\":%s\n", v)
	}
	if v, err := config.Float64("math::math.f64"); err != nil || v != 64.1 {
		fmt.Printf("Get failure: expected different value for math::math.f64 (expected: [%#v] got: [%#v])\n", 64.1, v)
	} else {
		fmt.Printf("get config \"math::math.f64\":%s\n", v)
	}

	// other::key test
	if v := config.String("other::name"); v != "ATC自动化测试^-^&($#……#" {
		fmt.Printf("Get failure: expected different value for other::name (expected: [%#v] got: [%#v])\n", "ATC自动化测试^-^&($#……#", v)
	} else {
		fmt.Printf("get config \"other::name\":%s\n", v)
	}
	if v := config.String("other::key1"); v != "test key" {
		fmt.Printf("Get failure: expected different value for other::key1 (expected: [%#v] got: [%#v])\n", "test key", v)
	} else {
		fmt.Printf("get config \"other::key1\":%s\n", v)
	}

	config.Set("other::key2", "test key2")
	if v := config.String("other::key2"); v != "test key2" {
		fmt.Printf("Get failure: expected different value for other::key2 (expected: [%#v] got: [%#v])\n", "test key2", v)
	} else {
		fmt.Printf("get config \"other::key2\":%s\n", v)
	}

	config, err = slconfig.ReloadConfig(config, *configPath)
	if err != nil {
		fmt.Printf("\nreload config file %v failed, err:%v\n", *configPath, err.Error())
		return
	}
	fmt.Printf("\nreload config file %v successed\n", *configPath)

	// after reload other::key test
	if v := config.String("other::name"); v != "ATC自动化测试^-^&($#……#" {
		fmt.Printf("reload Get failure: expected different value for other::name (expected: [%#v] got: [%#v])\n", "ATC自动化测试^-^&($#……#", v)
	} else {
		fmt.Printf("reload get config \"other::name\":%s\n", v)
	}
	if v := config.String("other::key1"); v != "test key" {
		fmt.Printf("reload Get failure: expected different value for other::key1 (expected: [%#v] got: [%#v])\n", "test key", v)
	} else {
		fmt.Printf("reload get config \"other::key1\":%s\n", v)
	}
	if v := config.String("other::key2"); v != "test key2" {
		fmt.Printf("reload Get failure: expected different value for other::key2 (expected: [%#v] got: [%#v])\n", "test key2", v)
	} else {
		fmt.Printf("reload get config \"other::key2\":%s\n", v)
	}
}
