package httpserver

import (
	"context"
	"net/http"
	"time"

	"github.com/thanhpp/aws-cognito-example/infra/cognitolib"
)

type Server struct {
	host       string
	port       string
	cognitoApp *cognitolib.CognitoApp
	httpServer *http.Server
}

func NewServer(host, port string, cognitoApp *cognitolib.CognitoApp) *Server {
	return &Server{
		host:       host,
		port:       port,
		cognitoApp: cognitoApp,
	}
}

func (s *Server) Start() error {
	server := &http.Server{
		Addr:    s.host + ":" + s.port,
		Handler: newRouter(s.cognitoApp),
	}
	s.httpServer = server

	return server.ListenAndServe()
}

func (s *Server) Stop() error {
	if s.httpServer == nil {
		return nil
	}

	shutdownCtx, cancel := context.WithTimeout(
		context.Background(),
		time.Second*5,
	)
	defer cancel()

	return s.httpServer.Shutdown(shutdownCtx)
}
