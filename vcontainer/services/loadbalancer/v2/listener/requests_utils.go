package listener

func NewCreateOpts(pProjectID, pLbID string, pOpts *CreateOpts) ICreateOptsBuilder {
	pOpts.ProjectID = pProjectID
	pOpts.LoadBalancerID = pLbID
	return pOpts
}
