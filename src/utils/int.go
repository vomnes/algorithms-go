package utils

// SwapInt swap two given integers
func SwapInt(a, b *int) {
	*a, *b = *b, *a
}

// ArrayIntEquals compare two int arrays return true if equals
func ArrayIntEquals(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
