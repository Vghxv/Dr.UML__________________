package Attribute

type TextStyle int

const (
	bold      = 1 << iota // 0x001
	italic    = 1 << iota // 0x010
	underline = 1 << iota // 0x100
)
