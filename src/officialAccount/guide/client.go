package guide

import (
	"errors"
	"github.com/ArtisanCloud/PowerLibs/object"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel"
	response2 "github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/src/officialAccount/guide/request"
	"github.com/ArtisanCloud/PowerWeChat/src/officialAccount/guide/response"
)

type Client struct {
	*kernel.BaseClient
}

// 为服务号添加顾问
// https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.addGuideAcct.html
func (comp *Client) CreateAdviser(guideAccount string, guideOpenID string, guideHeadImgURL string, guideNickname string) (*response2.ResponseOfficialAccount, error) {

	result := &response2.ResponseOfficialAccount{}

	params := &object.HashMap{}
	params, err := comp.SelectAccountAndOpenID(params, guideAccount, guideOpenID)
	if err != nil {
		return nil, err
	}

	if guideHeadImgURL != "" {
		(*params)["guide_headimgurl"] = guideHeadImgURL
	}
	if guideNickname != "" {
		(*params)["guide_nickname"] = guideNickname
	}
	_, err = comp.HttpPostJson("cgi-bin/guide/addguideacct", params, nil, nil, result)

	return result, err
}

// 获取顾问信息
// https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.getGuideAcct.html
func (comp *Client) GetAdviser(guideAccount string, guideOpenID string) (*response.ResponseGuideGetAdviser, error) {

	result := &response.ResponseGuideGetAdviser{}

	params := &object.HashMap{}
	params, err := comp.SelectAccountAndOpenID(params, guideAccount, guideOpenID)
	if err != nil {
		return nil, err
	}
	_, err = comp.HttpPostJson("cgi-bin/guide/getguideacct", params, nil, nil, result)

	return result, err
}

// 修改顾问的昵称或头像
// https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.updateGuideAcct.html
func (comp *Client) UpdateAdviser(guideAccount string, guideOpenID string, guideHeadImgURL string, guideNickname string) (*response2.ResponseOfficialAccount, error) {

	result := &response2.ResponseOfficialAccount{}

	params := &object.HashMap{}
	params, err := comp.SelectAccountAndOpenID(params, guideAccount, guideOpenID)
	if err != nil {
		return nil, err
	}

	if guideHeadImgURL != "" {
		(*params)["guide_headimgurl"] = guideHeadImgURL
	}
	if guideNickname != "" {
		(*params)["guide_nickname"] = guideNickname
	}
	_, err = comp.HttpPostJson("cgi-bin/guide/updateguideacct", params, nil, nil, result)

	return result, err
}

// 删除顾问
// https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.delGuideAcct.html
func (comp *Client) DeleteAdviser(guideAccount string, guideOpenID string) (*response2.ResponseOfficialAccount, error) {

	result := &response2.ResponseOfficialAccount{}

	params := &object.HashMap{}
	params, err := comp.SelectAccountAndOpenID(params, guideAccount, guideOpenID)
	if err != nil {
		return nil, err
	}
	_, err = comp.HttpPostJson("cgi-bin/guide/delguideacct", params, nil, nil, result)

	return result, err
}

// 获取服务号顾问列表
// https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.delGuideAcct.html
func (comp *Client) GetAdvisers(guideAccount string, guideOpenID string) (*response.ResponseGuideGetAdvisers, error) {

	result := &response.ResponseGuideGetAdvisers{}

	params := &object.HashMap{}
	params, err := comp.SelectAccountAndOpenID(params, guideAccount, guideOpenID)
	if err != nil {
		return nil, err
	}
	_, err = comp.HttpPostJson("cgi-bin/guide/getguideacctlist", params, nil, nil, result)

	return result, err
}

// 生成顾问二维码
// https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.delGuideAcct.html
func (comp *Client) CreateQrCode(guideAccount string, guideOpenID string, qrCodeInfo string) (*response2.ResponseOfficialAccount, error) {

	result := &response2.ResponseOfficialAccount{}

	params := &object.HashMap{}
	params, err := comp.SelectAccountAndOpenID(params, guideAccount, guideOpenID)
	if err != nil {
		return nil, err
	}
	if qrCodeInfo != "" {
		(*params)["qrcode_info"] = qrCodeInfo
	}
	_, err = comp.HttpPostJson("cgi-bin/guide/guidecreateqrcode", params, nil, nil, result)

	return result, err
}

