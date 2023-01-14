package keyboardrow

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findWords(t *testing.T) {

	testCases := []struct {
		name     string
		words    []string
		expected []string
	}{
		{
			name:     "Test #1",
			words:    []string{"Hello", "Alaska", "Dad", "Peace"},
			expected: []string{"Alaska", "Dad"},
		},
		{
			name:     "Test #2",
			words:    []string{"omk"},
			expected: []string{},
		},
		{
			name:     "Test #3",
			words:    []string{"adsdf", "sfd"},
			expected: []string{"adsdf", "sfd"},
		},
		{
			name:     "Test #4",
			words:    []string{"afghjkll"},
			expected: []string{"afghjkll"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := findWords(tc.words)

			require.Equalf(t, tc.expected, got, "expected: %v, got: %v", tc.expected, got)

		})
	}

}
