package mwr

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sso/internal/def"
	"sso/internal/handler/v1/response"
	"strings"
	"time"
)

type Auth struct {
	rb       *response.Builder
	authSrvc AuthSrvc
	userSrvc UserSrvc
}

func NewAuth(
	rb *response.Builder,
	authSrvc AuthSrvc,
	userSrvc UserSrvc,
) *Auth {
	return &Auth{
		rb:       rb,
		authSrvc: authSrvc,
		userSrvc: userSrvc,
	}
}

func (a *Auth) MwrFunc(next http.HandlerFunc) http.HandlerFunc {
	const op = "v1.mwr.Auth.MwrFunc"
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get(def.HeaderAuthorization.String())
		if authHeader == "" {
			a.rb.JsonFail(w, r, fmt.Errorf("%s: %w", op, def.ErrAuthMissing))
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			a.rb.JsonFail(w, r, fmt.Errorf("%s: %w", op, def.ErrInvalidAuthFormat))
			return
		}

		claims, err := a.authSrvc.DecodeAToken(r.Context(), parts[1])
		if err != nil {
			a.rb.JsonFail(w, r, fmt.Errorf("%s: %w", op, err))
			return
		}

		if claims.ExpiresAt.Time.Before(time.Now()) {
			a.rb.JsonFail(w, r, fmt.Errorf("%s: %w", op, def.ErrATokenExpired))
			return
		}

		user, err := a.userSrvc.GetByID(r.Context(), claims.UserID.Hex())
		if err != nil {
			if errors.Is(err, def.ErrNotFound) {
				a.rb.JsonFail(w, r, fmt.Errorf("%s: %w", op, def.ErrCannotLogin))
			}
			a.rb.JsonFail(w, r, fmt.Errorf("%s: %w", op, err))
			return
		}

		ctx := context.WithValue(r.Context(), def.ContextAuthUser, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
