package srvc

import (
	"context"
	"fmt"
	"sso/internal/def"
	"sso/internal/dto"
	"sso/internal/model"

	"github.com/gosimple/slug"
)

type Permission struct {
	permissionRepo PermissionRepo
}

func NewPermission(
	permissionRepo PermissionRepo,
) *Permission {
	return &Permission{
		permissionRepo: permissionRepo,
	}
}

func (r *Permission) List(ctx context.Context, page, count int, filters, sorts map[string]string) ([]model.Permission, *dto.Pagination, error) {
	const op = "srvc.Permission.List"

	permission, pagination, err := r.permissionRepo.List(ctx, page, count, filters, sorts)
	if err != nil {
		return nil, nil, fmt.Errorf("%s: %w", op, err)
	}

	return permission, pagination, nil
}

func (r *Permission) Create(ctx context.Context, name string) (*model.Permission, error) {
	const op = "srvc.Permission.Create"

	slug := slug.Make(name)
	exists, err := r.permissionRepo.IsExistsSlug(ctx, slug)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	if exists {
		return nil, fmt.Errorf("%s: %w", op, def.ErrAlreadyExists)
	}

	permission := model.Permission{
		Name: name,
		Slug: slug,
	}

	err = r.permissionRepo.Create(ctx, &permission)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &permission, nil
}

func (r *Permission) GetByID(ctx context.Context, id string) (*model.Permission, error) {
	const op = "srvc.Permission.GetByID"

	permission, err := r.permissionRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return permission, nil
}

func (r *Permission) Delete(ctx context.Context, id string) error {
	const op = "srvc.Permission.Delete"

	err := r.permissionRepo.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
