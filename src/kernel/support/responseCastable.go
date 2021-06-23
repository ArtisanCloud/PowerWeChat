package support

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"go/types"
	"io/ioutil"
	"net/http"
	"reflect"
)

type ResponseCastable struct {
}

func (responseCastable *ResponseCastable) castResponseToType(response *http.Response, castType string) (interface{}, error) {

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var data interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	switch castType {

	case "array":
		return data, nil
	case "object":
		return data, nil
	case "raw":
		return response, nil
	default:

	}

	return nil, nil
}

func (responseCastable *ResponseCastable) DetectAndCastResponseToType(response interface{}, toType string) (interface{}, error) {

	var returnResponse http.Response
	switch response.(type) {
	case http.Response:
		returnResponse = response.(http.Response)
		break
	case types.Array:
		returnResponse = http.Response{
			StatusCode: 200,
			Header:     nil,
		}
		postBodyBuf, _ := json.Marshal(response.(types.Array))
		returnResponse.Body = ioutil.NopCloser(bytes.NewBuffer(postBodyBuf))

		break
	case types.Object:
		returnResponse = http.Response{
			StatusCode: 200,
			Header:     nil,
		}
		postBodyBuf, _ := json.Marshal(response.(types.Object))
		returnResponse.Body = ioutil.NopCloser(bytes.NewBuffer(postBodyBuf))
		break
	case string:
		returnResponse = http.Response{
			StatusCode: 200,
			Header:     nil,
		}
		postBodyBuf, _ := json.Marshal(response.(string))
		returnResponse.Body = ioutil.NopCloser(bytes.NewBuffer(postBodyBuf))
		break
	default:
		return nil, errors.New(fmt.Sprintf("unsupported response type '%s'", reflect.TypeOf(response).Kind().String()))
	}

	return responseCastable.castResponseToType(&returnResponse, toType)
}
