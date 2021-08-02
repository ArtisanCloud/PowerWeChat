package oauth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ParaOAuthCallbackQRCode struct {
	Code string `form:"code" json:"code" xml:"code" binding:"required"`
	AppID string `form:"appid" json:"appid" xml:"appid" binding:"required"`
	State string `form:"state" json:"state" xml:"state"`
}


func ValidateRequestOAuthCallbackQRCode(context *gin.Context) {
	var form ParaOAuthCallback

	//logger.Info("validate make reservation", nil)
	if err := context.ShouldBind(&form); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, "Lack of parameters!")
	}

	context.Set("params", form)

	context.Next()
}