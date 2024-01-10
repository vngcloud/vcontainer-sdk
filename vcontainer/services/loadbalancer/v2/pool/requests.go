package pool

import (
	"github.com/vngcloud/vcontainer-sdk/vcontainer/services/common"
	lbCm "github.com/vngcloud/vcontainer-sdk/vcontainer/services/loadbalancer/v2"
)

// **************************************************** CreateOpts *****************************************************

type (
	CreateOptsAlgorithmOpt           string
	CreateOptsProtocolOpt            string
	CreateOptsHealthCheckProtocolOpt string
)

const (
	CreateOptsAlgorithmOptRoundRobin CreateOptsAlgorithmOpt = "ROUND_ROBIN"
	CreateOptsAlgorithmOptLeastConn  CreateOptsAlgorithmOpt = "LEAST_CONNECTIONS"
	CreateOptsAlgorithmOptSourceIP   CreateOptsAlgorithmOpt = "SOURCE_IP"

	CreateOptsProtocolOptTCP   CreateOptsProtocolOpt = "TCP"
	CreateOptsProtocolOptUDP   CreateOptsProtocolOpt = "UDP"
	CreateOptsProtocolOptHTTP  CreateOptsProtocolOpt = "HTTP"
	CreateOptsProtocolOptProxy CreateOptsProtocolOpt = "PROXY"

	CreateOptsHealthCheckProtocolOptTCP     CreateOptsHealthCheckProtocolOpt = "TCP"
	CreateOptsHealthCheckProtocolOptHTTP    CreateOptsHealthCheckProtocolOpt = "HTTP"
	CreateOptsHealthCheckProtocolOptHTTPs   CreateOptsHealthCheckProtocolOpt = "HTTPS"
	CreateOptsHealthCheckProtocolOptPINGUDP CreateOptsHealthCheckProtocolOpt = "PING-UDP"
)

type (
	CreateOpts struct {
		Algorithm     CreateOptsAlgorithmOpt `json:"algorithm"`
		PoolName      string                 `json:"poolName"`
		PoolProtocol  CreateOptsProtocolOpt  `json:"poolProtocol"`
		HealthMonitor HealthMonitor          `json:"healthMonitor"`
		Members       []Member               `json:"members"`

		common.CommonOpts
		lbCm.LoadBalancerV2Common
	}

	Member struct {
		Backup      bool   `json:"backup"`
		IpAddress   string `json:"ipAddress"`
		MonitorPort int    `json:"monitorPort"`
		Name        string `json:"name"`
		Port        int    `json:"port"`
		Weight      int    `json:"weight"`
	}

	HealthMonitor struct {
		HealthCheckProtocol CreateOptsHealthCheckProtocolOpt `json:"healthCheckProtocol"`
		HealthCheckPath     string                           `json:"healthCheckPath,omitempty"`
		HealthyThreshold    int                              `json:"healthyThreshold"`
		UnhealthyThreshold  int                              `json:"unhealthyThreshold"`
		Interval            int                              `json:"interval"`
		Timeout             int                              `json:"timeout"`
	}
)

func (s *CreateOpts) ToRequestBody() interface{} {
	// If health check protocol is TCP, health check path must be empty
	if s.HealthMonitor.HealthCheckProtocol == CreateOptsHealthCheckProtocolOptTCP {
		s.HealthMonitor.HealthCheckPath = ""
	}

	return s
}

// ****************************************** ListPoolsBasedLoadBalancerOpts *******************************************

type ListPoolsBasedLoadBalancerOpts struct {
	common.CommonOpts
	lbCm.LoadBalancerV2Common
}

// ***************************************************** DeleteOpts ****************************************************

type DeleteOpts struct {
	common.CommonOpts
	lbCm.LoadBalancerV2Common
	lbCm.PoolV2Common
}
