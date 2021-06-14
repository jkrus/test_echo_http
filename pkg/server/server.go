package server

import (
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type Server struct {
	httpServer *echo.Echo
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = echo.New()
	s.httpServer.Server = &http.Server{
		Addr:           "localhost:" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return s.httpServer.Server.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
