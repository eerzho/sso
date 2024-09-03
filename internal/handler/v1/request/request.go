package request

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

func Parse(r *http.Request, reqDTO interface{}) error {
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(reqDTO)
	if err != nil {
		return err
	}

	err = validate.Struct(reqDTO)
	if err != nil {
		return err
	}

	return nil
}
