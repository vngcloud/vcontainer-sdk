package volume

import (
	"github.com/vngcloud/vcontainer-sdk/client"
	"github.com/vngcloud/vcontainer-sdk/vcontainer/pagination"
	"github.com/vngcloud/vcontainer-sdk/vcontainer/services/blockstorage/v2/volume/obj"
)

func List(pSc *client.ServiceClient, pOpts IListOptsBuilder) *pagination.Pager {
	qp, err := pOpts.ToListQuery()
	url := listURL(pSc, pOpts)
	if err == nil {
		url = url + qp
	}
	return pagination.NewPager(pSc, url, pOpts,
		func() interface{} {
			return NewListResponse()
		},
		func(r interface{}) pagination.IPage {
			resp := r.(*ListResponse)
			return resp
		})
}

func Create(pSc *client.ServiceClient, pOpts ICreateOptsBuilder) (*obj.Volume, error) {
	response := NewCreateResponse()
	body := pOpts.ToRequestBody()
	_, err := pSc.Post(createURL(pSc, pOpts), &client.RequestOpts{
		JSONBody:     body,
		JSONResponse: response,
		OkCodes:      []int{202},
	})

	if err != nil {
		return nil, err
	}

	return response.ToVolumeObject(), nil
}

func Delete(pSc *client.ServiceClient, pOpts IDeleteOptsBuilder) error {
	_, err := pSc.Delete(deleteURL(pSc, pOpts), &client.RequestOpts{
		OkCodes: []int{202},
	})

	return err
}

func Get(pSc *client.ServiceClient, pOpts IGetOptsBuilder) (*obj.Volume, error) {
	response := NewGetResponse()
	_, err := pSc.Get(getURL(pSc, pOpts), &client.RequestOpts{
		JSONResponse: response,
		OkCodes:      []int{200},
	})

	if err != nil {
		return nil, err
	}

	return response.ToVolumeObject(), nil
}
