package findpair

import "testing"

func Test_findPair(t *testing.T) {
	type args struct {
		data  []int
		value int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"integers sequential", args{[]int{1, 2, 3, 4, 5}, 3}, true},
		{"middle", args{[]int{1, 2, 3, 4, 5}, 5}, true},
		{"not found big", args{[]int{1, 2, 3, 4, 5}, 20}, false},
		{"not found", args{[]int{1, 2, 3, 4, 5}, 10}, false},
		{"not found min", args{[]int{1, 2, 3, 4, 5}, 1}, false},
		{"not found too small", args{[]int{1, 2, 3, 4, 5}, 0}, false},
		{"empty slice", args{[]int{}, 0}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPair(tt.args.data, tt.args.value); got != tt.want {
				t.Errorf("findPair() = %v, want %v", got, tt.want)
			}
		})
	}
}
