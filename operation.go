package argparser

import "fmt"

type Operation interface {

	// Add a boolean switch
	AddBooleanSwitch(short rune, long string) Operation

	// Add a boolean switch with only a long name
	AddLongBooleanSwitch(long string) Operation

	// Add an increment switch
	AddIncrementSwitch(short rune, long string) Operation

	// Add an increment switch with only a long name
	AddLongIncrementSwitch(long string) Operation

	// Add an data switch
	AddDataSwitch(short rune, long string) Operation

	// Add an data switch with only a long name
	AddLongDataSwitch(long string) Operation

	// Set the executor
	//
	// NOTE: There may be only one executor.
	// If you set the executor multiple times, only the latest update will be preserved.
	SetExecutor(e func(op Operation, args []string) error) Operation

	// Complete the configuration and go back to its parent
	Complete() ArgParser

	// Get regular boolean switches
	BooleanSwitches() map[string]bool

	// Get increment switches
	IncrementSwitches() map[string]uint

	// Get data switches
	DataSwitches() map[string]string

	// Get data that do not belong to any data switches
	Data() []string
}

// Operation implementation
type operation struct {
	parent ArgParser

	executor func(op Operation, args []string) error
	data     []string

	booleanSwitches   map[string]bool
	incrementSwitches map[string]uint
	dataSwitches      map[string]string

	switchLongShortMap map[string]rune
	switchShortLongMap map[rune]string
}

func (op *operation) AddBooleanSwitch(short rune, long string) Operation {
	op.booleanSwitches[string(short)] = false
	op.booleanSwitches[long] = false

	op.switchLongShortMap[long] = short
	op.switchShortLongMap[short] = long

	return op
}

func (op *operation) AddLongBooleanSwitch(long string) Operation {
	op.booleanSwitches[long] = false
	return op
}

func (op *operation) AddIncrementSwitch(short rune, long string) Operation {
	op.incrementSwitches[string(short)] = 0
	op.incrementSwitches[long] = 0

	op.switchLongShortMap[long] = short
	op.switchShortLongMap[short] = long

	return op
}

func (op *operation) AddLongIncrementSwitch(long string) Operation {
	op.incrementSwitches[long] = 0
	return op
}

func (op *operation) AddDataSwitch(short rune, long string) Operation {
	op.dataSwitches[string(short)] = ""
	op.dataSwitches[long] = ""

	op.switchLongShortMap[long] = short
	op.switchShortLongMap[short] = long

	return op
}

func (op *operation) AddLongDataSwitch(long string) Operation {
	op.dataSwitches[long] = ""
	return op
}

func (op *operation) SetExecutor(e func(Operation, []string) error) Operation {
	op.executor = e
	return op
}

func (op *operation) Complete() ArgParser {
	return op.parent
}

func (op *operation) execute(args []string) error {
	if op.executor == nil {
		return fmt.Errorf("no executors")
	}

	return op.executor(op, args)
}

func (op *operation) BooleanSwitches() map[string]bool {
	return op.booleanSwitches
}

func (op *operation) IncrementSwitches() map[string]uint {
	return op.incrementSwitches
}

func (op *operation) DataSwitches() map[string]string {
	return op.dataSwitches
}

func (op *operation) Data() []string {
	return op.data
}
