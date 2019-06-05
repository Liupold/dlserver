package main

import (
	"fmt"

	"github.com/liupold/dlserver/pkg/demeter"
)

func main() {
	url := "https://mirror.downloadvn.com/videolan/vlc/3.0.6/win32/vlc-3.0.6-win32.exe"
	demeterObj := demeter.MkDemeter(url)
	fmt.Println(demeterObj.URL)
}
