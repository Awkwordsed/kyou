package main

import (
  "crypto/aes"
  "crypto/cipher"
  "crypto/rand"
  "encoding/hex"
  "fmt"
  "io/ioutil"
  "os"
  "io"
)

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

    
    // Create the file using the CreateFile function
    file, err := os.Create("filename.txt")
    if err != nil {
      fmt.Println("Error creating the file:", err)
      return
    }
    defer file.Close()

    // Write the variable's value to the file
    err = ioutil.WriteFile("filename.txt", []byte(enc), 0644)
    if err != nil {
      fmt.Println("Error writing to the file:", err)
      return
    }

    fmt.Println("Variable written to the file successfully.")


}
