package main

import (
	"flag" // for flag.Arg
	"gols/functions"
	"io"
	"os"
)

func main() {
	var writer io.Writer
	args := flag.Args()
	useColor := functions.IsTerminal(os.Stdout)

	functions.SimpleLS(writer, args, useColor)
}
