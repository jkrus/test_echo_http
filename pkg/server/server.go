package server

import (
	"context"
	"fmt"
	"github.com/gookit/validate"
	"github.com/labstack/echo/v4"
)

type Server struct {
	httpServer *echo.Echo
}

type customValidator struct {
	validator *validate.Validation
}

func (cv *customValidator) Validate(i interface{}) error {
	v := validate.Struct(i)
	v.Validate()
	if v.Errors.Empty() {
		return nil
	}
	return fmt.Errorf(v.Errors.Error())
}

func (s *Server) Run(adr string, handler *echo.Echo) error {
	s.httpServer = handler
	s.httpServer.Validator = &customValidator{}
	return s.httpServer.Start(adr)
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

/*s.httpServer.Server = &http.Server{
	Addr:           "localhost:" + port,
	Handler:        handler,
	MaxHeaderBytes: 1 << 20,
	ReadTimeout:    10 * time.Second,
	WriteTimeout:   10 * time.Second,
}*/
//	s.httpServer.Server.Handler = handler
