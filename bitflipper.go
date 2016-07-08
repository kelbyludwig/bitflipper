package main

import (
	"encoding/base64"
	"flag"
	"fmt"
)

var b64Data string

func main() {
	flag.StringVar(&b64Data, "b", "b64", "base64 encoded data to bitflip")
	flag.Parse()

	bytes, err := base64.StdEncoding.DecodeString(b64Data)

	if err != nil {
		fmt.Println("[ERR] Invalid base64 input data. Did you forget the 'b' flag?")
		return
	}

	bits := []byte{0x01, 0x02, 0x04, 0x08, 0x10, 0x20, 0x40, 0x80}

	modBytes := make([]byte, len(bytes))
	num := 0
	for i, b := range bytes {
		for _, bi := range bits {
			modByte := bi ^ b
			copy(modBytes, bytes)
			modBytes[i] = modByte
			modb64 := base64.StdEncoding.EncodeToString(modBytes)
			fmt.Println(modb64)
			num = num + 1
		}
	}

	fmt.Printf("Made %d modifications\n", num)

}
