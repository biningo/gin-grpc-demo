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

type C struct {
	pb.UnimplementedCServiceServer
}

//echo
func (C) EchoC(c context.Context, request *pb.CMsg) (*pb.CMsg, error) {
	return &pb.CMsg{Value: "C replay:"+request.Value}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":8003")
	if err != nil {
		log.Fatalf("listen error:%s", err)
	}
	defer listen.Close()

	serverC := grpc.NewServer()
	pb.RegisterCServiceServer(serverC,&C{})
	log.Println("start c service.....")
	if err := serverC.Serve(listen); err != nil {
		log.Fatalf("server error:%s", err)
	}
}
