package main

import "github.com/bradfitz/slice"
import (
	ch "characterFrequency"
	"fmt"
	hex "hexConversion"
    "bufio"
    "os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type cyphInfos struct {
    xorKey int
    chFrequency int
    data string
}

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

func printCyphInfos(info cyphInfos) {
    fmt.Printf("Freq: %d Key: %d Data: %s\n", info.chFrequency, info.xorKey, info.data)
}

func main() {
    finish := make(chan struct{})
    rawDataIn := make(chan []byte, 10)
    unxoredDataIn := make(chan cyphInfos, 10)
    dat, err := readLines("./4.txt")
    check(err)

    go func(in chan []byte) {
        var key byte

        for d := range in {
            for j := 0; j < 256; j++ {
                key = byte(j)
                var tmp string
                for i := 0; i < len(d); i++ {
                    tmp += string(key ^ d[i])
                }
                var data cyphInfos
                data.xorKey = j
                data.data = tmp
                unxoredDataIn <- data
            }
        }
        close(unxoredDataIn)
        finish <- struct{}{} // Signal main that we are done
    }(rawDataIn)

    go func() {
        decrypt(unxoredDataIn)
        finish <- struct{}{}
    }()

    for _, l := range dat {
        rawDataIn <- hex.HexASCIIStringToBytes(l)
    }
    close(rawDataIn)

    for i := 0; i < 2; i++ { // Wait for goroutine to finish --'
        <- finish
    }
}

func decrypt(dataIn chan cyphInfos) {
    finish := make(chan struct{})
    freqDataIn := make(chan cyphInfos, 10)


    go func(in chan cyphInfos) {
        for s := range in {
            s.chFrequency = ch.CharacterFrequency(s.data)
            freqDataIn <- s
        }
        close(freqDataIn)
        finish <- struct{}{} // Signal main that we are done
    }(dataIn)

    go func(in chan cyphInfos) {
        allInfo := make([]cyphInfos, 0)
        for data := range in {
            allInfo = append(allInfo, data)
        }

        slice.Sort(allInfo[:], func(i, j int) bool {
            return allInfo[i].chFrequency < allInfo[j].chFrequency
        })

        for i := 0; i < 20; i++ {
            printCyphInfos(allInfo[i])
        }

        finish <- struct{}{} // Signal main that we are done
    }(freqDataIn)

    for i := 0; i < 2; i++ { // Wait for goroutine to finish --'
        <- finish
    }
}
