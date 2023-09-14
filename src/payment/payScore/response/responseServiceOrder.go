package response

type PostPayment struct {
	Name        string `json:"name"`
	Amount      int    `json:"amount"`
	Description string `json:"description"`
	Count       int    `json:"count"`
}

type PostDiscount struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Amount      int    `json:"amount"`
	Count       int    `json:"count"`
}

type RisFund struct {
	Name        string `json:"name"`
	Amount      int    `json:"amount"`
	Description string `json:"description"`
}

type TimeRange struct {
	StartTime       string `json:"start_time"`
	EndTime         string `json:"end_time"`
	StartTimeRemark string `json:"start_time_remark"`
	EndTimeRemark   string `json:"end_time_remark"`
}

type Location struct {
	StartLocation string `json:"start_location"`
	EndLocation   string `json:"end_location"`
}

type ResponseServiceOrder struct {
	AppId               string         `json:"appid"`
	MchId               string         `json:"mchid"`
	OutOrderNo          string         `json:"out_order_no"`
	ServiceId           string         `json:"service_id"`
	ServiceIntroduction string         `json:"service_introduction"`
	State               string         `json:"state"`
	StateDescription    string         `json:"state_description"`
	PostPayments        []PostPayment  `json:"post_payments"`
	PostDiscounts       []PostDiscount `json:"post_discounts"`
	RiskFund            RisFund        `json:"risk_fund"`
	TimeRange           TimeRange      `json:"time_range"`
	Location            Location       `json:"location"`
	Attach              string         `json:"attach"`
	NotifyUrl           string         `json:"notify_url"`
	OrderId             string         `json:"order_id"`
	Package             string         `json:"package"`
}
