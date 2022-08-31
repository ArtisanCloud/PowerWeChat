package kernel

import (
	"fmt"
	"github.com/ArtisanCloud/PowerLibs/v2/http/contract"
	"github.com/ArtisanCloud/PowerLibs/v2/http/request"
	"github.com/ArtisanCloud/PowerLibs/v2/http/response"
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	response2 "github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/support"
	http2 "net/http"
	"os"
	"path/filepath"
	"time"
)

type BaseClient struct {
	*request.HttpRequest
	*response.HttpResponse

	*support.ResponseCastable

	ExternalRequest *http2.Request

	Signer *support.SHA256WithRSASigner

	App   *ApplicationInterface
	Token *AccessToken
}

type UploadForm struct {
	FileName string
	Contents []*UploadContent
}
type UploadContent struct {
	Name  string
	Value interface{}
}

func NewBaseClient(app *ApplicationInterface, token *AccessToken) (*BaseClient, error) {
	config := (*app).GetContainer().GetConfig()

	if token == nil {
		token = (*app).GetAccessToken()
	}
	httpRequest, err := request.NewHttpRequest(config)
	if err != nil {
		return nil, err
	}
	client := &BaseClient{
		HttpRequest: httpRequest,
		App:         app,
		Token:       token,
	}

	if (*config)["mch_id"] != nil && (*config)["serial_no"] != nil && (*config)["key_path"] != nil {
		client.Signer = &support.SHA256WithRSASigner{
			MchID:               (*config)["mch_id"].(string),
			CertificateSerialNo: (*config)["serial_no"].(string),
			PrivateKeyPath:      (*config)["key_path"].(string),
		}
	}

	return client, nil

}

func (client *BaseClient) HttpGet(url string, query *object.StringMap, outHeader interface{}, outBody interface{}) (interface{}, error) {
	return client.Request(
		url,
		"GET",
		&object.HashMap{
			"query": query,
		},
		false,
		outHeader,
		outBody,
	)
}

func (client *BaseClient) HttpPost(url string, data interface{}, outHeader interface{}, outBody interface{}) (interface{}, error) {
	return client.Request(
		url,
		"POST",
		&object.HashMap{
			"form_params": data,
		},
		false,
		outHeader,
		outBody,
	)
}

func (client *BaseClient) HttpPostJson(url string, data interface{}, query *object.StringMap, outHeader interface{}, outBody interface{}) (interface{}, error) {
	return client.Request(
		url,
		"POST",
		&object.HashMap{
			"query":       query,
			"form_params": data,
		},
		false,
		outHeader,
		outBody,
	)
}

func (client *BaseClient) HttpUpload(url string, files *object.HashMap, form *UploadForm, query interface{}, outHeader interface{}, outBody interface{}) (interface{}, error) {

	multipart := []*object.HashMap{}
	headers := &object.HashMap{}

	// 如果设置了filename，则初始化一个header，在每一个multipart里注入
	if form != nil {
		if form.FileName != "" {
			fileName := form.FileName
			(*headers)["Content-Disposition"] = fmt.Sprintf("form-data; name=\"media\"; filename=\"%s\"", fileName)
		}
	}

	// 遍历文件目录
	if files != nil {
		for name, path := range *files {

			_, err := os.Open(path.(string))
			if err != nil {
				return nil, err
			}
			if (*headers)["filename"] == nil {
				(*headers)["filename"] = filepath.Base(path.(string))
			}

			multipart = append(multipart, &object.HashMap{
				"name":    name,
				"value":   path,
				"headers": headers,
			})
		}
	}

	// 遍历表单的数据
	if form != nil {
		for _, content := range form.Contents {
			part := &object.HashMap{
				"name": content.Name,
				//"value": object.EncodeToBytes(content.Value),
				"value": content.Value,
			}
			multipart = append(multipart, part)
		}
	}

	return client.Request(url, "POST", &object.HashMap{
		"query":           query,
		"multipart":       multipart,
		"connect_timeout": 30,
		"timeout":         30,
		"read_timeout":    30,
	}, false, nil, outBody)
}

