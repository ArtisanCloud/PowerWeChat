package response

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/go-wechat/src/kernel/response"
)

type ResponseGetContactWay struct {
	*response.ResponseWX

	Type          int             `json:"type"`            // :1,
	Scene         int             `json:"scene"`           // 1,
	Style         int             `json:"style"`           // 1,
	Remark        string          `json:"remark"`          // "渠道客户",
	SkipVerify    bool            `json:"skip_verify"`     // true,
	State         string          `json:"state"`           // "teststate",
	User          []string        `json:"user"`            // : ["zhangsan", "lisi", "wangwu"],
	Party         []int           `json:"party"`           // : [2, 3],
	IsTemp        bool            `json:"is_temp"`         // true,
	ExpiresIn     int             `json:"expires_in"`      // 86400,
	ChatExpiresIn int             `json:"chat_expires_in"` // 86400,
	UnionID       string          `json:"unionid"`         // "oxTWIuGaIt6gTKsQRLau2M0AAAA",
	Conclusions   *object.HashMap `json:"conclusions"`
}
