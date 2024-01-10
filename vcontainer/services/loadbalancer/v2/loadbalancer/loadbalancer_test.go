package loadbalancer

import (
	"fmt"
	"github.com/vngcloud/vcontainer-sdk/vcontainer"
	"github.com/vngcloud/vcontainer-sdk/vcontainer/services/identity/v2/extensions/oauth2"
	"github.com/vngcloud/vcontainer-sdk/vcontainer/services/identity/v2/tokens"
	"testing"
)

func TestListBySubnetID(t *testing.T) {
	clientID := ""
	clientSecret := ""
	projectID := "pro-xxxxxxxxx-6858-466f-bf05-df2b33faa360"
	subnetID := "sub-xxxxxxxx-39fc-47c4-b40b-8df0ecb71045"

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
