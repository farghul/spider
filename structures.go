package main

import (
	"bufio"
	"os"
)

// Platform holds variations of a multisite WordPress install
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

// Location holds the WordPress location data
type Location struct {
	URL    string `json:"url"`
	Path   string `json:"path"`
	Server string `json:"server"`
}

// Person holds the Administrator email and SSH credentials
type Person struct {
	Admin string `json:"admin"`
	SSH   string `json:"ssh"`
}

// Constant declarations
const (
	automatic string = "\033[0m"
	fgYellow  string = "\033[33m"
	bgRed     string = "\033[41m"
	halt      string = "program halted "
	huh       string = "Unrecognized flag detected -"
	many      string = "Too many arguments supplied -"
	few       string = "Insufficient arguments supplied -"
)

// Variable declarations
var (
	wordpress Platform
	user      Person
	purpose   = os.Args
	sflag     = os.Args[1]
	dflag     = os.Args[2]
	siteSlug  = os.Args[3]
	reader    = bufio.NewReader(os.Stdin)
	servers   = []string{"-p", "-s", "-b", "-d", "-t", "-e", "-f", "-w", "-v"}
	// String variables used to create objects
	sourcePath, sourceURL, sourceID, sourceServer, destPath, destURL, destID, destServer string
)
