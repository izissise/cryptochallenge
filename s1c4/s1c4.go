package main

import (
	ch "characterFrequency"
	"fmt"
	hex "hexConversion"
	"io/ioutil"
	"sort"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("./4.txt")
	check(err)
	cyph := hex.HexASCIIStringToBytes(string(dat))

	var key byte
	allStrings := make([]string, 255)

	for j := 0; j < 255; j++ {
		key = byte(j)
		var tmp string
		for i := 0; i < len(cyph); i++ {
			tmp += string(key ^ cyph[i])
		}
		allStrings[j] = tmp
	}

	var subString []string
	for _, s := range allStrings {
		print := true
		var tmp string
		for i := 0; i < len(s); i++ {
			if !(strconv.IsPrint(int32(s[i]))) {
				if print {
					subString = append(subString, tmp)
				}
				print = false
				tmp = ""
			} else {
				print = true
				tmp = tmp + string(s[i])
			}
		}
	}

	freqMap := make(map[int]string)
	for _, s := range subString {
		perc := ch.CharacterFrequency(s)
		freqMap[perc] = s
	}

	var keys []int
	for k := range freqMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		fmt.Println(k, " -> ", freqMap[k])
	}
}
