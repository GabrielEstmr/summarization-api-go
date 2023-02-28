package test_support

import (
	"runtime"
	"testing"
)

type TestSupport struct {
	t *testing.T
}

func NewTestSupport(t *testing.T) *TestSupport {
	return &TestSupport{t}
}

func (thisStruct *TestSupport) AssertEquals(expected any, actual any) {
	if expected != actual {
		thisStruct.t.Error(expected, actual)
	}
}

func (thisStruct *TestSupport) AssertEqualsWithMsg(message string, expected any, actual any) {
	if expected != actual {
		thisStruct.t.Error(message, expected, actual)
	}
}

func (thisStruct *TestSupport) Mock(interfaceToMock interface{}, functionToMock runtime.Func) {

}
