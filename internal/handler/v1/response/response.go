package response

import (
	"encoding/json"
	"errors"
	"net/http"
	"sso/internal/def"
)

type (
	fail struct {
		Data    interface{} `json:"data,omitempty"`
		Message string      `json:"message"`
	}

	success struct {
		Data interface{} `json:"data,omitempty"`
	}

	list struct {
		Data       interface{} `json:"data"`
		Pagination interface{} `json:"pagination,omitempty"`
	}
)

var strangeCaseJson = `{"message": "` + http.StatusText(http.StatusInternalServerError) + `"}`

func JsonFail(w http.ResponseWriter, err error) {
	f := fail{Message: originalErr(err).Error()}

	code := http.StatusInternalServerError
	if errors.Is(err, def.ErrNotFound) {
		code = http.StatusNotFound
	} else if errors.Is(err, def.ErrAlreadyExists) {
		code = http.StatusBadRequest
	}

	Json(w, code, &f)
}

func JsonSuccess(w http.ResponseWriter, code int, data interface{}) {
	s := success{Data: data}

	Json(w, code, &s)
}

func JsonList(w http.ResponseWriter, data, pagination interface{}) {
	l := list{
		Data:       data,
		Pagination: pagination,
	}

	Json(w, http.StatusOK, &l)
}

func Json(w http.ResponseWriter, code int, body interface{}) {
	w.Header().Set(def.HeaderContentType.String(), "application/json")

	jsonBody, err := json.Marshal(body)
	if err != nil {
		http.Error(w, strangeCaseJson, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(code)
	w.Write(jsonBody)
}

func originalErr(err error) error {
	unwrappedErr := errors.Unwrap(err)
	if unwrappedErr == nil {
		return err
	}
	return originalErr(unwrappedErr)
}
