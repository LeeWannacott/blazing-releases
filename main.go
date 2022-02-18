package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
  "encoding/json"
  // "reflect"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func callGithubAPI() {
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
    Name string `json:"name"`
    ZipballURL string `json:"zipball_url"`
    TarballURL string `json:"tarball_url"`
  }
  fmt.Printf(sb)

  var tags []Tag
  json.Unmarshal([]byte(sb), &tags)
  fmt.Printf("tags : %+v", tags)
}

func getVersionLineNumbers(scanner *bufio.Scanner) ([]int, []string) {
	// ## 1.0.0 (2021-12-13)
	r, err := regexp.Compile(`## \d+\.\d+\.\d+`)
	if (err) != nil {
		log.Fatal(err)
	}
	lineCount := int(0)
	versionText := ""
	counter := 0
	var versionLineNumbers []int
	var changeLogText []string
	for scanner.Scan() {
		changeLogText = append(changeLogText, scanner.Text())
		if r.MatchString(scanner.Text()) {
			counter++
			versionText += fmt.Sprintln(lineCount, scanner.Text())
			versionLineNumbers = append(versionLineNumbers, lineCount)
		}
		lineCount++
	}
	if scanner.Err() != nil {
		fmt.Println(scanner.Err())
	}
	return versionLineNumbers, changeLogText
}

//
func main() {
	fmt.Println("quick release notes.")
	file, err := os.Open("./docs/CHANGELOG.md")
	if (err) != nil {
		log.Fatal(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	versionLineNumbers, changeLogText := getVersionLineNumbers(scanner)
	fmt.Println("Version line numbers:", versionLineNumbers)

	var releaseBody []string

	for i := range versionLineNumbers[:len(versionLineNumbers)-1] {
		releaseBodyLines := ""
  // fmt.Println("donkey kong", lineNumber, changeLogText[lineNumber])
		if i != (len(versionLineNumbers) - 1) {
			for j := versionLineNumbers[i]; j < versionLineNumbers[i+1]; j++ {
				// fmt.Println(j, changeLogText[j])
				releaseBodyLines += changeLogText[j] + "\n"
			}
		}
		releaseBody = append(releaseBody, releaseBodyLines)
		// fmt.Println("#####")
		fmt.Println(i, releaseBody[i])
	}
	// }
	callGithubAPI()
}
