package metrics

import "time"

// Collector collects metrics from the feditools
type Collector interface {
	Close() error

	HTTPRequestTiming(t time.Duration, status int, method, path string)
}
