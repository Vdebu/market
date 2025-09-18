package grpc

import (
	"fmt"
	"github.com/vdebu/market-proto/golang/order"
	"github.com/vdebu/market/order/internal/ports"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type Adapter struct {
	api                            ports.APIPort // 持有业务接口实例，对接核心业务逻辑
	port                           int           // gRPC服务监听的端口号
	order.UnimplementedOrderServer               // 嵌入自动生成的空实现，实现接口向前兼容
}

// NewAdapter 创建适配器实例
func NewAdapter(api ports.APIPort, port int) *Adapter {
	return &Adapter{
		api:  api,
		port: port,
	}
}

// Run 启动RPC服务器
func (a Adapter) Run() {
	var err error
	// 监听当前适配器中存储的端口
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen on port: %d, error: %v", a.port, err)
	}
	// 创建空的微服务器实例
	grpcServer := grpc.NewServer()
	// 向服务器上挂载服务(使用内置生成方法)
	// 1:挂载到的服务器实例，2:实际的服务适配器
	order.RegisterOrderServer(grpcServer, a)
	// 根据开发的情况选择是否开启反射
	if config.GetEnv() == "develop" {
		reflection.Register(grpcServer)
	}
	// 开始监听指定的端口
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve grpc on port: %d", a.port)
	}
}
