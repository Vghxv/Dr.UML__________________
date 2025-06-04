package duerror

type CorruptedFile struct {
	msg string
}

func NewCorruptedFile(msg string) *FileIOError {
	return &FileIOError{msg: msg}
}

func (e *CorruptedFile) Error() string {
	return e.msg
}
