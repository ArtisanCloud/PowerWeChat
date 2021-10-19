package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseRiskControlGetUserRiskRank struct {
	*response.ResponseMiniProgram
	RiskRank int64 `json:"risk_rank"`
	UnoinID  int64 `json:"unoin_id"`
}
