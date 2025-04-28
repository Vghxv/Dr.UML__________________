package Attribute

type TextStyle int

const (
	Bold      = 1 << iota // 0x001
	Italic    = 1 << iota // 0x010
	Underline = 1 << iota // 0x100
)
