package statsd

import (
	"github.com/cactus/go-statsd-client/v5/statsd"
	"github.com/feditools/login/internal/metrics"
	"strconv"
	"time"
)

// HTTPRequest send a metrics relating to a http request
func (m *Module) HTTPRequest(t time.Duration, status int, method, path string) {
	err := m.s.TimingDuration(
		metrics.StatHTTPRequest,
		t,
		m.rate,
		statsd.Tag{"status", strconv.Itoa(status)},
		statsd.Tag{"method", method},
		statsd.Tag{"path", path},
	)
	if err != nil {
		logger.WithField("func", "HTTPRequest").Warn(err.Error())
	}
}
