package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

type ResponseGetPublicKey struct {
	response.ResponsePayment

	MchID  string `xml:"mch_id"`
	PubKey string `xml:"pub_key"`
}
