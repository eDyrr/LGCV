package main

import "flag"

func main() {
	var folder string
	var email string
	flag.StringVar(&folder, "add", "/Users/edyrr/Desktop/Projects/C++", "add a new folder to scan for Git repos")
	flag.StringVar(&email, "email", "e6yr@proton.me", "the email to scan")
	flag.Parse()

	if folder != "" {
		scan(folder)
		return
	}

	stats(email)
}
