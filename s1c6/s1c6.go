package main

import (
	ch "characterFrequency"
	b64 "encoding/base64"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("./6.txt")
	check(err)
	phrase64 := make([]byte, b64.StdEncoding.DecodedLen(len(dat)))
	b64.StdEncoding.Decode(phrase64, []byte(dat))
	phrase := string(phrase64)

	hamdistMap := make(map[float32]int)
	var hamdistSort []float32
	for keysize := 2; keysize < 41; keysize++ {
		block1 := phrase[0:keysize]
		block2 := phrase[keysize:keysize]
		hd := float32(ch.HammingDistanceStr(block1, block2)) / float32(keysize)
		hamdistMap[hd] = keysize
		hamdistSort = append(hamdistSort, hd)
	}

	var keysizeList []int
	for _, kk := range hamdistSort {
		keysizeList = append(keysizeList, hamdistMap[kk])
	}

	for _, keySize := range keysizeList {
		blockKeyTmp := make([]string, keySize)
		for j := 0; j < keySize; j++ {
			for i := 0; i < len(phrase); i += keySize {
				blockKeyTmp[j] += string(phrase[i+j])
			}
		}

		allStrings := make([]string, 255)
		for j := 0; j < 255; j++ {
			key := byte(j)
			var tmp string
			for i := 0; i < len(blockKeyTmp); i++ {
				tmp += string(key ^ blockKeyTmp[i])
			}
			allStrings[j] = tmp
		}

		for i := 0; i < len(allStrings); i++ {

		}
	}

}
