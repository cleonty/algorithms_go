package countingsort

import (
	"sort"
	"testing"
)

func TestSortAges(t *testing.T) {
	type args struct {
		ages []int
	}
	tests := []struct {
		name string
		ages []int
	}{
		{"sorted", []int{1, 2, 3, 4, 5}},
		{"reversed", []int{5, 4, 3, 2, 1}},
		{"equal", []int{1, 2, 1, 2, 1, 4, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SortAges(tt.ages)
			if !sort.IntsAreSorted(tt.ages) {
				t.Errorf("SortAges() didn't sort ages %v", tt.ages)
			}
		})
	}
}
