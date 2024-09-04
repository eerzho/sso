package srvc

import (
	"context"
	"sso/internal/dto"
	"sso/internal/model"
)

type User struct {
	userRepo UserRepo
}

func NewUser(userRepo UserRepo) *User {
	return &User{
		userRepo: userRepo,
	}
}

func (u *User) List(ctx context.Context, page, count int, filters, sorts map[string]string) ([]model.User, *dto.Pagination, error) {
	return u.userRepo.List(ctx, page, count, filters, sorts)
}

func (u *User) Create(ctx context.Context, email, name, password string) (*model.User, error) {
	user := model.User{
		Email:    email,
		Name:     name,
		Password: password,
	}

	err := u.userRepo.Create(ctx, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *User) Show(ctx context.Context, id string) (*model.User, error) {
	return u.userRepo.GetByID(ctx, id)
}

func (u *User) Update(ctx context.Context, id, name string) (*model.User, error) {
	user, err := u.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	user.Name = name

	err = u.userRepo.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) Delete(ctx context.Context, id string) error {
	return u.userRepo.Delete(ctx, id)
}
