package request

import (
	"encoding/json"
	"errors"
	"io"
	"net"
	"net/http"
	"sso/internal/def"
	"sso/internal/model"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Parser struct {
	validate *validator.Validate
}

func NewParser() *Parser {
	return &Parser{
		validate: validator.New(validator.WithRequiredStructEnabled()),
	}
}

func (p *Parser) ParseBody(r *http.Request, req interface{}) error {
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(req)
	if err != nil {
		if errors.Is(err, io.EOF) {
			return def.ErrInvalidBody
		}
		return err
	}

	err = p.validate.Struct(req)
	if err != nil {
		return err
	}

	return nil
}

func (p *Parser) GetQuerySearch(r *http.Request) *Search {
	return &Search{
		Pagination: Pagination{
			Page:  p.GetQueryInt(r, "pagination[page]", 1),
			Count: p.GetQueryInt(r, "pagination[count]", 10),
		},
		Filters: p.GetQueryMap(r, "filters"),
		Sorts:   p.GetQueryMap(r, "sorts"),
	}
}

func (p *Parser) GetQueryMap(r *http.Request, key string) map[string]string {
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

func (p *Parser) GetQueryInt(r *http.Request, key string, defaultValue int) int {
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

func (p *Parser) GetHeaderIP(r *http.Request) string {
	forwardedFor := r.Header.Get(def.HeaderForwardedFor.String())
	if forwardedFor != "" {
		ips := strings.Split(forwardedFor, ",")

		return strings.TrimSpace(ips[0])
	}

	ip, _, _ := net.SplitHostPort(r.RemoteAddr)

	return ip
}

func (p *Parser) GetAuthUser(r *http.Request) (*model.User, error) {
	user, ok := r.Context().Value(def.ContextAuthUser).(*model.User)
	if !ok {
		return nil, def.ErrInvalidUserType
	}

	return user, nil
}
