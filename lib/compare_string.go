package lib

func CompareIgnoreCase(s1, s2 string) bool {
	var i, j int = 0, 0
	var len1, len2 int = len(s1), len(s2)

	for i < len1 && j < len2 {
		if s1[i] == ' ' {
			i++
			continue
		} else if s2[j] == ' ' {
			j++
			continue
		}

		char1 := s1[i]
		char2 := s2[j]

		if char1 >= 'A' && char1 <= 'Z' {
			char1 += 'a' - 'A'
		} else if char1 >= 'a' && char1 <= 'z' {
			char1 -= 'a' - 'A'
		}

		if char1 != char2 {
			return false
		}

		i++
		j++
	}

	for i < len1 && s1[i] == ' ' {
		i++
	}
	for j < len2 && s2[j] == ' ' {
		j++
	}

	return i == len1 && j == len2
}
