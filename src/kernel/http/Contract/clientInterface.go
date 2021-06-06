package Contract

import (
	"github.com/ArtisanCloud/go-libs/object"
)

type ClientInterface interface {
	Send(request RequestInterface, options object.HashMap) ResponseContract
	SendAsync(request RequestInterface, options object.HashMap) PromiseInterface

	Request(method string, uri string, options object.HashMap) ResponseContract
	RequestAsync(method string, uri string, options object.HashMap) PromiseInterface
}
