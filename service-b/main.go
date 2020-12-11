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

type B struct {
	pb.UnimplementedBServiceServer
}

//echo
func (B) EchoB(c context.Context, request *pb.BMsg) (*pb.BMsg, error) {
	return &pb.BMsg{Value: "B replay:"+request.Value}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":8002")
	if err != nil {
		log.Fatalf("listen error:%s", err)
	}
	defer listen.Close()

	serverB := grpc.NewServer()
	pb.RegisterBServiceServer(serverB,&B{})
	log.Println("start b service.....")
	if err := serverB.Serve(listen); err != nil {
		log.Fatalf("server error:%s", err)
	}
}
