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
	fourSpacesIndented = regexp.MustCompile(`^([ ]{4}).[^\s]+`)
	endsWithSpaces     = regexp.MustCompile(`\s+$`)
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s [inputfile]\n", os.Args[0])
		os.Exit(2)
	}
	filename := os.Args[1]

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

		if startsWithSpaces.MatchString(line) && !fourSpacesIndented.MatchString(line) {
			printError("Not indented using 4 spaces", filename, num)
		}

		if endsWithSpaces.MatchString(line) {
			printError("Trailing whitespaces at the end of the line", filename, num)
		}

		if strings.Trim(line, " ") == "{" {
			printError("First parenthese isn't on the end of the first line", filename, num)
		}
	}
}

func printError(err string, filename string, num int) {
	fmt.Printf("%s:%d\t%s\n", filename, num, err)
}
