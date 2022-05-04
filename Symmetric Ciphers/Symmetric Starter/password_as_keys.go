package main

import (
	"bufio"
	"bytes"
	"crypto/aes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

var URL string = "http://aes.cryptohack.org"

func main() {
	resp, _ := http.Get(URL + "/passwords_as_keys/encrypt_flag/")
	res := make(map[string]string)
	json.NewDecoder(resp.Body).Decode(&res)
	resp.Body.Close()

	file, _ := os.Open("./words.txt")
	file_rd := bufio.NewReader(file)
	for line := []byte{0}; len(line) > 0; {
		line, _ = file_rd.ReadBytes('\n')
		line = bytes.TrimSpace(line)

		password := md5.Sum(line)
		cipher, _ := aes.NewCipher(password[:])
		ct, _ := hex.DecodeString(res["ciphertext"])
		pt := make([]byte, len(ct))
		n := cipher.BlockSize()
		for i := 0; i < len(ct); i += n {
			cipher.Decrypt(pt[i:i+n], ct[i:i+n])
		}

		if bytes.HasPrefix(pt, []byte("crypto")) {
			fmt.Println(string(pt))
			break
		}
	}
}
