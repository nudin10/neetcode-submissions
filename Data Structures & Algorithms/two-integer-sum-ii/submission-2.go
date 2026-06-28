func twoSum(numbers []int, target int) []int {
	x, y := 0, len(numbers)-1

	for x < y {
		sum := numbers[x]+numbers[y]
		if sum == target {
			break
		}
		if sum < target {
			x++
		}
		if sum > target {
			y--
		}
	}

	return []int{x+1, y+1}
}
