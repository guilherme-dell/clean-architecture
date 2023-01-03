package usecase

import (
	"app/domain/entities/role"
	"errors"
	"github.com/go-playground/validator/v10"
)

type CreateRoleUseCase interface {
	Execute(*entities.RoleIn) (*entities.RoleOut, error)
}

type CreateRoleUseCaseRepository interface {
	Insert(*entities.Role) (*entities.Role, error)
}

type createNewRole struct {
	repository CreateRoleUseCaseRepository
}

func NewCreateRoleUseCase(rep CreateRoleUseCaseRepository) CreateRoleUseCase {
	return &createNewRole{rep}
}

func (c *createNewRole) Execute(roleInput *entities.RoleIn) (*entities.RoleOut, error) {

	validate := validator.New()

	err := validate.Struct(roleInput)
	if err != nil {
		return nil, err
	}

	roleConversion := &entities.Role{
		ID:      "",
		Nome:    roleInput.Nome,
		Local:   roleInput.Local,
		Cidade:  roleInput.Cidade,
		Horario: roleInput.Horario,
		Dia:     roleInput.Dia,
		Tipo:    roleInput.Tipo,
	}

	role, err := c.repository.Insert(roleConversion)
	if err != nil {
		return nil, errors.New("Erro ao marcar role :(")
	}
	roleMake := &entities.RoleOut{
		ID:      role.ID,
		Nome:    role.Nome,
		Local:   role.Local,
		Cidade:  role.Cidade,
		Horario: role.Horario,
		Dia:     role.Dia,
		Tipo:    role.Tipo,
	}

	return roleMake, nil
}
