package manager

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/cheggaaa/pb/v3"
	"github.com/liupold/dlserver/pkg/ghelp"
	"github.com/vbauerster/mpb"
	"github.com/vbauerster/mpb/decor"

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
		doneLN := demeter.GetTotalDone(&demeterObj)
		done := ghelp.ByteCountIEC(doneLN)
		total := ghelp.ByteCountIEC(demeterObj.Length)
		activeindicator := ""
		if demeterObj.Active {
			activeindicator = "*"
		}
		if demeterObj.Done {
			activeindicator = "Done"
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
	go func() {
		for {
			doneLN := demeter.GetTotalDone(demeterObj)
			bar.SetCurrent(int64(doneLN))
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

	demeterObj := &dlList[len(dlList)-1]
	StartDownload(len(dlList) - 1)
	p := mpb.New(
		mpb.WithWidth(60),
		mpb.WithRefreshRate(180*time.Millisecond))
	// name := demeterObj.Filename
	size := demeterObj.Length
	bar := p.AddBar(int64(size), mpb.BarStyle(" ██░ "),
		mpb.PrependDecorators(
			decor.Name("["),
			decor.Percentage(),
			decor.CountersKibiByte("]  % 6.1f / % 6.1f "),
		),
		mpb.AppendDecorators(
			// decor.EwmaETA(decor.ET_STYLE_GO, 60),
			decor.OnComplete(
				// ETA decorator with ewma age of 60, and width reservation of 4
				decor.EwmaETA(decor.ET_STYLE_GO, 60), "done",
			),
			decor.Name(" || "),
			decor.AverageSpeed(decor.UnitKiB, "% .2f"),
		),
	)
	for {
		start := time.Now()
		prevDone := demeter.GetTotalDone(demeterObj)
		time.Sleep(200 * time.Millisecond)
		doneLN := demeter.GetTotalDone(demeterObj)
		delta := doneLN - prevDone
		bar.IncrBy(delta, time.Since(start))

		if demeterObj.Done {
			bar.Abort(true)
			p.Wait()
			break
		}
	}
}
