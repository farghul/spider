package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Tell the program what to do based on the results of a --dry-run
func direct(answer, nav string) {
	if strings.ToLower(answer) == "y" {
		proceed(nav)
	} else {
		os.Exit(0)
	}
}

// Solicit user confirmation after completion of a --dry-run
func confirm() string {
	answer := solicit("Does this output seem acceptable? Shall we continue without the --dry-run flag? (y/n) ")
	return answer
}

// Execute the functions without a --dry-run condition
func proceed(action string) {
	switch action {
	case "lf":
		linkFix()
	case "ac":
		assetCopy()
	case "fr":
		uploadsFolder()
	case "fr2":
		uploadsFolderEscapes()
	case "hf":
		httpFind()
	}
}

// Run through server options
func discovery(trios [9][3]string) {
	for i, f := range servers {
		if f == sflag {
			sourceURL = trios[i][0]
			sourcePath = trios[i][1]
			sourceServer = trios[i][2]
		}

		if f == dflag {
			destURL = trios[i][0]
			destPath = trios[i][1]
			destServer = trios[i][2]
		}
	}
}

// Set the proper url for running WP CLI queries
func properQURL(path string) string {
	var url string

	if strings.Contains(path, "test") {
		url = wordpress.Test.URL
	} else if strings.Contains(path, "dev") {
		url = wordpress.Development.URL
	} else {
		url = wordpress.Blog.URL
	}

	return url
}

func execute(task string, args []string, opts ExecOptions) ([]byte, error) {
	cmd := exec.Command(task, args...)
	cmd.Env = append(os.Environ(), opts.Env...)
	cmd.Dir = opts.Dir

	if opts.Stream {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return nil, cmd.Run()
	}

	return cmd.CombinedOutput()
}

// Get user input via screen prompt
func solicit(prompt string) string {
	fmt.Print(prompt)
	response, _ := reader.ReadString('\n')
	return strings.TrimSpace(response)
}

// Navigate to the WordPress installation
func changedir(path string) {
	os.Chdir("/data/www-app/" + path)
}

// Read any file and return the contents as a byte variable
func read(file string) []byte {
	outcome, problem := os.ReadFile(file)
	inspect(problem)
	return outcome
}

// Check for errors, print the result if found
func inspect(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}

// Provide and highlight an informational message
func banner(message string) {
	fmt.Println(fgYellow)
	fmt.Println("**", automatic, message, fgYellow, "**", automatic)
}

// Alert prints a colourized error message
func alert(message string) {
	fmt.Println(bgRed, message, halt)
	fmt.Println(automatic)
	os.Exit(0)
}
