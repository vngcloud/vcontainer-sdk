package subnet

func NewGetOpts(pProjectID, pSubnetUUID string) IGetOptsBuilder {
	opts := new(GetOpts)
	opts.ProjectID = pProjectID
	opts.CommonSubnetUUID = pSubnetUUID
	return opts
}
