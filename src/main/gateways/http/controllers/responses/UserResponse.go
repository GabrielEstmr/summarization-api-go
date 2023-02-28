package main_gateways_http_controllers_responses

import (
	main_domains "mpindicator/main/domains"
	"time"
)

type UserResponse struct {
	Id          string
	Name        string
	Nick        string
	Email       string
	CreatedDate time.Time
}

func NewUserResponse(user main_domains.User) UserResponse {
	return UserResponse{
		Id:          user.Id,
		Name:        user.Name,
		Nick:        user.Nick,
		Email:       user.Email,
		CreatedDate: user.CreatedDate,
	}
}
