package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

type ResponseOCRVehicleLicense struct {
	response.ResponseMiniProgram

	VhicleType     string `json:"vhicle_type"`     // "小型普通客⻋",
	Owner          string `json:"owner"`           // "东莞市xxxxx机械厂",
	Addr           string `json:"addr"`            // "广东省东莞市xxxxx号",
	UseCharacter   string `json:"use_character"`   // "非营运",
	Model          string `json:"model"`           // "江淮牌HFCxxxxxxx",
	Vin            string `json:"vin"`             // "LJ166xxxxxxxx51",
	EngineNum      string `json:"engine_num"`      // "J3xxxxx3",
	RegisterDate   string `json:"register_date"`   // "2018-07-06",
	IssueDate      string `json:"issue_date"`      // "2018-07-01",
	PlateNumB      string `json:"plate_num_b"`     // "粤xxxxx",
	Record         string `json:"record"`          // "441xxxxxx3",
	PassengersNum  string `json:"passengers_num"`  // "7人",
	TotalQuality   string `json:"total_quality"`   // "2700kg",
	PrepareQuality string `json:"prepare_quality"` // "1995kg"
}
