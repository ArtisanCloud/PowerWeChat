package liveBroadcast

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	response2 "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/liveBroadcast/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/liveBroadcast/response"
	"strconv"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

// 创建直播间
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.createRoom.html
func (comp *Client) CreateRoom(ctx context.Context, options *request.RequestBroadcastCreateRoom) (*response.ResponseBroadcastCreateRoom, error) {
	result := &response.ResponseBroadcastCreateRoom{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "wxaapi/broadcast/room/create", options, nil, nil, result)

	return result, err
}

// 获取直播间列表及直播间信息
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.getLiveInfo.html
func (comp *Client) GetLiveInfo(ctx context.Context, options *request.RequestBroadcastGetLiveInfo) (*response.ResponseBroadcastGetLiveInfo, error) {
	result := &response.ResponseBroadcastGetLiveInfo{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "wxa/business/getliveinfo", options, nil, nil, result)

	return result, err
}

// 获取直播间回放
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.getLiveInfo.html
func (comp *Client) GetLiveReplay(ctx context.Context, options *request.RequestBroadcastGetLiveReplay) (*response.ResponseBroadcastGetLiveReplay, error) {
	result := &response.ResponseBroadcastGetLiveReplay{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "wxa/business/getliveinfo", options, nil, nil, result)

	return result, err
}

// 直播间导入商品
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.addGoods.html
func (comp *Client) AddGoods(ctx context.Context, options *request.RequestBroadcastAddGoods) (*response2.ResponseMiniProgram, error) {
	result := &response2.ResponseMiniProgram{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "wxaapi/broadcast/room/addgoods", options, nil, nil, result)

	return result, err
}

// 删除直播间
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.deleteRoom.html
func (comp *Client) DeleteRoom(ctx context.Context, id int) (*response2.ResponseMiniProgram, error) {
	result := &response2.ResponseMiniProgram{}

	options := &object.HashMap{
		"id": id,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "wxaapi/broadcast/room/deleteroom", options, nil, nil, result)

	return result, err
}

// 编辑直播间
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.editRoom.html
func (comp *Client) EditRoom(ctx context.Context, options *request.RequestBroadcastEditRoom) (*response2.ResponseMiniProgram, error) {
	result := &response2.ResponseMiniProgram{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "wxaapi/broadcast/room/editroom", options, nil, nil, result)

	return result, err
}

// 获取直播间推流地址
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.getPushUrl.html
func (comp *Client) GetPushUrl(ctx context.Context, roomID int) (*response.ResponseBroadcastGetPushUrl, error) {
	result := &response.ResponseBroadcastGetPushUrl{}

	options := &object.StringMap{
		"roomId": strconv.Itoa(roomID),
	}

	_, err := comp.BaseClient.HttpGet(ctx, "wxaapi/broadcast/room/getpushurl", options, nil, result)

	return result, err
}

// 获取直播间分享二维码
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.getSharedCode.html
func (comp *Client) GetSharedCode(ctx context.Context, roomID int, params string) (*response.ResponseBroadcastGetSharedCode, error) {
	result := &response.ResponseBroadcastGetSharedCode{}

	options := &object.StringMap{
		"roomId": strconv.Itoa(roomID),
		"params": params,
	}

	_, err := comp.BaseClient.HttpGet(ctx, "wxaapi/broadcast/room/getsharedcode", options, nil, result)

	return result, err
}

// 添加管理直播间小助手
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.addAssistant.html
func (comp *Client) AddAssistant(ctx context.Context, options *request.RequestBroadcastAddAssistant) (*response2.ResponseMiniProgram, error) {

	result := &response2.ResponseMiniProgram{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "wxaapi/broadcast/room/addassistant", options, nil, nil, result)

	return result, err
}

// 修改管理直播间小助手
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.modifyAssistant.html
func (comp *Client) ModifyAssistant(ctx context.Context, options *request.RequestBroadcastModifyAssistant) (*response2.ResponseMiniProgram, error) {
	result := &response2.ResponseMiniProgram{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "wxaapi/broadcast/room/modifyassistant", options, nil, nil, result)

	return result, err
}

