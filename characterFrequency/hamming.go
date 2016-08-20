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

func HammingDistanceStr(str1 string, str2 string) int {
    return HammingDistanceByte([]byte(str1), []byte(str2))
}

// HammingDistanceStr gie hamming distance between two strings
func HammingDistanceByte(str1 []byte, str2 []byte) int {
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

func HammingDistanceBlockSize(data []byte, blockSize int, full bool) float32 {
    nb := 4
    if full {
        nb = len(data) / blockSize
    }
    var total float32
    var block1 []byte
    var block2 []byte
    block1 = data[0:blockSize]
    for i := 0; i < nb; i++ {
        block2 = block1
        block1 = data[(blockSize * i):blockSize * (i + 1)]
        total += float32(HammingDistanceByte(block1, block2)) / float32(blockSize)
    }
    hd := total / float32(nb - 1)
    return hd
}
