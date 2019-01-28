package insertion

import (
	"reflect"
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name string
		data []int
	}{
		{"empty", []int{}},
		{"pair", []int{3, 2}},
		{"three", []int{3, 2, 1}},
		{"almost sorted", []int{1, 3, 2, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := make([]int, len(tt.data))
			copy(want, tt.data)
			sort.Ints(want)
			Sort(tt.data)
			if !reflect.DeepEqual(tt.data, want) {
				t.Errorf("Sort() = %v, want %v", tt.data, want)
			}
		})
	}
}

func Test_insert(t *testing.T) {
	type args struct {
		data  []int
		index int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"single", args{[]int{1}, 0}, []int{1}},
		{"pair", args{[]int{2, 1}, 1}, []int{1, 2}},
		{"three", args{[]int{3, 2, 1}, 2}, []int{1, 3, 2}},
		{"sorted", args{[]int{1, 2, 3}, 2}, []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			insert(tt.args.data, tt.args.index)
			if !reflect.DeepEqual(tt.args.data, tt.want) {
				t.Errorf("insert() = %v, want %v", tt.args.data, tt.want)
			}

		})
	}
}
