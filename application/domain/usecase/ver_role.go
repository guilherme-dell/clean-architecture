package usecase

import "app/domain/entities/role"

type RoleMethods interface {
	GetRoleById(id string) *entities.RoleIn
}
