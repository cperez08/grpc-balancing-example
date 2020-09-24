package posthttp

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/Pallinder/go-randomdata"
	"github.com/cperez08/grcp-svc/internal/domain/post"
)

// PostHandler ...
type PostHandler struct {
	Cli post.PostServiceClient
}

// NewPostHandler ...
func NewPostHandler(cli post.PostServiceClient) *PostHandler {
	return &PostHandler{Cli: cli}
}

// CreatePost ...
func (h *PostHandler) CreatePost(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	var p *post.Post
	if strings.Contains(path, "/ok") {
		p = &post.Post{
			Author:      randomdata.SillyName(),
			Description: randomdata.RandStringRunes(15),
			Tags:        []string{"grpc", "example", "meetup"},
			Tittle:      randomdata.RandStringRunes(10),
		}
	} else {
		p = &post.Post{
			Description: randomdata.RandStringRunes(15),
			Tags:        []string{"grpc", "example", "meetup"},
			Tittle:      randomdata.RandStringRunes(10),
		}
	}

	rs, err := h.Cli.CreatePost(context.Background(), p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`"message": "%s"`, err.Error())))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf(`{"message": "post %d created", "created_by": "%s"}`, rs.PostId, rs.Author)))
	return
}
