package shakeAround

import (
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/shakeAround/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/shakeAround/response"
)

type PageClient struct {
	*kernel.BaseClient
}

func NewPageClient(app kernel.ApplicationInterface) (*PageClient, error) {
	baseClient, err := kernel.NewBaseClient(&app, nil)
	if err != nil {
		return nil, err
	}
	return &PageClient{
		baseClient,
	}, nil
}

// 新增摇一摇出来的页面信息
// https://developers.weixin.qq.com/doc/offiaccount/Shake_Nearby/Pages_management/Page_management.html
func (comp *PageClient) Create(data *request.RequestPageInfo) (*response.ResponsePage, error) {

	result := &response.ResponsePage{}

	_, err := comp.HttpPostJson("shakearound/page/add", data, nil, nil, result)

	return result, err
}

// 编辑摇一摇出来的页面信息
// https://developers.weixin.qq.com/doc/offiaccount/Shake_Nearby/Pages_management/Edit_page_information.html
func (comp *PageClient) Update(data request.RequestPageUpdate) (*response.ResponsePage, error) {

	result := &response.ResponsePage{}

	_, err := comp.HttpPostJson("shakearound/page/update", data, nil, nil, result)

	return result, err
}

// 更新设备分组
// https://developers.weixin.qq.com/doc/offiaccount/Shake_Nearby/Pages_management/Query_page_list.html
func (comp *PageClient) List(begin int, count int) (*response.ResponsePageList, error) {

	result := &response.ResponsePageList{}

	params := &object.HashMap{
		"type":  2,
		"begin": begin,
		"count": count,
	}

	_, err := comp.HttpPostJson("shakearound/page/search", params, nil, nil, result)

	return result, err
}

// 删除已有的页面
// https://developers.weixin.qq.com/doc/offiaccount/Shake_Nearby/Pages_management/Delete_page.html
func (comp *PageClient) Delete(pageID string) (*response.ResponsePage, error) {

	result := &response.ResponsePage{}

	params := &object.HashMap{
		"page_id": pageID,
	}

	_, err := comp.HttpPostJson("shakearound/page/delete", params, nil, nil, result)

	return result, err
}
