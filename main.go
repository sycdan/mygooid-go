package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/alexflint/go-arg"
	Slug "github.com/gosimple/slug"
	"github.com/sycdan/mygooid-go/internal/renji"
	"github.com/sycdan/mygooid-go/internal/utils"
)

type Args struct {
	Name       string `arg:"positional" help:"Full legal name"`
	Secret     string `arg:"positional" help:"Memorable secret"`
	Purpose    string `arg:"positional" help:"What is the password for (default: General)"`
	Characters string `arg:"--characters, -c" help:"All characters allowed in the password (default: A-Z, a-z, 0-9, !@#%:-_+=?)"`
}

var reader *bufio.Reader

var ANCHORS = []string{"aqw1", "btx2", "cry3", "djs4", "ez5*", "fmn6", "gop7", "hkl8", "iuv9"}

const PURPOSE = "General"
const UPPERCASE = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const LOWERCASE = "abcdefghijklmnopqrstuvwxyz"
const NUMBERS = "0123456789"
const SYMBOLS = "!@#%:-_+=?"

func MakeGooid(args Args) string {
	prefixedHashes := slugifyAndHashFields(args, []string{"Name", "Secret", "Purpose"})
	return utils.HashText(strings.Join(prefixedHashes, ""))
}

func init() {
	fmt.Println("Gooid Generator\n---------------")
	reader = bufio.NewReader(os.Stdin)
}

func main() {
	args := Args{
		Characters: UPPERCASE + LOWERCASE + NUMBERS + SYMBOLS,
		Purpose:    PURPOSE,
	}

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

// Slugify and then hash the requested fields from the passed object, then return the hashes in the same order.
func slugifyAndHashFields(object interface{}, fields []string) []string {
	var hashes []string
	value := reflect.ValueOf(object)
	for _, fieldName := range fields {
		field := value.FieldByName(fieldName)
		if field.IsValid() {
			value := utils.ToString(field.Interface())
			slug := slugify(value)
			hash := utils.HashText(slug)
			hashes = append(hashes, hash)
		}
	}
	return hashes
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
