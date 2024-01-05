package listener

func NewCreateOpts(pProjectID, pLbID string, pOpts *CreateOpts) ICreateOptsBuilder {
	pOpts.ProjectID = pProjectID
	pOpts.LoadBalancerID = pLbID
	return pOpts
}

func NewGetBasedLoadBalancerOpts(pProjectID, pLbID string) IGetBasedLoadBalancerOptsBuilder {
	opts := new(GetBasedLoadBalancerOpts)
	opts.ProjectID = pProjectID
	opts.LoadBalancerID = pLbID
	return opts
}
