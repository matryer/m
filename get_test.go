package m_test

import (
	"testing"

	"github.com/cheekybits/is"
	"github.com/matryer/m"
)

var getTests = []struct {
	M  interface{}
	K  string
	V  interface{}
	OK bool
}{
	{
		map[string]interface{}{"name": "Tylor"},
		"name",
		"Tylor",
		true,
	},
	{
		map[string]interface{}{"name": nil},
		"name",
		nil,
		false,
	},
	{
		map[string]interface{}{"address": map[string]interface{}{"city": "London"}},
		"address.city",
		"London",
		true,
	},
	{
		map[string]interface{}{
			"address": map[string]interface{}{
				"postcode": map[string]interface{}{
					"inner": "NG19",
				},
			},
		},
		"address.postcode.inner",
		"NG19",
		true,
	},
	{
		nil,
		"address.postcode.inner",
		nil,
		false,
	},
	{
		map[string]interface{}{"address": map[string]interface{}{"city": "London"}},
		"address.country",
		nil,
		false,
	}, {
		map[string]interface{}{"places": []interface{}{map[string]interface{}{"city": "London"}, map[string]interface{}{"city": "San Francisco"}}},
		"places",
		[]interface{}{map[string]interface{}{"city": "London"}, map[string]interface{}{"city": "San Francisco"}},
		true,
	}, {
		map[string]interface{}{"places": []interface{}{map[string]interface{}{"city": "London"}, map[string]interface{}{"city": "San Francisco"}}},
		"places[0]",
		map[string]interface{}{"city": "London"},
		true,
	}, {
		map[string]interface{}{"places": []interface{}{map[string]interface{}{"city": "London"}, map[string]interface{}{"city": "San Francisco"}}},
		"places[1]",
		map[string]interface{}{"city": "San Francisco"},
		true,
	}, {
		map[string]interface{}{"places": []interface{}{map[string]interface{}{"city": "London"}, map[string]interface{}{"city": "San Francisco"}}},
		"places[2]",
		nil,
		false,
	}, {
		map[string]interface{}{"places": []interface{}{map[string]interface{}{"city": "London"}, map[string]interface{}{"city": "San Francisco"}}},
		"places[1].city",
		"San Francisco",
		true,
	}, {
		[]map[string]interface{}{{"city": "London"}, {"city": "San Francisco"}},
		"[1].city",
		"San Francisco",
		true,
	}, {
		[]map[string]interface{}{{"city": "London"}, map[string]interface{}{"city": "San Francisco"}},
		"places[1].city",
		nil,
		false,
	}, {
		[]map[string]interface{}{{"city": "London"}, map[string]interface{}{"city": "San Francisco"}},
		"1].city",
		nil,
		false,
	}, {
		[]map[string]interface{}{
			{
				"address": []map[string]interface{}{
					{
						"postcode": map[string]interface{}{
							"inner": "NG19",
						},
					},
				},
			},
		},
		"[0].address[0].postcode.inner",
		"NG19",
		true,
	},
	{
		map[string]interface{}{"name": "Tylor"},
		"name",
		"Tylor",
		true,
	},
	{
		map[string]interface{}{"items": nil},
		"items[0].something",
		nil,
		false,
	},
}

func TestGet(t *testing.T) {
	is := is.New(t)
	for _, test := range getTests {
		actual, ok := m.GetOK(test.M, test.K)
		is.Equal(actual, test.V)
		is.Equal(ok, test.OK)
	}
}
