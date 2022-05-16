package power

import (
	"github.com/ArtisanCloud/PowerLibs/object"
)

type HashMap object.HashMap

func (obj *HashMap) ToHashMap() *object.HashMap {

	hObj, _ := object.StructToHashMap(obj)

	return hObj
}

// ------------------------------- Conversion ---------------------------------------
func HashMapToPower(obj *object.HashMap) (newMap *HashMap, err error) {

	newMap = &HashMap{}

	if obj == nil {
		return newMap, err
	}

	for k, v := range *obj {
		(*newMap)[k] = v
	}

	return newMap, err

}


func PowerHashMapToObjectHashMap(obj *HashMap) (newMap *object.HashMap, err error) {

	newMap = &object.HashMap{}

	if obj == nil {
		return newMap, err
	}

	for k, v := range *obj {
		(*newMap)[k] = v
	}

	return newMap, err

}


