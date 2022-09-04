package src

import (
	"fmt"
	"strings"
)

func RGB(r, g, b int) string {
	return strings.ToUpper(fmt.Sprintf("%02x%02x%02x", scope(r), scope(g), scope(b)))
}

func scope(rgb int) int {
	if rgb < 0 {
		rgb = 0
	}
	if rgb > 255 {
		rgb = 255
	}
	return rgb
}
