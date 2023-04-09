package main

import (
	"flag"
	"gogitlocalcontribs/scan"
	"gogitlocalcontribs/stats"
)

// scan given a path crawls it and its subfolders
// searching for Git repositories


func main() {
	var folder string
	var email string

	flag.StringVar(&folder, "add", "", "add a new folder to scan for Git repositories")
	flag.StringVar(&email, "email", "youor@email.com", "the email to scan")
	flag.Parse()

	if folder != "" {
		scan.Scan(folder)
		return
	}

	stats.Stats(email)
}
