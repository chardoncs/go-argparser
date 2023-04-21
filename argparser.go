package argparser

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
	GetBooleanSwitches() map[string]bool

	// Get increment switches
	GetIncrementSwitches() map[string]uint

	// Get data switches
	GetDataSwitches() map[string]string

	// Get data that do not belong to any data switches
	GetEndData() []string
}

// Argument parser customized for PacAUR-Go
type ArgParser interface {

	// Add an operation
	AddOperation(short rune, long string) Operation

	// Parse the arguments
	Parse(args []string) error
}
