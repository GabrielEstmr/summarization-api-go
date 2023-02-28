package main_gateways_http_controllers_beans

import (
	main_gateways_http_controllers "mpindicator/main/gateways/http/controllers"
	main_usecases_beans "mpindicator/main/usecases/beans"
)

func GetIndicatorControllerBeanFunc() *main_gateways_http_controllers.IndicatorController {
	useCaseApiIndicatorTrigger := main_usecases_beans.GetBeans().ProcessApiIndicatorTriggerBean
	useCaseFindIndicatorByIdWithCacheAndLocked := main_usecases_beans.GetBeans().FindIndicatorByIdWithCacheAndLockedBean
	userController := main_gateways_http_controllers.NewIndicatorController(
		useCaseApiIndicatorTrigger, useCaseFindIndicatorByIdWithCacheAndLocked)
	return userController
}
