package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func makeReadme(filename string) error {
	react:= "<code><img height=\"20\" src=\"https://raw.githubusercontent.com/github/explore/80688e429a7d4ef2fca1e82350fe8e3517d3494d/topics/react/react.png\"></code>"
	javascript:= "<code><img height=\"20\" src=\"https://raw.githubusercontent.com/github/explore/80688e429a7d4ef2fca1e82350fe8e3517d3494d/topics/javascript/javascript.png\"></code>"
	typescript:= "<code><img height=\"20\" src=\"https://raw.githubusercontent.com/github/explore/80688e429a7d4ef2fca1e82350fe8e3517d3494d/topics/typescript/typescript.png\"></code>"
	node:= "<code><img height=\"20\" src=\"https://raw.githubusercontent.com/github/explore/80688e429a7d4ef2fca1e82350fe8e3517d3494d/topics/nodejs/nodejs.png\"></code>"
	graphql:= "<code><img height=\"20\" src=\"https://raw.githubusercontent.com/github/explore/5c058a388828bb5fde0bcafd4bc867b5bb3f26f3/topics/graphql/graphql.png\"></code>"
	python:= "<code><img height=\"20\" src=\"https://raw.githubusercontent.com/github/explore/80688e429a7d4ef2fca1e82350fe8e3517d3494d/topics/python/python.png\"></code>"
	goLang:= "<code><img height=\"20\" src=\"https://raw.githubusercontent.com/github/explore/80688e429a7d4ef2fca1e82350fe8e3517d3494d/topics/go/go.png\"></code>"
	vim:= "<code><img height=\"20\" src=\"https://raw.githubusercontent.com/github/explore/80688e429a7d4ef2fca1e82350fe8e3517d3494d/topics/vim/vim.png\"></code>"
	docker:= "<code><img height=\"20\" src=\"https://raw.githubusercontent.com/github/explore/80688e429a7d4ef2fca1e82350fe8e3517d3494d/topics/docker/docker.png\"></code>"
	kubernetes:= "<code><img height=\"20\" src=\"https://raw.githubusercontent.com/github/explore/80688e429a7d4ef2fca1e82350fe8e3517d3494d/topics/kubernetes/kubernetes.png\"></code>"
	aws:= "<code><img height=\"20\" src=\"https://raw.githubusercontent.com/github/explore/80688e429a7d4ef2fca1e82350fe8e3517d3494d/topics/aws/aws.png\"></code>"
	terraform:= "<code><img height=\"20\" src=\"https://raw.githubusercontent.com/github/explore/80688e429a7d4ef2fca1e82350fe8e3517d3494d/topics/terraform/terraform.png\"></code>"
	postgresql:= "<code><img height=\"20\" src=\"https://raw.githubusercontent.com/github/explore/80688e429a7d4ef2fca1e82350fe8e3517d3494d/topics/postgresql/postgresql.png\"></code>"
	redis:= "<code><img height=\"20\" src=\"https://raw.githubusercontent.com/github/explore/80688e429a7d4ef2fca1e82350fe8e3517d3494d/topics/redis/redis.png\"></code>"
	date := time.Now().Format("2 Jan 2006")
	intro := "### Hello 👋 Welcome to my Github Page. I am a Lead Software Engineer who enjoys building great products, scaling systems, and mentoring engineers through code. Feel free to reach me [here](https://www.linkedin.com/in/mikengu/) if you like to chat or connect 🙂"
	tools :=  "### Languages and Tools I use: \n" + react + " " + javascript + " " + typescript + " " + node + " " + graphql + " " + python + " " + goLang + " " + vim + " " + docker + " " + kubernetes + " " + aws + " " + terraform + " " + postgresql + " " + redis

	updateDate := "This was last updated on " + date + "."

	data := fmt.Sprintf("%s\n%s\n\n%s\n", intro, tools, updateDate)

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
