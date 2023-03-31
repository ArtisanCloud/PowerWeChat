package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type ResponseOCRDrivingLicense struct {
	response.ResponseMiniProgram
	IdNum        string `json:"id_num"`
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
