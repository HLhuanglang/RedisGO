package main

import (
	"fmt"

	"redisgo/base/config"
	"redisgo/core/server"
	"redisgo/core/transport"
)

func main() {
	//1,初始化配置
	config.SetupConfig("redis.conf")

	//2,启动tcp服务
	err := transport.ListenAndServeWithSignal(&transport.Config{
		Address: fmt.Sprintf("%s:%d", config.RedisGoCfg.Bind, config.RedisGoCfg.Port),
	}, server.NewLogicServer())
	if err != nil {
		fmt.Println(err)
	}
}
