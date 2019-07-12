package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/liupold/dlserver/pkg/ghelp"

	"github.com/liupold/dlserver/pkg/manager"
)

func greet() {
	fmt.Println(strings.Repeat("─", 60))
	fmt.Println("Dlserver Based CLI 			       [ALPHA 1.0]")
	fmt.Println("Build By: liupold (rohn.ch@gmail.com)")
	fmt.Println("Location: ", location)
	fmt.Println("Temp Location: ", tmplocation)
	fmt.Print(strings.Repeat("─", 60), "\n\n")

}

// InputHandler :
func InputHandler() {
	var inputData string
	fmt.Print("))> ")
	fmt.Scan(&inputData)
	if inputData == "art" {
		fmt.Printf(
			`
  _____  _                                __   ___  
 |  __ \| |                              /_ | / _ \ 
 | |  | | |___  ___ _ ____   _____ _ __   | || | | |
 | |  | | / __|/ _ \ '__\ \ / / _ \ '__|  | || | | |
 | |__| | \__ \  __/ |   \ V /  __/ |     | || |_| |
 |_____/|_|___/\___|_|    \_/ \___|_|     |_(_)___/ 
                                                                                          
					   
`)
	} else if inputData == "exit" {
		os.Exit(0)
	} else if inputData == "clear" {
		fmt.Print("\033[H\033[2J")
	} else if inputData == "list" {
		manager.ListAllDownload()

	} else if inputData == "home" {
		fmt.Print("\033[H\033[2J")
		greet()

	} else if inputData == "status" {
		manager.ShowProgressAll()

	} else if len(inputData) > 5 {
		if inputData[0:5] == "set::" {
			varr := inputData[5:]
			if varr == "location=" {
				location = inputData[14:]
			}
			if varr == "tmplocation=" {
				tmplocation = inputData[17:]
			}
		} else if inputData[0:5] == "get::" {
			varr := inputData[5:]
			if varr == "location" {
				fmt.Println(location)
			}
			if varr == "tmplocation" {
				fmt.Println(tmplocation)
			}
		} else if inputData[0:5] == "add::" {
			url := inputData[5:]
			manager.AddDownload(url, thcount, location, tmplocation)

		} else if inputData[0:6] == "stop::" {
			index, err := strconv.Atoi(inputData[6:])
			ghelp.ErrCheck(err)
			manager.StopDownload(index)

		} else if inputData[0:7] == "start::" {
			index, err := strconv.Atoi(inputData[7:])
			ghelp.ErrCheck(err)
			manager.StartDownload(index)

		} else {
			manager.SingleAddandWait(inputData, thcount, location, tmplocation)

		}
	} else {
		manager.SingleAddandWait(inputData, thcount, location, tmplocation)
	}
}
