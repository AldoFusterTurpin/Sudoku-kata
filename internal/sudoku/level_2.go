package sudoku

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
This kind of problems are tipycal solved using a recursive function to easily generate the tree of combinations. In theory, every recursive solution can be transformed to a non-recursive one using a queue, set or array to mantain the "elements_not_processed_so_far". Clear examples of that are the Dijkstra Algorithm (https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm#Algorithm) or my solution to the Advent Of Code 2022 Day 7 where I use a queue instead of a recursive call to maintain the "unvisited" folders :) (https://github.com/AldoFusterTurpin/AdventOfCode-2022)
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
				candidatesSolutions = getMoreCandidatesSolutions(candidatesSolutions, grid, i, j)
			}
		}
	}

	return getValidSolutionFromCandidateSolutionsIfAny(grid, candidatesSolutions)
}

func getValidSolutionFromCandidateSolutionsIfAny(grid Matrix, candidateSolutions CandidatesSolutions) Matrix {
	for _, s := range candidateSolutions {
		if ProposedSolutionIsValid(grid, s) {
			return s
		}
	}

	return nil
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
