package kernel

import (
	"crypto/sha1"
	"encoding/xml"
	"errors"
	"fmt"
	fmt2 "github.com/ArtisanCloud/PowerLibs/fmt"
	"github.com/ArtisanCloud/PowerLibs/http/request"
	"github.com/ArtisanCloud/PowerLibs/object"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	response2 "github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/support"
	"io"
	"log"
	http2 "net/http"
	"os"
)

type BaseClient struct {
	kernel.BaseClient

	App *ApplicationPaymentInterface
}

func NewBaseClient(app *ApplicationPaymentInterface) *BaseClient {
	config := (*app).GetContainer().GetConfig()

	client := &BaseClient{
		BaseClient: kernel.BaseClient{
			HttpRequest: request.NewHttpRequest(config),
			Signer: &support.SHA256WithRSASigner{
				MchID:               (*config)["mch_id"].(string),
				CertificateSerialNo: (*config)["serial_no"].(string),
				PrivateKeyPath:      (*config)["key_path"].(string),
			},
		},
		App: app,
	}
	return client

}

func (client *BaseClient) prepends() *object.HashMap {
	return &object.HashMap{}
}

func (client *BaseClient) PlainRequest(endpoint string, params *object.StringMap, method string, options *object.HashMap,
	returnRaw bool, outHeader interface{}, outBody interface{},
) (response interface{}, err error) {

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

	options = object.MergeHashMap(&object.HashMap{
		"headers": &object.HashMap{
			"Authorization": authorization,
		},
		"body": signBody,
	}, options)

	// to be setup middleware here
	//client.PushMiddleware(client.logMiddleware(), "access_token")

	// http client request
	returnResponse, err := client.PerformRequest(endpoint, method, options, returnRaw, outHeader, outBody)
	if err != nil {
		return nil, err
	}

	if returnRaw {
		return returnResponse, nil
	} else {
		var rs http2.Response = http2.Response{
			StatusCode: 200,
			Header:     nil,
		}
		rs.Body = returnResponse.GetBody()
		result, _ := client.CastResponseToType(&rs, response2.TYPE_RAW)
		return result, nil
	}

}

func (client *BaseClient) RequestV2(endpoint string, params *object.HashMap, method string, option *object.HashMap,
	returnRaw bool, outHeader interface{}, outBody interface{},
) (response interface{}, err error) {

	//config := (*client.App).GetConfig()

	base := &object.HashMap{
		// 微信的接口如果传入接口意外的参数，签名会失败所以这里需要区分对待参数
		"nonce_str":  object.RandStringBytesMask(32),
		//"mch_id":     config.GetString("mch_id", ""),
		//"sub_mch_id": config.GetString("sub_mch_id", ""),
		//"sub_appid":  config.GetString("sub_appid", ""),
	}
	params = object.MergeHashMap(params, base)
	params = object.FilterEmptyHashMap(params)

	options, err := client.AuthSignRequestV2(endpoint, method, params, option)
	if err != nil {
		return nil, err
	}

	// to be setup middleware here
	//client.PushMiddleware(client.logMiddleware(), "access_token")

	// http client request
	returnResponse, err := client.PerformRequest(endpoint, method, options, returnRaw, outHeader, outBody)
	if err != nil {
		return nil, err
	}

	if returnRaw {
		return returnResponse, nil
	} else {
		var rs http2.Response = http2.Response{
			StatusCode: 200,
			Header:     nil,
		}
		rs.Body = returnResponse.GetBody()
		result, _ := client.CastResponseToType(&rs, response2.TYPE_MAP)
		return result, nil
	}

}

