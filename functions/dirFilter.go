package functions

import (
	"os"
)

func dirFilter(entries []os.DirEntry) []os.DirEntry {
	filtered := []os.DirEntry{}

	for _, tgt := range entries {
		if isHidden(tgt) {
			continue
		} else {
			filtered = append(filtered, tgt)
		}
	}
	return filtered
}

func isHidden(tgt os.DirEntry) bool {
	return tgt.Name()[0] == '.'
}
