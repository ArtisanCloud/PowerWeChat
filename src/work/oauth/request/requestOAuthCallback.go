package request

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ParaOAuthCallback struct {
	Code string `form:"code" json:"code" xml:"code" binding:"required"`
}

func ValidateRequestOAuthCallback(context *gin.Context) {
	var form ParaOAuthCallback

	//logger.Info("validate make reservation", nil)
	if err := context.ShouldBind(&form); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, "Lack of parameters!")
	}

	context.Set("params", form)

	context.Next()
}
