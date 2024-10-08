package srvc

import (
	"context"
	"sso/internal/model"
	"time"
)

type (
	UserSrvc interface {
		GetByEmail(ctx context.Context, email string) (*model.User, error)
		GetByID(ctx context.Context, id string) (*model.User, error)
	}

	RefreshTokenSrvc interface {
		DeleteByUser(ctx context.Context, user *model.User) error
		CreateByUser(ctx context.Context, user *model.User, ip, hash string, expiresAt time.Time) (*model.RefreshToken, error)
		GetByUserAndID(ctx context.Context, user *model.User, id string) (*model.RefreshToken, error)
	}

	RoleSrvc interface {
		GetByID(ctx context.Context, id string) (*model.Role, error)
	}

	PermissionSrvc interface {
		GetByID(ctx context.Context, id string) (*model.Permission, error)
	}
)
