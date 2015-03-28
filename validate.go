package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

var (
	startsWithSpaces   = regexp.MustCompile(`^\s+`)
	fourSpacesIndented = regexp.MustCompile(`^([ ]{4}).[^\s]+`)
	endsWithSpaces     = regexp.MustCompile(`\s+$`)
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

		err := validateLine(line)

		if err != nil {
			fmt.Printf("%s:%d\t%s\n", filename, num, err.Error())
		}
	}
}

func validateLine(line string) error {
	if startsWithSpaces.MatchString(line) && !fourSpacesIndented.MatchString(line) {
		return errors.New("Not indented using 4 spaces")
	}

	if endsWithSpaces.MatchString(line) {
		return errors.New("Trailing whitespaces at the end of the line")
	}

	if strings.Trim(line, " ") == "{" {
		return errors.New("First parenthese isn't on the end of the first line")
	}

	return nil
}
