package kernel

import (
	"bytes"
	"context"
	"crypto"
	"crypto/sha1"
	"errors"
	"fmt"
	fmt2 "github.com/ArtisanCloud/PowerLibs/v3/fmt"
	"github.com/ArtisanCloud/PowerLibs/v3/http/contract"
	"github.com/ArtisanCloud/PowerLibs/v3/http/helper"
	contract2 "github.com/ArtisanCloud/PowerLibs/v3/logger/contract"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerLibs/v3/security/sign"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	response2 "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/support"
	"io"
	"log"
	http "net/http"
	"os"
)

type BaseClient struct {
	kernel.BaseClient

	RsaOAEP *sign.RSASigner

	App *ApplicationPaymentInterface
}

func NewBaseClient(app *ApplicationPaymentInterface) (*BaseClient, error) {
	config := (*app).GetConfig()
	baseURI := config.GetString("http.base_uri", "/")

	httpRequest, err := helper.NewRequestHelper(&helper.Config{
		BaseUrl: baseURI,
	})
	if err != nil {
		return nil, err
	}

	client := &BaseClient{
		BaseClient: kernel.BaseClient{
			HttpHelper: httpRequest,
			Signer: &support.SHA256WithRSASigner{
				MchID:               config.GetString("mch_id", ""),
				CertificateSerialNo: config.GetString("serial_no", ""),
				PrivateKeyPath:      config.GetString("key_path", ""),
			},
		},
		App: app,
	}

	client.RsaOAEP, err = sign.NewRSASigner(crypto.SHA1)
	if err != nil {
		return nil, err
	}
	RSAPublicKeyPath := config.GetString("certificate_key_path", "")
	if RSAPublicKeyPath != "" {
		client.RsaOAEP.RSAEncryptor.PublicKeyPath = RSAPublicKeyPath
		_, err = client.RsaOAEP.RSAEncryptor.LoadPublicKeyByPath()
		if err != nil {
			return nil, err
		}
	}

	// to be setup middleware here
	client.OverrideGetMiddlewares()
	client.RegisterHttpMiddlewares()

	return client, nil

}

func (client *BaseClient) prepends() *object.HashMap {
	return &object.HashMap{}
}

func (client *BaseClient) PlainRequest(ctx context.Context, endpoint string, params *object.StringMap, method string, options *object.HashMap,
	returnRaw bool, outHeader interface{}, outBody interface{},
) (response *http.Response, err error) {

	//config := (*client.App).GetConfig()
	base := &object.HashMap{}

	// init options
	if options == nil {
		options = &object.HashMap{}
	}

	options = object.MergeHashMap(base, client.prepends(), options)
	options = object.FilterEmptyHashMap(options)

	// check need sign body or not
	signBody := ""
	if "get" != object.Lower(method) {
		signBody, err = object.JsonEncode(options)
		if err != nil {
			return nil, err
		}
	}

	authorization, err := client.Signer.GenerateRequestSign(&support.RequestSignChain{
		Method:       method,
		CanonicalURL: endpoint,
		SignBody:     signBody,
	})

	if err != nil {
		return nil, err
	}

	df := client.HttpHelper.Df().
		WithContext(ctx).
		Uri(endpoint).Method(method)

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
		if signBody != "" {
			r := bytes.NewBufferString(signBody)
			df.Body(r)
		}

		// set header
		df.
			Header("content-type", "application/json").
			Header("Authorization", authorization)

	}

	returnResponse, err := df.Request()
	if err != nil {
		return returnResponse, err
	}

	// decode response body to outBody
	err = client.HttpHelper.ParseResponseBodyContent(returnResponse, outBody)

	return returnResponse, err

}

func (client *BaseClient) RequestV2(ctx context.Context, endpoint string, params *object.HashMap, method string, option *object.HashMap,
	returnRaw bool, outHeader interface{}, outBody interface{},
) (response *http.Response, err error) {

	config := (*client.App).GetConfig()

	base := &object.HashMap{
		// 微信的接口如果传入接口以外的参数，签名会失败所以这里需要区分对待参数
		"mch_id":     config.GetString("mch_id", ""),
		"nonce_str":  object.RandStringBytesMask(32),
		"sub_mch_id": config.GetString("sub_mch_id", ""),
		"sub_appid":  config.GetString("sub_appid", ""),
	}
	params = object.MergeHashMap(params, base)
	if (*params)["mchid"] == nil {
		(*params)["mch_id"] = config.GetString("mch_id", "")
	} else {
		(*params)["mch_id"] = nil
	}
	params = object.FilterEmptyHashMap(params)

	//options, err := client.AuthSignRequestV2(endpoint, method, params, option)
	options, err := client.AuthSignRequestV2(endpoint, method, params, option)
	if err != nil {
		return nil, err
	}

	// http client request
	df := client.HttpHelper.Df().
		WithContext(ctx).
		Uri(endpoint).Method(method)

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
		if (*options)["body"] != nil {
			r := bytes.NewBufferString((*options)["body"].(string))
			df.Body(r)
		}
	}

	returnResponse, err := df.Request()
	if err != nil {
		return returnResponse, err
	}

	return returnResponse, err

}

