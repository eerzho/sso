package mwr

import (
	"context"
	"sso/internal/dto"
	"sso/internal/model"
)

type (
	AuthSrvc interface {
		DecodeAToken(ctx context.Context, aToken string) (*dto.Claims, error)
	}

	UserSrvc interface {
		GetByID(ctx context.Context, id string) (*model.User, error)
	}
)
