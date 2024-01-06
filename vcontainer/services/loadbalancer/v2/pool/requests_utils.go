package pool

func NewCreateOpts(pProjectID, pLbID string, pOpts *CreateOpts) ICreateOptsBuilder {
	pOpts.ProjectID = pProjectID
	pOpts.LoadBalancerID = pLbID
	return pOpts
}

func NewListPoolsBasedLoadBalancerOpts(pProjectID, pLbID string) IListPoolsBasedLoadBalancerOptsBuilder {
	opts := new(ListPoolsBasedLoadBalancerOpts)
	opts.ProjectID = pProjectID
	opts.LoadBalancerID = pLbID
	return opts
}

func NewDeleteOpts(pProjectID, pLbID, pPoolID string) IDeleteOptsBuilder {
	opts := new(DeleteOpts)
	opts.ProjectID = pProjectID
	opts.LoadBalancerID = pLbID
	opts.PoolID = pPoolID
	return opts
}
