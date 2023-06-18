package sudoku

const emptyCell = 0

func copyMatrix(m [][]int) [][]int {
	rows := len(m)
	cp := make([][]int, rows)
	for i := 0; i < rows; i++ {
		cp[i] = make([]int, len(m[i]))
		copy(cp[i], m[i])
	}
	return cp
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
