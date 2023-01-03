package firestore

import (
	entities "app/domain/entities/role"
	"app/repository"
	"cloud.google.com/go/firestore"
	"context"
	"errors"
	"github.com/google/uuid"
	"log"
	"os"
)

type rep struct {
	client *firestore.Client
}

func startFireEnvp() {
	err := os.Setenv("FIRESTORE_EMULATOR_HOST", "firebase:8081")
	if err != nil {
		log.Fatalln(err)
	}
}

func createClientFireStore() *firestore.Client {
	startFireEnvp()
	ctx := context.Background()

	client, err := firestore.NewClient(ctx, "bexs-challenge1")
	if err != nil {
		log.Println(err)
	}
	return client
}

func NewFireStoreConnection() repository.RoleRepository {
	client := createClientFireStore()

	return &rep{client}
}

func (r *rep) Insert(role *entities.Role) (*entities.Role, error) {
	ctx := context.Background()
	if role.ID == "" {
		role.ID = uuid.NewString()
	}

	_, err := r.client.Collection("Roles").Doc(role.ID).Create(ctx, role)
	if err != nil {
		return nil, err
	}

	return role, nil
}

func (r *rep) GetById(id string) (*entities.Role, error) {
	ctx := context.Background()
	var roleFound *entities.Role

	role, err := r.client.Collection("Roles").Doc(id).Get(ctx)
	if err != nil {
		return nil, errors.New("Role não encontrado :(")
	}

	err = role.DataTo(&roleFound)
	if err != nil {
		return nil, err
	}

	return roleFound, nil
}

func (r *rep) GetAll() ([]*entities.Role, error) {
	ctx := context.Background()
	var roleConv *entities.Role
	var listaRoles []*entities.Role

	roles, err := r.client.Collection("Roles").Documents(ctx).GetAll()
	if err != nil {
		return nil, errors.New("Nenhum role cadastrado :p")
	}

	for _, role := range roles {
		err = role.DataTo(&roleConv)
		if err != nil {
			continue
		}
		roleConv.ID = role.Ref.ID
		listaRoles = append(listaRoles, roleConv)
	}

	return listaRoles, nil
}

func (r *rep) Delete(id string) error {
	ctx := context.Background()
	_, err := r.client.Collection("Roles").Doc(id).Delete(ctx)
	if err != nil {
		return errors.New("Role não encontrado :(")
	}
	return err
}

func (r *rep) Update(roleIn *entities.Role) (*entities.RoleOut, error) {
	ctx := context.Background()
	collection := r.client.Collection("Roles")
	role, err := collection.Doc(roleIn.ID).Get(ctx)
	if err != nil {
		return nil, errors.New("Role não encontrado")
	}
	collection.Doc(role.Ref.ID).Set(ctx, roleIn)
	return (*entities.RoleOut)(roleIn), err
}
