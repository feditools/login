package grpc

// Module represents a module that can be added to a http server
type Module interface {
	Name() string
	Register(s *Server) error
}
