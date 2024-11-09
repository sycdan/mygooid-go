package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/alexflint/go-arg"
	// "github.com/sycdan/mygooid-go/internal/maverick"
	"github.com/sycdan/mygooid-go/internal/renji"
	"github.com/sycdan/mygooid-go/internal/utils"
)

type Args struct {
	Name      string `arg:"positional" help:"Full legal name"`
	Secret    string `arg:"positional" help:"Memorable secret"`
	Purpose   string `arg:"positional" help:"What is the password for (default: General)"`
	Length    int    `arg:"positional" help:"Length of the password (default: 8)"`
	Uppercase string `arg:"--uppercase, -u" help:"All uppercase characters allowed in the password (default: A-Z)"`
	Lowercase string `arg:"--lowercase, -l" help:"All lowercase characters allowed in the password (default: a-z)"`
	Number    string `arg:"--number, -n" help:"All number characters allowed in the password (default: 0-9)"`
	Special   string `arg:"--special, -s" help:"All special characters allowed in the password (default: !@#%:-_+=?)"`
}

func NewArgs() *Args {
	return &Args{
		// name and secret wil be set later
		Purpose:   DEFAULT_PURPOSE,
		Length:    DEFAULT_LENGTH,
		Uppercase: DEFAULT_UPPERCASES,
		Lowercase: DEFAULT_LOWERCASES,
		Number:    DEFAULT_NUMBERS,
		Special:   DEFAULT_SYMBOLS,
	}
}

var reader *bufio.Reader

var ANCHORS = []string{"aqw1", "btx2", "cry3", "djs4", "ez5*", "fmn6", "gop7", "hkl8", "iuv9"}

const DEFAULT_PURPOSE = "General"
const DEFAULT_UPPERCASES = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const DEFAULT_LOWERCASES = "abcdefghijklmnopqrstuvwxyz"
const DEFAULT_NUMBERS = "0123456789"
const DEFAULT_SYMBOLS = "!@#%:-_+=?"
const DEFAULT_LENGTH = 8

func MakeGooid(args *Args) string {
	prefixedHashes := slugifyAndHashFields(args, []string{"Name", "Secret", "Purpose"})
	return utils.HashText(strings.Join(prefixedHashes, ""))
}

func GeneratePassword(args *Args) string {
	var password string

	gooid := MakeGooid(args)
	rng := renji.NewRenji(gooid)

	// Add one of each required (nonblank) character group to the password.

	if args.Uppercase != "" {
		password += string(args.Uppercase[rng.Intn(len(args.Uppercase))])
	}
	if args.Lowercase != "" {
		password += string(args.Lowercase[rng.Intn(len(args.Lowercase))])
	}
	if args.Number != "" {
		password += string(args.Number[rng.Intn(len(args.Number))])
	}
	if args.Special != "" {
		password += string(args.Special[rng.Intn(len(args.Special))])
	}

	// shuffler := maverick.NewMaverick([]string{}, rng)
	// shuffler.Shuffle()

	return password
}

func init() {
	fmt.Println("Gooid Generator\n---------------")
	reader = bufio.NewReader(os.Stdin)
}

func parseArgs() *Args {
	args := NewArgs()
	arg.MustParse(args)

	if args.Purpose == "" {
		utils.Die("Purpose must be specified")
	}

	if args.Length < 4 {
		utils.Die("Length must be >= 4")
	}

	return args
}

func main() {
	args := parseArgs()

	if args.Name == "" {
		args.Name = readInput("Enter your full legal name")
	}

	if args.Secret == "" {
		args.Secret = readInput("Enter a memorable secret")
	}

	password := GeneratePassword(args)
	fmt.Println(password)
}

// Slugify and then hash the requested fields from the passed object, then return the hashes in the same order.
func slugifyAndHashFields(object interface{}, fields []string) []string {
	var hashes []string
	value := reflect.ValueOf(object)

	// If the value is a pointer, dereference it to get the actual struct.
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	for _, fieldName := range fields {
		field := value.FieldByName(fieldName)
		if field.IsValid() {
			value := utils.ToString(field.Interface())
			slug := utils.Slugify(value)
			hash := utils.HashText(slug)
			hashes = append(hashes, hash)
		}
	}
	return hashes
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
