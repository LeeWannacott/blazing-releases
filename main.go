package main

import (
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
	data, err := os.ReadFile("./docs/CHANGELOG.md")
	check(err)
	// fmt.Println(1, string(data))
	r, _ := regexp.Compile("##")

	test := fmt.Sprintf("%s\n", r.FindAllString(string(data), len(data)))
	fmt.Print(len(data), test)
	// for _, line := range data {
	// fmt.Print(string(line))
	// lineInChangeLog := string(line)
	// fmt.Printf("%s\n", r.FindString(lineInChangeLog))
	// }

	file, err := os.Open("./docs/CHANGELOG.md")

	if (err) != nil {
		log.Fatal(err)
	}

	defer file.Close()
}
