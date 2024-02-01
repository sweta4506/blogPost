package service

import (
	"blogPost/proto"
	"context"
	"log"
	"time"
)

func CallCreatePost(client proto.BlogServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := proto.Post{
		Title:   "first post",
		Content: "hey,are you there",
		Author:  "Sweta",
		Tags:    []string{"sweta,grpc,go"},
	}
	res, err := client.CreatePost(ctx, &req)
	if err != nil {
		log.Fatalf("Failed to created Post")
	}
	log.Printf("Create Post response%v", res)
}

func CallReadPost(client proto.BlogServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := proto.PostIDRequest{
		PostId: int32(2),
	}
	res, err := client.ReadPost(ctx, &req)
	if err != nil {
		log.Fatalf("Failed to Read Post ,err:%s ", err.Error())
	}
	log.Printf("Read Post response: %v", res)
}
func CallUpdatePost(client proto.BlogServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := proto.UpdatePostRequest{
		PostId:  int32(2),
		Title:   "Updated post",
		Content: "hey,are you there",
		Author:  "Sweta",
		Tags:    []string{"sweta,grpc,go, udpate tag, test"},
	}
	res, err := client.UpdatePost(ctx, &req)
	if err != nil {
		log.Fatalf("Failed to Update Post")
	}
	log.Printf("Update Post response: %v", res)
}
func CallDeletePost(client proto.BlogServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := proto.PostIDRequest{
		PostId: int32(2),
	}
	res, err := client.DeletePost(ctx, &req)
	if err != nil {
		log.Fatalf("Failed to Delete Post")
	}
	log.Printf("Delete Post response: %v", res)
}
