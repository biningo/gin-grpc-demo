package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

/**
*@Author lyer
*@Date 2020/12/11 21:54
*@Describe
**/

func main() {
	r := gin.Default()
	r.GET("/a", handlerA)
	r.GET("/b", handlerB)
	r.GET("/c", handlerC)
	log.Println("start client...........")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("run error:%s", err)
	}
}
