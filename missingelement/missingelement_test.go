package missingelement

import "testing"

func Test_findMissing(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"one missing", args{[]int{2, 3, 4}}, 1},
		{"four missing", args{[]int{1, 2, 3}}, 4},
		{"two missing", args{[]int{1, 3, 4}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMissing(tt.args.data); got != tt.want {
				t.Errorf("findMissing() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMissing2(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"one missing", args{[]int{2, 3, 4}}, 1},
		{"four missing", args{[]int{1, 2, 3}}, 4},
		{"two missing", args{[]int{1, 3, 4}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMissing2(tt.args.data); got != tt.want {
				t.Errorf("findMissing2() = %v, want %v", got, tt.want)
			}
		})
	}
}
