package keywords

import "strings"

func Processor(line string) []string {

	words := make([]string, 0)

	for _, word := range strings.Fields(line) {
		if len(word) > 1 {
			words = append(words, word)
		}
	}

	return words
}

func Comparator() {

}
