package request

type RequestSoter struct {
	OpenID        string `json:"openid"`
	JsonString    string `json:"json_string"`
	JsonSignature string `json:"json_signature"`
}
