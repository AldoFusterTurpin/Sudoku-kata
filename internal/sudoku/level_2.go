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
			if grid[i][j] == -1 {
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
		for j := 0; j < len(grid[i]); j++ {
			cellIsEmpty := grid[i][j] == -1
			if cellIsEmpty {
				// try all possible numbers in the current cell
				for v := 1; v <= l; v++ {
					currentVal := grid[i][j]
					// We just put the new number in the matrix if the matrix is valid with the new number on it.
					grid[i][j] = v // put the new number
					if !MatrixIsValid(grid) {
						// if matrix is not valid, undo the last change and skip this iteration.
						grid[i][j] = currentVal
						continue
					}
          // TODO(Aldo): maybe do the matrix copy  on line 65 ?
          if b, rGrid := SolveBacktracking(copyMatrix(grid)); b { 
						return true, rGrid
					}
					// else
					grid[i][j] = -1
				}
				// After trying all the numbers, we return false as we did not find a valid solution.
				// We can think about that as a search: if at that point we have not found x value, we can return false.
				return false, nil
			}
		}
	}

	// At this point we have found a solution as we have not returned false so far.
	// i,e: we only reach this point if all previous positions were valid.
	return true, grid
}

func getMoreCandidatesSolutions(candidateSolutionsIn CandidatesSolutions, grid [][]int, i int, j int) CandidatesSolutions {
	var moreCandidateSolutions CandidatesSolutions

	if len(candidateSolutionsIn) == 0 {
		if grid[i][j] == -1 {
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
			if solution[i][j] == -1 {
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
