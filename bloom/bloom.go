package main

import (
	"errors"
	"fmt"
	"os"
	"io/ioutil"
	"log"
	"strings"
	"strconv"
	"../badhash"
)

type Bloom struct {
	filter []int
	hashes []badhash.Badhash
}

func (bloom *Bloom) initialize(hashCount int) {
	bloom.filter = make([]int, 256)
	bloom.hashes= make([]badhash.Badhash, hashCount)

	for i := 0; i < hashCount; i++ {
		bloom.hashes[i].Seed(i)
	}
}

func (bloom *Bloom) insert(input string) {
	for i := 0; i < len(bloom.hashes); i++ {
		bloom.filter[bloom.hashes[i].Sum([]byte(input))] = 1
	}
}

func (bloom Bloom) query(query string) bool {
	isIncluded := true

	for i := 0; i < len(bloom.hashes); i++ {
		if bloom.filter[bloom.hashes[i].Sum([]byte(query))] == 0 {
			isIncluded = false
		}
	}

	return isIncluded
}

func readFileToList(filename string) []string {
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(strings.TrimSpace(string(content)), "\n")

	var allWords []string

	for i := 0; i < len(lines); i++ {
		words := strings.Split(lines[i], " ")
		allWords = append(allWords, words...)
	}

	for i := 0; i < len(allWords); i++ {
		allWords[i] = strings.Trim(allWords[i], "\t.:'[]{}!@#$%^&*()_=+~`<>/?|,\";/\\-")
	}

	return allWords
}

func main() {
	var bloom Bloom

	hashCount, err := strconv.Atoi(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	if hashCount < 1 || hashCount > 32 {
		log.Fatal(errors.New("Invalid hashCount"))
	}

	inputFilename := os.Args[2]
	queryFilename := os.Args[3]

	inputs := readFileToList(inputFilename)
	queries := readFileToList(queryFilename)

	fmt.Println(inputs)
	fmt.Println(queries)

	bloom.initialize(hashCount)

	for i := 0; i < len(inputs); i++ {
		bloom.insert(inputs[i])
	}

	for i := 0; i < len(queries); i++ {
		fmt.Printf("%s: %t\n", queries[i], bloom.query(queries[i]))
	}
}

