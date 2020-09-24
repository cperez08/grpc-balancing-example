package client

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/cperez08/grcp-svc/internal/domain/post"
	posthttp "github.com/cperez08/grcp-svc/internal/http"
	"google.golang.org/grpc"
)

// StartClient starts a grpc and http client
func StartClient(cli post.PostServiceClient, conn *grpc.ClientConn) {
	// creates a helper http server as request entrypoint
	handler := posthttp.NewPostHandler(cli)
	http.HandleFunc("/", handler.CreatePost)
	log.Println("http server running at 8080 port")
	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			panic(err)
		}
	}()

	closer := make(chan os.Signal, 1)
	signal.Notify(closer, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	<-closer
	if err := conn.Close(); err != nil {
		log.Println("error closing grpc connection", err)
		return
	}

	log.Println("gracefully shutting down grpc/http client")
}
