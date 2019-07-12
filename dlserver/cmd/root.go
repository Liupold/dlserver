package cmd

var version = "0.1.0"
var location = "/home/rohnch/Dl"
var tmplocation = "/home/rohnch/cars"
var thcount = 8

// Start : start the cli
func Start() {
	greet()
	for {
		InputHandler()
	}
}
