package model

type TransferInfo struct {
	// HandoverUserID 转移成员的userid
	HandoverUserID string `json:"handover_userid,omitempty"`
	// TakeoverUserID 接收成员的userid
	TakeoverUserID string `json:"takeover_userid,omitempty"`
	Errcode        int    `json:"errcode,omitempty"`
}
