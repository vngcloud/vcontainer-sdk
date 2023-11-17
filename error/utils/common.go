package utils

import vconError "github.com/vngcloud/vcontainer-sdk/error"

// **************************************************** ErrUnknown *****************************************************

func NewErrUnknown(pInfo string) vconError.IErrorBuilder {
	err := new(ErrUnknown)
	if pInfo != "" {
		err.Info = pInfo
	}
	return err
}

func IsErrUnknown(pErr error) bool {
	_, ok := pErr.(*ErrUnknown)
	return ok
}

type ErrUnknown struct {
	vconError.BaseError
}

func (s *ErrUnknown) Error() string {
	s.DefaultError = "unknown error, something went wrong"
	return s.ChoseErrString()
}
