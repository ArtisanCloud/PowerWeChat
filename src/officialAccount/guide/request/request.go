package request

type FastReply struct {
	Content string `json:"content"`
}

type FastReplyList struct {
	GuideFastReplyList []*FastReply `json:"guide_fast_reply_list"`
}


type AutoReply struct {
	Content string `json:"content"`
	MsgType string `json:"msgtype"`
}




type BlackKeyword struct {
	Values []string `json:"values"`
}


// ------------------------------------------------------------

type Buyer struct {
	OpenID        string `json:"openid"`
	BuyerNickname string `json:"buyer_nickname"`
}

type BuyerList struct {
	GuideAccount string `json:"guide_account"`
	BuyerList    []Buyer `json:"buyer_list"`
}