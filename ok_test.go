package m_test

import (
	"testing"

	"github.com/cheekybits/is"
	"github.com/cheekybits/m"
)

func TestHas(t *testing.T) {
	is := is.New(t)

	d := map[string]interface{}{"name": "Mat", "book": "Go Programming Blueprints"}

	is.Equal(true, m.OK(d, "name"))
	is.Equal(true, m.OK(d, "book"))
	is.Equal(false, m.OK(d, "sausages"))

}
