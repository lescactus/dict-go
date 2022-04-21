package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildUrl(t *testing.T) {
	tests := []struct {
		desc string
		want string
		have [2]string
	}{
		{
			desc: "Valid entry",
			want: "https://api.dictionaryapi.dev/api/v2/entries/en/hello",
			have: [2]string{"en", "hello"},
		},
		{
			desc: "Valid entry",
			want: "https://api.dictionaryapi.dev/api/v2/entries/fr/bonjour",
			have: [2]string{"fr", "bonjour"},
		},
		{
			desc: "Missing word",
			want: "https://api.dictionaryapi.dev/api/v2/entries/en/",
			have: [2]string{"en", ""},
		},
		{
			desc: "Missing lang",
			want: "https://api.dictionaryapi.dev/api/v2/entries//hello",
			have: [2]string{"", "hello"},
		},
		{
			desc: "Missing lang and word",
			want: "https://api.dictionaryapi.dev/api/v2/entries//",
			have: [2]string{"", ""},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			url := buildUrl(test.have[0], test.have[1])
			assert.Equal(t, test.want, url)
		})
	}
}