package main

import (
	"context"
	"github.com/biningo/gin-grpc-demo/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

/**
*@Author lyer
*@Date 2020/12/11 22:11
*@Describe
**/

type A struct {
	pb.UnimplementedAServiceServer
}

//echo
func (A) EchoA(c context.Context, request *pb.AMsg) (*pb.AMsg, error) {
	return &pb.AMsg{Value: "A replay:"+request.Value}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":8001")
	if err != nil {
		log.Fatalf("listen error:%s", err)
	}
	defer listen.Close()

	serverA := grpc.NewServer()
	pb.RegisterAServiceServer(serverA,&A{})
	log.Println("start a service.....")
	if err := serverA.Serve(listen); err != nil {
		log.Fatalf("server error:%s", err)
	}
}
