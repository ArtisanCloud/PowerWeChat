package support

import (
	"github.com/ArtisanCloud/go-libs/http/contract"
)

type ResponseCastable struct {
}

func (responseCastable *ResponseCastable) castResponseToType(response contract.ResponseContract, castType string) interface{}{

	body:=response.GetBody()
	//(*body).

	switch castType {
	case "collection":

	case "array":

	case "object":

	case "raw":

	default:
	}

	return nil
}

func (responseCastable *ResponseCastable) detectAndCastResponseToType() {

}
