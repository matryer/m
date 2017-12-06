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
func Get(m interface{}, keypath string) interface{} {
	value, _ := GetOK(m, keypath)
	return value
}

// GetOK gets the value from the map by the given dot-notation
// keypath, or returns the second argument false if any of the data
// is missing.
// For more information on supported keypaths, see Get.
func GetOK(m interface{}, keypath string) (interface{}, bool) {
	return getOK(m, strings.Split(keypath, dot))
}

// getOK gets the value specified by the keys array from the
// map.
func getOK(m interface{}, keys []string) (interface{}, bool) {
	k := keys[0]
	if len(keys) > 1 {
		var sub interface{}
		var ok bool
		if sub, ok = get(m, k); !ok {
			return nil, false
		}

		switch sub.(type) {
		case map[string]interface{}, []map[string]interface{}:
			return getOK(sub, keys[1:])
		}
		return nil, false
	}
	value, ok := get(m, k)
	if value == nil {
		return nil, false
	}
	return value, ok
}

func parseArrayPath(k string) ([]string, int) {
	segs := strings.Split(k, openingBracket)
	if len(segs) == 1 {
		return segs, -1
	}

	i, err := strconv.ParseInt(segs[1][0:len(segs[1])-1], 10, 64)
	if err != nil {
		return segs, -1
	}

	return segs, int(i)
}

// get gets the key from the map.
// Supports array notation for slices.
func get(m interface{}, k string) (interface{}, bool) {
	if m == nil {
		return nil, false
	}

	if k == "" {
		return m, true
	}

	if k[len(k)-1] == closingBracket {
		segs, i := parseArrayPath(k)
		if i == -1 {
			return nil, false
		}
		sub, ok := get(m, segs[0])
		if !ok {
			return nil, false
		}
		val := reflect.ValueOf(sub)
		if !val.IsValid() {
			return nil, false
		}
		if val.Len() <= i || (val.Kind() != reflect.Slice && val.Kind() != reflect.Array) {
			return nil, false
		}
		return val.Index(int(i)).Interface(), true
	}

	mapType := reflect.TypeOf(map[string]interface{}(nil))
	if reflect.TypeOf(m).ConvertibleTo(mapType) {
		val := reflect.ValueOf(m).Convert(mapType)
		v, ok := val.Interface().(map[string]interface{})[k]
		return v, ok
	}

	return nil, false
}
