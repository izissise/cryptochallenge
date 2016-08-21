package main

import "crypto/aes"
import (
    "io/ioutil"
    b64 "encoding/base64"
    "fmt"
    ecb "aesecb"
)

func main() {
    dat, err := ioutil.ReadFile("./7.txt")
    if err != nil {
       panic(err)
    }
    phrase := make([]byte, b64.StdEncoding.DecodedLen(len(dat)))
    b64.StdEncoding.Decode(phrase, []byte(dat))

    out := make([]byte, len(phrase))

    b, err := aes.NewCipher([]byte("YELLOW SUBMARINE"))
    if err != nil {
        panic(err)
    }

    ecb.NewECBDecrypter(b).CryptBlocks(out, phrase)
    fmt.Println(string(out))
}
