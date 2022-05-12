package grpc

import (
	"context"
)

func (s *Server) authValid(ctx context.Context, authorization []string) bool {
	if len(authorization) < 1 {
		return false
	}

	applicationToken, err := s.db.ReadApplicationTokenByToken(ctx, authorization[0])
	if err != nil {
		logger.WithField("func", "authValid").Errorf("db read: %s", err.Error())
		return false
	}
	if applicationToken == nil {
		return false
	}

	return true
}
