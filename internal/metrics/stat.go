package metrics

const (
	// StatBaseDB is the base state name for database metrics
	StatBaseDB = "db"
	// StatDBQuery is the timing for a db query
	StatDBQuery = StatBaseDB + ".query"
	// StatDBQueryTiming is the timing for a db query
	StatDBQueryTiming = StatDBQuery + ".timing"
	// StatDBQueryCount is the counts for db queries
	StatDBQueryCount = StatDBQuery + ".count"

	// StatBaseDBCache is the base state name for database cache metrics
	StatBaseDBCache = "dbcache"
	// StatDBCacheQuery is the timing for a db query
	StatDBCacheQuery = StatBaseDBCache + ".query"
	// StatDBCacheQueryTiming is the timing for a db query
	StatDBCacheQueryTiming = StatDBCacheQuery + ".timing"
	// StatDBCacheQueryCount is the counts for db queries
	StatDBCacheQueryCount = StatDBCacheQuery + ".count"

	// StatBaseGRPC is the base state name for grpc metrics
	StatBaseGRPC = "grpc"
	// StatGRPCRequest is the timing for a grpc request
	StatGRPCRequest = StatBaseGRPC + ".request"
	// StatGRPCRequestTiming is the timing for a db query
	StatGRPCRequestTiming = StatGRPCRequest + ".timing"
	// StatGRPCRequestCount is the counts for db queries
	StatGRPCRequestCount = StatGRPCRequest + ".count"

	// StatBaseHTTP is the base state name for http metrics
	StatBaseHTTP = "http"
	// StatHTTPRequest is the timing for a http request
	StatHTTPRequest = StatBaseHTTP + ".request"

	// StatBaseSys is the base state name for system metrics
	StatBaseSys = "sys"
	// StatSysMem is the base state name for system memory metrics
	StatSysMem = StatBaseSys + ".mem"
	// StatSysMemAlloc is the gauge for memory allocation
	StatSysMemAlloc = StatSysMem + ".alloc"
	// StatSysMemAllocTotal is the gauge for memory total allocation
	StatSysMemAllocTotal = StatSysMem + ".alloc-total"
	// StatSysMemSys is the gauge for system memory usage
	StatSysMemSys = StatSysMem + ".sys"
	// StatSysMemNumGC is the counter for
	StatSysMemNumGC = StatSysMem + ".num-gc"
	// StatSysRoutines is the gauge for the number of active go routines
	StatSysRoutines = StatBaseSys + ".goroutines"
)
