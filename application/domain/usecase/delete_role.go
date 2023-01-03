package usecase

import "errors"

type DeleteRoleUseCase interface {
	Execute(id string) error
}

type DeleteRoleRepository interface {
	Delete(id string) error
}

type deleteRole struct {
	repository DeleteRoleRepository
}

func NewDeleteRoleUseCase(rep DeleteRoleRepository) DeleteRoleUseCase {
	return &deleteRole{rep}
}

func (d *deleteRole) Execute(id string) error {
	err := d.repository.Delete(id)
	if err != nil {
		if err.Error() == "Role não encontrado :(" {
			return err
		}
		return errors.New("Erro ao deletar role")
	}
	return err
}
