package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

type Photo struct {
	PhotoUrl string `json:"photo_url"`
}

type BusinessInfo struct {
	PoiID             string   `json:"poi_id"`
	SID               string   `json:"sid"`
	BusinessName      string   `json:"business_name"`
	BranchName        string   `json:"branch_name"`
	Province          string   `json:"province"`
	City              string   `json:"city"`
	Address           string   `json:"address"`
	Telephone         string   `json:"telephone"`
	Categories        []string `json:"categories"`
	OffsetType        int      `json:"offset_type"`
	Longitude         float64  `json:"longitude"`
	Latitude          float64  `json:"latitude"`
	PhotoList         []*Photo `json:"photo_list"`
	Recommend         string   `json:"recommend"`
	Special           string   `json:"special"`
	Introduction      string   `json:"introduction"`
	OpenTime          string   `json:"open_time"`
	AvgPrice          int      `json:"avg_price"`
	AvailableState    int      `json:"available_state"`
	UpdateStatus      int      `json:"update_status"`
	Status            int      `json:"status"`
	District          string   `json:"district"`
	QualificationNum  string   `json:"qualification_num"`
	QualificationName string   `json:"qualification_name"`
}

type Business struct {
	BaseInfo *BusinessInfo `json:"base_info"`
}

type ResponsePOIGet struct {
	response.ResponseOfficialAccount

	Business *Business `json:"business"`
}

// ------------------------------------------------------------

type ResponsePOIList struct {
	BaseInfo *BusinessInfo `json:"base_info"`

	BusinessList []*Business `json:"business_list"`
	TotalCount   string      `json:"total_count"`
}
