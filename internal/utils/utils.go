package utils

import (
	"crypto/sha256"
	"fmt"
	"os"
)

// Exit the program with a message.
func Die(reason string) {
	fmt.Println(reason)
	os.Exit(1)
}

// Compute the SHA256 hash of a string and return it in hex.
func HashText(text string) string {
	sum := sha256.Sum256([]byte(text))
	return fmt.Sprintf("%x", sum)
}

// Convert any value to its default string representation.
func ToString[T any](value T) string {
	return fmt.Sprintf("%v", value)
}
