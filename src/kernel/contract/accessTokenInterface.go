package contract

import (
	"github.com/ArtisanCloud/go-libs/object"
	response2 "github.com/ArtisanCloud/go-wechat/src/kernel/response"
	"net/http"
)

type (
	AccessTokenInterface interface {
		GetToken(refresh bool) (resToken *response2.ResponseGetToken, err error)
		Refresh() AccessTokenInterface
		ApplyToRequest(request *http.Request, requestOptions *object.HashMap) (*http.Request, error)
	}
)
