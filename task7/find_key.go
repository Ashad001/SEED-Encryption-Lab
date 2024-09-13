package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"os"
)

// Pad the key to 16 bytes with '#'
func padKey(key string) []byte {
	paddedKey := make([]byte, 16)
	copy(paddedKey, key)
	for i := len(key); i < 16; i++ {
		paddedKey[i] = '#'
	}
	return paddedKey
}

// Perform AES-128-CBC decryption
func aesDecrypt(ciphertext []byte, key []byte, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < aes.BlockSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	plaintext := make([]byte, len(ciphertext))
	mode.CryptBlocks(plaintext, ciphertext)

	// Remove padding
	paddingLen := int(plaintext[len(plaintext)-1])
	if paddingLen > aes.BlockSize || paddingLen > len(plaintext) {
		return nil, fmt.Errorf("padding size error")
	}
	plaintext = plaintext[:len(plaintext)-paddingLen]

	return plaintext, nil
}

func main() {
	ciphertextHex := "44b7e8d03b633600ea3aaabf42807627fabce9af94e4ba8c46da8efb360d0d2457ccb004914cb8dc99d59584d50eae4f"
	ivHex := "010203040506070809000a0b0c0d0e0f"
	plaintext := "my stories dont end until I stop running"

	ciphertext, _ := hex.DecodeString(ciphertextHex)
	iv, _ := hex.DecodeString(ivHex)

	// Open the wordlist file
	wordlist, err := os.Open("words.txt")
	if err != nil {
		fmt.Println("Error opening wordlist:", err)
		return
	}
	defer wordlist.Close()

	var word string
	key := make([]byte, 16)

	scanner := bufio.NewScanner(wordlist)
	for scanner.Scan() {
		word = scanner.Text()
		if len(word) == 0 {
			continue
		}

		// Pad the word to 16 characters using '#'
		key = padKey(word)

		decryptedtext, err := aesDecrypt(ciphertext, key, iv)
		if err != nil {
			continue
		}

		if string(decryptedtext) == plaintext {
			fmt.Println("Key found:", word)
			break
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading wordlist:", err)
	}
}
