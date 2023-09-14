package request

type PostPayment struct {
	Name        string `json:"name"`
	Amount      int    `json:"amount"`
	Description string `json:"description"`
	Count       int    `json:"count"`
}

type PostDiscount struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Count       int    `json:"count"`
}

type TimeRage struct {
	StartTime       string `json:"start_time"`
	EndTime         string `json:"end_time"`
	StartTimeRemark string `json:"start_time_remark"`
	EndTimeRemark   string `json:"end_time_remark"`
}

type Location struct {
	StartLocation string `json:"start_location"`
	EndLocation   string `json:"end_location"`
}

type RiskFund struct {
	Name        string `json:"name"`
	Amount      int    `json:"amount"`
	Description string `json:"description"`
}

type RequestServiceOrder struct {
	OutOrderNo          string         `json:"out_order_no"`
	Appid               string         `json:"appid"`
	ServiceId           string         `json:"service_id"`
	ServiceIntroduction string         `json:"service_introduction"`
	PostPayments        []PostPayment  `json:"post_payments"`
	PostDiscounts       []PostDiscount `json:"post_discounts"`
	TimeRange           TimeRage       `json:"time_range"`
	Location            Location       `json:"location"`
	RiskFund            RiskFund       `json:"risk_fund"`
	Attach              string         `json:"attach"`
	NotifyUrl           string         `json:"notify_url"`
	Openid              string         `json:"openid"`
	NeedUserConfirm     bool           `json:"need_user_confirm"`
}
