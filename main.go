package main

// Constant declarations
const (
	few  string = "Insufficient arguments supplied -"
	many string = "Too many arguments supplied -"
)

// Start of the Spider application
func main() {
	// whereisit()
	if len(purpose) < 4 {
		alert(few)
	} else if len(purpose) > 4 {
		alert(many)
	} else {
		quarterback() //calls this in the launch.go file
	}
}
