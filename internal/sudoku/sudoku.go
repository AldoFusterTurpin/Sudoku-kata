package sudoku

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
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

func SolveLevel_1(path string) error {
	grid, err := getMatrixFromPath(path + "/grid.csv")
	if err != nil {
		return err
	}

	proposedSolution, err := getMatrixFromPath(path + "/solution.csv")
	if err != nil {
		return err
	}

	if ProposedSolutionIsValid(grid, proposedSolution) {
		fmt.Println("The proposed solution is correct")
		return nil
	}

	fmt.Println("The proposed solution is incorrect")
	return nil
}

func SolveLevel_0(path string) {
	processPathFn := WalkFn()
	filepath.Walk(path, processPathFn)
}

// WalkFnLevel0() returns the function that will be called for each path.
// It always returns nil as the error handling (printing) is performed inside this function.
func WalkFn() filepath.WalkFunc {
	return func(fullPath string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		if info.IsDir() {
			return nil
		}

		err = validateMatrixFromPath(fullPath)
		if err != nil {
			log.Print(err)
		}
		return nil
	}
}

func validateMatrixFromPath(fullPath string) error {
	matrix, err := getMatrixFromPath(fullPath)
	if err != nil {
		return err
	}

	fileName := filepath.Base(fullPath)
	fmt.Printf("%v: ", fileName)

	if MatrixIsValid(matrix) {
		fmt.Println("The input comply with Sudoku's rules.")
		return nil
	}

	fmt.Println("The input doesn't comply with Sudoku's rules.")
	return nil
}

func getMatrixFromPath(fullPath string) ([][]int, error) {
	input, err := getContentOfFile(fullPath)
	if err != nil {
		return nil, err
	}

	matrix, err := convertInputToMatrix(input)
	if err != nil {
		return nil, err
	}
	return matrix, nil
}

func convertInputToMatrix(input string) ([][]int, error) {
	input = strings.ReplaceAll(input, "\n\n", "\n")
	input = strings.ReplaceAll(input, " ", "")
	lines := strings.Split(input, "\n")
	var matrix [][]int

	for _, line := range lines {
		if line == "" {
			continue
		}

		line = strings.TrimSuffix(line, ",")
		elements := strings.Split(line, ",")
		var row []int

		for _, element := range elements {
			v, err := strconv.Atoi(element)
			if err != nil {
				return nil, err
			}
			row = append(row, v)
		}

		matrix = append(matrix, row)
	}

	return matrix, nil
}

func getContentOfFile(filePath string) (string, error) {
	s, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(s), nil
}
