package manager

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/cheggaaa/pb/v3"
	"github.com/liupold/dlserver/pkg/ghelp"

	"github.com/liupold/dlserver/pkg/demeter"
)

var dlList []demeter.Demeter
var mainSyncG sync.WaitGroup
var stopBar = false
var waitBarActive = false

// AddDownload :
func AddDownload(url string, thCount int, finalLocation string, tmpLocation string) {
	demeterObj := demeter.MkDemeter(url, thCount, finalLocation, tmpLocation)
	dlList = append(dlList, demeterObj)
}

// StartDownload :
func StartDownload(downloadListIndex int) {
	if !dlList[downloadListIndex].Done {
		mainSyncG.Add(1)
		go demeter.Download(&dlList[downloadListIndex], &mainSyncG)
		fmt.Println("Downloading: ", dlList[downloadListIndex].Filename)
	} else {
		fmt.Println("Already Done! ( ͡° ͜ʖ ͡°)")
	}
}

// ListAllDownload :
func ListAllDownload() {
	for index, demeterObj := range dlList {
		var doneLN int64
		for _, doneln := range demeterObj.DonelnList {
			doneLN += doneln
		}
		done := ghelp.ByteCountIEC(int(doneLN))
		total := ghelp.ByteCountIEC(demeterObj.Length)
		activeindicator := ""
		if demeterObj.Active {
			activeindicator = "*"
		}
		fmt.Printf("%d: [%s] %s (%s/%s)\n", index, activeindicator, demeterObj.Filename, done, total)
	}
}

// StopDownload :
func StopDownload(downloadListIndex int) {
	dlList[downloadListIndex].Active = false
}

// WaitBar : animation when you get board running a 1000+ downloads
func WaitBar(demeterObj *demeter.Demeter, syncGroup *sync.WaitGroup) {
	syncGroup.Add(1)
	bar := pb.StartNew(demeterObj.Length)
	bar.Set(pb.Bytes, true)
	var doneLN int64
	go func() {
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
			if stopBar == true {
				bar.Finish()
				break
			}
		}
	}()
	syncGroup.Done()
}

// ShowProgressAll : progressbar until ctrl+c
func ShowProgressAll() {
	waitBarActive = true
	var barSyncGroup sync.WaitGroup
	for _, demeterObj := range dlList {
		WaitBar(&demeterObj, &barSyncGroup)
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			stopBar = true
			waitBarActive = false
		}
	}()
}

// SingleAddandWait :
func SingleAddandWait(url string, thCount int, finalLocation string, tmpLocation string) {
	AddDownload(url, thCount, finalLocation, tmpLocation)

	demeterObj := dlList[len(dlList)-1]
	StartDownload(len(dlList) - 1)

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
