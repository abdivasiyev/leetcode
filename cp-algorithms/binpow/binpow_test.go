package binpow

import "testing"

func TestBinPowRec(t *testing.T) {
	tests := []struct {
		name string
		a    int64
		n    int64
		want int64
	}{
		{
			name: "Test with even degree",
			a:    2,
			n:    4,
			want: 16,
		},
		{
			name: "Test with odd degree",
			a:    2,
			n:    5,
			want: 32,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinPowRecursive(tt.a, tt.n); got != tt.want {
				t.Errorf("BinPowRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkBinPowRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BinPowRecursive(2, 100_000)
	}
}

func TestBinPowIterative(t *testing.T) {
	tests := []struct {
		name string
		a    int64
		n    int64
		want int64
	}{
		{
			name: "Test with even degree",
			a:    2,
			n:    4,
			want: 16,
		},
		{
			name: "Test with odd degree",
			a:    2,
			n:    5,
			want: 32,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinPowIterative(tt.a, tt.n); got != tt.want {
				t.Errorf("BinPowIterative() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkBinPowIterative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BinPowIterative(2, 100_000)
	}
}
