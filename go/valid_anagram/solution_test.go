package validanagram

import "testing"

func Test_isAnagram(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		s        string
		t        string
		expected bool
	}{
		{
			name:     "Test #1",
			s:        "anagram",
			t:        "nagaram",
			expected: true,
		},
		{
			name:     "Test #2",
			s:        "rat",
			t:        "car",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := isAnagram(tc.s, tc.t)

			if got != tc.expected {
				t.Fatalf("expected: %v, got: %v\n", tc.expected, got)
			}
		})
	}
}
