package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

type ResponseOpenUserIDToUserID struct {
	response.ResponseWork
	UserIDList []OpenUserIDToUserIDResult `json:"userid_list,omitempty"`
	// InvalidUserIDList 不合法的open_userid列表
	InvalidUserIDList []string `json:"invalid_userid_list,omitempty"`
}

type OpenUserIDToUserIDResult struct {
	// UserID 转换成功的userid
	UserID string `json:"userid,omitempty"`
	// OpenUserID 转换成功的userid对应的服务商主体下的密文userid
	OpenUserID string `json:"open_userid,omitempty"`
}
