package find

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	text  = "BBC ABCDAB ABCDABCDABDE"
	word  = "ABCDABD"
	index = 15
)

func TestNative(t *testing.T) {
	assert.Equal(t, index, Native(text, word), "string.Native")
}

func TestRabinKarp(t *testing.T) {
	assert.Equal(t, index, RabinKarp(text, word), "string.RabinKarp")
}