func (client *BaseClient) Request(ctx context.Context, endpoint string, params *object.StringMap, method string, options *object.HashMap,
	returnRaw bool, outHeader interface{}, outBody interface{},
) (response *http.Response, err error) {

	config := (*client.App).GetConfig()

	// 签名访问的URL，请确保url后面不要跟参数，因为签名的参数，不包含?参数
	// 比如需要在请求的时候，把debug=false，这样url后面就不会多出 "?debug=true"
	options, err = client.AuthSignRequest(ctx, config, endpoint, method, params, options)
	if err != nil {
		return nil, err
	}
	// to be setup middleware here
	//client.PushMiddleware(client.logMiddleware(), "access_token")

	// http client request
	df := client.HttpHelper.Df().
		WithContext(ctx).
		Uri(endpoint).Method(method)

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
		if (*options)["body"] != nil {
			r := bytes.NewBufferString((*options)["body"].(string))
			df.Body(r)
		}

		// set header
		df.
			Header("content-type", "application/json").
			Header("Authorization", (*(*options)["headers"].(*object.HashMap))["Authorization"].(string))

		// 微信支付时候，需要配置平台证书的序列号
		// https://pay.weixin.qq.com/wiki/doc/apiv3/wechatpay/wechatpay4_1.shtml
		wechatPaySerial := config.GetString("wechat_pay_serial", "")
		if wechatPaySerial != "" {
			df.Header("Wechatpay-Serial", wechatPaySerial)
		}

	}

	returnResponse, err := df.Request()
	if err != nil {
		return returnResponse, err
	}

	// decode response body to outBody
	err = client.HttpHelper.ParseResponseBodyContent(returnResponse, outBody)

	return returnResponse, err

}

func (client *BaseClient) RequestRawXML(ctx context.Context, url string, params *object.StringMap, method string, options *object.HashMap, outHeader interface{}, outBody interface{}) (interface{}, error) {

	config := (*client.App).GetConfig()

	base := &object.StringMap{
		// 微信的接口如果传入接口以外的参数，签名会失败所以这里需要区分对待参数
		"mch_id":    config.GetString("mch_id", ""),
		"nonce_str": object.RandStringBytesMask(32),
	}
	params = object.MergeStringMap(params, base)
	params = object.FilterEmptyStringMap(params)

	// 签名访问的URL，请确保url后面不要跟参数，因为签名的参数，不包含?参数
	// 比如需要在请求的时候，把debug=false，这样url后面就不会多出 "?debug=true"
	options, err := client.AuthSignRequestSimple(config, url, method, params, options)
	if err != nil {
		return nil, err
	}
	// to be setup middleware here
	//client.PushMiddleware(client.logMiddleware(), "access_token")

	httpConfig := client.HttpHelper.GetClient().GetConfig()
	httpConfig.Cert.CertFile = config.GetString("cert_path", "")
	httpConfig.Cert.KeyFile = config.GetString("key_path", "")
	client.HttpHelper.GetClient().SetConfig(&httpConfig)

	// http client request
	df := client.HttpHelper.Df().
		WithContext(ctx).
		Url(url).Method(method)

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
		if (*options)["body"] != nil {
			r := bytes.NewBufferString((*options)["body"].(string))
			df.Body(r)
		}

		// set header
		df.
			Header("content-type", "application/xml").
			//Header("Authorization", (*(*options)["headers"].(*object.HashMap))["Authorization"].(string)).
			Header("Wechatpay-Serial", (*(*options)["headers"].(*object.HashMap))["Wechatpay-Serial"].(string))

	}

	returnResponse, err := df.Request()
	if err != nil {
		return returnResponse, err
	}

	// decode response body to outBody
	err = client.HttpHelper.ParseResponseBodyContent(returnResponse, outBody)

	return returnResponse, err

}

