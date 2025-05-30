package attribute

type Textstyle int

const (
	Bold                    = 1 << iota // 0x1
	Italic                  = 1 << iota // 0x2
	Underline               = 1 << iota // 0x4
	supportedTextStyleFlags = Bold | Italic | Underline
)

var AllTextstyleTypes = []struct {
	Value  Textstyle
	TSName string
}{
	{Bold, "Bold"},
	{Italic, "Italic"},
	{Underline, "Underline"},
}
