package support

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	response2 "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"go/types"
	"io/ioutil"
	"net/http"
	"reflect"
)

type ResponseCastable struct {
}

func (responseCastable *ResponseCastable) CastResponseToType(response *http.Response, castType string) (interface{}, error) {

	if castType == response2.TYPE_RAW {
		return response, nil
	}

	switch castType {

	case response2.TYPE_ARRAY:
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		data := []interface{}{}
		err = json.Unmarshal(body, data)

		return data, err
	case response2.TYPE_MAP:
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		data := &object.HashMap{}
		err = json.Unmarshal(body, data)

		return data, err

	case response2.TYPE_RAW:

		return response, nil
	default:
		return nil, errors.New("Config key \"response_type\" classname must be an instanceof %s'")
	}
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
	case *http.Response:
		returnResponse = *response.(*http.Response)
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
