package sudoku

import "fmt"

/*
Aldo mental notes:

Given an input like
5,3, , ,7, , , ,
6, , ,1,9,5, , ,
 ,9,8, , , , ,6,
8, , , ,6, , , ,3
4, , ,8, ,3, , ,1
7, , , ,2, , , ,6
 ,6, , , , ,2,8,
 , , ,4,1,9, , ,5
 , , , ,8, , ,7,9
we need to get all the possible combinations but excluding the ones that are not valid.

This problem is a clear candidate to be solved using Backtracking as we need to generate all the combinations (brute force) but removing the invalid ones as soon as possible.
This kind of problems are tipycal solved using a recursive function to easily generate the tree of combinations. In theory, every recursive solution can be transformed to a non-recursive one using a queue, set or array to mantain the "elements_not_processed_so_far".
Clear examples of that are the Dijkstra Algorithm (https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm#Algorithm) or my solution to the Advent Of Code 2022 Day 7 where I use a queue instead of a recursive call to maintain the "unvisited" folders :) (https://github.com/AldoFusterTurpin/AdventOfCode-2022)
Important note: recursive functions can be converted to iterative ones but Backtracking is not easily convertible to a recursive one.
*/
const emptyCell = 0

type Matrix [][]int

type CandidatesSolutions []Matrix

func copyMatrix(m [][]int) [][]int {
	rows := len(m)
	cp := make([][]int, rows)
	for i := 0; i < rows; i++ {
		cp[i] = make([]int, len(m[i]))
		copy(cp[i], m[i])
	}
	return cp
}

func Solve(grid [][]int) [][]int {
	var candidatesSolutions CandidatesSolutions

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == emptyCell {
				candidatesSolutions = append(candidatesSolutions, getMoreCandidatesSolutions(candidatesSolutions, grid, i, j)...)
			}
		}
	}

	// playing a bit
	fmt.Println(candidatesSolutions)

	return getValidSolutionFromCandidateSolutionsIfAny(grid, candidatesSolutions)
}

func SolveBacktracking(grid [][]int) (bool, [][]int) {
	l := len(grid)
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			cellIsEmpty := grid[i][j] == emptyCell
			if cellIsEmpty {
				for numberToTry := 1; numberToTry <= l; numberToTry++ {
					if placementIsValid(grid, i, j, numberToTry) {
						grid[i][j] = numberToTry
						if solved, rGrid := SolveBacktracking(grid); solved {
							return true, rGrid
						}
						grid[i][j] = emptyCell
					}
				}
				// After trying all the numbers, we return false as we did not find a valid number to place in that cell
				return false, nil
			}
		}
	}
	// At this point we have found a solution as we have not returned false so far.
	// i,e: we only reach this point if all previous positions were valid.
	return true, grid
}

// placementIsValid returns true if puting numberToTry in grid[i][j]
// would lead to a valid matrix
func placementIsValid(grid [][]int, i, j, numberToTry int) bool {
	gridCopy := copyMatrix(grid)
	gridCopy[i][j] = numberToTry
	return MatrixIsValid(gridCopy)
}

func getMoreCandidatesSolutions(candidateSolutionsIn CandidatesSolutions, grid [][]int, i int, j int) CandidatesSolutions {
	var moreCandidateSolutions CandidatesSolutions

	if len(candidateSolutionsIn) == 0 {
		if grid[i][j] == emptyCell {
			for k := 1; k <= 9; k++ {
				tmp := copyMatrix(grid)
				tmp[i][j] = k
				if ProposedSolutionIsPartiallyValid(grid, tmp) {
					moreCandidateSolutions = append(moreCandidateSolutions, tmp)
				}
			}
		}
	} else {
		for _, solution := range candidateSolutionsIn {
			if solution[i][j] == emptyCell {
				for k := 1; k <= 9; k++ {
					tmp := copyMatrix(solution)
					tmp[i][j] = k
					if ProposedSolutionIsPartiallyValid(grid, tmp) {
						moreCandidateSolutions = append(moreCandidateSolutions, tmp)
					}
				}
			}
		}
	}

	return moreCandidateSolutions
}

func getValidSolutionFromCandidateSolutionsIfAny(grid Matrix, candidateSolutions CandidatesSolutions) Matrix {
	for _, s := range candidateSolutions {
		if ProposedSolutionIsValid(grid, s) {
			return s
		}
	}

	return nil
}
