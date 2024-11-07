package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strings"

	"github.com/alexflint/go-arg"
	Slug "github.com/gosimple/slug"
	"github.com/sycdan/mygooid-go/internal/renji"
	"github.com/sycdan/mygooid-go/internal/utils"
)

type Args struct {
	Name   string `arg:"positional" help:"Full legal name"`
	Secret string `arg:"positional" help:"Memorable secret"`
}

var reader *bufio.Reader

var ANCHORS = []string{"aqw", "btx", "cry", "djs", "ez", "fmn", "gop", "hkl", "iuv"}

const UPPERCASE = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const LOWERCASE = "abcdefghijklmnopqrstuvwxyz"
const NUMBERS = "0123456789"
const SYMBOLS = "!@#%:-_+=?"

func MakeGooid(args Args) string {
	prefixedHashes := slugifyAndHashObject(args)
	sort.Strings(prefixedHashes)
	return utils.HashText(strings.Join(prefixedHashes, ""))
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
	rng := renji.NewRenji(gooid)
	rng.Float64()
	fmt.Println(gooid, rng)
}

// Hash every value in the passed object then hash the hashes (prefixed with their property names).
func slugifyAndHashObject(obj interface{}) []string {
	var prefixedHashes []string
	value := reflect.ValueOf(obj)
	if value.Kind() == reflect.Struct {
		for i := 0; i < value.NumField(); i++ {
			field := value.Type().Field(i)
			value := value.Field(i)
			slug := slugify(utils.ToString(value.Interface()))
			hash := utils.HashText(slug)
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
		utils.Die("Error reading input:" + err.Error())
	}
	return strings.TrimSpace(text)
}
