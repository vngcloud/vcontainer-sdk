package utils

import (
	"fmt"
	vconError "github.com/vngcloud/vcontainer-sdk/error"
)

// ********************************************** ErrUnexpectedStatusCode **********************************************

func NewErrUnexpectedStatusCode(pURL, pMethod string, pExpected []int, pActual int, pInfo string) *ErrUnexpectedStatusCode {
	err := new(ErrUnexpectedStatusCode)
	err.URL = pURL
	err.Method = pMethod
	err.Expected = pExpected
	err.Actual = pActual
	if pInfo != "" {
		err.Info = pInfo
	}
	return err
}

func IsErrUnexpectedStatusCode(pErr error) bool {
	_, ok := pErr.(*ErrUnexpectedStatusCode)
	return ok
}

type ErrUnexpectedStatusCode struct {
	URL      string
	Method   string
	Expected []int
	Actual   int

	vconError.BaseError
}

func (s *ErrUnexpectedStatusCode) Error() string {
	s.DefaultError = fmt.Sprintf("expected HTTP status code [%v] when accessing [%s %s], but got %d instead\n[%s]",
		s.Expected, s.Method, s.URL, s.Actual, s.Info)
	return s.ChoseErrString()
}
