package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/cheggaaa/pb/v3"
	"github.com/liupold/dlserver/pkg/demeter"
	"github.com/liupold/dlserver/pkg/ghelp"
)

func main() {
	for {
		var url string
		fmt.Print("URL ==>")
		fmt.Scan(&url)
		demeterObj := demeter.MkDemeter(url, 8, "/home/rohnch/Downloads/", "/home/rohnch/cars")
		// fmt.Printf("%+v\n", demeterObj)
		var MainSyncG sync.WaitGroup
		MainSyncG.Add(1)
		if demeterObj.Resumeable == false {
			fmt.Println("Download Boost Not Available for " + demeterObj.Filename)
		}
		go demeter.Download(&demeterObj, &MainSyncG)
		fmt.Println(ghelp.ByteCountIEC(demeterObj.Length))

		/* var doneLN int64
		for {
			for _, doneln := range demeterObj.DonelnList {
				doneLN += doneln
			}
			fmt.Printf("Done: %.2f%s \r", (float64(doneLN)/float64(demeterObj.Length))*100, "%")
			time.Sleep(time.Millisecond * 10)
			if demeterObj.Done {
				break
			}
			doneLN = 0
		}
		fmt.Printf("\n") */
		bar := pb.StartNew(demeterObj.Length)
		bar.Set(pb.Bytes, true)
		var doneLN int64
		for {
			for _, doneln := range demeterObj.DonelnList {
				doneLN += doneln
			}
			bar.SetCurrent(doneLN)
			time.Sleep(time.Millisecond)
			doneLN = 0
			if demeterObj.Done {
				bar.Finish()
				break
			}
		}
	}
}
