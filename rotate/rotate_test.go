package main

import (
	"reflect"
	"testing"
)

func Test_rotate(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"rotate by 0", args{[]int{0, 1, 2, 3, 4, 5}, 0}, []int{0, 1, 2, 3, 4, 5}},
		{"rotate by 2", args{[]int{0, 1, 2, 3, 4, 5}, 2}, []int{2, 3, 4, 5, 0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotate(tt.args.arr, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotate() = %v, want %v", got, tt.want)
			}
		})
	}
}
