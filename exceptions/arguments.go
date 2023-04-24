package exceptions

type EmptyArgumentError struct {
	msg string
}

func (err *EmptyArgumentError) Error() string {
	return err.msg
}

func NewEmptyArgumentError() *EmptyArgumentError {
	return &EmptyArgumentError{
		msg: "the argument list is empty",
	}
}
