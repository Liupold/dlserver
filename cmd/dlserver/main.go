package main

import (
	"fmt"

	"github.com/liupold/dlserver/pkg/demeter"
)

func main() {
	url := "https://mirror.downloadvn.com/videolan/vlc/3.0.6/win32/vlc-3.0.6-win32.exe"
	demeterObj := demeter.MkDemeter(url, 8, "Downloads/", "~/tmp")
	fmt.Printf("%+v\n", demeterObj)
	demeter.Download(demeterObj)
	print(&demeterObj)
}
