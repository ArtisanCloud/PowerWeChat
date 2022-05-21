package poi

import (
	"github.com/ArtisanCloud/PowerLibs/object"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel"
	response2 "github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/src/officialAccount/poi/request"
	"github.com/ArtisanCloud/PowerWeChat/src/officialAccount/poi/response"
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
func (comp *Client) Get(poiID int) (*response.ResponsePOIGet, error) {

	result := &response.ResponsePOIGet{}

	_, err := comp.HttpPostJson("cgi-bin/poi/getpoi", &object.HashMap{"poi_id": poiID}, nil, nil, result)

	return result, err
}

// 查询门店列表
// https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Store_Interface.html
func (comp *Client) List(offset int, limit int) (*response.ResponsePOIList, error) {

	result := &response.ResponsePOIList{}

	params := &object.HashMap{
		"begin": offset,
		"limit": limit,
	}

	_, err := comp.HttpPostJson("cgi-bin/poi/getpoilist", params, nil, nil, result)

	return result, err
}

// 查询门店列表
// https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Store_Interface.html
func (comp *Client) Create(data *request.BusinessInfo) (*response2.ResponseOfficialAccount, error) {

	result := &response2.ResponseOfficialAccount{}

	params := &object.HashMap{
		"business": &object.HashMap{
			"base_info": data,
		},
	}

	_, err := comp.HttpPostJson("cgi-bin/poi/addpoi", params, nil, nil, result)

	return result, err
}

// 查询门店列表并获取POI ID
// https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Store_Interface.html
func (comp *Client) CreateAndGetID(data *request.BusinessInfo) (int, error) {

	result, err := comp.Create(data)

	poi, err := comp.DetectAndCastResponseToType(result, "array")
	if err != nil {
		return 0, err
	}

	mapPOI := poi.(*object.HashMap)
	poiID := (*mapPOI)["poi_id"].(int)

	return poiID, err
}


// 修改门店服务信息
// https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Store_Interface.html
func (comp *Client) Update(data *request.BusinessInfo) (*response2.ResponseOfficialAccount, error) {

	result := &response2.ResponseOfficialAccount{}

	params := &object.HashMap{
		"business": &object.HashMap{
			"base_info": data,
		},
	}

	_, err := comp.HttpPostJson("cgi-bin/poi/updatepoi", params, nil, nil, result)

	return result, err
}

// 删除门店
// https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Store_Interface.html
func (comp *Client) Delete(poiID int) (*response.ResponsePOIGet, error) {

	result := &response.ResponsePOIGet{}

	_, err := comp.HttpPostJson("cgi-bin/poi/delpoi", &object.HashMap{"poi_id": poiID}, nil, nil, result)

	return result, err
}