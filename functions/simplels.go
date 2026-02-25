package functions

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
)

func SimpleLS(w io.Writer, args []string, useColor bool) {
	// Separate arguments into files and dirs
	// Sorting is handled by Partition()
	files, dirs := Partition(args)

	// Print filenames
	for _, file := range files {
		w.Write([]byte(file + "\n"))
	}

	if len(dirs) > 0 {
		w.Write([]byte("\n"))
	}

	multipleDirs := len(dirs) > 1

	// Print dirs and their contents
	for dirnum, dir := range dirs {
		tgts, err := os.ReadDir(dir)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Fprintf: %v\n", err)
			continue
		}

		if multipleDirs {
			w.Write([]byte(filepath.Base(dir) + ":\n"))
		}

		for _, tgt := range dirFilter(tgts) {
			w.Write([]byte(tgt.Name() + "\n"))
		}
		if dirnum < len(dirs)-1 {
			w.Write([]byte("\n"))
		}
	}
}

func Partition(args []string) (files []string, dirs []string) {
	files = []string{}
	dirs = []string{}
	for _, arg := range args {
		info, err := os.Lstat(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fprintf: %v\n", err)
			continue
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
