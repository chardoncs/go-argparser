package exceptions

// No operation error
type NoOperationError struct {
	msg string
}

func (err *NoOperationError) Error() string {
	return err.msg
}

func NewNoOperationError() *NoOperationError {
	return &NoOperationError{
		msg: "No operation specified",
	}
}

type OperationType int

const (
	OPERATION OperationType = iota + 1
	COMMAND
)

// Invalid operation error
type InvalidOperationError struct {
	msg, op string
	_type   OperationType
}

func (err *InvalidOperationError) Error() string {
	return err.msg
}

func (err *InvalidOperationError) Operation() string {
	return err.op
}

func (err *InvalidOperationError) Type() OperationType {
	return err._type
}

func NewInvalidOperationError(operation string, _type OperationType) *InvalidOperationError {
	return &InvalidOperationError{
		msg:   "Invalid operation",
		op:    operation,
		_type: _type,
	}
}
