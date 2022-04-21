package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type HTTPClient struct {
	URL string

	Client *http.Client
}

func NewHTTPClient(lang, word string) *HTTPClient {
	return &HTTPClient{
		URL:    buildUrl(lang, word),
		Client: &http.Client{},
	}
}

func (c *HTTPClient) FetchWordDefinition() ([]byte, error) {
	resp, err := c.Client.Get(c.URL)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Error while fetching word definition.")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	return body, err
}

func buildUrl(lang, word string) string {
	return fmt.Sprintf("%s/%s/%s", baseDictAPI, lang, word)
}
