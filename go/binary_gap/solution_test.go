package binary_gap

import "testing"

func TestBinaryGap(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name        string
		description string
		n           int
		expected    int
	}{
		{
			name:        "example 1",
			description: "22 in binary is \"10110\".",
			n:           22,
			expected:    2,
		},
		{
			name:        "example 2",
			description: "8 in binary is \"1000\".",
			n:           8,
			expected:    0,
		},
		{
			name:        "example 3",
			description: "5 in binary is \"101\".",
			n:           5,
			expected:    2,
		},
		{
			name:        "example 4",
			description: "12 in binary is \"1100\".",
			n:           12,
			expected:    1,
		},
		{
			name:        "example 5",
			description: "902 in binary is \"1110000110\".",
			n:           902,
			expected:    5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := binaryGap(tc.n)

			if got != tc.expected {
				t.Fatalf("expected %v, got %v.\n%s", tc.expected, got, tc.description)
			}
		})
	}
}

func BenchmarkBinaryGap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		binaryGap(i)
	}
}
