package argparser

// Argument parser interface
type ArgParser interface {

	// Add a command
	AddCommand(name string) Operation

	// Add an operation
	AddOperation(short rune, long string) Operation

	// Parse the arguments
	Parse(args []string) error
}
