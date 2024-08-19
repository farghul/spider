package main

import (
	"encoding/json"
	"strings"
)

// Quarterback function controls the flow of the program
func quarterback() {
	// read a json file and builds the objects
	sites := readit("local/env.json")
	json.Unmarshal(sites, &websites)
	var sspot, dspot int

	combinations := [9][2]string{
		{websites.Production.URL, websites.Production.Path},
		{websites.Staging.URL, websites.Staging.Path},
		{websites.Blog.URL, websites.Blog.Path},
		{websites.Development.URL, websites.Development.Path},
		{websites.Test.URL, websites.Test.Path},
		{websites.Engage.URL, websites.Engage.Path},
		{websites.Forms.URL, websites.Forms.Path},
		{websites.Working.URL, websites.Working.Path},
		{websites.Vanity.URL, websites.Vanity.Path},
	}

	for i, f := range choices {
		if f == sflag {
			sspot = i
		}

		if f == dflag {
			dspot = i
		}
	}

	source(combinations[sspot][0], combinations[sspot][1])
	first()
	destination(combinations[dspot][0], combinations[dspot][1])
	receiver()

	// switch sflag {
	// case "-s":
	// 	source(websites.Staging.URL, websites.Staging.Path)
	// case "-p":
	// 	source(websites.Production.URL, websites.Production.Path)
	// case "-d":
	// 	source(websites.Development.URL, websites.Development.Path)
	// case "-e":
	// 	source(websites.Engage.URL, websites.Engage.Path)
	// case "-f":
	// 	source(websites.Forms.URL, websites.Forms.Path)
	// case "-w":
	// 	source(websites.Working.URL, websites.Working.Path)
	// case "-v":
	// 	source(websites.Vanity.URL, websites.Vanity.Path)
	// default:
	// 	alert(huh)
	// }

	// switch dflag {
	// case "-s":
	// 	destination(websites.Staging.URL, websites.Staging.Path)
	// case "-p":
	// 	destination(websites.Production.URL, websites.Production.Path)
	// case "-d":
	// 	destination(websites.Development.URL, websites.Development.Path)
	// case "-e":
	// 	destination(websites.Engage.URL, websites.Engage.Path)
	// case "-f":
	// 	destination(websites.Forms.URL, websites.Forms.Path)
	// case "-w":
	// 	destination(websites.Working.URL, websites.Working.Path)
	// case "-v":
	// 	destination(websites.Vanity.URL, websites.Vanity.Path)
	// default:
	// 	alert(huh)
	// }

}

// Create the source object
func source(url, path string) {
	sourceURL, sourcePath = url, path          // Transfer local JSON contents to main code
	sourceList := construct(sourcePath)        // List of source sites in JSON format
	sourceID = aquireID(sourceURL, sourceList) // Creates a specific source object
}

// Create the destination object
func destination(url, path string) {
	destURL, destPath = url, path        // Transfer local JSON contents to main code
	destList := construct(destPath)      // List of destination sites in JSON format
	destID = aquireID(destURL, destList) // The specific destination object
}

// Run the first few functions up to the new site creation
func first() {
	banner("Exporting the database tables")
	exportDB()
	banner("Creating a user export file")
	exportUsers()
	banner("Creating the new WordPress site")
	// createSite(websites.Email.Admin)
}

// Query WordPress for a list of all sites and save as a csv variable
func construct(path string) string {
	url := properQURL(path)
	query := execute("-c", "wp", "site", "list", "--fields=blog_id,url", "--path=/data/www-app/"+path+"/current/web/wp", "--url="+url, "--skip-plugins", "--skip-themes", "--format=csv")
	result := strings.Replace(string(query), "blog_id,url\n", "", 1)
	result = strings.ReplaceAll(result, "\n", ",")
	result = strings.TrimSuffix(result, ",")

	return result
}

// Search the blog list to find the ID that matches the supplied URL
func aquireID(url, list string) string {
	var blid string
	blogs := strings.Split(list, ",")

	for order, item := range blogs {
		if strings.Contains(item, url+"/"+siteSlug) {
			blid = blogs[order-1]
		}
	}

	return blid
}

// Trigger the rest of the program after passing through the Quarterback
func receiver() {
	second()
	dryrun()
	last()
}

// Run the second round of functions after being able to grab the new site ID
func second() {
	banner("Backing up the database")
	// backupDB()
	banner("Replacing the destination blog_id with that of the source")
	siteID()
	banner("Importing the database tables")
	importDB()
}

// Pre-emptively run the data modifying functions in --dry-run mode
func dryrun() {
	banner("Updating URL's")
	linkFixDR()
	direct(confirm(), "lf")
	banner("Copying Assets")
	assetCopyDR()
	direct(confirm(), "ac")
	banner("Fixing Uploads")
	uploadsFolderDR()
	direct(confirm(), "fr")
	banner("Fixing Escapes")
	uploadsFolderEscapesDR()
	direct(confirm(), "fr2")
	banner("Fixing HTTP References")
	httpFindDR()
	direct(confirm(), "hf")
}

// Run the remaining functions
func last() {
	banner("Remaping the users to match their new ID")
	remap()
	banner("Flushing the WordPress cache")
	flush()
}
