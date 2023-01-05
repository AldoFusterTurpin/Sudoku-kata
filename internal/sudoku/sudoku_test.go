package sudoku_test

import (
	"testing"

	"github.com/AldoFusterTurpin/Sudoku-kata/internal/sudoku"
)

func TestIsMatrixValid(t *testing.T) {
	type testCase struct {
		input    [][]int
		expected bool
	}

	tests := map[string]testCase{
		"ShouldReturnFalseForInvalidMatrix": {
			input:    [][]int{{1, 2, 3, 4}, {2, 1, 4, 3}, {3, 4, 1, 2}, {4, 3, 2, 1}},
			expected: false,
		},
		"ShouldReturnFalseFoReallyWrongMatrix": {
			input:    [][]int{{1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}},
			expected: false,
		},
		"ShouldReturnTrueForValidMatrix": {
			input:    [][]int{{1, 2, 3, 4}, {3, 4, 1, 2}, {2, 3, 4, 1}, {4, 1, 2, 3}},
			expected: true,
		},
		"ShouldReturnFalseForNotValidRangeMatrix": {
			input:    [][]int{{1}, {1, 2}, {3, 2, 1}, {4, 1, 2, 3}},
			expected: false,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := sudoku.IsMatrixValid(tc.input)

			if tc.expected != got {
				t.Fatalf("expected %v, but got %v", tc.expected, got)
			}
		})
	}
}
