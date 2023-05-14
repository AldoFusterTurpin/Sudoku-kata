package sudoku_test

import (
	"reflect"
	"testing"

	"github.com/AldoFusterTurpin/Sudoku-kata/internal/sudoku"
)

func TestMatrixIsValid(t *testing.T) {
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
			got := sudoku.MatrixIsValid(tc.input)

			if tc.expected != got {
				t.Fatalf("expected %v, but got %v", tc.expected, got)
			}
		})
	}
}

func TestProposedSolutionIsValid(t *testing.T) {
	type testCase struct {
		grid             [][]int
		proposedSolution [][]int
		expected         bool
	}

	tests := map[string]testCase{
		"ShouldReturnTrueIfGridIsContainedInSolutionAndProposedSolutionIsValid_WrongStatement": {
			grid: [][]int{
				{5, 3, -1, -1, 7, -1, -1, -1, -1},
				{6, -1, -1, 1, 9, 5, -1, -1, -1},
				{-1, 9, 8, -1, -1, -1, -1, 6, -1},
				{8, -1, -1, -1, 6, -1, -1, -1, 3},
				{4, -1, -1, 8, -1, 3, -1, -1, 1},
				{7, -1, -1, -1, 2, -1, -1, -1, 6},
				{-1, 6, -1, -1, -1, -1, 2, 8, -1},
				{-1, -1, -1, 4, 1, 9, -1, -1, 5},
				{-1, -1, -1, -1, 8, -1, -1, 7, 9},
			},
			proposedSolution: [][]int{
				{5, 3, 4, 6, 7, 8, 9, 1, 2},
				{6, 7, 2, 1, 9, 5, 3, 4, 8},
				{1, 9, 8, 3, 4, 2, 5, 6, 7},
				{8, 5, 9, 7, 6, 1, 4, 2, 3},
				{4, 2, 6, 8, 5, 3, 7, 9, 1},
				{7, 1, 3, 9, 2, 4, 8, 5, 6},
				{9, 6, 1, 5, 3, 7, 2, 8, 4},
				{2, 8, 7, 4, 1, 9, 6, 3, 5},
				{3, 4, 5, 2, 8, 6, 1, 7, 9}},
			expected: true,
		},
		"ShouldReturnFalseIfNumberOfColumnsBetweenGridAndProposedSolutionIsDifferent": {
			grid: [][]int{{5, 3, -1, -1, 7, -1, -1, -1},
				{6, -1, -1, 1, 9, 5, -1, -1},
				{-1, 9, 8, -1, -1, -1, -1, 6},
				{8, -1, -1, -1, 6, -1, -1, -1},
				{4, -1, -1, 8, 7, 3, -1, -1},
				{7, -1, -1, -1, 2, -1, -1, -1},
				{-1, 6, -1, -1, -1, -1, 2, 8},
				{-1, -1, -1, 4, 1, 9, -1, -1},
				{-1, 1, -1, -1, 8, -1, -1, 7}},
			proposedSolution: [][]int{{5, 3, 4, 6, 7, 8, 9, 1, 2},
				{6, 7, 2, 1, 9, 5, 3, 4, 8},
				{1, 9, 8, 3, 4, 2, 5, 6, 7},
				{8, 5, 9, 7, 6, 1, 4, 2, 3},
				{4, 2, 6, 8, 5, 3, 7, 9, 1},
				{7, 1, 3, 9, 2, 4, 8, 5, 6},
				{9, 6, 1, 5, 3, 7, 2, 8, 4},
				{2, 8, 7, 4, 1, 9, 6, 3, 5},
				{3, 4, 5, 2, 8, 6, 1, 7, 9}},
			expected: false,
		},
		"ShouldReturnFalseIfNumberOfRowsBetweenGridAndProposedSolutionIsDifferent": {
			grid: [][]int{{5, 3, -1, -1, 7, -1, -1, -1, -1},
				{6, -1, -1, 1, 9, 5, -1, -1, -1},
				{-1, 9, 8, -1, -1, -1, -1, 6, -1},
				{8, -1, -1, -1, 6, -1, -1, -1, 3},
				{4, -1, -1, 8, 7, 3, -1, -1, 1},
				{7, -1, -1, -1, 2, -1, -1, -1, 6},
				{-1, 6, -1, -1, -1, -1, 2, 8, -1},
				{-1, -1, -1, 4, 1, 9, -1, -1, 5},
				{-1, 1, -1, -1, 8, -1, -1, 7, 9}},
			proposedSolution: [][]int{{5, 3, 4, 6, 7, 8, 9, 1, 2},
				{6, 7, 2, 1, 9, 5, 3, 4, 8},
				{1, 9, 8, 3, 4, 2, 5, 6, 7},
				{8, 5, 9, 7, 6, 1, 4, 2, 3},
				{4, 2, 6, 8, 5, 3, 7, 9, 1},
				{7, 1, 3, 9, 2, 4, 8, 5, 6},
				{9, 6, 1, 5, 3, 7, 2, 8, 4},
				{2, 8, 7, 4, 1, 9, 6, 3, 5}},
			expected: false,
		},
		"ShouldReturnFalseIfGridIsContainedInSolutionButProposedSolutionIsNotValid": {
			grid: [][]int{{5, 3, -1, -1, 7, -1, -1, -1, -1},
				{6, -1, -1, 1, 9, 5, -1, -1, -1},
				{-1, 9, 8, -1, -1, -1, -1, 6, -1},
				{8, -1, -1, -1, 6, -1, -1, -1, 3},
				{4, -1, -1, 8, 7, 3, -1, -1, 1},
				{7, -1, -1, -1, 2, -1, -1, -1, 6},
				{-1, 6, -1, -1, -1, -1, 2, 8, -1},
				{-1, -1, -1, 4, 1, 9, -1, -1, 5},
				{-1, 1, -1, -1, 8, -1, -1, 7, 9}},
			proposedSolution: [][]int{{5, 3, 4, 6, 7, 8, 9, 1, 2},
				{6, 7, 2, 1, 9, 5, 3, 4, 8},
				{1, 9, 8, 3, 4, 2, 5, 6, 7},
				{8, 5, 9, 7, 6, 1, 4, 2, 3},
				{4, 2, 6, 8, 7, 3, 7, 9, 1},
				{7, 1, 3, 9, 2, 4, 8, 5, 6},
				{9, 6, 1, 5, 3, 7, 2, 8, 4},
				{2, 8, 7, 4, 1, 9, 6, 3, 5},
				{3, 1, 5, 2, 8, 6, 1, 7, 9}},
			expected: false,
		},
		"ShouldReturnFalseIfGridIsContainedInSolutionButInvalidProposedSolution": {
			grid: [][]int{{5, 3, -1, -1, 7, -1, -1, -1, -1},
				{6, -1, -1, 1, 9, 5, -1, -1, -1},
				{-1, 9, 8, -1, -1, -1, -1, 6, -1},
				{8, -1, -1, -1, 6, -1, -1, -1, 3},
				{4, -1, -1, 8, 7, 3, -1, -1, 1},
				{7, -1, -1, -1, 2, -1, -1, -1, 6},
				{-1, 6, -1, -1, -1, -1, 2, 8, -1},
				{-1, -1, -1, 4, 1, 9, -1, -1, 5},
				{-1, 1, -1, -1, 8, -1, -1, 7, 9}},
			proposedSolution: [][]int{{4, 5, 3, 8, 2, 6, 1, 9, 7},
				{8, 9, 2, 5, 7, 1, 6, 3, 4},
				{1, 6, 7, 4, 9, 3, 5, 2, 8},
				{7, 1, 4, 9, 5, 2, 8, 6, 3},
				{5, 8, 6, 1, 3, 7, 2, 4, 9},
				{3, 2, 9, 6, 8, 4, 7, 5, 1},
				{9, 3, 5, 2, 1, 8, 4, 7, 6},
				{6, 7, 1, 3, 4, 5, 9, 8, 2},
				{2, 4, 8, 7, 6, 9, 3, 1, 5}},
			expected: false,
		},
		"ShouldReturnFalseIfGridIsContainedInSolutionButSolutionIsNotValidDueToRepeated3InRow0": {
			grid: [][]int{{5, 3, -1, -1, 7, -1, -1, -1, -1},
				{6, -1, -1, 1, 9, 5, -1, -1, -1},
				{-1, 9, 8, -1, -1, -1, -1, 6, -1},
				{8, -1, -1, -1, 6, -1, -1, -1, 3},
				{4, -1, -1, 8, 7, 3, -1, -1, 1},
				{7, -1, -1, -1, 2, -1, -1, -1, 6},
				{-1, 6, -1, -1, -1, -1, 2, 8, -1},
				{-1, -1, -1, 4, 1, 9, -1, -1, 5},
				{-1, 1, -1, -1, 8, -1, -1, 7, 9}},
			proposedSolution: [][]int{{5, 3, 3, 6, 7, 8, 9, 1, 2},
				{6, 7, 2, 1, 9, 5, 3, 4, 8},
				{1, 9, 8, 3, 4, 2, 5, 6, 7},
				{8, 5, 9, 7, 6, 1, 4, 2, 3},
				{4, 2, 6, 8, 7, 3, 7, 9, 1},
				{7, 1, 3, 9, 2, 4, 8, 5, 6},
				{9, 6, 1, 5, 3, 7, 2, 8, 4},
				{2, 8, 7, 4, 1, 9, 6, 3, 5},
				{3, 1, 5, 2, 8, 6, 1, 7, 9}},
			expected: false,
		},
		"ShouldReturnTrueIfGridIsContainedInSolutionAndProposedSolutionIsValid": {
			grid: [][]int{{6, -1, -1, -1, 5, 8, -1, -1, -1},
				{-1, -1, -1, -1, -1, -1, -1, -1, 5},
				{4, 3, -1, -1, -1, -1, -1, 8, -1},
				{5, -1, 6, -1, -1, -1, -1, -1, -1},
				{-1, 7, -1, 1, 6, -1, -1, 4, -1},
				{9, -1, 3, -1, 4, -1, -1, -1, -1},
				{-1, -1, -1, -1, 7, -1, 5, -1, 8},
				{-1, -1, -1, 9, -1, -1, -1, 1, 7},
				{-1, 9, -1, -1, 3, -1, -1, 6, 4}},
			proposedSolution: [][]int{{6, 2, 7, 4, 5, 8, 1, 3, 9},
				{1, 8, 9, 6, 2, 3, 4, 7, 5},
				{4, 3, 5, 7, 1, 9, 6, 8, 2},
				{5, 4, 6, 3, 9, 7, 8, 2, 1},
				{8, 7, 2, 1, 6, 5, 9, 4, 3},
				{9, 1, 3, 8, 4, 2, 7, 5, 6},
				{3, 6, 1, 2, 7, 4, 5, 9, 8},
				{2, 5, 4, 9, 8, 6, 3, 1, 7},
				{7, 9, 8, 5, 3, 1, 2, 6, 4}},
			expected: true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := sudoku.ProposedSolutionIsValid(tc.grid, tc.proposedSolution)

			if tc.expected != got {
				t.Fatalf("expected %v, but got %v", tc.expected, got)
			}
		})
	}

}

