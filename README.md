[![GoDoc](https://godoc.org/github.com/cheekybits/m?status.png)](http://godoc.org/github.com/cheekybits/m)

# m
Map utilities for Go

### `Get` and `Set` using JavaScript notation

Before:

```
city := data.(map[string]interface{})["addresses"].([]interface{})[0].(map[string]interface{})["city"]
```

After:

```
city := m.Get(data, "addresses[0].city")
```