package colors

import (
	"strconv"
	"strings"
)

func hexToDecimal(color string) int {
	value, _ := strconv.ParseInt(strings.ReplaceAll(color, "#", ""), 16, 64)

	return int(value)
}

// Bot Colors

func Monetizze() int {
	return hexToDecimal("#ffffff")
}
func Maite() int {
	return hexToDecimal("#8e5ba0")
}

// Special Colors

func Error() int {
	return hexToDecimal("#e30e0e")
}
func Success() int {
	return hexToDecimal("#07ed1e")
}

// Normal Colors

func White() int {
	return hexToDecimal("#ffffff")
}
