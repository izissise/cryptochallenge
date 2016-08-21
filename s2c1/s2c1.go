package main

import (
    ecb "aesecb"
    "fmt"
)

func main() {
    fmt.Println(string(ecb.PadDataPKCS([]byte("YELLOW SUBMARINE"), 20)))
}
