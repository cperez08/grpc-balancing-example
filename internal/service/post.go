package service

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	"github.com/cperez08/grcp-svc/internal/domain/post"
)

var host, _ = os.Hostname()

// postService ...
type postService struct {
	ID uint64
}

// NewPostService ...
func NewPostService() post.PostServiceServer {
	return &postService{ID: 0}
}

// CreatePost creates a new post
func (s *postService) CreatePost(ctx context.Context, in *post.Post) (*post.Post, error) {
	if in.Author == "" {
		log.Println("error creating post, author is empty")
		return nil, errors.New("author cannot be empty")
	}

	s.ID++
	in.PostId = s.ID
	in.CreatedAt = time.Now().UTC().String()
	log.Println("post created:", in, "in: ", host)
	return in, nil
}
