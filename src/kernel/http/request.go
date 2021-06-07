package http

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/go-libs/str"
	"github.com/ArtisanCloud/go-wechat/src/kernel/http/contract"
)

var _defaults *object.HashMap

func SetDefaultOptions(defaults *object.HashMap) {
	_defaults = defaults
}

func GetDefaultOptions() *object.HashMap {
	return _defaults
}

func (client *BaseClient) performRequest(url string, method string, options object.HashMap) contract.ResponseContract {
	// change method string format
	method = str.Lower(method)

	// merge options with default options

	response := client.GetHttpClient().Request(method, url, options)
	return response
}

func (client *BaseClient) SetHttpClient(httpClient *contract.ClientInterface) *BaseClient {
	client.httpClient = httpClient
	return client
}

func (client *BaseClient) GetHttpClient() contract.ClientInterface {
	return nil
}
