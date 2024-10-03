package main

/* ----- Dry Run Functions ----- */

// Correct the links with search-replace using --dry-run
func linkFixDR() {
	changedir(destPath)
	execute("-v", "wp", "search-replace", "--url="+destURL+"/"+siteSlug, "--all-tables-with-prefix", sourceURL, destURL, "--dry-run")
	execute("-v", "wp", "search-replace", "--ssh="+user.SSH+destServer+":/data/www-app/"+destPath, "--url="+destURL+"/"+siteSlug, "--all-tables-with-prefix", sourceURL, destURL, "--dry-run")
}

// Copy the site assets over using --dry-run
func assetCopyDR() {
	execute("-v", "rsync", "-a", "/data/www-assets/"+sourcePath+"/uploads/sites/"+sourceID+"/", "/data/www-assets/"+destPath+"/uploads/sites/"+destID+"/", "--stats", "--dry-run")
	execute("-v", "rsync", "-a", "/data/www-assets/"+sourcePath+"/uploads/sites/"+sourceID+"/", "--ssh="+user.SSH+destServer+":/data/www-assets/"+destPath+"/uploads/sites/"+destID+"/", "--stats", "--dry-run")
}

// Correct the references to the uploads folder using --dry-run
func uploadsFolderDR() {
	changedir(destPath)
	execute("-v", "wp", "search-replace", "--url="+destURL+"/"+siteSlug, "--all-tables-with-prefix", "app/uploads/sites/"+sourceID, "app/uploads/sites/"+destID, "--dry-run")
	execute("-v", "wp", "search-replace", "--ssh="+user.SSH+destServer+":/data/www-app/"+destPath, "--url="+destURL+"/"+siteSlug, "--all-tables-with-prefix", "app/uploads/sites/"+sourceID, "app/uploads/sites/"+destID, "--dry-run")
}

// Correct any unescaped folders due to Gutenberg Blocks using --dry-run
func uploadsFolderEscapesDR() {
	changedir(destPath)
	execute("-v", "wp", "search-replace", "--url="+destURL+"/"+siteSlug, "--all-tables-with-prefix", "app\\/uploads\\/sites\\/"+sourceID, "app\\/uploads\\/sites\\/"+destID, "--dry-run")
}

// Catch any lingering http addresses using --dry-run
func httpFindDR() {
	changedir(destPath)
	execute("-v", "wp", "search-replace", "--url="+destURL+"/"+siteSlug, "--all-tables-with-prefix", "http://", "https://", "--dry-run")
}
