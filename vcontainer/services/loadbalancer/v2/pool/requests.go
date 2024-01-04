package pool

import (
	"github.com/vngcloud/vcontainer-sdk/vcontainer/services/common"
	lbCm "github.com/vngcloud/vcontainer-sdk/vcontainer/services/loadbalancer/v2"
)

type (
	CreateOptsAlgorithmOpt string
	CreateOptsProtocolOpt  string
)

const (
	CreateOptsAlgorithmOptRoundRobin CreateOptsAlgorithmOpt = "ROUND_ROBIN"
	CreateOptsAlgorithmOptLeastConn  CreateOptsAlgorithmOpt = "LEAST_CONNECTIONS"
	CreateOptsAlgorithmOptSourceIP   CreateOptsAlgorithmOpt = "SOURCE_IP"

	CreateOptsProtocolOptTCP   CreateOptsProtocolOpt = "TCP"
	CreateOptsProtocolOptUDP   CreateOptsProtocolOpt = "UDP"
	CreateOptsProtocolOptHTTP  CreateOptsProtocolOpt = "HTTP"
	CreateOptsProtocolOptProxy CreateOptsProtocolOpt = "PROXY"
)

type CreateOpts struct {
	Algorithm    CreateOptsAlgorithmOpt `json:"algorithm"`
	PoolName     string                 `json:"poolName"`
	PoolProtocol CreateOptsProtocolOpt  `json:"poolProtocol"`

	common.CommonOpts
	lbCm.LoadBalancerV2Common
}

func (s *CreateOpts) ToRequestBody() interface{} {
	return s
}
