package response

type FreeCardSummary struct {
	RefDate     string `json:"ref_date"`
	CardID      string `json:"card_id"`
	CardType    int    `json:"card_type"`
	ViewCnt     int    `json:"view_cnt"`
	ViewUser    int    `json:"view_user"`
	ReceiveCnt  int    `json:"receive_cnt"`
	ReceiveUser int    `json:"receive_user"`
	VerifyCnt   int    `json:"verify_cnt"`
	VerifyUser  int    `json:"verify_user"`
	GivenCnt    int    `json:"given_cnt"`
	GivenUser   int    `json:"given_user"`
	ExpireCnt   int    `json:"expire_cnt"`
	ExpireUser  int    `json:"expire_user"`
}

type ResponseDataCubeFreeCardSummary struct {
	List []*FreeCardSummary `json:"list"`
}

// ----------------------------------------------------------------------

type MemberCardSummary struct {
	RefDate          string `json:"ref_date"`
	ViewCnt          int    `json:"view_cnt"`
	ViewUser         int    `json:"view_user"`
	ReceiveCnt       int    `json:"receive_cnt"`
	ReceiveUser      int    `json:"receive_user"`
	ActiveUser       int    `json:"active_user"`
	VerifyCNT        int    `json:"verify_cnt"`
	VerifyUser       int    `json:"verify_user"`
	TotalUser        int    `json:"total_user"`
	TotalReceiveUser int    `json:"total_receive_user"`
}

type ResponseDataCubeMemberCardSummary struct {
	List []*MemberCardSummary `json:"list"`
}



// ----------------------------------------------------------------------

type MemberCardSummaryByID struct {
	RefDate        string `json:"ref_date"`
	MerchantType   int    `json:"merchanttype"`
	CardID         string `json:"cardid"`
	SubMerchantID  int    `json:"submerchantid"`
	ViewCNT        int    `json:"view_cnt"`
	ViewUser       int    `json:"view_user"`
	ReceiveCNT     int    `json:"receive_cnt"`
	ReceiveUser    int    `json:"receive_user"`
	VerifyCnt      int    `json:"verify_cnt"`
	VerifyUser     int    `json:"verify_user"`
	ActiveCnt      int    `json:"active_cnt"`
	ActiveUser     int    `json:"active_user"`
	TotalUser      int    `json:"total_user"`
	NewUser        int    `json:"new_user"`
	PayOriginalFee int    `json:"payOriginalFee"`
	Fee            int    `json:"fee"`
}

type ResponseDataCubeMemberCardSummaryByID struct {
	List []*MemberCardSummaryByID `json:"list"`
}

// ----------------------------------------------------------------------




// ----------------------------------------------------------------------

