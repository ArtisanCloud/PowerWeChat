package http

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/go-libs/str"
	"github.com/ArtisanCloud/go-wechat/src/kernel/http/Contract"
)


func (client *BaseClient) performRequest(url string, method string, options object.HashMap) Contract.ResponseContract {
	// change method string format
	method = str.Lower(method)

	// merge options with default options

	response := client.GetHttpClient().Request(method, url, options)
	return response
}

func (client *BaseClient) SetHttpClient(httpClient *Contract.ClientInterface) *BaseClient {
	client.httpClient = httpClient
	return client
}

func (client *BaseClient) GetHttpClient() Contract.ClientInterface {
	return nil
}
