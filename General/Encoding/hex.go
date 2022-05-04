package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	hexstr := "63727970746f7b596f755f77696c6c5f62655f776f726b696e675f776974685f6865785f737472696e67735f615f6c6f747d"
	bytes, _ := hex.DecodeString(hexstr)
	fmt.Println(string(bytes))
}
