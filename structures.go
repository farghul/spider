package main

import "os"

// Platform holds the JSON data
type Platform struct {
	Test        Website `json:"test"`
	Blog        Website `json:"blog"`
	Forms       Website `json:"forms"`
	Engage      Website `json:"engage"`
	Vanity      Website `json:"vanity"`
	Staging     Website `json:"staging"`
	Working     Website `json:"working"`
	Production  Website `json:"production"`
	Development Website `json:"development"`
	Email       Person  `json:"email"`
}

// Website holds the JSON data
type Website struct {
	URL  string `json:"url"`
	Path string `json:"path"`
}

// Person holds the JSON data
type Person struct {
	Admin string `json:"admin"`
}

// Variable declarations
var (
	websites Platform
	purpose  = os.Args
	sflag    = os.Args[1]
	dflag    = os.Args[2]
	siteSlug = os.Args[3]
	// String variables used to create objects
	sourcePath, sourceURL, destPath, destURL, sourceID, destID string
)

// var combinations = [9][2]string{
// 	{websites.Production.URL, websites.Production.Path},
// 	{websites.Staging.URL, websites.Staging.Path},
// 	{websites.Blog.URL, websites.Blog.Path},
// 	{websites.Development.URL, websites.Development.Path},
// 	{websites.Test.URL, websites.Test.Path},
// 	{websites.Engage.URL, websites.Engage.Path},
// 	{websites.Forms.URL, websites.Forms.Path},
// 	{websites.Working.URL, websites.Working.Path},
// 	{websites.Vanity.URL, websites.Vanity.Path},
// }

var choices = []string{"-p", "-s", "-b", "-d", "-t", "-e", "-f", "-w", "-v"}
