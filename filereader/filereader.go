package filereader

import (
	"bufio"
	"os"
	"log"
	"io/ioutil"
	"strings"
	"fmt"
)

func splitAndTrim(line string) []string {
	var words []string

	words = strings.Split(line, " ")

	for i := 0; i < len(words); i++ {
		words[i] = strings.Trim(words[i], "\t.:'[]{}!@#$%^&*()_=+~`<>/?|,\";/\\-")
	}

	return words
}

func ReadStdinToList() []string {
	var words []string
	var scanner *bufio.Scanner

	scanner = bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		// dont forget to explode the words
		words = append(words, splitAndTrim(scanner.Text())...)
	}

	return words
}

func ReadFileToList(filename string) []string {
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(strings.TrimSpace(string(content)), "\n")

	var allWords []string

	for i := 0; i < len(lines); i++ {
		words := splitAndTrim(lines[i])
		allWords = append(allWords, words...)
	}

	return allWords
}

func ReadToList(filename string) []string {
	var words []string

	if (strings.Compare(filename, "-") == 0) {
		fmt.Println("Reading until ^D is seen")
		words = ReadStdinToList()
	} else {
		fmt.Printf("Reading file `%s`\n", filename)
		words = ReadFileToList(filename)
	}

	fmt.Printf("Reading done, `%d` words ingested\n", len(words))

	return words
}

