package main

import (
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
	currentDirectory, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	iteratePath(currentDirectory + "/data")
}

func iteratePath(path string) {
	processPathFn := processPathFn()
	filepath.Walk(path, processPathFn)
}

// processPathFn() returns the function that will be called for each path.
// It always returns nil as the error handling (printing)
// is performed inside this function.
func processPathFn() filepath.WalkFunc {
	return func(fullPath string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		if info.IsDir() {
			return nil
		}

		err = processFilePath(fullPath)
		if err != nil {
			log.Print(err)
		}
		return nil
	}
}

func processFilePath(fullPath string) error {
	input, err := getContentOfFile(fullPath)
	if err != nil {
		return err
	}

	matrix, err := convertInputToMatrix(input)
	if err != nil {
		return err
	}

	fileName := filepath.Base(fullPath)
	fmt.Printf("%v: ", fileName)

	if sudoku.IsMatrixValid(matrix) {
		fmt.Println("The input comply with Sudoku's rules.")
		return nil
	}

	fmt.Println("The input doesn't comply with Sudoku's rules.")
	return nil
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
