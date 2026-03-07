package metrics

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/Alia5/steaminputdb.com/api"
	"github.com/Alia5/steaminputdb.com/api/ctx"
	"github.com/google/uuid"
)

func metricPathLabel(r *http.Request) string {
	if r.Pattern == "" {
		return "/unmatched"
	}

	pattern := r.Pattern
	if sep := strings.IndexByte(pattern, ' '); sep >= 0 {
		pattern = pattern[sep+1:]
	}

	if pattern == "" {
		return "/unmatched"
	}

	return pattern
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		httpActiveConnections.Inc()
		defer httpActiveConnections.Dec()

		sw := &api.StatusWriter{ResponseWriter: w, Status: http.StatusOK}
		start := time.Now()

		sw.Header().Set("X-Request-ID", uuid.NewString())
		next.ServeHTTP(sw, r)

		dur := time.Since(start)
		pathLabel := metricPathLabel(r)

		httpRequestsTotal.WithLabelValues(
			r.Method, pathLabel, http.StatusText(sw.Status),
		).Inc()
		httpRequestDuration.WithLabelValues(
			r.Method, pathLabel,
		).Observe(dur.Seconds())

		c := r.Context()
		c = context.WithValue(c, ctx.KeyStatusCode, sw.Status)
		c = context.WithValue(c, ctx.KeyDuration, dur.Milliseconds())
		if sw.Error != nil {
			c = context.WithValue(c, ctx.KeyError, sw.Error)
		}
		*r = *r.WithContext(c)
	})
}
