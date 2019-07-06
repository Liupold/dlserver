package main

import (
	"fmt"
	"sync"

	"github.com/liupold/dlserver/pkg/demeter"
	"github.com/liupold/dlserver/pkg/ghelp"
)

func main() {
	for {
		var url string
		fmt.Print("URL -->")
		fmt.Scan(&url)
		demeterObj := demeter.MkDemeter(url, 8, "/home/rohnch/Downloads/", "/home/rohnch/cars")
		// fmt.Printf("%+v\n", demeterObj)
		var MainSyncG sync.WaitGroup
		MainSyncG.Add(1)
		demeter.Download(demeterObj, &MainSyncG)
		fmt.Println(ghelp.ByteCountIEC(demeterObj.Length))
		MainSyncG.Wait()
		// print(&demeterObj, "\n")
	}
}
