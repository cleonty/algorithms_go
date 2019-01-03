package findpair

import "testing"

func Test_findPair(t *testing.T) {
	type args struct {
		a     []int
		b     []int
		value int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"sorted lists", args{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}, 5}, true},
		{"unsorted lists", args{[]int{7, 1, 2, 3, 4, 5}, []int{8, 1, 2, 3, 4, 5}, 15}, true},
		{"not found", args{[]int{7, 1, 2, 3, 4, 5}, []int{8, 1, 2, 3, 4, 5}, 25}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPair(tt.args.a, tt.args.b, tt.args.value); got != tt.want {
				t.Errorf("findPair() = %v, want %v", got, tt.want)
			}
		})
	}
}
