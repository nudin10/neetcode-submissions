func hasDuplicate(nums []int) bool {
	numM := make(map[int]struct{})

	for _, num := range nums {
		if _, found := numM[num]; found {
			return true
		}
		numM[num] = struct{}{}
	}

	return false
}