// 获取顾问聊天记录
// https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.getGuideBuyerChatRecord.html
func (comp *Client) GetBuyerChatRecords(guideAccount string, guideOpenID string, openID string, beginTime string, endTime string) (*response.ResponseGuideGetChatRecords, error) {

	result := &response.ResponseGuideGetChatRecords{}

	params := &object.HashMap{}
	params, err := comp.SelectAccountAndOpenID(params, guideAccount, guideOpenID)
	if err != nil {
		return nil, err
	}
	if openID != "" {
		(*params)["openid"] = openID
	}
	if beginTime != "" {
		(*params)["begin_time"] = beginTime
	}
	if endTime != "" {
		(*params)["end_time"] = endTime
	}
	_, err = comp.HttpPostJson("cgi-bin/guide/getguidebuyerchatrecord", params, nil, nil, result)

	return result, err
}

// 设置快捷回复与关注自动回复
// https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.setGuideConfig.html
func (comp *Client) SetConfig(guideAccount string, guideOpenID string, isDelete bool,
	fastReplyListArray *request.FastReplyList, guideAutoReply *request.AutoReply, guideAutoReplyPlus *request.AutoReply) (*response2.ResponseOfficialAccount, error) {

	result := &response2.ResponseOfficialAccount{}

	params := &object.HashMap{
		"is_delete": isDelete,
	}
	params, err := comp.SelectAccountAndOpenID(params, guideAccount, guideOpenID)
	if err != nil {
		return nil, err
	}
	if fastReplyListArray != nil {
		(*params)["guide_fast_reply_list"] = fastReplyListArray
	}
	if guideAutoReply != nil {
		(*params)["guide_auto_reply"] = guideAutoReply
	}
	if guideAutoReplyPlus != nil {
		(*params)["guide_auto_reply_plus"] = guideAutoReplyPlus
	}
	_, err = comp.HttpPostJson("cgi-bin/guide/setguideconfig", params, nil, nil, result)

	return result, err
}

// 获取快捷回复与关注自动回复
// https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.getGuideConfig.html
func (comp *Client) GetConfig(guideAccount string, guideOpenID string) (*response.ResponseGuideGetConfig, error) {

	result := &response.ResponseGuideGetConfig{}

	params := &object.HashMap{}
	params, err := comp.SelectAccountAndOpenID(params, guideAccount, guideOpenID)
	if err != nil {
		return nil, err
	}
	_, err = comp.HttpPostJson("cgi-bin/guide/getguideconfig", params, nil, nil, result)

	return result, err
}

// 设置离线自动回复与敏感词
// https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.setGuideConfig.html
func (comp *Client) SetAdviserConfig(guideAccount string, guideOpenID string, isDelete bool,
	blackKeyword *request.BlackKeyword, guideAutoReply *request.AutoReply) (*response2.ResponseOfficialAccount, error) {

	result := &response2.ResponseOfficialAccount{}

	params := &object.HashMap{
		"is_delete": isDelete,
	}
	params, err := comp.SelectAccountAndOpenID(params, guideAccount, guideOpenID)
	if err != nil {
		return nil, err
	}
	if blackKeyword != nil {
		(*params)["black_keyword"] = blackKeyword
	}
	if guideAutoReply != nil {
		(*params)["guide_auto_reply"] = guideAutoReply
	}

	_, err = comp.HttpPostJson("cgi-bin/guide/setguideacctconfig", params, nil, nil, result)

	return result, err
}

// 获取离线自动回复与敏感词
// https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.getGuideAcctConfig.html
func (comp *Client) GetAdviserConfig() (*response.ResponseGuideGetAdviserConfig, error) {

	result := &response.ResponseGuideGetAdviserConfig{}

	_, err := comp.HttpPostJson("cgi-bin/guide/getguideacctconfig", nil, nil, nil, result)

	return result, err
}

