package http

import (
	"net/http"
)

func (s *Server) middlewareMetrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		metric := s.metrics.NewHTTPRequest(r.Method, r.URL.Path)
		l := logger.WithField("func", "middlewareMetrics")

		wx := NewResponseWriter(w)

		// Do Request
		next.ServeHTTP(wx, r)

		go func() {
			ended := metric.Done(wx.Status())
			l.Debugf("rendering %s took %d ms", r.URL.Path, ended.Milliseconds())
		}()
	})
}
