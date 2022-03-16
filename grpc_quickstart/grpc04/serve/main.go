package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc_test/grpc03/pb/person"
	"net"
	"time"
)

type personServe struct {
	person.UnimplementedSearchServiceServer
}

func (*personServe) Search(ctx context.Context, req *person.PersonReq) (*person.PersonRes, error) {
	name := req.GetName()
	res := &person.PersonRes{Name: "服务端收到了来自" + name + "的信息"}
	return res, nil
}
func (*personServe) SearchIn(server person.SearchService_SearchIntServer) error {
	for {
		req, err := server.Recv()
		fmt.Println(req)
		if err != nil {
			server.SendAndClose(&person.PersonRes{Name: "完美"})
			break
		}
	}
	return nil
}
func (*personServe) SearchOut(req *person.PersonReq, server person.SearchService_SearchOutServer) error {
	name := req.Name
	i := 0
	for {
		if i > 10 {
			break
		}
		time.Sleep(1 * time.Second)
		i++
		server.Send(&person.PersonRes{Name: "服务端拿到了" + name})
	}
	return nil
}
func (*personServe) SearchIO(server person.SearchService_SearchIOServer) error {
	i := 0
	str := make(chan string)
	go func() {
		for {
			i++
			req, _ := server.Recv()
			if i > 10 {
				str <- " 结束 "
				break
			}
			str <- req.Name
		}
	}()
	for {
		s := <-str
		if s == "结束" {
			break
		}
		err := server.Send(&person.PersonRes{Name: s})
		if err != nil {
			fmt.Println(err)
		}
	}
	return nil
}

func main() {
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println(err)
	}
	s := grpc.NewServer()
	person.RegisterSearchServiceServer(s, &personServe{})
	s.Serve(l)
}
