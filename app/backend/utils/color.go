package utils

type Color struct {
	R, G, B uint8
}

func (c *Color) ToHex() int {
	return int(c.R)<<16 | int(c.G)<<8 | int(c.B)
}

func FromHex(i int) Color {
	return Color{
		R: uint8((i >> 16) & 0xFF),
		G: uint8((i >> 8) & 0xFF),
		B: uint8(i & 0xFF),
	}
	// return Color{R: uint8(i >> 16), G: uint8(i >> 8), B: uint8(i)}
}
