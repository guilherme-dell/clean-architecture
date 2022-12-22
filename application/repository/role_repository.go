package repository

import "app/domain/entities/role"

type RoleRepository interface {
	MarcarRole(role entities.RoleIn) (*entities.RoleOut, error)
	GetRoleById(id string) (*entities.RoleOut, error)
}
