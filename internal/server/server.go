// internal/server/server.go
package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
	http   *http.Server
}

func NewServer(engine *gin.Engine) *Server {
	return &Server{
		engine: engine,
		http: &http.Server{
			Handler: engine,
		},
	}
}

func (s *Server) Start(port string) error {
	s.http.Addr = fmt.Sprintf(":%s", port)
	return s.http.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.http.Shutdown(ctx)
}
