func characterReplacement(s string, k int) int {
	var (
		longest, maxCharFreq, l int
		cf = make(map[byte]int)
	)

	for r := 0; r < len(s); r++ {
		cf[s[r]]++
		maxCharFreq = max(maxCharFreq, cf[s[r]])

		if (r-l+1) - maxCharFreq > k {
			cf[s[l]]--
			l++
		}

		longest = max(longest, r-l+1)
	}

	return longest
}