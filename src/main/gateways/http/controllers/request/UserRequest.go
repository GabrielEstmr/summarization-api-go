package main_gateways_http_controllers_requests

import (
	"errors"
	main_domains "mpindicator/main/domains"
)

const MSG_INVALID_USER_PARAMETERS = "Invalid user parameters"

type UserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (userResponse *UserRequest) Validate() error {
	if userResponse.Name == "" || userResponse.Email == "" {
		return errors.New(MSG_INVALID_USER_PARAMETERS)
	}
	return nil
}

func (userResponse *UserRequest) ToDomain() main_domains.User {
	return main_domains.User{
		Name:  userResponse.Name,
		Email: userResponse.Email,
	}
}
