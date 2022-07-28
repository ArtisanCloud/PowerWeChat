package request

// https://developer.work.weixin.qq.com/document/path/90236#文件消息
type RequestMessageSendInterface interface {
	GetAgentID() int
	SetAgentID(agentID int) RequestMessageSendInterface
}

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

func (request *RequestMessageSend) GetAgentID() int {
	return request.AgentID
}

func (request *RequestMessageSend) SetAgentID(agentID int) RequestMessageSendInterface {
	request.AgentID = agentID
	return request
}
