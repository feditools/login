package metrics

import "time"

// Collector collects metrics from the feditools
type Collector interface {
	Close() error

	HTTPRequest(t time.Duration, status int, method, path string)
	NewDBQuery(name string) DBQuery
	NewDBCacheQuery(name string) DBCacheQuery
	NewGRPCRequest(method string) GRPCRequest
}

// DBQuery is a new database query metric measurer
type DBQuery interface {
	Done(isError bool)
}

// DBCacheQuery is a new database cache query metric measurer
type DBCacheQuery interface {
	Done(hit bool, isError bool)
}

// GRPCRequest is a new grpc request metric measurer
type GRPCRequest interface {
	Done(isError bool)
}
