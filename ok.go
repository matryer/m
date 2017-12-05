package m

// OK gets whether the map has a value at the specified
// keypath.
// Nil values are considered not OK.
func OK(m interface{}, keypath string) bool {
	_, has := GetOK(m, keypath)
	return has
}
