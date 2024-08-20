package main

import (
	"bufio"
	"os"
)

// Platform holds variations of a multisite in JSON form
type Platform struct {
	Test        Location `json:"test"`
	Blog        Location `json:"blog"`
	Forms       Location `json:"forms"`
	Engage      Location `json:"engage"`
	Vanity      Location `json:"vanity"`
	Staging     Location `json:"staging"`
	Working     Location `json:"working"`
	Production  Location `json:"production"`
	Development Location `json:"development"`
}

// Website holds the WordPress location data in JSON form
type Location struct {
	URL    string `json:"url"`
	Path   string `json:"path"`
	Server string `json:"server"`
}

// Person holds the Administrator email in JSON form
type Person struct {
	Admin string `json:"admin"`
}

// Constant declarations
const (
	automatic string = "\033[0m"
	fgYellow  string = "\033[33m"
	bgRed     string = "\033[41m"
	halt      string = "program halted "
	huh       string = "Unrecognized flag detected -"
	few       string = "Insufficient arguments supplied -"
	many      string = "Too many arguments supplied -"
)

// Variable declarations
var (
	wordpress    Platform
	user         Person
	sspot, dspot int
	purpose      = os.Args
	sflag        = os.Args[1]
	dflag        = os.Args[2]
	siteSlug     = os.Args[3]
	reader       = bufio.NewReader(os.Stdin)
	choices      = []string{"-p", "-s", "-b", "-d", "-t", "-e", "-f", "-w", "-v"}
	// String variables used to create objects
	sourcePath, sourceURL, sourceID, destPath, destURL, destID string
)
