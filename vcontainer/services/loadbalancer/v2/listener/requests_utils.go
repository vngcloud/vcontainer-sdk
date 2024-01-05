package listener

func NewCreateOpts(pProjectID, pLbID string, pOpts *CreateOpts) ICreateOptsBuilder {
	pOpts.ProjectID = pProjectID
	pOpts.LoadBalancerID = pLbID
	return pOpts
}

func NewGetOpts(pProjectID, pLbID string) IGetOptsBuilder {
	opts := new(GetOpts)
	opts.ProjectID = pProjectID
	opts.LoadBalancerID = pLbID
	return opts
}