func (client *BaseClient) StreamDownload(ctx context.Context, requestDownload *power.RequestDownload, filePath string) (int64, error) {
	fileHandler, err := os.Create(filePath)
	if err != nil {
		return 0, err
	}
	//defer fileHandler.Close()

	config := (*client.App).GetConfig()

	method := http.MethodPost
	options, err := client.AuthSignRequest(ctx, config, requestDownload.DownloadURL, method, nil, nil)
	if err != nil {
		return 0, err
	}

	rs, err := client.HttpHelper.Df().Url(requestDownload.DownloadURL).Method(method).Json(options).Request()
	if err != nil {
		return 0, err
	}
	result, err := fileHandler.ReadFrom(rs.Body)
	if err != nil {
		return result, err
	}
	fmt2.Dump("http stream download file size:", result)

	// 校验已下载文件
	downloadedHandler, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	//defer downloadedHandler.Close()

	fileMd5 := sha1.New()
	totalSize, err := io.Copy(fileMd5, downloadedHandler)
	if err != nil {
		return 0, err
	}

	//fmt2.Dump(totalSize)

	if requestDownload.HashValue != "" {
		fmt2.Dump(fileMd5.Sum(nil), requestDownload.HashValue)
		if fmt.Sprintf("%x", fileMd5.Sum(nil)) != requestDownload.HashValue {
			return 0, errors.New("文件损坏")
		} else {
			log.Println("文件SHA-256校验成功")
		}
	}

	return totalSize, err
}

func (client *BaseClient) RequestArray(ctx context.Context, url string, method string, options *object.HashMap, outHeader interface{}, outBody interface{}) (*object.HashMap, error) {
	returnResponse, err := client.Request(ctx, url, nil, method, options, true, outHeader, outBody)
	if err != nil {
		return nil, err
	}
	result, err := client.CastResponseToType(returnResponse, response2.TYPE_RAW)

	return result.(*object.HashMap), err
}

func (client *BaseClient) SafeRequestV3(ctx context.Context, url string, params *object.StringMap, method string, option *object.HashMap, outHeader interface{}, outBody interface{}) (interface{}, error) {
	config := (*client.App).GetConfig()

	httpConfig := client.HttpHelper.GetClient().GetConfig()
	httpConfig.Cert.CertFile = config.GetString("cert_path", "")
	httpConfig.Cert.KeyFile = config.GetString("key_path", "")
	client.HttpHelper.GetClient().SetConfig(&httpConfig)

	// get xml string result from return raw as true
	rs, err := client.Request(
		ctx,
		url,
		params,
		method,
		option,
		true,
		outHeader,
		outBody,
	)

	if err != nil {
		return nil, err
	}

	// get out result
	client.HttpHelper.ParseResponseBodyContent(rs, outBody)

	return outBody, err
}

func (client *BaseClient) SafeRequest(ctx context.Context, url string, params *object.HashMap, method string, option *object.HashMap, outHeader interface{}, outBody interface{}) (interface{}, error) {
	config := (*client.App).GetConfig()

	httpConfig := client.HttpHelper.GetClient().GetConfig()
	httpConfig.Cert.CertFile = config.GetString("cert_path", "")
	httpConfig.Cert.KeyFile = config.GetString("key_path", "")
	client.HttpHelper.GetClient().SetConfig(&httpConfig)

	strOutBody := ""
	// get xml string result from return raw as true
	rs, err := client.RequestV2(
		ctx,
		url,
		params,
		method,
		option,
		true,
		outHeader,
		&strOutBody,
	)

	if err != nil {
		return nil, err
	}

	// get out result
	client.HttpHelper.ParseResponseBodyContent(rs, outBody)

	return outBody, err
}

func (client *BaseClient) Wrap(endpoint string) string {
	if (*client.App).InSandbox() {
		return "sandboxnew/" + endpoint
	} else {
		return endpoint
	}
}

