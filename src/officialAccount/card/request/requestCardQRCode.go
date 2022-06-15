package request

type CardQRCode struct {
	CardID   string `json:"card_id"`
	Code     string `json:"code"`
	OuterStr string `json:"outer_str"`
}

type MultiCard struct {
	CardList []*CardQRCode `json:"card_list"`
}

type RequestCreateQrCode struct {
	ActionName string `json:"action_name"`
	ActionInfo struct {
		MultipleCard *MultiCard  `json:"multiple_card"`
		Card         *CardQRCode `json:"card"`
	} `json:"action_info"`
}
