package utils

// SwapInt swap two given integers
func SwapInt(a, b *int) {
	*a, *b = *b, *a
}
