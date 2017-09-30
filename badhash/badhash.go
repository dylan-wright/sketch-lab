package badhash

import (
	"errors"
	"crypto/sha256"
)

type Badhash struct {
	index int
}

func (bh *Badhash) Seed (index int) error {
	if (index < 0 || 31 < index) {
		return errors.New("Index must be between 0 and 31")
	}

	bh.index = index

	return nil
}

func (bh Badhash) Sum (data []byte) int {
	return int(sha256.Sum256(data)[bh.index])
}