// 允许微信用户复制小程序页面路径
// https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.pushShowWxaPathMenu.html
func (comp *Client) AllowCopyMiniAppPath(wxaAppID string, wxUsername string) (*response2.ResponseOfficialAccount, error) {

	result := &response2.ResponseOfficialAccount{}

	params := &object.HashMap{
		"wxa_appid":   wxaAppID,
		"wx_username": wxUsername,
	}
	_, err := comp.HttpPostJson("cgi-bin/guide/pushshowwxapathmenu", params, nil, nil, result)

	return result, err
}

func (comp *Client) SelectAccountAndOpenID(params *object.HashMap, guideAccount string, guideOpenID string) (*object.HashMap, error) {

	if params == nil {
		params = &object.HashMap{}
	}

	if guideOpenID != "" {
		(*params)["guide_openid"] = guideOpenID

	} else if guideAccount != "" {
		(*params)["guide_account"] = guideAccount

	} else {
		return nil, errors.New("微信号和OPENID不能同时为空")
	}

	return params, nil

}

// 新建顾问分组
// https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.newGuideGroup.html
func (comp *Client) CreateGroup(groupName string) (*response.ResponseGuideCreateGroup, error) {

	result := &response.ResponseGuideCreateGroup{}

	params := &object.HashMap{
		"group_name": groupName,
	}
	_, err := comp.HttpPostJson("cgi-bin/guide/newguidegroup", params, nil, nil, result)

	return result, err
}

// 获取服务号下所有顾问分组的列表
// https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.getGuideGroupList.html
func (comp *Client) GetGuideGroups() (*response.ResponseGuideGetGroupList, error) {

	result := &response.ResponseGuideGetGroupList{}

	_, err := comp.HttpPostJson("cgi-bin/guide/getguidegrouplist", nil, nil, nil, result)

	return result, err
}

// 获取指定顾问分组信息，以及分组内顾问信息
// https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.getGroupInfo.html
func (comp *Client) GetGroups(groupName string) (*response.ResponseGuideGetGroups, error) {

	result := &response.ResponseGuideGetGroups{}

	params := &object.HashMap{
		"group_name": groupName,
	}
	_, err := comp.HttpPostJson("cgi-bin/guide/getgroupinfo", params, nil, nil, result)

	return result, err
}

// 分组内添加顾问
// https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.addGuide2GuideGroup.html
func (comp *Client) AddGroupGuide(groupID int, guideAccount string) (*response2.ResponseOfficialAccount, error) {

	result := &response2.ResponseOfficialAccount{}

	params := &object.HashMap{
		"group_id":      groupID,
		"guide_account": guideAccount,
	}
	_, err := comp.HttpPostJson("cgi-bin/guide/addguide2guidegroup", params, nil, nil, result)

	return result, err
}

// 分组内删除顾问
// https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.delGuide2GuideGroup.html
func (comp *Client) DeleteGroupGuide(groupID int, guideAccount string) (*response2.ResponseOfficialAccount, error) {

	result := &response2.ResponseOfficialAccount{}

	params := &object.HashMap{
		"group_id":      groupID,
		"guide_account": guideAccount,
	}
	_, err := comp.HttpPostJson("cgi-bin/guide/delguide2guidegroup", params, nil, nil, result)

	return result, err
}

// 获取顾问所在分组
// https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.getGroupByGuide.html
func (comp *Client) GetGuideGroup(guideAccount string) (*response.ResponseGuideGetGuideGroup, error) {

	result := &response.ResponseGuideGetGuideGroup{}

	params := &object.HashMap{
		"guide_account": guideAccount,
	}
	_, err := comp.HttpPostJson("cgi-bin/guide/getgroupbyguide", params, nil, nil, result)

	return result, err
}

