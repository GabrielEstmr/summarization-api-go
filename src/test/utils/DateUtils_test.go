package test_utils

import (
	main_domains "mpindicator/main/domains"
	main_utils "mpindicator/main/utils"
	test_support "mpindicator/test/support"
	"testing"
	"time"
)

const THIRTY_DAY_IN_THE_PAST_OFFSET = -2

func TestDateUtilsGetDateWithOffSetFromTodayAtStartOfDay(t *testing.T) {

	testSupport := test_support.NewTestSupport(t)

	offsetInDays := 3
	expectedDate := time.Now().AddDate(0, 0, offsetInDays)

	result := main_utils.GetDateWithOffSetFromTodayAtStartOfDay(offsetInDays)

	testSupport.AssertEqualsWithMsg("Date differ from expected date", expectedDate.Day(), result.Day())
	testSupport.AssertEqualsWithMsg("Date differ from expected date", expectedDate.Month(), result.Month())
	testSupport.AssertEqualsWithMsg("Date differ from expected date", expectedDate.Year(), result.Year())
}

func TestGetThirtyDaysInThePastAtStartOfDayAsString(t *testing.T) {

	testSupport := test_support.NewTestSupport(t)

	offsetInDays := THIRTY_DAY_IN_THE_PAST_OFFSET
	datePattern := main_domains.YYYYMMDD
	expectedDate := time.Now().AddDate(0, 0, offsetInDays).Format(datePattern.GetDescription())

	result := main_utils.GetThirtyDaysInThePastAtStartOfDayAsString(datePattern)

	testSupport.AssertEqualsWithMsg("Date differ from expected date", expectedDate, result)
}

func TestGetTodayAsString(t *testing.T) {

	testSupport := test_support.NewTestSupport(t)

	datePatternInput := main_domains.YYYYMMDD
	expectedDate := time.Now().Format(datePatternInput.GetDescription())

	result := main_utils.GetTodayAsString(datePatternInput)

	testSupport.AssertEqualsWithMsg("Date differ from expected date", expectedDate, result)
}
