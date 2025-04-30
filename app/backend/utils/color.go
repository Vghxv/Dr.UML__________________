package utils

import (
	"fmt"
)

type Color struct {
	R, G, B uint8
}

func (c Color) ToHex() string {
	return "#" + fmt.Sprintf("%02X%02X%02X", c.R, c.G, c.B)
}
