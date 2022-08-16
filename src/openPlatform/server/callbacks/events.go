package openplatform

import "encoding/xml"

const EVENT_AUTHORIZED string = "authorized"
const EVENT_UNAUTHORIZED string = "unauthorized"
const EVENT_UPDATE_AUTHORIZED string = "updateauthorized"
const EVENT_COMPONENT_VERIFY_TICKET string = "component_verify_ticket"
const EVENT_THIRD_FAST_REGISTERED string = "notify_third_fasteregister"

type Callback struct {
	XMLName xml.Name `xml:"xml"`
	Text    string   `xml:",chardata"`
	AppID   string   `xml:"AppId"`
	Encrypt string   `xml:"Encrypt"`
}

// ------------------------------------------------------------
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Before_Develop/component_verify_ticket.html
type EventVerifyTicket struct {
	XMLName               xml.Name `xml:"xml"`
	Text                  string   `xml:",chardata"`
	CreateTime            string   `xml:"CreateTime"`
	InfoType              string   `xml:"InfoType"`
	ComponentVerifyTicket string   `xml:"ComponentVerifyTicket"`
}

// 授权变更通知推送
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Before_Develop/authorize_event.html
type EventAuthorization struct {
	XMLName                      xml.Name `xml:"xml"`
	Text                         string   `xml:",chardata"`
	AppId                        string   `xml:"AppId"`
	CreateTime                   string   `xml:"CreateTime"`
	InfoType                     string   `xml:"InfoType"`
	AuthorizerAppid              string   `xml:"AuthorizerAppid"`
	AuthorizationCode            string   `xml:"AuthorizationCode"`
	AuthorizationCodeExpiredTime string   `xml:"AuthorizationCodeExpiredTime"`
	PreAuthCode                  string   `xml:"PreAuthCode"`
}
