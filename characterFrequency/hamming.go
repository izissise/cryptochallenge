package characterFrequency

// HammingDistance return the number of bits that need to be changed to go from x to y
func HammingDistance(x byte, y byte) int {
	dist := 0
	val := x ^ y

	for val != 0 {
		dist++
		val &= val - 1
	}
	return dist
}

// HammingDistanceStr gie hamming distance between two strings
func HammingDistanceStr(str1 string, str2 string) int {
	res := 0
	mlen := 0
	if len(str1) > len(str2) {
		mlen = len(str1)
	} else {
		mlen = len(str2)
	}

	for i := 0; i < mlen; i++ {
		var b1 byte
		var b2 byte
		if i < len(str1) {
			b1 = str1[i]
		}
		if i < len(str2) {
			b2 = str2[i]
		}
		res += HammingDistance(b1, b2)
	}
	return res
}
