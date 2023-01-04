package sudoku

func ValidateMatrix(matrix [][]int) bool {
	for _, line := range matrix {
		if !validateSlice(line) {
			return false
		}
	}

	// 1234
	// 5678

	nRows := len(matrix)
	nCols := len(matrix[0])

	for i := 0; i < nRows; i++ {
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
