package main

import (
	"log"
	"os"

	"github.com/cperez08/grcp-svc/internal/client"
	"github.com/cperez08/grcp-svc/internal/domain/post"

	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	var serverAddress = os.Getenv("POST_SVC_ADDRESS")

	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}

	log.Println("grpc client connected with the server")
	cli := post.NewPostServiceClient(conn)
	client.StartClient(cli, conn)
}
