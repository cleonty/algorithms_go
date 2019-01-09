package separate

import (
	"reflect"
	"testing"
)

func Test_separate(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want []int
	}{
		{"one element", []int{1}, []int{1}},
		{"both even", []int{2, 4}, []int{2, 4}},
		{"both odd", []int{1, 3}, []int{1, 3}},
		{"even odd even", []int{2, 3, 4}, []int{2, 4, 3}},
		{"odd odd even", []int{1, 3, 4}, []int{4, 3, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			separate(tt.data)
			if !reflect.DeepEqual(tt.data, tt.want) {
				t.Errorf("separate() = %v, want %v", tt.data, tt.want)
			}
		})
	}
}
