func lengthOfLongestSubstring(s string) int {
	var (
		maxL int
		chars = make(map[byte]int)
		l, r int
	)

	for r < len(s) {
		if c, ok := chars[s[r]]; ok {
			l = max(l, c+1)
		}

		chars[s[r]]=r
		maxL = max(maxL, r-l+1)
		
		r++
	}

	return maxL
}
