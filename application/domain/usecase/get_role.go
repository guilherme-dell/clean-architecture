package usecase

import (
	entities "app/domain/entities/role"
	"errors"
)

type GetRoleUseCase interface {
	Execute(id string) (*entities.RoleOut, error)
}

type GetRoleUseCaseRepository interface {
	GetById(id string) (*entities.Role, error)
}

type getRole struct {
	repository GetRoleUseCaseRepository
}

func NewGetRoleUseCase(rep GetRoleUseCaseRepository) GetRoleUseCase {
	return &getRole{rep}
}

func (g *getRole) Execute(id string) (*entities.RoleOut, error) {

	role, err := g.repository.GetById(id)
	if err != nil {
		if err.Error() == "Role não encontrado :(" {
			return nil, err
		}
		return nil, errors.New("Erro ao encontrar role :(")
	}

	roleFound := &entities.RoleOut{
		ID:      role.ID,
		Nome:    role.Nome,
		Local:   role.Local,
		Cidade:  role.Cidade,
		Horario: role.Horario,
		Dia:     role.Dia,
		Tipo:    role.Tipo,
	}

	return roleFound, nil
}
