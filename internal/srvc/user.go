package srvc

import (
	"fmt"
	"sso/internal/model"

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

func (u *User) List(page, count int, filters, sorts map[string]string) ([]model.User, error) {
	return u.users, nil
}

func (u *User) Create(email, name, password string) (*model.User, error) {
	user := model.User{
		ID: primitive.NewObjectID(),
		Email: email,
		Name: name,
		Password: password,
	}

	u.users = append(u.users, user)

	return &user, nil
}

func (u *User) Show(id string) (*model.User, error) {
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

func (u *User) Update(id, name string) (*model.User, error) {
	idObj, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	for i, user := range u.users {
		if user.ID == idObj {
			u.users[i].Name = name
			return &u.users[i], nil
		}
	}

	return nil, fmt.Errorf("not found")
}

func (u *User) Delete(id string) error {
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
