package kernel

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/ArtisanCloud/go-libs/exception"
	http2 "github.com/ArtisanCloud/go-libs/http"
	httpContract "github.com/ArtisanCloud/go-libs/http/contract"
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/go-wechat/src/kernel/contract"
	response2 "github.com/ArtisanCloud/go-wechat/src/kernel/response"
	"net/http"
	"time"
)

type AccessToken struct {
	App *ApplicationInterface

	*http2.HttpRequest
	*http2.HttpResponse

	RequestMethod      string
	EndpointToGetToken string
	QueryName          string
	Token              object.HashMap
	TokenKey           string
	CachePrefix        string

	*CacheToken

	GetCredentials func() *object.StringMap
}

func NewAccessToken(app *ApplicationInterface) *AccessToken {
	config := (*app).GetContainer().GetConfig()

	token := &AccessToken{
		App:         app,
		HttpRequest: http2.NewHttpRequest(config),

		RequestMethod:      "GET",
		EndpointToGetToken: "",
		QueryName:          "",
		Token:              nil,
		TokenKey:           "access_token",
		CachePrefix:        "ac.go.wechat.kernel.access_token.",
	}

	return token
}

func (accessToken *AccessToken) GetRefreshedToken() *response2.ResponseGetToken {
	return accessToken.GetToken(true)
}

func (accessToken *AccessToken) GetToken(refresh bool) (resToken *response2.ResponseGetToken) {
	//cacheKey := accessToken.getCacheKey()
	//cache := *accessToken.getCache()
	//
	//// get token from cache
	//if !refresh && cache.Has(cacheKey) {
	//	token = &object.HashMap{}
	//	err := cache.Get(cacheKey, token)
	//	if err == nil {
	//		return token
	//	}
	//}

	// request token from wx
	response := accessToken.requestToken(accessToken.GetCredentials())

	//// save token into cache
	resToken = response.(*response2.ResponseGetToken)
	var expireIn time.Duration = 7200 * time.Second
	if resToken.ExpiresIn > 0 {
		expireIn = time.Duration(resToken.ExpiresIn) * time.Second
	}
	accessToken.SetToken(resToken.AccessToken, expireIn)

	// tbd dispatch a event for AccessTokenRefresh

	return resToken
}

func (accessToken *AccessToken) SetToken(token string, lifeTime time.Duration) (tokenInterface contract.AccessTokenInterface, err error) {
	if lifeTime <= 0 {
		lifeTime = 7200 * time.Second
	}

	// tbd
	cache := accessToken.getCache()
	err = cache.Set(accessToken.getCacheKey(), map[string]interface{}{
		accessToken.TokenKey: token,
		"expires_in":         lifeTime,
	}, lifeTime)

	defer (&exception.Exception{}).HandleException(nil, "accessToken.set.token", nil)
	if !cache.Has(accessToken.getCacheKey()) {
		panic("Failed to cache access token.")
		return nil, err
	}
	return accessToken, err

}

func (accessToken *AccessToken) Refresh() contract.AccessTokenInterface {

	return nil
}

func (accessToken *AccessToken) requestToken(credentials *object.StringMap) httpContract.ResponseContract {

	// tbf
	return &response2.ResponseGetToken{
		AccessToken: "Np9hU7IcFX8fajwPvhI4-_ZIJh0kDIBambDyI8nYuNMNsO74lUAgkA75iETLu2mAv05EkAwZf9caFupjQuVATBo17lk7xqNnqw_e5NpfdjEK9F5moG755THiHOd_XSGFb6a9QTggn_t6ejA7CcxlKetr6ntcxdRomQC9rnJ-razLKrWrJilWERDHVtwxg4TGh86qFvOjRZ7OMBznnmrbbQ",
		ExpiresIn:   7200,
		ResponseWX: &response2.ResponseWX{
			ErrCode: 0,
			ErrMSG:  "ok",
		},
	}

	res := accessToken.sendRequest(credentials)
	token := res.(*response2.ResponseGetToken)
	defer (&exception.Exception{}).HandleException(nil, "accessToken.request.token", nil)
	if token == nil || token.AccessToken == "" {
		panic(fmt.Sprintf("Request access_token fail: %v", res))
		return nil
	}

	return token
}

func (accessToken *AccessToken) ApplyToRequest(request *http.Request, requestOptions *object.HashMap) *http.Request {

	// query Access Token map
	mapToken := accessToken.getQuery()
	q := request.URL.Query()
	for key, value := range *mapToken {
		q.Set(key, value)
	}
	request.URL.RawQuery = q.Encode()

	return request
}

func (accessToken *AccessToken) sendRequest(credential *object.StringMap) httpContract.ResponseContract {

	key := "json"
	if accessToken.RequestMethod == "GET" {
		key = "query"
	}
	options := &object.HashMap{
		key: credential,
	}

	res := &response2.ResponseGetToken{}

	accessToken.SetHttpClient(accessToken.GetHttpClient()).PerformRequest(
		accessToken.getEndpoint(),
		accessToken.RequestMethod,
		options,
		res,
	)
	return res
}

func (accessToken *AccessToken) getCacheKey() string {
	data, _ := json.Marshal(accessToken.GetCredentials())
	buffer := md5.Sum(data)
	return accessToken.CachePrefix + string(buffer[:])
}

func (accessToken *AccessToken) getQuery() *object.StringMap {
	// set the current token key
	var key string
	if accessToken.QueryName != "" {
		key = accessToken.QueryName
	} else {
		key = accessToken.TokenKey
	}

	// get token string map
	resToken := accessToken.GetToken(false)
	arrayReturn := &object.StringMap{
		key: resToken.AccessToken,
	}

	return arrayReturn

}

func (accessToken *AccessToken) getEndpoint() string {
	defer (&exception.Exception{}).HandleException(nil, "accessToken.get.endpoint", accessToken)
	if accessToken.EndpointToGetToken == "" {
		panic("No endpoint for access token request.")
		return ""
	}

	return accessToken.EndpointToGetToken
}

func (accessToken *AccessToken) getTokenKey() string {
	return accessToken.TokenKey
}
