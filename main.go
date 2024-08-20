package main

// Start of the Spider application
func main() {
	if len(purpose) < 4 {
		alert(few)
	} else if len(purpose) > 4 {
		alert(many)
	} else {
		quarterback() // calls this function in the launch.go file
	}
}
