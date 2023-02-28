package main_gateway_mongodb_documents

import (
	main_domains "mpindicator/main/domains"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserDocument struct {
	Id          primitive.ObjectID `json:"id,omitempty"`
	Name        string             `json:"name,omitempty"`
	Nick        string             `json:"nick,omitempty"`
	Email       string             `json:"email,omitempty"`
	Password    string             `json:"password,omitempty"`
	CreatedDate primitive.DateTime `json:"createdDate,omitempty"`
}

func NewUserDocument(user main_domains.User) UserDocument {
	return UserDocument{
		Name:     user.Email,
		Nick:     user.Nick,
		Email:    user.Email,
		Password: user.Password,
	}
}

func (userDocument *UserDocument) ToDomain() main_domains.User {
	return main_domains.User{
		Id:          userDocument.Id.String(),
		Name:        userDocument.Name,
		Nick:        userDocument.Nick,
		Email:       userDocument.Email,
		Password:    userDocument.Password,
		CreatedDate: userDocument.CreatedDate.Time(),
	}
}
