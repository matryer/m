package m

import (
	"reflect"
	"strconv"
	"strings"
)

// Set sets the value in the map at the given keypath.
// For more information on supported keypaths, see Get.
func Set(m map[string]interface{}, keypath string, value interface{}) {
	setOK(m, strings.Split(keypath, dot), value)
}

// SetOK sets the value from the map by the given dot-notation
// keypath, or returns false if any of the data is missing.
// For more information on supported keypaths, see Get.
func SetOK(m map[string]interface{}, keypath string, value interface{}) bool {
	return setOK(m, strings.Split(keypath, dot), value)
}

func setOK(m map[string]interface{}, keys []string, value interface{}) bool {
	k := keys[0]
	if len(keys) > 1 {
		var sub interface{}
		var ok bool
		if sub, ok = get(m, k); !ok {
			return false
		}
		var submap map[string]interface{}
		if submap, ok = sub.(map[string]interface{}); !ok {
			return false
		}
		return setOK(submap, keys[1:], value)
	}
	return set(m, k, value)
}

// set sets the value to the key.
// Supports array setting: arr[2]=val.
func set(m map[string]interface{}, k string, value interface{}) bool {

	if k[len(k)-1] == closingBracket {
		segs := strings.Split(k, openingBracket)
		i, err := strconv.ParseInt(segs[1][0:len(segs[1])-1], 10, 64)
		if err != nil {
			return false
		}
		sub, ok := get(m, segs[0])
		if !ok {
			return false
		}
		reflect.ValueOf(sub).Index(int(i)).Set(reflect.ValueOf(value))
		return true
	}

	m[k] = value
	return true
}
