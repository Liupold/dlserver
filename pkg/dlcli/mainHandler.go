package dlcli

import "fmt"

func greet() {
	fmt.Println("Welcome to dlserver!, :)")
}

func getInfo(ver string, builder string, bdate string) {
	fmt.Print("Version: ", ver, "\n\n")
	fmt.Println("Builder: ", builder)
	fmt.Println("Build date:", bdate)
}
