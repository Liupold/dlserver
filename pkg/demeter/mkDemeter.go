package demeter

import (
	"github.com/liupold/dlserver/pkg/ghelp"
)

// MkDemeter : make Demeter for download
func MkDemeter(url string) Demeter {
	demeterObj := Demeter{URL: url}
	tmpResp := ghelp.GetResp(url)
	tmpResp.Body.Close()
	demeterObj.TmpResp = tmpResp
	length := GetLength(demeterObj)
	resumeable := GetResumeable(demeterObj)
	filename := GetFilename(demeterObj)
	demeterObj.Filename = filename
	demeterObj.Length = length
	demeterObj.Resumeable = resumeable
	return demeterObj

}
