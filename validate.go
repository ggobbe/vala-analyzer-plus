package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

var (
	startsWithSpaces   = regexp.MustCompile(`^\s+`)
	fourSpacesIndented = regexp.MustCompile(`^( {4})*[^\s]+`)
	endsWithSpaces     = regexp.MustCompile(`\s+$`)
	openingParenthese  = regexp.MustCompile(`[_|\(| ]\(`)
)

func validateFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)

	num := 0
	for scanner.Scan() {
		line := scanner.Text()
		num++

		warnings := validateLine(line)

		if len(warnings) > 0 {
			for _, warning := range warnings {
				fmt.Printf("%s:%d\t%s\n", filename, num, warning)
			}
		}
	}
}

func validateLine(line string) []string {
	var warnings []string
	if strings.Trim(line, " ") == "{" {
		warnings = append(warnings, "First brace isn't on the end of the first line")
	}

	if endsWithSpaces.MatchString(line) {
		warnings = append(warnings, "Trailing whitespaces at the end of the line")
	}

	if startsWithSpaces.MatchString(line) && !fourSpacesIndented.MatchString(line) {
		warnings = append(warnings, "Not indented using 4 spaces")
	}

	if strings.Contains(line, "(") && !openingParenthese.MatchString(line) {
		warnings = append(warnings, "Opening parenthese not preceded by a whitespace")
	}

	return warnings
}
