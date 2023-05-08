package kernel

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	contract "github.com/ArtisanCloud/PowerLibs/v3/http/contract"
	"github.com/ArtisanCloud/PowerLibs/v3/http/helper"
	contract2 "github.com/ArtisanCloud/PowerLibs/v3/logger/contract"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerLibs/v3/os"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/support"
	"io"
	"io/ioutil"
	http "net/http"
	"strconv"
)

type BaseClient struct {
	HttpHelper *helper.RequestHelper

	BaseURI  string
	QueryRaw bool

	*support.ResponseCastable

	Signer *support.SHA256WithRSASigner

	App   *ApplicationInterface
	Token *AccessToken

	GetMiddlewareOfAccessToken        contract.RequestMiddleware
	GetMiddlewareOfLog                func(logger contract2.LoggerInterface) contract.RequestMiddleware
	GetMiddlewareOfRefreshAccessToken func(retry int) contract.RequestMiddleware
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

func (client *BaseClient) HttpGet(ctx context.Context, url string, query *object.StringMap, outHeader interface{}, outBody interface{}) (interface{}, error) {
	return client.Request(
		ctx,
		url,
		http.MethodGet,
		&object.HashMap{
			"query": query,
		},
		false,
		outHeader,
		outBody,
	)
}

func (client *BaseClient) HttpPost(ctx context.Context, url string, data interface{}, outHeader interface{}, outBody interface{}) (interface{}, error) {
	return client.Request(
		ctx,
		url,
		http.MethodPost,
		&object.HashMap{
			"form_params": data,
		},
		false,
		outHeader,
		outBody,
	)
}

func (client *BaseClient) HttpPostJson(ctx context.Context, url string, data interface{}, query *object.StringMap, outHeader interface{}, outBody interface{}) (interface{}, error) {
	return client.Request(
		ctx,
		url,
		http.MethodPost,
		&object.HashMap{
			"query":       query,
			"form_params": data,
		},
		false,
		outHeader,
		outBody,
	)
}

func (client *BaseClient) HttpUpload(ctx context.Context, url string, files *object.HashMap, form *UploadForm, query interface{}, outHeader interface{}, outBody interface{}) (interface{}, error) {

	// http client request
	df := client.HttpHelper.Df().WithContext(ctx).Uri(url).Method(http.MethodPost)

	var mems map[string]io.Reader
	// 遍历表单的数据
	if form != nil {
		mems = make(map[string]io.Reader)
		for _, content := range form.Contents {
			value, err := os.ConvertFileHandleToReader(content.Value)
			if err != nil {
				return nil, err
			}
			mems[content.Name] = value
		}
	}

	df.Multipart(func(multipart contract.MultipartDfInterface) {
		// 遍历文件目录
		if files != nil {
			for name, path := range *files {
				multipart.FileByPath(name, path.(string))
			}
		}

		for k, v := range mems {
			multipart.FileMem(form.FileName, k, v)
		}

	})

	// set query params
	if query != nil {
		queries := query.(*object.StringMap)
		if queries != nil {
			for k, v := range *queries {
				df.Query(k, v)
			}
		}
	}

	// set debug mode
	config := (*client.App).GetConfig()
	// 微信如果需要传debug模式
	debug := config.GetBool("debug", false)
	if debug {
		df.Query("debug", "1")
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

	return response, err

}

func (client *BaseClient) Request(ctx context.Context, url string, method string, options *object.HashMap,
	returnRaw bool, outHeader interface{}, outBody interface{},
) (*http.Response, error) {

	// http client request
	client.QueryRaw = returnRaw
	df := client.HttpHelper.Df().WithContext(ctx).Uri(url).Method(method)

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

}

func (client *BaseClient) RequestRaw(ctx context.Context, url string, method string, options *object.HashMap, outHeader interface{}, outBody interface{}) (*http.Response, error) {
	return client.Request(ctx, url, method, options, true, outHeader, outBody)
}

func (client *BaseClient) RegisterHttpMiddlewares() {

	// access token
	accessTokenMiddleware := client.GetMiddlewareOfAccessToken
	// log
	logMiddleware := client.GetMiddlewareOfLog

	// check invalid access token
	checkAccessTokenMiddleware := client.GetMiddlewareOfRefreshAccessToken

	config := (*client.App).GetConfig()
	logger := (*client.App).GetComponent("Logger").(contract2.LoggerInterface)
	client.HttpHelper.WithMiddleware(
		accessTokenMiddleware,
		logMiddleware(logger),
		checkAccessTokenMiddleware(3),
		helper.HttpDebugMiddleware(config.GetBool("http_debug", false)),
	)
}

// ----------------------------------------------------------------------

func (client *BaseClient) OverrideGetMiddlewares() {
	client.OverrideGetMiddlewareOfLog()
	client.OverrideGetMiddlewareOfAccessToken()
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
				request, err = accessToken.ApplyToRequest(request, config)
			}

			if err != nil {
				return nil, err
			}

			response, err = handle(request)
			if err != nil {
				return response, err
			}

			// handle 执行之后就可以操作 response 和 err
			// 后置中间件
			//fmt.Println("获取access token, 在请求后执行")
			return
		}
	}
}

