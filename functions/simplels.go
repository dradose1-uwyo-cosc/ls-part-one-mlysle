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
		info, err := os.Lstat(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fprintf: %v\n", err)
			continue
		}
		printTarget(w, useColor, info)
	}

	if len(dirs) > 0 && len(files) > 0 {
		io.WriteString(w, "\n")
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
			io.WriteString(w, filepath.Base(dir)+":\n")
		}

		for _, tgt := range dirFilter(tgts) {
			// info, err := tgt.Info()
			path := filepath.Join(dir, tgt.Name())
			info, err := os.Lstat(path)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Fprintf: %v\n", err)
				continue
			}
			printTarget(w, useColor, info)
		}

		if dirnum < len(dirs)-1 {
			io.WriteString(w, "\n")
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

func printTarget(w io.Writer, useColor bool, info os.FileInfo) {
	name := info.Name() + "\n"
	if !useColor {
		io.WriteString(w, name)
		return
	}

	if info.IsDir() {
		blue.ColorPrint(w, name)
	} else if info.Mode().IsRegular() && info.Mode()&0111 != 0 {
		green.ColorPrint(w, name)
	} else {
		io.WriteString(w, name)
	}
}
