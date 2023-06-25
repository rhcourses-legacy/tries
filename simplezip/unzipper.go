package simplezip

import (
	"fmt"

	"github.com/rhcourses/tries/tries"
)

// Unzipper is a data structure that stores a compressed string and everything that is
// needed to track the progress while decompressing it.

type Unzipper struct {
	compressed string

	tw       *tries.TrieWalker
	position int

	result string
}

// NewUnzipper creates a new unzipper from a compressed string and a trie.
func NewUnzipper(compressed string, trie *tries.Trie) *Unzipper {
	return &Unzipper{
		tw:         tries.NewTrieWalker(trie),
		compressed: compressed,
	}
}

// Done returns true if the unzipper is done.
func (uz *Unzipper) Done() bool {
	return uz.position == len(uz.compressed)
}

// Result returns the decompressed string.
func (uz *Unzipper) Result() string {
	return uz.result
}

// AdvanceSingleChar advances the unzipper by one character.
// It copies the current character to the result and advances the position.
//
// This function does not query the trie.
// It is meant to be used while unzipping a string when the trie walker has not moved.
func (uz *Unzipper) AdvanceSingleChar() {
	uz.result += string(uz.compressed[uz.position])
	uz.position++
}

// InsertCurrentData inserts the data of the current node into the result.
// If the current node has no data, it does nothing
// Returns true if data has been inserted.
//
// This function neither queries the trie nor advances the unzipper.
// It is meant to be used while unzipping a string after the trie walker has moved.
func (uz *Unzipper) InsertCurrentData() bool {
	data := uz.tw.Data()
	if len(data) != 0 {
		uz.result += fmt.Sprintf("%v", data[0])
		return true
	}
	return false
}

// CopyChars expects a number of characters and copies them to from the compressed string to the result.
//
// This function neither queries the trie nor advances the unzipper.
// It is meant to be used while unzipping a string when the trie walker has not moved.
func (uz *Unzipper) CopyChars(n int) {
	uz.result += uz.compressed[uz.position : uz.position+n]
}

// Run decompresses the string.
func (uz *Unzipper) Run() {
	for !uz.Done() {
		consumed := uz.tw.Walk(uz.compressed[uz.position:])
		if consumed == 0 {
			uz.AdvanceSingleChar()
			continue
		}
		if !uz.InsertCurrentData() {
			uz.CopyChars(consumed)
		}
		uz.tw.Reset()
		uz.position += consumed
	}
}
