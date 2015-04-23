package hexConversion

// HexASCIIToByte convert two character in hexa to its value
func HexASCIIToByte(hexaByte string) byte {
	var f = func(j byte) byte {
		if j >= '0' && j <= '9' {
			return j - '0'
		}
		if j >= 'a' && j <= 'f' {
			return j - 'a' + 10
		}
		return j - 'A' + 10
	}
	return f(hexaByte[0])*16 + f(hexaByte[1])
}

// HexASCIIStringToBytes convert a string of axii hex value to a byte slice
func HexASCIIStringToBytes(hex string) []byte {
	res := make([]byte, len(hex)/2)

	for i := 0; i < len(hex); i += 2 {
		res[i/2] = HexASCIIToByte(string(hex[i]) + string(hex[i+1]))
	}
	return res
}

// ByteToHex convert a byte to its two character hexa
func ByteToHex(b byte) string {
	b1 := b / 16
	b2 := b % 16

	var f = func(j byte) byte {
		if j < 10 {
			return j + '0'
		}
		return j - 10 + 'a'
	}
	return string(f(b1)) + string(f(b2))
}

// SliceByteToHexASCII convert a byte slice to its hexa ascii representation
func SliceByteToHexASCII(bs []byte) string {
	res := ""

	for _, c := range bs {
		res = res + ByteToHex(c)
	}
	return res
}
