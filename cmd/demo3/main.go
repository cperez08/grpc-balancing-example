package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	dmresolver "github.com/cperez08/dm-resolver/pkg/resolver"
	"github.com/cperez08/grcp-svc/internal/client"
	"github.com/cperez08/grcp-svc/internal/domain/post"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/resolver"

	"google.golang.org/grpc"
)

var (
	scheme        = "demo"
	serverAddress = os.Getenv("POST_SVC_ADDRESS")
	refreshRate   = time.Duration(15)
)

func init() {
	addr := strings.Split(serverAddress, ":")
	if len(addr) < 2 {
		panic("invalid service address make sure the format is host:port")
	}

	resolver.Register(dmresolver.NewDomainResolverBuilder(scheme, addr[0], addr[1], true, &refreshRate))
}

func main() {
	var conn *grpc.ClientConn
	target := fmt.Sprintf("%s:///%s", scheme, serverAddress)
	conn, err := grpc.Dial(target, grpc.WithInsecure(), grpc.WithBalancerName(roundrobin.Name))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}

	log.Println("grpc client connected with the server")
	cli := post.NewPostServiceClient(conn)
	client.StartClient(cli, conn)
}
