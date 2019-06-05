package ghelp

import "fmt"

func errCheck(err error) {
	/* check if err message is nil
	if err != nil panic */
	if err != nil {
		panic(err)
	}
}

func isError(err error) bool {
	/* chek if err message is nil
	panic or return true */
	if err != nil {
		panic(err)
	}
	return true
}

func printErr(err error) {
	/* only print err if not nil */
	if err != nil {
		fmt.Println(err)
	}
}
