package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/AldoFusterTurpin/Sudoku-kata/internal/sudoku"
)

func main() {
	var level = flag.Int("level", 0, "the Level of the Sudoku Kata you are solving")
	flag.Parse()

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	switch *level {
	case 0:
		solveLevel_0(cwd + "/data/level0")
	default:
		if err := solveLevel1(cwd + "/data/level1"); err != nil {
			fmt.Println(err)
		}
	}
}

func solveLevel1(path string) error {
	grid, err := getMatrixFromPath(path + "/grid.csv")
	if err != nil {
		return err
	}
	fmt.Println(grid)

	proposedSolution, err := getMatrixFromPath(path + "/solution.csv")
	if err != nil {
		return err
	}

	if sudoku.ProposedSolutionIsValid(grid, proposedSolution) {
		fmt.Println("The proposed solution is correct")
		return nil
	}

	fmt.Println("The proposed solution is incorrect")
	return nil
}

func solveLevel_0(path string) {
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

	if sudoku.MatrixIsValid(matrix) {
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
