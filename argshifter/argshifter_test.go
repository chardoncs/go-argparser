package argshifter_test

import (
	"testing"

	"github.com/chardon55/go-argparser/argshifter"
)

func processArgs(t *testing.T, args []string, expectedTypes []argshifter.ArgType, expectedValues []string) {
	length := len(expectedTypes)
	shifter := argshifter.NewArgShifter(args)

	argType := shifter.GetArgumentType()
	val, prs := shifter.Shift()

	var i int
	for i = 0; prs; i++ {
		if i >= length {
			t.Errorf("Index exceeded! %d", i)
			return
		}

		expectedType := expectedTypes[i]
		if argType != expectedType {
			t.Errorf("Type not matched! Expected: %d; actual: %d", expectedType, argType)
			return
		}

		expectedVal := expectedValues[i]
		if val != expectedVal {
			t.Errorf("Value not matched! Expected: %s; actual: %s", expectedVal, val)
			return
		}

		argType = shifter.GetArgumentType()
		val, prs = shifter.Shift()
	}

	if i < length {
		t.Errorf("Shifting not complete! %d/%d", i, length)
	}
}

func TestBasicArgs(t *testing.T) {
	args := []string{"apt", "install", "neofetch", "sl", "cowsay"}
	expectedTypes := []argshifter.ArgType{
		argshifter.Root,
		argshifter.Command,
		argshifter.Data,
		argshifter.Data,
		argshifter.Data,
	}
	expectedValues := args
	processArgs(t, args, expectedTypes, expectedValues)
}

func TestShortOptions1(t *testing.T) {
	args := []string{"apt", "install", "-y", "neofetch", "sl", "cowsay"}
	expectedTypes := []argshifter.ArgType{
		argshifter.Root,
		argshifter.Command,
		argshifter.ShortOption,
		argshifter.Data,
		argshifter.Data,
		argshifter.Data,
	}
	expectedValues := []string{"apt", "install", "y", "neofetch", "sl", "cowsay"}
	processArgs(t, args, expectedTypes, expectedValues)
}

func TestShortOptions2(t *testing.T) {
	args := []string{"tar", "-xzvf", "blahblah.tar.gz"}
	expectedTypes := []argshifter.ArgType{
		argshifter.Root,
		argshifter.ShortOption,
		argshifter.ShortOption,
		argshifter.ShortOption,
		argshifter.ShortOption,
		argshifter.Data,
	}
	expectedValues := []string{"tar", "x", "z", "v", "f", "blahblah.tar.gz"}
	processArgs(t, args, expectedTypes, expectedValues)
}

func TestShortOptions3(t *testing.T) {
	args := []string{"pacman", "-Syu"}
	expectedTypes := []argshifter.ArgType{
		argshifter.Root,
		argshifter.ShortOption,
		argshifter.ShortOption,
		argshifter.ShortOption,
	}
	expectedValues := []string{"pacman", "S", "y", "u"}
	processArgs(t, args, expectedTypes, expectedValues)
}

func TestShortOptions4(t *testing.T) {
	args := []string{"tar", "-x", "-z", "-vf", "blahblah.tar.gz"}
	expectedTypes := []argshifter.ArgType{
		argshifter.Root,
		argshifter.ShortOption,
		argshifter.ShortOption,
		argshifter.ShortOption,
		argshifter.ShortOption,
		argshifter.Data,
	}
	expectedValues := []string{"tar", "x", "z", "v", "f", "blahblah.tar.gz"}
	processArgs(t, args, expectedTypes, expectedValues)
}

func TestLongOption1(t *testing.T) {
	args := []string{"pacman", "--sync", "--refresh", "--update"}
	expectedTypes := []argshifter.ArgType{
		argshifter.Root,
		argshifter.LongOption,
		argshifter.LongOption,
		argshifter.LongOption,
	}
	expectedValues := []string{"pacman", "sync", "refresh", "update"}
	processArgs(t, args, expectedTypes, expectedValues)
}

// func TestLongOption2(t *testing.T) {
// 	args := []string{"pacman", "-S", "--ignore=sl", "--refresh", "--update"}
// 	expectedTypes := []argshifter.ArgType{
// 		argshifter.Root,
// 		argshifter.ShortOption,
// 		argshifter.LongOption,
// 		argshifter.Data,
// 		argshifter.LongOption,
// 		argshifter.LongOption,
// 	}
// 	expectedValues := []string{"pacman", "S", "ignore", "sl", "refresh", "update"}
// 	processArgs(t, args, expectedTypes, expectedValues)
// }
