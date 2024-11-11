package utils

import (
	"crypto/sha256"
	"fmt"
	"os"

	"github.com/gosimple/slug"
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

// Check if a specific string is in a slice.
func Contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

// Convert "Some Text" to "some-text".
func Slugify(text string) string {
	return slug.Make(text)
}

// Convert a string to a generic interface slice of characters.
func StringToInterfaceSlice(text string) []interface{} {
	var result []interface{}
	for _, runeValue := range text {
		result = append(result, string(runeValue))
	}
	return result
}
