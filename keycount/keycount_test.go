package keycount

import (
	"testing"
)

func Test_keyCountSlow(t *testing.T) {
	type args struct {
		data []int
		key  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"zero", args{[]int{0, 1, 2, 3, 4, 5}, -1}, 0},
		{"one", args{[]int{0, 1, 2, 3, 4, 5}, 1}, 1},
		{"five", args{[]int{1, 1, 1, 1, 1, 5}, 1}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := keyCountSlow(tt.args.data, tt.args.key); got != tt.want {
				t.Errorf("keyCountSlow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_keyCountFast(t *testing.T) {
	type args struct {
		data []int
		key  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"zero", args{[]int{0, 1, 2, 3, 4, 5}, -1}, 0},
		{"one", args{[]int{0, 1, 2, 3, 4, 5}, 1}, 1},
		{"five", args{[]int{1, 1, 1, 1, 1, 5}, 1}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := keyCountFast(tt.args.data, tt.args.key); got != tt.want {
				t.Errorf("keyCountFast() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findFirstIndex(t *testing.T) {
	type args struct {
		data  []int
		key   int
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"zero", args{[]int{0, 1, 2, 3, 4, 5}, 1, 0, 6}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findFirstIndex(tt.args.data, tt.args.key, tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("findFirstIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
