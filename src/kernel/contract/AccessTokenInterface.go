package contract

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/go-wechat/src/kernel/http/contract"
)

type AccessTokenInterface interface {
	GetToken() string
	Refresh() AccessTokenInterface
	ApplyToRequest(request contract.RequestInterface, requestOptions object.HashMap) contract.RequestInterface

}