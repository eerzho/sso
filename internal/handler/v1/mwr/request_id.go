package mwr

import (
	"context"
	"net/http"
	"sso/internal/def"

	"github.com/google/uuid"
)

type RequestID struct {
}

func NewRequestId() *RequestID {
	return &RequestID{}
}

func (req *RequestID) Mwr(next http.Handler) http.Handler {
	return req.MwrFunc(next.ServeHTTP)
}

func (req *RequestID) MwrFunc(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqID := uuid.New().String()

		w.Header().Set(string(def.RID), reqID)

		ctx := r.Context()
		ctx = context.WithValue(ctx, def.RID, reqID)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	}
}
