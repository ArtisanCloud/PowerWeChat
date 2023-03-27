package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"time"
)

type ResponseGetPublicKey struct {
	response.ResponsePayment

	MchID  string `xml:"mch_id"`
	PubKey string `xml:"pub_key"`
}

type Certificate struct {
	SerialNo           string    `json:"serial_no"`
	EffectiveTime      time.Time `json:"effective_time "`
	ExpireTime         time.Time `json:"expire_time "`
	EncryptCertificate struct {
		Algorithm      string `json:"algorithm"`
		Nonce          string `json:"nonce"`
		AssociatedData string `json:"associated_data"`
		Ciphertext     string `json:"ciphertext"`
	} `json:"encrypt_certificate"`
}

type ResponseGetCertificates struct {
	Data []*Certificate `json:"data"`
}
