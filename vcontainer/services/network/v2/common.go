package v2

type SecgroupV2Common struct {
	SecgroupUUID string
}

func (s *SecgroupV2Common) GetSecgroupUUID() string {
	return s.SecgroupUUID
}
