package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/liurongdev/firstGoWeb/global"
	pb "github.com/liurongdev/firstGoWeb/grpc/proto"
	"github.com/liurongdev/firstGoWeb/grpc/server"
	"github.com/liurongdev/firstGoWeb/middleware/logger"
	"github.com/liurongdev/firstGoWeb/middleware/redis"
	"github.com/liurongdev/firstGoWeb/route/user"
	"github.com/liurongdev/firstGoWeb/tool"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	"os"
	"sync"
)

func main() {
	fmt.Println("main function call...")
	start()
	//test()
}

func init() {
	fmt.Println("init function call...")
	global.InitViper("settings.dev.yml", "./config")
	global.InitMysql(global.GetMysqlConfig())
	redis.InitRedis(global.GetRedisConfig())
	logger.Init()
	grpclog.SetLoggerV2(grpclog.NewLoggerV2(os.Stdout, os.Stderr, os.Stderr))
}

func start() {
	port := flag.String("port", "8081", "端口")
	flag.Parse()
	if *port == "" {
		*port = global.Viper.GetString("settings.application.port")
	}
	fmt.Println("port:", *port)
	listen, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%s", *port))
	if err != nil {
		logger.Error(err.Error())
	}
	defer listen.Close()
	mux := cmux.New(listen)
	// 匹配 gRPC 流量（基于 HTTP/2）
	//严格显示客户端为http2格式
	grpcListener := mux.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	//grpcListener := mux.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
	httpListener := mux.Match(cmux.Any())
	// 匹配 HTTP 流量
	var wait sync.WaitGroup
	wait.Add(1)

	go startGrpcServer(grpcListener)
	go startGinServer(httpListener)
	// 启动 cmux
	logger.Info("Starting cmux on ", port)
	if err := mux.Serve(); err != nil {
		log.Fatalf("Failed to serve cmux: %v", err)
	}
	wait.Wait()
}

func startGrpcServer(listener net.Listener) {
	logger.Info("Starting ginServer on ")
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pb.RegisterHelloServiceServer(grpcServer, &server.HelloServiceServer{})
	if err := grpcServer.Serve(listener); err != nil {
		logger.Error(err.Error())
	}
}

func startGinServer(listener net.Listener) {
	// 初始化 Gin 路由引擎（默认包含 Logger 和 Recovery 中间件）
	r := gin.Default()
	user.Registry(r)
	httpServer := &http.Server{
		Handler: r,
	}
	logger.Info("Starting ginServer on ")
	if err := httpServer.Serve(listener); err != nil {
		logger.Error(err.Error())
	}
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
