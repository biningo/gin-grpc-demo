package main
import (
	"context"
	"github.com/biningo/gin-grpc-demo/pb"
	"github.com/gin-gonic/gin"
	"log"
)

/**
*@Author lyer
*@Date 2020/12/11 22:29
*@Describe
**/

func handlerA(ctx *gin.Context) {
	msg := ctx.Query("msg")
	msgBody := &pb.AMsg{Value: msg}
	client:=NewClientA()
	respBody, err := client.EchoA(context.Background(), msgBody)
	if err != nil {
		log.Fatalf("echoA error:%s", err)
	}
	ctx.JSON(200, gin.H{
		"data": respBody,
	})
}

func handlerB(ctx *gin.Context) {
	msg := ctx.Query("msg")
	msgBody := &pb.BMsg{Value: msg}
	client := NewClientB()
	respBody, err := client.EchoB(context.Background(), msgBody)
	if err != nil {
		log.Fatalf("echoB error:%s", err)
	}
	ctx.JSON(200, gin.H{
		"data": respBody,
	})
}

func handlerC(ctx *gin.Context) {
	msg := ctx.Query("msg")
	msgBody := &pb.CMsg{Value: msg}
	client := NewClientC()
	respBody, err := client.EchoC(context.Background(), msgBody)
	if err != nil {
		log.Fatalf("echoC error:%s", err)
	}
	ctx.JSON(200, gin.H{
		"data": respBody,
	})
}