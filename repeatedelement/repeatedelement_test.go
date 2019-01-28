package repeatedelement

import "testing"

func Test_findRepeatedElement(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"one", args{[]int{1, 2, 3, 1, 2}}, 1},
		{"not found", args{[]int{1, 2, 3, 4, 5}}, 0},
		{"single element", args{[]int{1}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRepeatedElement(tt.args.data); got != tt.want {
				t.Errorf("findRepeatedElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
