package lib

func CompareIgnoreCase(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		char1 := s1[i]
		char2 := s2[i]

		if char1 != char2 {
			if char1 >= 'A' && char1 <= 'Z' {
				char1 += 'a' - 'A'
			} else if char1 >= 'a' && char1 <= 'z' {
				char1 -= 'a' - 'A'
			}

			if char1 != char2 {
				return false
			}
		}
	}

	return true
}
