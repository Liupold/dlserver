package demeter

import (
	"strconv"

	"github.com/liupold/dlserver/pkg/ghelp"
)

/*
get information for demeter
add GetFilename func
*/

// GetResumeable : return bool, if support range header
func GetResumeable(demeterObj Demeter) bool {
	resp := demeterObj.TmpResp
	if resp.Header["Accept-Ranges"] == nil {
		return false
	} else if resp.Header["Accept-Ranges"][0] == "bytes" {
		return true
	}
	return false
}

// GetLength : get the content length of a request (file todonwload)
func GetLength(demeterObj Demeter) int {
	resp := demeterObj.TmpResp
	if resp.Header["Content-Length"] == nil {
		return -1
	}
	length, err := strconv.Atoi(resp.Header["Content-Length"][0])
	ghelp.ErrCheck(err)
	return length
}

// GetFilename : return the filename of (file to download)
func GetFilename(demeterObj Demeter) string {
	return ""
}
