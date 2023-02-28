package test_usecases

import (
	"context"
	main_domains "mpindicator/main/domains"
	main_gateways "mpindicator/main/gateways"
	main_usecases "mpindicator/main/usecases"
	main_utils "mpindicator/main/utils"
	test_support "mpindicator/test/support"
	"reflect"
	"testing"
	"time"
)

type IndicatorDatabaseGatewayMock struct {
}

func (thisMock *IndicatorDatabaseGatewayMock) FindById(ctx context.Context, id string) (main_domains.Indicator, error) {
	return main_domains.Indicator{}, nil
}

func (thisMock *IndicatorDatabaseGatewayMock) Save(ctx context.Context, indicator main_domains.Indicator) (string, error) {
	return "main_domains.Indicator{}", nil
}

func (thisMock *IndicatorDatabaseGatewayMock) SaveAll(ctx context.Context, indicators []main_domains.Indicator) ([]string, error) {
	var ids []string
	return ids, nil
}

func TestShouldReturnFindIndicatorById(t *testing.T) {

	testSupport := test_support.NewTestSupport(t)

	var indicatorDatabaseGateway main_gateways.IndicatorDatabaseGateway

	testee := IndicatorDatabaseGatewayMock{}
	indicatorDatabaseGateway = &testee

	expectedType := reflect.TypeOf(&main_usecases.FindIndicatorById{}).String()

	result := reflect.TypeOf(main_usecases.NewFindIndicatorById(&indicatorDatabaseGateway)).String()

	testSupport.AssertEqualsWithMsg("Date differ from expected date", expectedType, result)
}

func TestShouldReturnIndicator(t *testing.T) {

	testSupport := test_support.NewTestSupport(t)

	offsetInDays := 3
	expectedDate := time.Now().AddDate(0, 0, offsetInDays)

	result := main_utils.GetDateWithOffSetFromTodayAtStartOfDay(offsetInDays)

	testSupport.AssertEqualsWithMsg("Date differ from expected date", expectedDate.Day(), result.Day())
	testSupport.AssertEqualsWithMsg("Date differ from expected date", expectedDate.Month(), result.Month())
	testSupport.AssertEqualsWithMsg("Date differ from expected date", expectedDate.Year(), result.Year())
}

func TestShouldThrowErrorDueToErrorByIndicatorGateway(t *testing.T) {

	testSupport := test_support.NewTestSupport(t)

	offsetInDays := 3
	expectedDate := time.Now().AddDate(0, 0, offsetInDays)

	result := main_utils.GetDateWithOffSetFromTodayAtStartOfDay(offsetInDays)

	testSupport.AssertEqualsWithMsg("Date differ from expected date", expectedDate.Day(), result.Day())
	testSupport.AssertEqualsWithMsg("Date differ from expected date", expectedDate.Month(), result.Month())
	testSupport.AssertEqualsWithMsg("Date differ from expected date", expectedDate.Year(), result.Year())
}
