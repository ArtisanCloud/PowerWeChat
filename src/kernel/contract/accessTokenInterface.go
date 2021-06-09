package contract

import (
	"github.com/ArtisanCloud/go-libs/object"
	"net/http"
)

type (
	AccessTokenInterface interface {
		GetToken() string
		Refresh() AccessTokenInterface
		ApplyToRequest(request *http.Request, requestOptions *object.HashMap) *http.Request
	}
)