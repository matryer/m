package m

import (
	"reflect"
	"strconv"
	"strings"
)

const (
	dot            = "."
	openingBracket = "["
)

var closingBracket = "]"[0]

// Get gets the value from the map by the given dot-notation
// keypath. Returns nil if any of the value is missing.
// Supported keypaths:
//     key
//     key.subkey
//     key[i].subkey
//     key[i].subkey.subkey2.forever...
func Get(m map[string]interface{}, keypath string) interface{} {
	value, _ := GetOK(m, keypath)
	return value
}

// GetOK gets the value from the map by the given dot-notation
// keypath, or returns the second argument false if any of the data
// is missing.
// For more information on supported keypaths, see Get.
func GetOK(m map[string]interface{}, keypath string) (interface{}, bool) {
	return getOK(m, strings.Split(keypath, dot))
}

// getOK gets the value specified by the keys array from the
// map.
func getOK(m map[string]interface{}, keys []string) (interface{}, bool) {
	k := keys[0]
	if len(keys) > 1 {
		var sub interface{}
		var ok bool
		if sub, ok = get(m, k); !ok {
			return nil, false
		}
		var submap map[string]interface{}
		if submap, ok = sub.(map[string]interface{}); !ok {
			return nil, false
		}
		return getOK(submap, keys[1:])
	}
	value, ok := get(m, k)
	if value == nil {
		return nil, false
	}
	return value, ok
}

// get gets the key from the map.
// Supports array notation for slices.
func get(m map[string]interface{}, k string) (interface{}, bool) {
	if k[len(k)-1] == closingBracket {
		segs := strings.Split(k, openingBracket)
		i, err := strconv.ParseInt(segs[1][0:len(segs[1])-1], 10, 64)
		if err != nil {
			return nil, false
		}
		sub, ok := get(m, segs[0])
		if !ok {
			return nil, false
		}
		val := reflect.ValueOf(sub)
		if val.Len() <= int(i) {
			return nil, false
		}
		return val.Index(int(i)).Interface(), true
	}
	v, ok := m[k]
	return v, ok
}
