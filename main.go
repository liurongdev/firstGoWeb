package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/liurongdev/firstGoWeb/global"
	"github.com/liurongdev/firstGoWeb/middleware/logger"
	"github.com/liurongdev/firstGoWeb/middleware/redis"
	"github.com/liurongdev/firstGoWeb/route/user"
	"github.com/liurongdev/firstGoWeb/tool"
)

func main() {
	start()
	//test()
}

func start() {

	name := flag.String("name", "wang", "用户名称")
	fmt.Println(name)
	flag.Parse()
	global.InitViper("settings.dev.yml", "./config")
	global.InitMysql(global.GetMysqlConfig())
	redis.InitRedis(global.GetRedisConfig())
	logger.Init()
	// 初始化 Gin 路由引擎（默认包含 Logger 和 Recovery 中间件）
	r := gin.Default()
	user.Registry(r)
	fmt.Println("start success")
	// 启动服务器（默认监听 0.0.0.0:8080）
	r.Run()

}

func test() {
	//nums := [][]int{{5, 8}, {3, 9}, {5, 12}, {16, 5}}
	tool.Partition("aab")
}
