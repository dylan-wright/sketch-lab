package countmin

import (
	"../badhashes"
	"strings"
	"unsafe"
)

/**
 * The count-min sketch can be used to estimate the frequency of an element
 * in a multiset. For sets with a high range of possible values, such
 * as the set of valid internet domain names it would require a large amount
 * of space to store an individual counter for each element. This sketch
 * allows us to get a pretty good estimate of the frequency of elements in
 * such a set in a datastructure that can be easily stored in a database
 *
 * this datastructure uses space O(nh) where n is the number of hashing
 * algorithms and h is the range of those algorithms' output
 */
type CountMin struct {
	matrix [][]int
	hashes []badhashes.Badhash
}

func (cm *CountMin) Size() int {
	size := int(unsafe.Sizeof(*cm))
	for i := 0; i < len(cm.matrix); i++ {
		size += 4 * len(cm.matrix[i])
	}
	return size
}

func (cm *CountMin) Initialize(hashCount int, hashWidth int) {
	cm.matrix = make([][]int, 0)

	cm.hashes = make([]badhashes.Badhash, hashCount)

	for i := 0; i < hashCount; i++ {
		cm.matrix = append(cm.matrix, make([]int, 2 << uint(8 * hashWidth)))
		cm.hashes[i].Seed(i, hashWidth)
	}
}

func (cm *CountMin) Insert(input string) {
	for i := 0; i < len(cm.hashes); i++ {
		cm.matrix[i][cm.hashes[i].Sum([]byte(input))] += 1
	}
}

func (cm *CountMin) Count(query string) int {
	min := cm.matrix[0][cm.hashes[0].Sum([]byte(query))]

	for i := 1; i < len(cm.hashes); i++ {
		count := cm.matrix[i][cm.hashes[i].Sum([]byte(query))]

		if count < min {
			min = count
		}
	}

	return min
}

/**
 * helper function for determining the exact frequency of a string in the
 * input multiset. This can be used to determine error in the count min
 * datastructure
 */
func Count(query string, words []string) int {
	count := 0
	for i := 0; i < len(words); i++ {
		if strings.Compare(words[i], query) == 0 {
			count += 1
		}
	}
	return count
}
