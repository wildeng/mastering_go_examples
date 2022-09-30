package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"regexp"
)

func main() {
	var filename string

	flag.StringVar(&filename, "f", "", "Specify the path and name of the file to read")
	flag.Parse()

	if filename != "" {
		mails := FindEmail(filename)
		fmt.Println(mails)
	} else {
		fmt.Println("Filename missing! Exiting")
	}
}

func FindEmail(filename string) []string {
	var mailRegexp = regexp.MustCompile(`[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}`)
	var results []string
	b, _ := ioutil.ReadFile(filename)
	results = mailRegexp.FindAllString(string(b), -1)
	return results
}
