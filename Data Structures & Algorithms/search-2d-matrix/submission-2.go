func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
        return false
    }
	
	// this approach treats the matrix as one dimensional array
	rows, cols := len(matrix), len(matrix[0])
	low, high := 0, (rows*cols)-1

	for low <= high {
		mid := (high-low)/2 + low
		// get middle value across the rows and colums
		val := matrix[mid/cols][mid%cols]
		if val == target {
			return true
		} else if val < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return false
}
