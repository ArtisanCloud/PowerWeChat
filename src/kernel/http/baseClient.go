package http

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/go-wechat/src/kernel"
	"github.com/ArtisanCloud/go-wechat/src/kernel/http/Contract"
)

type BaseClient struct {
	App *kernel.ServiceContainer

	httpClient *Contract.ClientInterface
}

func (client *BaseClient) HttpGet(url string, query object.StringMap) interface{} {
	return client.Request(
		url,
		"GET",
		object.HashMap{
			"query": query,
		},
		false,
	)
}

func (client *BaseClient) HttpPost(url string, data object.HashMap) interface{} {
	return client.Request(
		url,
		"POST",
		object.HashMap{
			"form_params": data,
		},
		false,
	)
}

func (client *BaseClient) HttpPostJson(url string, data object.HashMap, query object.StringMap) interface{} {
	return client.Request(
		url,
		"POST",
		object.HashMap{
			"query":       query,
			"form_params": data,
		},
		false,
	)
}

func (client *BaseClient) Request(url string, method string, options object.HashMap, returnRaw bool) interface{} {

	// to be setup middleware here

	//
	response := client.performRequest(url, method, options)

	if returnRaw {
		return response
	} else {
		client.castResponseToType(response, client.App.Config[1]["response_type"].(string))
	}
	return nil
}

