func checkInclusion(s1 string, s2 string) bool {
	var cm1, cm2 [26]int
	for _, c := range s1 {
		cm1[c-'a']++
	}

	var l int
	for r := 0; r < len(s2); r++ {
		cm2[s2[r]-'a']++
		if r-l+1 > len(s1) {
			cm2[s2[l]-'a']--
			l++
		}
		if cm1 == cm2 {
			return true
		}
	}

	return false
}
