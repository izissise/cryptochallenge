package main

import (
// 	ch "characterFrequency"
	"fmt"
	hex "hexConversion"
// 	"sort"
)

func main() {
	cyph := hex.HexASCIIStringToBytes("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	var key byte
	allStrings := make([]string, 255)

	for j := 0; j < 255; j++ {
		key = byte(j)
		var tmp string
		for i := 0; i < len(cyph); i++ {
			tmp += string(key ^ cyph[i])
		}
		allStrings[j] = tmp
		fmt.Println(j, " -> ", tmp)
	}

// 	freqMap := make(map[int]string)
// 	for _, s := range allStrings {
// 		perc := ch.CharacterFrequency(s)
// 		freqMap[perc] = s
// 	}
//
// 	var keys []int
// 	for k := range freqMap {
// 		keys = append(keys, k)
// 	}
// 	sort.Ints(keys)
//
// 	for _, k := range keys {
// 		fmt.Println(k, " -> ", freqMap[k])
// 	}
}
