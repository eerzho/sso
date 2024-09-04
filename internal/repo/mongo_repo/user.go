package mongo_repo

import (
	"context"
	"fmt"
	"sso/internal/dto"
	"sso/internal/model"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	users []model.User
}

func NewUser() *User {
	return &User{
		users: make([]model.User, 0),
	}
}

func (u *User) List(ctx context.Context, page, count int, filters, sorts map[string]string) ([]model.User, *dto.Pagination, error) {
	pagination := dto.Pagination{
		Page:  page,
		Count: count,
		Total: len(u.users),
	}

	return u.users, &pagination, nil
}

func (u *User) Create(ctx context.Context, user *model.User) error {
	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	u.users = append(u.users, *user)

	return nil
}

func (u *User) GetByID(ctx context.Context, id string) (*model.User, error) {
	idObj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	for _, user := range u.users {
		if user.ID == idObj {
			return &user, nil
		}
	}

	return nil, fmt.Errorf("not found")
}

func (u *User) Update(ctx context.Context, user *model.User) error {
	for i, _ := range u.users {
		if u.users[i].ID == user.ID {
			user.UpdatedAt = time.Now()
			u.users[i] = *user
			return nil
		}
	}

	return fmt.Errorf("not found")
}

func (u *User) Delete(ctx context.Context, id string) error {
	idObj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	for i, user := range u.users {
		if user.ID == idObj {
			u.users = append(u.users[:i], u.users[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("not found")
}
