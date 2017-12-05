package m_test

import (
	"testing"

	"github.com/cheekybits/is"
	"github.com/matryer/m"
)

func TestHas(t *testing.T) {
	is := is.New(t)

	d := map[string]interface{}{"name": "Mat", "book": "Go Programming Blueprints"}

	is.Equal(true, m.OK(d, "name"))
	is.Equal(true, m.OK(d, "book"))
	is.Equal(false, m.OK(d, "sausages"))

	da := []map[string]interface{}{map[string]interface{}{"name": "Mat", "book": "Go Programming Blueprints"}, map[string]interface{}{"name": "John", "book": "Go Programming", "sausages": "please"}}

	is.Equal(true, m.OK(da, "[0].name"))
	is.Equal(true, m.OK(da, "[0].book"))
	is.Equal(false, m.OK(da, "[0].sausages"))

	is.Equal(true, m.OK(da, "[1].name"))
	is.Equal(true, m.OK(da, "[1].book"))
	is.Equal(true, m.OK(da, "[1].sausages"))
}
