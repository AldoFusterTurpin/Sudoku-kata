package sudoku_test

import (
	"testing"

	"github.com/AldoFusterTurpin/Sudoku-kata/internal/sudoku"
)

func TestValidateMatrixSolutionShouldReturnFalseForInvalidMatrix(t *testing.T) {
	input := [][]int{{1, 2, 3, 4}, {2, 1, 4, 3}, {3, 4, 1, 2}, {4, 3, 2, 1}}
	expected := false

	got := sudoku.ValidateMatrix(input)
	if got != expected {
		t.Errorf("expected %v, but got %v", expected, got)
	}
}

func TestValidateMatrixSolutionShouldReturnFalseFoReallyWrongMatrix(t *testing.T) {
	input := [][]int{{1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}}
	expected := false

	got := sudoku.ValidateMatrix(input)
	if got != expected {
		t.Errorf("expected %v, but got %v", expected, got)
	}
}

func TestValidateMatrixSolutionShouldReturnTrueForValidMatrix(t *testing.T) {
	input := [][]int{{1, 2, 3, 4}, {3, 4, 1, 2}, {2, 3, 4, 1}, {4, 1, 2, 3}}
	expected := true

	got := sudoku.ValidateMatrix(input)
	if got != expected {
		t.Errorf("expected %v, but got %v", expected, got)
	}
}

func TestValidateMatrixSolutionShouldReturnFalseForNotValidRangeMatrix(t *testing.T) {
	input := [][]int{{1}, {1, 2}, {3, 2, 1}, {4, 1, 2, 3}}
	expected := false

	got := sudoku.ValidateMatrix(input)
	if got != expected {
		t.Errorf("expected %v, but got %v", expected, got)
	}
}