// 删除管理直播间小助手
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.removeAssistant.html
func (comp *Client) RemoveAssistant(ctx context.Context, options *request.RequestBroadcastRemoveAssistant) (*response2.ResponseMiniProgram, error) {
	result := &response2.ResponseMiniProgram{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "wxaapi/broadcast/room/removeassistant", options, nil, nil, result)

	return result, err
}

// 查询管理直播间小助手
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.getAssistantList.html
func (comp *Client) GetAssistantList(ctx context.Context, roomID int) (*response.ResponseBroadcastGetAssistantList, error) {
	result := &response.ResponseBroadcastGetAssistantList{}

	options := &object.StringMap{
		"roomId": strconv.Itoa(roomID),
	}

	_, err := comp.BaseClient.HttpGet(ctx, "wxaapi/broadcast/room/getassistantlist", options, nil, result)

	return result, err
}

// 添加主播副号
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.addSubAnchor.html
func (comp *Client) AddSubAnchor(ctx context.Context, roomID int, userName string) (*response2.ResponseMiniProgram, error) {
	result := &response2.ResponseMiniProgram{}

	options := &object.HashMap{
		"roomId":   roomID,
		"username": userName,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "wxaapi/broadcast/room/addsubanchor", options, nil, nil, result)

	return result, err
}

// 修改主播副号
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.modifySubAnchor.html
func (comp *Client) ModifySubAnchor(ctx context.Context, roomID int, username string) (*response2.ResponseMiniProgram, error) {
	result := &response2.ResponseMiniProgram{}

	options := &object.HashMap{
		"roomId":   roomID,
		"username": username,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "wxaapi/broadcast/room/modifysubanchor", options, nil, nil, result)

	return result, err
}

// 删除主播副号
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.deleteSubAnchor.html
func (comp *Client) DeleteSubAnchor(ctx context.Context, roomID int) (*response2.ResponseMiniProgram, error) {
	result := &response2.ResponseMiniProgram{}

	options := &object.HashMap{
		"roomId": roomID,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "wxaapi/broadcast/room/deletesubanchor", options, nil, nil, result)

	return result, err
}

// 获取主播副号
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.getSubAnchor.html
func (comp *Client) GetSubAnchor(ctx context.Context, roomID int) (*response.ResponseBroadcastGetSubAnchor, error) {
	result := &response.ResponseBroadcastGetSubAnchor{}

	options := &object.StringMap{
		"roomId": strconv.Itoa(roomID),
	}

	_, err := comp.BaseClient.HttpGet(ctx, "wxaapi/broadcast/room/getsubanchor", options, nil, result)

	return result, err
}

// 开启/关闭直播间官方收录
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.updateFeedPublic.html
func (comp *Client) UpdateFeedPublic(ctx context.Context, roomID int, isFeedsPublic int) (*response2.ResponseMiniProgram, error) {
	result := &response2.ResponseMiniProgram{}

	options := &object.HashMap{
		"roomId":        roomID,
		"isFeedsPublic": isFeedsPublic,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "wxaapi/broadcast/room/updatefeedpublic", options, nil, nil, result)

	return result, err
}

// 开启/关闭回放功能
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.updateReplay.html
func (comp *Client) UpdateReplay(ctx context.Context, roomID int, closeReplay int) (*response2.ResponseMiniProgram, error) {
	result := &response2.ResponseMiniProgram{}
	options := &object.HashMap{
		"roomId":      roomID,
		"closeReplay": closeReplay,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "wxaapi/broadcast/room/updatereplay", options, nil, nil, result)

	return result, err
}

// 开启/关闭客服功能
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.updateKF.html
func (comp *Client) UpdateKF(ctx context.Context, roomID int, closeKF int) (*response2.ResponseMiniProgram, error) {
	result := &response2.ResponseMiniProgram{}

	options := &object.HashMap{
		"roomId":  roomID,
		"closeKf": closeKF,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "wxaapi/broadcast/room/updatekf", options, nil, nil, result)

	return result, err
}

