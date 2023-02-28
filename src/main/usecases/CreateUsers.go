package main_usecases

import (
	"context"
	"log"
	main_domains "mpindicator/main/domains"
	main_gateways "mpindicator/main/gateways"
)

type CreateUser struct {
	UserDatabaseGateway main_gateways.UserDatabaseGateway
}

func NewCreateUser(userDatabaseGateway *main_gateways.UserDatabaseGateway) *CreateUser {
	return &CreateUser{UserDatabaseGateway: *userDatabaseGateway}
}

func (thisUseCase *CreateUser) Execute(ctx context.Context, user main_domains.User) (string, error) {

	id, err := thisUseCase.UserDatabaseGateway.Save(ctx, user)

	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return id, nil
}
