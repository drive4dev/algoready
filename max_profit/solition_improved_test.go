package max_profit

import "testing"

func Test_maxProfitImp(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case 1",
			args: args{
				prices: []int{7, 1, 5, 3, 6, 4},
			},
			want: 5,
		},

		{
			name: "Case 2",
			args: args{
				prices: []int{7, 6, 4, 3, 1},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfitImp(tt.args.prices); got != tt.want {
				t.Errorf("maxProfitImp() = %v, want %v", got, tt.want)
			}
		})
	}
}
