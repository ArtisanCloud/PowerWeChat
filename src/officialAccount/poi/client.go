package poi

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	response2 "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/poi/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/poi/response"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) (*Client, error) {
	baseClient, err := kernel.NewBaseClient(&app, nil)
	if err != nil {
		return nil, err
	}
	return &Client{
		baseClient,
	}, nil
}

// 查询门店信息
// https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Store_Interface.html
func (comp *Client) Get(ctx *context.Context, poiID int) (*response.ResponsePOIGet, error) {

	result := &response.ResponsePOIGet{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/poi/getpoi", &object.HashMap{"poi_id": poiID}, nil, nil, result)

	return result, err
}

// 查询门店列表
// https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Store_Interface.html
func (comp *Client) List(ctx *context.Context, offset int, limit int) (*response.ResponsePOIList, error) {

	result := &response.ResponsePOIList{}

	params := &object.HashMap{
		"begin": offset,
		"limit": limit,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/poi/getpoilist", params, nil, nil, result)

	return result, err
}

// 查询门店列表
// https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Store_Interface.html
func (comp *Client) Create(ctx *context.Context, data *request.BusinessInfo) (*response2.ResponseOfficialAccount, error) {

	result := &response2.ResponseOfficialAccount{}

	params := &object.HashMap{
		"business": &object.HashMap{
			"base_info": data,
		},
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/poi/addpoi", params, nil, nil, result)

	return result, err
}

// 查询门店列表并获取POI ID
// https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Store_Interface.html
func (comp *Client) CreateAndGetID(ctx *context.Context, data *request.BusinessInfo) (int, error) {

	result, err := comp.Create(ctx, data)

	poi, err := comp.BaseClient.DetectAndCastResponseToType(result, response2.TYPE_MAP)
	if err != nil {
		return 0, err
	}

	mapPOI := poi.(*object.HashMap)
	poiID := (*mapPOI)["poi_id"].(int)

	return poiID, err
}

// 修改门店服务信息
// https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Store_Interface.html
func (comp *Client) Update(ctx *context.Context, data *request.BusinessInfo) (*response2.ResponseOfficialAccount, error) {

	result := &response2.ResponseOfficialAccount{}

	params := &object.HashMap{
		"business": &object.HashMap{
			"base_info": data,
		},
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/poi/updatepoi", params, nil, nil, result)

	return result, err
}

// 删除门店
// https://developers.weixin.qq.com/doc/offiaccount/WeChat_Stores/WeChat_Store_Interface.html
func (comp *Client) Delete(ctx *context.Context, poiID int) (*response.ResponsePOIGet, error) {

	result := &response.ResponsePOIGet{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/poi/delpoi", &object.HashMap{"poi_id": poiID}, nil, nil, result)

	return result, err
}
