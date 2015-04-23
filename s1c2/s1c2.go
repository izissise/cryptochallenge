package main

import "fmt"
import hex "hexConversion"

func main() {
	key1 := hex.HexASCIIStringToBytes("1c0111001f010100061a024b53535009181c")
	key2 := hex.HexASCIIStringToBytes("686974207468652062756c6c277320657965")
	var res string
	var res2 string

	for i := 0; i < len(key1); i++ {
		res += string(key1[i] ^ key2[i])
	}

	for i := 0; i < len(key1); i++ {
		res2 += hex.ByteToHex(res[i])
	}

	fmt.Println(res2)
}
