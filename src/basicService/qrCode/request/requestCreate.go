package request


type ActionInfo struct {
	Scene interface{} `json:"scene"`
}

type RequestQRCodeCreate struct {
	ActionName string      `json:"action_name"`
	ActionInfo *ActionInfo `json:"action_info"`
}
