package client

import (
	"context"
	"github.com/vngcloud/vcontainer-sdk/error/utils"
	"sync"

	"github.com/imroc/req/v3"
)

type (
	ProviderClient struct {
		IdentityEndpoint  string
		AccessToken       string
		HTTPClient        *req.Client
		ReauthFunc        func() error
		ThrowAway         bool
		Context           context.Context
		RetryBackoffFunc  RetryBackoffFunc
		MaxBackoffRetries uint

		mut       *sync.RWMutex
		reauthmut *reauthlock
	}

	RetryBackoffFunc func(context.Context, *utils.ErrUnexpectedStatusCode, error, uint) error

	RequestOpts struct {
		JSONBody         interface{}
		JSONResponse     interface{}
		OkCodes          []int
		MoreHeaders      map[string]string
		OmitHeaders      []string
		ErrorContext     error
		KeepResponseBody bool
	}

	reauthlock struct {
		sync.RWMutex
		ongoing *reauthFuture
	}

	reauthFuture struct {
		done chan struct{}
		err  error
	}

	requestState struct {
		requestType string
	}
)
