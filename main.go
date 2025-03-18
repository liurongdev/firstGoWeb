package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/liurongdev/firstGoWeb/global"
	"github.com/liurongdev/firstGoWeb/grpc/server"
	"github.com/liurongdev/firstGoWeb/middleware/logger"
	"github.com/liurongdev/firstGoWeb/middleware/redis"
	"github.com/liurongdev/firstGoWeb/route/user"
	"github.com/liurongdev/firstGoWeb/tool"
	"github.com/soheilhy/cmux"
	"net"
	"net/http"
)

func main() {
	start()
	//test()
}

func start() {
	port := global.Viper.GetString("settings.application.port")
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		logger.Error(err.Error())
	}

	mux := cmux.New(listen)
	// 匹配 gRPC 流量（基于 HTTP/2）
	grpcListener := mux.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
	// 匹配 HTTP 流量
	httpListener := mux.Match(cmux.Any())

	go startGinServer(httpListener)
	go startGinServer(grpcListener)
}

func startGinServer(httpListen net.Listener) {
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

	http.Serve(httpListen, r)

	//r.Run(fmt.Sprintf(":%s", port))

}

func startGrpcServer(grpcListen net.Listener) {
	server.StartGRPC(grpcListen)

}

func test() {

	//nums := [][]int{{5, 8}, {3, 9}, {5, 12}, {16, 5}}
	var node5 *tool.TreeNode = &tool.TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}
	var node3 *tool.TreeNode = &tool.TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}

	var node6 *tool.TreeNode = &tool.TreeNode{
		Val:   6,
		Left:  nil,
		Right: nil,
	}

	var node2 *tool.TreeNode = &tool.TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}

	var node4 *tool.TreeNode = &tool.TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}

	var node7 *tool.TreeNode = &tool.TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}

	node5.Left = node3
	node5.Right = node6

	node3.Left = node2
	node3.Right = node4

	node6.Right = node7

	tool.DeleteNode(node5, 3)
	fmt.Println("%v:", node5)
}
