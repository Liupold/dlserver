package ghelp

import (
	"os"
)

// RemoveIndexFile :  does what it says
func RemoveIndexFile(s []*os.File, index int) []*os.File {
	var newSclice []*os.File
	for indx, value := range s {
		if indx != index {
			newSclice = append(newSclice, value)
		}
	}
	return newSclice
}
