package main

import (
	"fmt"
	"github.com/vinicius73/personal-gitmoji/cmd"
)

func main() {
	content, err := cmd.GenerateCollection("personal-gitmoji-source.json")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(content))
}
