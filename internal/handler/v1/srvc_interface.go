package v1

import "sso/internal/model"

type UserSrvc interface {
	List(page, count int, filters, sorts map[string]string) ([]model.User, error)
	Create(email, name, password string) (*model.User, error)
	Show(id string) (*model.User, error)
	Update(id, name string) (*model.User, error)
	Delete(id string) error
}
