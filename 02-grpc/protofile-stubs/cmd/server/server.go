package main

import (
	"log"
	"net"

	"github.com/luankosaka1/golang-protofile-stubs/pb"
	"github.com/luankosaka1/golang-protofile-stubs/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	list, err := net.Listen("tcp", "localhost:50051")

	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, services.NewUserService())
	reflection.Register(grpcServer) // public method to server client

	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("Could not server: %v", err)
	}
}
