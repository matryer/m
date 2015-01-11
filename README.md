[![GoDoc](https://godoc.org/github.com/cheekybits/m?status.png)](http://godoc.org/github.com/cheekybits/m)

# m
Map utilities for Go

### `Get` and `Set` using JavaScript notation

Before (yeah it scrolls):

```
city := data.(map[string]interface{})["addresses"].([]interface{})[0].(map[string]interface{})["city"]
```

  * Panics if any piece is missing or wrong type

After:

```
city := m.Get(data, "addresses[0].city")
```

  * Returns nil if any piece is missing or wrong type
