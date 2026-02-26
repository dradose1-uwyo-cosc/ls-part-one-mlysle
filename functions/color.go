package functions

import (
	"io"
)

type color string

const (
	reset color = "\033[0m"
	blue  color = "\033[34m"
	green color = "\033[32m"
)

func (c color) ColorPrint(w io.Writer, s string) {
	io.WriteString(w, string(c)+s+string(reset))
}
