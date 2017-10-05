/**
 * program main for sketch lab
 * dev-meeting 10/12/2017
 * dylan.wright@squarehook.com
 *
 * inspired by article <TODO>
 */
package main

import(
	// for parsing flag
	"flag"
	// for comparing strings
	"strings"
	// for killing on error
	"log"
	// for printing output
	"fmt"
	// for creating new errors if args are not valid
	"errors"
	// an implementation of the bloom filter
	"./bloom"
	// an implementation of the count-min sketch
	"./count-min"
	// an implementation of the hyperloglog sketch
	//"./hyperloglog"
	// a helper package for reading file contents into a slice of strings
	"./filereader"
)

/**
 * run in bloom mode - input some strings and query them
 */
func bloomMode (hashCount int, hashWidth int, inputFilename string, queryFilename string) {
	var bloom bloom.Bloom

	if hashCount < 1 || hashCount > 32 {
		log.Fatal(errors.New("Invalid hashCount"))
	}

	inputs := filereader.ReadToList(inputFilename)
	queries := filereader.ReadToList(queryFilename)

	bloom.Initialize(hashCount, hashWidth)

	for i := 0; i < len(inputs); i++ {
		bloom.Insert(inputs[i])
	}

	for i := 0; i < len(queries); i++ {
		fmt.Printf("%s: %t\n", queries[i], bloom.Query(queries[i]))
	}

}

/**
 * run in count-min mode - input some strings and query them
 */
func countminMode (hashCount int, hashWidth int, inputFilename string, queryFilename string) {
	var countMin countmin.CountMin

	if hashCount < 1 || hashCount > 32 {
		log.Fatal(errors.New("Invalid hashCount"))
	}

	inputs := filereader.ReadToList(inputFilename)
	queries := filereader.ReadToList(queryFilename)

	countMin.Initialize(hashCount, hashWidth)

	for i := 0; i < len(inputs); i++ {
		countMin.Insert(inputs[i])
	}

	fmt.Printf("Sizeof datastructure `%d`", countMin.Size())

	for i := 0; i < len(queries); i++ {
		countMinCount := countMin.Count(queries[i])
		trueCount := countmin.Count(queries[i], inputs)

		fmt.Printf("%s: %d true: %d\n", queries[i], countMinCount, trueCount)
	}

}

/**
 * entrypoint
 */
func main() {
	var mode, inputFilename, queryFilename string
	var hashCount, hashWidth int

	flag.StringVar(&mode, "mode", "none", "Mode to run (bloom|count-min|hyperloglog")
	flag.StringVar(&inputFilename, "inputs", "-", "Filename to read inputs from (- for stdin)")
	flag.StringVar(&queryFilename, "queries", "-", "Filename to read queries from (- for stdin)")
	flag.IntVar(&hashCount, "hash-count", 1, "Number of hashes to use")
	flag.IntVar(&hashWidth, "hash-width", 1, "Number of bytes wide hashes should be")
	flag.Parse()

	if (strings.Compare(mode, "bloom") == 0) {
		bloomMode(hashCount, hashWidth, inputFilename, queryFilename)
	} else if (strings.Compare(mode, "count-min") == 0) {
		countminMode(hashCount, hashWidth, inputFilename, queryFilename)
	} else if (strings.Compare(mode, "hyperloglog") == 0) {
		log.Fatal(errors.New("Hyperloglog not implemented"))
	} else {
		log.Fatal(errors.New("Unknown mode: `" + mode + "`"))
	}
}

