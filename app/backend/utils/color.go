package utils

import (
	"fmt"
	"strconv"
)

type Color struct {
	R, G, B uint8
}

func (c *Color) ToHexString() string {
	return "#" + fmt.Sprintf("%02X%02X%02X", c.R, c.G, c.B)
}

// FromHexString parse string like "#FF00FF"
func FromHexString(s string) Color {
	if len(s) != 7 || s[0] != '#' {
		return Color{}
	}
	hex := s[1:]
	i, err := strconv.ParseInt(hex, 16, 32)
	if err != nil {
		return Color{}
	}
	return Color{
		R: uint8((i >> 16) & 0xFF),
		G: uint8((i >> 8) & 0xFF),
		B: uint8(i & 0xFF),
	}
}
