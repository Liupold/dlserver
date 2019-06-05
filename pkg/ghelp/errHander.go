package ghelp

import "fmt"

// ErrCheck : check if err message is nil
func ErrCheck(err error) {
	/* if err != nil panic */
	if err != nil {
		panic(err)
	}
}

// IsError : check if err message is nil
func IsError(err error) bool {
	/* chek if err message is nil
	panic or return bool*/
	if err != nil {
		return true
	}
	return false
}

//PrintErr : only print err if not nil
func PrintErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
