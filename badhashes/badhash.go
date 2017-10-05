package badhashes

import (
	"errors"
	"crypto/sha256"
)

type Badhash struct {
	index int
	width int
}

func (bh *Badhash) Seed (index int, width int) error {
	if (index < 0 || 31 < index) {
		return errors.New("Index must be between 0 and 31")
	}

	bh.index = index
	bh.width = width

	return nil
}

func (bh Badhash) Sum (data []byte) int {
	var sum int

	for i := 0; i < bh.width; i++ {
		sum = sum << 8 + int(sha256.Sum256(data)[bh.index + i])
	}

	return sum
}
