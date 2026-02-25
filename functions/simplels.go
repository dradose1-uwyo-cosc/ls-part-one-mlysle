package functions

import (
	"fmt"
	"io"
	"os"
	"sort"
)

func SimpleLS(w io.Writer, args []string, useColor bool) {
	files, dirs := Partition(args)
}

func Partition(args []string) (files []string, dirs []string) {
	files = []string{}
	dirs = []string{}
	for _, arg := range args {
		info, err := os.Lstat(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fprintf: %v\n", err)
		}
		if info.IsDir() {
			dirs = append(dirs, arg)
		} else {
			files = append(files, arg)
		}
	}
	sort.Strings(files)
	sort.Strings(dirs)
	return files, dirs
}
