package demeter

import (
	"github.com/liupold/dlserver/pkg/ghelp"
)

// MkDemeter : make Demeter for download
func MkDemeter(url string, thCount int, location string, tmpLocation string) Demeter {
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
	if resumeable {
		demeterObj.ThCount = thCount
	} else {
		demeterObj.ThCount = 1
	}
	demeterObj.Location = location
	demeterObj.TmpLocation = tmpLocation
	demeterObj.Active = false
	demeterObj.Done = false
	demeterObj.DonelnList = make([]int64, thCount)
	return demeterObj

}
