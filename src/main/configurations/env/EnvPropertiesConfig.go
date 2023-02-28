package main_configurations_env

import (
	"errors"
	"github.com/joho/godotenv"
	"log"
	"os"
	"sync"
)

const MSG_ENV_BEAN = "Invalid env property value."
const MSG_ERROR_READ_ENV_FILE = "Error to read .env file."

const YML_BASE_DIRECTORY_MAIN_REFERENCE = "./zresources"

var once sync.Once
var EnvConfigs *map[string]string

func GetBeanPropertyByName(envName EnvironmentProperty) string {
	GetEnvConfigBean()
	properties := *EnvConfigs
	property := properties[envName.GetDescription()]
	return property

}

func GetEnvConfigBean() *map[string]string {
	once.Do(func() { // <-- atomic, does not allow repeating

		if EnvConfigs == nil {
			EnvConfigs = getEnvConfig()
		} // <-- thread safe

	})
	return EnvConfigs
}

func getEnvConfig() *map[string]string {
	envNames := []string{
		MP_INDICATOR_APPLICATION_PROFILE.GetDescription(),
	}

	err := godotenv.Load(YML_BASE_DIRECTORY_MAIN_REFERENCE + "/.env")
	FailOnError(err, MSG_ERROR_READ_ENV_FILE)

	data := make(map[string]string)
	for _, value := range envNames {
		envValue := os.Getenv(value)
		if envValue == "" {
			log.Panicf("%s: %s", MSG_ENV_BEAN, errors.New(MSG_ENV_BEAN))
		}
		data[value] = envValue
	}
	return &data
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
