package mwr

import (
	"log/slog"
	"net/http"
	"time"
)

type RequestLogger struct {
	lg *slog.Logger
}

func NewRequestLogger(lg *slog.Logger) *RequestLogger {
	return &RequestLogger{
		lg: lg,
	}
}

func (rl *RequestLogger) Mwr(next http.Handler) http.Handler {
	return rl.MwrFunc(next.ServeHTTP)
}

func (rl *RequestLogger) MwrFunc(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handler.mwr.RequestLogger.Mwr"
		start := time.Now()

		lg := rl.lg.With(
			slog.String("op", op),
			slog.String("method", r.Method),
			slog.String("url", r.URL.Path),
		)

		lg.Info("start")
		next.ServeHTTP(w, r)
		lg.Info("end", slog.Float64("time", time.Since(start).Seconds()))
	}
}
