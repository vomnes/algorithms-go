package algorithms

// IterativeSearch is an iteration search function
// O()
func IterativeSearch(data []int, length, toFind int) bool {
	for index := 0; index < length; index++ {
		if data[index] == toFind {
			return true
		}
	}
	return false
}

// BinarySearch is the binary search algorithm implementation
// O(log n)
func BinarySearch(data []int, length, toFind int) bool {
	var index, l, r int
	l = 0
	r = length - 1
	for l <= r {
		index = (l + r) / 2
		if data[index] < toFind { // 1st Part
			l = index + 1
		} else if data[index] > toFind { // 2nd Part
			r = index - 1
		} else {
			return true
		}
	}
	return false
}
