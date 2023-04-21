package argshifter

import "strings"

type argShifter struct {
	args         []string
	cursor       int
	charCursor   int
	argumentType ArgType
}

func (as *argShifter) GetArgumentType() ArgType {
	return as.argumentType
}

func (as *argShifter) Peek() string {
	if as.argumentType != ShortOption {
		return as.args[as.cursor]
	}

	return string(([]rune(as.args[as.cursor]))[as.charCursor])
}

func (as *argShifter) toNext() {
	as.cursor++
	if as.cursor >= len(as.args) {
		return
	}

	prevType := as.argumentType
	curRunes := []rune(as.args[as.cursor])

	if len(curRunes) < 1 {
		as.argumentType = Invalid
	} else {
		if curRunes[0] == '-' {
			if len(curRunes) > 1 && curRunes[1] == '-' {
				as.argumentType = LongOption
			} else {
				as.argumentType = ShortOption
			}
		} else {
			if prevType == Root {
				as.argumentType = Command
			} else {
				as.argumentType = Data
			}
		}
	}

	if as.argumentType == ShortOption {
		as.charCursor = 1
	} else {
		as.charCursor = 0
	}
}

func (as *argShifter) Shift() (string, bool) {
	if as.cursor >= len(as.args) {
		return "", false
	}

	arg := as.args[as.cursor]
	argRunes := []rune(arg)
	if as.argumentType == ShortOption {
		defer func() {
			as.charCursor++
			if as.charCursor >= len(argRunes) {
				as.toNext()
			}
		}()
		return string(argRunes[as.charCursor]), true
	}

	defer as.toNext()
	return strings.TrimLeft(arg, "-"), true
}

func (as *argShifter) Reset() {
	as.cursor = 0
	as.charCursor = 0
	as.argumentType = Root
}

func (as *argShifter) GetArgs() []string {
	return as.args
}

func NewArgShifter(args []string) ArgShifter {
	return &argShifter{
		args:         args,
		cursor:       0,
		charCursor:   0,
		argumentType: Root,
	}
}
