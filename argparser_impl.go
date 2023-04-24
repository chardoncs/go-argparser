package argparser

import (
	"github.com/chardon55/go-argparser/argshifter"
	"github.com/chardon55/go-argparser/exceptions"
)

type switchName struct {
	Short string
	Long  string
}

// ArgParser implementation
type argParser struct {
	ops      map[string]*operation
	commands map[string]*operation
}

func (parser *argParser) newOperation() *operation {
	return &operation{
		parent:             parser,
		booleanSwitches:    make(map[string]bool),
		incrementSwitches:  make(map[string]uint),
		dataSwitches:       make(map[string]string),
		data:               []string{},
		switchLongShortMap: make(map[string]rune),
		switchShortLongMap: make(map[rune]string),
	}
}

func (parser *argParser) AddCommand(name string) Operation {
	op := parser.newOperation()

	parser.commands[name] = op
	return op
}

func (parser *argParser) AddOperation(short rune, long string) Operation {
	op := parser.newOperation()

	parser.ops[string(short)] = op
	parser.ops[long] = op
	return op
}

func (parser *argParser) Parse(args []string) error {
	shifter := argshifter.NewArgShifter(args)

	_, _, prs := shifter.Shift()
	if !prs {
		return exceptions.NewEmptyArgumentError()
	}

	// Get operation
	operationString, argType, prs := shifter.Shift()
	if !prs {
		return exceptions.NewNoOperationError()
	}

	var op *operation

	switch argType {
	case argshifter.Command:
		op, prs = parser.commands[operationString]
		if !prs {
			return exceptions.NewInvalidOperationError(operationString, exceptions.COMMAND)
		}

	case argshifter.ShortOption, argshifter.LongOption:
		op, prs = parser.ops[operationString]
		if !prs {
			return exceptions.NewInvalidOperationError(operationString, exceptions.OPERATION)
		}

	default:
		return exceptions.NewNoOperationError()
	}

	var dataSwitchNamePtr *switchName

	value, argType, valPrs := shifter.Shift()

	for valPrs {
		switch argType {
		case argshifter.Data:
			if dataSwitchNamePtr != nil {
				if len(dataSwitchNamePtr.Short) > 0 {
					op.DataSwitches()[dataSwitchNamePtr.Short] = value
				}
				op.dataSwitches[dataSwitchNamePtr.Long] = value

				dataSwitchNamePtr = nil
			} else {
				op.data = append(op.data, value)
			}

		case argshifter.ShortOption:
			longOp := op.switchShortLongMap[[]rune(value)[0]]

			if _, prs := op.booleanSwitches[value]; prs {
				op.booleanSwitches[value] = true
				op.booleanSwitches[longOp] = op.booleanSwitches[value]

			} else if _, prs := op.incrementSwitches[value]; prs {
				op.incrementSwitches[value]++
				op.incrementSwitches[longOp] = op.incrementSwitches[value]

			} else if _, prs := op.dataSwitches[value]; prs {
				dataSwitchNamePtr = &switchName{
					Short: value,
					Long:  longOp,
				}
			}

		case argshifter.LongOption:
			shortOpRune, opPrs := op.switchLongShortMap[value]
			shortOp := string(shortOpRune)

			if _, prs := op.booleanSwitches[value]; prs {
				op.booleanSwitches[value] = true

				if opPrs {
					op.booleanSwitches[shortOp] = op.booleanSwitches[value]
				}
			} else if _, prs := op.incrementSwitches[value]; prs {
				op.incrementSwitches[value]++

				if opPrs {
					op.incrementSwitches[shortOp] = op.incrementSwitches[value]
				}
			} else if _, prs := op.dataSwitches[value]; prs {
				var finalShort string
				if opPrs {
					finalShort = shortOp
				} else {
					finalShort = ""
				}

				dataSwitchNamePtr = &switchName{
					Short: finalShort,
					Long:  value,
				}
			}

		default:
			// Nothing to do
		}

		value, argType, valPrs = shifter.Shift()
	}

	return op.execute(args)
}

func NewArgParser() ArgParser {
	return &argParser{
		ops:      make(map[string]*operation),
		commands: make(map[string]*operation),
	}
}
