package srvc

import (
	"context"
	"fmt"
	"sso/internal/def"
	"sso/internal/dto"
	"sso/internal/model"

	"github.com/gosimple/slug"
)

type Role struct {
	roleRepo RoleRepo
}

func NewRole(
	roleRepo RoleRepo,
) *Role {
	return &Role{
		roleRepo: roleRepo,
	}
}

func (r *Role) List(ctx context.Context, page, count int, filters, sorts map[string]string) ([]model.Role, *dto.Pagination, error) {
	const op = "srvc.Role.List"

	role, pagination, err := r.roleRepo.List(ctx, page, count, filters, sorts)
	if err != nil {
		return nil, nil, fmt.Errorf("%s: %w", op, err)
	}

	return role, pagination, nil
}

func (r *Role) Create(ctx context.Context, name string) (*model.Role, error) {
	const op = "srvc.Role.Create"

	slug := slug.Make(name)
	exists, err := r.roleRepo.IsExistsSlug(ctx, slug)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	if exists {
		return nil, fmt.Errorf("%s: %w", op, def.ErrAlreadyExists)
	}

	role := model.Role{
		Name: name,
		Slug: slug,
	}

	err = r.roleRepo.Create(ctx, &role)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &role, nil
}

func (r *Role) GetByID(ctx context.Context, id string) (*model.Role, error) {
	const op = "srvc.Role.GetByID"

	role, err := r.roleRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return role, nil
}

func (r *Role) Delete(ctx context.Context, id string) error {
	const op = "srvc.Role.Delete"

	err := r.roleRepo.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (r *Role) GetByIDs(ctx context.Context, ids []string) ([]model.Role, error) {
	const op = "srvc.Role.GetByIDs"

	roles, err := r.roleRepo.GetByIDs(ctx, ids)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return roles, nil
}
