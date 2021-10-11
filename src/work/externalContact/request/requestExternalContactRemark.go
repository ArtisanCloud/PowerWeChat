package request

type RequestExternalContactRemark struct {
	UserID           string   `json:"userid"`
	ExternalUserID   string   `json:"external_userid"`
	Remark           string   `json:"remark"`
	Description      string   `json:"description"`
	RemarkCompany    string   `json:"remark_company"`
	RemarkMobiles    []string `json:"remark_mobiles"`
	RemarkPicMediaID string   `json:"remark_pic_mediaid"`
}
