package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("quick release notes.")
	file, err := os.Open("./docs/CHANGELOG.md")

	if (err) != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	r, err := regexp.Compile("^##")

	if (err) != nil {
		log.Fatal(err)
	}

	lineCount := 0
	text := ""
	for scanner.Scan() {
		lineCount++
		if r.MatchString(scanner.Text()) {
			text += fmt.Sprintln(lineCount, scanner.Text())
		}
	}
	fmt.Print(text)
}
