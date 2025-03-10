package main

import (
	"awesomeProject/global"
	"awesomeProject/route/user"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	global.InitMysql()
	fmt.Println("init mysql success")
	// 初始化 Gin 路由引擎（默认包含 Logger 和 Recovery 中间件）
	r := gin.Default()
	user.Registry(r)
	// 启动服务器（默认监听 0.0.0.0:8080）
	r.Run()
}
