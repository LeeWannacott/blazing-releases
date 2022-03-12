package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	// "reflect"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getTagsFromApi() {
	//https: //docs.github.com/en/rest/reference/releases
	resp, err := http.Get("https://api.github.com/repos/quick-lint/quick-lint-js/tags")
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)

	type Tag struct {
		Name       string `json:"name"`
		ZipballURL string `json:"zipball_url"`
		TarballURL string `json:"tarball_url"`
	}
	// fmt.Printf(sb)

	var tags []Tag
	json.Unmarshal([]byte(sb), &tags)
	//fmt.Printf("tags : %+v", tags)
}

func getVersionLineNumbers(scanner *bufio.Scanner) ([]int, []string, int) {
	// ## 1.0.0 (2021-12-13)
	r, err := regexp.Compile(`## \d+\.\d+\.\d+`)
	if (err) != nil {
		log.Fatal(err)
	}
	lineCount := 0
	versionText := ""
	counterForChangeLogLength := 0
	var versionLineNumbers []int
	var changeLogText []string
	for scanner.Scan() {
		counterForChangeLogLength++
		changeLogText = append(changeLogText, scanner.Text())
		if r.MatchString(scanner.Text()) {
			versionText += fmt.Sprintln(lineCount, scanner.Text())
			versionLineNumbers = append(versionLineNumbers, lineCount)
		}
		lineCount++
	}
	if scanner.Err() != nil {
		fmt.Println(scanner.Err())
	}
	return versionLineNumbers, changeLogText, counterForChangeLogLength
}

func main() {
	fmt.Println("quick release notes.")
	file, err := os.Open("./docs/CHANGELOG.md")
	if (err) != nil {
		log.Fatal(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	versionLineNumbers, changeLogText, changeLogLength := getVersionLineNumbers(scanner)
	fmt.Println("Version line numbers:", versionLineNumbers)

	// Store contributors and errors from end of changelog.
	contributorsAndErrors := ""
	for i := 4 + versionLineNumbers[len(versionLineNumbers)-1]; i < changeLogLength; i++ {
		contributorsAndErrors += changeLogText[i] + "\n"
	}

	var releaseBody []string
	for i, versionLineNumber := range versionLineNumbers[:len(versionLineNumbers)] {
		releaseBodyLines := ""
		// Last version (## 0.2.0) excluded with - 1
		if i < (len(versionLineNumbers) - 1) {
			for j := versionLineNumbers[i]; j < versionLineNumbers[i+1]; j++ {
				releaseBodyLines += changeLogText[j] + "\n"
			}
		}

		// Handle last version (## 0.2.0)
		if versionLineNumber == versionLineNumbers[len(versionLineNumbers)-1] {
			for j := 0; j < 4; j++ {
				releaseBodyLines += changeLogText[versionLineNumber+j] + "\n"
			}
		}
		releaseBody = append(releaseBody, releaseBodyLines+contributorsAndErrors)
		fmt.Println(i, releaseBody[i])
	}
	getTagsFromApi()
}
