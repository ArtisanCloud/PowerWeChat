package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type ResponseOCRBusinessLicense struct {
	response.ResponseMiniProgram
	RegNum              string         `json:"reg_num"`              //注册号
	Serial              string         `json:"serial"`               //编号
	LegalRepresentative string         `json:"legal_representative"` //法定代表人姓名
	EnterpriseName      string         `json:"enterprise_name"`      //企业名称
	TypeOfOrganization  string         `json:"type_of_organization"` //组成形式
	Address             string         `json:"address"`              //经营场所/企业住所
	TypeOfEnterprise    string         `json:"type_of_enterprise"`   //公司类型
	BusinessScope       string         `json:"business_scope"`       //经营范围
	RegisteredCapital   string         `json:"registered_capital"`   //注册资本
	PaidInCapital       string         `json:"paid_in_capital"`      //实收资本
	ValidPeriod         string         `json:"valid_period"`         //营业期限
	RegisteredDate      string         `json:"registered_date"`      //注册日期/成立日期
	CertPosition        *power.HashMap `json:"cert_position"`        //营业执照位置
	ImgSize             *power.HashMap `json:"img_size"`             //营业执照大小

}
