package m

import (
	"reflect"
	"strings"
)

// Set sets the value in the map at the given keypath.
// For more information on supported keypaths, see Get.
// If an expected object is missing, a map[string]interface{}
// will be created and assigned to the appropriate key.
func Set(m interface{}, keypath string, value interface{}) {
	setOK(m, strings.Split(keypath, dot), value)
}

// SetOK sets the value from the map by the given dot-notation
// keypath, or returns false if any of the data is missing.
// For more information on supported keypaths, see Get.
func SetOK(m interface{}, keypath string, value interface{}) bool {
	return setOK(m, strings.Split(keypath, dot), value)
}

func setOK(m interface{}, keys []string, value interface{}) bool {
	if m == nil {
		return false
	}
	k := keys[0]
	if len(keys) > 1 {
		var sub interface{}
		var ok bool
		if sub, ok = get(m, k); !ok {
			// Make sure it is an array index
			nextKey := keys[1]
			if nextKey[len(nextKey)-1] == closingBracket {
				sub = make([]map[string]interface{}, 0)
			} else {
				sub = make(map[string]interface{})
			}
			if !set(m, k, sub) {
				return false
			}
		}
		switch sub.(type) {
		case map[string]interface{}, []map[string]interface{}:
			return setOK(sub, keys[1:], value)
		}

		return false
	}
	return set(m, k, value)
}

// set sets the value to the key.
// Supports array setting: arr[2]=val.
func set(m interface{}, k string, value interface{}) bool {

	if k[len(k)-1] == closingBracket {
		segs, i := parseArrayPath(k)
		if i == -1 {
			return false
		}
		sub, ok := get(m, segs[0])
		if !ok {
			return false
		}

		// Grow array by 1, and don't panic
		val := reflect.ValueOf(sub)
		if val.Len() <= i {
			return false
		}

		val.Index(i).Set(reflect.ValueOf(value))
		return true
	}

	mapType := reflect.TypeOf(map[string]interface{}(nil))
	if reflect.TypeOf(m).ConvertibleTo(mapType) {
		val := reflect.ValueOf(m).Convert(mapType)
		val.Interface().(map[string]interface{})[k] = value
		return true
	}

	return false
}
