package demeter

import (
	"net/http"
)

// Demeter {TmpResp, Url, Length, Resumeable, Filename, RsespList} : struct with download details
type Demeter struct {
	TmpResp    *http.Response
	URL        string
	Length     int
	Resumeable bool
	Filename   string
	RespList   []*http.Response
}
