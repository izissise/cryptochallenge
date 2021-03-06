package aesecb

import "crypto/cipher"
import "crypto/aes"
import "reflect"
import xor "xorUtils"

type ecb struct {
    b         cipher.Block
    blockSize int
}

func newECB(b cipher.Block) *ecb {
    return &ecb{
        b:         b,
        blockSize: b.BlockSize(),
    }
}

type ecbEncrypter ecb

func NewECBEncrypter(b cipher.Block) cipher.BlockMode {
    return (*ecbEncrypter)(newECB(b))
}

func (x *ecbEncrypter) BlockSize() int { return x.blockSize }

func (x *ecbEncrypter) CryptBlocks(dst, src []byte) {
    if len(src)%x.blockSize != 0 {
        panic("crypto/cipher: input not full blocks")
    }
    if len(dst) < len(src) {
        panic("crypto/cipher: output smaller than input")
    }
    for len(src) > 0 {
        x.b.Encrypt(dst, src[:x.blockSize])
        src = src[x.blockSize:]
        dst = dst[x.blockSize:]
    }
}

type ecbDecrypter ecb

func NewECBDecrypter(b cipher.Block) cipher.BlockMode {
    return (*ecbDecrypter)(newECB(b))
}

func (x *ecbDecrypter) BlockSize() int { return x.blockSize }

func (x *ecbDecrypter) CryptBlocks(dst, src []byte) {
    if len(src)%x.blockSize != 0 {
        panic("crypto/cipher: input not full blocks")
    }
    if len(dst) < len(src) {
        panic("crypto/cipher: output smaller than input")
    }
    for len(src) > 0 {
        x.b.Decrypt(dst, src[:x.blockSize])
        src = src[x.blockSize:]
        dst = dst[x.blockSize:]
    }
}

func EcbOccurence(data []byte, blockSize int) int {
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
    return k
}

func PadDataPKCS(data []byte, sizeMultiple int) []byte {
   padding := byte(4)
   toAdd := sizeMultiple - (len(data) % sizeMultiple)
   if toAdd == sizeMultiple {
       return data
   }
   for i := 0; i < toAdd; i++ {
       data = append(data, padding)
   }
   return data
}

func Aes_Ecb_Cbc_Decrypt(data []byte, key []byte, ivs []byte) []byte {
    blockSize := 16
    b, err := aes.NewCipher(key)
    if err != nil {
        panic(err)
    }
    dec := NewECBDecrypter(b)

    var final []byte
    prev := ivs
    for i := 0; i < (len(data) / blockSize); i++ {
        block := data[blockSize * i : blockSize * (i + 1)]
        out := make([]byte, len(block))
        dec.CryptBlocks(out, block)
        final = append(final, xor.XorBlock(prev, out)...)
        prev = block
    }
    return final
}

func Aes_Ecb_Cbc_Encrypt(data []byte, key []byte, ivs []byte) []byte {
    blockSize := 16
    b, err := aes.NewCipher(key)
    if err != nil {
        panic(err)
    }
    enc := NewECBEncrypter(b)

    var final []byte
    prev := ivs
    for i := 0; i < (len(data) / blockSize); i++ {
        block := data[blockSize * i : blockSize * (i + 1)]
        block = xor.XorBlock(block, prev)
        out := make([]byte, len(block))
        enc.CryptBlocks(out, block)
        final = append(final, out...)
        prev = out
    }
    return final
}




