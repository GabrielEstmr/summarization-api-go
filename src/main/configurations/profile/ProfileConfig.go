package main_configurations_profile

import (
	"log"
	main_configurations_env "mpindicator/main/configurations/env"
	"sync"
)

const MSG_ERROR_TO_GET_PROFILE = "Error to get application profile"

var once sync.Once
var ApplicationProfileBean *ApplicationProfile

func GetProfileBean() *ApplicationProfile {

	once.Do(func() { // <-- atomic, does not allow repeating

		if ApplicationProfileBean == nil {
			ApplicationProfileBean = getProfile()
		} // <-- thread safe

	})
	return ApplicationProfileBean
}

func getProfile() *ApplicationProfile {
	profile := main_configurations_env.GetBeanPropertyByName(main_configurations_env.MP_INDICATOR_APPLICATION_PROFILE)
	appProfile, err := FindApplicationProfileByDescription(profile)
	FailOnError(err, MSG_ERROR_TO_GET_PROFILE)
	return &appProfile
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
