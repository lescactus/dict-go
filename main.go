package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
)

var (
	// Instantiate cli flags
	lang = flag.String("lang", "en", "Lang of the word")
	word = flag.String("word", "", "Word to lookup")
)

const (
	baseDictAPI = "https://api.dictionaryapi.dev/api/v2/entries"
)

func main() {
	flag.Parse()

	if *word == "" {
		fmt.Fprintln(os.Stderr, "You must specify a word to lookup.")
		os.Exit(1)
	}

	URL := buildUrl(*lang, *word)
	resp, err := http.Get(URL)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if resp.StatusCode != 200 {
		fmt.Fprintln(os.Stderr, "Error while fetching word definition.")
		os.Exit(1)
	}

	//fmt.Printf("%s", string(body))

	var e Entry

	err = json.NewDecoder(resp.Body).Decode(&e)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Printf("→ word: %s\n", e[0].Word)
	fmt.Printf("→ lang: %s\n\n", *lang)
	for _, m := range e[0].Meanings {
		partOfSpeech := m.PartOfSpeech

		for _, d := range m.Definitions {
			fmt.Printf("• %s: %s\n\n", partOfSpeech, d.Definition)
		}
	}
	
}

func buildUrl(lang, word string) string {
	return fmt.Sprintf("%s/%s/%s", baseDictAPI, lang, word)
}
