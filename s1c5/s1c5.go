package main

import (
	"fmt"
	hex "hexConversion"
)

func main() {
	phrase := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	key := "ICE"
	pos := 0
	res := ""

	for i := 0; i < len(phrase); i++ {
		res = res + string(phrase[i]^key[pos])
		pos = (pos + 1) % len(key)
	}

	fmt.Println(hex.SliceByteToHexASCII([]byte(res)))
}
