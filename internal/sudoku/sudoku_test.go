package sudoku_test

import (
	"testing"

	"github.com/AldoFusterTurpin/Sudoku-kata/internal/sudoku"
)

func TestIsValidMatrixShouldReturnFalseForInvalidMatrix(t *testing.T) {
	input := [][]int{{1, 2, 3, 4}, {2, 1, 4, 3}, {3, 4, 1, 2}, {4, 3, 2, 1}}
	expected := false

	got := sudoku.IsMatrixValid(input)
	if got != expected {
		t.Errorf("expected %v, but got %v", expected, got)
	}
}

func TestIsValidMatrixShouldReturnFalseFoReallyWrongMatrix(t *testing.T) {
	input := [][]int{{1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}}
	expected := false

	got := sudoku.IsMatrixValid(input)
	if got != expected {
		t.Errorf("expected %v, but got %v", expected, got)
	}
}

func TestIsValidMatrixShouldReturnTrueForValidMatrix(t *testing.T) {
	input := [][]int{{1, 2, 3, 4}, {3, 4, 1, 2}, {2, 3, 4, 1}, {4, 1, 2, 3}}
	expected := true

	got := sudoku.IsMatrixValid(input)
	if got != expected {
		t.Errorf("expected %v, but got %v", expected, got)
	}
}

func TestIsValidMatrixShouldReturnFalseForNotValidRangeMatrix(t *testing.T) {
	input := [][]int{{1}, {1, 2}, {3, 2, 1}, {4, 1, 2, 3}}
	expected := false

	got := sudoku.IsMatrixValid(input)
	if got != expected {
		t.Errorf("expected %v, but got %v", expected, got)
	}
}
