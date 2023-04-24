package argparser_test

import (
	"testing"

	"github.com/chardon55/go-argparser"
)

func TestArguments1(t *testing.T) {
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

			data := op.Data()

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

			is := op.IncrementSwitches()
			if is["s"] != 1 || is["search"] != 1 {
				t.Errorf("Switch not matched")
				return nil
			}

			return nil
		}).
		Complete().
		Parse(args)

	if err != nil {
		panic(err)
	}
}

func TestArguments2(t *testing.T) {
	args1 := []string{"conda", "install", "pandas", "seaborn", "jupyter", "-y"}

	err := argparser.NewArgParser().
		AddCommand("install").
		AddBooleanSwitch('y', "yes").
		SetExecutor(func(op argparser.Operation, args []string) error {
			if args1[0] != args[0] {
				t.Errorf("Incorrect program name. Expected: %s; actual: %s", args[0], args1[0])
				return nil
			}

			data := op.Data()

			actualLength := len(data)
			if actualLength < 3 {
				t.Errorf("Uncompleted data value collection. Length: %d", actualLength)
				return nil
			}

			for i, val := range data {
				if val != args[2+i] {
					t.Errorf("Unmatched value. Expected: %s; actual: %s", args[2+i], val)
					return nil
				}
			}

			bs := op.BooleanSwitches()

			if !bs["y"] || !bs["yes"] {
				t.Errorf("Switch not matched")
				return nil
			}

			return nil
		}).
		Complete().
		Parse(args1)

	if err != nil {
		panic(err)
	}
}
