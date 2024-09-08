package v1

import (
	"context"
	"sso/internal/dto"
	"sso/internal/model"
)

type (
	UserSrvc interface {
		List(ctx context.Context, page, count int, filters, sorts map[string]string) ([]model.User, *dto.Pagination, error)
		Create(ctx context.Context, email, name, password string) (*model.User, error)
		GetByID(ctx context.Context, id string) (*model.User, error)
		Update(ctx context.Context, id, name string) (*model.User, error)
		Delete(ctx context.Context, id string) error
	}

	AuthSrvc interface {
		Login(ctx context.Context, email, password, ip string) (*dto.Token, error)
		DecodeAToken(ctx context.Context, aToken string) (*dto.Claims, error)
		Refresh(ctx context.Context, aToken, rToken, ip string) (*dto.Token, error)
	}

	RoleSrvc interface {
		List(ctx context.Context, page, count int, filters, sorts map[string]string) ([]model.Role, *dto.Pagination, error)
		Create(ctx context.Context, name string) (*model.Role, error)
		GetByID(ctx context.Context, id string) (*model.Role, error)
		Delete(ctx context.Context, id string) error
	}
)