func (client *BaseClient) OverrideGetMiddlewareOfLog() {
	client.GetMiddlewareOfLog = func(logger contract2.LoggerInterface) contract.RequestMiddleware {
		return contract.RequestMiddleware(func(handle contract.RequestHandle) contract.RequestHandle {
			return func(request *http.Request) (response *http.Response, err error) {
				// 前置中间件
				//logger.Println("这里是前置中间件log, 在请求前执行")

				response, err = handle(request)
				if err != nil {
					return response, err
				}

				//// 后置中间件
				////logger.Println("这里是后置置中间件log, 在请求后执行")
				return
			}
		})
	}
}

func (client *BaseClient) OverrideGetMiddlewareOfRefreshAccessToken() {
	client.GetMiddlewareOfRefreshAccessToken = func(retry int) contract.RequestMiddleware {
		return contract.RequestMiddleware(func(handle contract.RequestHandle) contract.RequestHandle {
			return func(request *http.Request) (response *http.Response, err error) {
				// 前置中间件
				//fmt.Println("检查微信返回错误，token是否失效，执行前访问")

				response, err = handle(request)
				// handle 执行之后就可以操作 response 和 err

				if err != nil {
					return response, err
				}

				if response.StatusCode != http.StatusOK {
					return response, errors.New(fmt.Sprintf("http response code:%d", response.StatusCode))
				}

				if !client.QueryRaw {
					rs, err := client.CheckTokenNeedRefresh(request, response, retry)
					if err != nil {
						return rs, err
					} else if rs != nil {
						return rs, nil
					}
				}

				// 后置中间件
				//fmt.Println("检查微信返回错误，token是否失效，, 在请求后执行")
				return
			}
		})
	}
}

func (client *BaseClient) CheckTokenNeedRefresh(req *http.Request, rs *http.Response, retry int) (*http.Response, error) {

	mapResponse := &object.HashMap{}
	err := client.HttpHelper.ParseResponseBodyToMap(rs, mapResponse)
	if err != nil {
		return nil, err
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
		if retry > 0 && client.RetryDecider(conditions) {
			client.Token.Refresh()

			// clone 一个request
			token, err := client.Token.GetToken(false)
			q := req.URL.Query()
			q.Set(client.Token.TokenKey, token.AccessToken)
			req.URL.RawQuery = q.Encode()
			req2 := req.Clone(req.Context())
			if req.Body != nil {
				// 缓存请求body
				reqData, err := ioutil.ReadAll(req.Body)
				if err != nil {
					return nil, err
				}

				// 给两个request复制缓存下来的body数据
				req.Body = ioutil.NopCloser(bytes.NewBuffer(reqData))
				req2.Body = ioutil.NopCloser(bytes.NewReader(reqData))
			}

			res2, err := client.HttpHelper.GetClient().DoRequest(req2)
			if err != nil {
				return nil, err
			}

			return res2, err
			//b, err := ioutil.ReadAll(res2.Body)
			//rs.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			//content := string(b)
			//fmt2.Dump(content)
		}
	}

	return rs, nil
}

func (client *BaseClient) RetryDecider(conditions *object.HashMap) bool {
	code := (*conditions)["code"].(int)
	if code == 40001 || code == 40014 || code == 42001 {
		return true
	}
	return false
}
