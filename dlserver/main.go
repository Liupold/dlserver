package main

import "github.com/liupold/dlserver/dlserver/cmd"

func main() {
	cmd.Start()
}

/*
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
*/
