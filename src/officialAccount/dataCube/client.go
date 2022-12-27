package dataCube

import (
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/dataCube/response"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

// 获取用户增减数据
// https://developers.weixin.qq.com/doc/offiaccount/Analytics/User_Analysis_Data_Interface.html
func (comp *Client) GetUserSummary(from string, to string) (*response.ResponseDataCubeUserSummary, error) {

	result := &response.ResponseDataCubeUserSummary{}

	err := comp.Query("datacube/getusersummary", from, to, nil, result)

	return result, err
}

// 获取累计用户数据
// https://developers.weixin.qq.com/doc/offiaccount/Analytics/User_Analysis_Data_Interface.html
func (comp *Client) GetUserCumulate(from string, to string) (*response.ResponseDataCubeUserCumulate, error) {

	result := &response.ResponseDataCubeUserCumulate{}

	err := comp.Query("datacube/getusercumulate", from, to, nil, result)

	return result, err
}

// 获取图文群发每日数据
// https://developers.weixin.qq.com/doc/offiaccount/Analytics/Graphic_Analysis_Data_Interface.html
func (comp *Client) ArticleSummary(from string, to string) (*response.ResponseDataCubeArticleSummary, error) {

	result := &response.ResponseDataCubeArticleSummary{}

	err := comp.Query("datacube/getarticlesummary", from, to, nil, result)

	return result, err
}

// 获取图文群发总数据
// https://developers.weixin.qq.com/doc/offiaccount/Analytics/Graphic_Analysis_Data_Interface.html
func (comp *Client) ArticleTotal(from string, to string) (*response.ResponseDataCubeArticleTotal, error) {

	result := &response.ResponseDataCubeArticleTotal{}

	err := comp.Query("datacube/getarticletotal", from, to, nil, result)

	return result, err
}

// 获取图文统计数据
// https://developers.weixin.qq.com/doc/offiaccount/Analytics/Graphic_Analysis_Data_Interface.html
func (comp *Client) UserReadSummary(from string, to string) (*response.ResponseDataCubeUserReadSummary, error) {

	result := &response.ResponseDataCubeUserReadSummary{}

	err := comp.Query("datacube/getuserread", from, to, nil, result)

	return result, err
}

// 获取图文分享转发数据.
// https://developers.weixin.qq.com/doc/offiaccount/Analytics/Graphic_Analysis_Data_Interface.html
func (comp *Client) UserShareSummary(from string, to string) (*response.ResponseDataCubeUserShareSummary, error) {

	result := &response.ResponseDataCubeUserShareSummary{}

	err := comp.Query("datacube/getusershare", from, to, nil, result)

	return result, err
}

// 获取图文分享转发分时数据.
// https://developers.weixin.qq.com/doc/offiaccount/Analytics/Graphic_Analysis_Data_Interface.html
func (comp *Client) UserShareHourly(from string, to string) (*response.ResponseDataCubeUserShareHourly, error) {

	result := &response.ResponseDataCubeUserShareHourly{}

	err := comp.Query("datacube/getusersharehour", from, to, nil, result)

	return result, err
}

// 获取消息发送概况数据.
// https://developers.weixin.qq.com/doc/offiaccount/Analytics/Message_analysis_data_interface.html
func (comp *Client) UpstreamMessageSummary(from string, to string) (*response.ResponseDataCubeUpstreamMessageSummary, error) {

	result := &response.ResponseDataCubeUpstreamMessageSummary{}

	err := comp.Query("datacube/getupstreammsg", from, to, nil, result)

	return result, err
}

// 获取图文分享转发分时数据.
// https://developers.weixin.qq.com/doc/offiaccount/Analytics/Message_analysis_data_interface.html
func (comp *Client) UpstreamMessageHourly(from string, to string) (*response.ResponseDataCubeUpstreamMessageHourly, error) {

	result := &response.ResponseDataCubeUpstreamMessageHourly{}

	err := comp.Query("datacube/getupstreammsghour", from, to, nil, result)

	return result, err
}

// 获取消息发送周数据.
// https://developers.weixin.qq.com/doc/offiaccount/Analytics/Message_analysis_data_interface.html
func (comp *Client) UpstreamMessageWeekly(from string, to string) (*response.ResponseDataCubeUpstreamMessageWeekly, error) {

	result := &response.ResponseDataCubeUpstreamMessageWeekly{}

	err := comp.Query("datacube/getupstreammsgweek", from, to, nil, result)

	return result, err
}

// 获取消息发送月数据.
// https://developers.weixin.qq.com/doc/offiaccount/Analytics/Message_analysis_data_interface.html
func (comp *Client) UpstreamMessageMonthly(from string, to string) (*response.ResponseDataCubeUpstreamMessageMonthly, error) {

	result := &response.ResponseDataCubeUpstreamMessageMonthly{}

	err := comp.Query("datacube/getupstreammsgmonth", from, to, nil, result)

	return result, err
}

// 获取消息发送分布数据.
// https://developers.weixin.qq.com/doc/offiaccount/Analytics/Message_analysis_data_interface.html
func (comp *Client) UpstreamMessageDistSummary(from string, to string) (*response.ResponseDataCubeUpstreamMessageDistSummary, error) {

	result := &response.ResponseDataCubeUpstreamMessageDistSummary{}

	err := comp.Query("datacube/getupstreammsgdist", from, to, nil, result)

	return result, err
}

// 获取消息发送分布周数据.
// https://developers.weixin.qq.com/doc/offiaccount/Analytics/Message_analysis_data_interface.html
func (comp *Client) UpstreamMessageDistWeekly(from string, to string) (*response.ResponseDataCubeUpstreamMessageDistWeekly, error) {

	result := &response.ResponseDataCubeUpstreamMessageDistWeekly{}

	err := comp.Query("datacube/getupstreammsgdistweek", from, to, nil, result)

	return result, err
}

// 获取消息发送分布月数据.
// https://developers.weixin.qq.com/doc/offiaccount/Analytics/Message_analysis_data_interface.html
func (comp *Client) UpstreamMessageDistMonthly(from string, to string) (*response.ResponseDataCubeUpstreamMessageDistMonthly, error) {

	result := &response.ResponseDataCubeUpstreamMessageDistMonthly{}

	err := comp.Query("datacube/getupstreammsgdistmonth", from, to, nil, result)

	return result, err
}

// 获取接口分析数据.
// https://developers.weixin.qq.com/doc/offiaccount/Analytics/Analytics_API.html
func (comp *Client) InterfaceSummary(from string, to string) (*response.ResponseDataCubeInterfaceSummary, error) {

	result := &response.ResponseDataCubeInterfaceSummary{}

	err := comp.Query("datacube/getinterfacesummary", from, to, nil, result)

	return result, err
}

// 获取接口分析分时数据.
// https://developers.weixin.qq.com/doc/offiaccount/Analytics/Analytics_API.html
func (comp *Client) InterfaceSummaryHourly(from string, to string) (*response.ResponseDataCubeInterfaceSummaryHourly, error) {

	result := &response.ResponseDataCubeInterfaceSummaryHourly{}

	err := comp.Query("datacube/getinterfacesummaryhour", from, to, nil, result)

	return result, err
}

// 获取免费券数据接口.
// https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html
func (comp *Client) FreeCardSummary(from string, to string, condSource int, cardID string) (*response.ResponseDataCubeFreeCardSummary, error) {

	result := &response.ResponseDataCubeFreeCardSummary{}

	ext := &object.HashMap{
		"cond_source": condSource,
		"card_id":     cardID,
	}

	err := comp.Query("datacube/getcardcardinfo", from, to, ext, result)

	return result, err
}

// 拉取会员卡数据接口.
// https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html
func (comp *Client) MemberCardSummary(from string, to string, condSource int, cardID string) (*response.ResponseDataCubeMemberCardSummary, error) {

	result := &response.ResponseDataCubeMemberCardSummary{}

	ext := &object.HashMap{
		"cond_source": condSource,
		"card_id":     cardID,
	}

	err := comp.Query("datacube/getcardmembercardinfo", from, to, ext, result)

	return result, err
}

// 拉取单张会员卡数据接口.
// https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html
func (comp *Client) memberCardSummaryByID(from string, to string, cardID string) (*response.ResponseDataCubeMemberCardSummaryByID, error) {

	result := &response.ResponseDataCubeMemberCardSummaryByID{}

	ext := &object.HashMap{
		"card_id": cardID,
	}

	err := comp.Query("datacube/getcardmembercarddetail", from, to, ext, result)

	return result, err
}

func (comp *Client) Query(api string, from string, to string, ext *object.HashMap, result interface{}) (err error) {

	params := &object.HashMap{
		"begin_date": from,
		"end_date":   to,
	}
	params = object.MergeHashMap(params, ext)

	_, err = comp.BaseClient.HttpPostJson(api, params, nil, nil, result)

	return err

}
