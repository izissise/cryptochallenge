package main

import "github.com/bradfitz/slice"
import (
	ch "characterFrequency"
	b64 "encoding/base64"
	"io/ioutil"
    "fmt"
    xor "xorUtils"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type keySizeDistance struct {
  hamming float32
  keysize int
}

func byteDivideBlocks(data []byte, size int, out chan []byte) {
  var nb int
  nb = len(data) / size
  for i := 0; i < nb; i++ {
      out <- data[(size * i):(size * (i + 1))]
  }
}

func singleKeyBlock(blocks chan []byte, blockSize int, meningfulBytes int) [][]byte {
    keysBlock := make([][]byte, blockSize)
    var i int
    for d := range blocks {
        if (i > meningfulBytes) {
            break
        }
        for j := 0; j < blockSize; j++ {
            keysBlock[j] = append(keysBlock[j], d[j])
        }
        i++;
    }
    return keysBlock;
}

func findKeys(data []byte, keySizeMin int, keySizeMax int, outKeys chan []byte) {
    finish := make(chan struct{})
    nbGoRoutine := 0

    hamdistMap := make([]keySizeDistance, 0)
    for keysize := keySizeMin; keysize < keySizeMax; keysize++ {
        hd := ch.HammingDistanceBlockSize(data, keysize, false)
        hamdistMap = append(hamdistMap, keySizeDistance{hd, keysize})
    }

    slice.Sort(hamdistMap[:], func(i, j int) bool {
        return hamdistMap[i].hamming < hamdistMap[j].hamming
    })

    for i := 0; i < 3; i++ {
        kSize := hamdistMap[i].keysize
        blocks := make(chan []byte, 10)
        nbGoRoutine += 1
        go func() {
            finalKey := make([]byte, kSize)
            keyBlocks := singleKeyBlock(blocks, kSize, 200000)
            for i := 0; i < len(keyBlocks); i++ {
                finalKey[i] = xor.BestFreqKey(keyBlocks[i]) // Change algorithm here if not xor
            }
            outKeys <- finalKey
            finish <- struct{}{}
        }()
        nbGoRoutine += 1
        go func() {
          byteDivideBlocks(data, kSize, blocks)
          close(blocks)
          finish <- struct{}{}
        }()

    }

    for i := 0; i < nbGoRoutine; i++ { // Wait for goroutine to finish --'
        <- finish
    }
}

func decrypt(data []byte) ([]byte, []byte) {
    finish := make(chan struct{})
    nbGoRoutine := 0

    keys := make(chan []byte)

    nbGoRoutine += 1
    go func() {
        findKeys(data, 3, 41, keys)
        close(keys)
        finish <- struct{}{}
    }()

    var final []byte
    freq := 9999999
    var finalKey []byte
    for k := range keys {
        decrypted := xor.UnxorDataWithKey(data, k)
        tmpFreq := ch.CharacterFrequency(string(decrypted))
        if tmpFreq < freq {
            freq = tmpFreq
            final = decrypted
            finalKey = k
        }
    }

    for i := 0; i < nbGoRoutine; i++ { // Wait for goroutine to finish --'
        <- finish
    }

    return finalKey, final
}

func main() {
	dat, err := ioutil.ReadFile("./6.txt")
	check(err)
	phrase := make([]byte, b64.StdEncoding.DecodedLen(len(dat)))
	b64.StdEncoding.Decode(phrase, []byte(dat))

    key, data := decrypt(phrase)
    fmt.Printf("Key: %s\n%s\n", string(key), string(data))
}
