package broadcast

import (
	"github.com/ArtisanCloud/go-libs/http/response"
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
)

type Client struct {
	*kernel.BaseClient
}

// 添加管理直播间小助手
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.addAssistant.html
func (comp *Client) AddAssistant(data *power.HashMap) (*response.HttpResponse, error) {

	rs, err := comp.HttpPostJson("wxaapi/broadcast/room/addassistant", data, nil, nil, nil)

	return rs.(*response.HttpResponse), err
}

// 直播间导入商品
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.addGoods.html
func (comp *Client) AddGoods(data *power.HashMap) (*response.HttpResponse, error) {

	rs, err := comp.HttpPostJson("wxaapi/broadcast/room/addgoods", data, nil, nil, nil)

	return rs.(*response.HttpResponse), err
}

// 设置成员角色
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.addRole.html
func (comp *Client) AddRole(data *power.HashMap) (*response.HttpResponse, error) {

	rs, err := comp.HttpPostJson("wxaapi/broadcast/role/addrole", data, nil, nil, nil)

	return rs.(*response.HttpResponse), err
}

// 添加主播副号
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.addSubAnchor.html
func (comp *Client) AddSubAnchor(data *power.HashMap) (*response.HttpResponse, error) {

	rs, err := comp.HttpPostJson("wxaapi/broadcast/room/addsubanchor", data, nil, nil, nil)

	return rs.(*response.HttpResponse), err
}

// 创建直播间
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.createRoom.html
func (comp *Client) CreateRoom(data *power.HashMap) (*response.HttpResponse, error) {

	rs, err := comp.HttpPostJson("wxaapi/broadcast/room/create", data, nil, nil, nil)

	return rs.(*response.HttpResponse), err
}

// 解除成员角色
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.deleteRole.html
func (comp *Client) DeleteRole(data *power.HashMap) (*response.HttpResponse, error) {

	rs, err := comp.HttpPostJson("wxaapi/broadcast/role/deleterole", data, nil, nil, nil)

	return rs.(*response.HttpResponse), err
}

// 删除直播间
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.deleteRoom.html
func (comp *Client) DeleteRoom(data *power.HashMap) (*response.HttpResponse, error) {

	rs, err := comp.HttpPostJson("wxaapi/broadcast/room/deleteroom", data, nil, nil, nil)

	return rs.(*response.HttpResponse), err
}

// 删除主播副号
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.deleteSubAnchor.html
func (comp *Client) DeleteSubAnchor(data *power.HashMap) (*response.HttpResponse, error) {

	rs, err := comp.HttpPostJson("wxaapi/broadcast/room/deletesubanchor", data, nil, nil, nil)

	return rs.(*response.HttpResponse), err
}

// 编辑直播间
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.editRoom.html
func (comp *Client) EditRoom(data *power.HashMap) (*response.HttpResponse, error) {

	rs, err := comp.HttpPostJson("wxaapi/broadcast/room/editroom", data, nil, nil, nil)

	return rs.(*response.HttpResponse), err
}

// 查询管理直播间小助手
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.getAssistantList.html
func (comp *Client) GetAssistantList(data *power.HashMap) (*response.HttpResponse, error) {

	rs, err := comp.HttpPostJson("wxaapi/broadcast/room/getassistantlist", data, nil, nil, nil)

	return rs.(*response.HttpResponse), err
}

// 获取长期订阅用户
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.getFollowers.html
func (comp *Client) GetFollowers(data *power.HashMap) (*response.HttpResponse, error) {

	rs, err := comp.HttpPostJson("wxa/business/get_wxa_followers", data, nil, nil, nil)

	return rs.(*response.HttpResponse), err
}

// 获取直播间列表及直播间信息
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.getLiveInfo.html
func (comp *Client) GetLiveInfo(data *power.HashMap) (*response.HttpResponse, error) {

	rs, err := comp.HttpPostJson("wxa/business/getliveinfo", data, nil, nil, nil)

	return rs.(*response.HttpResponse), err
}

// 获取直播间推流地址
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.getPushUrl.html
func (comp *Client) GetPushUrl(data *power.HashMap) (*response.HttpResponse, error) {

	rs, err := comp.HttpPostJson("wxaapi/broadcast/room/getpushurl", data, nil, nil, nil)

	return rs.(*response.HttpResponse), err
}

// 获取直播间推流地址
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.getPushUrl.html
func (comp *Client) GetRoleList(data *power.HashMap) (*response.HttpResponse, error) {

	rs, err := comp.HttpPostJson("wxaapi/broadcast/role/getrolelist", data, nil, nil, nil)

	return rs.(*response.HttpResponse), err
}

// 获取直播间分享二维码
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.getSharedCode.html
func (comp *Client) GetSharedCode(data *power.HashMap) (*response.HttpResponse, error) {

	rs, err := comp.HttpPostJson("wxaapi/broadcast/room/getsharedcode", data, nil, nil, nil)

	return rs.(*response.HttpResponse), err
}

// 获取主播副号
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.getSubAnchor.html
func (comp *Client) GetSubAnchor(data *power.HashMap) (*response.HttpResponse, error) {

	rs, err := comp.HttpPostJson("wxaapi/broadcast/room/getsubanchor", data, nil, nil, nil)

	return rs.(*response.HttpResponse), err
}

// 商品添加并提审
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.goodsAdd.html
func (comp *Client) GoodsAdd(data *power.HashMap) (*response.HttpResponse, error) {

	rs, err := comp.HttpPostJson("wxaapi/broadcast/goods/add", data, nil, nil, nil)

	return rs.(*response.HttpResponse), err
}

