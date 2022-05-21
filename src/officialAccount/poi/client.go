package poi

import (
	"github.com/ArtisanCloud/PowerLibs/object"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/src/officialAccount/poi/reponse"
)

type Client struct {
	*kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) *Client {
	return &Client{
		kernel.NewBaseClient(&app, nil),
	}
}

// 查询门店信息
// https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Store_Interface.html
func (comp *Client) Get(poiID int) (*reponse.ResponsePOIGet, error) {

	result := &reponse.ResponsePOIGet{}

	_, err := comp.HttpPostJson("cgi-bin/poi/getpoi", &object.HashMap{"poi_id": poiID}, nil, nil, result)

	return result, err
}

// 查询门店列表
// https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Store_Interface.html
func (comp *Client) List(offset int, limit int) (*reponse.ResponsePOIList, error) {

	result := &reponse.ResponsePOIList{}

	params := &object.HashMap{
		"begin": offset,
		"limit": limit,
	}

	_, err := comp.HttpPostJson("cgi-bin/poi/getpoilist", params, nil, nil, result)

	return result, err
}

// 查询门店列表
// https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Store_Interface.html
func (comp *Client) Create(data *reponse.BusinessInfo) (*reponse.ResponsePOIList, error) {

	result := &reponse.ResponsePOIList{}

	params := &object.HashMap{
		"business": &object.HashMap{
			"base_info": data,
		},
	}

	_, err := comp.HttpPostJson("cgi-bin/poi/addpoi", params, nil, nil, result)

	return result, err
}
