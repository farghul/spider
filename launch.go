package main

import (
	"encoding/json"
	"strings"
)

// Quarterback function controls the flow of the program
func quarterback() {
	sites := readit("local/env.json")
	json.Unmarshal(sites, &wordpress)

	combinations := [9][2]string{
		{wordpress.Production.URL, wordpress.Production.Path},
		{wordpress.Staging.URL, wordpress.Staging.Path},
		{wordpress.Blog.URL, wordpress.Blog.Path},
		{wordpress.Development.URL, wordpress.Development.Path},
		{wordpress.Test.URL, wordpress.Test.Path},
		{wordpress.Engage.URL, wordpress.Engage.Path},
		{wordpress.Forms.URL, wordpress.Forms.Path},
		{wordpress.Working.URL, wordpress.Working.Path},
		{wordpress.Vanity.URL, wordpress.Vanity.Path},
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
	createSite(user.Admin)
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
	backupDB()
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
