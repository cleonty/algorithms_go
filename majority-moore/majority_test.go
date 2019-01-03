package majority

import "testing"

func Test_getMajority(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 bool
	}{
		{"all 1s", args{[]int{1, 1, 1, 1, 1, 1}}, 1, true},
		{"1s in the middle", args{[]int{2, 1, 1, 1, 1, 2}}, 1, true},
		{"half of array are ones", args{[]int{1, 2, 1, 2, 1, 2, 1, 2}}, 0, false},
		{"half of array + 1 are ones", args{[]int{1, 2, 1, 2, 1, 2, 1, 2, 1}}, 1, true},
		{"empty array", args{[]int{}}, 0, false},
		{"single element", args{[]int{1}}, 1, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := getMajority(tt.args.data)
			if got != tt.want {
				t.Errorf("getMajority() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getMajority() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
