package main

import (
	"fmt"
	"log"
	"os"

	"github.com/cperez08/grcp-svc/internal/client"
	"github.com/cperez08/grcp-svc/internal/domain/post"

	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
)

func main() {
	var conn *grpc.ClientConn
	var serverAddress = os.Getenv("POST_SVC_ADDRESS")

	target := fmt.Sprintf("dns:///%s", serverAddress)

	conn, err := grpc.Dial(target, grpc.WithInsecure(), grpc.WithBalancerName(roundrobin.Name))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}

	log.Println("grpc client connected with the server")
	cli := post.NewPostServiceClient(conn)
	client.StartClient(cli, conn)
}
