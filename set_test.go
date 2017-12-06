package m_test

import (
	"testing"

	"github.com/cheekybits/is"
	"github.com/matryer/m"
)

var setTests = []struct {
	M  interface{}
	K  string
	V  interface{}
	X  interface{}
	OK bool
}{
	{
		map[string]interface{}{"name": "Tylor"},
		"name",
		"Ryon",
		map[string]interface{}{"name": "Ryon"},
		true,
	},
	{
		map[string]interface{}{},
		"name",
		"added",
		map[string]interface{}{"name": "added"},
		true,
	},
	{
		map[string]interface{}{"address": map[string]interface{}{"city": "London"}},
		"address.city",
		"San Francisco",
		map[string]interface{}{"address": map[string]interface{}{"city": "San Francisco"}},
		true,
	},
	{
		nil,
		"address.postcode.inner",
		nil,
		nil,
		false,
	},
	{
		map[string]interface{}{"address": map[string]interface{}{"city": "London"}},
		"address.country",
		"UK",
		map[string]interface{}{"address": map[string]interface{}{"city": "London", "country": "UK"}},
		true,
	}, {
		map[string]interface{}{"places": []interface{}{map[string]interface{}{"city": "London"}, map[string]interface{}{"city": "San Francisco"}}},
		"places",
		[]interface{}{map[string]interface{}{"city": "London Town"}, map[string]interface{}{"city": "San Francisco Town"}},
		map[string]interface{}{"places": []interface{}{map[string]interface{}{"city": "London Town"}, map[string]interface{}{"city": "San Francisco Town"}}},
		true,
	}, {
		map[string]interface{}{"places": []interface{}{map[string]interface{}{"city": "London"}, map[string]interface{}{"city": "San Francisco"}}},
		"places[0]",
		map[string]interface{}{"city": "Mansfield"},
		map[string]interface{}{"places": []interface{}{map[string]interface{}{"city": "Mansfield"}, map[string]interface{}{"city": "San Francisco"}}},
		true,
	}, {
		map[string]interface{}{"places": []interface{}{map[string]interface{}{"city": "London"}, map[string]interface{}{"city": "San Francisco"}}},
		"places[1]",
		map[string]interface{}{"city": "San Francisco Town"},
		map[string]interface{}{"places": []interface{}{map[string]interface{}{"city": "London"}, map[string]interface{}{"city": "San Francisco Town"}}},
		true,
	}, {
		map[string]interface{}{"places": []interface{}{map[string]interface{}{"city": "London"}, map[string]interface{}{"city": "San Francisco"}}},
		"places[1].city",
		"San Francisco Town",
		map[string]interface{}{"places": []interface{}{map[string]interface{}{"city": "London"}, map[string]interface{}{"city": "San Francisco Town"}}},
		true,
	}, {
		map[string]interface{}{},
		"person.name",
		"David",
		map[string]interface{}{"person": map[string]interface{}{"name": "David"}},
		true,
	},
	{
		[]interface{}{
			map[string]interface{}{
				"city": "London",
			},
			map[string]interface{}{
				"city": "San Francisco",
			},
		},
		"[1].city",
		"San Francisco Town",
		[]interface{}{
			map[string]interface{}{
				"city": "London",
			},
			map[string]interface{}{
				"city": "San Francisco Town",
			},
		},
		true,
	}, {
		nil,
		"[0].postcode.inner",
		nil,
		nil,
		false,
	}, {
		map[string]interface{}{"name": "Tylor"},
		"name",
		"Ryon",
		map[string]interface{}{"name": "Ryon"},
		true,
	},
}

func TestSet(t *testing.T) {
	is := is.New(t)
	for _, test := range setTests {
		//log.Println(test)
		ok := m.SetOK(test.M, test.K, test.V)
		if test.X == nil {
			is.Nil(test.M)
		} else {
			is.Equal(test.M, test.X)
		}
		is.Equal(test.OK, ok)
	}
}
