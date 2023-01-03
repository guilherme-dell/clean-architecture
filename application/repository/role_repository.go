package repository

import entities "app/domain/entities/role"

type RoleRepository interface {
	Insert(*entities.Role) (*entities.Role, error)
	GetById(id string) (*entities.Role, error)
	GetAll() ([]*entities.Role, error)
	Delete(id string) error
	Update(*entities.Role) (*entities.RoleOut, error)
}
