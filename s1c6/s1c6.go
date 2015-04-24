package main

import (
	ch "characterFrequency"
	b64 "encoding/base64"
	"fmt"
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

	var keysize int

	for keysize = 2; keysize < 41; keysize++ {
		block1 := phrase[0:keysize]
		block2 := phrase[keysize:keysize]
		hd := float32(ch.HammingDistanceStr(block1, block2)) / float32(keysize)
		fmt.Println(hd, " -> ", keysize)
	}
	/*
		for i := 0; i < len(cyph); i++ {
			res = res + string(cyph[i]^key[pos])
			pos = (pos + 1) % len(key)
		}

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
		}*/
}