// 删除指定顾问分组
// https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/guide-account/shopping-guide.delGuideGroup.html
func (comp *Client) DeleteGroup(groupID int) (*response2.ResponseOfficialAccount, error) {

	result := &response2.ResponseOfficialAccount{}

	params := &object.HashMap{
		"group_id": groupID,
	}
	_, err := comp.HttpPostJson("cgi-bin/guide/delguidegroup", params, nil, nil, result)

	return result, err
}

// 为顾问分配客户
// https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/buyer-account/shopping-guide.addGuideBuyerRelation.html
func (comp *Client) CreateBuyerRelation(guideAccount string, guideOpenID string, buyerList *request.BuyerList) (*response.ResponseGuideBuyerRelation, error) {

	result := &response.ResponseGuideBuyerRelation{}

	params := &object.HashMap{
		"buyer_list": buyerList,
	}

	params, err := comp.SelectAccountAndOpenID(params, guideAccount, guideOpenID)
	if err != nil {
		return nil, err
	}

	_, err = comp.HttpPostJson("cgi-bin/guide/addguidebuyerrelation", params, nil, nil, result)

	return result, err
}

// 为顾问移除客户
// https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/buyer-account/shopping-guide.delGuideBuyerRelation.html
func (comp *Client) DeleteBuyerRelation(guideAccount string, guideOpenID string, openIDList []string) (*response.ResponseGuideBuyerRelation, error) {

	result := &response.ResponseGuideBuyerRelation{}

	params := &object.HashMap{
		"openid_list": openIDList,
	}

	params, err := comp.SelectAccountAndOpenID(params, guideAccount, guideOpenID)
	if err != nil {
		return nil, err
	}

	_, err = comp.HttpPostJson("cgi-bin/guide/delguidebuyerrelation", params, nil, nil, result)

	return result, err
}

// 获取顾问的客户列表
// https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/buyer-account/shopping-guide.getGuideBuyerRelationList.html
func (comp *Client) GetBuyerRelations(guideAccount string, guideOpenID string, page int, num int) (*response.ResponseGuideBuyerRelationList, error) {

	result := &response.ResponseGuideBuyerRelationList{}

	params := &object.HashMap{
		"page": page,
		"num":  num,
	}

	params, err := comp.SelectAccountAndOpenID(params, guideAccount, guideOpenID)
	if err != nil {
		return nil, err
	}

	_, err = comp.HttpPostJson("cgi-bin/guide/getguidebuyerrelationlist", params, nil, nil, result)

	return result, err
}

// 为客户更换顾问
// https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/buyer-account/shopping-guide.rebindGuideAcctForBuyer.html
func (comp *Client) RebindBuyerGuide(oldGuideTarget string, newGuideTarget string, openidList []string, useTargetOpenID bool) (*response.ResponseGuideBuyerRelation, error) {

	result := &response.ResponseGuideBuyerRelation{}

	params := &object.HashMap{
		"openid_list": openidList,
	}

	if useTargetOpenID {
		(*params)["old_guide_openid"] = oldGuideTarget
		(*params)["new_guide_openid"] = newGuideTarget
	} else {
		(*params)["old_guide_account"] = oldGuideTarget
		(*params)["new_guide_account"] = newGuideTarget
	}

	_, err := comp.HttpPostJson("cgi-bin/guide/rebindguideacctforbuyer", params, nil, nil, result)

	return result, err
}

// 修改客户昵称
// https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/buyer-account/shopping-guide.updateGuideBuyerRelation.html
func (comp *Client) UpdateBuyerRelation(guideAccount string, guideOpenID string, openID string, nickName string) (*response2.ResponseOfficialAccount, error) {

	result := &response2.ResponseOfficialAccount{}

	params := &object.HashMap{
		"openid":         openID,
		"buyer_nickname": nickName,
	}

	params, err := comp.SelectAccountAndOpenID(params, guideAccount, guideOpenID)
	if err != nil {
		return nil, err
	}

	_, err = comp.HttpPostJson("cgi-bin/guide/updateguidebuyerrelation", params, nil, nil, result)

	return result, err
}

