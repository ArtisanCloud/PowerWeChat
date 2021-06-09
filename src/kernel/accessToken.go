package kernel

import (
	"crypto/md5"
	"encoding/json"
	"github.com/ArtisanCloud/go-libs/exception"
	fmt2 "github.com/ArtisanCloud/go-libs/fmt"
	http2 "github.com/ArtisanCloud/go-libs/http"
	httpContract "github.com/ArtisanCloud/go-libs/http/contract"
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/go-wechat/src/kernel/contract"
	"net/http"
	"net/url"
	"time"
)

type ResponseGetToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	*ResponseWX
}

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

func (accessToken *AccessToken) GetRefreshedToken() *object.HashMap {
	return accessToken.GetToken(true)
}

func (accessToken *AccessToken) GetToken(refresh bool) (token *object.HashMap) {
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
	response := accessToken.requestToken(accessToken.GetCredentials(), true)

	//// save token into cache
	token = response.(*object.HashMap)
	strToken := (*token)[accessToken.TokenKey].(string)
	var expireIn time.Duration = 7200 * time.Second
	if (*token)[accessToken.TokenKey] == nil {
		expireIn = (*token)[accessToken.TokenKey].(time.Duration) * time.Second
	}
	fmt2.Dump(token, strToken, expireIn)
	//token, _ = accessToken.SetToken(strToken, expireIn)

	//typeResponse := reflect.TypeOf(response)
	////fmt.Printf("response name:%s, type: %s \n",typeResponse.Name(), typeResponse.Kind() )
	//if typeResponse.Name() == "ResponseInterface" {
	//	token = response.(*object.HashMap)
	//}

	return token
}

func (accessToken *AccessToken) SetToken(token string, lifeTime time.Duration) (*object.HashMap, error) {
	if lifeTime <= 0 {
		lifeTime = 7200 * time.Second
	}

	//cache := accessToken.GetCache()
	//cache.Set(accessToken.GetCacheKey(), map[string]interface{}{
	//	accessToken.TokenKey: token,
	//	"expires_in":       lifeTime,
	//})
	//
	//if !cache.Has(accessToken.GetCacheKey()) {
	//	err = errors.New("Failed to cache access token.")
	//}
	//return token, err
	return nil, nil

}

func (accessToken *AccessToken) Refresh() *contract.AccessTokenInterface {

	return nil
}

func (accessToken *AccessToken) requestToken(credentials *object.StringMap, toArray bool) interface{} {

	accessToken.sendRequest(credentials)

	return nil
}

func (accessToken *AccessToken) ApplyToRequest(request *http.Request, requestOptions *object.HashMap) *http.Request {
	// parse query
	// merge query with query again
	hashToken := accessToken.getQuery()

	strMapToString := object.ConvertStringMapToString(hashToken)
	uriToken, _ := url.Parse(strMapToString)
	request.URL.ResolveReference(uriToken)
	return request
}

func (accessToken *AccessToken) sendRequest(credential *object.StringMap) *httpContract.ResponseContract {

	key := "json"
	if accessToken.RequestMethod == "GET" {
		key = "query"
	}
	options := &object.HashMap{
		key: credential,
	}

	accessToken.SetHttpClient(accessToken.GetHttpClient()).PerformRequest(
			accessToken.getEndpoint(),
			accessToken.RequestMethod,
			options,
			&ResponseGetToken{},
		)
	return nil
}

func (accessToken *AccessToken) getCacheKey() string {
	data, _ := json.Marshal(accessToken.GetCredentials())
	buffer := md5.Sum(data)
	return accessToken.CachePrefix + string(buffer[:])
}

func (accessToken *AccessToken) getQuery() *object.StringMap {
	var key string
	if accessToken.QueryName != "" {
		key = accessToken.QueryName
	} else {
		key = accessToken.TokenKey
	}

	token := accessToken.GetToken(false)
	arrayReturn := &object.StringMap{
		key: (*token)[accessToken.TokenKey].(string),
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
