package argshifter

// Argument type enumeration
type ArgType int

const (
	Invalid ArgType = iota - 1
	Root
	Command
	Data
	ShortOption
	LongOption
	OptionWithData
)

// Argument shifter interface
type ArgShifter interface {

	// Get the type of the current argument
	ArgumentType() ArgType

	// Get the current argument but do not shift
	Peek() string

	// Get the current argument and switch to the next one
	//
	// Returns:
	//   string - the argument before the shift
	//   ArgType (int) - Argument type
	//   bool - is any argument available
	Shift() (string, ArgType, bool)

	// Walk to the next without checking the value
	Walk() bool

	// Reset the status of the argument shifter
	Reset()

	// Get arguments
	GetArgs() []string
}
