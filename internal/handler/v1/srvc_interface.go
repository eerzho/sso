package v1

import "sso/internal/model"

type UserSrvc interface {
	List() []model.User
	Show(id string) *model.User
}
