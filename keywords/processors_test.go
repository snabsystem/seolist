package keywords_test

import (
	"github.com/stretchr/testify/assert"
	"seolist/keywords"
	"testing"
)

func TestProcessor(t *testing.T) {

	tests := []struct {
		name string
		line string
		want []string
	}{
		{"skip one-chars words", "o", []string{}},
		{"empty", "", []string{}},
		{"one word", "one", []string{"one"}},
		{"two words", "one two", []string{"one", "two"}},
		{"multiple spaces between", "one         two", []string{"one", "two"}},
		{"skip one-chars words", "one o", []string{"one"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, keywords.Processor(tt.line), "Processor(%v)", tt.line)
		})
	}
}

func BenchmarkProcessor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		keywords.Processor("one two three")
	}
}
