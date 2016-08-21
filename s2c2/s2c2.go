package main

import (
    ecb "aesecb"
    b64 "encoding/base64"
    "io/ioutil"
    "fmt"
)

func main() {
    dat, err := ioutil.ReadFile("./10.txt")
    if err != nil {
        panic(err)
    }
    phrase := make([]byte, b64.StdEncoding.DecodedLen(len(dat)))
    b64.StdEncoding.Decode(phrase, []byte(dat))
    if err != nil {
        panic(err)
    }
    out := ecb.Aes_Ecb_Cbc_Decrypt(phrase, []byte("YELLOW SUBMARINE"),
                                   []byte{0 ,0 ,0 ,0 ,0 ,0 ,0 ,0 ,0 ,0 ,0 ,0 ,0 ,0 ,0 ,0})

    fmt.Println(string(out))
}
