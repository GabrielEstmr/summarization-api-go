package main_utils

import (
	main_domains "mpindicator/main/domains"
	"time"
)

// THIRTY_DAY_IN_THE_PAST_OFFSET has been set to 2 due to testing only propose
const THIRTY_DAY_IN_THE_PAST_OFFSET = -2

func GetDateWithOffSetFromTodayAtStartOfDay(offSet int) time.Time {
	myDateNow := time.Now()
	return myDateNow.AddDate(0, 0, offSet)
}

func GetThirtyDaysInThePastAtStartOfDayAsString(datePattern main_domains.DatePattern) string {
	return GetDateWithOffSetFromTodayAtStartOfDay(
		THIRTY_DAY_IN_THE_PAST_OFFSET).Format(datePattern.GetDescription())
}

func GetTodayAsString(datePattern main_domains.DatePattern) string {
	return time.Now().Format(datePattern.GetDescription())
}
