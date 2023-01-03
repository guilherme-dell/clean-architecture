package usecase

import (
	"app/domain/entities/role"
	"github.com/go-playground/validator/v10"
)

type UpdateRoleUseCase interface {
	Execute(*entities.Role) (*entities.RoleOut, error)
}
type UpdateRoleUseCaseRepository interface {
	Update(*entities.Role) (*entities.RoleOut, error)
}

type updateRole struct {
	repository UpdateRoleUseCaseRepository
}

func NewUpdateRoleUseCase(rep UpdateRoleUseCaseRepository) UpdateRoleUseCase {
	return &updateRole{rep}
}

func (u *updateRole) Execute(roleInput *entities.Role) (*entities.RoleOut, error) {
	validate := validator.New()

	err := validate.Struct(roleInput)
	if err != nil {
		return nil, err
	}

	role, err := u.repository.Update(roleInput)
	if err != nil {
		if err.Error() == "Role não encontrado" {
			return nil, err
		}
		return nil, err
	}
	roleChanged := &entities.RoleOut{
		ID:      role.ID,
		Nome:    role.Nome,
		Local:   role.Local,
		Cidade:  role.Cidade,
		Horario: role.Horario,
		Dia:     role.Dia,
		Tipo:    role.Tipo,
	}
	return roleChanged, nil
}
