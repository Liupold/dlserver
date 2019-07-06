package demeter

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"

	"github.com/liupold/dlserver/pkg/ghelp"
)

// WriteToFile :
func WriteToFile(respList []*http.Response, filePathList []string, doneLen *int64, wg *sync.WaitGroup) {
	var fileList []*os.File
	//open files
	defer wg.Done()
	for _, filename := range filePathList {
		f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0666)
		ghelp.ErrCheck(err)
		fileList = append(fileList, f)
		ghelp.ErrCheck(err)
	}

	for {
		for indx, file := range fileList {
			nWritten, err := io.CopyN(file, respList[indx].Body, 4096)
			if err == io.EOF {
				file.Close()
				fileList = ghelp.RemoveIndexFile(fileList, indx)

			} else {
				ghelp.ErrCheck(err)
			}
			*doneLen = +nWritten
		}
		if len(fileList) == 0 {
			break
		}
	}
}

// Download : dl the file
func Download(demeterObj Demeter, MainSyncG *sync.WaitGroup) {

	rangeList := make([][2]int, demeterObj.ThCount)
	lastIndex := demeterObj.ThCount - 1
	var dlRange [2]int
	partLength := demeterObj.Length / demeterObj.ThCount
	for i := 0; i < lastIndex; i++ {
		dlRange[0] = i * partLength
		dlRange[1] = (i+1)*partLength - 1
		rangeList[i] = dlRange

	}
	dlRange[0] = lastIndex * partLength
	dlRange[1] = demeterObj.Length
	rangeList[lastIndex] = dlRange
	// create resp list

	fmt.Println("Initiating Download")
	client := &http.Client{}
	var wg sync.WaitGroup
	wg.Add(len(rangeList))
	for _, dlRange := range rangeList {
		go func(dlRange [2]int) {
			defer wg.Done()
			req, _ := http.NewRequest("GET", demeterObj.URL, nil)
			stringRange := fmt.Sprintf("bytes=%d-%d", dlRange[0], dlRange[1])
			req.Header.Add("Range", stringRange)
			resp, err := client.Do(req)
			ghelp.ErrCheck(err)
			demeterObj.RespList = append(demeterObj.RespList, resp)
		}(dlRange)
	}
	wg.Wait()
	fmt.Println("Initiating Done")

	var filePathList []string

	// genereate tmp file names
	for i := 0; i < demeterObj.ThCount; i++ {
		tmpName := fmt.Sprintf("%s/%s.%d.part", demeterObj.TmpLocation, demeterObj.Filename, i)
		filePathList = append(filePathList, tmpName)
	}
	go WriteToFile(demeterObj.RespList, filePathList, &demeterObj.DoneLength, MainSyncG)

}
