package main

import "fmt"
import h "hexConversion"
import b64 "encoding/base64"

func main() {
	hex := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	var res string

	for i := 0; i < len(hex); i += 2 {
		res += string(h.HexASCIIToByte(string(hex[i]) + string(hex[i+1])))
	}
	fmt.Println(b64.StdEncoding.EncodeToString([]byte(res)))
}
