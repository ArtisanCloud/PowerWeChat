package request

type ParaOAuthCallback struct {
	Code string `form:"code" json:"code" xml:"code" binding:"required"`
}

