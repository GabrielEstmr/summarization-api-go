package main_gateways_http_controllers_beans

import (
	main_gateways_http_controllers "mpindicator/main/gateways/http/controllers"
	"sync"
)

var once sync.Once

var Beans *ControllerBeans = nil

type ControllerBeans struct {
	UserControllerBean      *main_gateways_http_controllers.UserController
	IndicatorControllerBean *main_gateways_http_controllers.IndicatorController
}

func GetBeans() *ControllerBeans {
	once.Do(func() {

		if Beans == nil {
			Beans = getFunctionBeans()
		}

	})
	return Beans
}

func getFunctionBeans() *ControllerBeans {

	bean := ControllerBeans{
		UserControllerBean:      GetUserControllerBeanFunc(),
		IndicatorControllerBean: GetIndicatorControllerBeanFunc(),
	}

	return &bean
}
