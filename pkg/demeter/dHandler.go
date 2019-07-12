package demeter

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"

	"github.com/liupold/dlserver/pkg/ghelp"
)

// WriteToFile :
func WriteToFile(resp *http.Response, filepath string, doneln *int64, paused *bool) {
	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	ghelp.ErrCheck(err)
	//open files

	for {
		nWritten, err := io.CopyN(file, resp.Body, 4096)
		*doneln += nWritten
		if err == io.EOF {
			break
		}
		if *paused {
			break
		}
		ghelp.ErrCheck(err)
	}
}

func getDone(demeterObj *Demeter, index int) int {
	tmpPath := getPartName(demeterObj, index)
	f, err := os.OpenFile(tmpPath, os.O_RDONLY, 0666)
	if os.IsNotExist(err) {
		return 0
	}
	fi, err := f.Stat()
	ghelp.ErrCheck(err)
	return int(fi.Size())

}

func getDlRanges(demeterObj *Demeter) [][2]int {
	ThCount := demeterObj.ThCount
	Length := demeterObj.Length
	rangeList := make([][2]int, ThCount)
	lastIndex := ThCount - 1
	var dlRange [2]int
	partLength := Length / ThCount
	for i := 0; i < lastIndex; i++ {
		alreadyDone := getDone(demeterObj, i)
		demeterObj.DonelnList[i] = int64(alreadyDone)
		dlRange[0] = (i * partLength) + alreadyDone
		dlRange[1] = (i+1)*partLength - 1
		rangeList[i] = dlRange

	}
	alreadyDone := getDone(demeterObj, lastIndex)
	dlRange[0] = (lastIndex * partLength) + alreadyDone
	dlRange[1] = Length
	demeterObj.DonelnList[lastIndex] = int64(alreadyDone)
	rangeList[lastIndex] = dlRange
	return rangeList

}

func getPartName(demeterObj *Demeter, index int) string {
	return fmt.Sprintf("%s/%s.%d.part", demeterObj.TmpLocation, demeterObj.Filename, index)
}

func fileMerge(demterObj *Demeter) {
	finalPath := demterObj.Location + demterObj.Filename
	finalFile, err := os.OpenFile(finalPath, os.O_RDWR|os.O_CREATE, 0666)
	ghelp.ErrCheck(err)
	for i := 0; i < demterObj.ThCount; i++ {
		tmpPath := getPartName(demterObj, i)
		tmpFile, err := os.OpenFile(tmpPath, os.O_RDWR|os.O_CREATE, 0666)
		ghelp.ErrCheck(err)
		io.Copy(finalFile, tmpFile)
		tmpFile.Close()

	}
	// check if the file is ok
	fi, err := finalFile.Stat()
	ghelp.ErrCheck(err)
	if int(fi.Size()) == demterObj.Length {
		for i := 0; i < demterObj.ThCount; i++ {
			tmpPath := getPartName(demterObj, i)
			os.Remove(tmpPath)
		}
	} else if demterObj.Length > 0 {
		err := errors.New("File Damaged ")
		ghelp.ErrCheck(err)
	}
	finalFile.Close()
}

// Download : dl the file
func Download(demeterObj *Demeter, MainSyncG *sync.WaitGroup) {
	defer MainSyncG.Done()
	client := &http.Client{}
	var wg sync.WaitGroup
	rangeList := getDlRanges(demeterObj)
	// demeterObj.DonelnList = make([]int64, demeterObj.ThCount)
	wg.Add(demeterObj.ThCount)
	for index, dlRange := range rangeList {
		if dlRange[0] < dlRange[1] {
			go func(index int, dlRange [2]int) {
				defer wg.Done()
				req, err := http.NewRequest("GET", demeterObj.URL, nil)
				ghelp.ErrCheck(err)

				stringRange := fmt.Sprintf("bytes=%d-%d", dlRange[0], dlRange[1])
				req.Header.Add("Range", stringRange)

				resp, err := client.Do(req)
				ghelp.ErrCheck(err)

				demeterObj.RespList = append(demeterObj.RespList, resp)

				filepath := getPartName(demeterObj, index)
				WriteToFile(resp, filepath, &demeterObj.DonelnList[index], &demeterObj.Paused)
			}(index, dlRange)
		} else if dlRange[0] == dlRange[1]+1 {
			wg.Done()
		} else if index == demeterObj.ThCount-1 {
			// special rule for the last part
			if dlRange[0] == dlRange[1] {
				wg.Done()
			}
		} else {
			fmt.Println("tmp file(s) got f++ked")
		}
	}
	wg.Wait()
	fileMerge(demeterObj)
	demeterObj.Done = true
}

//Pause : pause the download
func Pause(demeterObj *Demeter) {
	if demeterObj.Resumeable != true {
		fmt.Println("Warning: File will Not Resume")
	}
	demeterObj.Paused = true
}
