package main

import (
	"blogPost/client/service"
	"blogPost/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":3001"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Fialed to connect, err:%s", err.Error())
	}
	defer conn.Close()
	client := proto.NewBlogServiceClient(conn)
	service.CallCreatePost(client)
	service.CallReadPost(client)
	service.CallUpdatePost(client)
	service.CallDeletePost(client)
}
