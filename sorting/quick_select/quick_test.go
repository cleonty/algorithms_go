package quick

import (
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name  string
		data  []int
		index int
		elem  int
	}{
		{"pair", []int{3, 2}, 1, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := make([]int, len(tt.data))
			copy(want, tt.data)
			sort.Ints(want)
			elem := Select(tt.data, tt.index)
			if elem != want[tt.index] {
				t.Errorf("Select() = %v, want %v", elem, want[tt.index])
			}
		})
	}
}
