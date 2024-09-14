package v1

import (
	"fmt"
	"net/http"
	"sso/internal/handler/v1/mwr"
	"sso/internal/handler/v1/request"
	"sso/internal/handler/v1/response"
)

type permission struct {
	rp       *request.Parser
	rb       *response.Builder
	permissionSrvc PermissionSrvc
}

func newPermission(
	mux *http.ServeMux,
	prefix string,
	rp *request.Parser,
	rb *response.Builder,
	authMwr *mwr.Auth,
	permissionSrvc PermissionSrvc,
) {
	prefix += "/permissions"
	re := permission{
		rp:       rp,
		rb:       rb,
		permissionSrvc: permissionSrvc,
	}

	mux.HandleFunc("GET "+prefix, authMwr.MwrFunc(re.list))
	mux.HandleFunc("POST "+prefix, authMwr.MwrFunc(re.create))
	mux.HandleFunc("GET "+prefix+"/{id}", authMwr.MwrFunc(re.show))
	mux.HandleFunc("DELETE "+prefix+"/{id}", authMwr.MwrFunc(re.delete))
}

// @Summary permissions list
// @Tags permissions
// @Security BearerAuth
// @Router /v1/permissions [get]
// @Param pagination[page] query int false "page"
// @Param pagination[count] query int false "count"
// @Param sorts[created_at] query string false "created_at" Enums(asc, desc)
// @Param sorts[updated_at] query string false "updated_at" Enums(asc, desc)
// @Param sorts[name] query string false "name" Enums(asc, desc)
// @Param sorts[slug] query string false "slug" Enums(asc, desc)
// @Param filters[name] query string false "name"
// @Param filters[slug] query string false "slug"
// @Produce json
// @Success 200 {object} response.list{data=[]model.Permission,pagination=dto.Pagination}
func (re *permission) list(w http.ResponseWriter, r *http.Request) {
	const op = "v1.permission.list"

	search := re.rp.GetQuerySearch(r)
	permissions, pagination, err := re.permissionSrvc.List(
		r.Context(),
		search.Pagination.Page,
		search.Pagination.Count,
		search.Filters,
		search.Sorts,
	)
	if err != nil {
		re.rb.JsonFail(w, r, fmt.Errorf("%s: %w", op, err))
		return
	}

	re.rb.JsonList(w, r, permissions, pagination)
}

// @Summary create permission
// @Tags permissions
// @Security BearerAuth
// @Router /v1/permissions [post]
// @Accept json
// @Param body body request.PermissionCreate true "permission create request"
// @Produce json
// @Success 201 {object} response.success{data=model.Permission}
func (re *permission) create(w http.ResponseWriter, r *http.Request) {
	const op = "v1.permission.create"

	var req request.PermissionCreate
	err := re.rp.ParseBody(r, &req)
	if err != nil {
		re.rb.JsonFail(w, r, fmt.Errorf("%s: %w", op, err))
		return
	}

	permission, err := re.permissionSrvc.Create(r.Context(), req.Name)
	if err != nil {
		re.rb.JsonFail(w, r, fmt.Errorf("%s: %w", op, err))
		return
	}

	re.rb.JsonSuccess(w, r, http.StatusCreated, permission)
}

// @Summary get permission by id
// @Tags permissions
// @Security BearerAuth
// @Router /v1/permissions/{id} [get]
// @Param id path string true "permission id"
// @Produce json
// @Success 200 {object} response.success{data=model.Permission}
func (re *permission) show(w http.ResponseWriter, r *http.Request) {
	const op = "v1.permission.show"

	id := r.PathValue("id")
	permission, err := re.permissionSrvc.GetByID(r.Context(), id)
	if err != nil {
		re.rb.JsonFail(w, r, fmt.Errorf("%s: %w", op, err))
		return
	}

	re.rb.JsonSuccess(w, r, http.StatusOK, permission)
}

// @Summary delete permission by id
// @Tags permissions
// @Security BearerAuth
// @Router /v1/permissions/{id} [delete]
// @Param id path string true "permission id"
// @Success 204
func (re *permission) delete(w http.ResponseWriter, r *http.Request) {
	const op = "v1.permission.delete"

	id := r.PathValue("id")
	err := re.permissionSrvc.Delete(r.Context(), id)
	if err != nil {
		re.rb.JsonFail(w, r, fmt.Errorf("%s: %w", op, err))
		return
	}

	re.rb.JsonSuccess(w, r, http.StatusNoContent, nil)
}
