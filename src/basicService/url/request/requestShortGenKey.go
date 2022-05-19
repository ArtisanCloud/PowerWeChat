package request

type RequestShortGenKey struct {
	LongData      string `json:"long_data"`
	ExpireSeconds int    `json:"expire_seconds"`
}