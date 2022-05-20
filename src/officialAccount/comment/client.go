package comment

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel"
	response2 "github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
	request2 "github.com/ArtisanCloud/PowerWeChat/src/officialAccount/material/request"
	response3 "github.com/ArtisanCloud/PowerWeChat/src/officialAccount/material/response"
)

type Client struct {
	*kernel.BaseClient
}


// 新增永久图文素材
// https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Adding_Permanent_Assets.html
func (comp *Client) AddNews(options *request2.RequestMaterialAddNews) (*response3.ResponseMaterialAddNews, error) {

	result := &response3.ResponseMaterialAddNews{}

	_, err := comp.HttpPostJson("cgi-bin/material/add_news", options, nil, nil, result)

	return result, err
}


// 修改永久图文素材
// https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Editing_Permanent_Rich_Media_Assets.html
func (comp *Client) UpdateNews(options *request2.RequestMaterialUpdateNews) (*response2.ResponseOfficialAccount, error) {

	result := &response2.ResponseOfficialAccount{}

	_, err := comp.HttpPostJson("cgi-bin/material/update_news", options, nil, nil, result)

	return result, err
}