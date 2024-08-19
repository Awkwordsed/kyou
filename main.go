package main

import (
  "crypto/aes"
  "crypto/cipher"
  "crypto/rand"
  "encoding/hex"
  "fmt"
  "os"
  "io"
)

func shstat() {
	const myFilePath = "shadow"

	if _, err := os.Stat(myFilePath); err != nil {
		if os.IsNotExist(err) {
      // Create the file using the CreateFile function
      file, err := os.Create("shadow")
      if err != nil {
        fmt.Println("Error creating the file:", err)
        return
      }
      defer file.Close()

      fmt.Println("Made new shadow file")

	} else {
			panic(err)
		}
	} else {
	    fmt.Println("The shadow file exist") 
	}
}

func main() {
    // Data Preparation
    var data string
    fmt.Println("Input: ")
    fmt.Scanln(&data)
    plaintext := []byte(data)

    // Key Generation
    key := make([]byte, 32)
    if _, err := rand.Reader.Read(key); err != nil {
        fmt.Println("Error generating random encrypting key", err)
        return
    }
    
    // AES Block Cipher Creation
    block, err := aes.NewCipher(key)
    if err != nil {
        fmt.Println("Error creating AES block cipher", err)
        return
    }

    // GCM Mode Setup
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        fmt.Println("Error setting up GCM mode", err)
        return
    }

    // Nonce Generation
    nonce := make([]byte, gcm.NonceSize())
    if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
        fmt.Println("Error generating nonce", err)
        return
    }

    // Encrypting the data using GCM mode
    ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)

    // Convert to Hexadecimal
    enc := hex.EncodeToString(ciphertext)

    shstat()
    
    f, err := os.OpenFile("shadow", os.O_APPEND|os.O_WRONLY, 0600)
    if err != nil {
        panic(err)
    }
    defer f.Close()

    if _, err = f.WriteString(enc); err != nil {
        panic(err)
    }


    fmt.Println("Save successful.")


}
