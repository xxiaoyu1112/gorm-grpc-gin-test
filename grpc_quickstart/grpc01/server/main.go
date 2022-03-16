package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	hellogrpc2 "grpc_test/grpc01/pb"
	"net"
)

type server struct {
	hellogrpc2.UnimplementedHelloGRPCServer // 实现一个接口
}

// 在实现接口的基础上实现方法
func (s *server) SayHi(ctx context.Context, req *hellogrpc2.Req) (res *hellogrpc2.Res, err error) {
	fmt.Println(req.GetMessage())
	return &hellogrpc2.Res{Message: "我是从服务端返回的grpc的内容"}, nil
}

func main() {
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println(err)
	}
	s := grpc.NewServer()
	hellogrpc2.RegisterHelloGRPCServer(s, &server{})
	s.Serve(l)
}
