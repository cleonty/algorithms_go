package maxprofit

import "testing"

func Test_maxProfit(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"one", args{[]int{0, 1, 1, 0}}, 1},
		{"two peaks", args{[]int{0, 1, 1, 0, 2, 3, 1, 0}}, 3},
		{"two peaks in middle", args{[]int{2, 0, 1, 1, 0, 2, 3, 1, 2}}, 3},
		{"empty list", args{[]int{}}, 0},
		{"negative", args{[]int{-1, -2, -3, -1}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.data); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
