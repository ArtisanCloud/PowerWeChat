package request

// RequestServerMarket
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/service-market/serviceMarket.invokeService.html
type RequestServiceMarket struct {
	Service     string `json:"service"`
	Api         string `json:"api"`
	ClientMsgID string `json:"client_msg_id"`

	// 服务提供方接口定义的 JSON 格式的数据。
	// 因为微信文档说是string，但是给的例子又是json，所以这边不限制类型。 string或者hashmap类型都无所谓
	Data interface{} `json:"data"`
}
