package server

import "github.com/vngcloud/vcontainer-sdk/client"

func getServerURL(sc *client.ServiceClient, opts GetOptsBuilder) string {
	return sc.ServiceURL(
		opts.ExtractProjectID(),
		"servers",
		opts.ExtractInstanceID())
}
