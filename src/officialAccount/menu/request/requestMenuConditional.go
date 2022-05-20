package request


type RequestMatchRule struct {
	TagID              string `json:"tag_id"`
	Sex                string `json:"sex"`
	Country            string `json:"country"`
	Province           string `json:"province"`
	City               string `json:"city"`
	ClientPlatformType string `json:"client_platform_type"`
	Language           string `json:"language"`
}
