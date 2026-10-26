func maxArea(heights []int) int {
	if len(heights) <= 1 {
		return 0
	}

	var maxA, l, r int = 0 ,0, len(heights) - 1

	for  l < r {
		h := min(heights[l], heights[r])
		maxA = max(maxA, (r-l)*h)

		if heights[l] < heights[r] {
			l++
		} else {
			r--
		}
	}

	return maxA
}
