package power

import (
	"encoding/json"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
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

func PowerStringMapToObjectStringMap(obj *StringMap) (newMap *object.StringMap, err error) {

	newMap = &object.StringMap{}

	if obj == nil {
		return newMap, err
	}

	for k, v := range *obj {
		(*newMap)[k] = v
	}

	return newMap, err

}
