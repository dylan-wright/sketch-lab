package bloom

import (
	"../badhashes"
)

/**
 * The bloom filter is a datastructure that can be used to determine membership
 * of an element in a multiset. The bloom filter has the property where all
 * negatives are true negatives but false positives are possible with multisets
 * of high cardinality or if hashing algorithms with high collision rates are
 * used.
 */
type Bloom struct {
	filter []int
	hashes []badhashes.Badhash
}

func (bloom *Bloom) Initialize(hashCount int, hashWidth int) {
	bloom.filter = make([]int, 2 << uint(8 * hashWidth))

	bloom.hashes = make([]badhashes.Badhash, hashCount)

	for i := 0; i < hashCount; i++ {
		bloom.hashes[i].Seed(i, hashWidth)
	}
}

func (bloom *Bloom) Insert(input string) {
	for i := 0; i < len(bloom.hashes); i++ {
		bloom.filter[bloom.hashes[i].Sum([]byte(input))] = 1
	}
}

func (bloom Bloom) Query(query string) bool {
	isIncluded := true

	for i := 0; i < len(bloom.hashes); i++ {
		if bloom.filter[bloom.hashes[i].Sum([]byte(query))] == 0 {
			isIncluded = false
		}
	}

	return isIncluded
}
