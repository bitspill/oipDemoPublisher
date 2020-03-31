package main_test

import (
	"fmt"
	"testing"
)

func TestCleanupBoolSlice(t *testing.T) {
	b := []bool{false, false, false, false, false}
	fmt.Println(cleanupBoolSlice(b))
}

func cleanupBoolSlice(boolSlice []bool) ([]bool, error) {
	l := 0
	r := len(boolSlice)

	var c bool

	// binary search for rightmost confirmed tx
	for l < r {
		m := (l + r) >> 1 // calc middle
		fmt.Println(l, m, r, len(boolSlice))

		c = boolSlice[m]

		if c {
			l = m + 1
		} else {
			r = m
		}
	}

	// move unconfirmed hashes to the front of slice
	i := copy(boolSlice, boolSlice[l:])
	for j := i; j < len(boolSlice); j++ {
		// clear trailing pointers to allow GC of confirmed hashes
		boolSlice[j] = false
	}
	// resize slice to contain only unconfirmed
	boolSlice = boolSlice[:i]

	return boolSlice, nil
}
