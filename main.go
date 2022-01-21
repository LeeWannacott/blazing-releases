package main

import (
	"fmt"
	"os"
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
	fmt.Print(string(data))
}
