// server/main.go
package main

import (
	"blogPost/dtos"
	"blogPost/proto"
	"blogPost/server/daos"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":3001"
)

type blogServer struct {
	proto.BlogServiceServer // Correct way to embed the interface
}

func main() {

	// //Create a new Database instance
	// db, err := daos.NewDatabase()
	// if err != nil {
	// 	log.Fatalf("Failed to create database: %v", err)
	// }

	// // //Create a new BlogServiceServer instance
	// blogService := service.NewBlogServiceServer(db, &blogServer{})
	// log.Printf("blogService: %v", blogService)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start the server: %s", err.Error())
	}

	server := grpc.NewServer()

	// Register your gRPC server implementation with the gRPC server
	proto.RegisterBlogServiceServer(server, &blogServer{}) // Pass a value of type blogServer, not a pointer

	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *blogServer) CreatePost(ctx context.Context, req *proto.Post) (*proto.Post, error) {
	daoReq := dtos.Post{
		PostID:  req.GetPostId(),
		Title:   req.GetTitle(),
		Author:  req.GetAuthor(),
		Content: req.GetContent(),
		Tags:    req.GetTags(),
	}
	db, err := daos.NewDatabase()
	if err != nil {
		log.Fatalf("Failed to create database: %v", err)
	}
	postID, err := db.CreatePost(daoReq)
	if err != nil {
		return nil, fmt.Errorf("failed to create post: %v", err)
	}

	return &proto.Post{
		PostId:  postID,
		Title:   req.GetTitle(),
		Author:  req.GetAuthor(),
		Content: req.GetContent(),
		Tags:    req.GetTags(),
	}, nil
}
func (s *blogServer) ReadPost(ctx context.Context, req *proto.PostIDRequest) (*proto.Post, error) {
	db, err := daos.NewDatabase()
	if err != nil {
		log.Fatalf("Failed to create database: %v", err)
	}
	post, err := db.ReadPost(req.GetPostId())
	if err != nil {
		return nil, fmt.Errorf("failed to read post: %v", err)
	}

	return &proto.Post{
		PostId:          post.PostID,
		Title:           post.Title,
		Content:         post.Content,
		Author:          post.Author,
		PublicationDate: post.PublicationDate,
		Tags:            post.Tags,
	}, nil
}

func (s *blogServer) UpdatePost(ctx context.Context, req *proto.UpdatePostRequest) (*proto.Post, error) {
	daoReq := dtos.Post{
		PostID:  req.GetPostId(),
		Title:   req.GetTitle(),
		Author:  req.GetAuthor(),
		Content: req.GetContent(),
		Tags:    req.GetTags(),
	}
	if daoReq.PostID == 0 {
		return nil, fmt.Errorf("PostID is not valid")
	}
	db, err := daos.NewDatabase()
	if err != nil {
		log.Fatalf("Failed to create database: %v", err)
	}

	err = db.UpdatePost(&daoReq)
	if err != nil {
		return nil, fmt.Errorf("failed to update post: %v", err)
	}

	// Retrieve the updated post
	updatedPost, err := db.ReadPost(daoReq.PostID)
	if err != nil {
		return nil, fmt.Errorf("failed to read updated post: %v", err)
	}

	return &proto.Post{
		PostId:          updatedPost.PostID,
		Title:           updatedPost.Title,
		Content:         updatedPost.Content,
		Author:          updatedPost.Author,
		PublicationDate: updatedPost.PublicationDate,
		Tags:            updatedPost.Tags,
	}, nil
}

func (s *blogServer) DeletePost(ctx context.Context, req *proto.PostIDRequest) (*proto.DeleteResponse, error) {
	db, err := daos.NewDatabase()
	if err != nil {
		log.Fatalf("Failed to create database: %v", err)
	}
	err = db.DeletePost(req.GetPostId())
	if err != nil {
		return &proto.DeleteResponse{Success: false}, fmt.Errorf("failed to delete post: %v", err)
	}

	return &proto.DeleteResponse{Success: true}, nil
}
