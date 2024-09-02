package srvc

import (
	"sso/internal/model"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
}

func NewUser() *User {
	return &User{}
}

func (u *User) List() []model.User {
	users := []model.User{
		{
			ID:        primitive.NewObjectID(),
			Email:     "test@test.com",
			Password:  "password",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        primitive.NewObjectID(),
			Email:     "test@test.com",
			Password:  "password",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	return users
}

func (u *User) Show(id string) *model.User {
	user := model.User{
		ID:        primitive.NewObjectID(),
		Email:     "test@test.com",
		Password:  "password",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return &user
}
