package extract

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestExtract(t *testing.T) {
	buf := []byte("{\"foo\":\"bar\"}")
	value, err := Extract(buf, "foo")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, value, "bar")
}
