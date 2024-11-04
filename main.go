package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strings"

	"github.com/alexflint/go-arg"
	Slug "github.com/gosimple/slug"
)

type Args struct {
	Name   string `arg:"positional" help:"Full legal name"`
	Secret string `arg:"positional" help:"Memorable secret"`
}

var reader *bufio.Reader

func MakeGooid(args Args) string {
	prefixedHashes := slugifyAndHashObject(args)
	sort.Strings(prefixedHashes)
	return hashText(strings.Join(prefixedHashes, ""))
}

func init() {
	fmt.Println("Gooid Generator\n---------------")
	reader = bufio.NewReader(os.Stdin)
}

func main() {
	var args Args
	arg.MustParse(&args)

	if args.Name == "" {
		args.Name = readInput("Enter your full legal name")
	}

	if args.Secret == "" {
		args.Secret = readInput("Enter a memorable secret")
	}

	gooid := MakeGooid(args)
	fmt.Println(gooid)
}

func toString[T any](value T) string {
	return fmt.Sprintf("%v", value)
}

// Hash every value in the passed object then hash the hashes (prefixed with their property names).
func slugifyAndHashObject(obj interface{}) []string {
	var prefixedHashes []string
	value := reflect.ValueOf(obj)
	if value.Kind() == reflect.Struct {
		for i := 0; i < value.NumField(); i++ {
			field := value.Type().Field(i)
			value := value.Field(i)
			slug := slugify(toString(value.Interface()))
			hash := hashText(slug)
			prefixedHashes = append(prefixedHashes, slugify(field.Name+"-"+hash))
		}
	}
	return prefixedHashes
}

// Convert "Some Text" to "some-text".
func slugify(text string) string {
	return Slug.Make(text)
}

// Read all user input up until they press enter.
func readInput(prompt string) string {
	fmt.Println(prompt + ": ")
	text, err := reader.ReadString('\n')
	if err != nil {
		die("Error reading input:" + err.Error())
	}
	return strings.TrimSpace(text)
}

// Compute the SHA256 hash of a string and return it in hex.
func hashText(text string) string {
	sum := sha256.Sum256([]byte(text))
	return fmt.Sprintf("%x", sum)
}

func die(reason string) {
	fmt.Println(reason)
	os.Exit(1)
}
