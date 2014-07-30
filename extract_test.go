package extract

import (
	"encoding/json"
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

	value, err = Extract([]byte("{\"foo\":{\"bar\":{\"baz\":\"beep\"}}}"), "foo")
	check(t, err)
	assert.Equal(t, value, map[string]interface{}{
		"bar": map[string]interface{}{
			"baz": "beep",
		},
	})
}

var raw = []byte("{\"properties\":{\"selected\":\"2\",\"lastName\":\"\",\"username\":\"someone\",\"category\":\"Wedding Venues\",\"firstName\":\"\",\"product\":\"planner\",\"location\":\"\",\"platform\":\"ios\",\"email\":\"someone@yahoo.com\",\"member_id\":\"12312313123123\",\"filtered\":\"false\",\"viewed\":3},\"projectId\":\"foobarbaz\",\"userId\":\"123123123123123\",\"sessionId\":\"FF8D19D8-123123-449E-A0B9-2181C4886020\",\"requestId\":\"F3C49DEB-123123-4A54-BB72-D4BE591E4B29\",\"action\":\"Track\",\"event\":\"Vendor Category Viewed\",\"timestamp\":\"2014-04-23T20:55:19.000Z\",\"context\":{\"providers\":{\"Crittercism\":false,\"Amplitude\":false,\"Mixpanel\":false,\"Countly\":false,\"Localytics\":false,\"Google Analytics\":false,\"Flurry\":false,\"Tapstream\":false,\"Bugsnag\":false},\"appReleaseVersion\":\"2.3.1\",\"osVersion\":\"7.1\",\"os\":\"iPhone OS\",\"appVersion\":\"690\",\"screenHeight\":480,\"library-version\":\"0.10.3\",\"traits\":{\"lastName\":\"\",\"product\":\"planner\",\"member_id\":\"123123123123123\",\"firstName\":\"\",\"email\":\"someone@yahoo.com\",\"platform\":\"ios\",\"username\":\"someone\"},\"screenWidth\":320,\"deviceManufacturer\":\"Apple\",\"library\":\"analytics-ios\",\"idForAdvertiser\":\"1323232-A0ED-47AB-BE4F-274F2252E4B4\",\"deviceModel\":\"iPad3,4\"},\"requestTime\":\"2014-04-23T20:55:44.211Z\",\"version\":1,\"channel\":\"server\"}")

func BenchmarkExtract(b *testing.B) {
	for i := 0; i < b.N; i++ {
		value, err := Extract([]byte(raw), "projectId")
		if err != nil {
			b.Error(err)
		}
		use(value)
	}
}

func BenchmarkParse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var all interface{}
		err := json.Unmarshal(raw, &all)
		if err != nil {
			b.Error(err)
		}
		use(all.(map[string]interface{})["projectId"])
	}
}

func check(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}

func use(i interface{}) {
}
