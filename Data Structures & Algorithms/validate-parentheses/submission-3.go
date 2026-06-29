func isValid(s string) bool {
    var (
		brackets = make([]rune, 0)
	)

	if len(s) % 2 > 0 {
		return false
	}

	for _, c := range s {
		l := len(brackets)
		switch c {
			case '(', '{', '[':
			brackets = append(brackets, c)
			case ')', '}', ']':
			if l < 1 {
				return false
			}
			if c == ')' && brackets[l-1] != '(' {
				return false
			}
			if c == '}' && brackets[l-1] != '{' {
				return false
			}
			if c == ']' && brackets[l-1] != '[' {
				return false
			}
			brackets = brackets[:l-1]
		}
	}
	if len(brackets) > 0 {
		return false
	}
	return true
}
