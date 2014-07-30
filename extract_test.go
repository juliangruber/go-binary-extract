package extract

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestExtract(t *testing.T) {
	value, err := Extract([]byte("{\"foo\":\"bar\"}"), "foo")
	check(t, err)
	assert.Equal(t, value, "bar")

	value, err = Extract([]byte("{\"foo\":\"bar\",\"bar\":\"baz\"}"), "foo")
	check(t, err)
	assert.Equal(t, value, "bar")

	value, err = Extract([]byte("{\"foo\":\"bar\",\"bar\":\"baz\"}"), "bar")
	check(t, err)
	assert.Equal(t, value, "baz")

	value, err = Extract([]byte("{\"foo\":{\"beep\":\"boop\",\"bar\":\"oops\"},\"bar\":\"baz\"}"), "bar")
	check(t, err)
	assert.Equal(t, value, "baz")

	value, err = Extract([]byte("{\"foo\":[{\"bar\":\"oops\"}],\"bar\":\"baz\"}"), "bar")
	check(t, err)
	assert.Equal(t, value, "baz")

	value, err = Extract([]byte("{\"foo\":\",bar\",\"bar\":\"baz\"}"), "bar")
	check(t, err)
	assert.Equal(t, value, "baz")

	value, err = Extract([]byte("{\"foo\":{\"bar\":\"baz\"}}"), "foo")
	check(t, err)
	assert.Equal(t, value, map[string]interface{}{
		"bar": "baz",
	})

	value, err = Extract([]byte("{\"foo\":[\"bar\",\"baz\"]}"), "foo")
	check(t, err)
	assert.Equal(t, value, []interface{}{"bar", "baz"})

	value, err = Extract([]byte("{\"beep\":\"\\\"\",\"foo\":\"bar\"}"), "foo")
	check(t, err)
	assert.Equal(t, value, "bar")
}

func check(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}
