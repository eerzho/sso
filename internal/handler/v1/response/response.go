package response

import (
	"encoding/json"
	"net/http"
)

type (
	fail struct {
		Data    interface{} `json:"data,omitempty"`
		Message string      `json:"message"`
	}

	success struct {
		Data interface{} `json:"data"`
	}
)

func JsonFail(w http.ResponseWriter, err error) {
	code := http.StatusInternalServerError
	f := fail{Message: err.Error()}

	Json(w, code, f)
}

func JsonSuccess(w http.ResponseWriter, code int, data interface{}) {
	s := success{Data: data}

	Json(w, code, s)
}

func Json(w http.ResponseWriter, code int, body interface{}) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
