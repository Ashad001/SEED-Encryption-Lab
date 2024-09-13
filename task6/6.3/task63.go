package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"log"
)

func xorBytes(b1, b2 []byte) []byte {
	if len(b1) != len(b2) {
		log.Fatal("Byte slices must be of the same length")
	}
	result := make([]byte, len(b1))
	for i := range b1 {
		result[i] = b1[i] ^ b2[i]
	}
	return result
}

// padding function similar to PKCS7 padding used in AES
func pad(input string, blockSize int) []byte {
	padding := blockSize - (len(input) % blockSize)
	padded := append([]byte(input), bytes.Repeat([]byte{byte(padding)}, padding)...)
	return padded
}

func main() {
	// Original plaintext message "Yes" padded to 16 bytes (like PKCS7 padding)
	plaintext := pad("Yes", 16)
	pt1Hex := hex.EncodeToString(plaintext)

	// IVs as provided in the container
	iv1Hex := "5ddffd2abd1fb16589d2f667d9aade2c"
	iv2Hex := "3a56f26ebd1fb16589d2f667d9aade2c"

	// convert plaintext and IVs to byte slices
	pt1Bytes, err := hex.DecodeString(pt1Hex)
	if err != nil {
		log.Fatal(err)
	}
	iv1Bytes, err := hex.DecodeString(iv1Hex)
	if err != nil {
		log.Fatal(err)
	}
	iv2Bytes, err := hex.DecodeString(iv2Hex)
	if err != nil {
		log.Fatal(err)
	}

	// XOR the plaintext with respective IVs
	tempList := xorBytes(pt1Bytes, iv1Bytes)
	ip2List := xorBytes(tempList, iv2Bytes)

	fmt.Println(string(ip2List))
}
