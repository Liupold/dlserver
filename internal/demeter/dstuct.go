package demeter

import "net/http"

// Demeter {url, length, resumeable, filename, respList} : struct with download details
type Demeter struct {
	url        string
	length     int
	resumeable bool
	filename   string
	respList   []*http.Response
}
