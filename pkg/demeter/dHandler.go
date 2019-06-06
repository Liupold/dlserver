package demeter

import "fmt"

// Download : dl the file
func Download(demeterObj Demeter) {

	var dlRange [2]int
	partLength := demeterObj.Length / demeterObj.ThCount
	for i := 0; i < 7; i++ {
		dlRange[0] = i * partLength
		dlRange[1] = (i+1)*partLength - 1
		fmt.Println(dlRange)
	}
	dlRange[0] = 7 * partLength
	dlRange[1] = demeterObj.Length
	fmt.Println(dlRange)

}
