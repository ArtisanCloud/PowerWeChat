package power

import (
	"encoding/json"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
)

type HashMap object.HashMap

func MergeHashMap(toMap *HashMap, subMaps ...*HashMap) *HashMap {
	if toMap == nil {
		toMap = &HashMap{}
	}
	// 拍平subMaps
	for _, subMap := range subMaps {
		if subMap != nil {
			// 迭代每个HashMap
			for k, v := range *subMap {
				toV := (*toMap)[k]

				// if the key is not exist in toMap
				if toV == nil {
					(*toMap)[k] = v
					continue
				}

				// if the toMap by the key is ""
				switch toV.(type) {
				case string:
					if (*toMap)[k] == "" && v != "" {
						(*toMap)[k] = v
					}
					break
				}

			}
		}
	}
	return toMap
}

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

func StructToHashMap(obj interface{}) (newMap *HashMap, err error) {
	data, err := json.Marshal(obj) // Convert to a json string

	if err != nil {
		return
	}

	newMap = &HashMap{}
	err = json.Unmarshal(data, newMap) // Convert to a map
	return
}
