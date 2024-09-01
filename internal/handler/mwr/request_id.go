package mwr

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type RequestID struct {
	key string
}

func NewRequestId(key string) *RequestID {
	return &RequestID{
		key: key,
	}
}

func (req *RequestID) Mwr(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		const op = "handler.mwr.RequestId.Mwr"

		reqID := uuid.New().String()

		w.Header().Set(req.key, reqID)

		ctx := r.Context()
		ctx = context.WithValue(ctx, req.key, reqID)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
