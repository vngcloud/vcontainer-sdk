package secgroup_rule

import (
	"fmt"
	improc "github.com/imroc/req/v3"
	"testing"
)

func TestCreate(t *testing.T) {
	token := "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICJlckVaZFpkNkRsc21pdjhsMDZIaVB3bHZYWnotLVlGYXlZcVJiczlxc09rIn0.eyJleHAiOjE3MDQ2NDc0MDIsImlhdCI6MTcwNDY0NTYwMiwianRpIjoiZTZjYTNlODQtNjcxOS00YzQ0LWExOGYtNTRjOGIyODQ1YzYzIiwiaXNzIjoiaHR0cHM6Ly9zaWduaW4udm5nY2xvdWQudm4vYXV0aC9yZWFsbXMvaWFtIiwic3ViIjoiNTlmOTZjYzUtZjg5MS00ZWExLTliMGYtMWNjMGE2MmM4ZjMzIiwidHlwIjoiQmVhcmVyIiwiYXpwIjoiN2FlMjM0ZWQtOTU3Mi00OTAzLTg0NTUtNDA4ODhjNTU0YWM5Iiwic2NvcGUiOiIiLCJjbGllbnRIb3N0IjoiMTAuMTY2LjIuOSIsImF1dGhBY2NvdW50SWQiOjYwMTA4LCJjbGllbnRBZGRyZXNzIjoiMTAuMTY2LjIuOSIsImNsaWVudF9pZCI6IjdhZTIzNGVkLTk1NzItNDkwMy04NDU1LTQwODg4YzU1NGFjOSIsImF1dGhVc2VyVHlwZSI6InVzZXItc2EifQ.bnMI7sBoZM16dureTfANpOrp4EsC4eJ7F2Q99h4xjlPP0pnyqOjm92NmJdw_yV4MmoJdOE72mzkYh8WGqPyGzn2CAyqWqqyo6bV_0a4UbAuMGJMCytYyrT-m8wfC7KoR-_53_njlyvkt8B8imiM6JOwYS8nuFp8cNhxB1edI7j_X5sLOmdxpdDiLt6u7HRDpzHpoHZLkHzLikECcIhCvlnFFFUcMNCLU6LRxy4jSQedA18itmXYCaeuwBzNZ16vjxWtlvPO6BSSfYP55km0cIU5WJ3cOGPOo3cv1wzSG8xrI1LpyJSDi1QicvOG98ASLfZfSC4gtI1TPOHe2uxGHaw"

	req := NewCreateOpts("pro-462803f3-6858-466f-bf05-df2b33faa360", &CreateOpts{
		Direction:      "ingress",
		EtherType:      "IPv4",
		Protocol:       "tcp",
		PortRangeMax:   3000,
		PortRangeMin:   3000,
		RemoteIPPrefix: "10.21.0.0/24",
		Description:    "test",
		SecgroupUUID:   "secg-60726e94-07da-4eb9-8e12-bdbff3dc6544",
	})

	client := improc.NewClient()
	resp := client.Post("https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway/v2/pro-462803f3-6858-466f-bf05-df2b33faa360/secgroups/secg-60726e94-07da-4eb9-8e12-bdbff3dc6544/secgroupRules").
		SetHeader("Content-Type", "application/json").SetHeader("Authorization", "Bearer "+token).SetBodyJsonMarshal(req).Do()

	fmt.Printf("%+v\n", resp.StatusCode)
	fmt.Printf("%+v\n", resp.Response)

}
