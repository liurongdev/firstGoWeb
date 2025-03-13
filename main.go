package main

import (
	"awesomeProject/global"
	"awesomeProject/route/user"
	"awesomeProject/tool"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	//start()
	test()
	var name string = "zhang"
	fmt.Println(name)
}

func start() {

	global.InitViper("settings.dev.yml", "./config")
	global.InitMysql(global.GetMysqlConfig())
	// 初始化 Gin 路由引擎（默认包含 Logger 和 Recovery 中间件）
	r := gin.Default()
	user.Registry(r)
	fmt.Println("start success")
	// 启动服务器（默认监听 0.0.0.0:8080）
	r.Run()

}

func test() {
	nums := [][]int{{5, 8}, {3, 9}, {5, 12}, {16, 5}}
	tool.CountGoodRectangles(nums)
}
