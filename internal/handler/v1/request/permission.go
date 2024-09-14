package request

type PermissionCreate struct {
	Name   string `json:"name" validate:"required,min=5"`
}
