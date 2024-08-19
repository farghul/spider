package main

import (
	"os"
	"strings"
)

func exportDB() {
	changedir(sourcePath)
	inside := execute("-c", "wp", "db", "tables", "--all-tables-with-prefix", "--url="+sourceURL+"/"+siteSlug, "--format=csv")
	result := strings.ReplaceAll(string(inside), "\n", ",")
	result = strings.TrimSuffix(result, ",")
	execute("-v", "wp", "db", "export", "--tables="+result, "/data/temp/"+siteSlug+".sql")

}

// Create a user export file in JSON format
func exportUsers() {
	changedir(sourcePath)
	people := execute("-c", "wp", "user", "list", "--url="+sourceURL+"/"+siteSlug, "--format=csv")
	inspect(os.WriteFile("/data/temp/"+siteSlug+".json", people, 0666))
}

// Create a new WordPress site
func createSite(email string) {
	changedir(destPath)
	execute("-v", "wp", "site", "create", "--url=https://"+destURL, "--title="+siteSlug, "--slug="+siteSlug, "--email="+email)
}

// Backup the WordPress SQL database
func backupDB() {
	changedir(destPath)
	execute("-v", "wp", "db", "export", "--path=/data/www-app/"+destPath+"/current/web/wp", "/data/temp/backup.sql")
}

// Replace the destination (destID) blog_id's with that of the source (sourceID)
func siteID() {
	execute("-v", "sed", "-i", "s/wp_"+sourceID+"_/wp_"+destID+"_/g", "/data/temp/"+siteSlug+".sql")
}

// Import the WordPress SQL database tables
func importDB() {
	changedir(destPath)
	execute("-v", "wp", "db", "import", "/data/temp/"+siteSlug+".sql")
}

// Correct the site links with wp search-replace
func linkFix() {
	changedir(destPath)
	execute("-v", "wp", "search-replace", "--url="+destURL+"/"+siteSlug, "--all-tables-with-prefix", sourceURL, destURL)
}

// Copy the WordPress site assets over
func assetCopy() {
	changedir(destPath)
	execute("-v", "rsync", "-a", "/data/www-assets/"+sourcePath+"/uploads/sites/"+sourceID+"/", "/data/www-assets/"+destPath+"/uploads/sites/"+destID+"/")
	// rsync -a /data/www-app/test_blog_gov_bc_ca/current/web/app/uploads/sites/392/ /data/www-app/dev_blog_gov_bc_ca/current/web/app/uploads/sites/187/ --stats

}

// Correct the references to the uploads folder
func uploadsFolder() {
	changedir(destPath)
	execute("-v", "wp", "search-replace", "--url="+destURL+"/"+siteSlug, "--all-tables-with-prefix", "app/uploads/sites/"+sourceID, "app/uploads/sites/"+destID)
}

// Correct any unescaped folders due to Gutenberg Blocks
func uploadsFolderEscapes() {
	changedir(destPath)
	execute("-v", "wp", "search-replace", "--url="+destURL+"/"+siteSlug, "--all-tables-with-prefix", "app\\/uploads\\/sites\\/"+sourceID, "app\\/uploads\\/sites\\/"+destID)
}

// Catch any lingering http addresses and convert them to https
func httpFind() {
	changedir(destPath)
	execute("-v", "wp", "search-replace", "--url="+destURL+"/"+siteSlug, "--all-tables-with-prefix", "http://", "https://")
}

// Remap the users to match their new ID
func remap() {
	changedir(destPath)
	execute("-v", "wp", "user", "import-csv", "/data/temp/"+siteSlug+".json")
}

// Flush the WordPress cache
func flush() {
	changedir(destPath)
	execute("-v", "wp", "cache", "flush")
}
