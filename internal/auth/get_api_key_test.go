package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {

	type output struct {
		key    string
		hasErr bool
	}

	type test struct {
		input http.Header
		want  output
	}

	tests := []test{
		{input: http.Header{}, want: output{"", true}},
		{input: http.Header{"Authorization": {""}}, want: output{"", true}},
		{input: http.Header{"Authorization": {"   "}}, want: output{"", true}},
		{input: http.Header{"Authorization": {"   aaa"}}, want: output{"", true}},
		{input: http.Header{"Authorization": {"   aaa bbb"}}, want: output{"", true}},
		{input: http.Header{"Authorization": {"aaa bbb ccc"}}, want: output{"", true}},
		{input: http.Header{"Authorization": {"aaa bbb"}}, want: output{"", true}},
		{input: http.Header{"Authorization": {"ApiKey bbb"}}, want: output{"bbb", false}},
	}

	for _, test := range tests {
		key, err := GetAPIKey(test.input)
		if (err != nil) != test.want.hasErr {
			t.Fatalf("unexpected error: %v", err)
		}
		if key != test.want.key {
			t.Fatalf("got: %v, want: %v", key, test.want.key)
		}
	}

}
