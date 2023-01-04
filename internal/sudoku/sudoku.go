package sudoku

func ValidateMatrix(matrix [][]int) bool {
	for _, line := range matrix {
		if !valdiateSlice(line) {
			return false
		}
	}
	// 12
	// 12

	nRows := len(matrix)
	nCols := len(matrix[0])

	var col []int
	for i := 0; i < nRows; i++ {
		for j := 0; j < nCols; j++ {
			element := matrix[j][i]
		  col = append(col, element)
		}
    valdiateSlice(col)
	}
  
  return false
	//valdiateColumns

	//validateSquares
}

func valdiateSlice(slice []int) bool {
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
