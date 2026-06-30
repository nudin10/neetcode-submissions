func search(nums []int, target int) int {
	var low, high = 0, len(nums)-1

	for low <= high {
		mid := (high-low)/2 + low
		x := nums[mid]

		if x == target {
			return mid
		} else if x < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}
