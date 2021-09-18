package response

import (
  "github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseBroadcastGetAssistantList struct {
  *response.ResponseMiniProgram
  List []struct {
    Timestamp int    `json:"timestamp"`
    HeadImg   string `json:"headimg"`
    Nickname  string `json:"nickname"`
    Alias     string `json:"alias"`
    OpenID    string `json:"openid"`
  } `json:"list"`
  Count    int `json:"count"`
  MaxCount int `json:"maxCount"`
}
