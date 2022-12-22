package kernel

import (
	"bytes"
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
)

type BaseClient struct {
	Helper *helper.RequestHelper

	Middlewares []contract.RequestMiddleware

	*support.ResponseCastable

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
		Helper: h,
		App:    app,
		Token:  token,
	}

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

	// to be setup middleware here
	if len(client.Middlewares) == 0 {
		client.registerHttpMiddlewares()
	}
	// http client request
	response, err := client.Helper.Df().Url(url).Method(method).Json(options).Request()

	// decode response body to outBody
	bodyData, _ := ioutil.ReadAll(response.Body)
	response.Body = ioutil.NopCloser(bytes.NewBuffer(bodyData))
	err = object.JsonDecode(bodyData, outBody)

	// decode response header to outHeader
	//headerData, _ := ioutil.ReadAll(response.Header)
	//response.Header = ioutil.NopCloser(bytes.NewBuffer(headerData))
	//err = object.JsonDecode(headerData, outHeader)

	if err != nil {
		return nil, err
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

func (client *BaseClient) registerHttpMiddlewares() {

	client.Middlewares = []contract.RequestMiddleware{}

	// access token
	accessMiddleware := func(handle contract.RequestHandle) contract.RequestHandle {
		return func(request *http.Request) (response *http.Response, err error) {
			// 前置中间件
			fmt.Println("这里是前置中间件2, 在请求前执行")

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
			fmt.Println("这里是后置置中间件2, 在请求后执行")
			return
		}
	}
	// log
	logMiddleware := func(logger *log.Logger) contract.RequestMiddleware {
		return contract.RequestMiddleware(func(handle contract.RequestHandle) contract.RequestHandle {
			return func(request *http.Request) (response *http.Response, err error) {
				// 前置中间件
				logger.Println("这里是前置中间件log, 在请求前执行")

				response, err = handle(request)
				// handle 执行之后就可以操作 response 和 err

				// 后置中间件
				logger.Println("这里是后置置中间件log, 在请求后执行")
				return
			}
		})
	}

	client.Helper.WithMiddleware(accessMiddleware, logMiddleware(log.Default()))

}

// ----------------------------------------------------------------------

//func (client *BaseClient) CheckTokenNeedRefresh(rs contract.ResponseInterface) error {
//	data, err := rs.GetBodyData()
//	if err != nil {
//		return err
//	}
//	mapResponse := &object.HashMap{}
//	err = json.Unmarshal(data, mapResponse)
//	if err != nil {
//		return err
//	}
//
//	errCode := 0
//	if (*mapResponse)["errcode"] != nil {
//		switch (*mapResponse)["errcode"].(type) {
//		case float64:
//			errCode = int((*mapResponse)["errcode"].(float64))
//		case int:
//			errCode = (*mapResponse)["errcode"].(int)
//		case string:
//			errCode, err = strconv.Atoi((*mapResponse)["errcode"].(string))
//		default:
//
//		}
//
//		conditions := &object.HashMap{
//			"code": errCode,
//		}
//		if client.retryMiddleware().RetryDecider(conditions) {
//			client.Token.Refresh()
//		}
//	}
//
//	return nil
//}
