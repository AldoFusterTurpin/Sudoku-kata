package sudoku_test

import (
	"testing"

	"github.com/AldoFusterTurpin/Sudoku-kata/internal/sudoku"
)

func TestValidateMatrixSolutionShouldReturnFalseForInvalidMaxtrix(t *testing.T) {
	input := [][]int{{1, 2, 3, 4}, {2, 1, 4, 3}, {3, 4, 1, 2}, {4, 3, 2, 1}}
	expected := false

	got := sudoku.ValidateMatrix(input)
	if got != expected {
		t.Errorf("expected %vm but got %v", expected, got)
	}
}
