package request

import "github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/messageTemplate/request"

type Conclusions struct {
	Text        *request.TextOfMessage `json:"text"`
	Image       *request.Image         `json:"image"`
	Link        *request.Link          `json:"link"`
	MiniProgram *request.MiniProgram   `json:"miniprogram"`
}

type RequestAddContactWay struct {
	Type          int      `json:"type"`            // :1,
	Scene         int      `json:"scene"`           // 1,
	Style         int      `json:"style"`           // 1,
	Remark        string   `json:"remark"`          // "渠道客户",
	SkipVerify    bool     `json:"skip_verify"`     // true,
	State         string   `json:"state"`           // "teststate",
	User          []string `json:"user"`            // : ["zhangsan", "lisi", "wangwu"],
	Party         []int    `json:"party"`           // : [2, 3],
	IsTemp        bool     `json:"is_temp"`         // true,
	ExpiresIn     int      `json:"expires_in"`      // 86400,
	ChatExpiresIn int      `json:"chat_expires_in"` // 86400,
	UnionID       string   `json:"unionid"`         // "oxTWIuGaIt6gTKsQRLau2M0AAAA",

	Conclusions *Conclusions `json:"conclusions"`
}

// ------------------------------------------------------------------------
type RequestUpdateContactWay struct {
	ConfigID      string       `json:"config_id"`
	Remark        string       `json:"remark"`
	SkipVerify    bool         `json:"skip_verify"`
	Style         int          `json:"style"`
	State         string       `json:"state"`
	User          []string     `json:"user"`
	Party         []int        `json:"party"`
	ExpiresIn     int          `json:"expires_in"`
	ChatExpiresIn int          `json:"chat_expires_in"`
	UnionID       string       `json:"unionid"`
	Conclusions   *Conclusions `json:"conclusions"`
}
