package request

type RequestCardInvoice struct {
	CardID      string `json:"card_id"`
	EncryptCode string `json:"encrypt_code"`
}