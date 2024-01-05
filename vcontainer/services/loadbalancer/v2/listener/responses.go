package listener

import "github.com/vngcloud/vcontainer-sdk/vcontainer/services/loadbalancer/v2/listener/obj"

type CreateResponse struct {
}

func (s *CreateResponse) ToListenerObject() *obj.Listener {
	return nil
}

// **************************************************** GetResponse ****************************************************

type GetBasedLoadBalancerResponse struct {
	Data []struct {
		UUID                            string   `json:"uuid"`
		Name                            string   `json:"name"`
		Description                     *string  `json:"description"`
		Protocol                        string   `json:"protocol"`
		ProtocolPort                    int      `json:"protocolPort"`
		ConnectionLimit                 int      `json:"connectionLimit"`
		DefaultPoolId                   string   `json:"defaultPoolId"`
		DefaultPoolName                 string   `json:"defaultPoolName"`
		TimeoutClient                   int      `json:"timeoutClient"`
		TimeoutMember                   int      `json:"timeoutMember"`
		TimeoutConnection               int      `json:"timeoutConnection"`
		AllowedCidrs                    string   `json:"allowedCidrs"`
		Headers                         []string `json:"headers"`
		CertificateAuthorities          []string `json:"certificateAuthorities"`
		DisplayStatus                   string   `json:"displayStatus"`
		CreatedAt                       string   `json:"createdAt"`
		UpdatedAt                       string   `json:"updatedAt"`
		DefaultCertificateAuthority     *string  `json:"defaultCertificateAuthority"`
		ClientCertificateAuthentication *string  `json:"clientCertificateAuthentication"`
		ProgressStatus                  string   `json:"progressStatus"`
	} `json:"data"`
}

func (s *GetBasedLoadBalancerResponse) ToListListenerObject() []*obj.Listener {
	listeners := make([]*obj.Listener, 0, len(s.Data))
	for i := range s.Data {
		listeners = append(listeners, s.ToListenerObjectAt(i))
	}
	return listeners
}

func (s *GetBasedLoadBalancerResponse) ToListenerObjectAt(i int) *obj.Listener {
	item := s.Data[i]
	return &obj.Listener{
		ID:           item.UUID,
		Status:       item.DisplayStatus,
		Protocol:     item.Protocol,
		ProtocolPort: item.ProtocolPort,
	}
}
