package task

import (
	"context"
	"fmt"
)

type RemoveRole struct {
	ctx      context.Context
	userRepo UserRepo
	roleID   string
}

func NewRemoveRole(ctx context.Context, userRepo UserRepo, roleID string) *RemoveRole {
	return &RemoveRole{
		ctx:      ctx,
		userRepo: userRepo,
		roleID:   roleID,
	}
}

func (r *RemoveRole) Execute() error {
	const op = "task.RemoveRole.Execute"

	err := r.userRepo.RemoveRole(r.ctx, r.roleID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
