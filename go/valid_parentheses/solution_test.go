package validparentheses

import "testing"

func Test_isValid(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		s        string
		expected bool
	}{
		{
			name:     "Test #1",
			s:        "()",
			expected: true,
		},
		{
			name:     "Test #2",
			s:        "()[]{}",
			expected: true,
		},
		{
			name:     "Test #3",
			s:        "(]",
			expected: false,
		},
		{
			name:     "Test #4",
			s:        "(])",
			expected: false,
		},
		{
			name:     "Test #5",
			s:        "({([{{}()}])})",
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := isValid(tc.s)
			if got != tc.expected {
				t.Fatalf("sample: %s, expected: %v, got: %v\n", tc.s, tc.expected, got)
			}
		})
	}
}
