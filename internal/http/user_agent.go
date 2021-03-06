package http

import (
	"fmt"
	"sync"

	"github.com/feditools/login/internal/config"
	"github.com/spf13/viper"
)

var (
	initOnce      sync.Once
	userAgentLock sync.RWMutex

	userAgent string
)

// doInit sets the User-Agent for all subsequent requests.
func doInit() {
	userAgentLock.Lock()
	userAgent = fmt.Sprintf("Go-http-client/2.0 (%s/%s; +%s/)",
		viper.GetString(config.Keys.ApplicationName),
		viper.GetString(config.Keys.SoftwareVersion),
		viper.GetString(config.Keys.ServerExternalURL),
	)
	userAgentLock.Unlock()
}

// GetUserAgent returns the generated http User-Agent.
func GetUserAgent() string {
	initOnce.Do(doInit)

	userAgentLock.RLock()
	ua := userAgent
	userAgentLock.RUnlock()

	return ua
}
