package obj

type Cluster struct {
	ID                          string
	VpcID                       string
	SubnetID                    string
	MasterClusterSecGroupIDList []string
	MinionClusterSecGroupIDList []string
}
