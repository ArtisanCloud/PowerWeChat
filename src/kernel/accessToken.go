package kernel

import (
	"errors"
	"github.com/ArtisanCloud/go-libs/object"
	"reflect"
	"time"
)

type AccessToken struct {
	App *ServiceContainer

	RequestMethod      string
	EndpointToGetToken string
	QueryName          string
	Token              object.HashMap
	TokenKey           string
	CachePrefix        string

	GetCredentials func() *object.StringMap
}

func NewAccessToken(container *ServiceContainer) *AccessToken {
	token := &AccessToken{
		App:                container,
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
	typeResponse := reflect.TypeOf(response)
	//fmt.Printf("response name:%s, type: %s \n",typeResponse.Name(), typeResponse.Kind() )
	if typeResponse.Name() == "ResponseInterface" {
		token = response.(*object.HashMap)
	}

	return token
}

func (component *AccessToken) SetToken(token string, lifeTime time.Duration) (token *object.HashMap, err error) {
	if lifeTime <= 0 {
		lifeTime = 7200 * time.Second
	}

	cache := component.GetCache()
	cache.Set(component.GetCacheKey(), map[string]interface{}{
		component.TokenKey: token,
		"expires_in":       lifeTime,
	})

	if !cache.Has(component.GetCacheKey()) {
		err = errors.New("Failed to cache access token.")
	}
	return token, err

}

func (component *AccessToken) requestToken(credentials *object.StringMap, toArray bool) interface{} {
	//fmt.Printf("credentials:%v \n", credentials)

	return nil
}
