package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "grpc-user-crud/proto" 
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUserServiceServer
	users map[string]*pb.User
}


func (s *server) CreateUser(ctx context.Context, req *pb.User) (*pb.UserResponse, error) {
	s.users[req.Id] = req
	log.Println("User Created:", req)
	return &pb.UserResponse{Id: req.Id, Name: req.Name, Email: req.Email}, nil
}


func (s *server) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	user, exists := s.users[req.Id]
	if !exists {
		return nil, fmt.Errorf("User with ID %s not found", req.Id)
	}
	return &pb.UserResponse{Id: user.Id, Name: user.Name, Email: user.Email}, nil
}


func (s *server) UpdateUser(ctx context.Context, req *pb.User) (*pb.UserResponse, error) {
	user, exists := s.users[req.Id]
	if !exists {
		return nil, fmt.Errorf("User with ID %s not found", req.Id)
	}


	user.Name = req.Name
	user.Email = req.Email
	s.users[req.Id] = user

	log.Println("User Updated:", user)
	return &pb.UserResponse{Id: user.Id, Name: user.Name, Email: user.Email}, nil
}


func (s *server) DeleteUser(ctx context.Context, req *pb.UserRequest) (*pb.DeleteResponse, error) {
	_, exists := s.users[req.Id]
	if !exists {
		return nil, fmt.Errorf("User ID %s not found", req.Id)
	}

	delete(s.users, req.Id)
	log.Printf("User ID %s deleted", req.Id)

	return &pb.DeleteResponse{Message: fmt.Sprintf("User ID %s successfully deleted", req.Id)}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{users: make(map[string]*pb.User)})
	log.Println("gRPC Server running on port 50052")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
