package sudoku

import (
	"math"
)

// IsMatrixValid checks whether the matrix of a sudoku solution is valid or not.
func IsMatrixValid(matrix [][]int) bool {
	nRows := len(matrix)
	nCols := len(matrix[0])

	if nRows != nCols {
		return false
	}

	return validateRowsAndCols(matrix, nCols) && validateBoxes(matrix, nRows)
}

func validateRowsAndCols(matrix [][]int, nCols int) bool {
	for i, line := range matrix {
		if !validateSlice(line) {
			return false
		}

		var col []int

		for j := 0; j < nCols; j++ {
			element := matrix[j][i]
			col = append(col, element)
		}

		if !validateSlice(col) {
			return false
		}
	}
	return true
}

func validateSlice(slice []int) bool {
	uniques := make(map[int]struct{})

	for _, element := range slice {
		_, present := uniques[element]

		if present {
			return false
		}

		uniques[element] = struct{}{}
	}

	return true
}

func validateBoxes(matrix [][]int, nRows int) bool {
	numberOfBoxes := int(math.Sqrt(float64(nRows)))

	var startI, startJ int
	endI := numberOfBoxes
	endJ := numberOfBoxes

	for endI <= nRows && endJ <= nRows {
		box := sliceFromBox(matrix, startI, startJ, endI, endJ)

		if !validateSlice(box) {
			return false
		}

		startI += numberOfBoxes
		startJ += numberOfBoxes
		endI += numberOfBoxes
		endJ += numberOfBoxes
	}

	return true
}

func sliceFromBox(matrix [][]int, startRowIndex, startColIndex, endRowIndex, endColIndex int) []int {
	var box []int
	for i := startRowIndex; i < endRowIndex; i++ {

		for j := startColIndex; j < endColIndex; j++ {
			element := matrix[i][j]
			box = append(box, element)
		}
	}

	return box
}
