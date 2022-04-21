package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const (
	baseDictAPI = "https://api.dictionaryapi.dev/api/v2/entries"
)

var (
	lang string
	word string

	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:   "dict-go",
		Short: "Simple cli used to lookup for the definition of a given word",
		Long: `dict-go is a simple cli used to lookup for the definition of a given word. 
	You need to provide the word you are looking for and the language (optional - default is "en").

	The source code is available at https://github.com/lescactus/dict-go.`,
		Run: func(cmd *cobra.Command, args []string) {
			client := NewHTTPClient(lang, word)

			def, err := client.FetchWordDefinition()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			var e Entry

			err = json.Unmarshal(def, &e)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			fmt.Printf("→ word: %s\n", e[0].Word)
			fmt.Printf("→ lang: %s\n\n", lang)
			for _, m := range e[0].Meanings {
				partOfSpeech := m.PartOfSpeech

				for _, d := range m.Definitions {
					fmt.Printf("• %s: %s\n\n", partOfSpeech, d.Definition)
				}
			}
		},
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&lang, "lang", "l", "en", "Lang of the word (optional)")
	rootCmd.PersistentFlags().StringVarP(&word, "word", "w", "", "Word to lookup")
	rootCmd.MarkPersistentFlagRequired("word")
}
