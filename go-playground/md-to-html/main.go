package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/bhanuprakaash/go-tour.git/go-playground/md-to-html/parser"
)

func main() {
	var (
		input        string
		output       string
		outputWriter strings.Builder
	)
	flag.StringVar(&input, "input", "example.md", "provide the markdown file")
	flag.StringVar(&output, "output", "output.html", "provide the output html file")
	flag.Parse()

	inputFile, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	parser := parser.NewParser(&outputWriter)

	for scanner.Scan() {
		line := scanner.Text()
		parser.ParseLine(line)
	}

	parser.Finish()

	if err := scanner.Err(); err != nil {
		fmt.Println("scan error: ", err)
	}

	fmt.Print(outputWriter.String())

	err = os.WriteFile(output, []byte(outputWriter.String()), 0644)
	if err != nil {
		log.Fatal(err)
	}

}
