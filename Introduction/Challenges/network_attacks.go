package main

import (
	"encoding/json"
	"fmt"
	"net/textproto"
)

func main() {
	conn, _ := textproto.Dial("tcp", "socket.cryptohack.org:11112")
	defer conn.Close()

	for l := " "; len(l) > 0; {
		l, _ = conn.ReadLine()
	}

	conn.PrintfLine(`{"buy":"flag"}`)

	res_line, _ := conn.ReadLine()
	res := make(map[string]any)
	json.Unmarshal([]byte(res_line), &res)
	fmt.Println(res["flag"])
}
