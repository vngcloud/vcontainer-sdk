package oauth2

import (
	"github.com/vngcloud/vcontainer-sdk/client"
)

func authURL(pSc *client.ServiceClient) string {
	return pSc.ServiceURL("auth", "token")
}
