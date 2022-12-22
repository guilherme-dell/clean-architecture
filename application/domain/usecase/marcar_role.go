package usecase

import (
	entities "app/domain/entities/role"
	"app/repository"
)

type MarcarRole interface {
	Criar(*entities.RoleIn) (*entities.RoleOut, error)
}

type criarRole struct {
	rep repository.RoleRepository
}
