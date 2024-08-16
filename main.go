package main

import "os"

// Constant declarations
const (
	few  string = "Insufficient arguments supplied -"
	many string = "Too many arguments supplied -"
)

// Start of the Spider application
func main() {
	if len(os.Args) < 4 {
		alert(few)
	} else if len(os.Args) > 4 {
		alert(many)
	} else {
		quarterback()
	}
}
