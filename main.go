package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	gobible "github.com/gobible/gobible"
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("Usage: gobible-plaintext <file>")
		os.Exit(1)
	}

	b := gobible.New(args[0])

	os.Mkdir("out", 0755)

	for _, book := range b.Books {
		file := "out/" + strings.ReplaceAll(book.Name, " ", "") + ".txt"
		string := book.Name + "\n"
		for _, chapter := range book.Chapters {
			for _, verse := range chapter.Verses {
				string += strconv.Itoa(chapter.Number) + ":" + strconv.Itoa(verse.Number) + " " + verse.Text + "\n"
			}
		}
		err := os.WriteFile(file, []byte(string), 0644)
		if err != nil {
			fmt.Println("error >", err)
		}
	}
}
