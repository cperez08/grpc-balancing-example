package main

import (
	"log"
	"net"

	"github.com/cperez08/grcp-svc/internal/domain/post"
	"github.com/cperez08/grcp-svc/internal/service"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	post.RegisterPostServiceServer(grpcServer, service.NewPostService())
	log.Println("post gRPC server running at port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
