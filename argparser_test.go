package argparser_test

import (
	"testing"

	"github.com/chardon55/pacaurgo/core/cli/argparser"
)

func TestCollection1(t *testing.T) {
	args := []string{"pacman", "-Ss", "sl", "cowsay"}

	err := argparser.NewArgParser().
		AddOperation('S', "sync").
		AddIncrementSwitch('s', "search").
		AddLongIncrementSwitch("color").
		SetExecutor(func(op argparser.Operation, args1 []string) error {
			if args1[0] != args[0] {
				t.Errorf("Incorrect program name. Expected: %s; actual: %s", args[0], args1[0])
				return nil
			}

			data := op.GetEndData()

			actualLength := len(data)
			if actualLength < 2 {
				t.Errorf("Uncompleted data value collection. Length: %d", actualLength)
				return nil
			}

			for i, val := range data {
				if val != args[2+i] {
					t.Errorf("Unmatched value. Expected: %s; actual: %s", args[2+i], val)
					return nil
				}
			}

			return nil
		}).
		Complete().
		Parse(args)

	if err != nil {
		panic(err)
	}
}
