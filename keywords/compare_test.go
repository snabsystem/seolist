package keywords_test

import (
	"github.com/stretchr/testify/assert"
	"seolist/keywords"
	"testing"
)

func TestFromFile(t *testing.T) {

	keys, err := keywords.FromFile("test.csv", nil)

	assert.NoError(t, err)
	assert.NotEqual(t, 2, len(keys))
	assert.Equal(t, 3, len(keys))
}