func (client *BaseClient) AuthSignRequest(context context.Context, config *kernel.Config, endpoint string, method string, params *object.StringMap, options *object.HashMap) (*object.HashMap, error) {

	var err error

	base := &object.HashMap{}
	isPartnerPay := context.Value("isPartnerPay")
	if isPartnerPay == nil {
		(*base)["appid"] = config.GetString("app_id", "")
		(*base)["mchid"] = config.GetString("mch_id", "")
	}

	// init options
	if options == nil {
		options = &object.HashMap{}
	}

	// init query parameters into body
	if params != nil {
		endpoint += "?" + object.GetJoinedWithKSort(params)
		(*options)["query"] = params
	} else {
		(*options)["query"] = nil
	}

	options = object.MergeHashMap(base, client.prepends(), options)
	options = object.FilterEmptyHashMap(options)

	// check need sign body or not
	signBody := ""
	if "get" != object.Lower(method) {
		signBody, err = object.JsonEncode(options)
		if err != nil {
			return nil, err
		}
	}

	authorization, err := client.Signer.GenerateRequestSign(&support.RequestSignChain{
		Method:       method,
		CanonicalURL: endpoint,
		SignBody:     signBody,
	})

	if err != nil {
		return nil, err
	}

	options = object.MergeHashMap(&object.HashMap{
		"headers": &object.HashMap{
			"Authorization": authorization,
		},
		"body": signBody,
	}, options)

	return options, err
}

func (client *BaseClient) AuthSignRequestSimple(config *kernel.Config, endpoint string, method string, params *object.StringMap, options *object.HashMap) (*object.HashMap, error) {

	var err error

	secretKey, err := (*client.App).GetKey(endpoint)
	if err != nil {
		return nil, err
	}

	// convert StringMap to Power StringMap
	powerStrMapParams, err := power.StringMapToPower(params)
	if err != nil {
		return nil, err
	}

	(*powerStrMapParams)["sign_type"] = "MD5"

	// generate md5 signature with power StringMap
	(*powerStrMapParams)["sign"] = support.GenerateSignMD5(powerStrMapParams, secretKey)

	// convert signature to xml content
	var signBody = ""
	if "get" != object.Lower(method) {
		// check need sign body or not
		objPara, err := power.PowerStringMapToObjectStringMap(powerStrMapParams)
		if err != nil {
			return nil, err
		}
		signBody = object.StringMap2Xml(objPara)
	}

	// set body content
	options = object.MergeHashMap(&object.HashMap{
		"headers": &object.HashMap{
			//"Authorization":    authorization,
			"Wechatpay-Serial": config.GetString("serial_no", ""),
		},
		"body": signBody,
	}, options)

	return options, err
}

func (client *BaseClient) AuthSignRequestV2(endpoint string, method string, params *object.HashMap, options *object.HashMap) (*object.HashMap, error) {

	var err error

	secretKey, err := (*client.App).GetKey(endpoint)
	if err != nil {
		return nil, err
	}

	strMapParams, err := object.HashMapToStringMap(params)
	if err != nil {
		return nil, err
	}

	// convert StringMap to Power StringMap
	powerStrMapParams, err := power.StringMapToPower(strMapParams)
	if err != nil {
		return nil, err
	}

	// generate md5 signature with power StringMap
	(*powerStrMapParams)["sign"] = support.GenerateSignMD5(powerStrMapParams, secretKey)

	// convert signature to xml content
	var signBody = ""
	if "get" != object.Lower(method) {
		// check need sign body or not
		objPara, err := power.PowerStringMapToObjectStringMap(powerStrMapParams)
		if err != nil {
			return nil, err
		}
		signBody = object.StringMap2Xml(objPara)
	}

	// set body content
	options = object.MergeHashMap(&object.HashMap{
		"body": signBody,
	}, options)

	return options, err
}

// ----------------------------------------------------------------------
func (client *BaseClient) RegisterHttpMiddlewares() {

	// access token
	//accessTokenMiddleware := client.GetMiddlewareOfAccessToken
	// check invalid access token
	//checkAccessTokenMiddleware := client.GetMiddlewareOfRefreshAccessToken
	// log
	logMiddleware := client.GetMiddlewareOfLog

	config := (*client.App).GetConfig()
	logger := (*client.App).GetComponent("Logger").(contract2.LoggerInterface)
	client.HttpHelper.WithMiddleware(
		//accessTokenMiddleware,
		logMiddleware(logger),
		//checkAccessTokenMiddleware(3),
		helper.HttpDebugMiddleware(config.GetBool("http_debug", false)),
	)
}

func (client *BaseClient) OverrideGetMiddlewares() {
	client.OverrideGetMiddlewareOfLog()
	//client.OverrideGetMiddlewareOfAccessToken()
	//client.OverrideGetMiddlewareOfRefreshAccessToken()
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
