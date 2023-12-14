package leetcode

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
		output   []int
	}{
		{[]int{1, 1, 2, 3, 3, 4, 4, 5}, 5, []int{1, 2, 3, 4, 5}},
		{[]int{1, 1, 1, 2, 2, 3, 4, 4, 5}, 5, []int{1, 2, 3, 4, 5}},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			result := RemoveDuplicates(test.input)

			if result != test.expected {
				t.Errorf("Expected %d, but got %d", test.expected, result)
			}

			if !reflect.DeepEqual(test.input[:result], test.output) {
				t.Errorf("Expected %v, but got %v", test.output, test.input[:result])
			}
		})
	}
}
