package main

import (
	"os"
	"strings"
)

// Export the source WordPress site database to an sql file
func exportDB() {
	changedir(sourcePath)
	inside, err := execute("wp", []string{"db", "tables", "--all-tables-with-prefix", "--url=" + sourceURL + "/" + siteSlug, "--format=csv"}, ExecOptions{Stream: false})
	inspect(err)
	result := strings.ReplaceAll(string(inside), "\n", ",")
	result = strings.TrimSuffix(result, ",")
	execute("wp", []string{"db", "export", "--tables=" + result, "/data/temp/" + siteSlug + ".sql"}, ExecOptions{Stream: true})
}

// Create a user export file in JSON format
func exportUsers() {
	changedir(sourcePath)
	people, err := execute("wp", []string{"user", "list", "--url=" + sourceURL + "/" + siteSlug, "--format=csv"}, ExecOptions{Stream: false})
	inspect(err)
	inspect(os.WriteFile("/data/temp/"+siteSlug+".json", people, 0666))
}

// Create a new WordPress site
func createSite(email string) {
	changedir(destPath)
	execute("wp", []string{"site", "create", "--url=https://" + destURL, "--title=" + siteSlug, "--slug=" + siteSlug, "--email=" + email}, ExecOptions{Stream: true})
}

// Backup the entire destination WordPress SQL database
func backupDB() {
	changedir(destPath)
	execute("wp", []string{"db", "export", "--path=/data/www-app/" + destPath + "/current/web/wp", "/data/temp/backup.sql"}, ExecOptions{Stream: true})
}

// Replace the source (sourceID) with that of the destination (destID)
func siteID() {
	execute("sed", []string{"-i", "s/wp_" + sourceID + "_/wp_" + destID + "_/g", "/data/temp/" + siteSlug + ".sql"}, ExecOptions{Stream: true})
}

// Import the WordPress SQL database tables into the destination
func importDB() {
	changedir(destPath)
	execute("wp", []string{"db", "import", "/data/temp/" + siteSlug + ".sql"}, ExecOptions{Stream: true})
}

// Correct the site links with wp search-replace
func linkFix() {
	changedir(destPath)
	execute("wp", []string{"search-replace", "--url=" + destURL + "/" + siteSlug, "--all-tables-with-prefix", sourceURL, destURL}, ExecOptions{Stream: true})
}

// Copy the WordPress site assets over
func assetCopy() {
	changedir(destPath)
	execute("wp", []string{"rsync", "-a", "/data/www-assets/" + sourcePath + "/uploads/sites/" + sourceID + "/", "/data/www-assets/" + destPath + "/uploads/sites/" + destID + "/"}, ExecOptions{Stream: true})
}

// Correct the references to the uploads folder
func uploadsFolder() {
	changedir(destPath)
	execute("wp", []string{"search-replace", "--url=" + destURL + "/" + siteSlug, "--all-tables-with-prefix", "app/uploads/sites/" + sourceID, "app/uploads/sites/" + destID}, ExecOptions{Stream: true})
}

// Correct any unescaped folders due to Gutenberg Blocks
func uploadsFolderEscapes() {
	changedir(destPath)
	execute("wp", []string{"search-replace", "--url=" + destURL + "/" + siteSlug, "--all-tables-with-prefix", "app\\/uploads\\/sites\\/" + sourceID, "app\\/uploads\\/sites\\/" + destID}, ExecOptions{Stream: true})
}

// Catch any lingering http addresses and convert them to https
func httpFind() {
	changedir(destPath)
	execute("wp", []string{"search-replace", "--url=" + destURL + "/" + siteSlug, "--all-tables-with-prefix", "http://", "https://"}, ExecOptions{Stream: true})
}

// Remap the users to match their new ID
func remap() {
	changedir(destPath)
	execute("wp", []string{"user", "import-csv", "/data/temp/" + siteSlug + ".json"}, ExecOptions{Stream: true})
}

// Flush the WordPress cache
func flush() {
	changedir(destPath)
	execute("wp", []string{"cache", "flush"}, ExecOptions{Stream: true})
}
