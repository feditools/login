package grpc

import (
	"context"
	"github.com/feditools/login/internal/config"
	"github.com/feditools/login/internal/metrics"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"net"
)

// Server is a http 2 web server
type Server struct {
	metrics metrics.Collector
	tcp     net.Listener
	srv     *grpc.Server
}

// NewServer creates a new grpc web server
func NewServer(_ context.Context, m metrics.Collector) (*Server, error) {
	server := &Server{
		metrics: m,
	}

	var err error
	server.tcp, err = net.Listen("tcp", viper.GetString(config.Keys.ServerGRPCBind))
	if err != nil {
		return nil, err
	}

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(ensureValidToken),
	}
	server.srv = grpc.NewServer(opts...)

	return server, nil
}

// Start starts the web server
func (s *Server) Start() error {
	l := logger.WithField("func", "Start")
	l.Infof("listening on %s", s.tcp.Addr())
	return s.srv.Serve(s.tcp)
}

// Stop shuts down the web server
func (s *Server) Stop() error {
	s.srv.Stop()
	return nil
}

// Server returns the server, used by modules to register themselves
func (s *Server) Server() *grpc.Server {
	return s.srv
}
