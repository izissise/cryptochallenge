package main

import hex "hexConversion"
import (
	"bufio"
	"fmt"
	"os"
    "reflect"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	dat, err := readLines("./8.txt")
	if err != nil {
		panic(err)
	}
	blockSize := 16

	probl := -1
	occur := 0
	for i := 0; i < len(dat); i++ {
		data := hex.HexASCIIStringToBytes(dat[i])
		var block1 []byte
		var block2 []byte
        var k int
		for i := 0; i < len(data) / blockSize; i++ {
            block2 = data[(blockSize * i) : blockSize * (i + 1)]
			for j := 1; j < (len(data) - (blockSize * i)) / blockSize; j++ {
                block1 = data[(blockSize * (i + j)) : blockSize * (i + j + 1)]
                if reflect.DeepEqual(block1, block2) {
                    k++
                }
            }
		}
		if k > occur {
			probl = i
			occur = k
		}
	}
	fmt.Printf("Id: %d K: %d Data: %s\n", probl, occur, dat[probl])
}
