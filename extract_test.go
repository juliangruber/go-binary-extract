package extract

import (
	"errors"
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

	value, err = Extract([]byte("{\"foo\":\"bar\\\"baz\"}"), "foo")
	check(t, err)
	assert.Equal(t, value, "bar\"baz")

	value, err = Extract([]byte("{\"_a\":0,\"a_\":1,\"_a_\":2,\"a\":3}"), "a")
	check(t, err)
	assert.Equal(t, value, float64(3))

	value, err = Extract([]byte("{\"foo\""), "foo")
	if err == nil {
		t.Error(errors.New("missing error"))
	}

	value, err = Extract([]byte("{\"foo\":\"bar\"}"), "bar")
	if err == nil {
		t.Error(errors.New("missing error"))
	}
}

func check(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}
