package count_good_numbers

import "testing"

func Test_countGoodNumbers(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test #1",
			args: args{
				n: 1,
			},
			want: 5,
		},
		{
			name: "Test #2",
			args: args{
				n: 4,
			},
			want: 400,
		},
		{
			name: "Test #3",
			args: args{
				n: 50,
			},
			want: 564908303,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countGoodNumbers(tt.args.n); got != tt.want {
				t.Errorf("countGoodNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_countGoodNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countGoodNumbers(1_000_000_000_000_000)
	}
}
