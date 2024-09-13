package main

import (
	"encoding/hex"
	"fmt"
)

func xorBytes(b1, b2 []byte) []byte {
	if len(b1) != len(b2) {
		panic("Byte slices must be of the same length")
	}
	result := make([]byte, len(b1))
	for i := range b1 {
		result[i] = b1[i] ^ b2[i]
	}
	return result
}

func main() {
	plaintext1 := "This is a known message!" // P1 (in plaintext)
	ciphertext1Hex := "a469b1c502c1cab966965e50425438e1bb1b5f9037a4c159" // C1 (hex)
	ciphertext2Hex := "bf73bcd3509299d566c35b5d450337e1bb175f903fafc159" // C2 (hex)

	// Convert P1 to bytes
	plaintext1Bytes := []byte(plaintext1)

	// Decode hex ciphertexts to bytes
	ciphertext1Bytes, err := hex.DecodeString(ciphertext1Hex)
	if err != nil {
		panic(err)
	}
	ciphertext2Bytes, err := hex.DecodeString(ciphertext2Hex)
	if err != nil {
		panic(err)
	}

	// Calculate the keystream
	keystream := xorBytes(ciphertext1Bytes, plaintext1Bytes)

	// Recover the plaintext P2
	plaintext2Bytes := xorBytes(ciphertext2Bytes, keystream)

	// Convert recovered plaintext P2 to string
	plaintext2 := string(plaintext2Bytes)

	fmt.Printf("Recovered Plaintext P2: %s\n", plaintext2)
}
