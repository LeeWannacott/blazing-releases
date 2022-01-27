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

func callGithubAPI() {
	//https: //docs.github.com/en/rest/reference/releases
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
	fmt.Print("Version line numbers:", versionLineNumbers)

	print("bingo")
	// textBody := ""
	// fmt.Print(changeLogText)
	// for _, changeLogLine := range changeLogText {
	// fmt.Println(i, changeLogLine)
	for i, lineNumber := range versionLineNumbers {
		fmt.Println("donkey kong", lineNumber, changeLogText[lineNumber])
		if i != (len(versionLineNumbers) - 1) {
			for j := versionLineNumbers[i]; j < versionLineNumbers[i+1]; j++ {
				fmt.Println(j, changeLogText[j])
			}
		}
	}
	// }
	// callGithubAPI()
}
