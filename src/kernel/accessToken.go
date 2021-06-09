package kernel

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/ArtisanCloud/go-libs/exception"
	httpContract "github.com/ArtisanCloud/go-libs/http/contract"
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/go-wechat/src/kernel/contract"
	"reflect"
	"time"
)

type AccessToken struct {
	App *ApplicationInterface

	RequestMethod      string
	EndpointToGetToken string
	QueryName          string
	Token              object.HashMap
	TokenKey           string
	CachePrefix        string

	GetCredentials func() *object.StringMap
}

func NewAccessToken(app *ApplicationInterface) *AccessToken {
	token := &AccessToken{
		App:                app,
		RequestMethod:      "GET",
		EndpointToGetToken: "",
		QueryName:          "",
		Token:              nil,
		TokenKey:           "access_token",
		CachePrefix:        "ac.go.wechat.kernel.access_token.",
	}

	return token
}

func (component *AccessToken) GetRefreshedToken() *object.HashMap {
	return component.GetToken()
}

func (component *AccessToken) GetToken() (token *object.HashMap) {
	//cacheKey := token.GetCacheKey()
	//cache:= GetCache()

	response := component.requestToken(component.GetCredentials(), true)
	return nil
	typeResponse := reflect.TypeOf(response)
	//fmt.Printf("response name:%s, type: %s \n",typeResponse.Name(), typeResponse.Kind() )
	if typeResponse.Name() == "ResponseInterface" {
		token = response.(*object.HashMap)
	}

	return token
}

func (component *AccessToken) SetToken(token string, lifeTime time.Duration) (*object.HashMap, error) {
	if lifeTime <= 0 {
		lifeTime = 7200 * time.Second
	}

	//cache := component.GetCache()
	//cache.Set(component.GetCacheKey(), map[string]interface{}{
	//	component.TokenKey: token,
	//	"expires_in":       lifeTime,
	//})
	//
	//if !cache.Has(component.GetCacheKey()) {
	//	err = errors.New("Failed to cache access token.")
	//}
	//return token, err
	return nil, nil

}

func (component *AccessToken) Refresh() *contract.AccessTokenInterface {

	return nil
}

func (component *AccessToken) requestToken(credentials *object.StringMap, toArray bool) interface{} {
	fmt.Printf("credentials:%v \n", credentials)

	return nil
}

func (component *AccessToken) ApplyToRequest(request *httpContract.RequestInterface, requestOptions object.HashMap) httpContract.RequestInterface {
	// parse query
	//strURL := request.GetUri().GetQuery()
	//strURL = parsestring(strURL)

	// merge query with query again
	//query = http

	return request
}

func (component *AccessToken) sendRequest(credential *object.HashMap) *httpContract.ResponseContract {
	return nil
}

func (component *AccessToken) getCacheKey() string {
	data, _ := json.Marshal(component.GetCredentials())
	buffer := md5.Sum(data)
	return component.CachePrefix + string(buffer[:])
}

func (component *AccessToken) getQuery() interface{} {
	if component.QueryName != "" {
		return []string{component.QueryName}
	} else {
		arrayReturn := []object.HashMap{}
		token := component.GetToken()
		result := object.HashMap{
			component.TokenKey: (*token)[component.TokenKey],
		}
		arrayReturn = append(arrayReturn, result)

		return arrayReturn

	}
}

func (component *AccessToken) getEndpoint() string {
	defer (&exception.Exception{}).HandleException(nil, "accessToken.get.endpoint", component)
	if component.EndpointToGetToken == "" {
		panic("No endpoint for access token request.")
		return ""
	}

	return component.EndpointToGetToken
}

func (component *AccessToken) getTokenKey() string {
	return component.TokenKey
}
