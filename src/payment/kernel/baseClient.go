package kernel

import (
	"encoding/json"
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

func (client *BaseClient) prepends() *object.HashMap {
	return &object.HashMap{}
}

func (client *BaseClient) Request(endpoint string, params *object.StringMap, method string, options *object.HashMap,
	returnRaw bool, outHeader interface{}, outBody interface{},
) (interface{}, error) {

	config := (*client.App).GetConfig()
	base := &object.HashMap{
		"mch_id":     config.GetString("mch_id", ""),
		"nonce_str":  str.UniqueID(""),
		"sub_mch_id": config.GetString("sub_mch_id", ""),
		"sub_appid":  config.GetString("sub_appid", ""),
	}

	options = object.MergeHashMap(base, client.prepends(), options)
	options = object.FilterEmptyHashMap(options)

	signer := &support.SHA256WithRSASigner{
		MchID:               config.GetString("mch_id", ""),
		CertificateSerialNo: "2655A2CD634B06C2A86B28780228A997D047B01C",
		PrivateKeyPath:      config.GetString("key_path", ""),
	}

	bufferBody, err := json.Marshal(options)
	signBody := string(bufferBody)
	if err != nil {
		return nil, err
	}
	authorization, err := support.GenerateSign(signer, support.GenerateSigner{
		Method:       "POST",
		CanonicalURL: "/v3/pay/transactions/jsapi",
		SignBody:     signBody,
	})

	options = object.MergeHashMap(&object.HashMap{
		"headers": &object.HashMap{
			"Authorization": authorization,
		},
		"body": signBody,
	}, options)

	// to be setup middleware here
	//client.PushMiddleware(client.logMiddleware(), "access_token")

	// http client request
	returnResponse := client.PerformRequest(endpoint, method, options, returnRaw, outHeader, outBody)

	if returnRaw {
		return returnResponse, nil
	} else {
		responseType := config.GetString("response_type", "array")
		var rs http2.Response = http2.Response{
			StatusCode: 200,
			Header:     nil,
		}
		rs.Body = returnResponse.GetBody()
		result, _ := client.CastResponseToType(&rs, responseType)
		return result, nil
	}

}

func (client *BaseClient) RequestRaw(url string, params *object.StringMap, method string, options *object.HashMap, outHeader interface{}, outBody interface{}) (interface{}, error) {
	return client.Request(url, params, method, options, true, outHeader, outBody)
}

func (client *BaseClient) RequestArray(url string, method string, options *object.HashMap, outHeader interface{}, outBody interface{}) (*object.HashMap, error) {
	returnResponse, err := client.RequestRaw(url, nil, method, options, outHeader, outBody)
	if err != nil {
		return nil, err
	}
	result, err := client.CastResponseToType(returnResponse.(*http2.Response), "array")

	return result.(*object.HashMap), err
}

func (client *BaseClient) SafeRequest(url string, query interface{}, outHeader interface{}, outBody interface{}) (interface{}, error) {
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
