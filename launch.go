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
			sourceURL = combinations[i][0]
			sourcePath = combinations[i][1]
		}

		if f == dflag {
			destURL = combinations[i][0]
			destPath = combinations[i][1]
		}
	}

	sourceList := construct(sourcePath)
	sourceID = aquireID(sourceURL, sourceList)
	first()
	destList := construct(destPath)
	destID = aquireID(destURL, destList)
	receiver()
}

// Run the first few functions up to the new site creation
func first() {
	banner("Exporting the " + sourceURL + "/" + siteSlug + " database tables")
	exportDB()
	banner("Exporting the " + sourceURL + "/" + siteSlug + " users")
	exportUsers()
	banner("Creating the new " + destURL + "/" + siteSlug + "WordPress site")
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
	banner("Backing up the entire WordPress database")
	backupDB()
	banner("Replacing " + sourceID + " with " + destID)
	siteID()
	banner("Importing the " + sourceURL + "/" + siteSlug + " database tables")
	importDB()
}

// Pre-emptively run the data modifying functions in --dry-run mode
func dryrun() {
	banner("Updating URL's to " + destURL)
	linkFixDR()
	direct(confirm(), "lf")
	banner("Copying Assets to /data/www-assets/" + destPath + "/uploads/sites/" + destID)
	assetCopyDR()
	direct(confirm(), "ac")
	banner("Updating references to app/uploads/sites/" + destID)
	uploadsFolderDR()
	direct(confirm(), "fr")
	banner("Fixing unescaped folders due to Gutenberg Blocks in app/uploads/sites/" + destID)
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
