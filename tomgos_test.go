package tomgos

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeyToCamelCase(t *testing.T) {
	testcases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "simple key",
			input:    "key",
			expected: "Key",
		},
		{
			name:     "simple key with dash (-)",
			input:    "this-key",
			expected: "ThisKey",
		},
		{
			name:     "simple key with underscore (_)",
			input:    "this_key",
			expected: "ThisKey",
		},
		{
			name:     "simple key with number (0-9)",
			input:    "this-is-007",
			expected: "ThisIs007",
		},
		{
			name:     "simple key with uppercase (A-Z)",
			input:    "wHatEvEr",
			expected: "Whatever",
		},
		{
			name:     "dash first key",
			input:    "-key",
			expected: "Key",
		},
		{
			name:     "underscore first key",
			input:    "_key",
			expected: "Key",
		},
		{
			name:     "number first key",
			input:    "9key",
			expected: "9key",
		},
		{
			name:     "complex key combination",
			input:    "-this-is_kINda-r4Nd0m-KEY",
			expected: "ThisIsKindaR4nd0mKey",
		},
	}

	assert := assert.New(t)
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			output := keyToCamelCase(tc.input)

			assert.Equal(tc.expected, output)
		})
	}
}
