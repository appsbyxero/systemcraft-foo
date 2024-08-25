package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func MatchInFile(filePath string, pattern *regexp.Regexp) bool {
	var file *os.File
	var err error

	if filePath == "" {
		file = os.Stdin
	} else {
		file, err = os.Open(filePath)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	}

	scanner := bufio.NewScanner(file)
	matched := false
	for scanner.Scan() {
		b := scanner.Bytes()
		fmt.Println(string(b))
	}

	return matched
}

// This passes the first test!
func main() {
	exitCode := 0
	pattern := os.Args[1]
	matchPattern, err := regexp.Compile(pattern)
	if err != nil {
		log.Fatal(err)
	}

	path := os.Args[2]
	if !MatchInFile(path, matchPattern) {
		exitCode = 1
	}

	os.Exit(exitCode)
}