func TestProposedSolutionIsPartiallyValid(t *testing.T) {
	type testCase struct {
		grid             [][]int
		proposedSolution [][]int
		expected         bool
	}

	tests := map[string]testCase{
		"ShouldReturnTrueIfGridIsContainedInSolutionAndProposedSolutionIsPartialComplete": {
			grid: [][]int{{6, -1, -1, -1, 5, 8, -1, -1, -1},
				{-1, -1, -1, -1, -1, -1, -1, -1, 5},
				{4, 3, -1, -1, -1, -1, -1, 8, -1},
				{5, -1, 6, -1, -1, -1, -1, -1, -1},
				{-1, 7, -1, 1, 6, -1, -1, 4, -1},
				{9, -1, 3, -1, 4, -1, -1, -1, -1},
				{-1, -1, -1, -1, 7, -1, 5, -1, 8},
				{-1, -1, -1, 9, -1, -1, -1, 1, 7},
				{-1, 9, -1, -1, 3, -1, -1, 6, 4}},
			proposedSolution: [][]int{{-1, 2, 7, 4, 5, 8, 1, 3, 9},
				{1, 8, 9, 6, 2, 3, 4, 7, 5},
				{4, 3, 5, 7, 1, 9, 6, 8, 2},
				{5, 4, 6, 3, 9, 7, 8, 2, 1},
				{8, 7, 2, 1, 6, 5, 9, 4, 3},
				{9, 1, 3, 8, 4, 2, 7, 5, 6},
				{3, 6, 1, 2, 7, 4, 5, 9, 8},
				{2, 5, 4, 9, 8, 6, 3, 1, 7},
				{7, 9, 8, 5, 3, 1, 2, 6, 4}},
			expected: true,
		},
		"Should return false if grid is contained in solution but proposed solution is NOT partially complete due to repeated number 6 in column 0": {
			grid: [][]int{{6, -1, -1, -1, 5, 8, -1, -1, -1},
				{6, -1, -1, -1, -1, -1, -1, -1, 5},
				{4, 3, -1, -1, -1, -1, -1, 8, -1},
				{5, -1, 6, -1, -1, -1, -1, -1, -1},
				{-1, 7, -1, 1, 6, -1, -1, 4, -1},
				{9, -1, 3, -1, 4, -1, -1, -1, -1},
				{-1, -1, -1, -1, 7, -1, 5, -1, 8},
				{-1, -1, -1, 9, -1, -1, -1, 1, 7},
				{-1, 9, -1, -1, 3, -1, -1, 6, 4}},
			proposedSolution: [][]int{{-1, 2, 7, 4, 5, 8, 1, 3, 9},
				{1, 8, 9, 6, 2, 3, 4, 7, 5},
				{4, 3, 5, 7, 1, 9, 6, 8, 2},
				{5, 4, 6, 3, 9, 7, 8, 2, 1},
				{8, 7, 2, 1, 6, 5, 9, 4, 3},
				{9, 1, 3, 8, 4, 2, 7, 5, 6},
				{3, 6, 1, 2, 7, 4, 5, 9, 8},
				{2, 5, 4, 9, 8, 6, 3, 1, 7},
				{7, 9, 8, 5, 3, 1, 2, 6, 4}},
			expected: false,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := sudoku.ProposedSolutionIsPartiallyValid(tc.grid, tc.proposedSolution)

			if tc.expected != got {
				t.Fatalf("expected %v, but got %v", tc.expected, got)
			}
		})
	}
}

