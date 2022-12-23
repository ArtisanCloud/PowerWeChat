package kernel

import (
	"bytes"
	"encoding/json"
	"fmt"
	contract "github.com/ArtisanCloud/PowerLibs/v3/http/contract"
	"github.com/ArtisanCloud/PowerLibs/v3/http/helper"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/support"
	"io/ioutil"
	"log"
	http "net/http"
	"os"
	"path/filepath"
	"strconv"
)

type BaseClient struct {
	HttpHelper *helper.RequestHelper

	BaseURI string

	*support.ResponseCastable

	Signer *support.SHA256WithRSASigner

	App   *ApplicationInterface
	Token *AccessToken

	GetMiddlewareOfAccessToken        contract.RequestMiddleware
	GetMiddlewareOfLog                func(logger *log.Logger) contract.RequestMiddleware
	GetMiddlewareOfRefreshAccessToken contract.RequestMiddleware
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

	config := (*app).GetConfig()
	baseURI := config.GetString("http.base_uri", "/")

	if token == nil {
		token = (*app).GetAccessToken()
	}
	h, err := helper.NewRequestHelper(&helper.Config{
		BaseUrl: baseURI,
	})

	if err != nil {
		return nil, err
	}
	client := &BaseClient{
		HttpHelper: h,
		App:        app,
		Token:      token,
	}
	// to be setup middleware here
	client.OverrideGetMiddlewares()
	client.RegisterHttpMiddlewares()

	mchID := config.GetString("mch_id", "")
	serialNO := config.GetString("serial_no", "")
	keyPath := config.GetString("key_path", "")
	if mchID != "" && serialNO != "" && keyPath != "" {
		client.Signer = &support.SHA256WithRSASigner{
			MchID:               mchID,
			CertificateSerialNo: serialNO,
			PrivateKeyPath:      keyPath,
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
) (*http.Response, error) {

	// http client request
	df := client.HttpHelper.Df().Uri(url).Method(method)

	// 检查是否需要有请求参数配置
	if options != nil {
		// set query key values
		if (*options)["query"] != nil {
			queries := (*options)["query"].(*object.StringMap)
			if queries != nil {
				for k, v := range *queries {
					df.Query(k, v)
				}
			}
		}

		config := (*client.App).GetConfig()
		// 微信如果需要传debug模式
		debug := config.GetBool("debug", false)
		if debug {
			df.Query("debug", "1")
		}

		// set body json
		if (*options)["form_params"] != nil {
			df.Json((*options)["form_params"])
		}
	}

	response, err := df.Request()
	if err != nil {
		return response, err
	}

	// decode response body to outBody
	if outBody != nil {
		err = client.HttpHelper.ParseResponseBodyContent(response, outBody)
		if err != nil {
			return nil, err
		}
	}

	// decode response header to outHeader
	if outHeader != nil {
		strHeader, err := object.JsonEncode(response.Header)
		if err != nil {
			return nil, err
		}
		err = object.JsonDecode([]byte(strHeader), outHeader)
		if err != nil {
			return nil, err
		}
	}

	return response, err

	//if returnRaw {
	//	return response, err
	//} else {
	//	// tbf
	//	returnResponse, err := client.CastResponseToType(response, response2.TYPE_MAP)
	//	return returnResponse, err
	//}
}

func (client *BaseClient) RequestRaw(url string, method string, options *object.HashMap, outHeader interface{}, outBody interface{}) (*http.Response, error) {
	return client.Request(url, method, options, true, outHeader, outBody)
}

func (client *BaseClient) RegisterHttpMiddlewares() {

	// access token
	accessTokenMiddleware := client.GetMiddlewareOfAccessToken
	// log
	logMiddleware := client.GetMiddlewareOfLog

	// check invalid access token
	checkAccessTokenMiddleware := client.GetMiddlewareOfRefreshAccessToken

	client.HttpHelper.WithMiddleware(
		accessTokenMiddleware,
		logMiddleware(log.Default()),
		checkAccessTokenMiddleware,
	)

}

// ----------------------------------------------------------------------

func (client *BaseClient) OverrideGetMiddlewares() {
	client.OverrideGetMiddlewareOfAccessToken()
	client.OverrideGetMiddlewareOfLog()
	client.OverrideGetMiddlewareOfRefreshAccessToken()
}

func (client *BaseClient) OverrideGetMiddlewareOfAccessToken() {
	client.GetMiddlewareOfAccessToken = func(handle contract.RequestHandle) contract.RequestHandle {
		return func(request *http.Request) (response *http.Response, err error) {
			// 前置中间件
			//fmt.Println("获取access token, 在请求前执行")

			accessToken := (*client.App).GetAccessToken()

			if accessToken != nil {
				config := (*client.App).GetContainer().Config
				_, err = accessToken.ApplyToRequest(request, config)
			}

			if err != nil {
				return nil, err
			}

			response, err = handle(request)
			// handle 执行之后就可以操作 response 和 err

			// 后置中间件
			//fmt.Println("获取access token, 在请求后执行")
			return
		}
	}
}

func (client *BaseClient) OverrideGetMiddlewareOfLog() {
	client.GetMiddlewareOfLog = func(logger *log.Logger) contract.RequestMiddleware {
		return contract.RequestMiddleware(func(handle contract.RequestHandle) contract.RequestHandle {
			return func(request *http.Request) (response *http.Response, err error) {
				// 前置中间件
				//logger.Println("这里是前置中间件log, 在请求前执行")

				response, err = handle(request)
				// handle 执行之后就可以操作 response 和 err

				// 后置中间件
				//logger.Println("这里是后置置中间件log, 在请求后执行")
				return
			}
		})
	}
}

func (client *BaseClient) OverrideGetMiddlewareOfRefreshAccessToken() {
	client.GetMiddlewareOfRefreshAccessToken = func(handle contract.RequestHandle) contract.RequestHandle {
		return func(request *http.Request) (response *http.Response, err error) {
			// 前置中间件
			//fmt.Println("检查微信返回错误，token是否失效，执行前访问")

			response, err = handle(request)
			// handle 执行之后就可以操作 response 和 err

			err = client.CheckTokenNeedRefresh(response)
			if err != nil {
				return nil, err
			}

			//client.HttpHelper.Df()

			// 后置中间件
			//fmt.Println("检查微信返回错误，token是否失效，, 在请求后执行")
			return
		}
	}
}

func (client *BaseClient) CheckTokenNeedRefresh(rs *http.Response) error {
	data, err := ioutil.ReadAll(rs.Body)
	rs.Body = ioutil.NopCloser(bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	mapResponse := &object.HashMap{}
	err = json.Unmarshal(data, mapResponse)
	if err != nil {
		return err
	}

	errCode := 0
	if (*mapResponse)["errcode"] != nil {
		switch (*mapResponse)["errcode"].(type) {
		case float64:
			errCode = int((*mapResponse)["errcode"].(float64))
		case int:
			errCode = (*mapResponse)["errcode"].(int)
		case string:
			errCode, err = strconv.Atoi((*mapResponse)["errcode"].(string))
		default:

		}

		conditions := &object.HashMap{
			"code": errCode,
		}
		if client.RetryDecider(conditions) {
			client.Token.Refresh()
		}
	}

	return nil
}

func (client *BaseClient) RetryDecider(conditions *object.HashMap) bool {
	code := (*conditions)["code"].(int)
	if code == 40001 || code == 40014 || code == 42001 {
		return true
	}
	return false
}
