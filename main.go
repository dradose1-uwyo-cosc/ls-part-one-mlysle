package main

import (
	"flag" // for flag.Arg
	"gols/functions"
	"io"
	"os"
)

func main() {
	flag.Parse()
	var writer io.Writer = os.Stdout
	args := flag.Args()
	useColor := functions.IsTerminal(os.Stdout)

	functions.SimpleLS(writer, args, useColor)
}
