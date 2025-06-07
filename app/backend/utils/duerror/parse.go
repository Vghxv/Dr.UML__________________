package duerror

type ParsingError struct {
	msg string
}

func (e *ParsingError) Error() string {
	return e.msg
}

func NewParsingError(msg string) *ParsingError {
	return &ParsingError{msg: msg}
}
