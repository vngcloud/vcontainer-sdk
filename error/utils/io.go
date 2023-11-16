package utils

import (
	"fmt"

	vconError "github.com/vngcloud/vcontainer-sdk/error"
)

// ************************************************** ErrMissingInput **************************************************

func NewErrMissingInput(pArgument, pInfo string) *ErrMissingInput {
	err := new(ErrMissingInput)
	err.Argument = pArgument
	if pInfo != "" {
		err.Info = pInfo
	}
	return err
}

func IsErrMissingInput(pErr error) bool {
	_, ok := pErr.(*ErrMissingInput)
	return ok
}

type ErrMissingInput struct {
	Argument string

	vconError.BaseError
}

func (s *ErrMissingInput) Error() string {
	s.DefaultError = fmt.Sprintf("missing input for argument [%s]", s.Argument)
	return s.ChoseErrString()
}
