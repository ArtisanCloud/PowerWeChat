package kernel

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/ArtisanCloud/PowerLibs/v3/cache"
	"github.com/ArtisanCloud/PowerLibs/v3/http/helper"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/contract"
	response2 "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"net/http"
	"time"
)

type AccessToken struct {
	App *ApplicationInterface

	HttpHelper *helper.RequestHelper

	RequestMethod      string
	EndpointToGetToken string
	QueryName          string
	Token              *object.HashMap
	TokenKey           string
	CachePrefix        string

	*InteractsWithCache

	GetCredentials func() *object.StringMap
	GetEndpoint    func() (string, error)
}

func NewAccessToken(app *ApplicationInterface) (*AccessToken, error) {
	config := (*app).GetConfig()
	baseURI := config.GetString("http.base_uri", "/")

	var cacheClient cache.CacheInterface = nil
	c := config.Get("cache", nil)
	if c != nil {
		cacheClient = c.(cache.CacheInterface)
	}

	h, err := helper.NewRequestHelper(&helper.Config{
		BaseUrl: baseURI,
	})
	if err != nil {
		return nil, err
	}
	token := &AccessToken{
		App:        app,
		HttpHelper: h,

		RequestMethod:      "GET",
		EndpointToGetToken: "",
		QueryName:          "",
		Token:              nil,
		TokenKey:           "access_token",
		CachePrefix:        "powerwechat.access_token.",
		InteractsWithCache: NewInteractsWithCache(cacheClient),
	}

	token.OverrideGetEndpoint()

	return token, nil
}

func (accessToken *AccessToken) GetRefreshedToken() (*response2.ResponseGetToken, error) {
	return accessToken.GetToken(true)
}

func (accessToken *AccessToken) GetToken(refresh bool) (resToken *response2.ResponseGetToken, err error) {
	cacheKey := accessToken.GetCacheKey()
	cache := accessToken.GetCache()

	// get token from cache
	if !refresh && cache.Has(cacheKey) {
		value, err := cache.Get(cacheKey, nil)
		if err == nil && value != nil {

			token := (object.HashMap)(value.(map[string]interface{}))

			resToken = &response2.ResponseGetToken{
				ExpiresIn: token["expires_in"].(float64),
			}

			if accessToken.TokenKey == "access_token" && token["access_token"] != nil {
				resToken.AccessToken = token[accessToken.TokenKey].(string)

			} else if accessToken.TokenKey == "component_access_token" && token["component_access_token"] != nil {
				resToken.AccessToken = token[accessToken.TokenKey].(string)
				resToken.ComponentAccessToken = token[accessToken.TokenKey].(string)

			} else if accessToken.TokenKey == "authorizer_access_token" && token["authorizer_access_token"] != nil {
				resToken.AccessToken = token[accessToken.TokenKey].(string)
				resToken.AuthorizerAccessToken = token[accessToken.TokenKey].(string)

			} else if accessToken.TokenKey == "authorizer_refresh_token" && token["authorizer_refresh_token"] != nil {
				resToken.AccessToken = token[accessToken.TokenKey].(string)
				resToken.AuthorizerRefreshToken = token[accessToken.TokenKey].(string)

			} else {
				return nil, errors.New("no token found in cache")
			}

			return resToken, err
		}
	}

	// request token from power
	resToken, err = accessToken.requestToken(accessToken.GetCredentials())
	if err != nil {
		return nil, err
	}
	_, err = accessToken.SetToken(resToken)

	// tbd dispatch an event for AccessTokenRefresh

	return resToken, err
}

func (accessToken *AccessToken) SetToken(token *response2.ResponseGetToken) (tokenInterface contract.AccessTokenInterface, err error) {
	if token.ExpiresIn <= 0 {
		token.ExpiresIn = 7200
	}

	// set token into cache
	cache := accessToken.GetCache()
	err = cache.Set(accessToken.GetCacheKey(), token, time.Duration(token.ExpiresIn)*time.Second)

	if err != nil {
		return nil, err
	}

	if !cache.Has(accessToken.GetCacheKey()) {
		return nil, errors.New("failed to cache access token")
	}
	return accessToken, err

}

