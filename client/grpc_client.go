package main

import (
	"github.com/biningo/gin-grpc-demo/pb"
	"google.golang.org/grpc"
	"log"
)

/**
*@Author lyer
*@Date 2020/12/11 22:29
*@Describe
**/

func NewClientA() pb.AServiceClient{
	conn,err:=grpc.Dial("localhost:8001",grpc.WithInsecure(),grpc.WithBlock())
	if err!=nil{
		log.Fatalf("did not connect: %v", err)
	}
	client:=pb.NewAServiceClient(conn)
	return client
}


func NewClientB() pb.BServiceClient{
	conn,err:=grpc.Dial("localhost:8002",grpc.WithInsecure(),grpc.WithBlock())
	if err!=nil{
		log.Fatalf("did not connect: %v", err)
	}
	client:=pb.NewBServiceClient(conn)
	return client
}


func NewClientC() pb.CServiceClient{
	conn,err:=grpc.Dial("localhost:8003",grpc.WithInsecure(),grpc.WithBlock())
	if err!=nil{
		log.Fatalf("did not connect: %v", err)
	}
	client:=pb.NewCServiceClient(conn)
	return client
}