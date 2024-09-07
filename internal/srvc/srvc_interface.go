package srvc

import (
	"context"
	"sso/internal/model"
	"time"
)

type (
	UserSrvc interface {
		GetByEmail(ctx context.Context, email string) (*model.User, error)
	}

	RefreshTokenSrvc interface {
		DeleteByUser(ctx context.Context, user *model.User) error
		CreateByUser(ctx context.Context, user *model.User, ip, hash string, expiresAt time.Time) (*model.RefreshToken, error)
	}
)
