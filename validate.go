package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

type styleWarning struct {
	code    int
	message string
}

var (
	singleBrace        = regexp.MustCompile(`^\s*\{\s*$`)
	endsWithSpaces     = regexp.MustCompile(`\s+$`)
	startsWithSpaces   = regexp.MustCompile(`^\s+`)
	fourSpacesIndented = regexp.MustCompile(`^( {4})*[^\s]+`)
	openingParenthese  = regexp.MustCompile(`[^_|\(| ]\(`)
	equalWithSpaces    = regexp.MustCompile(`[^ |!|<|>|=]=|=[^=| |>]`)
	moreThan120Chars   = regexp.MustCompile(`.{121}`)
	glibNotNecessary   = regexp.MustCompile(`GLib.print`)

	warningMessages = map[int]string{
		1: "First brace isn't on the end of the first line",
		2: "Trailing whitespaces at the end of the line",
		3: "Not indented using 4 spaces",
		4: "Opening parenthese not preceeded by a whitespace",
		5: "Equals sign not surrounded by whitespaces",
		6: "The length of the line is superior to 120 characters",
		7: "Referring to GLib is not necessary",
	}
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
				fmt.Printf("%s:%d\t%s\n", filename, num, warningMessages[warning])
			}
		}
	}
}

func validateLine(line string) []int {
	var warnings []int
	if singleBrace.MatchString(line) {
		warnings = append(warnings, 1)
	}

	if endsWithSpaces.MatchString(line) {
		warnings = append(warnings, 2)
	}

	if startsWithSpaces.MatchString(line) && !fourSpacesIndented.MatchString(line) {
		warnings = append(warnings, 3)
	}

	if openingParenthese.MatchString(line) {
		warnings = append(warnings, 4)
	}

	if equalWithSpaces.MatchString(line) {
		warnings = append(warnings, 5)
	}

	if moreThan120Chars.MatchString(line) {
		warnings = append(warnings, 6)
	}

	if glibNotNecessary.MatchString(line) {
		warnings = append(warnings, 7)
	}

	return warnings
}
