func searchMatrix(matrix [][]int, target int) bool {
	for _, row := range matrix {
		x := len(row)-1
		if row[x] >= target && row[0] <= target {
			low, high := 0, x
			for low <= high {
				mid := (high-low)/2 + low
				if row[mid] == target {
					return true
				} else if row[mid] < target{
					low = mid + 1
				} else {
					high = mid - 1
				}
			}	
		} 
	}

	return false
}