func (accessToken *AccessToken) Refresh() contract.AccessTokenInterface {
	accessToken.GetToken(true)

	return accessToken
}

func (accessToken *AccessToken) requestToken(credentials *object.StringMap) (*response2.ResponseGetToken, error) {

	res, err := accessToken.sendRequest(credentials)
	if err != nil {
		return nil, err
	}
	token := res

	if token == nil || (token.AccessToken == "" &&
		token.ComponentAccessToken == "" &&
		token.AuthorizerAccessToken == "" &&
		token.AuthorizerRefreshToken == "") {
		return nil, errors.New(fmt.Sprintf("Request access_token fail: %v", res))
	}

	return token, nil
}

func (accessToken *AccessToken) ApplyToRequest(request *http.Request, requestOptions *object.HashMap) (*http.Request, error) {

	// query Access Token power
	mapToken, err := accessToken.getQuery()
	if err != nil {
		return nil, err
	}
	q := request.URL.Query()
	for key, value := range *mapToken {
		q.Set(key, value)
	}
	request.URL.RawQuery = q.Encode()

	return request, err
}

func (accessToken *AccessToken) sendRequest(credential *object.StringMap) (*response2.ResponseGetToken, error) {

	key := "json"
	if accessToken.RequestMethod == "GET" {
		key = "query"
	}
	options := &object.HashMap{
		key: credential,
	}

	res := &response2.ResponseGetToken{}

	strEndpoint, err := accessToken.GetEndpoint()
	if err != nil {
		return nil, err
	}

	df := accessToken.HttpHelper.Df().Uri(strEndpoint).
		Method(accessToken.RequestMethod)

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

		// set body json
		if (*options)["form_params"] != nil {
			df.Json((*options)["form_params"])
		}
	}

	rs, err := df.Request()
	if err != nil {
		return nil, err
	}

	// decode response body to outBody
	err = accessToken.HttpHelper.ParseResponseBodyContent(rs, res)

	return res, err
}

// GetCacheKey 缓存唯一key算法说明：
// 1. 使用appid和secret进行字符串拼接。例如: appid=testappid secret=testsecret, 那么结果为: testappidtestsecret
// 2. 计算字符串的md5。"testappidtestsecret"的md5值为"edc5f6181730baffc0b88cf96658aeff"
// 3. 加上PowerWeChat前缀命名空间："powerwechat.access_token."，最终结果为："powerwechat.access_token.edc5f6181730baffc0b88cf96658aeff"
func (accessToken *AccessToken) GetCacheKey() string {
	credentials := *accessToken.GetCredentials()
	data := fmt.Sprintf("%s%s%s", credentials["appid"], credentials["secret"], credentials["neededText"])
	buffer := md5.Sum([]byte(data))
	cacheKey := accessToken.CachePrefix + hex.EncodeToString(buffer[:])

	// tbf
	//fmt2.Dump(cacheKey)

	return cacheKey
}

func (accessToken *AccessToken) getQuery() (*object.StringMap, error) {
	// set the current token key
	var key string
	if accessToken.QueryName != "" {
		key = accessToken.QueryName
	} else {
		key = accessToken.TokenKey
	}

	// get token string power
	resToken, err := accessToken.GetToken(false)
	if err != nil {
		return nil, err
	}
	arrayReturn := &object.StringMap{
		key: resToken.AccessToken,
	}

	return arrayReturn, err

}

func (accessToken *AccessToken) OverrideGetEndpoint() {

	accessToken.GetEndpoint = func() (string, error) {
		if accessToken.EndpointToGetToken == "" {
			return "", errors.New("no endpoint for access token request")
		}

		return accessToken.EndpointToGetToken, nil
	}

}

func (accessToken *AccessToken) getTokenKey() string {
	return accessToken.TokenKey
}
