package main_usecases_beans

import (
	main_usecases "mpindicator/main/usecases"
	"sync"
)

var once sync.Once

var Beans *UseCaseBeans = nil

type UseCaseBeans struct {
	CreateUserBean                          *main_usecases.CreateUser
	ProcessApiIndicatorTriggerBean          *main_usecases.ApiIndicatorTrigger
	FindIndicatorByIdWithCacheAndLockedBean *main_usecases.FindIndicatorByIdWithCacheAndLocked
}

func GetBeans() *UseCaseBeans {
	once.Do(func() {

		if Beans == nil {
			Beans = getFunctionBeans()
		}

	})
	return Beans
}

func getFunctionBeans() *UseCaseBeans {

	bean := UseCaseBeans{
		CreateUserBean:                          GetCreateUserBeanFunc(),
		ProcessApiIndicatorTriggerBean:          GetProcessApiIndicatorTriggerBeanFunc(),
		FindIndicatorByIdWithCacheAndLockedBean: CreateFindIndicatorByIdWithCacheAndLockedBeanFunc(),
	}

	return &bean
}