// 开启/关闭直播间全局禁言
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.updateComment.html
func (comp *Client) UpdateComment(ctx context.Context, roomID int, banComment int) (*response2.ResponseMiniProgram, error) {
	result := &response2.ResponseMiniProgram{}

	options := &object.HashMap{
		"roomId":     roomID,
		"banComment": banComment,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "wxaapi/broadcast/room/updatecomment", options, nil, nil, result)

	return result, err
}

// 上下架商品
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.goodsSale.html
func (comp *Client) GoodsSale(ctx context.Context, roomID int, goodsID int, onSale int) (*response2.ResponseMiniProgram, error) {
	result := &response2.ResponseMiniProgram{}

	options := &object.HashMap{
		"roomId":  roomID,
		"goodsId": goodsID,
		"onSale":  onSale,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "wxaapi/broadcast/goods/onsale", options, nil, nil, result)

	return result, err
}

// 删除直播间商品
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.goodsAudit.html
func (comp *Client) GoodsDeleteInRoom(ctx context.Context, roomID int, goodsID int) (*response2.ResponseMiniProgram, error) {
	result := &response2.ResponseMiniProgram{}

	options := &object.HashMap{
		"roomId":  roomID,
		"goodsId": goodsID,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "wxaapi/broadcast/goods/deleteInRoom", options, nil, nil, result)

	return result, err
}

// 推送商品
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.goodsPush.html
func (comp *Client) GoodsPush(ctx context.Context, roomID int, goodsID int) (*response2.ResponseMiniProgram, error) {
	result := &response2.ResponseMiniProgram{}

	options := &object.HashMap{
		"roomId":  roomID,
		"goodsId": goodsID,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "wxaapi/broadcast/goods/push", options, nil, nil, result)

	return result, err
}

// 直播间商品排序
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.goodsSort.html
func (comp *Client) GoodsSort(ctx context.Context, roomID int, goods []request.RequestBroadcastGoodsSort) (*response2.ResponseMiniProgram, error) {
	result := &response2.ResponseMiniProgram{}

	options := &object.HashMap{
		"roomId": roomID,
		"goods":  goods,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "wxaapi/broadcast/goods/sort", options, nil, nil, result)

	return result, err
}

// 下载商品讲解视频
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.goodsVideo.html
func (comp *Client) GoodsVideo(ctx context.Context, roomID int, goodsID int) (*response.ResponseBroadcastGoodsVideo, error) {
	result := &response.ResponseBroadcastGoodsVideo{}

	options := &object.HashMap{
		"roomId":  roomID,
		"goodsId": goodsID,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "wxaapi/broadcast/goods/getVideo", options, nil, nil, result)

	return result, err
}

// 商品添加并提审
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.goodsAdd.html
func (comp *Client) GoodsAdd(ctx context.Context, options *request.RequestBroadcastGoodsAdd) (*response.ResponseBroadcastGoodsAdd, error) {
	result := &response.ResponseBroadcastGoodsAdd{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "wxaapi/broadcast/goods/add", options, nil, nil, result)

	return result, err
}

// 撤回商品审核
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.goodsResetAudit.html
func (comp *Client) GoodsResetAudit(ctx context.Context, options *request.RequestBroadcastGoodsResetAudit) (*response2.ResponseMiniProgram, error) {
	result := &response2.ResponseMiniProgram{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "wxaapi/broadcast/goods/resetaudit", options, nil, nil, result)

	return result, err
}

// 重新提交审核
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.goodsAudit.html
func (comp *Client) GoodsAudit(ctx context.Context, goodsID int) (*response.ResponseBroadcastGoodsAudit, error) {
	result := &response.ResponseBroadcastGoodsAudit{}

	options := &power.HashMap{
		"goodsId": goodsID,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "wxaapi/broadcast/goods/audit", options, nil, nil, result)

	return result, err
}

