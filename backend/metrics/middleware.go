package metrics

import (
	"context"
	"net/http"
	"time"

	"github.com/Alia5/steaminputdb.com/api"
	"github.com/google/uuid"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		httpActiveConnections.Inc()
		defer httpActiveConnections.Dec()

		sw := &api.StatusWriter{ResponseWriter: w, Status: http.StatusOK}
		start := time.Now()

		sw.Header().Set("X-Request-ID", uuid.NewString())
		next.ServeHTTP(sw, r)

		dur := time.Since(start)

		httpRequestsTotal.WithLabelValues(
			r.Method, r.URL.Path, http.StatusText(sw.Status),
		).Inc()
		httpRequestDuration.WithLabelValues(
			r.Method, r.URL.Path,
		).Observe(dur.Seconds())

		ctx := r.Context()
		//revive:disable-next-line
		ctx = context.WithValue(ctx, "status_code", sw.Status)
		//revive:disable-next-line
		ctx = context.WithValue(ctx, "duration", dur.Milliseconds())
		*r = *r.WithContext(ctx)
	})
}
