package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type ResponseGetSimpleUserList struct {
	response.ResponseWork

	UserList []*UserSimpleDetail `json:"userlist"`
}

type UserSimpleDetail struct {
	UserID     string `json:"userid"`      //  "zhangsan",
	Name       string `json:"name"`        //  "张三",
	Department []int  `json:"department"`  //  [1, 2],
	OpenUserID string `json:"open_userid"` //  "xxxxxx"
}
