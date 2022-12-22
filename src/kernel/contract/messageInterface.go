package contract

import (
	"github.com/ArtisanCloud/PowerLibs/v3/object"
)

type MessageInterface interface {
	GetType() string
	TransformForJsonRequest(appends *object.HashMap, withType bool) (*object.HashMap, error)

	// default return string
	TransformToXml(appends *object.HashMap, returnAsArray bool) (interface{}, error)
}
