package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	filepath "path/filepath"

	. "github.com/Charleslee522/gostack/src"
)

func run() {
	inputFile := flag.String("input", "./input.txt", "input file")
	flag.Parse()

	filename, _ := filepath.Abs(*inputFile)
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("input.Get err   #%v ", err)
	}

	tokenizer := NewTokenizer(string(content[:]))
	parser := NewParser(*tokenizer)
	elements := parser.Parse()
	executor := NewExecutor(elements)
	result := executor.RunIntoString()

	fmt.Println(result)
}

func main() {
	run()
}
