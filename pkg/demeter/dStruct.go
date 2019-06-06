package demeter

import (
	"net/http"
)

// Demeter {TmpResp, Url, Length, Resumeable, Filename, RsespList} : struct with download details
type Demeter struct {
	Filename    string
	Length      int
	Running     bool
	RespList    []*http.Response
	Resumeable  bool
	ThCount     int // change this to int8 if needed
	TmpResp     *http.Response
	URL         string
	Location    string
	TmpLocation string
}
