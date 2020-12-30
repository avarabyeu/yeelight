package color

import (
	"errors"
	"image/color"
)

func HexStringToRgbInt(s string) (value int, err error) {
	c := color.RGBA{}
	c.A = 0xff

	if s[0] != '#' {
		return value, errors.New("missing '#' at the beginning of the string")
	}

	switch len(s) {
	case 7:
		c.R = hexToByte(s[1])<<4 + hexToByte(s[2])
		c.G = hexToByte(s[3])<<4 + hexToByte(s[4])
		c.B = hexToByte(s[5])<<4 + hexToByte(s[6])
	case 4:
		c.R = hexToByte(s[1]) * 17
		c.G = hexToByte(s[2]) * 17
		c.B = hexToByte(s[3]) * 17
	default:
		err = errors.New("incorrect value")
	}

	value = 256*256*int(c.R) + 256*int(c.G) + int(c.B)

	return
}

func hexToByte(b byte) byte {
	switch {
	case b >= '0' && b <= '9':
		return b - '0'
	case b >= 'a' && b <= 'f':
		return b - 'a' + 10
	case b >= 'A' && b <= 'F':
		return b - 'A' + 10
	}

	return 0
}
