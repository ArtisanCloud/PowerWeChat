package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

type PreAuthCode struct {
	response.ResponseWork
	// PreAuthCode 预授权码,最长为512字节
	PreAuthCode string `json:"pre_auth_code,omitempty"`
	// ExpiresIn 有效期
	ExpiresIn int64 `json:"expires_in,omitempty"`
}
