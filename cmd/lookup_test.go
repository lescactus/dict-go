package cmd

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHTTPClient(t *testing.T) {
	tests := []struct {
		desc string
		want *HTTPClient
		have [2]string
	}{
		{
			desc: "Valid entry",
			have: [2]string{"en", "hello"},
			want: &HTTPClient{
				URL:    "https://api.dictionaryapi.dev/api/v2/entries/en/hello",
				Client: &http.Client{},
			},
		},
		{
			desc: "Valid entry",
			have: [2]string{"fr", "bonjour"},
			want: &HTTPClient{
				URL:    "https://api.dictionaryapi.dev/api/v2/entries/fr/bonjour",
				Client: &http.Client{},
			},
		},
		{
			desc: "Missing word",
			have: [2]string{"en", ""},
			want: &HTTPClient{
				URL:    "https://api.dictionaryapi.dev/api/v2/entries/en/",
				Client: &http.Client{},
			},
		},
		{
			desc: "Missing lang",
			have: [2]string{"", "hello"},
			want: &HTTPClient{
				URL:    "https://api.dictionaryapi.dev/api/v2/entries//hello",
				Client: &http.Client{},
			},
		},
		{
			desc: "Missing lang and word",
			have: [2]string{"", ""},
			want: &HTTPClient{
				URL:    "https://api.dictionaryapi.dev/api/v2/entries//",
				Client: &http.Client{},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			c := NewHTTPClient(test.have[0], test.have[1])

			assert.Equal(t, test.want, c)
			assert.NotEmpty(t, c)
		})
	}
}

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

func TestFetchWordDefinition(t *testing.T) {
	expected := []byte(`[{"word":"hello","phonetics":[{"audio":"https://api.dictionaryapi.dev/media/pronunciations/en/hello-au.mp3","sourceUrl":"https://commons.wikimedia.org/w/index.php?curid=75797336","license":{"name":"BY-SA 4.0","url":"https://creativecommons.org/licenses/by-sa/4.0"}},{"text":"/həˈləʊ/","audio":"https://api.dictionaryapi.dev/media/pronunciations/en/hello-uk.mp3","sourceUrl":"https://commons.wikimedia.org/w/index.php?curid=9021983","license":{"name":"BY 3.0 US","url":"https://creativecommons.org/licenses/by/3.0/us"}},{"text":"/həˈloʊ/","audio":""}],"meanings":[{"partOfSpeech":"noun","definitions":[{"definition":"\"Hello!\" or an equivalent greeting.","synonyms":[],"antonyms":[]}],"synonyms":["greeting"],"antonyms":[]},{"partOfSpeech":"verb","definitions":[{"definition":"To greet with \"hello\".","synonyms":[],"antonyms":[]}],"synonyms":[],"antonyms":[]},{"partOfSpeech":"interjection","definitions":[{"definition":"A greeting (salutation) said when meeting someone or acknowledging someone’s arrival or presence.","synonyms":[],"antonyms":[],"example":"Hello, everyone."},{"definition":"A greeting used when answering the telephone.","synonyms":[],"antonyms":[],"example":"Hello? How may I help you?"},{"definition":"A call for response if it is not clear if anyone is present or listening, or if a telephone conversation may have been disconnected.","synonyms":[],"antonyms":[],"example":"Hello? Is anyone there?"},{"definition":"Used sarcastically to imply that the person addressed or referred to has done something the speaker or writer considers to be foolish.","synonyms":[],"antonyms":[],"example":"You just tried to start your car with your cell phone. Hello?"},{"definition":"An expression of puzzlement or discovery.","synonyms":[],"antonyms":[],"example":"Hello! What’s going on here?"}],"synonyms":[],"antonyms":["bye","goodbye"]}],"license":{"name":"CC BY-SA 3.0","url":"https://creativecommons.org/licenses/by-sa/3.0"},"sourceUrls":["https://en.wiktionary.org/wiki/hello"]}]`)

	// Start a local http server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			w.Write(expected)
		} else if r.URL.Path == "/404" {
			w.WriteHeader(404)
		}
	}))

	// Close the http server
	defer server.Close()

	// Use Client & URL from the local test server
	c := HTTPClient{
		Client: server.Client(),
		URL:    server.URL,
	}

	body, err := c.FetchWordDefinition()
	assert.NoError(t, err)
	assert.NotEmpty(t, body)
	assert.Equal(t, expected, body)

	// Test for non 200 return code
	c404 := HTTPClient{
		Client: server.Client(),
		URL:    server.URL + "/404",
	}
	_, err = c404.FetchWordDefinition()
	assert.Error(t, err)
}
