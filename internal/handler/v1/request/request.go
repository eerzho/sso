package request

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

func ParseBody(r *http.Request, req interface{}) error {
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(req)
	if err != nil {
		return err
	}

	err = validate.Struct(req)
	if err != nil {
		return err
	}

	return nil
}

func GetQuerySearch(r *http.Request) *Search {
	return &Search{
		Pagination: Pagination{
			Page:  GetQueryInt(r, "pagination[page]", 1),
			Count: GetQueryInt(r, "pagination[count]", 10),
		},
		Filters: GetQueryMap(r, "filters"),
		Sorts:   GetQueryMap(r, "sorts"),
	}
}

func GetQueryMap(r *http.Request, key string) map[string]string {
	valuesMap := make(map[string]string)
	
	for queryKey, values := range r.URL.Query() {
		if strings.HasPrefix(queryKey, key+"[") && strings.HasSuffix(queryKey, "]") {
			fieldName := strings.TrimSuffix(strings.TrimPrefix(queryKey, key+"["), "]")
			if len(values) > 0 {
				valuesMap[fieldName] = values[0]
			}
		}
	}

	return valuesMap
}

func GetQueryInt(r *http.Request, key string, defaultValue int) int {
	valueStr := r.URL.Query().Get(key)
	if valueStr == "" {
		return defaultValue
	}

	valueInt, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultValue
	}

	return valueInt
}
