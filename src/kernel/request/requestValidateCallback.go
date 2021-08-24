package request

// https://work.weixin.qq.com/api/doc/90000/90135/90930

type RequestValidateCallback struct {
	// touser、toparty、totag不能同时为空
	MsgSignature string `xml:"msg_signature"`
	Timestamp    string `xml:"timestamp"`
	Nonce        string `xml:"nonce"`
	EchoStr      string `xml:"echostr"`
}
