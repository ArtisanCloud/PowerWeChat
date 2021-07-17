package kernel

import (
	"github.com/ArtisanCloud/go-libs/http/request"
	"github.com/ArtisanCloud/go-libs/http/response"
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/go-libs/str"
	"github.com/ArtisanCloud/power-wechat/src/kernel/support"
	http2 "net/http"
)

type BaseClient struct {
	*request.HttpRequest
	*response.HttpResponse

	*support.ResponseCastable

	App *ApplicationPaymentInterface
}

func NewBaseClient(app *ApplicationPaymentInterface) *BaseClient {
	config := (*app).GetContainer().GetConfig()

	client := &BaseClient{
		HttpRequest: request.NewHttpRequest(config),
		App:         app,
	}
	return client

}

func (client *BaseClient) prepends() *object.StringMap {
	return &object.StringMap{}
}

func (client *BaseClient) Request(endpoint string, params *object.StringMap, method string, options *object.HashMap,
	returnRaw bool, outHeader interface{}, outBody interface{},
) interface{} {

	config:= (*client.App).GetConfig()
	base := &object.StringMap{
		"mch_id":     config.GetString("mch_id", ""),
		"nonce_str":  str.UniqueID(""),
		"sub_mch_id": config.GetString("sub_mch_id", ""),
		"sub_appid":  config.GetString("sub_appid", ""),
	}

	params = object.MergeStringMap(base, params, client.prepends())
	params = object.FilterEmptyStringMap(params)

	secretKey, _ := (*client.App).GetKey(endpoint)
	signType := "MD5"
	if (*params)["sign_type"] != "" {
		signType = (*params)["sign_type"]
	}
	encryptMethod := support.GetEncryptMethod(signType, secretKey)
	(*params)["sign"] = support.GenerateSign(params, secretKey, encryptMethod)

	options = object.MergeHashMap(&object.HashMap{
		//"body": object.StringMap2Xml(params),
		"body": params,
	}, options)

	// to be setup middleware here
	//client.PushMiddleware(client.logMiddleware(), "access_token")

	// http client request
	returnResponse := client.PerformRequest(endpoint, method, options, false, outHeader, outBody)

	if returnRaw {
		return returnResponse
	} else {
		responseType := config.GetString("response_type", "array")
		var rs http2.Response = http2.Response{
			StatusCode: 200,
			Header:     nil,
		}
		rs.Body = returnResponse.GetBody()
		result, _ := client.CastResponseToType(&rs, responseType)
		return result
	}

}

func (client *BaseClient) RequestRaw(url string, params *object.StringMap, method string, options *object.HashMap, outHeader interface{}, outBody interface{}) interface{} {
	return client.Request(url, params, method, options, true, outHeader, outBody)
}

func (client *BaseClient) RequestArray(url string, method string, options *object.HashMap, outHeader interface{}, outBody interface{}) *object.HashMap {
	returnResponse := client.RequestRaw(url, nil, method, options, outHeader, outBody)
	result, _ := client.CastResponseToType(returnResponse.(*http2.Response), "array")

	return result.(*object.HashMap)
}

func (client *BaseClient) SafeRequest(url string, query interface{}, outHeader interface{}, outBody interface{}) interface{} {
	return client.Request(
		url,
		nil,
		"GET",
		&object.HashMap{
			"query": query,
		},
		false,
		outHeader,
		outBody,
	)
}
func (client *BaseClient) Wrap(endpoint string) string {
	if (*client.App).InSandbox() {
		return "sandboxnew/" + endpoint
	} else {
		return endpoint
	}
}

// ----------------------------------------------------------------------
type MiddlewareLogMiddleware struct {
	*BaseClient
}

func (client *BaseClient) logMiddleware() interface{} {
	return &MiddlewareLogMiddleware{
		client,
	}
}
