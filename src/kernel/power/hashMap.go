package power

import (
	"encoding/json"
	"github.com/ArtisanCloud/PowerLibs/object"
)

type HashMap object.HashMap

func (obj *HashMap) ToHashMap() *object.HashMap {

	hObj, _ := object.StructToHashMap(obj)

	return hObj
}

// ------------------------------- Conversion ---------------------------------------
func HashMapToPower(obj interface{}) (newMap *HashMap, err error) {
	data, err := json.Marshal(obj) // Convert to a json string

	if err != nil {
		return
	}

	newMap = &HashMap{}
	err = json.Unmarshal(data, newMap) // Convert to a map
	return
}
