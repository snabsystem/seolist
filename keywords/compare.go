package keywords

import (
	"bufio"
	"os"
	"strings"
)

type FieldsFunc func(string) []string

// FromFile loads a CSV file and returns a slice of strings
// each string is a keyword, words less than 2 characters are ignored
func FromFile(filename string, fn FieldsFunc) ([][]string, error) {

	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	if fn == nil {
		fn = strings.Fields
	}

	keys := make([][]string, 0)

	for line := range readLines(f) {

		words := fn(line)

		if len(words) > 0 {
			keys = append(keys, words)
		}
	}

	return keys, nil
}

func readLines(f *os.File) <-chan string {

	ch := make(chan string)

	go func() {
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			ch <- scanner.Text()
		}
		close(ch)
	}()

	return ch
}
