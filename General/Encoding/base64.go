package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	hexstr := "72bca9b68fc16ac7beeb8f849dca1d8a783e8acf9679bf9269f7bf"
	bytes, _ := hex.DecodeString(hexstr)
	fmt.Println(base64.RawStdEncoding.EncodeToString(bytes))
}
