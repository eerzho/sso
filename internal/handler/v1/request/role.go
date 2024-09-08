package request

type RoleCreate struct {
	Name string `json:"name" validate:"required,min=5"`
}