func (client *BaseClient) Request(endpoint string, params *object.StringMap, method string, options *object.HashMap,
	returnRaw bool, outHeader interface{}, outBody interface{},
) (response interface{}, err error) {

	config := (*client.App).GetConfig()

	// 签名访问的URL，请确保url后面不要跟参数，因为签名的参数，不包含?参数
	// 比如需要在请求的时候，把debug=false，这样url后面就不会多出 "?debug=true"
	options, err = client.AuthSignRequest(config, endpoint, method, params, options)
	if err != nil {
		return nil, err
	}
	// to be setup middleware here
	//client.PushMiddleware(client.logMiddleware(), "access_token")

	// http client request
	returnResponse, err := client.PerformRequest(endpoint, method, options, returnRaw, outHeader, outBody)
	if err != nil {
		return nil, err
	}

	if returnRaw {
		return returnResponse, nil
	} else {
		var rs http2.Response = http2.Response{
			StatusCode: 200,
			Header:     nil,
		}
		rs.Body = returnResponse.GetBody()
		result, _ := client.CastResponseToType(&rs, response2.TYPE_MAP)
		return result, nil
	}

}

func (client *BaseClient) RequestRaw(url string, params *object.StringMap, method string, options *object.HashMap, outHeader interface{}, outBody interface{}) (interface{}, error) {
	return client.Request(url, params, method, options, true, outHeader, outBody)
}

func (client *BaseClient) StreamDownload(requestDownload *power.RequestDownload, filePath string) (int64, error) {
	fileHandler, err := os.Create(filePath)
	if err != nil {
		return 0, err
	}
	defer fileHandler.Close()

	config := (*client.App).GetConfig()

	method := "GET"
	options, err := client.AuthSignRequest(config, requestDownload.DownloadURL, method, nil, nil)
	if err != nil {
		return 0, err
	}

	_, err = client.PerformRequest(requestDownload.DownloadURL, method, options, true, nil, fileHandler)
	if err != nil {
		return 0, err
	}

	// 校验下载文件
	downloadedHandler, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer downloadedHandler.Close()

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

func (client *BaseClient) RequestArray(url string, method string, options *object.HashMap, outHeader interface{}, outBody interface{}) (*object.HashMap, error) {
	returnResponse, err := client.RequestRaw(url, nil, method, options, outHeader, outBody)
	if err != nil {
		return nil, err
	}
	result, err := client.CastResponseToType(returnResponse.(*http2.Response), "array")

	return result.(*object.HashMap), err
}

func (client *BaseClient) SafeRequest(url string, params *object.HashMap, method string, option *object.HashMap, outHeader interface{}, outBody interface{}) (interface{}, error) {
	config := (*client.App).GetConfig()
	option = object.MergeHashMap(option, &object.HashMap{
		"cert":    config.GetString("cert_path", ""),
		"ssl_key": config.GetString("key_path", ""),
	})

	strOutBody := ""
	// get xml string result from return raw as true
	_, err := client.RequestV2(
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
	err = xml.Unmarshal([]byte(strOutBody), outBody)

	return outBody, err
}

func (client *BaseClient) Wrap(endpoint string) string {
	if (*client.App).InSandbox() {
		return "sandboxnew/" + endpoint
	} else {
		return endpoint
	}
}

func (client *BaseClient) AuthSignRequest(config *kernel.Config, endpoint string, method string, params *object.StringMap, options *object.HashMap) (*object.HashMap, error) {

	var err error

	base := &object.HashMap{
		"appid": config.GetString("app_id", ""),
		"mchid": config.GetString("mch_id", ""),
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

func (client *BaseClient) AuthSignRequestV2(endpoint string, method string, params *object.HashMap, options *object.HashMap) (*object.HashMap, error) {

	var err error

	secretKey, err := (*client.App).GetKey(endpoint)
	if err != nil {
		return nil, err
	}

	strMapParams,err := object.HashMapToStringMap(params)
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
		if err!= nil{
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
type MiddlewareLogMiddleware struct {
	*BaseClient
}

func (client *BaseClient) logMiddleware() interface{} {
	return &MiddlewareLogMiddleware{
		client,
	}
}
