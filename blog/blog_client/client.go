package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/grpc_blog_api/blog/blogpb"

	"google.golang.org/grpc"
)

func main() {

	// Create a connection
	opts := grpc.WithInsecure()
	conn, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("Could not connect %v", err)
	}
	defer conn.Close()

	client := blogpb.NewBlogServiceClient(conn)
	fmt.Printf("Client created: %v \n", client)

	// Create Blog
	log.Println("Creating blog")
	blog := &blogpb.Blog{
		AuthorId: "Jaiesh",
		Title:    "GRPC Blog",
		Content:  "GRPC is simple",
	}
	createBlogResp, err := client.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: blog})
	if err != nil {
		log.Fatalf("Error while creating blog %v", err)
	}
	log.Printf("Blog has been created: %v", createBlogResp)

	// Read Blog - Success
	log.Println("Reading the blog")
	readBlogResponse, err := client.ReadBLog(context.Background(), &blogpb.ReadBlogRequest{BlogId: createBlogResp.GetBlog().GetId()})
	if err != nil {
		log.Fatalf("Error while reading blog %v", err)
	}
	log.Printf("Blog was read: %v", readBlogResponse)

	// Update Blog
	log.Println("Updating the blog")
	blog.AuthorId = "Timo"
	blog.Title = "Edited Title"
	blog.Id = createBlogResp.GetBlog().GetId()
	updateRes, err := client.UpdateBlog(context.Background(), &blogpb.UpdateBlogRequest{Blog: blog})
	if err != nil {
		log.Fatalf("Error while updating blog %v", err)
	}
	log.Printf("Blog was update: %v", updateRes)

	// Delete Blog - Success
	log.Println("Deleting the blog")
	deleteBlogRes, err := client.DeleteBlog(context.Background(), &blogpb.DeleteBlogRequest{BlogId: createBlogResp.GetBlog().GetId()})
	if err != nil {
		log.Fatalf("Error while deleting blog %v", err)
	}
	log.Printf("Blog was deleted: %v", deleteBlogRes)

	// List blogs
	log.Println("Listing all blogs")
	stream, err := client.ListBlog(context.Background(), &blogpb.ListBlogRequest{})
	if err != nil {
		log.Fatalf("Error while calling ListBlog RPC %v", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Something happened: %v", err)
		}
		fmt.Println(res.GetBlog())
	}
}
