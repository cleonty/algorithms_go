package quick

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
