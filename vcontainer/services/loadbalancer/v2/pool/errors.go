package pool

import (
	"fmt"
	vconError "github.com/vngcloud/vcontainer-sdk/error"
)

func NewErrPoolInUse(pPoolID, pInfo string) vconError.IErrorBuilder {
	err := new(ErrPoolInUse)
	err.PoolID = pPoolID
	if pInfo != "" {
		err.Info = pInfo
	}
	return err
}

func IsErrPoolInUse(pErr error) bool {
	_, ok := pErr.(*ErrPoolInUse)
	return ok
}

type ErrPoolInUse struct {
	PoolID string
	vconError.BaseError
}

func (s *ErrPoolInUse) Error() string {
	s.DefaultError = fmt.Sprintf("pool %s is in use by listener", s.PoolID)
	return s.ChoseErrString()
}
