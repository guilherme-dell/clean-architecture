package usecase

import (
	entities "app/domain/entities/role"
	"errors"
)

type GetAllRolesUseCase interface {
	Execute() ([]*entities.RoleOut, error)
}

type GetAllRolesRepository interface {
	GetAll() ([]*entities.Role, error)
}

type getAllRoles struct {
	repository GetAllRolesRepository
}

func NewGetAllRolesUseCase(rep GetAllRolesRepository) GetAllRolesUseCase {
	return &getAllRoles{rep}
}

func (g *getAllRoles) Execute() ([]*entities.RoleOut, error) {
	roles, err := g.repository.GetAll()

	if err != nil {
		return nil, err
	}
	if roles == nil {
		return nil, errors.New("Nenhum role cadastrado")
	}

	var rolesFound []*entities.RoleOut

	for _, role := range roles {
		temp := &entities.RoleOut{
			ID:      role.ID,
			Nome:    role.Nome,
			Local:   role.Local,
			Cidade:  role.Cidade,
			Horario: role.Horario,
			Dia:     role.Dia,
			Tipo:    role.Tipo,
		}
		rolesFound = append(rolesFound, temp)
	}
	return rolesFound, nil
}
