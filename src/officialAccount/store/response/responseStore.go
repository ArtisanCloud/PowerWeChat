package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	response2 "github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/poi/response"
)

type Inner struct {
	Name string `json:"name"`
}

type Exter struct {
	InnerList []*Inner `json:"inner_list"`
}

type Qualify struct {
	ExterList []*Exter `json:"exter_list"`
}

type Category struct {
	ID            int      `json:"id"`
	Name          string   `json:"name"`
	Level         int      `json:"level"`
	Children      []int    `json:"children"`
	Father        int      `json:"father,omitempty"`
	Qualify       *Qualify `json:"qualify,omitempty"`
	Scene         int      `json:"scene,omitempty"`
	SensitiveType int      `json:"sensitive_type,omitempty"`
}

type AllCategoryInfo struct {
	Categories []*Category `json:"categories"`
}

type Store struct {
	AllCategoryInfo *AllCategoryInfo `json:"all_category_info"`
}

type ResponseStoreCategory struct {
	response.ResponseOfficialAccount

	Data Store `json:"data"`
}

// -------------------------------------------------------

type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type Result struct {
	ID       string    `json:"id"`
	Name     string    `json:"name,omitempty"`
	FullName string    `json:"fullname"`
	PinYin   []string  `json:"pinyin,omitempty"`
	Location *Location `json:"location"`
	CIDX     []int     `json:"cidx,omitempty"`
}

type ResponseStoreDistrict struct {
	response.ResponseOfficialAccount

	Status      int         `json:"status"`
	Message     string      `json:"message"`
	DataVersion string      `json:"data_version"`
	Result      [][]*Result `json:"result"`
}

// -------------------------------------------------------

type DataPIO struct {
	Item []interface{} `json:"item"`
}

type ResponseStoreSearchMapPIO struct {
	response.ResponseOfficialAccount

	Data *DataPIO `json:"data"`
}

// -------------------------------------------------------

type ResponseStoreGetStatus struct {
	response.ResponseOfficialAccount

	FirstCatID        int    `json:"first_catid"`
	SecondCatID       int    `json:"second_catid"`
	QualificationList string `json:"qualification_list"`
	HeadImgMediaID    string `json:"headimg_mediaid"`
	NickName          string `json:"nickname"`
	Intro             string `json:"intro"`
	OrgCode           string `json:"org_code"`
	OtherFiles        string `json:"other_files"`
}

// -------------------------------------------------------

type CreateStoreFromMap struct {
	BaseID int `json:"base_id"`
	RichID int `json:"rich_id"`
}

type ResponseStoreCreateFromMap struct {
	response.ResponseOfficialAccount

	Error interface{}        `json:"error"`
	Data  CreateStoreFromMap `json:"data"`
}

// -------------------------------------------------------

type CreateStore struct {
	AuditID int `json:"audit_id"`
}

type ResponseStoreCreate struct {
	response.ResponseOfficialAccount

	Data CreateStore `json:"data"`
}

// -------------------------------------------------------

type UpdateStore struct {
	HasAuditId int `json:"has_audit_id"`
	AuditId    int `json:"audit_id"`
}

type ResponseStoreUpdate struct {
	response.ResponseOfficialAccount

	Data UpdateStore `json:"data"`
}

// -------------------------------------------------------

type ResponseStoreInfo struct {
	response.ResponseOfficialAccount

	Business response2.Business `json:"business"`
}

// -------------------------------------------------------

type ResponseStoreList struct {
	response.ResponseOfficialAccount

	BusinessList []*response2.Business `json:"business_list"`
	TotalCount   int                   `json:"total_count"`
}
