package sortmultiple

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type multisorter func(...[]int) []int

var functions = map[string]multisorter{
	"SortArrayBigSort":              SortArrayBigSort,
	"SortArrayMap":                  SortArrayMap,
	"SortArrayMapPreAllocated":      SortArrayMapPreAllocated,
	"SortMapNoBigArray":             SortMapNoBigArray,
	"SortMapPreAllocatedNoBigArray": SortMapPreAllocatedNoBigArray,
	"SortArrayMerge":                SortArrayMerge,
}

func Test_allFunctionsShouldRunCorrectly(t *testing.T) {

	parameters := []struct {
		name     string
		lists    [][]int
		expected []int
	}{
		{
			name: "3 normal lists",
			lists: [][]int{
				{-1, 2, 4, 9, -2, 4, 10},
				{-1, 3, 6, 9, -2, 12, 7},
				{-8, 8, 8, 5, 1, 9, 11, -2, 12, 7},
			},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
		},
		{
			name: "nil slices",
			lists: [][]int{
				nil,
				nil,
				nil,
			},
			expected: []int{},
		},
		{
			name: "all negative slices",
			lists: [][]int{
				{-1, -2, -3, -4},
				{-5, -3, -3, -4},
				{-1, -6, -32, -4},
			},
			expected: []int{},
		},
		{
			name: "same number",
			lists: [][]int{
				{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
				{2, 2, 2, 2, 2, 2},
				{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
			},
			expected: []int{2},
		},
		{
			name: "only one slice",
			lists: [][]int{
				{5, 4, 4, 5, 5, 4, 5, 4},
			},
			expected: []int{4, 5},
		},
		{
			name: "one slice with a nil slice",
			lists: [][]int{
				{5, 4, 4, 5, 5, 4, 5, 4},
				nil,
			},
			expected: []int{4, 5},
		},
	}

	for _, param := range parameters {
		t.Run(param.name, func(t *testing.T) {
			for fName, fn := range functions {
				t.Run(fName, func(t *testing.T) {
					copy := copyLists(param.lists)
					res := fn(copy...)
					assert.Equal(t, param.expected, res, fName)
				})
			}
		})
	}

}

func copyLists(lists [][]int) [][]int {
	copy := make([][]int, len(lists))
	for i := range copy {
		copy[i] = append(copy[i], lists[i]...)
	}
	return copy
}
