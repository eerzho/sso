package request

type (
	Login struct {
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	Refresh struct {
		AToken string `json:"access_token" validate:"required"`
		RToken string `json:"refresh_token" validate:"required"`
	}
)
