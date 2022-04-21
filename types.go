package main

type Entry []struct {
	Word       string      `json:"word"`
	Phonetics  []Phonetics `json:"phonetics"`
	Meanings   []Meanings  `json:"meanings"`
	License    License     `json:"license"`
	SourceUrls []string    `json:"sourceUrls"`
}
type License struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Phonetics struct {
	Audio     string  `json:"audio"`
	SourceURL string  `json:"sourceUrl,omitempty"`
	License   License `json:"license,omitempty"`
	Text      string  `json:"text,omitempty"`
}
type Definitions struct {
	Definition string        `json:"definition"`
	Synonyms   []interface{} `json:"synonyms"`
	Antonyms   []interface{} `json:"antonyms"`
}
type Meanings struct {
	PartOfSpeech string        `json:"partOfSpeech"`
	Definitions  []Definitions `json:"definitions"`
	Synonyms     []string      `json:"synonyms"`
	Antonyms     []interface{} `json:"antonyms"`
}
