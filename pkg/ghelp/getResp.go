package ghelp

import "net/http"

// GetResp : return http GET fro a url and handles Error
func GetResp(url string) *http.Response {
	/* return resp struct for a given url */
	resp, err := http.Get(url)
	ErrCheck(err)
	return resp
}
