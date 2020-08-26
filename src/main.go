package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func makeReadme(filename string) error {
	date := time.Now().Format("2 Jan 2006")
	intro := "### Hello 👋 Welcome to my Github Page. I am just your average programmer that enjoys helping others through code. Feel free to reach me [here](https://www.linkedin.com/in/mikengu/) if you like to chat 🙂"
	updateDate := "This was last updated on " + date + "."

	data := fmt.Sprintf("%s\n%s\n", intro, updateDate)

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, data)
	if err != nil {
		return err
	}
	return file.Sync()
}

func main() {
	makeReadme("../README.md")
}
