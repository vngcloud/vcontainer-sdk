package loadbalancer

func NewCreateOpts(pProjectID string, pOpts *CreateOpts) ICreateOptsBuilder {
	pOpts.ProjectID = pProjectID
	return pOpts
}

func NewGetOpts(pProjectID, pLoadBalancerID string) IGetOptsBuilder {
	opts := new(GetOpts)
	opts.ProjectID = pProjectID
	opts.LoadBalancerID = pLoadBalancerID
	return opts
}

func NewDeleteOpts(pProjectID, pLoadBalancerID string) IDeleteOptsBuilder {
	opts := new(DeleteOpts)
	opts.ProjectID = pProjectID
	opts.LoadBalancerID = pLoadBalancerID
	return opts
}

func NewGetBySubnetIDOpts(pProjectID, pSubnetID string) IListBySubnetIDOptsBuilder {
	opts := new(ListBySubnetIDOptsBuilder)
	opts.ProjectID = pProjectID
	opts.SubnetID = pSubnetID
	return opts
}
