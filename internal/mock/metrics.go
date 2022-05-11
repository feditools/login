package mock

import (
	"github.com/feditools/login/internal/metrics"
	"time"
)

// GRPCRequest is a new database query metric measurer
type GRPCRequest struct{}

// Done is called when the grpc request is complete
func (GRPCRequest) Done(isError bool) {
	return
}

// DBQuery is a new database query metric measurer
type DBQuery struct{}

// Done is called when the db query is complete
func (DBQuery) Done(isError bool) {
	return
}

// DBCacheQuery is a new database cache query metric measurer
type DBCacheQuery struct{}

// Done is called when the db cache query is complete
func (DBCacheQuery) Done(hit, isError bool) {
	return
}

// MetricsCollector is a mock metrics collection
type MetricsCollector struct{}

// NewGRPCRequest creates a new grpc metrics collector
func (c MetricsCollector) NewGRPCRequest(method string) metrics.GRPCRequest {
	return &GRPCRequest{}
}

// Close does nothing
func (MetricsCollector) Close() error {
	return nil
}

// DBQuery does nothing
func (MetricsCollector) DBQuery(t time.Duration, name string, error bool) {
	return
}

// NewDBQuery creates a new db query metrics collector
func (MetricsCollector) NewDBQuery(name string) metrics.DBQuery {
	return &DBQuery{}
}

// NewDBCacheQuery creates a new db cache query metrics collector
func (c MetricsCollector) NewDBCacheQuery(name string) metrics.DBCacheQuery {
	return &DBCacheQuery{}
}

// HTTPRequest does nothing
func (MetricsCollector) HTTPRequest(t time.Duration, status int, method, path string) {
	return
}

// NewMetricsCollector creates a new mock metrics collector
func NewMetricsCollector() (metrics.Collector, error) {
	return &MetricsCollector{}, nil
}
