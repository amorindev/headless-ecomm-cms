package initializer

import (
	"context"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/roles/domain"
	"github.com/amorindev/headless-ecomm-cms/pkg/app/roles/errors"
	"github.com/amorindev/headless-ecomm-cms/pkg/app/roles/port"
)

type RoleInitializer struct {
	RoleRepo port.RoleRepo
}

func NewRoleItz(roleRepo port.RoleRepo) *RoleInitializer {
    return &RoleInitializer{
		RoleRepo: roleRepo,
	}
}

func (ri *RoleInitializer) SeedEssentialRoles(ctx context.Context) error {
	roles := []*domain.Role{
		domain.NewRole("Admin"),
		domain.NewRole("User"),
	}

	for _, role := range roles {
		exists, err := ri.RoleRepo.Exists(ctx, role.Name)
		if err != nil {
			if err != errors.ErrRoleNotFound {
				return err
			}
		}
		if !exists {
			if err := ri.RoleRepo.Insert(ctx, role); err != nil {
				return err
			}
		}
	}
	return nil
}
