func lengthOfLongestSubstring(s string) int {
	var (
		maxL, l , r int
		chars = make(map[byte]int)
	)

	for r < len(s) {
		if i, ok := chars[s[r]]; ok {
			l = max(l, i+1)
		}

		chars[s[r]] = r
		maxL = max(maxL, r-l+1)
		r++
	}

	return maxL
}
