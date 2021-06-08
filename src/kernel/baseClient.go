package kernel

import (
	"github.com/ArtisanCloud/go-libs/http"
	"github.com/ArtisanCloud/go-libs/object"
)

type BaseClient struct {
	*http.HttpRequest
	*http.HttpResponse

	App *ServiceContainer

	AccessToken *AccessToken
}

func NewBaseClient(app *ServiceContainer, token *AccessToken) *BaseClient {
	client := &BaseClient{
		HttpRequest: http.NewHttpRequest(app.GetConfig()),
		App:         app,
		AccessToken: token,
	}
	return client

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
	if client.Middlewares != nil {
		client.registerHttpMiddlewares()
	}
	//
	response := client.PerformRequest(url, method, options)

	if returnRaw {
		return response
	} else {
		client.CastResponseToType(response, (*client.App.Config)["response_type"].(string))
	}
	return response
}

func (client *BaseClient) registerHttpMiddlewares() {
	// retry
	//client.pushMiddleware(client.retryMiddleware(), "retry")
	//// access token
	//client.pushMiddleware(client.accessTokenMiddleware(), "access_token")
	//// log
	//client.pushMiddleware(client.logMiddleware(), "log")
}

func (client *BaseClient) accessTokenMiddleware() func() {
	return func() {

	}
}
