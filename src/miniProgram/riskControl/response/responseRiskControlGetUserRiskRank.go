package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type ResponseRiskControlGetUserRiskRank struct {
	response.ResponseMiniProgram
	RiskRank int64 `json:"risk_rank"`
	UnoinID  int64 `json:"unoin_id"`
}
