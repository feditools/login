package statsd

import (
	"fmt"
	"github.com/cactus/go-statsd-client/v5/statsd"
	"github.com/feditools/login/internal/metrics"
	"time"
)

// GRPCRequest is a new database cache query metric measurer
type GRPCRequest struct {
	s    statsd.Statter
	rate float32

	method string
	start  time.Time
}

// NewGRPCRequest creates a new db query metrics collector
func (m *Module) NewGRPCRequest(method string) metrics.GRPCRequest {
	return &GRPCRequest{
		s:    m.s,
		rate: m.rate,

		method: method,
		start:  time.Now(),
	}
}

// Done is called when the db query is complete
func (g *GRPCRequest) Done(isError bool) {
	l := logger.WithField("type", "GRPCRequest").WithField("func", "Done")

	t := time.Since(g.start)

	err := g.s.TimingDuration(
		metrics.StatGRPCRequestTiming,
		t,
		g.rate,
		statsd.Tag{"method", g.method},
		statsd.Tag{"error", fmt.Sprintf("%v", isError)},
	)
	if err != nil {
		l.WithField("type", "timing").Warn(err.Error())
	}

	err = g.s.Inc(
		metrics.StatGRPCRequestCount,
		1,
		g.rate,
		statsd.Tag{"method", g.method},
		statsd.Tag{"error", fmt.Sprintf("%v", isError)},
	)
	if err != nil {
		l.WithField("type", "count").Warn(err.Error())
	}
}
