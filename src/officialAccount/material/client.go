package material

import (
	"github.com/ArtisanCloud/PowerLibs/object"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel"
	response2 "github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
	request2 "github.com/ArtisanCloud/PowerWeChat/src/officialAccount/material/request"
	response3 "github.com/ArtisanCloud/PowerWeChat/src/officialAccount/material/response"
)

type Client struct {
	*kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) *Client {
	return &Client{
		kernel.NewBaseClient(&app, nil),
	}
}

// 新增永久图文素材
// https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Adding_Permanent_Assets.html
func (comp *Client) AddNews(options *request2.RequestMaterialAddNews) (*response3.ResponseMaterialAddNews, error) {

	result := &response3.ResponseMaterialAddNews{}

	_, err := comp.HttpPostJson("cgi-bin/material/add_news", options, nil, nil, result)

	return result, err
}

// 新增其他类型永久素材
// https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Adding_Permanent_Assets.html
func (comp *Client) AddMaterial(options *request2.RequestMaterialAddMaterial) (*response3.ResponseMaterialAddMaterial, error) {

	result := &response3.ResponseMaterialAddMaterial{}

	jsonDescription, err := object.JsonEncode(&object.HashMap{
		"title":        options.Title,
		"introduction": options.Introduction,
	})
	if err != nil {
		return nil, err
	}
	body := &object.HashMap{
		"type":        options.Type,
		"media":       options.Media,
		"Description": jsonDescription,
	}

	_, err = comp.HttpPostJson("cgi-bin/material/add_material", body, nil, nil, result)

	return result, err
}

// 获取永久素材
// https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Getting_Permanent_Assets.html
func (comp *Client) GetMaterial(mediaID int) (*response3.ResponseMaterialGetMaterial, error) {

	result := &response3.ResponseMaterialGetMaterial{}

	options := &object.HashMap{
		"media_id": mediaID,
	}

	_, err := comp.HttpPostJson("cgi-bin/material/get_material", options, nil, nil, result)

	return result, err
}

// 删除永久素材
// https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Deleting_Permanent_Assets.html
func (comp *Client) DelMaterial(mediaID int) (*response2.ResponseOfficialAccount, error) {

	result := &response2.ResponseOfficialAccount{}

	options := &object.HashMap{
		"media_id": mediaID,
	}

	_, err := comp.HttpPostJson("cgi-bin/material/del_material", options, nil, nil, result)

	return result, err
}

// 修改永久图文素材
// https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Editing_Permanent_Rich_Media_Assets.html
func (comp *Client) UpdateNews(options *request2.RequestMaterialUpdateNews) (*response2.ResponseOfficialAccount, error) {

	result := &response2.ResponseOfficialAccount{}

	_, err := comp.HttpPostJson("cgi-bin/material/update_news", options, nil, nil, result)

	return result, err
}

// 获取素材总数
// https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Get_the_total_of_all_materials.html
func (comp *Client) GetMaterialCount() (*response3.ResponseMaterialGetMaterialCount, error) {

	result := &response3.ResponseMaterialGetMaterialCount{}

	_, err := comp.HttpPostJson("cgi-bin/material/get_materialcount", nil, nil, nil, result)

	return result, err

}

// 获取素材列表
// https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Get_materials_list.html
func (comp *Client) BatchGetMaterial(options *request2.RequestMaterialBatchGetMaterial) (*response3.ResponseMaterialBatchGetMaterial, error) {

	result := &response3.ResponseMaterialBatchGetMaterial{}

	_, err := comp.HttpPostJson("cgi-bin/material/batchget_material", options, nil, nil, result)

	return result, err
}
