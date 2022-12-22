package license

import (
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	response2 "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/openPlatform/codeTemplate/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/openWork/auth"
)

type Account struct {
	*kernel.BaseClient
}

func NewAccount(app *kernel.ApplicationInterface) (*Account, error) {
	token := (*app).GetComponent("ProviderAccessToken").(*auth.AccessToken)
	baseAccount, err := kernel.NewBaseClient(app, token.AccessToken)
	if err != nil {
		return nil, err
	}
	return &Account{
		baseAccount,
	}, nil
}

// 获取代码草稿列表
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/code_template/gettemplatedraftlist.html
func (comp *Account) GetDrafts() (*response.ResponseGetDrafts, error) {

	result := &response.ResponseGetDrafts{}

	_, err := comp.HttpGet("wxa/gettemplatedraftlist", nil, nil, result)

	return result, err

}

// 将草稿添加到代码模板库
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/code_template/addtotemplate.html#请求地址
func (comp *Account) CreateFromDraft(draftID int, templateType int) (*response2.ResponseOpenPlatform, error) {

	result := &response2.ResponseOpenPlatform{}

	params := &object.HashMap{
		"draft_id":      draftID,
		"template_type": templateType,
	}

	_, err := comp.HttpPostJson("wxa/addtotemplate", params, nil, nil, result)

	return result, err

}

// 获取代码模板列表
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/code_template/gettemplatelist.html#请求地址
func (comp *Account) List() (*response.ResponseList, error) {

	result := &response.ResponseList{}

	_, err := comp.HttpGet("wxa/gettemplatelist", nil, nil, result)

	return result, err

}

// 删除指定代码模板
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/code_template/deletetemplate.html#请求地址
func (comp *Account) Delete(templateID string) (*response2.ResponseOpenPlatform, error) {

	result := &response2.ResponseOpenPlatform{}

	params := &object.HashMap{
		"template_id": templateID,
	}

	_, err := comp.HttpPostJson("wxa/deletetemplate", params, nil, nil, result)

	return result, err

}
