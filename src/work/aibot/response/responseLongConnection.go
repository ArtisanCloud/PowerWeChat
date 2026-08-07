package response

type LongConnectionHeaders struct {
	ReqID string `json:"req_id,omitempty"`
}

type ResponseLongConnection struct {
	Cmd     string                 `json:"cmd,omitempty"`
	Headers *LongConnectionHeaders `json:"headers,omitempty"`
	Body    map[string]any         `json:"body,omitempty"`
	ErrCode int                    `json:"errcode,omitempty"`
	ErrMsg  string                 `json:"errmsg,omitempty"`
}

func (resp *ResponseLongConnection) IsError() bool {
	if resp == nil {
		return true
	}
	return resp.ErrCode != 0
}
