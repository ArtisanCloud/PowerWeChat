package response

import "github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"

type ResponseOCRIDCard struct {
	*response.ResponseOfficialAccount

	Type        string `json:"type"`
	Name        string `json:"name"`
	ID          string `json:"id"`
	Addr        string `json:"addr"`
	Gender      string `json:"gender"`
	Nationality string `json:"nationality"`
}

// ----------------------------------------

type ResponseOCRBankCard struct {
	*response.ResponseOfficialAccount

	Number string `json:"number"`
}

// ----------------------------------------

type ResponseOCRVehicleLicense struct {
	*response.ResponseOfficialAccount

	IDNum        string `json:"id_num"`
	Name         string `json:"name"`
	Sex          string `json:"sex"`
	Nationality  string `json:"nationality"`
	Address      string `json:"address"`
	BirthDate    string `json:"birth_date"`
	IssueDate    string `json:"issue_date"`
	CarClass     string `json:"car_class"`
	ValidFrom    string `json:"valid_from"`
	ValidTo      string `json:"valid_to"`
	OfficialSeal string `json:"official_seal"`
}

// ----------------------------------------

type ResponseOCRDriving struct {
	*response.ResponseOfficialAccount

	IDNum        string `json:"id_num"`
	Name         string `json:"name"`
	Sex          string `json:"sex"`
	Nationality  string `json:"nationality"`
	Address      string `json:"address"`
	BirthDate    string `json:"birth_date"`
	IssueDate    string `json:"issue_date"`
	CarClass     string `json:"car_class"`
	ValidFrom    string `json:"valid_from"`
	ValidTo      string `json:"valid_to"`
	OfficialSeal string `json:"official_seal"`
}

// ----------------------------------------

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Position struct {
	LeftTop     *Point `json:"left_top"`
	RightTop    *Point `json:"right_top"`
	RightBottom *Point `json:"right_bottom"`
	LeftBottom  *Point `json:"left_bottom"`
}

type CertPosition struct {
	Pos *Position `json:"pos"`
}

type ImgSize struct {
	W int `json:"w"`
	H int `json:"h"`
}

type ResponseOCRBizLicense struct {
	*response.ResponseOfficialAccount

	RegNum              string        `json:"reg_num"`
	Serial              string        `json:"serial"`
	LegalRepresentative string        `json:"legal_representative"`
	EnterpriseName      string        `json:"enterprise_name"`
	TypeOfOrganization  string        `json:"type_of_organization"`
	Address             string        `json:"address"`
	TypeOfEnterprise    string        `json:"type_of_enterprise"`
	BusinessScope       string        `json:"business_scope"`
	RegisteredCapital   string        `json:"registered_capital"`
	PaidInCapital       string        `json:"paid_in_capital"`
	ValidPeriod         string        `json:"valid_period"`
	RegisteredDate      string        `json:"registered_date"`
	CertPosition        *CertPosition `json:"cert_position"`
	ImgSize             *ImgSize      `json:"img_size"`
}

// ----------------------------------------

type Item struct {
	Text string   `json:"text"`
	Pos  Position `json:"pos"`
}

type ResponseOCRCommon struct {
	*response.ResponseOfficialAccount

	Items   []Item  `json:"items"`
	ImgSize ImgSize `json:"img_size"`
}

// ----------------------------------------

type ResponseOCRPlateNumber struct {
	*response.ResponseOfficialAccount

	Number string `json:"number"`
}
