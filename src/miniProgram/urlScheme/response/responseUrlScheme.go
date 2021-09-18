package response

type ResponseUrlScheme struct {
  ErrCode  int    `json:"errcode"`
  ErrMsg   string `json:"errmsg"`
  OpenLink string `json:"openlink"`
}