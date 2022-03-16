package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc_test/grpc03/pb/person"
	"sync"
	"time"
)

func main() {
	l, err := grpc.Dial("127.0.0.1:8888", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	client := person.NewSearchServiceClient(l)

	// 普通传入
	//res, err := client.Search(context.Background(), &person.PersonReq{Name: "客户端"})
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(res)

	// 流式传入
	//c, err := client.SearchIn(context.Background())
	//if err != nil {
	//	fmt.Println(err)
	//}
	//i := 0
	//for {
	//	if i > 10 {
	//		res, err := c.CloseAndRecv()
	//		if err != nil {
	//			fmt.Println(err)
	//		}
	//		fmt.Println(res)
	//		break
	//	}
	//	time.Sleep(1 * time.Second)
	//	c.Send(&person.PersonReq{Name: "客户端的信息"})
	//	i++
	//}

	// 流式返回
	//c, err := client.SearchOut(context.Background(), &person.PersonReq{Name: "客户端"})
	//if err != nil {
	//	fmt.Println(err)
	//}
	//for {
	//	req, err := c.Recv()
	//	if err != nil {
	//		fmt.Println(err)
	//		break
	//	}
	//	fmt.Println(req)
	//}

	// 流式出入
	c, err := client.SearchIO(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for {
			time.Sleep(1 * time.Second)
			err := c.Send(&person.PersonReq{Name: "客户端"})
			if err != nil {
				wg.Done()
				break
			}
		}
	}()
	go func() {
		for {
			req, err := c.Recv()
			if err != nil {
				fmt.Println(err)
				wg.Done()
				break
			}
			fmt.Println(req)
		}
	}()
	wg.Wait()
}
