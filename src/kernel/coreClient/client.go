package coreClient

import (
	"crypto/tls"
	"fmt"
	"github.com/ArtisanCloud/PowerLibs/v2/http/contract"
	"github.com/ArtisanCloud/PowerLibs/v2/http/drivers/gout"
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"net"
	"net/http"
	"time"
)

// CoreClient 是SDK的封装Client, 实现中间件和通用请求逻辑
type CoreClient interface {
	SetDefaultOptions(defaults *object.HashMap)
	GetDefaultOptions() *object.HashMap
	SetHttpClient(httpClient contract.ClientInterface) CoreClient
	GetHttpClient() contract.ClientInterface
	PerformRequest(url string, method string, options *object.HashMap, returnRaw bool, outHeader interface{}, outBody interface{}) (contract.ResponseInterface, error)
}

var _defaults *object.HashMap

type RequestClient struct {
	httpClient contract.ClientInterface
	BaseURI    string
	doRequest  HandleFunc
}

func NewClient(config *object.HashMap, middlewares ...Middleware) (CoreClient, error) {
	var httpClient *http.Client

	if (*config)["cert_path"] != nil && (*config)["key_path"] != nil {
		certPath := (*config)["cert_path"].(string)
		keyPath := (*config)["key_path"].(string)
		if certPath != "" && keyPath != "" {
			var err error
			httpClient, err = newTLSHttpClient(certPath, keyPath)
			if err != nil {
				return nil, err
			}
		}
	}

	client := gout.NewClient(config, httpClient)

	// build handle request func
	var handle HandleFunc = func(req *PreRequest, res *AfterResponse) error {
		result, err := client.Request(req.Method, req.Url, req.Options, req.ReturnRaw, res.OutHeader, res.OutBody)
		if err != nil {
			return err
		}
		res.Result = result
		return nil
	}
	for _, md := range middlewares {
		handle = md(handle)
	}

	return &RequestClient{
		httpClient: client,
		doRequest:  handle,
	}, nil
}

func (r *RequestClient) SetDefaultOptions(defaults *object.HashMap) {
	_defaults = defaults
}

func (r *RequestClient) GetDefaultOptions() *object.HashMap {
	return _defaults
}

func (r *RequestClient) SetHttpClient(httpClient contract.ClientInterface) CoreClient {
	r.httpClient = httpClient
	return r
}

func (r *RequestClient) GetHttpClient() contract.ClientInterface {
	return r.httpClient
}

func (r *RequestClient) PerformRequest(url string, method string, options *object.HashMap, returnRaw bool, outHeader interface{}, outBody interface{}) (contract.ResponseInterface, error) {
	req := PreRequest{
		Url:       url,
		Method:    method,
		Options:   options,
		ReturnRaw: returnRaw,
	}
	res := AfterResponse{
		OutHeader: outHeader,
		OutBody:   outBody,
	}

	err := r.doRequest(&req, &res)
	return res.Result, err
}

type PreRequest struct {
	Url       string
	Method    string
	Options   *object.HashMap
	ReturnRaw bool
}

type AfterResponse struct {
	OutHeader interface{}
	OutBody   interface{}
	Result    contract.ResponseInterface
}

type HandleFunc func(req *PreRequest, res *AfterResponse) error

type Middleware func(handle HandleFunc) HandleFunc

func ExampleMiddleware(handle HandleFunc) HandleFunc {
	return func(req *PreRequest, res *AfterResponse) error {
		// do something before request

		return handle(req, res)

		// do something after request
	}
}

func newTLSHttpClient(certFile string, keyFile string) (httpClient *http.Client, err error) {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		fmt.Print("can not init cert...")
		return nil, err
	}
	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: false,
		ClientAuth:         tls.RequireAndVerifyClientCert,
		MinVersion:         tls.VersionTLS12,
	}

	httpClient = &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			Dial: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).Dial,
			TLSHandshakeTimeout: 10 * time.Second,
			TLSClientConfig:     tlsConfig,
		},
		Timeout: 60 * time.Second,
	}
	return
}