// 查询客户所属顾问
// https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/buyer-account/shopping-guide.getGuideBuyerRelationByBuyer.html
func (comp *Client) GetBuyerRelation(openID string) (*response.ResponseGuideGetBuyerRelation, error) {

	result := &response.ResponseGuideGetBuyerRelation{}

	params := &object.HashMap{
		"openid": openID,
	}

	_, err := comp.HttpPostJson("cgi-bin/guide/getguidebuyerrelationbybuyer", params, nil, nil, result)

	return result, err
}

// 查询指定顾问和客户的关系
// https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/buyer-account/shopping-guide.getGuideBuyerRelation.html
func (comp *Client) GetBuyerRelationByGuide(guideAccount string, guideOpenID string, openID string) (*response.ResponseGuideGetBuyerRelation, error) {

	result := &response.ResponseGuideGetBuyerRelation{}

	params := &object.HashMap{
		"openid": openID,
	}

	params, err := comp.SelectAccountAndOpenID(params, guideAccount, guideOpenID)
	if err != nil {
		return nil, err
	}

	_, err = comp.HttpPostJson("cgi-bin/guide/getguidebuyerrelation", params, nil, nil, result)

	return result, err
}

// 新建可查询的标签类型
// https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/tag-account/shopping-guide.newGuideTagOption.html
func (comp *Client) NewTagOption(tagName string, tagValues []string) (*response2.ResponseOfficialAccount, error) {

	result := &response2.ResponseOfficialAccount{}

	params := &object.HashMap{
		"tag_name":   tagName,
		"tag_values": tagValues,
	}

	_, err := comp.HttpPostJson("cgi-bin/guide/newguidetagoption", params, nil, nil, result)

	return result, err
}

// 删除指定标签类型
// https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/tag-account/shopping-guide.delguidetagoption.html
func (comp *Client) DeleteTagOption(tagName string) (*response2.ResponseOfficialAccount, error) {

	result := &response2.ResponseOfficialAccount{}

	params := &object.HashMap{
		"tag_name": tagName,
	}

	_, err := comp.HttpPostJson("cgi-bin/guide/delguidetagoption", params, nil, nil, result)

	return result, err
}

// 为标签添加可选值
// https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/tag-account/shopping-guide.addGuideTagOption.html
func (comp *Client) CreateTagOption(tagName string, tagValues []string) (*response2.ResponseOfficialAccount, error) {

	result := &response2.ResponseOfficialAccount{}

	params := &object.HashMap{
		"tag_name":   tagName,
		"tag_values": tagValues,
	}

	_, err := comp.HttpPostJson("cgi-bin/guide/addguidetagoption", params, nil, nil, result)

	return result, err
}

// 获取标签和可选值
// https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/tag-account/shopping-guide.getGuideTagOption.html
func (comp *Client) GetTagOption(tagName string, tagValues []string) (*response.ResponseGuideTagOption, error) {

	result := &response.ResponseGuideTagOption{}

	_, err := comp.HttpPostJson("cgi-bin/guide/getguidetagoption", nil, nil, nil, result)

	return result, err
}

// 为客户设置标签
// https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/tag-account/shopping-guide.addGuideBuyerTag.html
func (comp *Client) SetBuyersTag(guideAccount string, guideOpenID string, openIDList []string, tagValue string) (*response.ResponseGuideBuyerRelation, error) {

	result := &response.ResponseGuideBuyerRelation{}

	params := &object.HashMap{
		"tag_value":   tagValue,
		"openid_list": openIDList,
	}

	params, err := comp.SelectAccountAndOpenID(params, guideAccount, guideOpenID)
	if err != nil {
		return nil, err
	}

	_, err = comp.HttpPostJson("cgi-bin/guide/addguidebuyertag", params, nil, nil, result)

	return result, err
}