// 重新提交审核
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.goodsAudit.html
func (comp *Client) GoodsAudit(data *power.HashMap) (*response.HttpResponse, error) {

	rs, err := comp.HttpPostJson("wxaapi/broadcast/goods/audit", data, nil, nil, nil)

	return rs.(*response.HttpResponse), err
}

// 重新提交审核
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.goodsAudit.html
func (comp *Client) GoodsDelete(data *power.HashMap) (*response.HttpResponse, error) {

	rs, err := comp.HttpPostJson("wxaapi/broadcast/goods/deleteInRoom", data, nil, nil, nil)

	return rs.(*response.HttpResponse), err
}

// 获取商品状态
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.goodsInfo.html
func (comp *Client) GoodsInfo(data *power.HashMap) (*response.HttpResponse, error) {

	rs, err := comp.HttpPostJson("wxa/business/getgoodswarehouse", data, nil, nil, nil)

	return rs.(*response.HttpResponse), err
}

// 获取商品列表
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.goodsList.html
func (comp *Client) GoodsList(data *power.HashMap) (*response.HttpResponse, error) {

	rs, err := comp.HttpPostJson("wxaapi/broadcast/goods/getapproved", data, nil, nil, nil)

	return rs.(*response.HttpResponse), err
}


// 推送商品
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.goodsPush.html
func (comp *Client) GoodsPush(data *power.HashMap) (*response.HttpResponse, error) {

	rs, err := comp.HttpPostJson("wxaapi/broadcast/goods/push", data, nil, nil, nil)

	return rs.(*response.HttpResponse), err
}


// 撤回商品审核
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.goodsResetAudit.html
func (comp *Client) GoodsResetAudit(data *power.HashMap) (*response.HttpResponse, error) {

	rs, err := comp.HttpPostJson("wxaapi/broadcast/goods/resetaudit", data, nil, nil, nil)

	return rs.(*response.HttpResponse), err
}


// 上下架商品
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.goodsSale.html
func (comp *Client) GoodsSale(data *power.HashMap) (*response.HttpResponse, error) {

	rs, err := comp.HttpPostJson("wxaapi/broadcast/goods/onsale", data, nil, nil, nil)

	return rs.(*response.HttpResponse), err
}


// 直播间商品排序
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.goodsSort.html
func (comp *Client) GoodsSort(data *power.HashMap) (*response.HttpResponse, error) {

	rs, err := comp.HttpPostJson("wxaapi/broadcast/goods/sort", data, nil, nil, nil)

	return rs.(*response.HttpResponse), err
}


// 更新商品
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.goodsUpdate.html
func (comp *Client) GoodsUpdate(data *power.HashMap) (*response.HttpResponse, error) {

	rs, err := comp.HttpPostJson("wxaapi/broadcast/goods/update", data, nil, nil, nil)

	return rs.(*response.HttpResponse), err
}



// 下载商品讲解视频
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.goodsVideo.html
func (comp *Client) GoodsVideo(data *power.HashMap) (*response.HttpResponse, error) {

	rs, err := comp.HttpPostJson("wxaapi/broadcast/goods/getVideo", data, nil, nil, nil)

	return rs.(*response.HttpResponse), err
}




// 修改管理直播间小助手
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.modifyAssistant.html
func (comp *Client) ModifyAssistant(data *power.HashMap) (*response.HttpResponse, error) {

	rs, err := comp.HttpPostJson("wxaapi/broadcast/room/modifyassistant", data, nil, nil, nil)

	return rs.(*response.HttpResponse), err
}



// 修改主播副号
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.modifySubAnchor.html
func (comp *Client) ModifySubAnchor(data *power.HashMap) (*response.HttpResponse, error) {

	rs, err := comp.HttpPostJson("wxaapi/broadcast/room/modifysubanchor", data, nil, nil, nil)

	return rs.(*response.HttpResponse), err
}


// 向长期订阅用户群发直播间开始事件
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.pushMessage.html
func (comp *Client) PushMessage(data *power.HashMap) (*response.HttpResponse, error) {

	rs, err := comp.HttpPostJson("wxa/business/push_message", data, nil, nil, nil)

	return rs.(*response.HttpResponse), err
}


// 删除管理直播间小助手
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.removeAssistant.html
func (comp *Client) RemoveAssistant(data *power.HashMap) (*response.HttpResponse, error) {

	rs, err := comp.HttpPostJson("wxaapi/broadcast/room/removeassistant", data, nil, nil, nil)

	return rs.(*response.HttpResponse), err
}


// 开启/关闭直播间全局禁言
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.updateComment.html
func (comp *Client) UpdateComment(data *power.HashMap) (*response.HttpResponse, error) {

	rs, err := comp.HttpPostJson("wxaapi/broadcast/room/updatecomment", data, nil, nil, nil)

	return rs.(*response.HttpResponse), err
}

// 开启/关闭直播间官方收录
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.updateFeedPublic.html
func (comp *Client) UpdateFeedPublic(data *power.HashMap) (*response.HttpResponse, error) {

	rs, err := comp.HttpPostJson("wxaapi/broadcast/room/updatefeedpublic", data, nil, nil, nil)

	return rs.(*response.HttpResponse), err
}


// 开启/关闭客服功能
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.updateKF.html
func (comp *Client) UpdateKF(data *power.HashMap) (*response.HttpResponse, error) {

	rs, err := comp.HttpPostJson("wxaapi/broadcast/room/updatekf", data, nil, nil, nil)

	return rs.(*response.HttpResponse), err
}



// 开启/关闭回放功能
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.updateReplay.html
func (comp *Client) UpdateReplay(data *power.HashMap) (*response.HttpResponse, error) {

	rs, err := comp.HttpPostJson("wxaapi/broadcast/room/updatereplay", data, nil, nil, nil)

	return rs.(*response.HttpResponse), err
}



