package main_gateways_http_controllers_beans

import (
	main_gateways_http_controllers "mpindicator/main/gateways/http/controllers"
	main_usecases_beans "mpindicator/main/usecases/beans"
)

func GetUserControllerBeanFunc() *main_gateways_http_controllers.UserController {
	useCase := main_usecases_beans.GetBeans().CreateUserBean
	userController := main_gateways_http_controllers.NewUserController(useCase)
	return userController
}
