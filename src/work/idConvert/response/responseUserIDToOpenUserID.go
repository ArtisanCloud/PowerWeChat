package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

type ResponseUserIDToOpenUserID struct {
	response.ResponseWork
	// OpenUserIDList 该服务商主体下的密文userid
	OpenUserIDList []UserIDToOpenUserIDResult `json:"open_userid_list,omitempty"`
}

// UserIDToOpenUserIDResult 该服务商主体下的密文userid
type UserIDToOpenUserIDResult struct {
	// UserID 转换成功的userid
	UserID string `json:"userid,omitempty"`
	// OpenUserID 转换成功的userid对应的服务商主体下的密文userid
	OpenUserID string `json:"open_userid,omitempty"`
}
