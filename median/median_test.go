package median

import "testing"

func Test_median(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name        string
		args        args
		want        float64
		shouldPanic bool
	}{
		{"three elements", args{[]int{1, 2, 3}}, 2, false},
		{"two equal elements", args{[]int{1, 1}}, 1, false},
		{"two not equal elements", args{[]int{1, 2}}, 1.5, false},
		{"single element", args{[]int{1}}, 1, false},
		{"empty list", args{[]int{}}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); tt.shouldPanic && r == nil {
					t.Errorf("should have panicked")
				}
			}()
			if got := median(tt.args.data); got != tt.want {
				t.Errorf("median() = %v, want %v", got, tt.want)
			}
		})
	}
}
