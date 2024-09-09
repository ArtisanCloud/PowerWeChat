package contract

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	response2 "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"net/http"
)

type (
	AccessTokenInterface interface {
		GetToken(ctx context.Context, refresh bool) (resToken *response2.ResponseGetToken, err error)
		Refresh(ctx context.Context) AccessTokenInterface
		ApplyToRequest(request *http.Request, requestOptions *object.HashMap) (*http.Request, error)
	}
)
