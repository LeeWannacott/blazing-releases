package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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
	// ## 1.0.0 (2021-12-13)
	r, err := regexp.Compile(`## \d+\.\d+\.\d+`)

	if (err) != nil {
		log.Fatal(err)
	}

	lineCount := 0
	versionText := ""
	var versionLineNumbers []int
	for scanner.Scan() {
		lineCount++
		if r.MatchString(scanner.Text()) {
			versionText += fmt.Sprintln(lineCount, scanner.Text())
			versionLineNumbers = append(versionLineNumbers, lineCount)
		}
	}
	fmt.Print(versionText)
	fmt.Print("Version line numbers:", versionLineNumbers)

	// https://docs.github.com/en/rest/reference/releases
	resp, err := http.Get("https://api.github.com/repos/leewannacott/table-sort-js/releases")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	log.Printf(sb)
}
