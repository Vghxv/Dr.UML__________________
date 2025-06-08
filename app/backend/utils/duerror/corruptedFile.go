package duerror

type CorruptedFile struct {
	msg string
}

func NewCorruptedFile(msg string) *CorruptedFile {
	return &CorruptedFile{msg: msg}
}

func (e *CorruptedFile) Error() string {
	return e.msg
}
