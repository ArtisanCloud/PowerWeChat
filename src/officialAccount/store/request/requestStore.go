package request

type BaseInfo struct {
	Name         string `json:"name"`
	Longitude    string `json:"longitude"`
	Latitude     string `json:"latitude"`
	Province     string `json:"province"`
	City         string `json:"city"`
	District     string `json:"district"`
	Address      string `json:"address"`
	Category     string `json:"category"`
	Telephone    string `json:"telephone"`
	Photo        string `json:"photo"`
	License      string `json:"license"`
	Introduction string `json:"introduct"`
	DistrictID   string `json:"districtid"`
}

type RequestStoreCreate struct {
	PoiID             string `json:"poi_id"`
	MapPoiID          string `json:"map_poi_id"`
	PicList           string `json:"pic_list"`
	ContractPhone     string `json:"contract_phone"`
	Credential        string `json:"credential"`
	QualificationList string `json:"qualification_list"`
}

type RequestStoreUpdate struct {
	MapPoiID      string `json:"map_poi_id"`
	PoiID         string `json:"poi_id"`
	Hour          string `json:"hour"`
	ContractPhone string `json:"contract_phone"`
	PicList       string `json:"pic_list"`
}

