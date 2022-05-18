package wellknown

import (
	"encoding/json"
	"fmt"
	nethttp "net/http"
	"strconv"

	"github.com/feditools/go-lib/http"
	"github.com/gorilla/handlers"
	"github.com/tyrm/go-util/middleware"
)

func (m *Module) returnErrorPage(w nethttp.ResponseWriter, _ *nethttp.Request, code int, errStr string) {
	l := logger.WithField("func", "returnErrorPage")

	errorResp := map[string]interface{}{
		"error": map[string]interface{}{
			"detail": errStr,
			"status": strconv.FormatInt(int64(code), 10),
			"title":  nethttp.StatusText(code),
		},
	}

	w.WriteHeader(code)
	w.Header().Set(http.HeaderContentType, http.MimeAppJSON)
	err := json.NewEncoder(w).Encode(errorResp)
	if err != nil {
		l.Errorf("writing response: %s", err.Error())
	}
}

func (m *Module) methodNotAllowedHandler() nethttp.Handler {
	// wrap in middleware since middleware isn't run on error pages
	return m.wrapInMiddlewares(nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		m.returnErrorPage(w, r, nethttp.StatusMethodNotAllowed, r.Method)
	}))
}

func (m *Module) notFoundHandler() nethttp.Handler {
	// wrap in middleware since middleware isn't run on error pages
	return m.wrapInMiddlewares(nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		m.returnErrorPage(w, r, nethttp.StatusNotFound, fmt.Sprintf("page not found: %s", r.URL.Path))
	}))
}

func (m *Module) wrapInMiddlewares(h nethttp.Handler) nethttp.Handler {
	return m.srv.MiddlewareMetrics(
		handlers.CompressHandler(
			middleware.BlockFlocMux(
				h,
			),
		),
	)
}
