package request

// https://open.work.weixin.qq.com/api/doc/90000/90135/90236#文件消息

type RequestMessageSend struct {
	// touser、toparty、totag不能同时为空
	ToUser  string `json:"touser"`  // "UserID1|UserID2|UserID3",
	ToParty string `json:"toparty"` // "PartyID1|PartyID2",
	ToTag   string `json:"totag"`   // "TagID1 | TagID2",
	MsgType string `json:"msgtype"` // "text",
	AgentID int    `json:"agentid"` // 1,

	Safe                   int `json:"safe"`                     //  0,
	EnableIDTrans          int `json:"enable_id_trans"`          //  0,
	EnableDuplicateCheck   int `json:"enable_duplicate_check"`   //  0,
	DuplicateCheckInterval int `json:"duplicate_check_interval"` //  1800
}
