package loadbalancer

import (
	"fmt"
	"github.com/vngcloud/vcontainer-sdk/vcontainer"
	"github.com/vngcloud/vcontainer-sdk/vcontainer/services/identity/v2/extensions/oauth2"
	"github.com/vngcloud/vcontainer-sdk/vcontainer/services/identity/v2/tokens"
	"testing"
)

func TestListBySubnetID(t *testing.T) {
	clientID := "7ae234ed-9572-4903-8455-40888c554ac9"
	clientSecret := "5648cd7b-24c8-41b3-b0ef-c70633d88ccb"
	projectID := "pro-462803f3-6858-466f-bf05-df2b33faa360"
	subnetID := "sub-403b36d2-39fc-47c4-b40b-8df0ecb71045"

	identityURL := "https://iamapis.vngcloud.vn/accounts-api/v2"
	provider, _ := vcontainer.NewClient(identityURL)
	vcontainer.Authenticate(provider, &oauth2.AuthOptions{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		AuthOptionsBuilder: &tokens.AuthOptions{
			IdentityEndpoint: identityURL,
		},
	})

	vlb, _ := vcontainer.NewServiceClient(
		"https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway/v2",
		provider,
		"vlb")

	opt := new(ListBySubnetIDOptsBuilder)
	opt.ProjectID = projectID
	opt.SubnetID = subnetID

	resp, err := ListBySubnetID(vlb, opt)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	fmt.Printf("%v\n", resp)
}
