package v2

type LoadBalancerV2Common struct {
	LoadBalancerID string
}

func (s *LoadBalancerV2Common) GetLoadBalancerID() string {
	return s.LoadBalancerID
}
