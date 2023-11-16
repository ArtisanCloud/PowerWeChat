package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

type CorpToken struct {
	response.ResponseWork
	// 授权方（企业）access_token,最长为512字节
	AccessToken string `json:"access_token,omitempty"`
	// ExpiresIn 授权方（企业）access_token超时时间
	ExpiresIn int64 `json:"expires_in,omitempty"`
}