func TestSolveSudoku(t *testing.T) {
	type testCase struct {
		grid             [][]int
		expectedSolution [][]int
	}

	tests := map[string]testCase{
		"ShouldReturnTheSudokuSolved_IfItHasASolution": {
			grid: [][]int{{5, 3, -1, -1, 7, -1, -1, -1, -1},
				{6, -1, -1, 1, 9, 5, -1, -1, -1},
				{-1, 9, 8, -1, -1, -1, -1, 6, -1},
				{8, -1, -1, -1, 6, -1, -1, -1, 3},
				{4, -1, -1, 8, -1, 3, -1, -1, 1},
				{7, -1, -1, -1, 2, -1, -1, -1, 6},
				{-1, 6, -1, -1, -1, -1, 2, 8, -1},
				{-1, -1, -1, 4, 1, 9, -1, -1, 5},
				{-1, -1, -1, -1, 8, -1, -1, 7, 9}},
			expectedSolution: [][]int{
				{5, 3, 4, 6, 7, 8, 9, 1, 2},
				{6, 7, 2, 1, 9, 5, 3, 4, 8},
				{1, 9, 8, 3, 4, 2, 5, 6, 7},
				{8, 5, 9, 7, 6, 1, 4, 2, 3},
				{4, 2, 6, 8, 5, 3, 7, 9, 1},
				{7, 1, 3, 9, 2, 4, 8, 5, 6},
				{9, 6, 1, 5, 3, 7, 2, 8, 4},
				{2, 8, 7, 4, 1, 9, 6, 3, 5},
				{3, 4, 5, 2, 8, 6, 1, 7, 9}},
		},
		"ShouldReturnTheSudokuSolved_IfItHasASolution_2": {
			grid: [][]int{
        {5, 3, 4, 6, 7, 8, 9, 1, 2},
        {6, 7, 2, 1, 9, 5, 3, 4, 8},
        {1, 9, 8, 3, 4, 2, 5, 6, 7},
        {8, 5, 9, 7, 6, 1, 4, 2, 3},
        {4, 2, 6, 8, 5, 3, 7, 9, 1},
        {7, 1, 3, 9, 2, 4, 8, 5, 6},
        {9, 6, 1, 5, 3, 7, 2, 8, 4},
        {2, 8, 7, 4, 1, 9, 6, 3, 5},
        {3, 4, 5, 2, 8, 6, 1, 7, -1}},
      expectedSolution: [][]int{
				{5, 3, 4, 6, 7, 8, 9, 1, 2},
				{6, 7, 2, 1, 9, 5, 3, 4, 8},
				{1, 9, 8, 3, 4, 2, 5, 6, 7},
				{8, 5, 9, 7, 6, 1, 4, 2, 3},
				{4, 2, 6, 8, 5, 3, 7, 9, 1},
				{7, 1, 3, 9, 2, 4, 8, 5, 6},
				{9, 6, 1, 5, 3, 7, 2, 8, 4},
				{2, 8, 7, 4, 1, 9, 6, 3, 5},
				{3, 4, 5, 2, 8, 6, 1, 7, 9},
        },
		},
		"ShouldReturnTheSudokuSolved With a hard sudoku": {
			// https://abcnews.go.com/blogs/headlines/2012/06/can-you-solve-the-hardest-ever-sudoku
			// TODO: solve this test case
			grid: [][]int{
				{8, -1, -1, -1, -1, -1, -1, -1, -1},
				{-1, -1, 3, 6, -1, -1, -1, -1, -1},
				{-1, 7, -1, -1, 9, -1, 2, -1, -1},
				{-1, 5, -1, -1, -1, 7, -1, -1, -1},
				{-1, -1, -1, -1, 4, 5, 7, -1, -1},
				{-1, -1, -1, 1, -1, -1, -1, 3, -1},
				{-1, -1, 1, -1, -1, -1, -1, 6, 8},
				{-1, -1, 8, 5, -1, -1, -1, 1, -1},
				{-1, 9, -1, -1, -1, -1, 4, -1, -1}},
			expectedSolution: [][]int{
				{8, 1, 2, 7, 5, 3, 6, 4, 9},
				{9, 4, 3, 6, 8, 2, 1, 7, 5},
				{6, 7, 5, 4, 9, 1, 2, 8, 3},
				{1, 5, 4, 2, 3, 7, 8, 9, 6},
				{3, 6, 9, 8, 4, 5, 7, 2, 1},
				{2, 8, 7, 1, 6, 9, 5, 3, 4},
				{5, 2, 1, 9, 7, 4, 3, 6, 8},
				{4, 3, 8, 5, 2, 6, 9, 1, 7},
				{7, 9, 6, 3, 1, 8, 4, 5, 2}},
		},
		// TODO Returning nil in this test, due to invalid grid.
		/* "ShouldReturnTheSudokuSolved_Lelele": {
			grid: [][]int{{5, 3, 4, 4, 7, 4, 4, 4, -1},
				{6, 4, 4, 1, 9, 5, 4, 4, 9},
				{4, 9, 8, 4, 4, 4, 4, 6, 4},
				{8, 4, 4, 4, 6, 4, 4, 4, 3},
				{4, 4, 4, 8, 4, 3, 4, 4, 1},
				{7, 4, 4, 4, 2, 4, 4, 4, 6},
				{4, 6, 4, 4, 4, 4, 2, 8, 4},
				{4, 4, 4, 4, 1, 9, 4, 4, 5},
				{4, 4, 4, 4, 8, 4, -1, 7, 9}},
			expectedSolution: [][]int{
				{5, 3, 4, 6, 7, 8, 9, 1, 2},
				{6, 7, 2, 1, 9, 5, 3, 4, 8},
				{1, 9, 8, 3, 4, 2, 5, 6, 7},
				{8, 5, 9, 7, 6, 1, 4, 2, 3},
				{4, 2, 6, 8, 5, 3, 7, 9, 1},
				{7, 1, 3, 9, 2, 4, 8, 5, 6},
				{9, 6, 1, 5, 3, 7, 2, 8, 4},
				{2, 8, 7, 4, 1, 9, 6, 3, 5},
				{3, 4, 5, 2, 8, 6, 1, 7, 9}},
		}, */
		"ShouldReturnNil_IfCannotBeSolved": {
			grid: [][]int{
        {5, 3, -1, -1, 7, -1, -1, -1, -1},
				{6, -1, -1, 1, 9, 5, -1, -1, -1},
				{-1, 9, 8, -1, -1, -1, -1, 6, -1},
				{8, -1, -1, -1, 6, -1, -1, -1, 3},
				{4, -1, -1, 8, -1, 3, -1, -1, 1},
				{7, -1, -1, -1, 2, -1, -1, -1, 6},
				{-1, 6, -1, -1, -1, -1, 2, 8, -1},
				{-1, -1, -1, 3, 1, 9, -1, -1, -1},
				{-1, -1, -1, -1, 8, -1, -1, 7, 9}},
			expectedSolution: nil,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			_, got := sudoku.SolveBacktracking(tc.grid)

			isEqual := reflect.DeepEqual(got, tc.expectedSolution)
			if !isEqual {
				t.Fatalf("expected %v, but got %v", tc.expectedSolution, got)
			}
		})
	}
}
