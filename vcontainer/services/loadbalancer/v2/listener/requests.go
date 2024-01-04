package listener

import (
	"github.com/vngcloud/vcontainer-sdk/vcontainer/services/common"
	lbCm "github.com/vngcloud/vcontainer-sdk/vcontainer/services/loadbalancer/v2"
)

type (
	CreateOptsListenerProtocolOpt string
)

const (
	CreateOptsListenerProtocolOptTCP   CreateOptsListenerProtocolOpt = "TCP"
	CreateOptsListenerProtocolOptUDP   CreateOptsListenerProtocolOpt = "UDP"
	CreateOptsListenerProtocolOptHTTP  CreateOptsListenerProtocolOpt = "HTTP"
	CreateOptsListenerProtocolOptHTTPS CreateOptsListenerProtocolOpt = "HTTPS"
)

type CreateOpts struct {
	AllowedCidrs         string                        `json:"allowedCidrs"`
	DefaultPoolId        string                        `json:"defaultPoolId"`
	ListenerName         string                        `json:"listenerName"`
	ListenerProtocol     CreateOptsListenerProtocolOpt `json:"listenerProtocol"`
	ListenerProtocolPort int                           `json:"listenerProtocolPort"`
	TimeoutClient        int                           `json:"timeoutClient"`
	TimeoutConnection    int                           `json:"timeoutConnection"`
	TimeoutMember        int                           `json:"timeoutMember"`

	common.CommonOpts
	lbCm.LoadBalancerV2Common
}

func (s *CreateOpts) ToRequestBody() interface{} {
	return s
}
