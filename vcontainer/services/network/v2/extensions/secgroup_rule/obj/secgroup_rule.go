package obj

type SecgroupRule struct {
	UUID           string
	SecgroupUUID   string
	Direction      string
	EtherType      string
	PortRangeMax   int
	PortRangeMin   int
	Protocol       string
	Description    string
	RemoteIPPrefix string
}
