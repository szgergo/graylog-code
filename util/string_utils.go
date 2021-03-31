package util

func StringArrayEquals(a []string, b[]string) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil && b != nil {
		return false
	}
	if a != nil && b == nil {
		return false
	}
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		isEqual := false
		for j := range b {
			if a[i] == b[j] {
				isEqual = true
			}
		}
		if !isEqual {
			return false
		}
	}
	return true
}
