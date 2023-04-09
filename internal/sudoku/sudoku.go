package sudoku

import (
	"math"
)

// MatrixIsValid checks whether the matrix of a sudoku solution is valid or not.
func MatrixIsValid(matrix [][]int) bool {
	nRows := len(matrix)
	nCols := len(matrix[0])

	if nRows != nCols {
		return false
	}

	return areRowsAndColsValid(matrix, nCols) && areBoxesValid(matrix, nRows)
}

func areRowsAndColsValid(matrix [][]int, nCols int) bool {
	for i, line := range matrix {
		if !isSliceValid(line) {
			return false
		}

		var col []int

		for j := 0; j < nCols; j++ {
			element := matrix[j][i]
			col = append(col, element)
		}

		if !isSliceValid(col) {
			return false
		}
	}
	return true
}

func isSliceValid(slice []int) bool {
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

func areBoxesValid(matrix [][]int, nRows int) bool {
	numberOfBoxes := int(math.Sqrt(float64(nRows)))

	var startI, startJ int
	endI := numberOfBoxes
	endJ := numberOfBoxes

	for endI <= nRows && endJ <= nRows {
		box := createBoxSliceFromMatrix(matrix, startI, startJ, endI, endJ)

		if !isSliceValid(box) {
			return false
		}

		startI += numberOfBoxes
		startJ += numberOfBoxes
		endI += numberOfBoxes
		endJ += numberOfBoxes
	}

	return true
}

func createBoxSliceFromMatrix(matrix [][]int, startRowIndex, startColIndex, endRowIndex, endColIndex int) []int {
	var box []int
	for i := startRowIndex; i < endRowIndex; i++ {

		for j := startColIndex; j < endColIndex; j++ {
			element := matrix[i][j]
			box = append(box, element)
		}
	}

	return box
}

// Precondition: grid and proposedSolution have a valid shape (n x n with n > 0)
func ProposedSolutionIsValid(grid [][]int, proposedSolution [][]int) bool {
	return gridCorrespondsToProposedSolution(grid, proposedSolution) && MatrixIsValid(proposedSolution)
}

func gridCorrespondsToProposedSolution(grid [][]int, proposedSolution [][]int) bool {
	gridRows := len(grid)
	proposedSolutionRows := len(proposedSolution)

	if gridRows != proposedSolutionRows {
		return false
	}

	gridCols := len(grid[0])
	proposedSolutionCols := len(proposedSolution[0])

	if gridCols != proposedSolutionCols {
		return false
	}

	for i := 0; i < gridRows; i++ {
		for j := 0; j < gridCols; j++ {
			cellContainsNumber := grid[i][j] != -1
			if cellContainsNumber && grid[i][j] != proposedSolution[i][j] {
				return false
			}
		}
	}
	return true
}
