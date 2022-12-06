package middleware

import (
	"encoding/json"
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/contract"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/coreClient"
	"github.com/pkg/errors"
	"net/http"
	"regexp"
	"time"
)

type AccessTokenSource interface {
	RefreshToken() error
	Expired() bool
	GetToken() (string, error)
}

type RequestMapOptions interface {
	ApplyToRequest(request *http.Request, requestOptions *object.HashMap) (*http.Request, error)
}

const retention = time.Second * 10

type AccessToken struct {
	accessToken contract.AccessTokenInterface
	expireAt    time.Time
}

func (a *AccessToken) RefreshToken() error {
	resToken, err := a.accessToken.GetToken(true)
	a.expireAt = time.Now().Add(time.Duration(resToken.ExpiresIn) * time.Second)
	if err != nil {
		return errors.Wrap(err, "refresh token failed")
	} else {
		return nil
	}
}

func (a *AccessToken) Expired() bool {
	return a.expireAt.IsZero() || a.expireAt.Before(time.Now().Add(-retention))
}

func (a *AccessToken) GetToken() (string, error) {
	resToken, err := a.accessToken.GetToken(false)
	if err != nil {
		return "", errors.Wrap(err, "get access token failed")
	} else {
		return resToken.AccessToken, nil
	}
}

func NewAccessToken(source contract.AccessTokenInterface) *AccessToken {
	return &AccessToken{
		accessToken: source,
	}
}

const accessTokenQueryKey = "access_token"

var tokenInvalidErrCodeRegex, _ = regexp.Compile(`errcode": (40001)|(40014)|(42001 )`)

func QueryAccessTokenMiddleware(tokenSource AccessTokenSource) coreClient.Middleware {
	return func(handle coreClient.HandleFunc) coreClient.HandleFunc {
		return func(req *coreClient.PreRequest, res *coreClient.AfterResponse) error {
			// set access token
			if tokenSource.Expired() {
				if err := tokenSource.RefreshToken(); err != nil {
					return err
				}
			}

			setToken := func() error {
				token, err := tokenSource.GetToken()
				if err != nil {
					return err
				}
				queries := &object.StringMap{}
				if !object.IsObjectNil((*req.Options)["query"]) {
					queries = (*req.Options)["query"].(*object.StringMap)
				}
				(*queries)[accessTokenQueryKey] = token
				(*req.Options)["query"] = queries
				return nil
			}

			if err := setToken(); err != nil {
				return err
			}

			// 序列化请求的 options
			reqData, _ := json.Marshal(req.Options)

			// do request
			if err := handle(req, res); err != nil {
				return err
			}

			// check response & retry
			data, err := res.Result.GetBodyData()
			if err != nil {
				return err
			}
			if tokenInvalidErrCodeRegex.Match(data) {
				// refresh token & retry
				if err = tokenSource.RefreshToken(); err != nil {
					return err
				}
				if err = setToken(); err != nil {
					return err
				}

				// 这里 req 的 option 经过上次请求, 已经发生了变化, 因此使用序列化 options 重置
				var retryData *object.HashMap
				err = json.Unmarshal(reqData, retryData)
				if err != nil {
					return err
				}
				req.Options = retryData

				if err = handle(req, res); err != nil {
					return err
				}
			}

			return nil
		}
	}
}
