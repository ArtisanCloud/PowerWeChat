package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"time"
)

type ResponseMediaUpload struct {
	response.ResponsePayment
	MediaId string `json:"media_id"`
}

type ReturnAddressInfo struct {
	ReturnAddress string `json:"return_address"`
	Longitude     string `json:"longitude"`
	Latitude      string `json:"latitude"`
}

type SharePowerInfo struct {
	ReturnTime        time.Time         `json:"return_time"`
	ReturnAddressInfo ReturnAddressInfo `json:"return_address_info"`
}

type AdditionalInfo struct {
	Type           string         `json:"type"`
	SharePowerInfo SharePowerInfo `json:"share_power_info"`
}

type ServiceOrderInfo struct {
	OrderId    string `json:"order_id"`
	OutOrderNo string `json:"out_order_no"`
	State      string `json:"state"`
}

type ComplaintMediaList struct {
	MediaType string   `json:"media_type"`
	MediaUrl  []string `json:"media_url"`
}

type ComplainOrderInfo struct {
	TransactionId string `json:"transaction_id"`
	OutTradeNo    string `json:"out_trade_no"`
	Amount        int    `json:"amount"`
}

type ResponseQueryComplaint struct {
	response.ResponsePayment

	ComplaintId           string               `json:"complaint_id"`
	ComplaintTime         time.Time            `json:"complaint_time"`
	ComplaintDetail       string               `json:"complaint_detail"`
	ComplaintState        string               `json:"complaint_state"`
	PayerPhone            string               `json:"payer_phone"`
	ComplaintOrderInfo    []ComplainOrderInfo  `json:"complaint_order_info"`
	ComplaintFullRefunded bool                 `json:"complaint_full_refunded"`
	IncomingUserResponse  bool                 `json:"incoming_user_response"`
	UserComplaintTimes    int                  `json:"user_complaint_times"`
	ComplaintMediaList    []ComplaintMediaList `json:"complaint_media_list"`
	ProblemDescription    string               `json:"problem_description"`
	ProblemType           string               `json:"problem_type"`
	ApplyRefundAmount     int                  `json:"apply_refund_amount"`
	UserTagList           []string             `json:"user_tag_list"`
	ServiceOrderInfo      []ServiceOrderInfo   `json:"service_order_info"`
	AdditionalInfo        AdditionalInfo       `json:"additional_info"`
}

type ResponseComplaints struct {
	response.ResponsePayment

	Data       []ResponseQueryComplaint `json:"data"`
	Limit      int                      `json:"limit"`
	Offset     int                      `json:"offset"`
	TotalCount int                      `json:"total_count"`
}
