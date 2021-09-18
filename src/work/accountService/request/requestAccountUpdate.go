package request

type RequestAccountUpdate struct {
	OpenKFID string `json:"open_kfid"`
	Name     string `json:"name"`
	MediaID  string `json:"media_id"`
}
