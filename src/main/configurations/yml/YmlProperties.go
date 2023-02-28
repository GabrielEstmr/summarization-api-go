package main_configurations_yml

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	main_configurations_profile "mpindicator/main/configurations/profile"
	"strings"
	// main_utils "mpindicator/main/utils"
	"os"
	"sync"
)

const MSG_YML_BEANS = "Initializing yml properties bean."
const MSG_ERROR_MANDATORY_ENV = "Mandatory env variable not found: %s"
const MSG_ERROR_MANDATORY_SEPARATOR = "Mandatory separator not found: %s"
const MSG_ERROR_READ_YML = "Error to read yml file."
const MSG_ERROR_PARSE_YML = "Error to parse yml file."

const YML_BASE_DIRECTORY_MAIN_REFERENCE = "./zresources"
const IDX_START_ENV_SEPARATOR = "${"
const IDX_END_ENV_SEPARATOR = "}"
const YML_FILE_DEFAULT_BASE_NAME = "/application-properties-%s.yaml"

var once sync.Once
var YmlConfigs *map[string]Property

type Property struct {
	Value string
}

func GetBeanPropertyByName(propertyName string) string {
	properties := *YmlConfigs
	property := properties[propertyName]
	return property.Value
}

func GetYmlConfigBean() *map[string]Property {

	once.Do(func() { // <-- atomic, does not allow repeating

		if YmlConfigs == nil {
			YmlConfigs = getYmlConfig()
		} // <-- thread safe

	})
	return YmlConfigs
}

func getYmlConfig() *map[string]Property {

	log.Println(MSG_YML_BEANS)
	profile := main_configurations_profile.GetProfileBean().GetLowerCaseDescription()
	ymlPath := YML_BASE_DIRECTORY_MAIN_REFERENCE + fmt.Sprintf(
		YML_FILE_DEFAULT_BASE_NAME, profile)

	yFile, err := os.ReadFile(ymlPath)
	failOnError(err, MSG_ERROR_READ_YML)

	data := make(map[string]Property)
	err2 := yaml.Unmarshal(yFile, &data)
	failOnError(err2, MSG_ERROR_PARSE_YML)

	for key, property := range data {
		log.Println(property)
		for {
			newValue, hasIdx := replaceEnvIdxToValue(data[key].Value)
			if hasIdx {
				data[key] = Property{newValue}
			}
			if !hasEnvIdxToSubstitute(data[key].Value) {
				break
			}
		}
	}
	return &data
}

func replaceEnvIdxToValue(value string) (string, bool) {
	hasIdx := hasEnvIdxToSubstitute(value)
	if hasEnvIdxToSubstitute(value) {
		before, afterTemp := cutOrPanic(value, IDX_START_ENV_SEPARATOR)
		envIndex, after := cutOrPanic(afterTemp, IDX_END_ENV_SEPARATOR)
		envValue := getEnvOrPanic(envIndex)
		newValue := buildNewValue(before, envValue, after)
		return newValue, hasIdx
	}
	return value, hasIdx
}

func hasEnvIdxToSubstitute(value string) bool {
	idxStartEnv := strings.Index(value, IDX_START_ENV_SEPARATOR)
	idxEndEnv := strings.Index(value, IDX_END_ENV_SEPARATOR)
	return idxStartEnv != -1 && idxEndEnv != -1
}

func buildNewValue(before string, envValue string, after string) string {
	return before + envValue + after
}

func cutOrPanic(value string, sep string) (string, string) {
	before, after, foundPrefix := strings.Cut(value, sep)
	if !foundPrefix {
		panic(fmt.Sprintf(MSG_ERROR_MANDATORY_SEPARATOR, value))
	}
	return before, after
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func getEnvOrPanic(env string) string {
	res := os.Getenv(env)
	if len(res) == 0 {
		panic(fmt.Sprintf(MSG_ERROR_MANDATORY_ENV + env))
	}
	return res
}
