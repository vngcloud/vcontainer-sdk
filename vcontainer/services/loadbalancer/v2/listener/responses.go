package listener

import "github.com/vngcloud/vcontainer-sdk/vcontainer/services/loadbalancer/v2/listener/obj"

type CreateResponse struct {
	UUID string `json:"uuid"`
}

func (s *CreateResponse) ToListenerObject() *obj.Listener {
	if s == nil {
		return nil
	}

	return &obj.Listener{
		ID: s.UUID,
	}
}

// **************************************************** GetResponse ****************************************************

type GetBasedLoadBalancerResponse struct {
	Data []struct {
		UUID                            string   `json:"uuid"`
		Name                            string   `json:"name"`
		Description                     string   `json:"description,omitempty"`
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
	var listeners []*obj.Listener

	if s == nil || s.Data == nil || len(s.Data) == 0 {
		return listeners
	}

	for _, itemListener := range s.Data {
		listeners = append(listeners, &obj.Listener{
			ID:           itemListener.UUID,
			Name:         itemListener.Name,
			Protocol:     itemListener.Protocol,
			ProtocolPort: itemListener.ProtocolPort,
			Status:       itemListener.DisplayStatus,
		})
	}
	return listeners
}
