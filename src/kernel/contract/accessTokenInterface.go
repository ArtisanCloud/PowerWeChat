package contract

import (
	"github.com/ArtisanCloud/go-libs/http/contract"
	"github.com/ArtisanCloud/go-libs/object"
)

type (
	AccessTokenInterface interface {
		GetToken() string
		Refresh() AccessTokenInterface
		ApplyToRequest(request contract.RequestInterface, requestOptions object.HashMap) contract.RequestInterface
	}
)