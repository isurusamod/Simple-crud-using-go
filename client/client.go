package main

import (
	"context"
	"fmt"
	"log"

	pb "grpc-user-crud/proto" 
	"google.golang.org/grpc"
)

func main() {
	
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewUserServiceClient(conn)

	
	user := &pb.User{
		Id:    "",
		Name:  "",
		Email: "",
	}
	
	resp, err := c.CreateUser(context.Background(), user)
	if err != nil {
		log.Fatalf("Could not create user: %v", err)
	}

	fmt.Printf("User created: %v\n", resp)
}
