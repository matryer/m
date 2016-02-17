[![GoDoc](https://godoc.org/github.com/matryer/m?status.png)](http://godoc.org/github.com/matryer/m)

# m
Map utilities for Go

### `Get` and `Set` using JavaScript notation

Before (yeah it scrolls):

```
// set the city
data.(map[string]interface{})["addresses"].([]interface{})[0].(map[string]interface{})["city"] = "London"
// get the city
city := data.(map[string]interface{})["addresses"].([]interface{})[0].(map[string]interface{})["city"]
```

  * Panics if any piece is missing or wrong type

After:

```
m.Set(data, "addresses[0].city", "London")
city := m.Get(data, "addresses[0].city")
```

  * Returns nil if any piece is missing or wrong type
  * Use [`GetOK`](http://godoc.org/github.com/matryer/m#GetOK) for a second argument bool
