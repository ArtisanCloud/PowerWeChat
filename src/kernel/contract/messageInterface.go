package contract

import "github.com/ArtisanCloud/go-libs/object"

type MessageInterface interface {

	GetType() string
	TransformForJsonRequest() *object.HashMap
	TransformToXml() (interface{},error)
}
