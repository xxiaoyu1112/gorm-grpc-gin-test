package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	hello_grpc2 "grpc_test/grpc01/pb"
)

func main() {
	conn, err := grpc.Dial("localhost:8888", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	client := hello_grpc2.NewHelloGRPCClient(conn)
	req, _ := client.SayHi(context.Background(), &hello_grpc2.Req{Message: "我从客户端来"})
	fmt.Println(req.GetMessage())
}