// 查询客户标签
// https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/tag-account/shopping-guide.getGuideBuyerTag.html
func (comp *Client) GetBuyerTags(guideAccount string, guideOpenID string, openIDList []string, tagValue string) (*response.ResponseGuideGetBuyerTags, error) {

	result := &response.ResponseGuideGetBuyerTags{}

	params := &object.HashMap{
		"tag_value":   tagValue,
		"openid_list": openIDList,
	}

	params, err := comp.SelectAccountAndOpenID(params, guideAccount, guideOpenID)
	if err != nil {
		return nil, err
	}

	_, err = comp.HttpPostJson("cgi-bin/guide/addguidebuyertag", params, nil, nil, result)

	return result, err
}

// 根据标签值筛选客户
// https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/tag-account/shopping-guide.queryGuideBuyerByTag.html
func (comp *Client) GetBuyerByTag(guideAccount string, guideOpenID string, pushCount int, tagValues []string) (*response.ResponseGuideGetBuyerByTag, error) {

	result := &response.ResponseGuideGetBuyerByTag{}

	params := &object.HashMap{}

	params, err := comp.SelectAccountAndOpenID(params, guideAccount, guideOpenID)
	if err != nil {
		return nil, err
	}

	if pushCount > 0 {
		(*params)["push_count"] = pushCount
	}
	if len(tagValues) > 0 {
		(*params)["tag_values"] = tagValues
	}

	_, err = comp.HttpPostJson("cgi-bin/guide/queryguidebuyerbytag", params, nil, nil, result)

	return result, err
}

// 删除客户标签
// https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/tag-account/shopping-guide.delGuideBuyerTag.html
func (comp *Client) DeleteBuyerTag(guideAccount string, guideOpenID string, tagValue string, openIDList []string) (*response.ResponseGuideBuyerRelation, error) {

	result := &response.ResponseGuideBuyerRelation{}

	params := &object.HashMap{
		"tag_value": tagValue,
	}

	params, err := comp.SelectAccountAndOpenID(params, guideAccount, guideOpenID)
	if err != nil {
		return nil, err
	}

	if len(openIDList) > 0 {
		(*params)["openid_list"] = openIDList
	}

	_, err = comp.HttpPostJson("cgi-bin/guide/delguidebuyertag", params, nil, nil, result)

	return result, err
}

// 设置自定义客户信息
// https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/tag-account/shopping-guide.addGuideBuyerDisplayTag.html
func (comp *Client) SetBuyerDisplayTags(guideAccount string, guideOpenID string, openID string, displayTagList []string) (*response2.ResponseOfficialAccount, error) {

	result := &response2.ResponseOfficialAccount{}

	params := &object.HashMap{
		"openid":           openID,
		"display_tag_list": displayTagList,
	}

	params, err := comp.SelectAccountAndOpenID(params, guideAccount, guideOpenID)
	if err != nil {
		return nil, err
	}

	_, err = comp.HttpPostJson("cgi-bin/guide/addguidebuyerdisplaytag", params, nil, nil, result)

	return result, err
}

// 获取自定义客户信息
// https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/tag-account/shopping-guide.getGuideBuyerDisplayTag.html
func (comp *Client) GetBuyerDisplayTags(guideAccount string, guideOpenID string, openID string) (*response.ResponseGuideGetBuyerDisplayTags, error) {

	result := &response.ResponseGuideGetBuyerDisplayTags{}

	params := &object.HashMap{
		"openid": openID,
	}

	params, err := comp.SelectAccountAndOpenID(params, guideAccount, guideOpenID)
	if err != nil {
		return nil, err
	}

	_, err = comp.HttpPostJson("cgi-bin/guide/getguidebuyerdisplaytag", params, nil, nil, result)

	return result, err
}

// 添加小程序卡片素材
// https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/model-account/shopping-guide.setGuideCardMaterial.html
func (comp *Client) CreateCardMaterial(mediaID string, title string, path string, appID string, mType int) (*response2.ResponseOfficialAccount, error) {

	result := &response2.ResponseOfficialAccount{}

	params := &object.HashMap{
		"media_id": mediaID,
		"type":     mType,
		"title":    title,
		"path":     path,
		"appid":    appID,
	}

	_, err := comp.HttpPostJson("cgi-bin/guide/setguidecardmaterial", params, nil, nil, result)

	return result, err
}

