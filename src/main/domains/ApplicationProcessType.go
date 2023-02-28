package main_domains

import "errors"

type ApplicationProcessType string

const invalidAppIndicatorTypeDescription = "Invalid Indicator description."

const appIndicatorTypeFindIndicatorDescription = "find-indicator"
const appIndicatorTypeFindUserDescription = "find-user"
const appIndicatorTypeProcessIndicatorDescription = "process-indicator"
const appIndicatorTypeReprocessIndicatorDescription = "reprocess-indicator"

const (
	APP_INDICATOR_TYPE_FIND_INDICATOR      ApplicationProcessType = appIndicatorTypeFindIndicatorDescription
	APP_INDICATOR_TYPE_FIND_USER           ApplicationProcessType = appIndicatorTypeFindUserDescription
	APP_INDICATOR_TYPE_PROCESS_INDICATOR   ApplicationProcessType = appIndicatorTypeProcessIndicatorDescription
	APP_INDICATOR_TYPE_REPROCESS_INDICATOR ApplicationProcessType = appIndicatorTypeReprocessIndicatorDescription
)

func (s ApplicationProcessType) GetDescription() string {
	switch s {
	case APP_INDICATOR_TYPE_FIND_INDICATOR:
		return appIndicatorTypeFindIndicatorDescription
	case APP_INDICATOR_TYPE_FIND_USER:
		return appIndicatorTypeFindUserDescription
	case APP_INDICATOR_TYPE_PROCESS_INDICATOR:
		return appIndicatorTypeProcessIndicatorDescription
	case APP_INDICATOR_TYPE_REPROCESS_INDICATOR:
		return appIndicatorTypeReprocessIndicatorDescription
	}
	return "unknown"
}

func FindApplicationProcessTypeByDescription(description string) (ApplicationProcessType, error) {
	switch description {
	case appIndicatorTypeFindIndicatorDescription:
		return APP_INDICATOR_TYPE_FIND_INDICATOR, nil
	case appIndicatorTypeFindUserDescription:
		return APP_INDICATOR_TYPE_FIND_USER, nil
	case appIndicatorTypeProcessIndicatorDescription:
		return APP_INDICATOR_TYPE_PROCESS_INDICATOR, nil
	case appIndicatorTypeReprocessIndicatorDescription:
		return APP_INDICATOR_TYPE_REPROCESS_INDICATOR, nil
	}
	return "", errors.New(invalidAppIndicatorTypeDescription)
}
