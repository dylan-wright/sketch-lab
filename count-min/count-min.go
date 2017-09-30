package main

import (
	"errors"
	"../badhash"
	"fmt"
	"os"
	"io/ioutil"
	"log"
	"strings"
	"strconv"
)

type CountMin struct {
	matrix [][]int
	hashes []badhash.Badhash
}

func (cm *CountMin) initialize(hashCount int) {
	cm.matrix = make([][]int, 0)

	cm.hashes = make([]badhash.Badhash, hashCount)

	for i := 0; i < hashCount; i++ {
		cm.matrix = append(cm.matrix, make([]int, 256))
		cm.hashes[i].Seed(i)
	}
}

func (cm *CountMin) insert(input string) {
	for i := 0; i < len(cm.hashes); i++ {
		cm.matrix[i][cm.hashes[i].Sum([]byte(input))] += 1
	}
}

func (cm *CountMin) count(query string) int {
	min := cm.matrix[0][cm.hashes[0].Sum([]byte(query))]

	for i := 1; i < len(cm.hashes); i++ {
		count := cm.matrix[i][cm.hashes[i].Sum([]byte(query))]

		if count < min {
			min = count
		}
	}

	return min
}

func count(query string, words []string) int {
	count := 0
	for i := 0; i < len(words); i++ {
		if strings.Compare(words[i], query) == 0 {
			count += 1
		}
	}
	return count
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
	var countMin CountMin

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

	countMin.initialize(hashCount)

	for i := 0; i < len(inputs); i++ {
		countMin.insert(inputs[i])
	}

	for i := 0; i < len(queries); i++ {
		countMinCount := countMin.count(queries[i])
		trueCount := count(queries[i], inputs)

		fmt.Printf("%s: %d true: %d\n", queries[i], countMinCount, trueCount)
	}
}

