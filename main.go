package main

import (
	"github.com/howen/goimg/route"
	"github.com/howen/goimg/server"
	"github.com/howen/goimg/config"
)

func main() {

	// 初始化路由
	route.InitRoute()

	// 开始监听
	server.RunHttp(config.HttpAddr())
}
