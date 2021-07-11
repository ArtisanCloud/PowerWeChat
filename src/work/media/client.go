package media

import (
	"github.com/ArtisanCloud/go-libs/http/contract"
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type Client struct {
	*kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) *Client {
	return &Client{
		kernel.NewBaseClient(&app, nil),
	}
}

func (comp *Client) Get(mediaID string) contract.ResponseContract {

	result := &response.ResponseWX{}
	rawResponse := comp.RequestRaw("cgi-bin/media/get", "GET", &object.HashMap{
		"query": object.StringMap{
			"media_id": mediaID,
		},
	}, result)

	response:= rawResponse.(contract.ResponseContract)
	header := (*response.GetHeaders()).Header()
	if header.Get("Content-Type") != "text/plain"{
		//config := (*comp.App).GetConfig()
		//return comp.CastResponseToType(response,config.Get("response_type", "array"))
		return nil
	}

	return response
}
