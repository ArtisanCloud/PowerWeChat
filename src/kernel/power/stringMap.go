package power

import (
	"encoding/json"
	"github.com/ArtisanCloud/go-libs/object"
)

type StringMap object.StringMap

func (obj *StringMap) ToStringMap() *object.StringMap {

	hObj, _ := object.StructToStringMap(obj)

	return hObj
}

func StringMapToPower(obj interface{}) (newMap *StringMap, err error) {
	data, err := json.Marshal(obj) // Convert to a json string

	if err != nil {
		return
	}

	newMap = &StringMap{}
	err = json.Unmarshal(data, newMap) // Convert to a string map
	return
}
