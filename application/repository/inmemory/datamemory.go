package inmemory

import (
	entities "app/domain/entities/role"
	"app/repository"
	"errors"
	"fmt"
	"github.com/google/uuid"
)

type BancoTemp struct {
	banco map[string]*entities.RoleOut
}

func NewBancotemp() repository.RoleRepository {
	return &BancoTemp{
		banco: make(map[string]*entities.RoleOut),
	}

}

func (b *BancoTemp) Insert(role *entities.Role) (*entities.Role, error) {
	if role.ID == "" {
		role.ID = uuid.NewString()
	}

	b.banco[role.ID] = (*entities.RoleOut)(role)
	for key, element := range b.banco {
		fmt.Println("Chave:", key, "Valor:", element)
	}

	return role, nil
}

func (b *BancoTemp) GetById(id string) (*entities.Role, error) {

	for key, element := range b.banco {
		if key == id {
			return (*entities.Role)(element), nil
		}
	}
	return nil, errors.New("Role não encontrado :(")
}

func (b *BancoTemp) GetAll() ([]*entities.Role, error) {
	var roles []*entities.Role

	if len(b.banco) == 0 {
		return nil, errors.New("Nenhum role cadastrado :p")
	}
	for _, element := range b.banco {
		roles = append(roles, (*entities.Role)(element))
	}
	return roles, nil
}

func (b *BancoTemp) Delete(id string) error {
	if len(b.banco) == 0 {
		return errors.New("Nenhum role foi cadastrado :(")
	}
	if _, ok := b.banco[id]; ok {
		delete(b.banco, id)
		return nil
	} else {
		return errors.New("O ID do role não foi encontrado :(")
	}
}

func (b *BancoTemp) Update(roleIn *entities.Role) (*entities.RoleOut, error) {
	if len(b.banco) == 0 {
		return nil, errors.New("Nenhum role foi cadastrado :(")
	}
	if _, ok := b.banco[roleIn.ID]; ok {
		b.banco[roleIn.ID] = (*entities.RoleOut)(roleIn)
		return (*entities.RoleOut)(roleIn), nil
	} else {
		return nil, errors.New("Role não encontrado")
	}
}
