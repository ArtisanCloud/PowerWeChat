package support

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ArtisanCloud/go-libs/object"
	"go/types"
	"io/ioutil"
	"net/http"
	"reflect"
)

type ResponseCastable struct {
}

func (responseCastable *ResponseCastable) CastResponseToType(response *http.Response, castType string) (interface{}, error) {

	return response, nil

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	switch castType {

	case "array":
		data := &object.HashMap{}
		err = json.Unmarshal(body, data)

		return data, err
	case "power":
		var data interface{}
		err = json.Unmarshal(body, &data)
		return data, err
	case "raw":
		return response, nil
	default:

	}

	return nil, nil
}

func (responseCastable *ResponseCastable) DetectAndCastResponseToType(response interface{}, toType string) (interface{}, error) {

	var returnResponse http.Response = http.Response{
		StatusCode: 200,
		Header:     nil,
	}
	switch response.(type) {
	case http.Response:
		returnResponse = response.(http.Response)
		break
	case types.Array:

		postBodyBuf, _ := json.Marshal(response.(types.Array))
		returnResponse.Body = ioutil.NopCloser(bytes.NewBuffer(postBodyBuf))

		break
	case object.HashMap:
		postBodyBuf, _ := json.Marshal(response.(object.HashMap))
		returnResponse.Body = ioutil.NopCloser(bytes.NewBuffer(postBodyBuf))

		break
	case *object.HashMap:
		postBodyBuf, _ := json.Marshal(response.(*object.HashMap))
		returnResponse.Body = ioutil.NopCloser(bytes.NewBuffer(postBodyBuf))

		break
	case string:
		postBodyBuf, _ := json.Marshal(response.(string))
		returnResponse.Body = ioutil.NopCloser(bytes.NewBuffer(postBodyBuf))

		break
	default:
		return nil, errors.New(fmt.Sprintf("unsupported response type '%s'", reflect.TypeOf(response).Kind().String()))
	}

	return responseCastable.CastResponseToType(&returnResponse, toType)
}
