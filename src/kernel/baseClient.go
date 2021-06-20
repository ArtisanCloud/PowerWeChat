package kernel

import (
	"fmt"
	"github.com/ArtisanCloud/go-libs/http"
	"github.com/ArtisanCloud/go-libs/object"
	http2 "net/http"
)

type BaseClient struct {
	*http.HttpRequest
	*http.HttpResponse

	App   *ApplicationInterface
	Token *AccessToken
}

func NewBaseClient(app *ApplicationInterface, token *AccessToken) *BaseClient {
	config := (*app).GetContainer().GetConfig()

	if token == nil {
		token = (*app).GetAccessToken()
	}

	client := &BaseClient{
		HttpRequest: http.NewHttpRequest(config),
		App:         app,
		Token:       token,
	}
	return client

}

func (client *BaseClient) HttpGet(url string, query object.StringMap, outResponse interface{}) interface{} {
	return client.Request(
		url,
		"GET",
		&object.HashMap{
			"query": query,
		},
		false,
		outResponse,
	)
}

func (client *BaseClient) HttpPost(url string, data interface{}, outResponse interface{}) interface{} {
	return client.Request(
		url,
		"POST",
		&object.HashMap{
			"form_params": data,
		},
		false,
		outResponse,
	)
}

func (client *BaseClient) HttpPostJson(url string, data interface{}, query interface{}, outResponse interface{}) interface{} {
	return client.Request(
		url,
		"POST",
		&object.HashMap{
			"query":       query,
			"form_params": data,
		},
		false,
		outResponse,
	)
}

func (client *BaseClient) Request(url string, method string, options *object.HashMap, returnRaw bool, outResponse interface{}) interface{} {

	// to be setup middleware here
	if client.Middlewares == nil {
		client.registerHttpMiddlewares()
	}
	// http client request
	response := client.PerformRequest(url, method, options, outResponse)

	if returnRaw {
		return response
	} else {
		config := *(*client.App).GetContainer().Config
		client.CastResponseToType(response, config["response_type"])
	}
	return response
}

func (client *BaseClient) registerHttpMiddlewares() {

	client.Middlewares = []interface{}{}

	// retry
	//client.PushMiddleware(client.retryMiddleware(), "retry")
	// access token
	client.PushMiddleware(client.accessTokenMiddleware(), "access_token")
	// log
	//client.PushMiddleware(client.logMiddleware(), "log")

}

// ----------------------------------------------------------------------
type MiddlewareAccessToken struct {
	*BaseClient
}
type MiddlewareLogMiddleware struct {
	*BaseClient
}
type MiddlewareRetry struct {
	*BaseClient
}

func (d *MiddlewareAccessToken) ModifyRequest(req *http2.Request) (err error) {
	accessToken := (*d.App).GetAccessToken()

	if accessToken != nil {
		config := (*d.App).GetContainer().Config
		_, err = accessToken.ApplyToRequest(req, config)
	}

	return err
}
func (d *MiddlewareLogMiddleware) ModifyRequest(req *http2.Request) error {
	fmt.Println("logMiddleware")
	return nil
}
func (d *MiddlewareRetry) ModifyRequest(req *http2.Request) error {
	fmt.Println("retryMiddleware")
	return nil
}

func (client *BaseClient) accessTokenMiddleware() interface{} {
	return &MiddlewareAccessToken{
		client,
	}
}
func (client *BaseClient) logMiddleware() interface{} {
	return &MiddlewareLogMiddleware{
		client,
	}
}
func (client *BaseClient) retryMiddleware() interface{} {
	return &MiddlewareRetry{
		client,
	}
}