func (client *BaseClient) Request(url string, method string, options *object.HashMap,
	returnRaw bool, outHeader interface{}, outBody interface{},
) (interface{}, error) {

	// to be setup middleware here
	if client.Middlewares == nil {
		client.registerHttpMiddlewares()
	}
	// http client request
	response, err := client.PerformRequest(url, method, options, returnRaw, outHeader, outBody)

	if err != nil {
		return nil, err
	}

	if returnRaw {
		return response, err
	} else {
		// tbf
		var rs http2.Response = http2.Response{
			StatusCode: response.GetStatusCode(),
			Header:     response.GetHeader(),
			Body:       response.GetBody(),
		}
		returnResponse, err := client.CastResponseToType(&rs, response2.TYPE_MAP)
		return returnResponse, err
	}
}

func (client *BaseClient) RequestRaw(url string, method string, options *object.HashMap, outHeader interface{}, outBody interface{}) (interface{}, error) {
	return client.Request(url, method, options, true, outHeader, outBody)
}

func (client *BaseClient) registerHttpMiddlewares() {

	client.Middlewares = []interface{}{}

	// retry
	retryMiddleware := client.retryMiddleware()
	client.PushMiddleware(retryMiddleware, retryMiddleware.Name)
	// access token
	accessTokenMiddleware := client.accessTokenMiddleware()
	client.PushMiddleware(accessTokenMiddleware, "access_token")
	// log
	//logMiddleware:=client.logMiddleware()
	//client.PushMiddleware(logMiddleware, logMiddleware.Name)

}

// ----------------------------------------------------------------------
type Middleware struct {
	contract.MiddlewareInterface
	*BaseClient
	Name string
}

func (d *Middleware) GetName() string {
	return d.Name
}

func (d *Middleware) SetName(name string) {
	d.Name = name
}

func (d *Middleware) Retries() int {
	config := (*d.BaseClient.App).GetConfig()
	return config.GetInt("http.max_retries", 1)
}

func (d *Middleware) Delay() time.Duration {
	config := (*d.BaseClient.App).GetConfig()
	second := time.Duration(config.GetInt("http.retry_delay", 500))
	return second * time.Second
}

func (d *Middleware) RetryDecider(conditions *object.HashMap) bool {
	return false
}

type MiddlewareAccessToken struct {
	*Middleware
}
type MiddlewareLogMiddleware struct {
	*Middleware
}
type MiddlewareRetry struct {
	*Middleware
}

// --- MiddlewareAccessToken ---
func (d *MiddlewareAccessToken) ModifyRequest(req *http2.Request) (err error) {
	accessToken := (*d.App).GetAccessToken()

	if accessToken != nil {
		config := (*d.App).GetContainer().Config
		_, err = accessToken.ApplyToRequest(req, config)
	}

	return err
}

// --- MiddlewareLogMiddleware ---
func (d *MiddlewareLogMiddleware) ModifyRequest(req *http2.Request) error {
	fmt.Println("logMiddleware")
	return nil
}

// --- MiddlewareRetry ---
func (d *MiddlewareRetry) ModifyRequest(req *http2.Request) error {
	return nil
}
func (d *MiddlewareRetry) RetryDecider(conditions *object.HashMap) bool {
	code := (*conditions)["code"].(int)
	if code == 40001 || code == 40014 || code == 42001 {
		d.BaseClient.Token.Refresh()

		return true
	}
	return false
}

// ---
func (client *BaseClient) accessTokenMiddleware() *MiddlewareAccessToken {
	return &MiddlewareAccessToken{
		&Middleware{
			BaseClient: client,
			Name:       "access_token",
		},
	}
}
func (client *BaseClient) logMiddleware() *MiddlewareLogMiddleware {
	return &MiddlewareLogMiddleware{
		&Middleware{
			BaseClient: client,
			Name:       "log",
		},
	}
}
func (client *BaseClient) retryMiddleware() *MiddlewareRetry {
	return &MiddlewareRetry{
		&Middleware{
			BaseClient: client,
			Name:       "retry",
		},
	}
}