// 删除商品
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.goodsAudit.html
func (comp *Client) GoodsDelete(ctx context.Context, goodsID int) (*response2.ResponseMiniProgram, error) {
	result := &response2.ResponseMiniProgram{}

	options := &power.HashMap{
		"goodsId": goodsID,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "wxaapi/broadcast/goods/delete", options, nil, nil, result)

	return result, err
}

// 更新商品
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.goodsUpdate.html
func (comp *Client) GoodsUpdate(ctx context.Context, options *request.RequestBroadcastGoodsUpdate) (*response2.ResponseMiniProgram, error) {
	result := &response2.ResponseMiniProgram{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "wxaapi/broadcast/goods/update", options, nil, nil, result)

	return result, err
}

// 获取商品状态
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.goodsInfo.html
func (comp *Client) GoodsInfo(ctx context.Context, goodsIDs []int) (*response.ResponseBroadcastGoodsInfo, error) {
	result := &response.ResponseBroadcastGoodsInfo{}

	options := &power.HashMap{
		"goods_ids": goodsIDs,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "wxa/business/getgoodswarehouse", options, nil, nil, result)

	return result, err
}

// 获取商品列表
// https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/liveplayer/commodity-api.html#7
func (comp *Client) GoodsList(ctx context.Context, offset, count, status string) (*response.ResponseBroadcastGoodsList, error) {
	result := &response.ResponseBroadcastGoodsList{}

	options := &object.StringMap{
		"offset": offset,
		"count":  count,
		"status": status,
	}

	_, err := comp.BaseClient.HttpGet(ctx, "wxaapi/broadcast/goods/getapproved", options, nil, result)

	return result, err
}

// 设置成员角色
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.addRole.html
func (comp *Client) AddRole(ctx context.Context, userName string, role int) (*response2.ResponseMiniProgram, error) {
	result := &response2.ResponseMiniProgram{}

	options := &object.HashMap{
		"username": userName,
		"role":     role,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "wxaapi/broadcast/role/addrole", options, nil, nil, result)

	return result, err
}

// 解除成员角色
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.deleteRole.html
func (comp *Client) DeleteRole(ctx context.Context, userName string, role int) (*response2.ResponseMiniProgram, error) {
	result := &response2.ResponseMiniProgram{}

	options := &object.HashMap{
		"username": userName,
		"role":     role,
	}
	_, err := comp.BaseClient.HttpPostJson(ctx, "wxaapi/broadcast/role/deleterole", options, nil, nil, result)

	return result, err
}

// 获取直播间推流地址
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.getPushUrl.html
// 获取角色列表
func (comp *Client) GetRoleList(ctx context.Context, options *request.RequestBroadcastGetRoleList) (*response.ResponseBroadcastGetRoleList, error) {
	result := &response.ResponseBroadcastGetRoleList{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "wxaapi/broadcast/role/getrolelist", options, nil, nil, result)

	return result, err
}

// 获取长期订阅用户
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.getFollowers.html
func (comp *Client) GetFollowers(ctx context.Context, limit int, pageBreak int) (*response.ResponseBroadcastGetFollowers, error) {
	result := &response.ResponseBroadcastGetFollowers{}

	options := &power.HashMap{
		"limit":      limit,
		"page_break": pageBreak,
	}
	_, err := comp.BaseClient.HttpPostJson(ctx, "wxa/business/get_wxa_followers", options, nil, nil, result)

	return result, err
}

// 向长期订阅用户群发直播间开始事件
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/livebroadcast/liveBroadcast.pushMessage.html
func (comp *Client) PushMessage(ctx context.Context, roomID int, userOpenID []string) (*response2.ResponseMiniProgram, error) {
	result := &response2.ResponseMiniProgram{}

	options := &power.HashMap{
		"room_id":     roomID,
		"user_openid": userOpenID,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "wxa/business/push_message", options, nil, nil, result)

	return result, err
}
