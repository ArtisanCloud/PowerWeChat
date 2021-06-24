package response

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/go-wechat/src/kernel/response"
)

type ResponseBatchGetByUser struct {
	response.ResponseWX
	ExternalContactList []*ResponseExternalContact `json:"external_contact_list"`
}

type ResponseExternalContact struct {
	ExternalContact *ResponseExternalContactDetail `json:"external_contact"`
	FollowInfo      *FollowInfo                    `json:"follow_info"`
}

type ResponseExternalContactDetail struct {
	ExternalUserID  string                   `json:"external_userid"`  // :"woAJ2GCAAAXtWyujaWJHDDGi0mACHAAA",
	Name            string                   `json:"name"`             // :"李四",
	Position        string                   `json:"position"`         // :"Manager",
	Avatar          string                   `json:"avatar"`           // :"http://p.qlogo.cn/bizmail/IcsdgagqefergqerhewSdage/0",
	CorpName        string                   `json:"corp_name"`        // :"腾讯",
	CorpFullName    string                   `json:"corp_full_name"`   // :"腾讯科技有限公司",
	Type            string                   `json:"type"`             // :2,
	Gender          string                   `json:"gender"`           // :1,
	UnionID         string                   `json:"unionid"`          // :"ozynqsulJFCZ2z1aYeS8h-nuasdAAA",
	ExternalProfile *ResponseExternalProfile `json:"external_profile"` // :
}

type ResponseExternalProfile struct {
	ExternalProfile []*object.HashMap `json:"external_attr"` // :
}

type FollowInfo struct {
	Remark      string              `json:"remark"`      // "王助理",
	Description string              `json:"description"` // "采购问题咨询",
	CreateTime  int `json:"createtime"`  // 1525881637,
	TagID       []string            `json:"tag_id"`      // ["etAJ2GCAAAXtWyujaWJHDDGi0mACHAAA"],
	State       string              `json:"state"`       // "外联二维码1",
	OperUserID  string              `json:"oper_userid"` // "woAJ2GCAAAd1asdasdjO4wKmE8AabjBBB",
	AddWay      int                 `json:"add_way"`     // 3
}