// 查询小程序卡片素材
// https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/model-account/shopping-guide.getGuideCardMaterial.html
func (comp *Client) GetCardMaterial(mType int) (*response.ResponseGuideGetCardMaterial, error) {

	result := &response.ResponseGuideGetCardMaterial{}

	params := &object.HashMap{
		"type": mType,
	}

	_, err := comp.HttpPostJson("cgi-bin/guide/getguidecardmaterial", params, nil, nil, result)

	return result, err
}

// 删除小程序卡片素材
// https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/model-account/shopping-guide.delGuideCardMaterial.html
func (comp *Client) DeleteCardMaterial(title string, path string, appID string, mType int) (*response2.ResponseOfficialAccount, error) {

	result := &response2.ResponseOfficialAccount{}

	params := &object.HashMap{
		"type":  mType,
		"title": title,
		"path":  path,
		"appid": appID,
	}

	_, err := comp.HttpPostJson("cgi-bin/guide/delguidecardmaterial", params, nil, nil, result)

	return result, err
}

// 添加图片素材
// https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/model-account/shopping-guide.setGuideImageMaterial.html
func (comp *Client) CreateImageMaterial(mediaID string, mType int) (*response2.ResponseOfficialAccount, error) {

	result := &response2.ResponseOfficialAccount{}

	params := &object.HashMap{
		"media_id": mediaID,
		"type":     mType,
	}

	_, err := comp.HttpPostJson("cgi-bin/guide/setguideimagematerial", params, nil, nil, result)

	return result, err
}

// 查询图片素材
// https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/model-account/shopping-guide.getGuideImageMaterial.html
func (comp *Client) GetImageMaterial(mType int) (*response.ResponseGuideGetImageMaterial, error) {

	result := &response.ResponseGuideGetImageMaterial{}

	params := &object.HashMap{
		"type": mType,
	}

	_, err := comp.HttpPostJson("cgi-bin/guide/getguideimagematerial", params, nil, nil, result)

	return result, err
}

// 删除图片素材
// https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/model-account/shopping-guide.delGuideImageMaterial.html
func (comp *Client) DeleteImageMaterial(mType int, picURL string) (*response2.ResponseOfficialAccount, error) {

	result := &response2.ResponseOfficialAccount{}

	params := &object.HashMap{
		"type":   mType,
		"picurl": picURL,
	}

	_, err := comp.HttpPostJson("cgi-bin/guide/delguideimagematerial", params, nil, nil, result)

	return result, err
}

// 添加文字素材
// https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/model-account/shopping-guide.setGuideWordMaterial.html
func (comp *Client) CreateWordMaterial(mType int, word string) (*response2.ResponseOfficialAccount, error) {

	result := &response2.ResponseOfficialAccount{}

	params := &object.HashMap{
		"type": mType,
		"word": word,
	}

	_, err := comp.HttpPostJson("cgi-bin/guide/setguidewordmaterial", params, nil, nil, result)

	return result, err
}

// 查询文字素材
// https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/model-account/shopping-guide.getGuideWordMaterial.html
func (comp *Client) GetWordMaterial(mType int, start int, num int) (*response.ResponseGuideGetWordMaterial, error) {

	result := &response.ResponseGuideGetWordMaterial{}

	params := &object.HashMap{
		"type":  mType,
		"start": start,
		"num":   num,
	}

	_, err := comp.HttpPostJson("cgi-bin/guide/getguideimagematerial", params, nil, nil, result)

	return result, err
}

// 删除文字素材
// https://developers.weixin.qq.com/doc/offiaccount/Shopping_Guide/model-account/shopping-guide.delGuideWordMaterial.html
func (comp *Client) DeleteWordMaterial(mType int, word string) (*response2.ResponseOfficialAccount, error) {

	result := &response2.ResponseOfficialAccount{}

	params := &object.HashMap{
		"type": mType,
		"word": word,
	}

	_, err := comp.HttpPostJson("cgi-bin/guide/delguidewordmaterial", params, nil, nil, result)

	return result, err
}
