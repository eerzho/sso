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
	}
)
