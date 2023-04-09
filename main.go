package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/AldoFusterTurpin/Sudoku-kata/internal/sudoku"
)

const (
	pathDataLevel0 = "/data/level0"
	pathDataLevel1 = "/data/level1"
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
		sudoku.SolveLevel_0(cwd + pathDataLevel0)
	case 1:
		if err := sudoku.SolveLevel_1(cwd + pathDataLevel1); err != nil {
			fmt.Println(err)
		}
	default:
		fmt.Printf("Unknown Level %v. Valid values: 0 or 1", *level)
	}
}
