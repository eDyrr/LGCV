package main

import (
	"flag"
	"fmt"
)

func main() {
	var folder string
	var email string
	flag.StringVar(&folder, "add", "", "add a new folder to scan for Git repos")
	flag.StringVar(&email, "email", "e6yr@proton.me", "the email to scan")
	flag.Parse()

	if folder != "" {
		scan(folder)
		return
	}
	stats(email)
	fmt.Println("after scan 3")
}
