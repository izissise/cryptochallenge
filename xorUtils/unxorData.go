package xorUtils

import "github.com/bradfitz/slice"
import (
    ch "characterFrequency"
)

type xorInfo struct {
  key int
  data []byte
}

type xorFreq struct {
  freq int
  info xorInfo
}

func UnxorKey(data []byte, key int) []byte {
    bkey := byte(key)
    out := make([]byte, len(data))
    for i := 0; i < len(data); i++ {
        out[i] = bkey ^ data[i]
    }
    return out
}

func UnxorAllKey(data []byte, out chan xorInfo) {
    for i := 0; i < 256; i++ {
        out <- xorInfo{i, UnxorKey(data, i)}
    }
}

func XorFreq(in chan xorInfo, out chan xorFreq) {
    for s := range in {
        chFrequency := ch.CharacterFrequency(string(s.data))
        out <- xorFreq{chFrequency, s}
    }
}

func AllKeyFreq(data []byte) ([]xorFreq) {
    finish := make(chan struct{})
    nbGoRoutine := 0
    unxor := make(chan xorInfo)
    freqs := make(chan xorFreq)
    freqMap := make([]xorFreq, 0)

    nbGoRoutine++
    go func() {
        XorFreq(unxor, freqs)
        close(freqs)
        finish <- struct{}{}
    }()

    nbGoRoutine++
    go func() {
        UnxorAllKey(data, unxor)
        close(unxor)
        finish <- struct{}{}
    }()

    nbGoRoutine++
    go func() {
        for freq := range freqs {
            freqMap = append(freqMap, freq)
        }
        finish <- struct{}{}
    }()

    for i := 0; i < nbGoRoutine; i++ { // Wait for goroutine to finish --'
        <- finish
    }

    slice.Sort(freqMap[:], func(i, j int) bool {
        return freqMap[i].freq < freqMap[j].freq
    })
    return freqMap
}

func BestFreqKey(data []byte) byte {
    return byte((AllKeyFreq(data)[0]).info.key)
}

func UnxorDataWithKey(data []byte, key []byte) []byte {
    out := make([]byte, len(data))
    pos := 0

    for i := 0; i < len(data); i++ {
        out = append(out, data[i]^key[pos])
        pos = (pos + 1) % len(key)
    }
    return out
}

