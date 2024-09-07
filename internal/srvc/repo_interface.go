package srvc

import (
	"context"
	"sso/internal/dto"
	"sso/internal/model"
)

type (
	UserRepo interface {
		List(ctx context.Context, page, count int, filters, sorts map[string]string) ([]model.User, *dto.Pagination, error)
		Create(ctx context.Context, user *model.User) error
		GetByID(ctx context.Context, id string) (*model.User, error)
		Update(ctx context.Context, user *model.User) error
		Delete(ctx context.Context, id string) error
		IsExistsEmail(ctx context.Context, email string) (bool, error)
		GetByEmail(ctx context.Context, email string) (*model.User, error)
	}

	RefreshTokenRepo interface {
		DeleteByUser(ctx context.Context, user *model.User) error
		Create(ctx context.Context, refreshToken *model.RefreshToken) error
		GetByUserAndID(ctx context.Context, user *model.User, id string) (*model.RefreshToken, error)
	}
)
