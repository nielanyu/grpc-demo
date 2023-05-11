package main

import (
	"grpc-demo/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	server := grpc.NewServer()
	service.RegisterGetUserinfoServiceServer(server, service.UserService{})

	listener, err := net.Listen("tcp", ":8002")
	if err != nil {
		log.Fatal("服务监听端口失败", err)
	}
	err = server.Serve(listener)
	if err != nil {
		log.Fatal("服务挂了", err)
	}
}
